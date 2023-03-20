// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation } from "./libs/Attestation.sol";
import { AttestationReport } from "./libs/AttestationReport.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DomainContext } from "./context/DomainContext.sol";
import { ExecutionAttestation, InterfaceDestination } from "./interfaces/InterfaceDestination.sol";
import { DestinationEvents } from "./events/DestinationEvents.sol";
import { ExecutionHub } from "./hubs/ExecutionHub.sol";

contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev All snapshot roots from the saved attestations
    bytes32[] private roots;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) DomainContext(_domain) {}

    /// @notice Initializes Origin contract:
    /// - msg.sender is set as contract owner
    function initialize() external initializer {
        // Initialize SystemContract: msg.sender is set as "owner"
        __SystemContract_initialize();
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
        // This will revert if signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(att, _attSignature);
        // Check that Notary is active on local domain
        require(domain == localDomain, "Wrong Notary domain");
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, notary);
        emit AttestationAccepted(domain, notary, _attPayload, _attSignature);
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
        // This will revert if the report signer is not an active Guard
        address guard = _verifyAttestationReport(report, _arSignature);
        // This will revert if attestation signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(report.attestation(), _attSignature);
        // Reported Attestation was signed by the Notary => open dispute
        _openDispute(guard, domain, notary);
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
    ▏*║                     INTERNAL LOGIC: ATTESTATION                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isIgnoredAgent(uint32 _domain, address)
        internal
        view
        virtual
        override
        returns (bool)
    {
        // Destination only keeps track of local Notaries and Guards
        return _domain != localDomain && _domain != 0;
    }
}
