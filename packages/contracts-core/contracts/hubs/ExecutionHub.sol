// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation, ExecutionAttestation } from "../libs/Attestation.sol";
import { SYSTEM_ROUTER, TREE_DEPTH } from "../libs/Constants.sol";
import { MerkleLib } from "../libs/Merkle.sol";
import { Header, Message, MessageLib, Tips } from "../libs/Message.sol";
import { TypeCasts } from "../libs/TypeCasts.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DisputeHub } from "./DisputeHub.sol";
import { ExecutionHubEvents } from "../events/ExecutionHubEvents.sol";
import { IExecutionHub } from "../interfaces/IExecutionHub.sol";
import { IMessageRecipient } from "../interfaces/IMessageRecipient.sol";
import { SystemRegistry } from "../system/SystemRegistry.sol";

/**
 * @notice ExecutionHub is responsible for executing the messages that are
 * proven against the Snapshot Merkle Roots.
 * The Snapshot Merkle Roots themselves are supposed to be dealt with in the child contracts.
 * On the Synapse Chain Notaries are submitting the snapshots that are later used for proving.
 * On the other chains Notaries are submitting the attestations that are later used for proving.
 */
abstract contract ExecutionHub is DisputeHub, SystemRegistry, ExecutionHubEvents, IExecutionHub {
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

    /// @dev Tracks all saved attestations
    // (root => attestation)
    mapping(bytes32 => ExecutionAttestation) private rootAttestations;

    /// @dev gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           EXECUTE MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IExecutionHub
    function execute(
        bytes memory _message,
        bytes32[TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) external {
        // This will revert if payload is not a formatted message payload
        Message message = _message.castToMessage();
        Header header = message.header();
        bytes32 msgLeaf = message.leaf();
        // Check proofs validity and mark message as executed
        ExecutionAttestation memory execAtt = _prove(
            header,
            msgLeaf,
            _originProof,
            _snapProof,
            _stateIndex
        );
        // Store message tips
        Tips tips = message.tips();
        _storeTips(execAtt.notary, tips);
        // Get the specified recipient address
        uint32 origin = header.origin();
        address recipient = _checkForSystemRouter(header.recipient());
        // Pass the message to the recipient
        IMessageRecipient(recipient).handle(
            origin,
            header.nonce(),
            header.sender(),
            execAtt.submittedAt,
            message.body().clone()
        );
        emit Executed(origin, msgLeaf);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                  INTERNAL LOGIC: MESSAGE EXECUTION                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
     * @return execAtt      Attestation data for derived snapshot root
     */
    function _prove(
        Header _header,
        bytes32 _msgLeaf,
        bytes32[TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) internal returns (ExecutionAttestation memory execAtt) {
        // TODO: split into a few smaller functions?
        // Check that message has not been executed before
        require(messageStatus[_msgLeaf] == MESSAGE_STATUS_NONE, "!MessageStatus.None");
        // Ensure message was meant for this domain
        require(_header.destination() == localDomain, "!destination");
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        bytes32 originRoot = MerkleLib.branchRoot(_msgLeaf, _originProof, _header.nonce() - 1);
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if state index is out of range
        bytes32 snapshotRoot = _snapshotRoot(originRoot, _header.origin(), _snapProof, _stateIndex);
        // Fetch the attestation data for the snapshot root
        execAtt = rootAttestations[snapshotRoot];
        // Check if snapshot root has been submitted
        require(!execAtt.isEmpty(), "Invalid snapshot root");
        // Check that snapshot proof length matches the height of Snapshot Merkle Tree
        require(_snapProof.length == execAtt.height, "Invalid proof length");
        // Check if Notary who submitted the attestation is still active
        // TODO: check for dispute status instead
        require(_isActiveAgent(localDomain, execAtt.notary), "Inactive notary");
        // Check if optimistic period has passed
        require(
            block.timestamp >= _header.optimisticSeconds() + execAtt.submittedAt,
            "!optimisticSeconds"
        );
        // Mark message as executed against the snapshot root
        messageStatus[_msgLeaf] = snapshotRoot;
    }

    /// @dev Saves a snapshot root with the attestation data provided by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _saveAttestation(Attestation _att, address _notary) internal {
        bytes32 root = _att.root();
        require(rootAttestations[root].isEmpty(), "Root already exists");
        rootAttestations[root] = _att.toExecutionAttestation(_notary);
    }

    /// @dev Gets a saved attestation for the given snapshot root.
    /// Will return an empty struct, if the snapshot root hasn't been previously saved.
    function _getRootAttestation(bytes32 _root)
        internal
        view
        returns (ExecutionAttestation memory)
    {
        return rootAttestations[_root];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         INTERNAL LOGIC: TIPS                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _storeTips(address _notary, Tips _tips) internal {
        // TODO: implement tips logic
        emit TipsStored(_notary, _tips.unwrap().clone());
    }
}
