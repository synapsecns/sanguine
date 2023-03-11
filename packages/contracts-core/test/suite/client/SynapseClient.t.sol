// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./Client.t.sol";
import { SynapseClient } from "../../../contracts/client/SynapseClient.sol";
import { SynapseClientHarness } from "../../harnesses/client/SynapseClientHarness.t.sol";

// solhint-disable func-name-mixedcase
contract SynapseClientTest is ClientTest {
    function deployBasicClient(uint32 domain, uint32 optimisticSeconds)
        public
        virtual
        override
        returns (BasicClient)
    {
        SynapseClient client = new SynapseClientHarness(
            address(suiteOrigin(domain)),
            address(suiteDestination(domain)),
            optimisticSeconds
        );
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 anotherDomain = domains[d];
            if (anotherDomain != domain) {
                client.setTrustedSender(anotherDomain, mockTrustedSender(anotherDomain));
            }
        }
        client.transferOwnership(owner);
        return BasicClient(address(client));
    }

    // Most tests are inherited from ClientTest

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TESTS: OWNER ONLY (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setTrustedSender_revert_notOwner(address caller) public virtual {
        vm.assume(caller != owner);
        SynapseClient client = suiteSynapseClient(DOMAIN_LOCAL);
        expectRevertNotOwner();
        vm.prank(caller);
        client.setTrustedSender(0, bytes32(0));
    }

    function test_setTrustedSenders_revert_notOwner(address caller) public virtual {
        vm.assume(caller != owner);
        SynapseClient client = suiteSynapseClient(DOMAIN_LOCAL);
        uint32[] memory testDomains = new uint32[](1);
        bytes32[] memory testSenders = new bytes32[](1);
        expectRevertNotOwner();
        vm.prank(caller);
        client.setTrustedSenders(testDomains, testSenders);
    }

    function test_setTrustedSenders_revert_mismatchedArrays() public {
        SynapseClient client = suiteSynapseClient(DOMAIN_LOCAL);
        uint32[] memory testDomains = new uint32[](1);
        bytes32[] memory testSenders = new bytes32[](2);
        vm.expectRevert("!arrays");
        vm.prank(owner);
        client.setTrustedSenders(testDomains, testSenders);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: OWNER ONLY                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setTrustedSender() public {
        SynapseClient client = suiteSynapseClient(DOMAIN_LOCAL);
        bytes32 testSender = "very trusted sender";
        vm.prank(owner);
        client.setTrustedSender(DOMAIN_REMOTE, testSender);
        assertEq(client.trustedSender(DOMAIN_REMOTE), testSender, "Failed to set trusted sender");
    }

    function test_setTrustedSenders() public {
        SynapseClient client = suiteSynapseClient(DOMAIN_LOCAL);
        uint256 amount = 10;
        uint32[] memory testDomains = new uint32[](amount);
        bytes32[] memory testSenders = new bytes32[](amount);
        for (uint32 i = 0; i < amount; ++i) {
            testDomains[i] = 69 + 42 * i;
            testSenders[i] = mockTrustedSender(i + 1);
        }
        vm.prank(owner);
        client.setTrustedSenders(testDomains, testSenders);
        for (uint32 i = 0; i < amount; ++i) {
            assertEq(
                client.trustedSender(testDomains[i]),
                testSenders[i],
                "Failed to set trusted sender"
            );
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            CHAIN GETTERS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function suiteSynapseClient(uint32 domain) public view returns (SynapseClient) {
        return SynapseClient(address(basicClients[domain]));
    }
}
