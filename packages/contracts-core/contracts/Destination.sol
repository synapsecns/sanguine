// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "./libs/Merkle.sol";
import "./libs/Message.sol";
import "./libs/State.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DomainContext } from "./context/DomainContext.sol";
import { InterfaceDestination, ORIGIN_TREE_DEPTH } from "./interfaces/InterfaceDestination.sol";
import { DestinationAttestation, AttestationHub } from "./hubs/AttestationHub.sol";
import { Attestation, StatementHub } from "./hubs/StatementHub.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Destination is StatementHub, AttestationHub, SystemRegistry, InterfaceDestination {
    // TODO: Attach library functions to custom types globally
    using HeaderLib for Header;
    using MessageLib for Message;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: move Events to a separate contract, once the old Destination is deprecated

    /**
     * @notice Emitted when a snapshot is accepted by the Destination contract.
     * @param domain        Domain where the signed Notary is active
     * @param notary        Notary who signed the attestation
     * @param attestation   Raw payload with attestation data
     * @param attSignature  Notary signature for the attestation
     */
    event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature);

    /**
     * @notice Emitted when message is executed.
     * @param remoteDomain  Remote domain where message originated
     * @param messageHash   The keccak256 hash of the message that was executed
     * @param tips          Raw payload with tips paid for the off-chain agents
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash, bytes tips);

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           EXECUTE MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceDestination
    function execute(
        bytes memory _message,
        bytes32[ORIGIN_TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _snapIndex
    ) external {
        // TODO: implement
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
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

    /**
     * @notice Attempts to prove the validity of the cross-chain message.
     * First, the origin Merkle Root is reconstructed using the origin proof.
     * Then the origin state's "left leaf" is reconstructed using the origin domain.
     * After that the snapshot Merkle Root is reconstructed using the snapshot proof.
     * Finally, the optimistic period is checked for the derived snapshot root.
     * @dev Reverts if any of the checks fail.
     * @param _msg          Typed memory view over message payload
     * @param _originProof  Proof of inclusion of message in the Origin Merkle Tree
     * @param _snapProof    Proof of inclusion of Origin State Left Leaf into Snapshot Merkle Tree
     * @param _snapIndex    Index of Origin State Left Leaf in the Snapshot Merkle Tree
     * @return snapshotRoot Derived merkle root of the Snapshot Merkle Tree
     * @return destAtt      Rest of attestation data that Destination keeps track of
     */
    function _prove(
        Message _msg,
        bytes32[ORIGIN_TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _snapIndex
    ) internal view returns (bytes32 snapshotRoot, DestinationAttestation memory destAtt) {
        Header header = _msg.header();
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        bytes32 originRoot = MerkleLib.branchRoot(_msg.leaf(), _originProof, header.nonce() - 1);
        // Reconstruct left sub-leaf of the Origin State: (merkleRoot, originDomain)
        bytes32 leftLeaf = StateLib.leftLeaf(originRoot, header.origin());

        // TODO: implement branchRoot function that takes dynamic sized proof as variable
        // snapshotRoot = MerkleLib.branchRoot(leftLeaf, _snapProof, _snapIndex);

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
            block.timestamp >= header.optimisticSeconds() + destAtt.destTimestamp,
            "!optimisticSeconds"
        );
    }
}
