// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";

// solhint-disable no-empty-blocks
contract MessageRecipientMock is IMessageRecipient {
    /// @notice Prevents this contract from being included in the coverage report
    function testMessageRecipientMock() external {}

    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) external {}
}
