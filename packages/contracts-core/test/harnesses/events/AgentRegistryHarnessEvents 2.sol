// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

abstract contract AgentRegistryHarnessEvents {
    event AfterAgentAdded(uint32 domain, address account);

    event AfterAgentRemoved(uint32 domain, address account);
}
