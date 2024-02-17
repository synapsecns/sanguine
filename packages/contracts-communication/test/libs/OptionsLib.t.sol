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
        OptionsLib.NativeDrop[] memory nativeDrops = new OptionsLib.NativeDrop[](1);
        nativeDrops[0] = OptionsLib.NativeDrop({recipient: address(0x1), amount: 1000});
        bytes memory expected = abi.encode(version, gasLimit, nativeDrops);
        bytes memory actual = libHarness.encodeOptions(version, gasLimit, nativeDrops);
        assertEq(actual, expected);
    }

    function test_decodeOptions() public {
        uint8 version = 1;
        // 200k gas limit
        uint256 gasLimit = 200000;
        // 100k wei
        uint256 gasAirdrop = 100000;
        OptionsLib.NativeDrop[] memory nativeDrops = new OptionsLib.NativeDrop[](1);
        nativeDrops[0] = OptionsLib.NativeDrop({recipient: address(0x1), amount: gasAirdrop});
        bytes memory data = abi.encode(version, gasLimit, nativeDrops);
        (uint8 actualVersion, uint256 actualGasLimit, OptionsLib.NativeDrop[] memory actualNativeDrops) = libHarness.decodeOptions(data);

        assertEq(actualVersion, version);
        assertEq(actualGasLimit, gasLimit);
        assertEq(actualNativeDrops.length, nativeDrops.length);
        assertEq(actualNativeDrops[0].recipient, nativeDrops[0].recipient);
        assertEq(actualNativeDrops[0].amount, nativeDrops[0].amount);
    }
}
