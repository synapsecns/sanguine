// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { BondingManager } from "./BondingManager.sol";

interface IAttestationCollector {
    function addAgent(uint32 _domain, address _account) external returns (bool);

    function removeAgent(uint32 _domain, address _account) external returns (bool);
}

/**
 * @notice MVP for BondingManager. Controls agents status for local chain registries.
 * Doesn't do anything cross-chain related.
 */
contract BondingMVP is BondingManager {
    address public attestationCollector;

    constructor(uint32 _domain) DomainContext(_domain) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Adds an agent to all system registries on the local chain.
     * Does no cross-chain calls whatsoever.
     */
    function addAgent(uint32 _domain, address _account) external onlyOwner {
        _updateAgentStatus({ _domain: _domain, _agent: _account, _bonded: true });
        // TODO: remove when AttestationCollector is merged with BondingPrimary
        if (attestationCollector != address(0)) {
            IAttestationCollector(attestationCollector).addAgent(_domain, _account);
        }
    }

    /**
     * @notice Removes an agent from all system registries on the local chain.
     * Does no cross-chain calls whatsoever.
     */
    function removeAgent(uint32 _domain, address _account) external onlyOwner {
        _updateAgentStatus({ _domain: _domain, _agent: _account, _bonded: false });
        // TODO: remove when AttestationCollector is merged with BondingPrimary
        if (attestationCollector != address(0)) {
            IAttestationCollector(attestationCollector).removeAgent(_domain, _account);
        }
    }

    /// @notice Sets Attestation Collector address.
    /// @dev AttestationCollector.owner() needs to be BondingMVP.
    function setAttestationCollector(address _attestationCollector) external onlyOwner {
        attestationCollector = _attestationCollector;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating the list of off-chain agents needs to be synced.
     * @dev Must be called from the BondingManager on Synapse Chain.
     */
    function syncAgents(
        uint256,
        uint32,
        ISystemRouter.SystemEntity,
        uint256,
        bool,
        AgentInfo[] memory
    ) external view override onlySystemRouter {
        revert("Cross-chain disabled");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _updateAgentStatus(
        uint32 _domain,
        address _agent,
        bool _bonded
    ) internal {
        AgentInfo[] memory infos = new AgentInfo[](1);
        infos[0] = AgentInfo(_domain, _agent, _bonded);
        // RequestID is ignored at the moment
        _updateLocalRegistries({
            _data: _dataSyncAgents({ _requestID: 0, _removeExisting: false, _infos: infos }),
            _forwardUpdate: false, // cross-chain interactions are disabled
            _callOrigin: 0 // there was no system call that initiated the bonding
        });
    }

    /**
     * @notice Forward data with an agent status update (due to
     * a system call from `_callOrigin`).
     * @dev If BondingManager is deployed on Synapse Chain, all other chains should be notified.
     * Otherwise, only Synapse Chain should be notified.
     */
    function _forwardUpdateData(bytes memory, uint32) internal pure override {
        // Don't do anything: cross-chain interactions are disabled
    }

    /**
     * @notice Perform all required security checks for a cross-chain
     * system call for slashing an agent.
     */
    function _assertCrossChainSlashing(
        uint256,
        uint32,
        ISystemRouter.SystemEntity
    ) internal pure override {
        revert("Cross-chain disabled");
    }
}
