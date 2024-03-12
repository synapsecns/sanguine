// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20, InterchainERC20Harness} from "../harnesses/InterchainERC20Harness.sol";
import {MockInterchainFactory} from "../mocks/MockInterchainFactory.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract InterchainERC20Test is Test {
    InterchainERC20Harness public token;
    MockInterchainFactory public factory;

    address public admin;
    address public emergencyPauser;
    address public governor;
    address public processor;

    function setUp() public virtual {
        admin = makeAddr("Admin");
        emergencyPauser = makeAddr("EmergencyPauser");
        governor = makeAddr("Governor");
        processor = makeAddr("Processor");

        factory = new MockInterchainFactory();
        address deployedHarness =
            factory.deployInterchainTokenHarness("Token Name", "Token Symbol", 18, admin, processor);
        token = InterchainERC20Harness(deployedHarness);
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
        assertEq(token.PROCESSOR(), processor);
        assertEq(token.decimals(), 18);
    }

    function test_differentDecimals() public {
        assertEq(InterchainERC20(factory.deployInterchainToken("A", "B", 6, admin, processor)).decimals(), 6);
        assertEq(InterchainERC20(factory.deployInterchainToken("A", "B", 32, admin, processor)).decimals(), 32);
    }

    function test_revert_adminZero() public {
        vm.expectRevert(InterchainERC20.InterchainERC20__AdminZero.selector);
        factory.deployInterchainToken("A", "B", 18, address(0), processor);
    }
}
