// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/State.sol";

/**
 * @notice Exposes State methods for testing against golang.
 */
contract StateHarness {
    using StateLib for bytes;
    using StateLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToState(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        State _state = StateLib.castToState(_payload);
        return _state.unwrap().clone();
    }

    function root(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToState().root();
    }

    function origin(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToState().origin();
    }

    function nonce(bytes memory _payload) public pure returns (uint32) {
        return _payload.castToState().nonce();
    }

    function blockNumber(bytes memory _payload) public pure returns (uint40) {
        return _payload.castToState().blockNumber();
    }

    function timestamp(bytes memory _payload) public pure returns (uint40) {
        return _payload.castToState().timestamp();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         STATE FORMATTERS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatState(
        bytes32 _root,
        uint32 _origin,
        uint32 _nonce,
        uint40 _blockNumber,
        uint40 _timestamp
    ) public pure returns (bytes memory) {
        return StateLib.formatState(_root, _origin, _nonce, _blockNumber, _timestamp);
    }

    function isState(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isState();
    }
}
