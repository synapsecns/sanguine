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

    function castToState(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        State state = StateLib.castToState(payload);
        return state.unwrap().clone();
    }

    function equals(bytes memory a, bytes memory b) public pure returns (bool) {
        return a.castToState().equals(b.castToState());
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToState().leaf();
    }

    function subLeafs(bytes memory payload) public pure returns (bytes32, bytes32) {
        return payload.castToState().subLeafs();
    }

    function leftLeaf(bytes32 root, uint32 origin) public pure returns (bytes32) {
        return StateLib.leftLeaf(root, origin);
    }

    function rightLeaf(
        uint32 nonce,
        uint40 blockNumber,
        uint40 timestamp
    ) public pure returns (bytes32) {
        return StateLib.rightLeaf(nonce, blockNumber, timestamp);
    }

    function root(bytes memory payload) public pure returns (bytes32) {
        return payload.castToState().root();
    }

    function origin(bytes memory payload) public pure returns (uint32) {
        return payload.castToState().origin();
    }

    function nonce(bytes memory payload) public pure returns (uint32) {
        return payload.castToState().nonce();
    }

    function blockNumber(bytes memory payload) public pure returns (uint40) {
        return payload.castToState().blockNumber();
    }

    function timestamp(bytes memory payload) public pure returns (uint40) {
        return payload.castToState().timestamp();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ORIGIN STATE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatOriginState(
        OriginState memory originState,
        bytes32 root,
        uint32 origin,
        uint32 nonce
    ) public pure returns (bytes memory) {
        return originState.formatOriginState(root, origin, nonce);
    }

    function originState() public view returns (OriginState memory state) {
        return StateLib.originState();
    }

    function equalToOrigin(bytes memory payload, OriginState memory originState)
        public
        pure
        returns (bool)
    {
        return payload.castToState().equalToOrigin(originState);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             SUMMIT STATE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSummitState(SummitState memory summitState) public pure returns (bytes memory) {
        return summitState.formatSummitState();
    }

    function toSummitState(bytes memory payload) public pure returns (SummitState memory state) {
        return payload.castToState().toSummitState();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           STATE FORMATTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatState(
        bytes32 root,
        uint32 origin,
        uint32 nonce,
        uint40 blockNumber,
        uint40 timestamp
    ) public pure returns (bytes memory) {
        return StateLib.formatState(root, origin, nonce, blockNumber, timestamp);
    }

    function isState(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isState();
    }
}
