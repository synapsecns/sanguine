// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AgentSet } from "../libs/AgentSet.sol";
import { Auth } from "../libs/Auth.sol";
import { Signature } from "../libs/ByteString.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentRegistryEvents } from "../events/AgentRegistryEvents.sol";
import { IAgentRegistry } from "../interfaces/IAgentRegistry.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice Registry used for verifying signatures of any of the Agents.
 * Both Guards and Notaries could be stored in a single AgentRegistry.
 * An option to ignore certain agents is available, see {_isIgnoredAgent}.
 * @dev Following assumptions are implied:
 * 1. Guard is active on all domains at once.
 * 2. Notary is active on a single domain.
 * 3. Same account can't be both a Guard and a Notary.
 */
abstract contract AgentRegistry is AgentRegistryEvents, IAgentRegistry {
    using AgentSet for AgentSet.DomainAddressSet;
    using EnumerableSet for EnumerableSet.UintSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Refers to the current epoch. Whenever a full agent reset is required
     * by BondingManager, a new epoch starts. This saves us from iterating over all
     * agents and deleting them, which could be gas consuming.
     * @dev Variable is private as the child contracts are not supposed to modify it.
     * Use _currentEpoch() getter if needed.
     */
    uint256 private epoch;

    /**
     * @notice All active domains, i.e. domains having at least one active Notary.
     * Note: guards are stored with domain = 0, but we don't want to mix
     * "domains with at least one active Notary" and "zero domain with at least one active Guard",
     * so we are NOT storing domain == 0 in this set.
     */
    // (epoch => [domains with at least one active Notary])
    mapping(uint256 => EnumerableSet.UintSet) internal domains;

    /**
     * @notice DomainAddressSet implies that every agent is stored as a (domain, account) tuple.
     * Guard is active on all domains => Guards are stored as (domain = 0, account).
     * Notary is active on one (non-zero) domain => Notaries are stored as (domain > 0, account).
     */
    // (epoch => [set of active agents for all domains])
    mapping(uint256 => AgentSet.DomainAddressSet) internal agents;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Ensures that there is at least one active Notary for the given domain.
     */
    modifier haveActiveNotary(uint32 _domain) {
        require(_isActiveDomain(_domain), "No active notaries");
        _;
    }

    /**
     * @notice Ensures that there is at least one active Guard.
     */
    modifier haveActiveGuard() {
        // Guards are stored with `_domain == 0`
        require(amountAgents({ _domain: 0 }) != 0, "No active guards");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentRegistry
    function allAgents(uint32 _domain) external view returns (address[] memory) {
        return agents[_currentEpoch()].values(_domain);
    }

    /// @inheritdoc IAgentRegistry
    function allDomains() external view returns (uint32[] memory domains_) {
        uint256[] memory values = domains[_currentEpoch()].values();
        // Use assembly to perform uint256 -> uint32 downcast
        // See OZ's EnumerableSet.values()
        // solhint-disable-next-line no-inline-assembly
        assembly {
            domains_ := values
        }
    }

    /// @inheritdoc IAgentRegistry
    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain) {
        return _isActiveAgent(_account);
    }

    /// @inheritdoc IAgentRegistry
    function isActiveAgent(uint32 _domain, address _account) external view returns (bool) {
        return _isActiveAgent(_domain, _account);
    }

    /// @inheritdoc IAgentRegistry
    function isActiveDomain(uint32 _domain) external view returns (bool) {
        return _isActiveDomain(_domain);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             PUBLIC VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentRegistry
    function amountAgents(uint32 _domain) public view returns (uint256) {
        return agents[_currentEpoch()].length(_domain);
    }

    /// @inheritdoc IAgentRegistry
    function amountDomains() public view returns (uint256) {
        return domains[_currentEpoch()].length();
    }

    /// @inheritdoc IAgentRegistry
    function getAgent(uint32 _domain, uint256 _agentIndex) public view returns (address) {
        return agents[_currentEpoch()].at(_domain, _agentIndex);
    }

    /// @inheritdoc IAgentRegistry
    function getDomain(uint256 _domainIndex) public view returns (uint32) {
        return uint32(domains[_currentEpoch()].at(_domainIndex));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Tries to add an agent to the domain. If added, emits a corresponding event,
     * updates the list of active domains if necessary, and triggers a corresponding hook.
     * Note: use _domain == 0 to add a Guard, _domain > 0 to add a Notary.
     */
    function _addAgent(uint32 _domain, address _account) internal returns (bool wasAdded) {
        // Some Registries may want to ignore certain agents
        if (_isIgnoredAgent(_domain, _account)) return false;
        // Do the storage read just once
        uint256 _epoch = _currentEpoch();
        // Add to the list of agents for the domain in the current epoch
        wasAdded = agents[_epoch].add(_domain, _account);
        if (wasAdded) {
            emit AgentAdded(_domain, _account);
            // Consider adding domain to the list of "active domains" only if a Notary was added
            if (_domain != 0) {
                // We can skip the "already exists" check here, as EnumerableSet.add() does that
                if (domains[_epoch].add(_domain)) {
                    // Emit the event if domain was added to the list of active domains
                    emit DomainActivated(_domain);
                }
            }
            // Trigger the hook after the work is done
            _afterAgentAdded(_domain, _account);
        }
    }

    /**
     * @dev Tries to remove an agent from the domain. If removed, emits a corresponding event,
     * updates the list of active domains if necessary, and triggers a corresponding hook.
     * Note: use _domain == 0 to remove a Guard, _domain > 0 to remove a Notary.
     */
    function _removeAgent(uint32 _domain, address _account) internal returns (bool wasRemoved) {
        // Some Registries may want to ignore certain agents
        if (_isIgnoredAgent(_domain, _account)) return false;
        // Do the storage read just once
        uint256 _epoch = _currentEpoch();
        // Remove from the list of agents for the domain in the current epoch
        wasRemoved = agents[_epoch].remove(_domain, _account);
        if (wasRemoved) {
            emit AgentRemoved(_domain, _account);
            // Consider removing domain to the list of "active domains" only if a Notary was removed
            if (_domain != 0 && amountAgents(_domain) == 0) {
                // Remove domain for the "active list", if that was the last agent
                domains[_epoch].remove(_domain);
                emit DomainDeactivated(_domain);
            }
            // Trigger the hook after the work is done
            _afterAgentRemoved(_domain, _account);
        }
    }

    /**
     * @dev Tries to slash an agent active on the domain by removing it.
     * If slashed, emits a corresponding event, and triggers a corresponding hook if verified locally.
     * Hook will not be triggered, if agent was slashed elsewhere.
     * Note: use _domain == 0 to slash a Guard, _domain > 0 to slash a Notary.
     */
    function _slashAgent(
        uint32 _domain,
        address _account,
        bool _verified
    ) internal returns (bool wasSlashed) {
        wasSlashed = _removeAgent(_domain, _account);
        if (wasSlashed) {
            emit AgentSlashed(_domain, _account);
            if (_verified) _afterAgentSlashed(_domain, _account);
        }
    }

    /**
     * @dev Removes all active agents from all domains.
     * Note: iterating manually over all agents in order to delete them all is super inefficient.
     * Deleting sets (which contain mappings inside) is literally not possible.
     * So we're switching to fresh sets instead.
     */
    function _resetAgents() internal {
        ++epoch;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                HOOKS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable no-empty-blocks

    /// @dev Hook that is always called after a new agent was added for the domain.
    function _afterAgentAdded(uint32 _domain, address _account) internal virtual {}

    /// @dev Hook that is always called after an existing agent was removed from the domain.
    function _afterAgentRemoved(uint32 _domain, address _account) internal virtual {}

    /// @dev Hook that is called after an existing agent was slashed,
    /// when verification of an invalid agent statement was done in this contract.
    function _afterAgentSlashed(uint32 _domain, address _account) internal virtual {}

    // solhint-enable no-empty-blocks

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Returns current epoch, i.e. an index that is used to determine the currently
     * used sets for active agents and domains.
     */
    function _currentEpoch() internal view returns (uint256) {
        return epoch;
    }

    /**
     * @dev Recovers a signer from digest and signature, and checks if they are
     * active on the given domain.
     * Note: domain == 0 refers to a Guard, while _domain > 0 refers to a Notary.
     */
    function _checkAgentAuth(
        uint32 _domain,
        bytes32 _digest,
        Signature _signature
    ) internal view returns (address agent) {
        agent = Auth.recoverSigner(_digest, _signature);
        require(_isActiveAgent(_domain, agent), "Signer is not authorized");
    }

    /**
     * @dev Checks if agent is active on any of the domains.
     * Note: this returns if agent is active, and the domain where they're active.
     */
    function _isActiveAgent(address _account) internal view returns (bool, uint32) {
        // Check the list of global agents in the current epoch
        return agents[_currentEpoch()].contains(_account);
    }

    /**
     * @dev Checks if agent is active on the given domain.
     * Note: domain == 0 refers to a Guard, while _domain > 0 refers to a Notary.
     */
    function _isActiveAgent(uint32 _domain, address _account) internal view returns (bool) {
        // Check the list of the domain's agents in the current epoch
        return agents[_currentEpoch()].contains(_domain, _account);
    }

    /**
     * @dev Checks if there is at least one active Notary for the given domain.
     * Note: will return false for `_domain == 0`, even if there are active Guards.
     */
    function _isActiveDomain(uint32 _domain) internal view returns (bool) {
        return domains[_currentEpoch()].contains(_domain);
    }

    /**
     * @dev Child contracts should override this function to prevent
     * certain agents from being added and removed.
     * For instance, Origin might want to ignore all agents from the local domain.
     * Note: It is assumed that no agent can change its "ignored" status in any AgentRegistry.
     * In other words, do not use any values that might change over time, when implementing.
     * Otherwise, unexpected behavior might be expected. For instance, if an agent was added,
     * and then it became "ignored", it would be not possible to remove such agent.
     * Note: domain == 0 refers to a Guard, while _domain > 0 refers to a Notary.
     */
    function _isIgnoredAgent(uint32 _domain, address _account) internal view virtual returns (bool);
}
