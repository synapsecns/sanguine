// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsLib, OptionsV1} from "../libs/Options.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

contract OptionsLibMocks {
    function decodeOptions(bytes memory data) external view returns (OptionsV1 memory) {
        return OptionsLib.decodeOptionsV1(data);
    }

    function encodeOptions(OptionsV1 memory options) external pure returns (bytes memory) {
        return OptionsLib.encodeOptionsV1(options);
    }

    function addressToBytes32(address convertable) external pure returns (bytes32) {
        return TypeCasts.addressToBytes32(convertable);
    }
}
