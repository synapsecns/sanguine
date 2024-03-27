// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AppConfigV1, AppConfigLib} from "../../contracts/libs/AppConfig.sol";

contract AppConfigLibHarness {
    function encodeVersionedAppConfig(uint8 version, bytes calldata appConfig) external pure returns (bytes memory) {
        return AppConfigLib.encodeVersionedAppConfig(version, appConfig);
    }

    function decodeVersionedAppConfig(bytes calldata data) external pure returns (uint8, bytes memory) {
        (uint8 version, bytes memory appConfig) = AppConfigLib.decodeVersionedAppConfig(data);
        return (version, appConfig);
    }

    function encodeAppConfigV1(AppConfigV1 memory appConfig) external pure returns (bytes memory) {
        return AppConfigLib.encodeAppConfigV1(appConfig);
    }

    function decodeAppConfigV1(bytes calldata data) external pure returns (AppConfigV1 memory) {
        return AppConfigLib.decodeAppConfigV1(data);
    }
}
