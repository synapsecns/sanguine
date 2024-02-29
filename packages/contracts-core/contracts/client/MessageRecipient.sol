// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {
    CallerNotDestination,
    IncorrectNonce,
    IncorrectSender,
    IncorrectRecipient,
    ZeroProofMaturity
} from "../libs/Errors.sol";
import {Request, RequestLib} from "../libs/stack/Request.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";
import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";

abstract contract MessageRecipient is IMessageRecipient {
    struct MessageRequest {
        uint96 gasDrop;
        uint64 gasLimit;
        uint32 version;
    }

    /// @notice Local chain Origin: used for sending messages
    address public immutable origin;

    /// @notice Local chain Destination: used for receiving messages
    address public immutable destination;

    constructor(address origin_, address destination_) {
        origin = origin_;
        destination = destination_;
    }

    /// @inheritdoc IMessageRecipient
    function receiveBaseMessage(
        uint32 origin_,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) external payable {
        if (msg.sender != destination) revert CallerNotDestination();
        if (nonce == 0) revert IncorrectNonce();
        if (sender == 0) revert IncorrectSender();
        if (proofMaturity == 0) revert ZeroProofMaturity();
        _receiveBaseMessageUnsafe(origin_, nonce, sender, proofMaturity, version, content);
    }

    /**
     * @dev Child contracts should implement the logic for receiving a Base Message in an "unsafe way".
     * Following checks HAVE been performed:
     *  - receiveBaseMessage() was called by Destination (i.e. this is a legit base message).
     *  - Nonce is not zero.
     *  - Message sender on origin chain is not a zero address.
     *  - Proof maturity is not zero.
     * Following checks HAVE NOT been performed (thus "unsafe"):
     *  - Message sender on origin chain could be anything non-zero at this point.
     *  - Proof maturity could be anything non-zero at this point.
     */
    function _receiveBaseMessageUnsafe(
        uint32 origin_,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) internal virtual;

    /**
     * @dev Sends a message to given destination chain. Full `msg.value` is used to pay for the message tips.
     * `_getMinimumTipsValue()` could be used to calculate the minimum required tips value, and should be also
     * exposed as a public view function to estimate the tips value before sending a message off-chain.
     * This function is not exposed in MessageRecipient, as the message encoding is implemented by the child contract.
     * @param destination_          Domain of the destination chain
     * @param recipient             Address of the recipient on destination chain
     * @param optimisticPeriod      Optimistic period for the message
     * @param tipsValue             Tips to be paid for sending the message
     * @param request               Message execution request on destination chain
     * @param content               The message content
     */
    function _sendBaseMessage(
        uint32 destination_,
        bytes32 recipient,
        uint32 optimisticPeriod,
        uint256 tipsValue,
        MessageRequest memory request,
        bytes memory content
    ) internal returns (uint32 messageNonce, bytes32 messageHash) {
        if (recipient == 0) revert IncorrectRecipient();
        return InterfaceOrigin(origin).sendBaseMessage{value: tipsValue}(
            destination_, recipient, optimisticPeriod, _encodeRequest(request), content
        );
    }

    /**
     * @dev Returns the minimum tips value for sending a message to given destination chain.
     * @param destination_          Domain of the destination chain
     * @param request               Message execution request on destination chain
     * @param contentLength         Length of the message content
     */
    function _getMinimumTipsValue(uint32 destination_, MessageRequest memory request, uint256 contentLength)
        internal
        view
        returns (uint256 tipsValue)
    {
        return InterfaceOrigin(origin).getMinimumTipsValue(destination_, _encodeRequest(request), contentLength);
    }

    /**
     * @dev Encodes a message execution request into format that Origin contract is using.
     * @param request               Message execution request on destination chain
     * @return paddedRequest        Encoded request
     */
    function _encodeRequest(MessageRequest memory request) internal pure returns (uint256 paddedRequest) {
        return Request.unwrap(RequestLib.encodeRequest(request.gasDrop, request.gasLimit, request.version));
    }
}
