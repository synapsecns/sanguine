// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    OriginState,
    State,
    StateLib,
    SummitState,
    TypedMemView
} from "../../../contracts/libs/State.sol";

/**
 * @notice Exposes State methods for testing against golang.
 */
contract StateHarness {
    using StateLib for bytes;
    using StateLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToState(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        State _state = StateLib.castToState(_payload);
        return _state.unwrap().clone();
    }

    function equals(bytes memory _a, bytes memory _b) public pure returns (bool) {
        return _a.castToState().equals(_b.castToState());
    }

    function leaf(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToState().leaf();
    }

    function subLeafs(bytes memory _payload) public pure returns (bytes32, bytes32) {
        return _payload.castToState().subLeafs();
    }

    function leftLeaf(bytes32 _root, uint32 _origin) public pure returns (bytes32) {
        return StateLib.leftLeaf(_root, _origin);
    }

    function rightLeaf(
        uint32 _nonce,
        uint40 _blockNumber,
        uint40 _timestamp
    ) public pure returns (bytes32) {
        return StateLib.rightLeaf(_nonce, _blockNumber, _timestamp);
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
    ▏*║                             ORIGIN STATE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatOriginState(
        OriginState memory _originState,
        uint32 _origin,
        uint32 _nonce
    ) public pure returns (bytes memory) {
        return _originState.formatOriginState(_origin, _nonce);
    }

    function originState(bytes32 _root) public view returns (OriginState memory state) {
        return StateLib.originState(_root);
    }

    function equalToOrigin(bytes memory _payload, OriginState memory _originState)
        public
        pure
        returns (bool)
    {
        return _payload.castToState().equalToOrigin(_originState);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             SUMMIT STATE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSummitState(SummitState memory _summitState) public pure returns (bytes memory) {
        return _summitState.formatSummitState();
    }

    function toSummitState(bytes memory _payload) public pure returns (SummitState memory state) {
        return _payload.castToState().toSummitState();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           STATE FORMATTERS                           ║*▕
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
