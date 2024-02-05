// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20} from "../../src/tokens/InterchainERC20.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract InterchainERC20Test is Test {
    InterchainERC20 public token;

    address public admin;
    address public emergencyPauser;
    address public governor;

    function setUp() public virtual {
        admin = makeAddr("Admin");
        emergencyPauser = makeAddr("EmergencyPauser");
        governor = makeAddr("Governor");

        token = new InterchainERC20("Token Name", "Token Symbol", admin);
        vm.startPrank(admin);
        token.grantRole(token.EMERGENCY_PAUSER_ROLE(), emergencyPauser);
        token.grantRole(token.GOVERNOR_ROLE(), governor);
        vm.stopPrank();
    }

    function authSetTotalBurnLimit(address bridge, uint256 limit) public {
        vm.prank(governor);
        token.setTotalBurnLimit(bridge, limit);
    }

    function authSetTotalMintLimit(address bridge, uint256 limit) public {
        vm.prank(governor);
        token.setTotalMintLimit(bridge, limit);
    }

    function authPause() public {
        vm.prank(emergencyPauser);
        token.pause();
    }

    function authUnpause() public {
        vm.prank(emergencyPauser);
        token.unpause();
    }

    function test_constructor() public {
        assertEq(token.name(), "Token Name");
        assertEq(token.symbol(), "Token Symbol");
    }
}
