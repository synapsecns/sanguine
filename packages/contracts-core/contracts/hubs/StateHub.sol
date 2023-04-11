// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {HistoricalTree} from "../libs/Merkle.sol";
import {OriginState, State, StateLib} from "../libs/State.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DomainContext} from "../context/DomainContext.sol";
import {StateHubEvents} from "../events/StateHubEvents.sol";
import {IStateHub} from "../interfaces/IStateHub.sol";

/**
 * @notice Hub to accept, save and verify states for a local contract.
 * The State logic is fully outsourced to the State library, which defines
 * - What a "state" is
 * - How "state" getters work
 * - How to compare "states" to one another
 */
abstract contract StateHub is DomainContext, StateHubEvents, IStateHub {
    using StateLib for bytes;

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
        OriginState memory state = _originStates[nonce];
        return state.formatOriginState({root_: root, origin_: localDomain, nonce_: nonce});
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Initializes the saved states list by inserting a state for an empty Merkle Tree.
    function _initializeStates() internal {
        // This should only be called once, when the contract is initialized
        // This will revert if _tree.roots is non-empty
        bytes32 savedRoot = _tree.initializeRoots();
        // Save root for empty merkle _tree with block number and timestamp of initialization
        _saveState(savedRoot, StateLib.originState());
    }

    /// @dev Inserts leaf into the Merkle Tree and saves the updated origin State.
    function _insertAndSave(bytes32 leaf) internal {
        bytes32 newRoot = _tree.insert(leaf);
        _saveState(newRoot, StateLib.originState());
    }

    /// @dev Saves an updated state of the Origin contract
    function _saveState(bytes32 root, OriginState memory state) internal {
        // State nonce is its index in `_originStates` array
        uint32 stateNonce = uint32(_originStates.length);
        _originStates.push(state);
        // Emit event with raw state data
        emit StateSaved(state.formatOriginState({root_: root, origin_: localDomain, nonce_: stateNonce}));
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns nonce of the next sent message: the amount of saved States so far.
    /// This always equals to "total amount of sent messages" plus 1.
    function _nextNonce() internal view returns (uint32) {
        return uint32(_originStates.length);
    }

    /// @dev Checks if a state is valid, i.e. if it matches the historical one.
    /// Reverts, if state refers to another Origin contract.
    function _isValidState(State state) internal view returns (bool) {
        // Check if state refers to this contract
        require(state.origin() == localDomain, "Wrong origin");
        // Check if nonce exists
        uint32 nonce = state.nonce();
        if (nonce >= _originStates.length) return false;
        // Check if state root matches the historical one
        if (state.root() != _tree.root(nonce)) return false;
        // Check if state metadata matches the historical one
        return state.equalToOrigin(_originStates[nonce]);
    }
}
