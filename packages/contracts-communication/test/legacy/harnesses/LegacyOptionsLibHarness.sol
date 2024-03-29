// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyOptionsLib} from "../../../contracts/legacy/libs/LegacyOptions.sol";

contract LegacyOptionsLibHarness {
    function encodeLegacyOptions(uint256 gasLimit) external pure returns (bytes memory) {
        return LegacyOptionsLib.encodeLegacyOptions(gasLimit);
    }

    function decodeLegacyOptions(bytes calldata legacyOpts) external pure returns (uint256) {
        return LegacyOptionsLib.decodeLegacyOptions(legacyOpts);
    }
}
