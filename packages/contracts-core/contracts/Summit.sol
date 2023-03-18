// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { SummitState } from "./libs/State.sol";
import { AgentInfo } from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager } from "./manager/AgentManager.sol";
import { DomainContext } from "./context/DomainContext.sol";
import { SummitEvents } from "./events/SummitEvents.sol";
import { InterfaceSummit } from "./interfaces/InterfaceSummit.sol";
import { SnapshotHub } from "./hubs/SnapshotHub.sol";
import { Attestation, AttestationReport, Snapshot, StatementHub } from "./hubs/StatementHub.sol";

/**
 * @notice Accepts snapshots signed by Guards and Notaries. Verifies Notaries attestations.
 */
contract Summit is StatementHub, SnapshotHub, AgentManager, SummitEvents, InterfaceSummit {
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
    ▏*║                            ADDING AGENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool isAdded) {
        isAdded = _addAgent(_domain, _account);
        if (isAdded) {
            _syncAgentLocalRegistries(AgentInfo(_domain, _account, true));
        }
    }

    function removeAgent(uint32 _domain, address _account)
        external
        onlyOwner
        returns (bool isRemoved)
    {
        isRemoved = _removeAgent(_domain, _account);
        if (isRemoved) {
            _syncAgentLocalRegistries(AgentInfo(_domain, _account, false));
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ACCEPT STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function submitSnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        external
        returns (bool wasAccepted)
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
            _acceptNotarySnapshot(snapshot, agent);
        }
        emit SnapshotAccepted(domain, agent, _snapPayload, _snapSignature);
        return true;
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
        /// @dev Summit is BondingManager, so we need to slash Agent on local Registries,
        /// as well as relay this information to all other chains.
        /// There was no system call that triggered slashing, so callOrigin is set to ZERO.
        _updateLocalRegistries({
            _data: _dataSlashAgent(_domain, _agent),
            _forwardUpdate: true,
            _callOrigin: 0
        });
    }
}
