// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation } from "./libs/Attestation.sol";
import { AttestationReport } from "./libs/AttestationReport.sol";
import { AGENT_ROOT_OPTIMISTIC_PERIOD } from "./libs/Constants.sol";
import { AgentStatus, DestinationStatus } from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DestinationEvents } from "./events/DestinationEvents.sol";
import { IAgentManager } from "./interfaces/IAgentManager.sol";
import { ExecutionAttestation, InterfaceDestination } from "./interfaces/InterfaceDestination.sol";
import { ILightManager } from "./interfaces/ILightManager.sol";
import { DisputeHub, ExecutionHub } from "./hubs/ExecutionHub.sol";
import { DomainContext, Versioned } from "./system/SystemContract.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev All snapshot roots from the saved attestations
    bytes32[] private roots;

    /// @inheritdoc InterfaceDestination
    /// @dev Invariant: this is either current LightManager root,
    /// or the pending root to be passed to LightManager once its optimistic period is over.
    bytes32 public nextAgentRoot;

    /// @inheritdoc InterfaceDestination
    DestinationStatus public destStatus;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain, IAgentManager _agentManager)
        DomainContext(_domain)
        SystemRegistry(_agentManager)
        Versioned("0.0.3")
    {}

    /// @notice Initializes Destination contract:
    /// - msg.sender is set as contract owner
    function initialize(bytes32 _agentRoot) external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        // Set Agent Merkle Root in Light Manager
        nextAgentRoot = _agentRoot;
        ILightManager(address(agentManager)).setAgentRoot(_agentRoot);
        destStatus.agentRootTime = uint48(block.timestamp);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ACCEPT STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceDestination
    function submitAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool wasAccepted)
    {
        // First, try passing current agent merkle root
        (bool rootPassed, bool rootPending) = passAgentRoot();
        // Don't accept attestation, if the agent root was updated in LightManager,
        // as the following agent check will fail.
        if (rootPassed) return false;
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(_attPayload);
        // This will revert if signer is not an known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, _attSignature);
        // Check that Notary is active
        _verifyActive(status);
        // Check that Notary domain is local domain
        require(status.domain == localDomain, "Wrong Notary domain");
        // Check that Notary who submitted the attestation is not in dispute
        require(!_inDispute(notary), "Notary is in dispute");
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, notary);
        // Save Agent Root if required, and update the Destination's Status
        destStatus = _saveAgentRoot(rootPending, att.agentRoot(), notary);
        emit AttestationAccepted(status.domain, notary, _attPayload, _attSignature);
        return true;
    }

    /// @inheritdoc InterfaceDestination
    function submitAttestationReport(
        bytes memory _arPayload,
        bytes memory _arSignature,
        bytes memory _attSignature
    ) external returns (bool wasAccepted) {
        // Call the hook and check if we can accept the statement
        if (!_beforeStatement()) return false;
        // This will revert if payload is not an attestation report
        AttestationReport report = _wrapAttestationReport(_arPayload);
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory guardStatus, address guard) = _verifyAttestationReport(
            report,
            _arSignature
        );
        // Check that Guard is active
        _verifyActive(guardStatus);
        // This will revert if attestation signer is not a known Notary
        (AgentStatus memory notaryStatus, address notary) = _verifyAttestation(
            report.attestation(),
            _attSignature
        );
        // Notary needs to be Active/Unstaking
        _verifyActiveUnstaking(notaryStatus);
        // Reported Attestation was signed by the Notary => open dispute
        _openDispute(guard, notaryStatus.domain, notary);
        return true;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        AGENT ROOT QUARANTINE                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceDestination
    function passAgentRoot() public returns (bool rootPassed, bool rootPending) {
        bytes32 oldRoot = agentManager.agentRoot();
        bytes32 newRoot = nextAgentRoot;
        // Check if agent root differs from the current one in LightManager
        if (oldRoot == newRoot) return (false, false);
        DestinationStatus memory status = destStatus;
        // Invariant: Notary who supplied `newRoot` was registered as active against `oldRoot`
        // So we just need to check the Dispute status of the Notary
        if (_inDispute(status.notary)) {
            // Remove the pending agent merkle root, as its signer is in dispute
            nextAgentRoot = oldRoot;
            return (false, false);
        }
        // Check if agent root optimistic period is over
        if (status.agentRootTime + AGENT_ROOT_OPTIMISTIC_PERIOD > block.timestamp) {
            // We didn't pass anything, but there is a pending root
            return (false, true);
        }
        // `newRoot` signer was not disputed, and the root optimistic period is over.
        // Finally, pass the Agent Merkle Root to LightManager
        ILightManager(address(agentManager)).setAgentRoot(newRoot);
        return (true, false);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceDestination
    function attestationsAmount() external view returns (uint256) {
        return roots.length;
    }

    /// @inheritdoc InterfaceDestination
    function getAttestation(uint256 _index)
        external
        view
        returns (bytes32 root, ExecutionAttestation memory execAtt)
    {
        require(_index < roots.length, "Index out of range");
        root = roots[_index];
        execAtt = _getRootAttestation(root);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc DisputeHub
    function _beforeStatement() internal override returns (bool acceptNext) {
        (bool rootPassed, ) = passAgentRoot();
        // We don't accept statements if the root was updated just now,
        // as all the agent checks will fail otherwise.
        return !rootPassed;
    }

    /// @dev Opens a Dispute between a Guard and a Notary.
    /// This is overridden to allow disputes only between a Guard and a LOCAL Notary.
    function _openDispute(
        address _guard,
        uint32 _domain,
        address _notary
    ) internal override {
        // Only disputes for local Notaries could be initiated in Destination
        require(_domain == localDomain, "Not a local Notary");
        super._openDispute(_guard, _domain, _notary);
    }

    /// @dev Resolves a Dispute for a slashed agent, if it hasn't been done already.
    /// This is overridden to resolve disputes only between a Guard and a LOCAL Notary.
    function _resolveDispute(uint32 _domain, address _slashedAgent) internal virtual override {
        // Disputes could be only opened between a Guard and a local Notary
        if (_domain == 0 || _domain == localDomain) {
            super._resolveDispute(_domain, _slashedAgent);
        }
    }

    /// @dev Saves Agent Merkle Root from the accepted attestation, if there is
    /// no pending root to be passed to LightManager.
    /// Returns the updated "last snapshot root / last agent root" status struct.
    function _saveAgentRoot(
        bool _rootPending,
        bytes32 _agentRoot,
        address _notary
    ) internal returns (DestinationStatus memory status) {
        status = destStatus;
        // Update the timestamp for the latest snapshot root
        status.snapRootTime = uint48(block.timestamp);
        // Don't update agent root, if there is already a pending one
        // Update the data for latest agent root only if it differs from the saved one
        if (!_rootPending && nextAgentRoot != _agentRoot) {
            status.agentRootTime = uint48(block.timestamp);
            status.notary = _notary;
            nextAgentRoot = _agentRoot;
            emit AgentRootAccepted(_agentRoot);
        }
    }
}
