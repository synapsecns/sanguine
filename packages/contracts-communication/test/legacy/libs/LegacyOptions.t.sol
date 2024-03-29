// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyOptionsLibHarness, LegacyOptionsLib} from "../harnesses/LegacyOptionsLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract LegacyOptionsLibTest is Test {
    bytes public constant LEGACY_OPTIONS_FIXTURE =
        hex"000100000000000000000000000000000000000000000000000000000000003d0900";
    uint256 public constant GAS_LIMIT_FIXTURE = 4_000_000;

    LegacyOptionsLibHarness public libHarness;

    function setUp() public {
        libHarness = new LegacyOptionsLibHarness();
    }

    function expectRevertInvalidOptions(bytes memory legacyOpts) internal {
        vm.expectRevert(abi.encodeWithSelector(LegacyOptionsLib.LegacyOptionsLib__InvalidOptions.selector, legacyOpts));
    }

    function test_encodeLegacyOptions() public {
        bytes memory encoded = libHarness.encodeLegacyOptions(GAS_LIMIT_FIXTURE);
        assertEq(encoded, LEGACY_OPTIONS_FIXTURE);
    }

    function test_decodeLegacyOptions() public {
        uint256 gasLimit = libHarness.decodeLegacyOptions(LEGACY_OPTIONS_FIXTURE);
        assertEq(gasLimit, GAS_LIMIT_FIXTURE);
    }

    function test_encodeLegacyOptionsRoundtrip(uint256 gasLimit) public {
        bytes memory encoded = libHarness.encodeLegacyOptions(gasLimit);
        uint256 newGasLimit = libHarness.decodeLegacyOptions(encoded);
        assertEq(newGasLimit, gasLimit);
    }

    function test_decodeLegacyOptions_revert_invalidVersion(uint16 version) public {
        vm.assume(version != LegacyOptionsLib.LEGACY_OPTIONS_VERSION);
        bytes memory invalidOpts = abi.encodePacked(version, uint256(1));
        expectRevertInvalidOptions(invalidOpts);
        libHarness.decodeLegacyOptions(invalidOpts);
    }

    function test_decodeLegacyOptions_revert_incorrectLength(bytes memory invalidOpts) public {
        vm.assume(invalidOpts.length != LEGACY_OPTIONS_FIXTURE.length);
        expectRevertInvalidOptions(invalidOpts);
        libHarness.decodeLegacyOptions(invalidOpts);
    }
}
