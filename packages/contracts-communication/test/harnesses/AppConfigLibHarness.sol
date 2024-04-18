// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AppConfigV1, AppConfigLib} from "../../contracts/libs/AppConfig.sol";

contract AppConfigLibHarness {
    function decodeAppConfigV1(bytes calldata data) external view returns (AppConfigV1 memory) {
        return AppConfigLib.decodeAppConfigV1(data);
    }

    function encodeAppConfigV1(AppConfigV1 memory appConfig) external pure returns (bytes memory) {
        return AppConfigLib.encodeAppConfigV1(appConfig);
    }
}
