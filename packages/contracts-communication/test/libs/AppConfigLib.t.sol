// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Test} from "forge-std/Test.sol";

import {AppConfigLib, AppConfigV1, AppConfigLibHarness} from "../harnesses/AppConfigLibHarness.sol";

contract AppConfigLibTest is Test {
    struct MockAppConfigV2 {
        uint256 requiredResponses;
        uint256 optimisticPeriod;
        bytes32 newField;
    }

    AppConfigLibHarness public libHarness;

    function setUp() public {
        libHarness = new AppConfigLibHarness();
    }

    function test_encodeVersionedAppConfigRoundtrip(uint8 version, bytes memory appConfig) public {
        bytes memory encoded = libHarness.encodeVersionedAppConfig(version, appConfig);
        (uint8 newVersion, bytes memory newAppConfig) = libHarness.decodeVersionedAppConfig(encoded);
        assertEq(newVersion, version);
        assertEq(newAppConfig, appConfig);
    }

    function test_encodeAppConfigV1Roundtrip(AppConfigV1 memory appConfig) public {
        bytes memory encoded = libHarness.encodeAppConfigV1(appConfig);
        AppConfigV1 memory decoded = libHarness.decodeAppConfigV1(encoded);
        assertEq(decoded.requiredResponses, appConfig.requiredResponses);
        assertEq(decoded.optimisticPeriod, appConfig.optimisticPeriod);
    }

    function test_decodeAppConfigV1_decodesV2(MockAppConfigV2 memory appConfig) public {
        bytes memory encoded =
            libHarness.encodeVersionedAppConfig(AppConfigLib.APP_CONFIG_V1 + 1, abi.encode(appConfig));
        AppConfigV1 memory decoded = libHarness.decodeAppConfigV1(encoded);
        assertEq(decoded.requiredResponses, appConfig.requiredResponses);
        assertEq(decoded.optimisticPeriod, appConfig.optimisticPeriod);
    }

    function test_decodeAppConfigV1_revertLowerVersion() public {
        AppConfigV1 memory appConfig = AppConfigV1(3, 100);
        uint8 incorrectVersion = AppConfigLib.APP_CONFIG_V1 - 1;
        bytes memory encoded = libHarness.encodeVersionedAppConfig(incorrectVersion, abi.encode(appConfig));
        vm.expectRevert(abi.encodeWithSelector(AppConfigLib.AppConfigLib__IncorrectVersion.selector, incorrectVersion));
        libHarness.decodeAppConfigV1(encoded);
    }
}
