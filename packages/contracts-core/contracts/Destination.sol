// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "./libs/memory/Attestation.sol";
import {ByteString} from "./libs/memory/ByteString.sol";
import {AGENT_ROOT_OPTIMISTIC_PERIOD, SYNAPSE_DOMAIN} from "./libs/Constants.sol";
import {IndexOutOfRange, NotaryInDispute} from "./libs/Errors.sol";
import {ChainGas, GasData} from "./libs/stack/GasData.sol";
import {AgentStatus, DestinationStatus} from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "./base/AgentSecured.sol";
import {DestinationEvents} from "./events/DestinationEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceDestination} from "./interfaces/InterfaceDestination.sol";
import {InterfaceLightManager} from "./interfaces/InterfaceLightManager.sol";
import {IStatementInbox} from "./interfaces/IStatementInbox.sol";
import {ExecutionHub} from "./hubs/ExecutionHub.sol";

/// @notice `Destination` contract is used for receiving messages from other chains. It relies on
/// Notary-signed statements to get the truthful states of the remote chains. These states are then
/// used to verify the validity of the messages sent from the remote chains.
/// `Destination` is responsible for the following:
/// - Accepting the Attestations from the local Inbox contract.
/// - Using these Attestations to execute the messages (see parent `ExecutionHub`).
/// - Passing the Agent Merkle Roots from the Attestations to the local LightManager contract,
///   if deployed on a non-Synapse chain.
/// - Keeping track of the remote domains GasData submitted by Notaries, that could be later consumed
///   by the local `GasOracle` contract.
contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    using AttestationLib for bytes;
    using ByteString for bytes;

    // TODO: this could be further optimized in terms of storage
    struct StoredAttData {
        bytes32 agentRoot;
        bytes32 dataHash;
    }

    struct StoredGasData {
        GasData gasData;
        uint32 notaryIndex;
        uint40 submittedAt;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev Invariant: this is either current LightManager root,
    /// or the pending root to be passed to LightManager once its optimistic period is over.
    bytes32 internal _nextAgentRoot;

    /// @inheritdoc InterfaceDestination
    DestinationStatus public destStatus;

    /// @dev Stored lookup data for all accepted Notary Attestations
    StoredAttData[] internal _storedAttestations;

    /// @dev Remote domains GasData submitted by Notaries
    mapping(uint32 => StoredGasData) internal _storedGasData;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain, address agentManager_, address inbox_)
        AgentSecured("0.0.3", domain, agentManager_, inbox_)
    {} // solhint-disable-line no-empty-blocks

    /// @notice Initializes Destination contract:
    /// - msg.sender is set as contract owner
    function initialize(bytes32 agentRoot) external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        // Initialize ReeentrancyGuard
        __ReentrancyGuard_init();
        // Set Agent Merkle Root in Light Manager
        if (localDomain != SYNAPSE_DOMAIN) {
            _nextAgentRoot = agentRoot;
            InterfaceLightManager(address(agentManager)).setAgentRoot(agentRoot);
            destStatus.agentRootTime = uint40(block.timestamp);
        }
        // No need to do anything on Synapse Chain, as the agent root is set in BondingManager
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function acceptAttestation(
        uint32 notaryIndex,
        uint256 sigIndex,
        bytes memory attPayload,
        bytes32 agentRoot,
        ChainGas[] memory snapGas
    ) external onlyInbox returns (bool wasAccepted) {
        if (_isInDispute(notaryIndex)) revert NotaryInDispute();
        // First, try passing current agent merkle root
        (bool rootPassed, bool rootPending) = passAgentRoot();
        // Don't accept attestation, if the agent root was updated in LightManager,
        // as the following agent check will fail.
        if (rootPassed) return false;
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, notaryIndex, sigIndex);
        _storedAttestations.push(StoredAttData({agentRoot: agentRoot, dataHash: att.dataHash()}));
        // Save Agent Root if required, and update the Destination's Status
        destStatus = _saveAgentRoot(rootPending, agentRoot, notaryIndex);
        _saveGasData(snapGas, notaryIndex);
        return true;
    }

    // ═══════════════════════════════════════════ AGENT ROOT QUARANTINE ═══════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function passAgentRoot() public returns (bool rootPassed, bool rootPending) {
        // Agent root is not passed on Synapse Chain, as it could be accessed via BondingManager
        if (localDomain == SYNAPSE_DOMAIN) return (false, false);
        bytes32 oldRoot = IAgentManager(agentManager).agentRoot();
        bytes32 newRoot = _nextAgentRoot;
        // Check if agent root differs from the current one in LightManager
        if (oldRoot == newRoot) return (false, false);
        DestinationStatus memory status = destStatus;
        // Invariant: Notary who supplied `newRoot` was registered as active against `oldRoot`
        // So we just need to check the Dispute status of the Notary
        if (_isInDispute(status.notaryIndex)) {
            // Remove the pending agent merkle root, as its signer is in dispute
            _nextAgentRoot = oldRoot;
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
    function getAttestation(uint256 index) external view returns (bytes memory attPayload, bytes memory attSignature) {
        if (index >= _roots.length) revert IndexOutOfRange();
        bytes32 snapRoot = _roots[index];
        SnapRootData memory rootData = _rootData[snapRoot];
        StoredAttData memory storedAtt = _storedAttestations[index];
        attPayload = AttestationLib.formatAttestation({
            snapRoot_: snapRoot,
            dataHash_: storedAtt.dataHash,
            nonce_: rootData.attNonce,
            blockNumber_: rootData.attBN,
            timestamp_: rootData.attTS
        });
        // Attestation signatures are not required on Synapse Chain, as the attestations could be accessed via Summit.
        if (localDomain != SYNAPSE_DOMAIN) {
            attSignature = IStatementInbox(inbox).getStoredSignature(rootData.sigIndex);
        }
    }

    /// @inheritdoc InterfaceDestination
    function getGasData(uint32 domain) external view returns (GasData gasData, uint256 dataMaturity) {
        StoredGasData memory storedGasData = _storedGasData[domain];
        // Check if there is a stored gas data for the domain, and if the notary who provided the data is not in dispute
        if (storedGasData.submittedAt != 0 && !_isInDispute(storedGasData.notaryIndex)) {
            gasData = storedGasData.gasData;
            dataMaturity = block.timestamp - storedGasData.submittedAt;
        }
        // Return empty values if there is no data for the domain, or if the notary who provided the data is in dispute
    }

    /// @inheritdoc InterfaceDestination
    function nextAgentRoot() external view returns (bytes32) {
        // Return current agent root on Synapse Chain for consistency
        return localDomain == SYNAPSE_DOMAIN ? IAgentManager(agentManager).agentRoot() : _nextAgentRoot;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Saves Agent Merkle Root from the accepted attestation, if there is
    /// no pending root to be passed to LightManager.
    /// Returns the updated "last snapshot root / last agent root" status struct.
    function _saveAgentRoot(bool rootPending, bytes32 agentRoot, uint32 notaryIndex)
        internal
        returns (DestinationStatus memory status)
    {
        status = destStatus;
        // Update the timestamp for the latest snapshot root
        status.snapRootTime = uint40(block.timestamp);
        // No need to save agent roots on Synapse Chain, as they could be accessed via BondingManager
        // Don't update agent root, if there is already a pending one
        // Update the data for latest agent root only if it differs from the saved one
        if (localDomain != SYNAPSE_DOMAIN && !rootPending && _nextAgentRoot != agentRoot) {
            status.agentRootTime = uint40(block.timestamp);
            status.notaryIndex = notaryIndex;
            _nextAgentRoot = agentRoot;
            emit AgentRootAccepted(agentRoot);
        }
    }

    /// @dev Saves updated values from the snapshot's gas data list.
    function _saveGasData(ChainGas[] memory snapGas, uint32 notaryIndex) internal {
        uint256 statesAmount = snapGas.length;
        for (uint256 i = 0; i < statesAmount; i++) {
            ChainGas chainGas = snapGas[i];
            uint32 domain = chainGas.domain();
            // Don't save gas data for the local domain
            if (domain == localDomain) continue;
            StoredGasData memory storedGasData = _storedGasData[domain];
            // Check that the gas data is not already saved
            GasData gasData = chainGas.gasData();
            if (GasData.unwrap(gasData) == GasData.unwrap(storedGasData.gasData)) continue;
            // Save the gas data
            _storedGasData[domain] =
                StoredGasData({gasData: gasData, notaryIndex: notaryIndex, submittedAt: uint40(block.timestamp)});
        }
    }
}
