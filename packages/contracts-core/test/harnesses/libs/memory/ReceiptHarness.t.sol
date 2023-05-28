// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Receipt, ReceiptLib, MemView, MemViewLib} from "../../../../contracts/libs/memory/Receipt.sol";

// solhint-disable ordering
contract ReceiptHarness {
    using ReceiptLib for bytes;
    using ReceiptLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToReceipt(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Receipt receipt = ReceiptLib.castToReceipt(payload);
        return receipt.unwrap().clone();
    }

    /// @notice Returns receipt's origin field
    function origin(bytes memory payload) public pure returns (uint32) {
        return payload.castToReceipt().origin();
    }

    /// @notice Returns receipt's destination field
    function destination(bytes memory payload) public pure returns (uint32) {
        return payload.castToReceipt().destination();
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToReceipt().messageHash();
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToReceipt().snapshotRoot();
    }

    /// @notice Returns receipt's "state index" field
    function stateIndex(bytes memory payload) public pure returns (uint8) {
        return payload.castToReceipt().stateIndex();
    }

    /// @notice Returns receipt's "attestation notary" field
    function attNotary(bytes memory payload) public pure returns (address) {
        return payload.castToReceipt().attNotary();
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(bytes memory payload) public pure returns (address) {
        return payload.castToReceipt().firstExecutor();
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(bytes memory payload) public pure returns (address) {
        return payload.castToReceipt().finalExecutor();
    }

    function equals(bytes memory a, bytes memory b) public pure returns (bool) {
        return a.ref().castToReceipt().equals(b.ref().castToReceipt());
    }

    function isReceipt(bytes memory payload) public pure returns (bool) {
        return payload.ref().isReceipt();
    }

    function hashValid(bytes memory payload) public pure returns (bytes32) {
        return payload.ref().castToReceipt().hashValid();
    }

    function hashInvalid(bytes memory payload) public pure returns (bytes32) {
        return payload.ref().castToReceipt().hashInvalid();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatReceipt(
        uint32 origin_,
        uint32 destination_,
        bytes32 messageHash_,
        bytes32 snapshotRoot_,
        uint8 stateIndex_,
        address attNotary_,
        address firstExecutor_,
        address finalExecutor_
    ) public pure returns (bytes memory) {
        return ReceiptLib.formatReceipt(
            origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_
        );
    }
}
