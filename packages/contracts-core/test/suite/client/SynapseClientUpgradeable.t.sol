// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./SynapseClient.t.sol";
import {
    SynapseClientUpgradeableHarness
} from "../../harnesses/client/SynapseClientUpgradeableHarness.t.sol";

import {
    TransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// solhint-disable func-name-mixedcase
contract SynapseClientUpgradeableTest is SynapseClientTest {
    function deployBasicClient(uint32 domain, uint32 optimisticSeconds)
        public
        virtual
        override
        returns (BasicClient)
    {
        SynapseClientUpgradeableHarness impl = new SynapseClientUpgradeableHarness(
            address(suiteOrigin(domain)),
            address(suiteDestination(domain)),
            optimisticSeconds
        );
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(impl),
            proxyAdmin,
            bytes("")
        );
        SynapseClientUpgradeableHarness client = SynapseClientUpgradeableHarness(address(proxy));
        client.initialize();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 anotherDomain = domains[d];
            if (anotherDomain != domain) {
                client.setTrustedSender(anotherDomain, mockTrustedSender(anotherDomain));
            }
        }
        client.transferOwnership(owner);
        return BasicClient(address(proxy));
    }

    function test_setTrustedSender_revert_notOwner(address caller) public override {
        // Exclude calls from proxy admin
        vm.assume(caller != proxyAdmin);
        super.test_setTrustedSender_revert_notOwner(caller);
    }

    function test_setTrustedSenders_revert_notOwner(address caller) public override {
        // Exclude calls from proxy admin
        vm.assume(caller != proxyAdmin);
        super.test_setTrustedSenders_revert_notOwner(caller);
    }

    // Remaining tests are inherited from SynapseClientTest
}
