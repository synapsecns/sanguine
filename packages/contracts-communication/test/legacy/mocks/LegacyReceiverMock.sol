// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ILegacyReceiver} from "../../../contracts/legacy/interfaces/ILegacyReceiver.sol";

// solhint-disable no-empty-blocks
contract LegacyReceiverMock is ILegacyReceiver {
    function executeMessage(bytes32 srcAddress, uint256 srcChainId, bytes memory message, address executor) external {}
}
