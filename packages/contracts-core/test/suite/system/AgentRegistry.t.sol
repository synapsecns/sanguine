// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentRegistryTools } from "../../tools/system/AgentRegistryTools.t.sol";
import { EnumerableSetTools } from "../../tools/libs/EnumerableSetTools.t.sol";

import { AgentRegistryHarness } from "../../harnesses/system/AgentRegistryHarness.t.sol";

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

// solhint-disable func-name-mixedcase
contract AgentRegistryTest is AgentRegistryTools, EnumerableSetTools {
    AgentRegistryHarness internal registry;
    uint32[] internal domainsExtended;
    address[] internal allAgents;

    uint256 internal constant AGENTS_PER_DOMAIN =
        GUARDS < NOTARIES_PER_CHAIN ? GUARDS : NOTARIES_PER_CHAIN;

    function setUp() public override {
        super.setUp();
        registry = new AgentRegistryHarness();
        domainsExtended.push(0);
        for (uint256 d = 0; d < DOMAINS; ++d) {
            domainsExtended.push(domains[d]);
        }
        for (uint256 d = 0; d < DOMAINS + 1; ++d) {
            uint32 domain = domainsExtended[d];
            for (uint256 i = 0; i < AGENTS_PER_DOMAIN; ++i) {
                allAgents.push(suiteAgent(domain, i));
            }
        }
        expectedStates = new uint256[][](ELEMENTS);
        expectedStates[0] = [1, 2, 3]; // (2) is removed
        expectedStates[1] = [1, 3]; // (0) is removed (but is not stored)
        expectedStates[2] = [1, 3]; // (1) is removed
        expectedStates[3] = [3]; // (3) is removed
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TESTS: SINGLE AGENT                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent(uint32 domain, address account) public {
        // Agent will not be added twice
        bool toBeAdded = !registry.isActiveAgent(domain, account);
        // Check if it is the first agent for the domain
        bool isFirst = registry.amountAgents(domain) == 0;
        if (toBeAdded) expectAgentAdded(domain, account, isFirst);
        assertEq(registry.addAgent(domain, account), toBeAdded, "!addAgent: return value");
        assertTrue(registry.isActiveAgent(domain, account), "!addAgent: not added (domain)");
        (bool isActive, uint32 _domain) = registry.isActiveAgent(account);
        assertTrue(isActive, "!isActiveAgent: isActive");
        assertEq(_domain, domain, "!isActiveAgent: domain");
    }

    function test_removeAgent(uint32 domain, address account) public {
        (bool wasActive, uint32 initialDomain) = registry.isActiveAgent(account);
        // Agent will not be removed twice
        bool toBeRemoved = registry.isActiveAgent(domain, account);
        // Check if it is the last agent for the domain
        bool isLast = registry.amountAgents(domain) == 1;
        if (toBeRemoved) expectAgentRemoved(domain, account, isLast);
        assertEq(registry.removeAgent(domain, account), toBeRemoved, "!removeAgent: return value");
        assertFalse(registry.isActiveAgent(domain, account), "!removeAgent: not removed (domain)");
        (bool isActive, uint32 _domain) = registry.isActiveAgent(account);
        assertEq(isActive, wasActive && !toBeRemoved, "!removeAgent: isActive (global)");
        assertEq(_domain, toBeRemoved ? 0 : initialDomain, "!removeAgent: domain (global)");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        TESTS: MULTIPLE AGENTS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_suiteAgents() public {
        for (uint256 d = 0; d < domainsExtended.length; ++d) {
            uint32 domain = domainsExtended[d];
            for (uint256 i = 0; i < AGENTS_PER_DOMAIN; ++i) {
                test_addAgent({ domain: domain, account: suiteAgent(domain, i) });
            }
        }
    }

    function test_removeAgent_suiteAgents() public {
        // Add suite Agents first
        test_addAgent_suiteAgents();
        // Remove them one by one
        for (uint256 d = 0; d < domainsExtended.length; ++d) {
            uint32 domain = domainsExtended[d];
            for (uint256 i = 0; i < AGENTS_PER_DOMAIN; ++i) {
                test_removeAgent({ domain: domain, account: suiteAgent(domain, i) });
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: DUPLICATE AGENTS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_onTwoDomains() public {
        // Add suite Agents first
        test_addAgent_suiteAgents();
        // Iterate over all domains, plus zero domain for guards
        for (uint256 d = 0; d < domainsExtended.length; ++d) {
            uint32 domain = domainsExtended[d];
            for (uint256 i = 0; i < AGENTS_PER_DOMAIN; ++i) {
                address agent = suiteAgent(domain, i);
                for (uint256 f = 0; f < domainsExtended.length; ++f) {
                    uint32 domainForeign = domainsExtended[f];
                    if (domain == domainForeign) continue;
                    // Should not be able to add active Agent to another domain
                    assertFalse(
                        registry.addAgent(domainForeign, agent),
                        "!addAgent: added on two domains"
                    );
                }
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           TESTS: MODIFIERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_haveActiveNotary(uint32 domain) public {
        vm.assume(domain != 0);
        address agent = address(1);
        test_addAgent(domain, agent);
        // Should not revert with an active agent
        registry.onlyActiveNotary(domain);
        test_removeAgent(domain, agent);
        // Should revert with no active agents
        vm.expectRevert("No active notaries");
        registry.onlyActiveNotary(domain);
    }

    function test_haveActiveNotary_zeroDomain() public {
        uint32 domain = 0;
        address agent = address(1);
        test_addAgent(domain, agent);
        // Should revert, as this is not a Notary
        vm.expectRevert("No active notaries");
        registry.onlyActiveNotary(domain);
    }

    function test_haveActiveGuard() public {
        uint32 domain = 0;
        address agent = address(1);
        test_addAgent(domain, agent);
        // Should not revert with an active agent
        registry.onlyActiveGuard();
        test_removeAgent(domain, agent);
        // Should revert with no active agents
        vm.expectRevert("No active guards");
        registry.onlyActiveGuard();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           TESTS: SIGNATURE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_checkAgentAuth() public {
        // Add suite Agents first
        test_addAgent_suiteAgents();
        // Check signatures for all the added agents
        for (uint256 i = 0; i < allAgents.length; ++i) {
            _checkSignatures(allAgents[i]);
        }
        // Remove agents one by one and check all signatures
        for (uint256 d = 0; d < domainsExtended.length; ++d) {
            uint32 domain = domainsExtended[d];
            for (uint256 i = 0; i < AGENTS_PER_DOMAIN; ++i) {
                address agent = suiteAgent(domain, i);
                test_removeAgent(domain, agent);
                for (uint256 j = 0; j < allAgents.length; ++j) {
                    _checkSignatures(allAgents[j]);
                }
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: RESET                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_resetAgents() public {
        // Add suite Agents first
        test_addAgent_suiteAgents();
        assertEq(registry.currentEpoch(), 0, "!currentEpoch: before reset");
        registry.removeAllAgents();
        assertEq(registry.currentEpoch(), 1, "!currentEpoch: after reset");
        for (uint256 i = 0; i < allAgents.length; ++i) {
            address agent = allAgents[i];
            (bool isActive, uint32 _domain) = registry.isActiveAgent(agent);
            assertFalse(isActive, "!isActiveAgent(account): isActive after reset");
            assertEq(_domain, 0, "!isActiveAgent(account): domain after reset");
            for (uint256 d = 0; d < domainsExtended.length; ++d) {
                uint32 domain = domainsExtended[d];
                assertFalse(
                    registry.isActiveAgent(domain, agent),
                    "!isActiveAgent(domain, account): after reset"
                );
            }
        }
        _checkAgentViews(0);
    }

    function test_resetAgents_addSuiteAgents() public {
        test_resetAgents();
        test_addAgent_suiteAgents();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: IGNORE MODE                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_ignoreAgent(uint32 domain, address account) public {
        registry.setIgnoredAgent(domain, account);
        // Should not add agent in ignore mode
        registry.toggleIgnoreMode(true);
        assertFalse(registry.addAgent(domain, account), "!addAgent: added");
        (bool isActive, uint32 _domain) = registry.isActiveAgent(account);
        assertFalse(isActive, "!isActiveAgent(account): isActive added");
        assertEq(_domain, 0, "!isActiveAgent(account): domain added");
        assertFalse(
            registry.isActiveAgent(domain, account),
            "!isActiveAgent(domain, account): added"
        );
        if (domain != 0) assertFalse(registry.isActiveDomain(domain), "!isActiveDomain: added");
        // Should add agent w/o ignore mode
        registry.toggleIgnoreMode(false);
        test_addAgent(domain, account);
        // Should not remove agent in ignore mode
        registry.toggleIgnoreMode(true);
        assertFalse(registry.removeAgent(domain, account), "!removeAgent: removed");
        (isActive, _domain) = registry.isActiveAgent(account);
        assertTrue(isActive, "!isActiveAgent(account): isActive removed");
        assertEq(_domain, domain, "!isActiveAgent(account): domain removed");
        assertTrue(
            registry.isActiveAgent(domain, account),
            "!isActiveAgent(domain, account): removed"
        );
        if (domain != 0) assertTrue(registry.isActiveDomain(domain), "!isActiveDomain: removed");
        // Should remove agent w/o ignore mode
        registry.toggleIgnoreMode(false);
        test_removeAgent(domain, account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_agentViews() public {
        test_addAgent_suiteAgents();
        _checkAgentViews(NOTARIES_PER_CHAIN);
        test_removeAgent_suiteAgents();
        _checkAgentViews(0);
    }

    // solhint-disable-next-line code-complexity
    function test_domainViews() public {
        uint256 amount = domainsExtended.length;
        // Sanity check: test is setup correctly
        require(amount == ELEMENTS, "!EnumerableSetTools setup");
        for (uint256 d = 0; d < amount; ++d) {
            uint32 domain = domainsExtended[d];
            test_addAgent(domain, suiteAgent(domain, 0));
            for (uint256 i = 0; i < amount; ++i) {
                uint32 domainToCheck = domainsExtended[i];
                // Domain should be active only if a Notary was added there
                assertEq(
                    registry.isActiveDomain(domainToCheck),
                    domainToCheck != 0 && i <= d,
                    "!isActiveDomain"
                );
            }
            // Zero domain in stored first in domainsExtended,
            // so amount of active domains should be `d`
            assertEq(registry.amountDomains(), d, "!amountDomains: adding[0]");
            test_addAgent(domain, suiteAgent(domain, 1));
            assertEq(registry.amountDomains(), d, "!amountDomains: adding[1]");
        }
        for (uint256 d = 0; d < amount; ++d) {
            // Check getters
            uint32[] memory values = registry.allDomains();
            assertEq(values.length, expectedStates[d].length, "!values.length");
            for (uint256 i = 0; i < values.length; ++i) {
                uint32 expectedDomain = domainsExtended[expectedStates[d][i]];
                assertEq(values[i], expectedDomain, "!values");
                assertEq(registry.getDomain(i), expectedDomain, "getDomain");
            }
            // Remove all agents from the predetermined domain
            uint256 index = removalOrder[d];
            uint32 domain = domainsExtended[index];
            test_removeAgent(domain, suiteAgent(domain, 0));
            if (domain == 0) {
                assertFalse(registry.isActiveDomain(domain), "!isActiveDomain: zero domain");
            } else {
                assertTrue(registry.isActiveDomain(domain), "!isActiveDomain: 1 agent left");
            }
            test_removeAgent(domain, suiteAgent(domain, 1));
            if (domain == 0) {
                assertFalse(registry.isActiveDomain(domain), "!isActiveDomain: zero domain");
            } else {
                assertFalse(registry.isActiveDomain(domain), "!isActiveDomain: 0 agents left");
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks AgentRegistry agent-related views. It is assumed that
     * exactly `activeAgents` is active for every domain.
     */
    function _checkAgentViews(uint256 activeAgents) internal {
        for (uint256 d = 0; d < domainsExtended.length; ++d) {
            uint32 domain = domainsExtended[d];
            assertEq(registry.amountAgents(domain), activeAgents, "!amountAgents");
            address[] memory agents = registry.allAgents(domain);
            assertEq(agents.length, activeAgents, "!agents.length");
            // Check active agents
            for (uint256 i = 0; i < activeAgents; ++i) {
                address agent = suiteAgent(domain, i);
                assertEq(agents[i], agent, "!agents");
                assertEq(registry.getAgent(domain, i), agent, "!getAgent");
                (bool isActive, uint32 _domain) = registry.isActiveAgent(agent);
                assertTrue(isActive, "!isActiveAgent(account): isActive should be active");
                assertEq(_domain, domain, "!isActiveAgent(account): domain should be active");
                assertTrue(
                    registry.isActiveAgent(domain, agent),
                    "!isActiveAgent(account): should be active"
                );
            }
            // Check remaining agents: should be inactive
            for (uint256 i = activeAgents; i < AGENTS_PER_DOMAIN; ++i) {
                address agent = suiteAgent(domain, i);
                (bool isActive, uint32 _domain) = registry.isActiveAgent(agent);
                assertFalse(isActive, "!isActiveAgent(account): isActive should be inactive");
                assertEq(_domain, 0, "!isActiveAgent(account): domain should be inactive");
                assertFalse(
                    registry.isActiveAgent(domain, agent),
                    "!isActiveAgent(account): should be inactive"
                );
            }
        }
    }

    function _checkSignatures(address account) internal {
        bytes memory data = "test";
        bytes32 digest = ECDSA.toEthSignedMessageHash(keccak256(data));
        bytes memory signature = signMessage(account, data);
        for (uint256 d = 0; d < domainsExtended.length; ++d) {
            uint32 domain = domainsExtended[d];
            if (registry.isActiveAgent(domain, account)) {
                assertEq(
                    registry.checkAgentAuth(domain, digest, signature),
                    account,
                    "!checkAgentAuth"
                );
            } else {
                vm.expectRevert("Signer is not authorized");
                registry.checkAgentAuth(domain, digest, signature);
            }
        }
    }
}
