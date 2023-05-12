// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IMessageRecipient} from "../../../contracts/interfaces/IMessageRecipient.sol";

// solhint-disable no-empty-blocks
contract MessageRecipientMock is IMessageRecipient {
    /// @notice Prevents this contract from being included in the coverage report
    function testMessageRecipientMock() external {}

    function receiveBaseMessage(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) external payable {}
}
