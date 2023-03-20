// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation, AttestationLib, SummitAttestation } from "../libs/Attestation.sol";
import { MerkleList } from "../libs/MerkleList.sol";
import { Snapshot, SnapshotLib, SummitSnapshot } from "../libs/Snapshot.sol";
import { State, StateLib, SummitState } from "../libs/State.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { SnapshotHubEvents } from "../events/SnapshotHubEvents.sol";
import { ISnapshotHub } from "../interfaces/ISnapshotHub.sol";

/**
 * @notice Hub to accept and save snapshots, as well as verify attestations.
 */
abstract contract SnapshotHub is SnapshotHubEvents, ISnapshotHub {
    using AttestationLib for bytes;
    using SnapshotLib for uint256[];
    using StateLib for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev All States submitted by any of the Guards
    SummitState[] private guardStates;

    /// @dev All Snapshots submitted by any of the Guards
    SummitSnapshot[] private guardSnapshots;

    /// @dev All Snapshots submitted by any of the Notaries
    SummitSnapshot[] private notarySnapshots;

    /// @dev All Attestations created from Notary-submitted Snapshots
    /// Invariant: attestations.length == notarySnapshots.length
    SummitAttestation[] private attestations;

    /// @dev Pointer for the given State Leaf of the origin
    /// with ZERO as a sentinel value for "state not submitted yet".
    // (origin => (stateLeaf => {state index in guardStates PLUS 1}))
    mapping(uint32 => mapping(bytes32 => uint256)) private leafPtr;

    /// @dev Pointer for the latest Agent State of a given origin
    /// with ZERO as a sentinel value for "no states submitted yet".
    // (origin => (agent => {latest state index in guardStates PLUS 1}))
    mapping(uint32 => mapping(address => uint256)) private latestStatePtr;

    /// @dev gap for upgrade safety
    uint256[44] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ISnapshotHub
    function isValidAttestation(bytes memory _attPayload) external view returns (bool isValid) {
        // This will revert if payload is not a formatted attestation
        Attestation attestation = _attPayload.castToAttestation();
        return _isValidAttestation(attestation);
    }

    /// @inheritdoc ISnapshotHub
    function getAttestation(uint32 _nonce) external view returns (bytes memory attPayload) {
        require(_nonce < attestations.length, "Nonce out of range");
        return attestations[_nonce].formatSummitAttestation(_nonce);
    }

    /// @inheritdoc ISnapshotHub
    function getLatestAgentState(uint32 _origin, address _agent)
        external
        view
        returns (bytes memory stateData)
    {
        SummitState memory latestState = _latestState(_origin, _agent);
        if (latestState.nonce == 0) return bytes("");
        return latestState.formatSummitState();
    }

    /// @inheritdoc ISnapshotHub
    function getGuardSnapshot(uint256 _index) external view returns (bytes memory snapshotPayload) {
        require(_index < guardSnapshots.length, "Index out of range");
        return _restoreSnapshot(guardSnapshots[_index]);
    }

    /// @inheritdoc ISnapshotHub
    function getNotarySnapshot(uint256 _nonce)
        external
        view
        returns (bytes memory snapshotPayload)
    {
        require(_nonce < notarySnapshots.length, "Nonce out of range");
        return _restoreSnapshot(notarySnapshots[_nonce]);
    }

    /// @inheritdoc ISnapshotHub
    function getNotarySnapshot(bytes memory _attPayload)
        external
        view
        returns (bytes memory snapshotPayload)
    {
        // This will revert if payload is not a formatted attestation
        Attestation attestation = _attPayload.castToAttestation();
        require(_isValidAttestation(attestation), "Invalid attestation");
        // Attestation is valid => attestations[nonce] exists
        // notarySnapshots.length == attestations.length => notarySnapshots[nonce] exists
        return _restoreSnapshot(notarySnapshots[attestation.nonce()]);
    }

    /// @inheritdoc ISnapshotHub
    function getSnapshotProof(uint256 _nonce, uint256 _stateIndex)
        external
        view
        returns (bytes32[] memory snapProof)
    {
        require(_nonce < notarySnapshots.length, "Nonce out of range");
        snapProof = new bytes32[](attestations[_nonce].height);
        SummitSnapshot memory snap = notarySnapshots[_nonce];
        uint256 statesAmount = snap.getStatesAmount();
        require(_stateIndex < statesAmount, "Index out of range");
        // Reconstruct the leafs of Snapshot Merkle Tree
        bytes32[] memory hashes = new bytes32[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            // Get value for "index in guardStates PLUS 1"
            uint256 statePtr = snap.getStatePtr(i);
            // We are never saving zero values when accepting Guard/Notary snapshots, so this holds
            assert(statePtr != 0);
            SummitState memory guardState = guardStates[statePtr - 1];
            State state = guardState.formatSummitState().castToState();
            hashes[i] = state.leaf();
            if (i == _stateIndex) {
                // First element of the proof is "right sub-leaf"
                (, snapProof[0]) = state.subLeafs();
            }
        }
        // This will fill the remaining values in the `snapProof` array
        MerkleList.calculateProof(hashes, _stateIndex, snapProof);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ACCEPT DATA                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Accepts a Snapshot signed by a Guard.
    /// It is assumed that the Guard signature has been checked outside of this contract.
    function _acceptGuardSnapshot(Snapshot _snapshot, address _guard) internal {
        // Snapshot Signer is a Guard: save the states for later use.
        uint256 statesAmount = _snapshot.statesAmount();
        uint256[] memory statePtrs = new uint256[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            statePtrs[i] = _saveState(_snapshot.state(i), _guard);
            // Guard either submitted a fresh state, or reused state submitted by another Guard
            // In any case, the "state pointer" would never be zero
            assert(statePtrs[i] != 0);
        }
        // Save Guard snapshot for later retrieval
        _saveGuardSnapshot(statePtrs);
    }

    /// @dev Accepts a Snapshot signed by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    /// Returns the attestation created from the Notary snapshot.
    function _acceptNotarySnapshot(Snapshot _snapshot, address _notary)
        internal
        returns (bytes memory attPayload)
    {
        // Snapshot Signer is a Notary: construct an Attestation Merkle Tree,
        // while checking that the states were previously saved.
        uint256 statesAmount = _snapshot.statesAmount();
        uint256[] memory statePtrs = new uint256[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            State state = _snapshot.state(i);
            uint256 statePtr = _statePtr(state);
            // Notary can only used states previously submitted by any fo the Guards
            require(statePtr != 0, "State doesn't exist");
            statePtrs[i] = statePtr;
            // Check that Notary hasn't used a fresher state for this origin before
            uint32 origin = state.origin();
            require(state.nonce() > _latestState(origin, _notary).nonce, "Outdated nonce");
            // Update Notary latest state for origin
            latestStatePtr[origin][_notary] = statePtrs[i];
        }
        // Derive attestation merkle root and save it for a Notary attestation.
        // Save Notary snapshot for later retrieval
        return _saveNotarySnapshot(_snapshot, statePtrs);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         SAVE STATEMENT DATA                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Initializes the saved attestations list by inserting empty values.
    function _initializeAttestations() internal {
        // This should only be called once, when the contract is initialized
        assert(attestations.length == 0);
        // Insert empty non-meaningful values, that can't be used to prove anything
        attestations.push(AttestationLib.emptySummitAttestation());
        notarySnapshots.push(SnapshotLib.emptySummitSnapshot());
    }

    /// @dev Saves the Guard snapshot.
    function _saveGuardSnapshot(uint256[] memory statePtrs) internal {
        guardSnapshots.push(statePtrs.toSummitSnapshot());
    }

    /// @dev Saves the Notary snapshot and the attestation created from it.
    /// Returns the created attestation.
    function _saveNotarySnapshot(Snapshot _snapshot, uint256[] memory statePtrs)
        internal
        returns (bytes memory attPayload)
    {
        // Attestation nonce is its index in `attestations` array
        uint32 attNonce = uint32(attestations.length);
        SummitAttestation memory summitAtt = _snapshot.toSummitAttestation();
        attPayload = summitAtt.formatSummitAttestation(attNonce);
        /// @dev Add a single element to both `attestations` and `notarySnapshots`,
        /// enforcing the (attestations.length == notarySnapshots.length) invariant.
        attestations.push(summitAtt);
        notarySnapshots.push(statePtrs.toSummitSnapshot());
        // Emit event with raw attestation data
        emit AttestationSaved(attPayload);
    }

    /// @dev Saves the state signed by a Guard.
    function _saveState(State _state, address _guard) internal returns (uint256 statePtr) {
        uint32 origin = _state.origin();
        // Check that Guard hasn't submitted a fresher State before
        require(_state.nonce() > _latestState(origin, _guard).nonce, "Outdated nonce");
        bytes32 stateHash = _state.leaf();
        statePtr = leafPtr[origin][stateHash];
        // Save state only if it wasn't previously submitted
        if (statePtr == 0) {
            // Extract data that needs to be saved
            SummitState memory state = _state.toSummitState();
            guardStates.push(state);
            // State is stored at (length - 1), but we are tracking "index PLUS 1" as "pointer"
            statePtr = guardStates.length;
            leafPtr[origin][stateHash] = statePtr;
            // Emit event with raw state data
            emit StateSaved(_state.unwrap().clone());
        }
        // Update latest guard state for origin
        latestStatePtr[origin][_guard] = statePtr;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CHECK STATEMENT DATA                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns the amount of saved attestations (created from Notary snapshots) so far.
    function _attestationsAmount() internal view returns (uint256) {
        return attestations.length;
    }

    /// @dev Checks if attestation was previously submitted by a Notary (as a signed snapshot).
    function _isValidAttestation(Attestation _att) internal view returns (bool) {
        // Check if nonce exists
        uint32 nonce = _att.nonce();
        if (nonce >= attestations.length) return false;
        // Check if Attestation matches the historical one
        return _att.equalToSummit(attestations[nonce]);
    }

    /// @dev Restores Snapshot payload from a list of state pointers used for the snapshot.
    function _restoreSnapshot(SummitSnapshot memory _snapshot)
        internal
        view
        returns (bytes memory)
    {
        uint256 statesAmount = _snapshot.getStatesAmount();
        State[] memory states = new State[](statesAmount);
        for (uint256 i = 0; i < statesAmount; ++i) {
            // Get value for "index in guardStates PLUS 1"
            uint256 statePtr = _snapshot.getStatePtr(i);
            // We are never saving zero values when accepting Guard/Notary snapshots, so this holds
            assert(statePtr != 0);
            SummitState memory state = guardStates[statePtr - 1];
            // Get the state that Agent used for the snapshot
            states[i] = state.formatSummitState().castToState();
        }
        return SnapshotLib.formatSnapshot(states);
    }

    /// @dev Returns the pointer for a matching Guard State, if it exists.
    function _statePtr(State _state) internal view returns (uint256) {
        return leafPtr[_state.origin()][_state.leaf()];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          LATEST STATE VIEWS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns the latest state submitted by the Agent for the origin.
    /// Will return an empty struct, if the Agent hasn't submitted a single origin State yet.
    function _latestState(uint32 _origin, address _agent)
        internal
        view
        returns (SummitState memory state)
    {
        // Get value for "index in guardStates PLUS 1"
        uint256 latestPtr = latestStatePtr[_origin][_agent];
        // Check if the Agent has submitted at least one State for origin
        if (latestPtr != 0) {
            state = guardStates[latestPtr - 1];
        }
        // An empty struct is returned if the Agent hasn't submitted a single State for origin yet.
    }
}
