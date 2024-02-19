// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;
import {Test} from "forge-std/Test.sol";

import { OptionsLib, OptionsLibHarness } from "../harnesses/OptionsLibHarness.sol";

contract OptionsLibTest is Test {
    OptionsLibHarness public libHarness;

    function setUp() public {
        libHarness = new OptionsLibHarness();
    }

    function test_encodeOptions() public {
        uint8 version = 1;
        // 200k gas limit
        uint256 gasLimit = 200000;
        // 100k wei
        uint256 gasAirdrop = 100000;
        bytes memory expected = abi.encode(version, gasLimit, gasAirdrop);
        bytes memory actual = libHarness.encodeOptions(version, gasLimit, gasAirdrop);
        assertEq(actual, expected);
    }

    function test_decodeOptions() public {
        uint8 version = 1;
        // 200k gas limit
        uint256 gasLimit = 200000;
        // 100k wei
        uint256 gasAirdrop = 100000;
        bytes memory data = abi.encode(version, gasLimit, gasAirdrop);
        (uint8 actualVersion, uint256 actualGasLimit, uint256 actualGasAirdrop) = libHarness.decodeOptions(data);
        assertEq(actualVersion, version);
        assertEq(actualGasLimit, gasLimit);
        assertEq(actualGasAirdrop, gasAirdrop);
    }
}
