// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { SYSTEM_ROUTER } from "./libs/Constants.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Header, Message, MessageLib, Tips } from "./libs/Message.sol";
import { StateLib } from "./libs/State.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DomainContext } from "./context/DomainContext.sol";
import { DestinationEvents } from "./events/DestinationEvents.sol";
import { InterfaceDestination, ORIGIN_TREE_DEPTH } from "./interfaces/InterfaceDestination.sol";
import { IMessageRecipient } from "./interfaces/IMessageRecipient.sol";
import { DestinationAttestation, AttestationHub } from "./hubs/AttestationHub.sol";
import { Attestation, AttestationReport, StatementHub } from "./hubs/StatementHub.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Destination is
    StatementHub,
    AttestationHub,
    SystemRegistry,
    DestinationEvents,
    InterfaceDestination
{
    using MessageLib for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    bytes32 internal constant MESSAGE_STATUS_NONE = bytes32(0);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice (messageHash => status)
    /// TODO: Store something else as "status"? Notary/timestamp?
    /// - Message hasn't been executed: MESSAGE_STATUS_NONE
    /// - Message has been executed: snapshot root used for proving when executed
    /// @dev Messages coming from different origins will always have a different hash
    /// as origin domain is encoded into the formatted message.
    /// Thus we can use hash as a key instead of an (origin, hash) tuple.
    mapping(bytes32 => bytes32) public messageStatus;

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
        // This will revert if payload is not an attestation, or signer is not an active Notary
        (Attestation att, uint32 domain, address notary) = _verifyAttestation(
            _attPayload,
            _attSignature
        );
        // Check that Notary is active on local domain
        require(domain == localDomain, "Wrong Notary domain");
        // This will revert if snapshot root has been previously submitted
        _acceptAttestation(att, notary);
        emit AttestationAccepted(domain, notary, _attPayload, _attSignature);
        return true;
    }

    /// @inheritdoc InterfaceDestination
    function submitAttestationReport(
        bytes memory _arPayload,
        bytes memory _arSignature,
        bytes memory _attSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not an attestation report, or report signer is not an active Guard
        (AttestationReport report, address guard) = _verifyAttestationReport(
            _arPayload,
            _arSignature
        );
        // This will revert if attestation signer is not an active Notary
        (uint32 domain, address notary) = _verifyAttestation(report.attestation(), _attSignature);
        _openDispute(guard, domain, notary);
        return true;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           EXECUTE MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceDestination
    function execute(
        bytes memory _message,
        bytes32[ORIGIN_TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) external {
        // This will revert if payload is not a formatted message payload
        Message message = _message.castToMessage();
        Header header = message.header();
        bytes32 msgLeaf = message.leaf();
        // Check proofs validity and mark message as executed
        DestinationAttestation memory destAtt = _prove(
            header,
            msgLeaf,
            _originProof,
            _snapProof,
            _stateIndex
        );
        // Store message tips
        Tips tips = message.tips();
        _storeTips(destAtt.notary, tips);
        // Get the specified recipient address
        uint32 origin = header.origin();
        address recipient = _checkForSystemRouter(header.recipient());
        // Pass the message to the recipient
        IMessageRecipient(recipient).handle(
            origin,
            header.nonce(),
            header.sender(),
            destAtt.destTimestamp,
            message.body().clone()
        );
        emit Executed(origin, msgLeaf);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Attempts to prove the validity of the cross-chain message.
     * First, the origin Merkle Root is reconstructed using the origin proof.
     * Then the origin state's "left leaf" is reconstructed using the origin domain.
     * After that the snapshot Merkle Root is reconstructed using the snapshot proof.
     * Finally, the optimistic period is checked for the derived snapshot root.
     * @dev Reverts if any of the checks fail.
     * @param _header       Typed memory view over message header payload
     * @param _msgLeaf      Message Leaf that was inserted in the Origin Merkle Tree
     * @param _originProof  Proof of inclusion of Message Leaf in the Origin Merkle Tree
     * @param _snapProof    Proof of inclusion of Origin State Left Leaf into Snapshot Merkle Tree
     * @param _stateIndex   Index of Origin State in the Snapshot
     * @return destAtt      Attestation data for derived snapshot root
     */
    function _prove(
        Header _header,
        bytes32 _msgLeaf,
        bytes32[ORIGIN_TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) internal returns (DestinationAttestation memory destAtt) {
        // TODO: split into a few smaller functions?
        // Check that message has not been executed before
        require(messageStatus[_msgLeaf] == MESSAGE_STATUS_NONE, "!MessageStatus.None");
        // Ensure message was meant for this domain
        require(_header.destination() == localDomain, "!destination");
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        bytes32 originRoot = MerkleLib.branchRoot(_msgLeaf, _originProof, _header.nonce() - 1);
        // Reconstruct left sub-leaf of the Origin State: (merkleRoot, originDomain)
        bytes32 leftLeaf = StateLib.leftLeaf(originRoot, _header.origin());
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // Index of "leftLeaf" is twice the state position in the snapshot
        /// @dev We ask to provide state index instead of "leftLeaf" index to enforce
        /// choice of State's left leaf for root reconstruction
        bytes32 snapshotRoot = MerkleLib.branchRoot(leftLeaf, _snapProof, _stateIndex << 1);
        // Fetch the attestation data for the snapshot root
        destAtt = _rootAttestation(snapshotRoot);
        // Check if snapshot root has been submitted
        require(!destAtt.isEmpty(), "Invalid snapshot root");
        // Check that snapshot proof length matches the height of Snapshot Merkle Tree
        require(_snapProof.length == destAtt.height, "Invalid proof length");
        // Check if Notary who submitted the snapshot is still active
        require(_isActiveAgent(localDomain, destAtt.notary), "Inactive notary");
        // Check if optimistic period has passed
        require(
            block.timestamp >= _header.optimisticSeconds() + destAtt.destTimestamp,
            "!optimisticSeconds"
        );
        // Mark message as executed against the snapshot root
        messageStatus[_msgLeaf] = snapshotRoot;
    }

    function _storeTips(address _notary, Tips _tips) internal {
        // TODO: implement tips logic
        emit TipsStored(_notary, _tips.unwrap().clone());
    }

    /**
     * @notice Returns adjusted "recipient" field.
     * @dev By default, "recipient" field contains the recipient address padded to 32 bytes.
     * But if SYSTEM_ROUTER value is used for "recipient" field, recipient is Synapse Router.
     * Note: tx will revert in Origin if anyone but SystemRouter uses SYSTEM_ROUTER as recipient.
     */
    function _checkForSystemRouter(bytes32 _recipient) internal view returns (address recipient) {
        // Check if SYSTEM_ROUTER was specified as message recipient
        if (_recipient == SYSTEM_ROUTER) {
            /**
             * @dev Route message to SystemRouter.
             * Note: Only SystemRouter contract on origin chain can send a message
             * using SYSTEM_ROUTER as "recipient" field (enforced in Origin.sol).
             */
            recipient = address(systemRouter);
        } else {
            // Cast bytes32 to address otherwise
            recipient = TypeCasts.bytes32ToAddress(_recipient);
        }
    }

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
