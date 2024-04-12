// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1BaseTest} from "./InterchainClientV1.Base.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainClientV1ManagementTest is InterchainClientV1BaseTest {
    function test_constructor() public {
        assertEq(icClient.INTERCHAIN_DB(), icDB);
        assertEq(icClient.owner(), owner);
    }

    function test_setExecutionFees_emitsEvent() public {
        expectEventExecutionFeesSet(execFees);
        setExecutionFees(execFees);
    }

    function test_setExecutionFees_setsExecutionFees() public {
        setExecutionFees(execFees);
        assertEq(icClient.executionFees(), execFees);
    }

    function test_setExecutionFees_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        icClient.setExecutionFees(execFees);
    }

    function test_setLinkedClient_emitsEvent() public {
        expectEventLinkedClientSet(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
    }

    function test_setLinkedClient_setsLinkedClient() public {
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        assertEq(icClient.getLinkedClient(REMOTE_CHAIN_ID), MOCK_REMOTE_CLIENT);
    }

    function test_setLinkedClient_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        icClient.setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
    }
}
