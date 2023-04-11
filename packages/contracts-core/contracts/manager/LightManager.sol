// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AGENT_TREE_HEIGHT} from "../libs/Constants.sol";
import {MerkleLib} from "../libs/Merkle.sol";
import {AgentFlag, AgentStatus, SlashStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager, ISystemRegistry} from "./AgentManager.sol";
import {DomainContext} from "../context/DomainContext.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {Versioned} from "../Version.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is Versioned, AgentManager, InterfaceLightManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════
    // Latest known Agent Merkle Root
    bytes32 private _latestAgentRoot;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) private _agentMap;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) DomainContext(domain) Versioned("0.0.3") {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    function initialize(ISystemRegistry origin_, ISystemRegistry destination_) external initializer {
        __AgentManager_init(origin_, destination_);
        __Ownable_init();
    }

    // ═══════════════════════════════════════════════ AGENTS LOGIC ════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function updateAgentStatus(address agent, AgentStatus memory status, bytes32[] memory proof) external {
        // Reconstruct the agent leaf: flag should be Active
        bytes32 leaf = _agentLeaf(status.flag, status.domain, agent);
        bytes32 root = _latestAgentRoot;
        // Check that proof matches the latest merkle root
        require(MerkleLib.proofRoot(status.index, leaf, proof, AGENT_TREE_HEIGHT) == root, "Invalid proof");
        // Update the agent status against this root
        _agentMap[root][agent] = status;
        emit StatusUpdated(status.flag, status.domain, agent);
        // Notify local Registries, if agent flag is Slashed
        if (status.flag == AgentFlag.Slashed) {
            // Prover is msg.sender
            _notifySlashing(DESTINATION | ORIGIN, status.domain, agent, msg.sender);
        }
    }

    /// @inheritdoc InterfaceLightManager
    function setAgentRoot(bytes32 agentRoot_) external {
        require(msg.sender == address(destination), "Only Destination sets agent root");
        _setAgentRoot(agentRoot_);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function agentRoot() public view override returns (bytes32) {
        return _latestAgentRoot;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    function _afterRegistrySlash(uint32 domain, address agent, address prover) internal override {
        // Issue a system call to BondingManager on SynChain
        _callAgentManager({
            domain: SYNAPSE_DOMAIN,
            optimisticPeriod: BONDING_OPTIMISTIC_PERIOD,
            payload: _remoteSlashPayload(domain, agent, prover)
        });
    }

    /// @dev Updates the Agent Merkle Root that Light Manager is tracking.
    function _setAgentRoot(bytes32 _agentRoot) internal {
        if (_latestAgentRoot != _agentRoot) {
            _latestAgentRoot = _agentRoot;
            emit RootUpdated(_agentRoot);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the status for the agent: whether or not they have been added
    /// using latest Agent merkle Root.
    function _agentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[_latestAgentRoot][agent];
    }

    /// @dev Returns data for a system call: remoteRegistrySlash()
    function _remoteSlashPayload(uint32 domain, address agent, address prover) internal pure returns (bytes memory) {
        // (proofMaturity, callOrigin, systemCaller) are omitted; (domain, agent, prover)
        return abi.encodeWithSelector(InterfaceBondingManager.remoteRegistrySlash.selector, domain, agent, prover);
    }
}
