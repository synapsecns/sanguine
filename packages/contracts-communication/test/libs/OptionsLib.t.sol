// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Test} from "forge-std/Test.sol";

import {OptionsLib, OptionsLibHarness, OptionsV1} from "../harnesses/OptionsLibHarness.sol";

contract OptionsLibTest is Test {
    struct MockOptionsV2 {
        uint256 gasLimit;
        uint256 gasAirdrop;
        bytes32 newField;
    }

    OptionsLibHarness public libHarness;

    function setUp() public {
        libHarness = new OptionsLibHarness();
    }

    function test_encodeOptions() public {
        uint8 version = 1;
        // 200k gas limit
        uint256 gasLimit = 200_000;
        // 100k wei
        uint256 gasAirdrop = 100_000;
        bytes memory expected = abi.encode(version, gasLimit, gasAirdrop);
        bytes memory actual = libHarness.encodeOptions(version, gasLimit, gasAirdrop);
        assertEq(actual, expected);
    }

    function test_decodeOptions() public {
        uint8 version = 1;
        // 200k gas limit
        uint256 gasLimit = 200_000;
        // 100k wei
        uint256 gasAirdrop = 100_000;
        bytes memory data = abi.encode(version, gasLimit, gasAirdrop);
        (uint8 actualVersion, uint256 actualGasLimit, uint256 actualGasAirdrop) = libHarness.decodeOptions(data);
        assertEq(actualVersion, version);
        assertEq(actualGasLimit, gasLimit);
        assertEq(actualGasAirdrop, gasAirdrop);
    }

    function test_encodeVersionedOptionsRoundtrip(uint8 version, bytes memory options) public {
        bytes memory encoded = libHarness.encodeVersionedOptions(version, options);
        (uint8 newVersion, bytes memory newOptions) = libHarness.decodeVersionedOptions(encoded);
        assertEq(newVersion, version);
        assertEq(newOptions, options);
    }

    function test_encodeOptionsV1Roundtrip(OptionsV1 memory options) public {
        bytes memory encoded = libHarness.encodeOptionsV1(options);
        OptionsV1 memory decoded = libHarness.decodeOptionsV1(encoded);
        assertEq(decoded.gasLimit, options.gasLimit);
        assertEq(decoded.gasAirdrop, options.gasAirdrop);
    }

    function test_decodeOptionsV1_decodesV2(MockOptionsV2 memory options) public {
        bytes memory encoded = libHarness.encodeVersionedOptions(OptionsLib.OPTIONS_V1 + 1, abi.encode(options));
        OptionsV1 memory decoded = libHarness.decodeOptionsV1(encoded);
        assertEq(decoded.gasLimit, options.gasLimit);
        assertEq(decoded.gasAirdrop, options.gasAirdrop);
    }

    function test_decodeOptionsV1_revertLowerVersion() public {
        OptionsV1 memory options = OptionsV1(200_000, 100_000);
        uint8 incorrectVersion = OptionsLib.OPTIONS_V1 - 1;
        bytes memory encoded = libHarness.encodeVersionedOptions(incorrectVersion, abi.encode(options));
        vm.expectRevert(abi.encodeWithSelector(OptionsLib.OptionsLib__IncorrectVersion.selector, incorrectVersion));
        libHarness.decodeOptionsV1(encoded);
    }
}
