// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";

import { AgentRegistryHarness } from "../../harnesses/system/AgentRegistryHarness.t.sol";

abstract contract AgentRegistryTools is SynapseTestSuite {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectAgentAdded(
        uint32 _domain,
        uint256 _agentIndex,
        bool _isFirst
    ) public {
        expectAgentAdded(_domain, suiteAgent(_domain, _agentIndex), _isFirst);
    }

    function expectAgentAdded(
        uint32 _domain,
        address _account,
        bool _isFirst
    ) public {
        vm.expectEmit(true, true, true, true);
        emit AgentAdded(_domain, _account);
        if (_isFirst) {
            vm.expectEmit(true, true, true, true);
            emit DomainActivated(_domain);
        }
        vm.expectEmit(true, true, true, true);
        emit AfterAgentAdded(_domain, _account);
    }

    function expectAgentRemoved(
        uint32 _domain,
        uint256 _agentIndex,
        bool _isLast
    ) public {
        expectAgentRemoved(_domain, suiteAgent(_domain, _agentIndex), _isLast);
    }

    function expectAgentRemoved(
        uint32 _domain,
        address _account,
        bool _isLast
    ) public {
        vm.expectEmit(true, true, true, true);
        emit AgentRemoved(_domain, _account);
        if (_isLast) {
            vm.expectEmit(true, true, true, true);
            emit DomainDeactivated(_domain);
        }
        vm.expectEmit(true, true, true, true);
        emit AfterAgentRemoved(_domain, _account);
    }
}
