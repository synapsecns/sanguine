// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";
import {AgentSetHarness} from "../../harnesses/libs/AgentSetHarness.t.sol";
import {EnumerableSetTools} from "../../tools/libs/EnumerableSetTools.t.sol";

// solhint-disable func-name-mixedcase
contract AgentSetTest is EnumerableSetTools, Test {
    struct agent {
        bool isActive;
        uint32 domain;
    }

    // Straightforward implementation of set to test against
    mapping(address => agent) internal agents;

    AgentSetHarness internal libHarness;
    uint32[] internal domains;

    uint256 internal constant AGENTS_PER_DOMAIN = 10;

    function setUp() public {
        libHarness = new AgentSetHarness();
        // Add a few domains to test
        domains.push(0); // Guards domain
        domains.push(1); // Generic domain
        domains.push(1 << 16); // Generic domain
        domains.push(type(uint32).max - 1); // Test max value
        domains.push(type(uint32).max); // Test max value
        // Create values for removal ordering tests
        createExpectedStates();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TESTS: SINGLE AGENT                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent(uint32 domain, address account) public {
        // Check if agent was active on this domain
        bool wasActiveDomain = _isActive(domain, account);
        // Agent will not added only if they were not previously active
        bool toBeAdded = !_isActive(account);
        _addAgent(domain, account);
        assertEq(libHarness.add(domain, account), toBeAdded, "!add: return value");
        // Agent should be marked as active on given domain, if it was added or active there before
        assertEq(libHarness.contains(domain, account), toBeAdded || wasActiveDomain, "!add: contains(domain, account)");
        (bool isActive, uint32 domain_) = libHarness.contains(account);
        // Agent should be marked as active globally
        assertEq(isActive, agents[account].isActive, "!add: contains(account), isActive");
        assertEq(domain_, agents[account].domain, "!add: contains(account), domain");
    }

    function test_removeAgent(uint32 domain, address account) public {
        // Agent will be removed only they were active on the given domain
        bool toBeRemoved = _isActive(domain, account);
        _removeAgent(domain, account);
        assertEq(libHarness.remove(domain, account), toBeRemoved, "!remove: return value");
        // Agent should be marked as not active on given domain
        assertFalse(libHarness.contains(domain, account), "!remove: contains(domain, account)");
        // Agent should be marked as active globally only if they were active on another domain
        (bool isActive, uint32 domain_) = libHarness.contains(account);
        assertEq(isActive, agents[account].isActive, "!remove: contains(account), isActive");
        assertEq(domain_, agents[account].domain, "!remove: contains(account), domain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        TESTS: MULTIPLE AGENTS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_suiteAgents() public {
        for (uint256 d = 0; d < domains.length; ++d) {
            uint32 domain = domains[d];
            for (uint256 index = 0; index < AGENTS_PER_DOMAIN; ++index) {
                test_addAgent(domain, _mockAgent(domain, index));
            }
        }
    }

    function test_addAgent_activeSameDomain() public {
        test_addAgent_suiteAgents();
        test_addAgent_suiteAgents();
    }

    function test_addAgent_activeAnotherDomain() public {
        test_addAgent_suiteAgents();
        for (uint256 d = 0; d < domains.length; ++d) {
            uint32 domain = domains[d];
            uint32 foreignDomain = domains[(d + 1) % domains.length];
            for (uint256 index = 0; index < AGENTS_PER_DOMAIN; ++index) {
                // Try to add an agent that is active on a foreign domain
                test_addAgent(domain, _mockAgent(foreignDomain, index));
            }
        }
    }

    function test_removeAgent_suiteAgents() public {
        test_addAgent_suiteAgents();
        for (uint256 d = 0; d < domains.length; ++d) {
            uint32 domain = domains[d];
            for (uint256 index = 0; index < AGENTS_PER_DOMAIN; ++index) {
                test_removeAgent(domain, _mockAgent(domain, index));
            }
        }
    }

    function test_removeAgent_notActiveAnywhere() public {
        test_removeAgent_suiteAgents();
        for (uint256 d = 0; d < domains.length; ++d) {
            uint32 domain = domains[d];
            for (uint256 index = 0; index < AGENTS_PER_DOMAIN; ++index) {
                test_removeAgent(domain, _mockAgent(domain, index));
            }
        }
    }

    function test_removeAgent_activeAnotherDomain() public {
        test_addAgent_suiteAgents();
        for (uint256 d = 0; d < domains.length; ++d) {
            uint32 domain = domains[d];
            uint32 foreignDomain = domains[(d + 1) % domains.length];
            for (uint256 index = 0; index < AGENTS_PER_DOMAIN; ++index) {
                // Try to remove an agent that is active on a foreign domain
                test_removeAgent(domain, _mockAgent(foreignDomain, index));
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: REMOVAL & ORDERING                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line code-complexity
    function test_removalOrdering() public {
        for (uint256 d = 0; d < domains.length; ++d) {
            uint32 domain = domains[d];
            for (uint256 index = 0; index < ELEMENTS; ++index) {
                test_addAgent(domain, _mockAgent(domain, index));
            }
        }
        for (uint256 step = 0; step < ELEMENTS; ++step) {
            // First, check the expected state for al domains
            for (uint256 d = 0; d < domains.length; ++d) {
                uint32 domain = domains[d];
                uint256 amount = ELEMENTS - step;
                address[] memory expectedAgents = new address[](amount);
                for (uint256 i = 0; i < amount; ++i) {
                    uint256 index = expectedStates[step][i];
                    expectedAgents[i] = _mockAgent(domain, index);
                }
                // Check amount of agents per domain
                assertEq(libHarness.length(domain), amount, "!length");
                // Check ordering
                address[] memory values = libHarness.values(domain);
                assertEq(values.length, amount, "!values: length");
                for (uint256 i = 0; i < amount; ++i) {
                    assertEq(values[i], expectedAgents[i], "!values: ordering");
                    assertEq(libHarness.at(domain, i), expectedAgents[i], "!at");
                }
            }
            // Then, remove a predetermined agent from every domain
            for (uint256 d = 0; d < domains.length; ++d) {
                uint32 domain = domains[d];
                uint256 index = removalOrder[step];
                test_removeAgent(domain, _mockAgent(domain, index));
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _addAgent(uint32 domain, address account) internal {
        if (!_isActive(account)) {
            agents[account] = agent({isActive: true, domain: domain});
        }
    }

    function _removeAgent(uint32 domain, address account) internal {
        if (_isActive(domain, account)) {
            delete agents[account];
        }
    }

    function _isActive(address account) internal view returns (bool) {
        return agents[account].isActive;
    }

    function _isActive(uint32 domain, address account) internal view returns (bool) {
        return agents[account].isActive && agents[account].domain == domain;
    }

    function _mockAgent(uint32 domain, uint256 index) internal pure returns (address) {
        return address((uint160(domain) << 128) | uint160(index));
    }
}
