// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {GasData, State, StateLib, MemView, MemViewLib} from "../../../../contracts/libs/memory/State.sol";

// solhint-disable ordering
/// @notice Exposes State methods for testing against golang.
contract StateHarness {
    using StateLib for bytes;
    using StateLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToState(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        State state = StateLib.castToState(payload);
        return state.unwrap().clone();
    }

    function equals(bytes memory a, bytes memory b) public pure returns (bool) {
        return a.castToState().equals(b.castToState());
    }

    function hashInvalid(bytes memory payload) public pure returns (bytes32) {
        return payload.castToState().hashInvalid();
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToState().leaf();
    }

    function subLeafs(bytes memory payload) public pure returns (bytes32, bytes32) {
        return payload.castToState().subLeafs();
    }

    function leftLeaf(bytes32 root_, uint32 origin_) public pure returns (bytes32) {
        return StateLib.leftLeaf(root_, origin_);
    }

    function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, GasData gasData_)
        public
        pure
        returns (bytes32)
    {
        return StateLib.rightLeaf(nonce_, blockNumber_, timestamp_, gasData_);
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

    function gasData(bytes memory payload) public pure returns (GasData) {
        return payload.castToState().gasData();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatState(
        bytes32 root_,
        uint32 origin_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_,
        GasData gasData_
    ) public pure returns (bytes memory) {
        return StateLib.formatState(root_, origin_, nonce_, blockNumber_, timestamp_, gasData_);
    }

    function isState(bytes memory payload) public pure returns (bool) {
        return payload.ref().isState();
    }
}
