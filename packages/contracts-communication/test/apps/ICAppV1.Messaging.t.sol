// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AppConfigV1, APP_CONFIG_GUARD_DISABLED} from "../../contracts/libs/AppConfig.sol";

import {InterchainAppV1MessagingTest} from "./InterchainAppV1.Messaging.t.sol";
import {IInterchainAppV1Harness} from "../interfaces/IInterchainAppV1Harness.sol";
import {ICAppV1Harness} from "../harnesses/ICAppV1Harness.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract ICAppV1MessagingTest is InterchainAppV1MessagingTest {
    /// @dev This should deploy the Interchain App V1 contract and give `governor`
    /// privileges to setup its interchain configuration.
    function deployICAppV1() internal override returns (IInterchainAppV1Harness app) {
        app = new ICAppV1Harness(address(this));
        IAccessControl(address(app)).grantRole(IC_GOVERNOR_ROLE, governor);
    }

    function test_freshAppConfig() public {
        ICAppV1Harness app = new ICAppV1Harness(address(this));
        AppConfigV1 memory config = app.getAppConfigV1();
        assertEq(config.requiredResponses, 0);
        assertEq(config.optimisticPeriod, 0);
        assertEq(config.guardFlag, APP_CONFIG_GUARD_DISABLED);
        assertEq(config.guard, address(0));
    }

    function test_freshAppModules() public {
        ICAppV1Harness app = new ICAppV1Harness(address(this));
        assertEq(app.getModules(), new address[](0));
    }
}
