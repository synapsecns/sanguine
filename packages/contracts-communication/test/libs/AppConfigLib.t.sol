// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

import {AppConfigLib, AppConfigV1, AppConfigLibHarness} from "../harnesses/AppConfigLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract AppConfigLibTest is Test {
    struct MockAppConfigV2 {
        uint256 requiredResponses;
        uint256 optimisticPeriod;
        uint256 guardFlag;
        address guard;
        bytes32 newField;
    }

    AppConfigLibHarness public libHarness;

    function setUp() public {
        libHarness = new AppConfigLibHarness();
    }

    function test_encodeAppConfigV1Roundtrip(AppConfigV1 memory appConfig) public view {
        bytes memory encoded = libHarness.encodeAppConfigV1(appConfig);
        AppConfigV1 memory decoded = libHarness.decodeAppConfigV1(encoded);
        assertEq(decoded.requiredResponses, appConfig.requiredResponses);
        assertEq(decoded.optimisticPeriod, appConfig.optimisticPeriod);
        assertEq(decoded.guardFlag, appConfig.guardFlag);
        assertEq(decoded.guard, appConfig.guard);
    }

    function test_decodeAppConfigV1_decodesV2(MockAppConfigV2 memory appConfig) public view {
        bytes memory encoded =
            VersionedPayloadLib.encodeVersionedPayload(AppConfigLib.APP_CONFIG_V1 + 1, abi.encode(appConfig));
        AppConfigV1 memory decoded = libHarness.decodeAppConfigV1(encoded);
        assertEq(decoded.requiredResponses, appConfig.requiredResponses);
        assertEq(decoded.optimisticPeriod, appConfig.optimisticPeriod);
    }

    function test_decodeAppConfigV1_revertLowerVersion() public {
        AppConfigV1 memory appConfig = AppConfigV1(3, 100, 0, address(0));
        uint16 VersionInvalid = AppConfigLib.APP_CONFIG_V1 - 1;
        bytes memory encoded = VersionedPayloadLib.encodeVersionedPayload(VersionInvalid, abi.encode(appConfig));
        vm.expectRevert(abi.encodeWithSelector(AppConfigLib.AppConfigLib__VersionInvalid.selector, VersionInvalid));
        libHarness.decodeAppConfigV1(encoded);
    }
}
