// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceV1Test} from "./SynapseExecutionServiceV1.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseExecutionServiceV1ProxyTest is SynapseExecutionServiceV1Test {
    function test_initialize() public {
        service.initialize(admin);
        assertTrue(service.hasRole(service.DEFAULT_ADMIN_ROLE(), admin));
    }

    function test_initialize_revert_calledTwice() public {
        service.initialize(admin);
        expectRevertInvalidInitialization();
        service.initialize(address(1));
    }

    function test_initialize_revert_calledImplementation() public {
        expectRevertInvalidInitialization();
        implementation.initialize(admin);
    }
}
