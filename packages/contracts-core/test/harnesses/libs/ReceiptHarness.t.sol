// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Receipt, ReceiptBody, ReceiptLib, Tips, MemView, MemViewLib} from "../../../contracts/libs/Receipt.sol";

// solhint-disable ordering
contract ReceiptHarness {
    using ReceiptLib for bytes;
    using ReceiptLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ═══════════════════════════════════════════ RECEIPT BODY GETTERS ════════════════════════════════════════════════

    function castToReceiptBody(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        ReceiptBody receiptBody = ReceiptLib.castToReceiptBody(payload);
        return receiptBody.unwrap().clone();
    }

    /// @notice Returns receipt's origin field
    function origin(bytes memory payload) public pure returns (uint32) {
        return payload.castToReceiptBody().origin();
    }

    /// @notice Returns receipt's destination field
    function destination(bytes memory payload) public pure returns (uint32) {
        return payload.castToReceiptBody().destination();
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToReceiptBody().messageHash();
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(bytes memory payload) public pure returns (bytes32) {
        return payload.castToReceiptBody().snapshotRoot();
    }

    /// @notice Returns receipt's "state index" field
    function stateIndex(bytes memory payload) public pure returns (uint8) {
        return payload.castToReceiptBody().stateIndex();
    }

    /// @notice Returns receipt's "attestation notary" field
    function attNotary(bytes memory payload) public pure returns (address) {
        return payload.castToReceiptBody().attNotary();
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(bytes memory payload) public pure returns (address) {
        return payload.castToReceiptBody().firstExecutor();
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(bytes memory payload) public pure returns (address) {
        return payload.castToReceiptBody().finalExecutor();
    }

    function isReceiptBody(bytes memory payload) public pure returns (bool) {
        return payload.ref().isReceiptBody();
    }

    function hashInvalid(bytes memory payload) public pure returns (bytes32) {
        return payload.ref().castToReceiptBody().hashInvalid();
    }

    function equals(bytes memory a, bytes memory b) public pure returns (bool) {
        return a.ref().castToReceiptBody().equals(b.ref().castToReceiptBody());
    }

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToReceipt(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Receipt receipt = ReceiptLib.castToReceipt(payload);
        return receipt.unwrap().clone();
    }

    /// @notice Returns Receipt's body field
    function body(bytes memory payload) public view returns (bytes memory) {
        return payload.castToReceipt().body().unwrap().clone();
    }

    /// @notice Returns Receipt's tips field
    function tips(bytes memory payload) public pure returns (uint256) {
        return Tips.unwrap(payload.castToReceipt().tips());
    }

    function isReceipt(bytes memory payload) public pure returns (bool) {
        return payload.ref().isReceipt();
    }

    function hashValid(bytes memory payload) public pure returns (bytes32) {
        return payload.ref().castToReceipt().hashValid();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatReceiptBody(
        uint32 origin_,
        uint32 destination_,
        bytes32 messageHash_,
        bytes32 snapshotRoot_,
        uint8 stateIndex_,
        address attNotary_,
        address firstExecutor_,
        address finalExecutor_
    ) public pure returns (bytes memory) {
        return ReceiptLib.formatReceiptBody(
            origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_
        );
    }

    function formatReceipt(bytes memory bodyPayload, Tips tips_) public pure returns (bytes memory) {
        return ReceiptLib.formatReceipt(bodyPayload, tips_);
    }
}
