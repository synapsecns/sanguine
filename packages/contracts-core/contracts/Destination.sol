// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation } from "./libs/Attestation.sol";
import { AttestationReport } from "./libs/AttestationReport.sol";
import { AgentStatus } from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DestinationEvents } from "./events/DestinationEvents.sol";
import { IAgentManager } from "./interfaces/IAgentManager.sol";
import { ExecutionAttestation, InterfaceDestination } from "./interfaces/InterfaceDestination.sol";
import { ExecutionHub } from "./hubs/ExecutionHub.sol";
import { DomainContext, Versioned } from "./system/SystemContract.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev All snapshot roots from the saved attestations
    bytes32[] private roots;

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
    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ACCEPT STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceDestination
    function submitAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool wasAccepted)
    {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(_attPayload);
        // This will revert if signer is not an known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, _attSignature);
        // Check that Notary is active
        _verifyActive(status);
        // Check that Notary domain is local domain
        require(status.domain == localDomain, "Wrong Notary domain");
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, notary);
        emit AttestationAccepted(status.domain, notary, _attPayload, _attSignature);
        return true;
    }

    /// @inheritdoc InterfaceDestination
    function submitAttestationReport(
        bytes memory _arPayload,
        bytes memory _arSignature,
        bytes memory _attSignature
    ) external returns (bool wasAccepted) {
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
    ▏*║                            DISPUTE LOGIC                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
}
