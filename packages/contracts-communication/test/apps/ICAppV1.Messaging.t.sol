// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainAppV1MessagingTest} from "./InterchainAppV1.Messaging.t.sol";
import {IInterchainAppV1Harness} from "../interfaces/IInterchainAppV1Harness.sol";
import {ICAppV1Harness} from "../harnesses/ICAppV1Harness.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

contract ICAppV1MessagingTest is InterchainAppV1MessagingTest {
    /// @dev This should deploy the Interchain App V1 contract and give `governor`
    /// privileges to setup its interchain configuration.
    function deployICAppV1() internal override returns (IInterchainAppV1Harness app) {
        app = new ICAppV1Harness(address(this));
        IAccessControl(address(app)).grantRole(IC_GOVERNOR_ROLE, governor);
    }
}
