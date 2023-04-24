// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "./libs/Attestation.sol";
import {AttestationReport} from "./libs/AttestationReport.sol";
import {ByteString} from "./libs/ByteString.sol";
import {AGENT_ROOT_OPTIMISTIC_PERIOD} from "./libs/Constants.sol";
import {AgentStatus, DestinationStatus} from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DestinationEvents} from "./events/DestinationEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceDestination} from "./interfaces/InterfaceDestination.sol";
import {InterfaceLightManager} from "./interfaces/InterfaceLightManager.sol";
import {DisputeHub, ExecutionHub} from "./hubs/ExecutionHub.sol";
import {SystemBase, Versioned} from "./system/SystemBase.sol";
import {SystemRegistry} from "./system/SystemRegistry.sol";

contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    using AttestationLib for bytes;
    using ByteString for bytes;

    // TODO: this could be further optimized in terms of storage
    struct StoredAttData {
        bytes32 agentRoot;
        bytes32 r;
        bytes32 s;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    /// @dev Invariant: this is either current LightManager root,
    /// or the pending root to be passed to LightManager once its optimistic period is over.
    bytes32 public nextAgentRoot;

    /// @inheritdoc InterfaceDestination
    DestinationStatus public destStatus;

    /// @dev Stored lookup data for all accepted Notary Attestations
    StoredAttData[] internal _storedAttestations;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain, IAgentManager agentManager_)
        SystemBase(domain)
        SystemRegistry(agentManager_)
        Versioned("0.0.3")
    {} // solhint-disable-line no-empty-blocks

    /// @notice Initializes Destination contract:
    /// - msg.sender is set as contract owner
    function initialize(bytes32 agentRoot) external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        // Set Agent Merkle Root in Light Manager
        nextAgentRoot = agentRoot;
        InterfaceLightManager(address(agentManager)).setAgentRoot(agentRoot);
        destStatus.agentRootTime = uint48(block.timestamp);
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function acceptAttestation(
        address notary,
        AgentStatus memory status,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {
        // First, try passing current agent merkle root
        (bool rootPassed, bool rootPending) = passAgentRoot();
        // Don't accept attestation, if the agent root was updated in LightManager,
        // as the following agent check will fail.
        if (rootPassed) return false;
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // Check that Notary who submitted the attestation is not in dispute
        require(!_inDispute(notary), "Notary is in dispute");
        (bytes32 r, bytes32 s, uint8 v) = attSignature.castToSignature().toRSV();
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, status.index, v);
        bytes32 agentRoot = att.agentRoot();
        _storedAttestations.push(StoredAttData(agentRoot, r, s));
        // Save Agent Root if required, and update the Destination's Status
        destStatus = _saveAgentRoot(rootPending, agentRoot, notary);
        emit AttestationAccepted(status.domain, notary, attPayload, attSignature);
        return true;
    }

    // ═══════════════════════════════════════════ AGENT ROOT QUARANTINE ═══════════════════════════════════════════════

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
        InterfaceLightManager(address(agentManager)).setAgentRoot(newRoot);
        return (true, false);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    // solhint-disable-next-line ordering
    function attestationsAmount() external view returns (uint256) {
        return _roots.length;
    }

    /// @inheritdoc InterfaceDestination
    function getSignedAttestation(uint256 index)
        external
        view
        returns (bytes memory attPayload, bytes memory attSignature)
    {
        require(index < _roots.length, "Index out of range");
        bytes32 snapRoot = _roots[index];
        SnapRootData memory rootData = _rootData[snapRoot];
        StoredAttData memory storedAtt = _storedAttestations[index];
        attPayload = AttestationLib.formatAttestation({
            snapRoot_: snapRoot,
            agentRoot_: storedAtt.agentRoot,
            nonce_: rootData.attNonce,
            blockNumber_: rootData.attBN,
            timestamp_: rootData.attTS
        });
        attSignature = ByteString.formatSignature({r: storedAtt.r, s: storedAtt.s, v: rootData.notaryV});
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Opens a Dispute between a Guard and a Notary.
    /// This is overridden to allow disputes only between a Guard and a LOCAL Notary.
    function _openDispute(address guard, uint32 domain, address notary) internal override {
        // Only disputes for local Notaries could be initiated in Destination
        require(domain == localDomain, "Not a local Notary");
        super._openDispute(guard, domain, notary);
    }

    /// @dev Resolves a Dispute for a slashed agent, if it hasn't been done already.
    /// This is overridden to resolve disputes only between a Guard and a LOCAL Notary.
    function _resolveDispute(uint32 domain, address slashedAgent) internal virtual override {
        // Disputes could be only opened between a Guard and a local Notary
        if (domain == 0 || domain == localDomain) {
            super._resolveDispute(domain, slashedAgent);
        }
    }

    /// @dev Saves Agent Merkle Root from the accepted attestation, if there is
    /// no pending root to be passed to LightManager.
    /// Returns the updated "last snapshot root / last agent root" status struct.
    function _saveAgentRoot(bool rootPending, bytes32 agentRoot, address notary)
        internal
        returns (DestinationStatus memory status)
    {
        status = destStatus;
        // Update the timestamp for the latest snapshot root
        status.snapRootTime = uint48(block.timestamp);
        // Don't update agent root, if there is already a pending one
        // Update the data for latest agent root only if it differs from the saved one
        if (!rootPending && nextAgentRoot != agentRoot) {
            status.agentRootTime = uint48(block.timestamp);
            status.notary = notary;
            nextAgentRoot = agentRoot;
            emit AgentRootAccepted(agentRoot);
        }
    }
}
