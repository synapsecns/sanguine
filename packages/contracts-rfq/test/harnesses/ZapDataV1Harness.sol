// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ZapDataV1} from "../../contracts/libs/ZapDataV1.sol";

contract ZapDataV1Harness {
    uint16 public constant VERSION = ZapDataV1.VERSION;
    uint16 public constant AMOUNT_NOT_PRESENT = ZapDataV1.AMOUNT_NOT_PRESENT;

    function validateV1(bytes calldata encodedZapData) public pure {
        ZapDataV1.validateV1(encodedZapData);
    }

    function encodeV1(
        uint16 amountPosition_,
        address finalToken_,
        address forwardTo_,
        address target_,
        bytes memory payload_
    )
        public
        pure
        returns (bytes memory encodedZapData)
    {
        return ZapDataV1.encodeV1(amountPosition_, finalToken_, forwardTo_, target_, payload_);
    }

    function version(bytes calldata encodedZapData) public pure returns (uint16) {
        return ZapDataV1.version(encodedZapData);
    }

    function finalToken(bytes calldata encodedZapData) public pure returns (address) {
        return ZapDataV1.finalToken(encodedZapData);
    }

    function forwardTo(bytes calldata encodedZapData) public pure returns (address) {
        return ZapDataV1.forwardTo(encodedZapData);
    }

    function target(bytes calldata encodedZapData) public pure returns (address) {
        return ZapDataV1.target(encodedZapData);
    }

    function payload(bytes calldata encodedZapData, uint256 amount) public pure returns (bytes memory) {
        return ZapDataV1.payload(encodedZapData, amount);
    }
}
