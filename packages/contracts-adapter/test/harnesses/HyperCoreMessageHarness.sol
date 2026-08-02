// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {HyperCoreMessage} from "../../src/libs/HyperCoreMessage.sol";

// solhint-disable no-empty-blocks
contract HyperCoreMessageHarness {
    function testHyperCoreMessageHarness() external {}

    function encodeHyperCoreMessage(
        address recipient,
        address srcToken,
        uint256 amount
    )
        public
        pure
        returns (bytes memory)
    {
        return HyperCoreMessage.encodeHyperCoreMessage(recipient, srcToken, amount);
    }

    function decodeHyperCoreMessage(bytes calldata payload)
        public
        pure
        returns (address recipient, address srcToken, uint256 amount)
    {
        return HyperCoreMessage.decodeHyperCoreMessage(payload);
    }
}
