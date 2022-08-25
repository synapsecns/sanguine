// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { LocalDomainContext } from "./context/LocalDomainContext.sol";
import { GlobalNotaryRegistry } from "./registry/GlobalNotaryRegistry.sol";
import { GuardRegistry } from "./registry/GuardRegistry.sol";
import { DestinationHub } from "./hubs/DestinationHub.sol";
import { SystemContract } from "./system/SystemContract.sol";
import { Version0 } from "./Version0.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
import { Header } from "./libs/Header.sol";
import { Tips } from "./libs/Tips.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
import { SystemMessage } from "./libs/SystemMessage.sol";
import { IMessageRecipient } from "./interfaces/IMessageRecipient.sol";
// ============ External Imports ============
import { TypedMemView } from "./libs/TypedMemView.sol";

/**
 * @title Destination
 * @notice Track merkle root state of Origin contracts on other chains,
 * prove and dispatch messages to end recipients.
 */
contract Destination is
    Version0,
    SystemContract,
    LocalDomainContext,
    DestinationHub,
    GlobalNotaryRegistry,
    GuardRegistry
{
    // ============ Libraries ============

    using Message for bytes;
    using Message for bytes29;
    using Header for bytes29;
    using TypedMemView for bytes29;

    /**
     * @notice Information stored for every blacklisted Notary.
     * TODO: finalize structure
     * @param isBlacklisted		Whether the Notary is blacklisted
     * @param guard				Guard who reported the Notary
     * @param blacklistedAt		Timestamp when Notary was blacklisted
     */
    struct Blacklist {
        address guard; // 160 bits
        uint96 blacklistedAt; // 96 bits
    }

    // ============ Public Storage ============

    // re-entrancy guard
    uint8 private entered;

    uint256 internal mirrorCount;

    // domain => [leaf => status]
    // Status is either NONE, EXECUTED (see below) or merkle root that was used for proving.
    mapping(uint32 => mapping(bytes32 => bytes32)) public messageStatus;

    // notary => blacklist info
    mapping(address => Blacklist) public blacklistedNotaries;

    // gap for upgrade safety
    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    // ============ Constants ============

    bytes32 internal constant MESSAGE_STATUS_NONE = bytes32(0);
    bytes32 internal constant MESSAGE_STATUS_EXECUTED = bytes32(uint256(1));

    // ============ Events ============

    /**
     * @notice Emitted when message is executed
     * @param messageHash The keccak256 hash of the message that was executed
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash);

    /**
     * @notice Emitted when a root's confirmation is modified by governance
     * @param root The root for which confirmAt has been set
     * @param previousConfirmAt The previous value of confirmAt
     * @param newConfirmAt The new value of confirmAt
     */
    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );

    event NotaryBlacklisted(
        address indexed notary,
        address indexed guard,
        address indexed reporter,
        bytes report
    );

    // ============ Constructor ============

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _localDomain) LocalDomainContext(_localDomain) {}

    // ============ Initializer ============

    /**
     * @notice Initialize the mirror
     * @dev Performs the following action:
     *      - initializes inherited contracts
     *      - initializes re-entrancy guard
     */
    function initialize() external initializer {
        __SystemContract_initialize();
        entered = 1;
    }

    // ============ External Functions ============

    /**
     * @notice First attempts to prove the validity of provided formatted
     * `message`. If the message is successfully proven, then tries to execute
     * message.
     * @dev Reverts if `prove` call returns false
     * @param _message Formatted message (refer to Message library)
     * @param _proof Merkle proof of inclusion for message's leaf
     * @param _index Index of leaf in origin's merkle tree
     */
    function proveAndExecute(
        uint32 _remoteDomain,
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) external {
        require(prove(_remoteDomain, _message, _proof, _index), "!prove");
        execute(_message);
    }

    /**
     * @notice Given formatted message, attempts to dispatch
     * message payload to end recipient.
     * @dev Recipient must implement a `handle` method (refer to IMessageRecipient.sol)
     * Reverts if formatted message's destination domain is not the Mirror's domain,
     * if message has not been proven,
     * or if recipient reverted upon receiving the message.
     * @param _message Formatted message
     */
    function execute(bytes memory _message) public {
        bytes29 messageView = _message.castToMessage();
        bytes29 header = messageView.header();
        uint32 originDomain = header.origin();
        // ensure message was meant for this domain
        require(header.destination() == _localDomain(), "!destination");
        // ensure message has been proven
        bytes32 messageHash = messageView.keccak();
        bytes32 root = messageStatus[originDomain][messageHash];
        require(
            root != MESSAGE_STATUS_NONE && root != MESSAGE_STATUS_EXECUTED,
            "!exists || executed"
        );
        require(
            acceptableRoot(originDomain, header.optimisticSeconds(), root),
            "!optimisticSeconds"
        );
        // check re-entrancy guard
        require(entered == 1, "!reentrant");
        entered = 0;
        _storeTips(messageView.tips());
        // update message status as executed
        messageStatus[originDomain][messageHash] = MESSAGE_STATUS_EXECUTED;
        address recipient = _checkForSystemMessage(header.recipient());
        IMessageRecipient(recipient).handle(
            originDomain,
            header.nonce(),
            header.sender(),
            mirrorRoots[originDomain][root].submittedAt,
            messageView.body().clone()
        );
        emit Executed(originDomain, messageHash);
        // reset re-entrancy guard
        entered = 1;
    }

    // ============ External Owner Functions ============

    /**
     * @notice Set Notary role
     * @dev MUST ensure that all roots signed by previous Notary have
     * been relayed before calling. Only callable by owner (Governance)
     * @param _notary New Notary
     */
    function setNotary(uint32 _domain, address _notary) external onlyOwner {
        // TODO: proper implementation
        _addNotary(_domain, _notary);
    }

    /**
     * @notice Set confirmAt for a given root
     * @dev To be used if in the case that fraud is proven
     * and roots need to be deleted / added. Only callable by owner (Governance)
     * @param _root The root for which to modify confirm time
     * @param _confirmAt The new confirmation time. Set to 0 to "delete" a root.
     */
    function setConfirmation(
        uint32 _originDomain,
        bytes32 _root,
        uint256 _confirmAt
    ) external onlyOwner {
        uint256 _previousConfirmAt = mirrorRoots[_originDomain][_root].submittedAt;
        mirrorRoots[_originDomain][_root].submittedAt = uint96(_confirmAt);
        emit SetConfirmation(_originDomain, _root, _previousConfirmAt, _confirmAt);
    }

    // ============ Public Functions ============

    /**
     * @notice Attempts to prove the validity of message given its leaf, the
     * merkle proof of inclusion for the leaf, and the index of the leaf.
     * @dev Reverts if message's MessageStatus != None (i.e. if message was
     * already proven or executed)
     * @dev For convenience, we allow proving against any previous root.
     * This means that witnesses never need to be updated for the new root
     * @param _message Formatted message
     * @param _proof Merkle proof of inclusion for leaf
     * @param _index Index of leaf in origin's merkle tree
     * @return Returns true if proof was valid and `prove` call succeeded
     **/
    function prove(
        uint32 _originDomain,
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) public returns (bool) {
        bytes32 _leaf = keccak256(_message);
        // ensure that mirror is active
        require(mirrors[_originDomain].latestNonce != 0, "Mirror not active");
        // ensure that message has not been proven or executed
        require(messageStatus[_originDomain][_leaf] == MESSAGE_STATUS_NONE, "!MessageStatus.None");
        // calculate the expected root based on the proof
        bytes32 _calculatedRoot = MerkleLib.branchRoot(_leaf, _proof, _index);
        // if the root is valid, save it for later optimistic period checking
        if (mirrorRoots[_originDomain][_calculatedRoot].submittedAt != 0) {
            messageStatus[_originDomain][_leaf] = _calculatedRoot;
            return true;
        }
        return false;
    }

    // ============ Internal Functions ============

    /**
     * @notice Blacklists Notary:
     * - New attestations signed by Notary are not accepted
     * - Any old roots attested by Notary can not be used for proving/executing
     * @dev _notary is always an active Notary, _guard is always an active Guard.
     * @param _domain   Origin domain where fraud was allegedly committed by Notary
     * @param _notary   Notary address who allegedly committed fraud attestation
     * @param _guard    Guard address that reported the Notary
     * @param _report   Payload with Report data and signature
     */
    function _blacklistNotary(
        uint32 _domain,
        address _notary,
        address _guard,
        bytes memory _report
    ) internal override {
        _removeNotary(_domain, _notary);
        emit NotaryBlacklisted(_notary, _guard, msg.sender, _report);
        blacklistedNotaries[_notary] = Blacklist({
            guard: _guard,
            blacklistedAt: uint96(block.timestamp)
        });
        // TODO: Send system message indicating that a Notary was reported?
    }

    function _checkForSystemMessage(bytes32 _recipient) internal view returns (address recipient) {
        // Check if SYSTEM_ROUTER was specified as message recipient
        if (_recipient == SystemMessage.SYSTEM_ROUTER) {
            /**
             * @dev Route message to SystemRouter.
             *      Note: Only SystemRouter contract on origin chain
             *      can send such a message (enforced in Origin.sol).
             */
            recipient = address(systemRouter);
        } else {
            // Cast bytes32 to address otherwise
            recipient = TypeCasts.bytes32ToAddress(_recipient);
        }
    }

    function _storeTips(bytes29 _tips) internal virtual {
        // TODO: implement storing & claiming logic
    }
}
