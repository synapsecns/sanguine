// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {TypeCasts} from "../libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessageRecipient} from "./MessageRecipient.sol";

contract TestClient is MessageRecipient {
    event MessageReceived(
        uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, uint32 version, bytes content
    );

    event MessageSent(uint32 destination, uint32 nonce, bytes32 sender, bytes32 recipient, bytes content);

    // solhint-disable-next-line no-empty-blocks
    constructor(address origin_, address destination_) MessageRecipient(origin_, destination_) {}

    function sendMessage(
        uint32 destination_,
        address recipientAddress,
        uint32 optimisticPeriod,
        uint64 gasLimit,
        uint32 version,
        bytes memory content
    ) external payable {
        bytes32 recipient = TypeCasts.addressToBytes32(recipientAddress);
        MessageRequest memory request = MessageRequest({gasDrop: 0, gasLimit: gasLimit, version: version});
        (uint32 nonce,) = _sendBaseMessage({
            destination_: destination_,
            recipient: recipient,
            optimisticPeriod: optimisticPeriod,
            tipsValue: msg.value,
            request: request,
            content: content
        });
        emit MessageSent(destination_, nonce, TypeCasts.addressToBytes32(address(this)), recipient, content);
    }

    /// @inheritdoc MessageRecipient
    function _receiveBaseMessageUnsafe(
        uint32 origin_,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) internal override {
        emit MessageReceived(origin_, nonce, sender, proofMaturity, version, content);
    }
}
