// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IAgentRegistry {
    /**
     * @notice Returns the amount of active agents for the given domain.
     * Note: will return the amount of active Guards, if `_domain == 0`.
     */
    function amountAgents(uint32 _domain) external view returns (uint256);

    /**
     * @notice Returns the amount of active domains.
     * @dev This always excludes the zero domain, which is used for storing the guards.
     */
    function amountDomains() external view returns (uint256);

    /**
     * @notice Returns i-th agent for a given domain.
     * @dev Will revert if index is out of range.
     * Note: domain == 0 refers to a Guard, while _domain > 0 refers to a Notary.
     */
    function getAgent(uint32 _domain, uint256 _agentIndex) external view returns (address);

    /**
     * @notice Returns i-th domain from the list of active domains.
     * @dev Will revert if index is out of range.
     * Note: this never returns the zero domain, which is used for storing the guards.
     */
    function getDomain(uint256 _domainIndex) external view returns (uint32);

    /**
     * @notice Returns all active Agents for a given domain in an array.
     * Note: will return the list of active Guards, if `_domain == 0`.
     * @dev This copies storage into memory, so can consume a lof of gas, if
     * amount of agents is large (see EnumerableSet.values())
     */
    function allAgents(uint32 _domain) external view returns (address[] memory);

    /**
     * @notice Returns all domains having at least one active Notary in an array.
     * @dev This always excludes the zero domain, which is used for storing the guards.
     */
    function allDomains() external view returns (uint32[] memory domains_);

    /**
     * @notice Returns true if the agent is active on any domain.
     * Note: that includes both Guards and Notaries.
     * @return isActive Whether the account is an active agent on any of the domains
     * @return domain   Domain, where the account is an active agent
     */
    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain);

    /**
     * @notice Returns true if the agent is active on the given domain.
     * Note: domain == 0 refers to a Guard, while _domain > 0 refers to a Notary.
     */
    function isActiveAgent(uint32 _domain, address _account) external view returns (bool);

    /**
     * @notice Returns true if there is at least one active notary for the domain
     * Note: will return false for `_domain == 0`, even if there are active Guards.
     */
    function isActiveDomain(uint32 _domain) external view returns (bool);
}
