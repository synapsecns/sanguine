// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/memory/Attestation.sol";
import {
    IncorrectAttestation, IncorrectState, IndexOutOfRange, NonceOutOfRange, OutdatedNonce
} from "../libs/Errors.sol";
import {ChainGas, GasData, GasDataLib} from "../libs/stack/GasData.sol";
import {MerkleMath} from "../libs/merkle/MerkleMath.sol";
import {Snapshot, SnapshotLib} from "../libs/memory/Snapshot.sol";
import {State, StateLib} from "../libs/memory/State.sol";
import {ChainContext} from "../libs/ChainContext.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "../base/AgentSecured.sol";
import {SnapshotHubEvents} from "../events/SnapshotHubEvents.sol";
import {ISnapshotHub} from "../interfaces/ISnapshotHub.sol";
import {IStatementInbox} from "../interfaces/IStatementInbox.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// @notice `SnapshotHub` is a parent contract for `Summit`. It is responsible for the following:
/// - Accepting and storing Guard and Notary snapshots to keep track of all the remote `Origin` states.
/// - Generating and storing Attestations derived from Notary snapshots, as well as verifying their validity.
abstract contract SnapshotHub is AgentSecured, SnapshotHubEvents, ISnapshotHub {
    using AttestationLib for bytes;
    using SafeCast for uint256;
    using StateLib for bytes;

    /// @notice Struct that represents stored State of Origin contract
    /// @param guardIndex   Index of Guard who submitted this State to Summit
    /// @param notaryIndex  Index of Notary who submitted this State to Summit
    struct SummitState {
        bytes32 root;
        uint32 origin;
        uint32 nonce;
        uint40 blockNumber;
        uint40 timestamp;
        GasData gasData;
        uint32 guardIndex;
        uint32 notaryIndex;
    }
    // TODO: revisit packing

    struct SummitSnapshot {
        // TODO: compress this - indexes might as well be uint32/uint64
        uint256[] statePtrs;
        uint256 sigIndex;
    }

    struct SummitAttestation {
        bytes32 snapRoot;
        bytes32 agentRoot;
        bytes32 snapGasHash;
        uint40 blockNumber;
        uint40 timestamp;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev All States submitted by any of the Guards
    SummitState[] private _states;

    /// @dev All Snapshots submitted by any of the Guards
    SummitSnapshot[] private _guardSnapshots;

    /// @dev All Snapshots submitted by any of the Notaries
    SummitSnapshot[] private _notarySnapshots;

    /// @dev All Attestations created from Notary-submitted Snapshots
    /// Invariant: _attestations.length == _notarySnapshots.length
    SummitAttestation[] private _attestations;

    /// @dev Pointer for the given State Leaf of the origin
    /// with ZERO as a sentinel value for "state not submitted yet".
    // (origin => (stateLeaf => {state index in _states PLUS 1}))
    mapping(uint32 => mapping(bytes32 => uint256)) private _leafPtr;

    /// @dev Pointer for the latest Agent State of a given origin
    /// with ZERO as a sentinel value for "no states submitted yet".
    // (origin => (agent index => {latest state index in _states PLUS 1}))
    mapping(uint32 => mapping(uint32 => uint256)) private _latestStatePtr;

    /// @dev Latest nonce that a Notary created
    // (notary index => latest nonce)
    mapping(uint32 => uint32) private _latestAttNonce;

    /// @dev gap for upgrade safety
    uint256[43] private __GAP; // solhint-disable-line var-name-mixedcase

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISnapshotHub
    function isValidAttestation(bytes memory attPayload) external view returns (bool isValid) {
        // This will revert if payload is not a formatted attestation
        Attestation attestation = attPayload.castToAttestation();
        return _isValidAttestation(attestation);
    }

    /// @inheritdoc ISnapshotHub
    function getAttestation(uint32 attNonce)
        external
        view
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas)
    {
        if (attNonce >= _attestations.length) revert NonceOutOfRange();
        SummitAttestation memory summitAtt = _attestations[attNonce];
        attPayload = _formatSummitAttestation(summitAtt, attNonce);
        agentRoot = summitAtt.agentRoot;
        snapGas = _restoreSnapGas(_notarySnapshots[attNonce]);
    }

    /// @inheritdoc ISnapshotHub
    function getLatestAgentState(uint32 origin, address agent) external view returns (bytes memory stateData) {
        SummitState memory latestState = _latestState(origin, _agentStatus(agent).index);
        if (latestState.nonce == 0) return bytes("");
        return _formatSummitState(latestState);
    }

    /// @inheritdoc ISnapshotHub
    function getLatestNotaryAttestation(address notary)
        external
        view
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas)
    {
        uint32 latestAttNonce = _latestAttNonce[_agentStatus(notary).index];
        if (latestAttNonce != 0) {
            SummitAttestation memory summitAtt = _attestations[latestAttNonce];
            attPayload = _formatSummitAttestation(summitAtt, latestAttNonce);
            agentRoot = summitAtt.agentRoot;
            snapGas = _restoreSnapGas(_notarySnapshots[latestAttNonce]);
        }
    }

    /// @inheritdoc ISnapshotHub
    function getGuardSnapshot(uint256 index)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {
        if (index >= _guardSnapshots.length) revert IndexOutOfRange();
        return _restoreSnapshot(_guardSnapshots[index]);
    }

    /// @inheritdoc ISnapshotHub
    function getNotarySnapshot(uint256 index)
        public
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {
        uint256 nonce = index + 1;
        if (nonce >= _notarySnapshots.length) revert IndexOutOfRange();
        return _restoreSnapshot(_notarySnapshots[nonce]);
    }

    /// @inheritdoc ISnapshotHub
    // solhint-disable-next-line ordering
    function getNotarySnapshot(bytes memory attPayload)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {
        // This will revert if payload is not a formatted attestation
        Attestation attestation = attPayload.castToAttestation();
        if (!_isValidAttestation(attestation)) revert IncorrectAttestation();
        // Attestation is valid => _attestations[nonce] exists
        // _notarySnapshots.length == _attestations.length => _notarySnapshots[nonce] exists
        return _restoreSnapshot(_notarySnapshots[attestation.nonce()]);
    }

    /// @inheritdoc ISnapshotHub
    function getSnapshotProof(uint32 attNonce, uint8 stateIndex) external view returns (bytes32[] memory snapProof) {
        if (attNonce == 0 || attNonce >= _notarySnapshots.length) revert NonceOutOfRange();
        SummitSnapshot memory snap = _notarySnapshots[attNonce];
        uint256 statesAmount = snap.statePtrs.length;
        if (stateIndex >= statesAmount) revert IndexOutOfRange();
        // Reconstruct the leafs of Snapshot Merkle Tree: two for each state
        bytes32[] memory hashes = new bytes32[](2 * statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            // Get value for "index in _states PLUS 1"
            uint256 statePtr = snap.statePtrs[i];
            // We are never saving zero values when accepting Guard/Notary snapshots, so this holds
            assert(statePtr != 0);
            State state = _formatSummitState(_states[statePtr - 1]).castToState();
            (hashes[2 * i], hashes[2 * i + 1]) = state.subLeafs();
        }
        // Index of State's left leaf is twice the state index
        return MerkleMath.calculateProof(hashes, 2 * stateIndex);
    }

    // ════════════════════════════════════════ INTERNAL LOGIC: ACCEPT DATA ════════════════════════════════════════════

    /// @dev Accepts a Snapshot signed by a Guard.
    /// It is assumed that the Guard signature has been checked outside of this contract.
    function _acceptGuardSnapshot(Snapshot snapshot, uint32 guardIndex, uint256 sigIndex) internal {
        // Snapshot Signer is a Guard: save the states for later use.
        uint256 statesAmount = snapshot.statesAmount();
        uint256[] memory statePtrs = new uint256[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            statePtrs[i] = _saveState(snapshot.state(i), guardIndex);
            // Guard either submitted a fresh state, or reused state submitted by another Guard
            // In any case, the "state pointer" would never be zero
            assert(statePtrs[i] != 0);
        }
        // Save Guard snapshot for later retrieval
        _saveGuardSnapshot(statePtrs, sigIndex);
    }

    /// @dev Accepts a Snapshot signed by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    /// Returns the attestation created from the Notary snapshot.
    function _acceptNotarySnapshot(Snapshot snapshot, bytes32 agentRoot, uint32 notaryIndex, uint256 sigIndex)
        internal
        returns (bytes memory attPayload)
    {
        // Snapshot Signer is a Notary: construct a Snapshot Merkle Tree,
        // while checking that the states were previously saved.
        uint256 statesAmount = snapshot.statesAmount();
        uint256[] memory statePtrs = new uint256[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            State state = snapshot.state(i);
            uint256 statePtr = _statePtr(state);
            // Notary can only used states previously submitted by any fo the Guards
            if (statePtr == 0) revert IncorrectState();
            statePtrs[i] = statePtr;
            // Check that Notary hasn't used a fresher state for this origin before
            uint32 origin = state.origin();
            if (state.nonce() <= _latestState(origin, notaryIndex).nonce) revert OutdatedNonce();
            // Save Notary if they are the first to use this state
            if (_states[statePtr - 1].notaryIndex == 0) _states[statePtr - 1].notaryIndex = notaryIndex;
            // Update Notary latest state for origin
            _latestStatePtr[origin][notaryIndex] = statePtrs[i];
        }
        // Derive the snapshot merkle root and save it for a Notary attestation.
        // Save Notary snapshot for later retrieval
        return _saveNotarySnapshot(snapshot, statePtrs, agentRoot, notaryIndex, sigIndex);
    }

    // ════════════════════════════════════ INTERNAL LOGIC: SAVE STATEMENT DATA ════════════════════════════════════════

    /// @dev Initializes the saved _attestations list by inserting empty values.
    function _initializeAttestations() internal {
        // This should only be called once, when the contract is initialized
        assert(_attestations.length == 0);
        // Insert empty non-meaningful values, that can't be used to prove anything
        _attestations.push(_toSummitAttestation(0, 0, 0));
        _notarySnapshots.push(SummitSnapshot(new uint256[](0), 0));
    }

    /// @dev Saves the Guard snapshot.
    function _saveGuardSnapshot(uint256[] memory statePtrs, uint256 sigIndex) internal {
        _guardSnapshots.push(SummitSnapshot(statePtrs, sigIndex));
    }

    /// @dev Saves the Notary snapshot and the attestation created from it.
    /// Returns the created attestation.
    function _saveNotarySnapshot(
        Snapshot snapshot,
        uint256[] memory statePtrs,
        bytes32 agentRoot,
        uint32 notaryIndex,
        uint256 sigIndex
    ) internal returns (bytes memory attPayload) {
        // Attestation nonce is its index in `_attestations` array
        // TODO: consider using more than 32 bits for attestation nonces
        uint32 attNonce = _attestations.length.toUint32();
        bytes32 snapGasHash = GasDataLib.snapGasHash(snapshot.snapGas());
        SummitAttestation memory summitAtt = _toSummitAttestation(snapshot.calculateRoot(), agentRoot, snapGasHash);
        attPayload = _formatSummitAttestation(summitAtt, attNonce);
        _latestAttNonce[notaryIndex] = attNonce;
        /// @dev Add a single element to both `_attestations` and `_notarySnapshots`,
        /// enforcing the (_attestations.length == _notarySnapshots.length) invariant.
        _attestations.push(summitAtt);
        _notarySnapshots.push(SummitSnapshot(statePtrs, sigIndex));
        // Emit event with raw attestation data
        emit AttestationSaved(attPayload);
    }

    /// @dev Saves the state signed by a Guard.
    function _saveState(State state, uint32 guardIndex) internal returns (uint256 statePtr) {
        uint32 origin = state.origin();
        // Check that Guard hasn't submitted a fresher State before
        if (state.nonce() <= _latestState(origin, guardIndex).nonce) revert OutdatedNonce();
        bytes32 stateHash = state.leaf();
        statePtr = _leafPtr[origin][stateHash];
        // Save state only if it wasn't previously submitted
        if (statePtr == 0) {
            // Extract data that needs to be saved
            SummitState memory summitState = _toSummitState(state, guardIndex);
            _states.push(summitState);
            // State is stored at (length - 1), but we are tracking "index PLUS 1" as "pointer"
            statePtr = _states.length;
            _leafPtr[origin][stateHash] = statePtr;
            // Emit event with raw state data
            emit StateSaved(state.unwrap().clone());
        }
        // Update latest guard state for origin
        _latestStatePtr[origin][guardIndex] = statePtr;
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Checks if attestation was previously submitted by a Notary (as a signed snapshot).
    function _isValidAttestation(Attestation att) internal view returns (bool) {
        // Check if nonce exists
        uint32 nonce = att.nonce();
        if (nonce >= _attestations.length) return false;
        // Check if Attestation matches the historical one
        return _areEqual(att, _attestations[nonce]);
    }

    /// @dev Restores Snapshot payload from a list of state pointers used for the snapshot.
    function _restoreSnapshot(SummitSnapshot memory snapshot)
        internal
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {
        uint256 statesAmount = snapshot.statePtrs.length;
        State[] memory states = new State[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            // Get value for "index in _states PLUS 1"
            uint256 statePtr = snapshot.statePtrs[i];
            // We are never saving zero values when accepting Guard/Notary snapshots, so this holds
            assert(statePtr != 0);
            // Get the state that Agent used for the snapshot
            states[i] = _formatSummitState(_states[statePtr - 1]).castToState();
        }
        snapPayload = SnapshotLib.formatSnapshot(states);
        snapSignature = IStatementInbox(inbox).getStoredSignature(snapshot.sigIndex);
    }

    /// @dev Restores the gas data from the snapshot.
    function _restoreSnapGas(SummitSnapshot memory snapshot) internal view returns (uint256[] memory snapGas) {
        uint256 statesAmount = snapshot.statePtrs.length;
        snapGas = new uint256[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            // Get value for "index in _states PLUS 1"
            uint256 statePtr = snapshot.statePtrs[i];
            // We are never saving zero values when accepting Guard/Notary snapshots, so this holds
            assert(statePtr != 0);
            // Get the state that Agent used for the snapshot
            snapGas[i] = ChainGas.unwrap(
                GasDataLib.encodeChainGas({
                    gasData_: _states[statePtr - 1].gasData,
                    domain_: _states[statePtr - 1].origin
                })
            );
        }
    }

    /// @dev Returns indexes of agents who provided state data for the Notary snapshot with the given nonce.
    function _stateAgents(uint32 nonce, uint8 stateIndex)
        internal
        view
        returns (uint32 guardIndex, uint32 notaryIndex)
    {
        uint256 statePtr = _notarySnapshots[nonce].statePtrs[stateIndex];
        return (_states[statePtr - 1].guardIndex, _states[statePtr - 1].notaryIndex);
    }

    /// @dev Returns the pointer for a matching Guard State, if it exists.
    function _statePtr(State state) internal view returns (uint256) {
        return _leafPtr[state.origin()][state.leaf()];
    }

    /// @dev Returns the latest state submitted by the Agent for the origin.
    /// Will return an empty struct, if the Agent hasn't submitted a single origin State yet.
    function _latestState(uint32 origin, uint32 agentIndex) internal view returns (SummitState memory state) {
        // Get value for "index in _states PLUS 1"
        uint256 latestPtr = _latestStatePtr[origin][agentIndex];
        // Check if the Agent has submitted at least one State for origin
        if (latestPtr != 0) {
            state = _states[latestPtr - 1];
        }
        // An empty struct is returned if the Agent hasn't submitted a single State for origin yet.
    }

    // ═════════════════════════════════════════════ STRUCT FORMATTING ═════════════════════════════════════════════════

    /// @dev Returns a formatted payload for a stored SummitState.
    function _formatSummitState(SummitState memory summitState) internal pure returns (bytes memory) {
        return StateLib.formatState({
            root_: summitState.root,
            origin_: summitState.origin,
            nonce_: summitState.nonce,
            blockNumber_: summitState.blockNumber,
            timestamp_: summitState.timestamp,
            gasData_: summitState.gasData
        });
    }

    /// @dev Returns a SummitState struct to save in the contract.
    function _toSummitState(State state, uint32 guardIndex) internal pure returns (SummitState memory summitState) {
        summitState.root = state.root();
        summitState.origin = state.origin();
        summitState.nonce = state.nonce();
        summitState.blockNumber = state.blockNumber();
        summitState.timestamp = state.timestamp();
        summitState.gasData = state.gasData();
        summitState.guardIndex = guardIndex;
        // summitState.notaryIndex is left as ZERO
    }

    /// @dev Returns a formatted payload for a stored SummitAttestation.
    function _formatSummitAttestation(SummitAttestation memory summitAtt, uint32 nonce)
        internal
        pure
        returns (bytes memory)
    {
        return AttestationLib.formatAttestation({
            snapRoot_: summitAtt.snapRoot,
            dataHash_: AttestationLib.dataHash(summitAtt.agentRoot, summitAtt.snapGasHash),
            nonce_: nonce,
            blockNumber_: summitAtt.blockNumber,
            timestamp_: summitAtt.timestamp
        });
    }

    /// @dev Returns an Attestation struct to save in the Summit contract.
    /// Current block number and timestamp are used.
    // solhint-disable-next-line ordering
    function _toSummitAttestation(bytes32 snapRoot, bytes32 agentRoot, bytes32 snapGasHash)
        internal
        view
        returns (SummitAttestation memory summitAtt)
    {
        summitAtt.snapRoot = snapRoot;
        summitAtt.agentRoot = agentRoot;
        summitAtt.snapGasHash = snapGasHash;
        summitAtt.blockNumber = ChainContext.blockNumber();
        summitAtt.timestamp = ChainContext.blockTimestamp();
    }

    /// @dev Checks that an Attestation and its Summit representation are equal.
    function _areEqual(Attestation att, SummitAttestation memory summitAtt) internal pure returns (bool) {
        // forgefmt: disable-next-item
        return 
            att.snapRoot() == summitAtt.snapRoot &&
            att.dataHash() == AttestationLib.dataHash(summitAtt.agentRoot, summitAtt.snapGasHash) &&
            att.blockNumber() == summitAtt.blockNumber &&
            att.timestamp() == summitAtt.timestamp;
    }
}
