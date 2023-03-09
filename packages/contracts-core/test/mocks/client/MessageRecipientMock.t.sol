// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";

// solhint-disable no-empty-blocks
contract MessageRecipientMock is IMessageRecipient {
    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootTimestamp,
        bytes memory _message
    ) external {}
}
