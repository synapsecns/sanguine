// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

library AgentSet {
    /**
     * @notice Information about an active Agent, optimized to fit in one word of storage.
     * @dev We are storing both Notaries (domain > 0) and Guards (domain == 0) this way.
     * @param domain    Domain where Agent is active
     * @param index     Agent position in _agents[domain] array, plus 1 because index 0
     *                  means Agent is not active on any domain
     */
    struct AgentIndex {
        uint32 domain;
        uint224 index;
    }

    /**
     * @notice Information about all active agents for all domains.
     * @dev We are storing both Notaries (domain > 0) and Guards (domain == 0) this way.
     * @param agents    List of active agents for each domain
     * @param indexes   Information about every active agent
     */
    struct DomainAddressSet {
        // (domain => [list of agents for the domain])
        mapping(uint32 => address[]) _agents;
        // (agent => agentIndex)
        mapping(address => AgentIndex) _indexes;
    }

    /**
     * @notice Add an agent to a given domain's set of active agents. O(1)
     * @dev Will not add the agent, if it is already active on another domain.
     *
     * Returns true if the agent was added to the domain, that is
     * if it was not already active on any domain.
     */
    function add(DomainAddressSet storage set, uint32 domain, address account) internal returns (bool) {
        (bool isActive,) = contains(set, account);
        if (isActive) return false;
        set._agents[domain].push(account);
        // The agent is stored at length-1, but we add 1 to all indexes
        // and use 0 as a sentinel value
        set._indexes[account] = AgentIndex({domain: domain, index: uint224(set._agents[domain].length)});
        return true;
    }

    /**
     * @notice Remove an agent from a given domain's set of active agents. O(1)
     * @dev Will not remove the agent, if it is not active on the given domain.
     *
     * Returns true if the agent was removed from the domain, that is
     * if it was active on that domain.
     */
    function remove(DomainAddressSet storage set, uint32 domain, address account) internal returns (bool) {
        AgentIndex memory agentIndex = set._indexes[account];
        // Do nothing if agent is not active, or is active but on another domain
        if (agentIndex.index == 0 || agentIndex.domain != domain) return false;
        uint256 toDeleteIndex = agentIndex.index - 1;
        // To delete an Agent from the array in O(1),
        // we swap the Agent to delete with the last one in the array,
        // and then remove the last Agent (sometimes called as 'swap and pop').
        address[] storage agents = set._agents[domain];
        uint256 lastIndex = agents.length - 1;
        if (lastIndex != toDeleteIndex) {
            address lastAgent = agents[lastIndex];
            // Move the last Agent to the index where the Agent to delete is
            agents[toDeleteIndex] = lastAgent;
            // Update the index for the moved Agent (use deleted agent's value)
            set._indexes[lastAgent].index = agentIndex.index;
        }
        // Delete the slot where the moved Agent was stored
        agents.pop();
        // Delete the index for the deleted slot
        delete set._indexes[account];
        return true;
    }

    /**
     * @notice Returns true if the agent is active on any domain,
     * and the domain where the agent is active. O(1)
     */
    function contains(DomainAddressSet storage set, address account)
        internal
        view
        returns (bool isActive, uint32 domain)
    {
        AgentIndex memory agentIndex = set._indexes[account];
        if (agentIndex.index != 0) {
            isActive = true;
            domain = agentIndex.domain;
        }
    }

    /**
     * @notice Returns true if the agent is active on the given domain. O(1)
     */
    function contains(DomainAddressSet storage set, uint32 domain, address account) internal view returns (bool) {
        // Read from storage just once
        AgentIndex memory agentIndex = set._indexes[account];
        // Check that agent domain matches, and that agent is active
        return agentIndex.domain == domain && agentIndex.index != 0;
    }

    /**
     * @notice Returns a number of active agents for the given domain. O(1)
     */
    function length(DomainAddressSet storage set, uint32 domain) internal view returns (uint256) {
        return set._agents[domain].length;
    }

    /**
     * @notice Returns the agent stored at position `index` in the given domain's set. O(1).
     * @dev Note that there are no guarantees on the ordering of agents inside the
     * array, and it may change when more agents are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(DomainAddressSet storage set, uint32 domain, uint256 index) internal view returns (address) {
        return set._agents[domain][index];
    }

    /**
     * @notice Return the entire set of domain's agents in an array.
     *
     * @dev This operation will copy the entire storage to memory, which can be quite expensive.
     * This is designed to mostly be used by view accessors that are queried without any gas fees.
     * Developers should keep in mind that this function has an unbounded cost, and using it as part
     * of a state-changing function may render the function uncallable if the set grows to a point
     * where copying to memory consumes too much gas to fit in a block.
     */
    function values(DomainAddressSet storage set, uint32 domain) internal view returns (address[] memory) {
        return set._agents[domain];
    }
}
