// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { LocalDomainContext } from "./context/LocalDomainContext.sol";
import { DestinationHub } from "./hubs/DestinationHub.sol";
import { SystemContract } from "./system/SystemContract.sol";
import { DestinationEvents } from "./events/DestinationEvents.sol";
import { Version0 } from "./Version0.sol";
import { IMessageRecipient } from "./interfaces/IMessageRecipient.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
import { Header } from "./libs/Header.sol";
import { Tips } from "./libs/Tips.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
import { SystemMessage } from "./libs/SystemMessage.sol";

/**
 * @title Destination
 * @notice Track merkle root state of Origin contracts on other chains,
 * prove and dispatch messages to end recipients.
 */
contract Destination is
    Version0,
    DestinationEvents,
    SystemContract,
    LocalDomainContext,
    DestinationHub
{
    using Message for bytes;
    using Message for bytes29;
    using Header for bytes29;
    using TypedMemView for bytes29;

    /**
     * @notice Information stored for every blacklisted Notary.
     * TODO: finalize structure
     * @param guard				Guard who reported the Notary
     * @param blacklistedAt		Timestamp when Notary was blacklisted
     */
    struct Blacklist {
        address guard; // 160 bits
        uint96 blacklistedAt; // 96 bits
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    bytes32 internal constant MESSAGE_STATUS_NONE = bytes32(0);

    // We're using uint256 instead of bool/uint8 here, because reading/writing is
    // cheaper for the types that are using a full word of storage
    uint256 internal constant NOT_ENTERED = 1;
    // TODO: set up execute() tests with message recipients actually spending some gas
    // to determine if ENTERED = 0 or ENTERED = 2 leads to lower median gas usage
    uint256 internal constant ENTERED = 0;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // re-entrancy guard
    uint256 private status;

    // domain => [leaf => status]
    // Message wasn't executed => MESSAGE_STATUS_NONE
    // Message was executed => merkle root used for proving when executed
    mapping(uint32 => mapping(bytes32 => bytes32)) public messageStatus;

    // notary => blacklist info
    mapping(address => Blacklist) public blacklistedNotaries;

    // gap for upgrade safety
    uint256[47] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) LocalDomainContext(_domain) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Initialize the mirror
     * @dev Performs the following action:
     *      - initializes inherited contracts
     *      - initializes re-entrancy guard
     */
    function initialize() external initializer {
        __SystemContract_initialize();
        status = NOT_ENTERED;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    EXTERNAL FUNCTIONS: RESTRICTED                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Set Notary role
     * @dev MUST ensure that all roots signed by previous Notary have
     * been relayed before calling. Only callable by owner (Governance)
     * @param _notary New Notary
     */
    function setNotary(uint32 _domain, address _notary) external onlyOwner {
        // TODO: proper implementation
        // Notaries and Guards should be added/removed by a BondingManager contract
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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Given formatted message, attempts to dispatch
     * message payload to end recipient.
     * @dev Recipient must implement a `handle` method (refer to IMessageRecipient.sol)
     * Reverts if formatted message's destination domain is not the Destination's domain,
     * if message proof is invalid, or its optimistic period not yet passed.
     * Also reverts if the recipient reverted upon receiving the message.
     * @param _message  Formatted message
     * @param _proof    Merkle proof of inclusion for message's leaf
     * @param _index    Index of leaf in origin's merkle tree
     */
    function execute(
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) external {
        bytes29 message = _message.castToMessage();
        bytes29 header = message.header();
        uint32 originDomain = header.origin();
        // ensure message was meant for this domain
        require(header.destination() == _localDomain(), "!destination");
        bytes32 leaf = message.keccak();
        // ensure message can be proven against a confirmed root,
        // and that message's optimistic period has passed
        bytes32 root = _prove(originDomain, leaf, _proof, _index, header.optimisticSeconds());
        // check re-entrancy guard
        require(status == NOT_ENTERED, "!reentrant");
        status = ENTERED;
        _storeTips(message.tips());
        // it should not be possible to construct a merkle tree with a root = 0x0, but even then
        // attestations with empty root would be rejected: see DestinationHub._handleAttestation()
        // update message status as executed, new status is never bytes32(0)
        messageStatus[originDomain][leaf] = root;
        address recipient = _checkForSystemMessage(header.recipient());
        IMessageRecipient(recipient).handle(
            originDomain,
            header.nonce(),
            header.sender(),
            mirrorRoots[originDomain][root].submittedAt,
            message.body().clone()
        );
        emit Executed(originDomain, leaf);
        // reset re-entrancy guard
        status = NOT_ENTERED;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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

    // solhint-disable-next-line no-empty-blocks
    function _storeTips(bytes29 _tips) internal virtual {
        // TODO: implement storing & claiming logic
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _checkForSystemMessage(bytes32 _recipient) internal view returns (address recipient) {
        // Check if SYSTEM_ROUTER was specified as message recipient
        if (_recipient == SystemMessage.SYSTEM_ROUTER) {
            /**
             * @dev Route message to SystemRouter.
             * Note: Only SystemRouter contract on origin chain
             * can send such a message (enforced in Origin.sol).
             */
            recipient = address(systemRouter);
        } else {
            // Cast bytes32 to address otherwise
            recipient = TypeCasts.bytes32ToAddress(_recipient);
        }
    }

    /**
     * @notice Attempts to prove the validity of message given its leaf, the
     * merkle proof of inclusion for the leaf, and the index of the leaf.
     * @dev Reverts if message's MessageStatus != None (i.e. if message was
     * already proven or executed)
     * @dev For convenience, we allow proving against any previous root.
     * This means that witnesses never need to be updated for the new root
     * @param _originDomain         Domain of Origin
     * @param _leaf                 Leaf (hash) of the message
     * @param _proof                Merkle proof of inclusion for leaf
     * @param _index                Index of leaf in Origin's merkle tree
     * @param _optimisticSeconds    Optimistic period of the message
     * @return root                 Merkle root used for proving message inclusion
     **/
    function _prove(
        uint32 _originDomain,
        bytes32 _leaf,
        bytes32[32] calldata _proof,
        uint256 _index,
        uint32 _optimisticSeconds
    ) internal view returns (bytes32 root) {
        // ensure that mirror is active
        require(mirrors[_originDomain].latestNonce != 0, "Mirror not active");
        // ensure that message has not been executed
        require(messageStatus[_originDomain][_leaf] == MESSAGE_STATUS_NONE, "!MessageStatus.None");
        // calculate the expected root based on the proof
        root = MerkleLib.branchRoot(_leaf, _proof, _index);
        // Sanity check: this either returns true or reverts
        assert(acceptableRoot(_originDomain, _optimisticSeconds, root));
    }
}
