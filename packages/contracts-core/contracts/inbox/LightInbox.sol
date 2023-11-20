// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/memory/Attestation.sol";
import {IncorrectDataHash, SynapseDomainForbidden} from "../libs/Errors.sol";
import {ChainGas, GasDataLib} from "../libs/stack/GasData.sol";
import {AgentStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {StatementInbox} from "./StatementInbox.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {InterfaceDestination} from "../interfaces/InterfaceDestination.sol";
import {InterfaceLightInbox} from "../interfaces/InterfaceLightInbox.sol";

/// @notice `LightInbox` is the child of `StatementInbox` contract, that is used chains other than the Synapse Chain.
/// In addition to the functionality of `StatementInbox`, it also:
/// - Accepts Notary Attestations and passes them to the `Destination` contract.
/// - Accepts Attestation Reports and initiates a dispute between the Notary and the Guard.
contract LightInbox is StatementInbox, InterfaceLightInbox {
    using AttestationLib for bytes;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_) MessagingBase("0.0.3", synapseDomain_) {
        if (localDomain == synapseDomain) revert SynapseDomainForbidden();
    }

    /// @notice Initializes `LightInbox` contract:
    /// - Sets `owner_` as the owner of the contract
    /// - Sets `agentManager`, `origin` and `destination` addresses
    function initialize(address agentManager_, address origin_, address destination_, address owner_)
        external
        initializer
    {
        __StatementInbox_init(agentManager_, origin_, destination_, owner_);
    }

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc InterfaceLightInbox
    function submitAttestation(
        bytes memory attPayload,
        bytes memory attSignature,
        bytes32 agentRoot_,
        uint256[] memory snapGas_
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not an known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Check that Notary is active
        status.verifyActive();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(status.domain);
        // Cast uint256[] to ChainGas[] using assembly. This prevents us from doing unnecessary copies.
        // Note that this does NOT clear the highest bits, but it's ok as the dirty highest bits
        // will lead to hash mismatch in snapGasHash() and thus to attestation rejection.
        ChainGas[] memory snapGas;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            snapGas := snapGas_
        }
        // Check that hash of provided data matches the attestation's dataHash
        if (
            att.dataHash()
                != AttestationLib.dataHash({agentRoot_: agentRoot_, snapGasHash_: GasDataLib.snapGasHash(snapGas)})
        ) {
            revert IncorrectDataHash();
        }
        // Store Notary signature for the attestation
        uint256 sigIndex = _saveSignature(attSignature);
        // This will revert if Notary is in Dispute
        wasAccepted = InterfaceDestination(destination).acceptAttestation({
            notaryIndex: status.index,
            sigIndex: sigIndex,
            attPayload: attPayload,
            agentRoot: agentRoot_,
            snapGas: snapGas
        });
        if (wasAccepted) {
            emit AttestationAccepted(status.domain, notary, attPayload, attSignature);
        }
    }

    /// @inheritdoc InterfaceLightInbox
    function submitAttestationReport(bytes memory attPayload, bytes memory arSignature, bytes memory attSignature)
        external
        returns (bool wasAccepted)
    {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory guardStatus,) = _verifyAttestationReport(att, arSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if attestation signer is not a known Notary
        (AgentStatus memory notaryStatus,) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        _saveReport(attPayload, arSignature);
        // This will revert if either actor is already in dispute
        IAgentManager(agentManager).openDispute(guardStatus.index, notaryStatus.index);
        return true;
    }
}
