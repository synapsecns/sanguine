// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AGENT_TREE_HEIGHT } from "../libs/Constants.sol";
import { MerkleLib } from "../libs/Merkle.sol";
import { AgentFlag, AgentStatus, SlashStatus } from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager, IAgentManager, ISystemRegistry } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { IBondingManager } from "../interfaces/IBondingManager.sol";
import { ILightManager } from "../interfaces/ILightManager.sol";
import { Versioned } from "../Version.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is Versioned, AgentManager, ILightManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Latest known Agent Merkle Root
    bytes32 private latestAgentRoot;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) private agentMap;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    function initialize(ISystemRegistry _origin, ISystemRegistry _destination)
        external
        initializer
    {
        __AgentManager_init(_origin, _destination);
        __Ownable_init();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             AGENTS LOGIC                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ILightManager
    function updateAgentStatus(
        address _agent,
        AgentStatus memory _status,
        bytes32[] memory _proof
    ) external {
        // Reconstruct the agent leaf: flag should be Active
        bytes32 leaf = _agentLeaf(_status.flag, _status.domain, _agent);
        bytes32 root = latestAgentRoot;
        // Check that proof matches the latest merkle root
        require(
            MerkleLib.proofRoot(_status.index, leaf, _proof, AGENT_TREE_HEIGHT) == root,
            "Invalid proof"
        );
        // Update the agent status against this root
        agentMap[root][_agent] = _status;
        emit StatusUpdated(_status.flag, _status.domain, _agent);
        // Notify local Registries, if agent flag is Slashed
        if (_status.flag == AgentFlag.Slashed) {
            // Prover is msg.sender
            _notifySlashing(DESTINATION | ORIGIN, _status.domain, _agent, msg.sender);
        }
    }

    /// @inheritdoc ILightManager
    function setAgentRoot(bytes32 _agentRoot) external {
        require(msg.sender == address(destination), "Only Destination sets agent root");
        _setAgentRoot(_agentRoot);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SLASHING LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function registrySlash(
        uint32 _domain,
        address _agent,
        address _prover
    ) external {
        // Check that Agent hasn't been already slashed and initiate the slashing
        _initiateSlashing(_domain, _agent, _prover);
        // On chains other than Synapse Chain only Origin could slash Agents
        if (msg.sender == address(origin)) {
            _notifySlashing(DESTINATION, _domain, _agent, _prover);
            // Issue a system call to BondingManager on SynChain
            _callAgentManager({
                _domain: SYNAPSE_DOMAIN,
                _optimisticSeconds: BONDING_OPTIMISTIC_PERIOD,
                _data: _remoteSlashData(_domain, _agent, _prover)
            });
        } else {
            revert("Unauthorized caller");
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function agentRoot() public view override returns (bytes32) {
        return latestAgentRoot;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Updates the Agent Merkle Root that Light Manager is tracking.
    function _setAgentRoot(bytes32 _agentRoot) internal {
        if (latestAgentRoot != _agentRoot) {
            latestAgentRoot = _agentRoot;
            emit RootUpdated(_agentRoot);
        }
    }

    /// @dev Returns the status for the agent: whether or not they have been added
    /// using latest Agent merkle Root.
    function _agentStatus(address _agent) internal view override returns (AgentStatus memory) {
        return agentMap[latestAgentRoot][_agent];
    }

    /// @dev Returns data for a system call: remoteRegistrySlash()
    function _remoteSlashData(
        uint32 _domain,
        address _agent,
        address _prover
    ) internal pure returns (bytes memory) {
        // (_rootSubmittedAt, _callOrigin, _systemCaller, _domain, _agent, _prover)
        return
            abi.encodeWithSelector(
                IBondingManager.remoteRegistrySlash.selector,
                0,
                0,
                0,
                _domain,
                _agent,
                _prover
            );
    }
}
