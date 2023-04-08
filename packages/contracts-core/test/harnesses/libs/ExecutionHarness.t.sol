// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Execution, ExecutionLib, MessageStatus, TypedMemView} from "../../../contracts/libs/Execution.sol";

// solhint-disable ordering
contract ExecutionHarness {
    using ExecutionLib for bytes;
    using ExecutionLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToExecution(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Execution execution = ExecutionLib.castToExecution(payload);
        return execution.unwrap().clone();
    }

    /// @notice Returns execution's status.
    function status(bytes memory payload) public pure returns (MessageStatus) {
        return payload.castToExecution().status();
    }

    /// @notice Returns execution's origin field
    function origin(bytes memory payload) public pure returns (uint32) {
        return payload.castToExecution().origin();
    }

    /// @notice Returns execution's destination field
    function destination(bytes memory payload) public pure returns (uint32) {
        return payload.castToExecution().destination();
    }

    /// @notice Returns execution's "message hash" field
    function messageHash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToExecution().messageHash();
    }

    /// @notice Returns execution's "snapshot root" field
    function snapshotRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToExecution().snapshotRoot();
    }

    /// @notice Returns execution's "first executor" field
    function firstExecutor(bytes memory payload) public pure returns (address) {
        return payload.castToExecution().firstExecutor();
    }

    /// @notice Returns execution's "final executor" field
    function finalExecutor(bytes memory payload) public pure returns (address) {
        return payload.castToExecution().finalExecutor();
    }

    /// @notice Returns baseMessage's tips field
    function tips(bytes memory payload) public view returns (bytes memory) {
        return payload.castToExecution().tips().unwrap().clone();
    }

    function isExecution(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isExecution();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatExecution(
        MessageStatus status_,
        uint32 origin_,
        uint32 destination_,
        bytes32 messageHash_,
        bytes32 snapshotRoot_,
        address firstExecutor_,
        address finalExecutor_,
        bytes memory tipsPayload
    ) public pure returns (bytes memory) {
        return ExecutionLib.formatExecution(
            status_, origin_, destination_, messageHash_, snapshotRoot_, firstExecutor_, finalExecutor_, tipsPayload
        );
    }
}
