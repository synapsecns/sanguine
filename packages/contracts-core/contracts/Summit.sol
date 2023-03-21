// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager } from "./manager/AgentManager.sol";
import { DomainContext } from "./context/DomainContext.sol";
import { SummitEvents } from "./events/SummitEvents.sol";
import { IAgentManager } from "./interfaces/IAgentManager.sol";
import { InterfaceSummit } from "./interfaces/InterfaceSummit.sol";
import { ExecutionHub } from "./hubs/ExecutionHub.sol";
import { SnapshotHub, SummitAttestation, SummitState } from "./hubs/SnapshotHub.sol";
import { Attestation, AttestationLib, AttestationReport, Snapshot } from "./hubs/StatementHub.sol";
import { DomainContext, Versioned } from "./system/SystemContract.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Summit is ExecutionHub, SnapshotHub, SummitEvents, InterfaceSummit {
    using AttestationLib for bytes;

    constructor(uint32 _domain, IAgentManager _agentManager)
        DomainContext(_domain)
        SystemRegistry(_agentManager)
        Versioned("0.0.3")
    {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        _initializeAttestations();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ACCEPT STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function submitSnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        external
        returns (bytes memory attPayload)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = _wrapSnapshot(_snapPayload);
        // This will revert if the signer is not an active Agent
        (uint32 domain, address agent) = _verifySnapshot(snapshot, _snapSignature);
        if (domain == 0) {
            // This will revert if Guard has previously submitted
            // a fresher state than one in the snapshot.
            _acceptGuardSnapshot(snapshot, agent);
        } else {
            // This will revert if any of the states from the Notary snapshot
            // haven't been submitted by any of the Guards before.
            attPayload = _acceptNotarySnapshot(snapshot, agent);
            // Save attestation derived from Notary snapshot
            _saveAttestation(attPayload.castToAttestation(), agent);
        }
        emit SnapshotAccepted(domain, agent, _snapPayload, _snapSignature);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VERIFY STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not an attestation
        Attestation att = _wrapAttestation(_attPayload);
        // This will revert if the attestation signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(att, _attSignature);
        isValid = _isValidAttestation(att);
        if (!isValid) {
            emit InvalidAttestation(_attPayload, _attSignature);
            // Slash Notary and notify local AgentManager
            _slashAgent(domain, notary);
        }
    }

    /// @inheritdoc InterfaceSummit
    function verifyAttestationReport(bytes memory _arPayload, bytes memory _arSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not an attestation report
        AttestationReport report = _wrapAttestationReport(_arPayload);
        // This will revert if the report signer is not an active Guard
        address guard = _verifyAttestationReport(report, _arSignature);
        // Report is valid, if the reported attestation is invalid
        isValid = !_isValidAttestation(report.attestation());
        if (!isValid) {
            emit InvalidAttestationReport(_arPayload, _arSignature);
            // Slash Guard and notify local AgentManager
            _slashAgent(0, guard);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function getLatestState(uint32 _origin) external view returns (bytes memory statePayload) {
        // TODO: implement once Agent Merkle Tree is done
    }
}
