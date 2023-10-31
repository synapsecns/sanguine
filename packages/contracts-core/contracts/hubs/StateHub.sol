// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {IncorrectOriginDomain} from "../libs/Errors.sol";
import {GasData, GasDataLib} from "../libs/stack/GasData.sol";
import {HistoricalTree} from "../libs/merkle/MerkleTree.sol";
import {State, StateLib} from "../libs/memory/State.sol";
import {ChainContext} from "../libs/ChainContext.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "../base/AgentSecured.sol";
import {StateHubEvents} from "../events/StateHubEvents.sol";
import {IStateHub} from "../interfaces/IStateHub.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// @notice `StateHub` is a parent contract for `Origin`. It is responsible for the following:
/// - Keeping track of the historical Origin Merkle Tree containing all the message hashes.
/// - Keeping track of the historical Origin States, as well as verifying their validity.
abstract contract StateHub is AgentSecured, StateHubEvents, IStateHub {
    using SafeCast for uint256;
    using StateLib for bytes;

    struct OriginState {
        uint40 blockNumber;
        uint40 timestamp;
        GasData gasData;
    }
    // Bits left for tight packing: 80

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev Historical Merkle Tree
    /// Note: Takes two storage slots
    HistoricalTree private _tree;

    /// @dev All historical contract States
    OriginState[] private _originStates;

    /// @dev gap for upgrade safety
    uint256[47] private __GAP; // solhint-disable-line var-name-mixedcase

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IStateHub
    function isValidState(bytes memory statePayload) external view returns (bool isValid) {
        // This will revert if payload is not a formatted state
        State state = statePayload.castToState();
        return _isValidState(state);
    }

    /// @inheritdoc IStateHub
    function statesAmount() external view returns (uint256) {
        return _originStates.length;
    }

    /// @inheritdoc IStateHub
    function suggestLatestState() external view returns (bytes memory stateData) {
        // This never underflows, assuming the contract was initialized
        return suggestState(_nextNonce() - 1);
    }

    /// @inheritdoc IStateHub
    function suggestState(uint32 nonce) public view returns (bytes memory stateData) {
        // This will revert if nonce is out of range
        bytes32 root = _tree.root(nonce);
        return _formatOriginState(_originStates[nonce], root, localDomain, nonce);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Initializes the saved states list by inserting a state for an empty Merkle Tree.
    function _initializeStates() internal {
        // This should only be called once, when the contract is initialized
        // This will revert if _tree.roots is non-empty
        bytes32 savedRoot = _tree.initializeRoots();
        // Save root for empty merkle _tree with block number and timestamp of initialization
        _saveState(savedRoot, _toOriginState());
    }

    /// @dev Inserts leaf into the Merkle Tree and saves the updated origin State.
    function _insertAndSave(bytes32 leaf) internal {
        bytes32 newRoot = _tree.insert(leaf);
        _saveState(newRoot, _toOriginState());
    }

    /// @dev Saves an updated state of the Origin contract
    function _saveState(bytes32 root, OriginState memory state) internal {
        uint32 nonce = _nextNonce();
        _originStates.push(state);
        // Emit event with raw state data
        emit StateSaved(_formatOriginState(state, root, localDomain, nonce));
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns nonce of the next sent message: the amount of saved States so far.
    /// This always equals to "total amount of sent messages" plus 1.
    function _nextNonce() internal view returns (uint32) {
        // TODO: consider using more than 32 bits for origin nonces
        return _originStates.length.toUint32();
    }

    /// @dev Checks if a state is valid, i.e. if it matches the historical one.
    /// Reverts, if state refers to another Origin contract.
    function _isValidState(State state) internal view returns (bool) {
        // Check if state refers to this contract
        if (state.origin() != localDomain) revert IncorrectOriginDomain();
        // Check if nonce exists
        uint32 nonce = state.nonce();
        if (nonce >= _originStates.length) return false;
        // Check if state root matches the historical one
        if (state.root() != _tree.root(nonce)) return false;
        // Check if state metadata matches the historical one
        return _areEqual(state, _originStates[nonce]);
    }

    // ═════════════════════════════════════════════ STRUCT FORMATTING ═════════════════════════════════════════════════

    /// @dev Returns a formatted payload for a stored OriginState.
    function _formatOriginState(OriginState memory originState, bytes32 root, uint32 origin, uint32 nonce)
        internal
        pure
        returns (bytes memory)
    {
        return StateLib.formatState({
            root_: root,
            origin_: origin,
            nonce_: nonce,
            blockNumber_: originState.blockNumber,
            timestamp_: originState.timestamp,
            gasData_: originState.gasData
        });
    }

    /// @dev Child contract should implement the logic for getting the current gas data from the gas oracle
    /// to be saved as part of the Origin State.
    // solhint-disable-next-line ordering
    function _fetchGasData() internal view virtual returns (GasData);

    /// @dev Returns a OriginState struct to save in the contract.
    function _toOriginState() internal view returns (OriginState memory originState) {
        originState.blockNumber = ChainContext.blockNumber();
        originState.timestamp = ChainContext.blockTimestamp();
        originState.gasData = _fetchGasData();
    }

    /// @dev Checks that a state and its Origin representation are equal.
    function _areEqual(State state, OriginState memory originState) internal pure returns (bool) {
        return state.blockNumber() == originState.blockNumber && state.timestamp() == originState.timestamp;
    }
}
