// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager } from "./manager/AgentManager.sol";
import { DomainContext } from "./context/DomainContext.sol";
import { SummitEvents } from "./events/SummitEvents.sol";
import { InterfaceSummit } from "./interfaces/InterfaceSummit.sol";
import { ExecutionHub } from "./hubs/ExecutionHub.sol";
import { SnapshotHub, SummitAttestation, SummitState } from "./hubs/SnapshotHub.sol";
import { Attestation, AttestationLib, AttestationReport, Snapshot } from "./hubs/StatementHub.sol";

contract Summit is ExecutionHub, SnapshotHub, SummitEvents, InterfaceSummit {
    using AttestationLib for bytes;

    constructor(uint32 _domain) DomainContext(_domain) {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __SystemContract_initialize();
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
            // Slash Notary and trigger a hook to send a slashAgent system call
            _slashAgent(domain, notary, true);
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
            // Slash Guard (domain == 0) and trigger a hook to send a slashAgent system call
            _slashAgent(0, guard, true);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function getLatestState(uint32 _origin) external view returns (bytes memory statePayload) {
        uint256 guardsAmount = amountAgents(0);
        SummitState memory latestState;
        for (uint256 i = 0; i < guardsAmount; ++i) {
            address guard = getAgent(0, i);
            SummitState memory state = _latestState(_origin, guard);
            if (state.nonce > latestState.nonce) {
                latestState = state;
            }
        }
        // Check if we found anything
        if (latestState.nonce != 0) {
            statePayload = latestState.formatSummitState();
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Hook that is called after an existing agent was slashed,
    /// when verification of an invalid agent statement was done in this contract.
    function _afterAgentSlashed(uint32 _domain, address _agent) internal virtual override {
        // TODO: implement
    }

    function _isIgnoredAgent(uint32, address) internal pure override returns (bool) {
        // Summit keeps track of every agent
        return false;
    }
}
