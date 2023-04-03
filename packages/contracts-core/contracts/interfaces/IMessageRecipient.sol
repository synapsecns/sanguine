// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IMessageRecipient {
    /**
     * @notice Message recipient needs to implement this function in order to
     * receive cross-chain messages.
     * @dev Message recipient needs to ensure that merkle proof for the message
     * is at least as old as the optimistic period that the recipient is using.
     * Note: as this point it is checked that the "message optimistic period" has passed,
     * however the period value itself could be anything, and thus could differ from the one
     * that the recipient would like to enforce.
     * @param origin            Domain where message originated
     * @param nonce             Message nonce on the origin domain
     * @param sender            Sender address on origin chain
     * @param rootSubmittedAt   Time when merkle root (used for proving this message) was submitted
     * @param content           Raw bytes content of message
     */
    function receiveBaseMessage(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 rootSubmittedAt,
        bytes memory content
    ) external payable;
}
