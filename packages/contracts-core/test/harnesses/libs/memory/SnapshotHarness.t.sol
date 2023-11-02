// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    ChainGas,
    Snapshot,
    SnapshotLib,
    State,
    StateLib,
    MemView,
    MemViewLib
} from "../../../../contracts/libs/memory/Snapshot.sol";

// solhint-disable ordering

/**
 * @notice Exposes Snapshot methods for testing against golang.
 */
contract SnapshotHarness {
    using StateLib for bytes;
    using SnapshotLib for bytes;
    using SnapshotLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToSnapshot(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Snapshot snapshot = SnapshotLib.castToSnapshot(payload);
        return snapshot.unwrap().clone();
    }

    function hashValid(bytes memory payload) public pure returns (bytes32) {
        return payload.castToSnapshot().hashValid();
    }

    function state(bytes memory payload, uint8 stateIndex) public view returns (bytes memory) {
        return payload.castToSnapshot().state(stateIndex).unwrap().clone();
    }

    function statesAmount(bytes memory payload) public pure returns (uint256) {
        return payload.castToSnapshot().statesAmount();
    }

    function snapGas(bytes memory payload) public pure returns (ChainGas[] memory) {
        return payload.castToSnapshot().snapGas();
    }

    function calculateRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToSnapshot().calculateRoot();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatSnapshot(bytes[] memory statePayloads) public view returns (bytes memory) {
        uint256 length = statePayloads.length;
        State[] memory states = new State[](length);
        for (uint256 i = 0; i < length; ++i) {
            states[i] = statePayloads[i].castToState();
        }
        return SnapshotLib.formatSnapshot(states);
    }

    function isSnapshot(bytes memory payload) public pure returns (bool) {
        return payload.ref().isSnapshot();
    }
}
