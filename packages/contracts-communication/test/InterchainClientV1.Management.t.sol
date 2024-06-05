// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1BaseTest} from "./InterchainClientV1.Base.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainClientV1ManagementTest is InterchainClientV1BaseTest {
    function test_constructor() public view {
        assertEq(icClient.INTERCHAIN_DB(), icDB);
        assertEq(icClient.owner(), owner);
    }

    function test_setDefaultGuard() public {
        expectEventDefaultGuardSet(defaultGuard);
        setDefaultGuard(defaultGuard);
        assertEq(icClient.defaultGuard(), defaultGuard);
    }

    function test_setDefaultGuard_revert_zeroAddress() public {
        expectRevertGuardZeroAddress();
        setDefaultGuard(address(0));
    }

    function test_setDefaultGuard_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        icClient.setDefaultGuard(defaultGuard);
    }

    function test_setDefaultModule() public {
        expectEventDefaultModuleSet(defaultModule);
        setDefaultModule(defaultModule);
        assertEq(icClient.defaultModule(), defaultModule);
    }

    function test_setDefaultModule_revert_zeroAddress() public {
        expectRevertModuleZeroAddress();
        setDefaultModule(address(0));
    }

    function test_setDefaultModule_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        icClient.setDefaultModule(defaultModule);
    }

    function test_setLinkedClient() public {
        expectEventLinkedClientSet(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        assertEq(icClient.getLinkedClient(REMOTE_CHAIN_ID), MOCK_REMOTE_CLIENT);
    }

    function test_setLinkedClient_success_zeroClient() public {
        expectEventLinkedClientSet(REMOTE_CHAIN_ID, 0);
        setLinkedClient(REMOTE_CHAIN_ID, 0);
        assertEq(icClient.getLinkedClient(REMOTE_CHAIN_ID), 0);
    }

    function test_setLinkedClient_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        icClient.setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
    }
}
