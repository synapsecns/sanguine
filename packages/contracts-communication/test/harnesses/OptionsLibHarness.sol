// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1, OptionsLib} from "../../contracts/libs/Options.sol";

contract OptionsLibHarness {
    function decodeOptionsV1(bytes calldata data) external view returns (OptionsV1 memory) {
        return OptionsLib.decodeOptionsV1(data);
    }

    function encodeOptionsV1(OptionsV1 memory options) external pure returns (bytes memory) {
        return OptionsLib.encodeOptionsV1(options);
    }
}
