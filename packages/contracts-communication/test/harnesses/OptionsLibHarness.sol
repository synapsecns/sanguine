// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1, OptionsLib} from "../../contracts/libs/Options.sol";

contract OptionsLibHarness {
    function encodeVersionedOptions(uint8 version, bytes calldata options) external pure returns (bytes memory) {
        return OptionsLib.encodeVersionedOptions(version, options);
    }

    function decodeVersionedOptions(bytes calldata data) external pure returns (uint8, bytes memory) {
        (uint8 version, bytes memory options) = OptionsLib.decodeVersionedOptions(data);
        return (version, options);
    }

    function encodeOptionsV1(OptionsV1 memory options) external pure returns (bytes memory) {
        return OptionsLib.encodeOptionsV1(options);
    }

    function decodeOptionsV1(bytes calldata data) external pure returns (OptionsV1 memory) {
        return OptionsLib.decodeOptionsV1(data);
    }
}
