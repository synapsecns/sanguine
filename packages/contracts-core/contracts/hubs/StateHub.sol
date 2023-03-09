// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { EMPTY_ROOT } from "../libs/Constants.sol";
import { OriginState, State, StateLib } from "../libs/State.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DomainContext } from "../context/DomainContext.sol";
import { StateHubEvents } from "../events/StateHubEvents.sol";
import { IStateHub } from "../interfaces/IStateHub.sol";

/**
 * @notice Hub to accept, save and verify states for a local contract.
 * The State logic is fully outsourced to the State library, which defines
 * - What a "state" is
 * - How "state" getters work
 * - How to compare "states" to one another
 */
abstract contract StateHub is DomainContext, StateHubEvents, IStateHub {
    using StateLib for bytes;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev All historical contract States
    OriginState[] private originStates;

    /// @dev gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IStateHub
    function isValidState(bytes memory _statePayload) external view returns (bool isValid) {
        // This will revert if payload is not a formatted state
        State state = _statePayload.castToState();
        return _isValidState(state);
    }

    /// @inheritdoc IStateHub
    function statesAmount() external view returns (uint256) {
        return originStates.length;
    }

    /// @inheritdoc IStateHub
    function suggestLatestState() external view returns (bytes memory stateData) {
        // This never underflows, assuming the contract was initialized
        return suggestState(_nextNonce() - 1);
    }

    /// @inheritdoc IStateHub
    function suggestState(uint32 _nonce) public view returns (bytes memory stateData) {
        require(_nonce < _nextNonce(), "Nonce out of range");
        OriginState memory state = originStates[_nonce];
        return state.formatOriginState({ _origin: localDomain, _nonce: _nonce });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SAVE STATE DATA                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Initializes the saved states list by inserting a state for an empty Merkle Tree.
    function _initializeStates() internal {
        // This should only be called once, when the contract is initialized
        assert(originStates.length == 0);
        // Save root for empty merkle tree with block number and timestamp of initialization
        _saveState(StateLib.originState(EMPTY_ROOT));
    }

    /// @dev Saves an updated state of the Origin contract
    function _saveState(OriginState memory _state) internal {
        // State nonce is its index in `originStates` array
        uint32 stateNonce = uint32(originStates.length);
        originStates.push(_state);
        // Emit event with raw state data
        emit StateSaved(_state.formatOriginState({ _origin: localDomain, _nonce: stateNonce }));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VERIFY STATE DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns nonce of the next dispatched message: the amount of saved States so far.
    /// This always equals to "total amount of dispatched messages" plus 1.
    function _nextNonce() internal view returns (uint32) {
        return uint32(originStates.length);
    }

    /// @dev Checks if a state is valid, i.e. if it matches the historical one.
    /// Reverts, if state refers to another Origin contract.
    function _isValidState(State _state) internal view returns (bool) {
        // Check if state refers to this contract
        require(_state.origin() == localDomain, "Wrong origin");
        // Check if nonce exists
        uint32 nonce = _state.nonce();
        if (nonce >= originStates.length) return false;
        // Check if state matches the historical one
        return _state.equalToOrigin(originStates[nonce]);
    }
}
