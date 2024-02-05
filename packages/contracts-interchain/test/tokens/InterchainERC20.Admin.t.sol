// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20, InterchainERC20Test} from "./InterchainERC20.t.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainERC20AdminTest is InterchainERC20Test {
    function test_grantRole_revert_notAdmin(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != admin);
        bytes32 role = token.EMERGENCY_PAUSER_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, 0)
        );
        token.grantRole(role, address(1337));
    }

    function test_setTotalBurnLimit_revert_notGovernor(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != governor);
        bytes32 requiredRole = token.GOVERNOR_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.setTotalBurnLimit(address(1337), 1);
    }

    function test_setTotalMintLimit_revert_notGovernor(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != governor);
        bytes32 requiredRole = token.GOVERNOR_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.setTotalMintLimit(address(1337), 1);
    }

    function test_pause_revert_notEmergencyPauser(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != emergencyPauser);
        bytes32 requiredRole = token.EMERGENCY_PAUSER_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.pause();
    }

    function test_unpause_revert_notEmergencyPauser(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != emergencyPauser);
        authPause();
        bytes32 requiredRole = token.EMERGENCY_PAUSER_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.unpause();
    }
}
