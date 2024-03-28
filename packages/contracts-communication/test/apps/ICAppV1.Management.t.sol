// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainAppV1ManagementTest} from "./InterchainAppV1.Management.t.sol";
import {IInterchainAppV1Harness} from "../interfaces/IInterchainAppV1Harness.sol";
import {ICAppV1Harness} from "../harnesses/ICAppV1Harness.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

contract ICAppV1ManagementTest is InterchainAppV1ManagementTest {
    /// @dev This should deploy the Interchain App V1 contract and give `governor`
    /// privileges to setup its interchain configuration.
    function deployICAppV1() internal override returns (IInterchainAppV1Harness app) {
        app = new ICAppV1Harness(address(this));
        IAccessControl(address(app)).grantRole(IC_GOVERNOR_ROLE, governor);
    }

    function expectRevertUnauthorizedGovernor(address caller) internal override {
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, IC_GOVERNOR_ROLE)
        );
    }
}
