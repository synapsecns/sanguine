// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../tools/NotaryManagerTools.t.sol";

// solhint-disable func-name-mixedcase

contract NotaryManagerTest is NotaryManagerTools {
    function test_setup() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            assertEq(suiteNotaryManager(domain).owner(), owner, "!owner");
            assertEq(suiteNotaryManager(domain).origin(), address(suiteOrigin(domain)), "!origin");
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                  TESTS: RESTRICTED ACCESS (REVERTS)                  ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setOrigin_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            expectRevertNotOwner();
            vm.prank(caller);
            suiteNotaryManager(domain).setOrigin(address(this));
        }
    }

    function test_setOrigin_revert_notContract() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            vm.expectRevert("!contract origin");
            vm.prank(owner);
            suiteNotaryManager(domain).setOrigin(owner);
        }
    }

    function test_setNotary_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            expectRevertNotOwner();
            vm.prank(caller);
            suiteNotaryManager(domain).setNotary(address(0));
        }
    }

    function test_slashNotary_revert_notOrigin(address caller) public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            vm.assume(caller != address(suiteOrigin(domain)));
            vm.expectRevert("!origin");
            vm.prank(caller);
            suiteNotaryManager(domain).slashNotary(payable(address(0)));
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: RESTRICTED ACCESS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setOrigin() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            expectNewOrigin(address(this));
            vm.prank(owner);
            suiteNotaryManager(domain).setOrigin(address(this));
            assertEq(
                suiteNotaryManager(domain).origin(),
                address(this),
                "Failed to set new origin"
            );
        }
    }

    function test_setNotary() public {
        address newNotary = address(1234);
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            expectNewNotary(newNotary);
            vm.prank(owner);
            suiteNotaryManager(domain).setNotary(newNotary);
            assertEq(suiteNotaryManager(domain).notary(), newNotary, "!newNotary: notaryManager");
            assertTrue(suiteOrigin(domain).isNotary(newNotary), "!newNotary: origin");
        }
    }

    function test_slashNotary() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            expectFakeSlashed(broadcaster);
            vm.prank(address(suiteOrigin(domain)));
            suiteNotaryManager(domain).slashNotary(payable(broadcaster));
        }
    }
}
