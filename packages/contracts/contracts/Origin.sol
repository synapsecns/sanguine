// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Version0 } from "./Version0.sol";
import { DomainNotaryRegistry } from "./registry/DomainNotaryRegistry.sol";
import { GuardRegistry } from "./registry/GuardRegistry.sol";
import { Attestation } from "./libs/Attestation.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Header } from "./libs/Header.sol";
import { Message } from "./libs/Message.sol";
import { Tips } from "./libs/Tips.sol";
import { SystemMessage } from "./system/SystemMessage.sol";
import { SystemContract } from "./system/SystemContract.sol";
import { MerkleTreeManager } from "./Merkle.sol";
import { INotaryManager } from "./interfaces/INotaryManager.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
// ============ External Imports ============
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Origin
 * @author Illusory Systems Inc.
 * @notice Accepts messages to be dispatched to remote chains,
 * constructs a Merkle tree of the messages,
 * and accepts signatures from a bonded Notary
 * which notarize the Merkle tree roots.
 * Accepts submissions of fraudulent signatures
 * by the Notary and slashes the Notary in this case.
 */
contract Origin is
    Version0,
    MerkleTreeManager,
    SystemContract,
    DomainNotaryRegistry,
    GuardRegistry
{
    // ============ Libraries ============

    using Attestation for bytes29;
    using MerkleLib for MerkleLib.Tree;

    using Tips for bytes;
    using Tips for bytes29;

    // ============ Enums ============

    // States:
    //   0 - UnInitialized - before initialize function is called
    //   note: the contract is initialized at deploy time, so it should never be in this state
    //   1 - Active - as long as the contract has not become fraudulent
    //   2 - Failed - after a valid fraud proof has been submitted;
    //   contract will no longer accept new messages
    enum States {
        UnInitialized,
        Active,
        Failed
    }

    // ============ Constants ============

    // Maximum bytes per message = 2 KiB
    // (somewhat arbitrarily set to begin)
    uint256 public constant MAX_MESSAGE_BODY_BYTES = 2 * 2**10;

    // ============ Public Storage Variables ============

    // domain => next available nonce for the domain
    uint32 public nonce;
    // contract responsible for Notary bonding, slashing and rotation
    INotaryManager public notaryManager;
    // Current state of contract
    States public state;

    // ============ Upgrade Gap ============

    // gap for upgrade safety
    uint256[47] private __GAP;

    // ============ Events ============

    /**
     * @notice Emitted when a new message is dispatched via Nomad
     * @param messageHash Hash of message; the leaf inserted to the Merkle tree
     *        for the message
     * @param leafIndex Index of message's leaf in merkle tree
     * @param destinationAndNonce Destination and destination-specific
     *        nonce combined in single field ((destination << 32) & nonce)
     * @param tips Tips paid for the remote off-chain agents
     * @param message Raw bytes of message
     */
    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes tips,
        bytes message
    );

    /**
     * @notice Emitted when proof of an improper attestation is submitted,
     * which sets the contract to FAILED state
     * @param notary       Notary who signed improper attestation
     * @param attestation   Attestation data and signature
     */
    event ImproperAttestation(address notary, bytes attestation);

    /**
     * @notice Emitted when the Notary is slashed
     * (should be paired with ImproperAttestation event)
     * @param notary The address of the notary
     * @param reporter The address of the entity that reported the notary misbehavior
     */
    event NotarySlashed(address indexed notary, address indexed reporter);

    /**
     * @notice Emitted when the NotaryManager contract is changed
     * @param notaryManager The address of the new notaryManager
     */
    event NewNotaryManager(address notaryManager);

    // ============ Constructor ============

    constructor(uint32 _localDomain)
        SystemContract(_localDomain)
        DomainNotaryRegistry(_localDomain)
    {} // solhint-disable-line no-empty-blocks

    // ============ Initializer ============

    function initialize(INotaryManager _notaryManager) public initializer {
        __SystemContract_initialize();
        _setNotaryManager(_notaryManager);
        _addNotary(notaryManager.notary());
        state = States.Active;
        // insert a historical root so nonces start at 1 rather then 0. Here we insert the default root of a sparse merkle tree
        historicalRoots.push(hex"27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757");
    }

    // ============ Modifiers ============

    /**
     * @notice Ensures that function is called by the NotaryManager contract
     */
    modifier onlyNotaryManager() {
        require(msg.sender == address(notaryManager), "!notaryManager");
        _;
    }

    /**
     * @notice Ensures that contract state != FAILED when the function is called
     */
    modifier notFailed() {
        require(state != States.Failed, "failed state");
        _;
    }

    // ============ External: Notary & NotaryManager Configuration  ============

    /**
     * @notice Set a new Notary
     * @dev To be set when rotating Notary after Fraud
     * @param _notary the new Notary
     */
    function setNotary(address _notary) external onlyNotaryManager {
        /**
         * TODO: do this properly
         * @dev 1. New Notaries should be added to all System Contracts
         *      from "secondary" Bonding contracts (global Notary/Guard registry)
         *      1a. onlyUpdaterManager -> onlyBondingManager (or w/e the name would be)
         *      2. There is supposed to be more than one active Notary
         *      2a. setUpdater() -> addNotary()
         *      3. No need to reset the `state`, as Origin is not supposed
         *      to be in Failed state in the first place (see _fail() for reasoning).
         */
        _addNotary(_notary);
        // set the Origin state to Active
        // now that Notary has been rotated
        state = States.Active;
    }

    /**
     * @notice Set a new NotaryManager contract
     * @dev Origin(s) will initially be initialized using a trusted NotaryManager contract;
     * we will progressively decentralize by swapping the trusted contract with a new implementation
     * that implements Notary bonding & slashing, and rules for Notary selection & rotation
     * @param _notaryManager the new NotaryManager contract
     */
    function setNotaryManager(address _notaryManager) external onlyOwner {
        _setNotaryManager(INotaryManager(_notaryManager));
    }

    // ============ External Functions  ============

    /**
     * @notice Dispatch the message to the destination domain & recipient
     * @dev Format the message, insert its hash into Merkle tree,
     * enqueue the new Merkle root, and emit `Dispatch` event with message information.
     * @param _destinationDomain Domain of destination chain
     * @param _recipientAddress Address of recipient on destination chain as bytes32
     * @param _messageBody Raw bytes content of message
     */
    function dispatch(
        uint32 _destinationDomain,
        bytes32 _recipientAddress,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable notFailed {
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        require(_tips.tipsView().totalTips() == msg.value, "!tips");
        // get the next nonce for the destination domain, then increment it
        nonce = nonce + 1;
        bytes32 _sender = _checkForSystemMessage(_recipientAddress);
        // format the message into packed bytes
        bytes memory _header = Header.formatHeader(
            localDomain,
            _sender,
            nonce,
            _destinationDomain,
            _recipientAddress,
            _optimisticSeconds
        );
        // format the message into packed bytes
        bytes memory _message = Message.formatMessage(_header, _tips, _messageBody);
        // insert the hashed message into the Merkle tree
        bytes32 _messageHash = keccak256(_message);
        // new root is added to the historical roots
        _insertHash(_messageHash);
        // Emit Dispatch event with message information
        // note: leafIndex is count() - 1 since new leaf has already been inserted
        emit Dispatch(
            _messageHash,
            count() - 1,
            _destinationAndNonce(_destinationDomain, nonce),
            _tips,
            _message
        );
    }

    /**
     * @notice Suggest an attestation for the Notary to sign and submit.
     * @dev If no messages have been sent, null bytes returned for both
     * @return _nonce Current nonce
     * @return _root Current merkle root
     */
    function suggestAttestation() external view returns (uint32 _nonce, bytes32 _root) {
        uint256 length = historicalRoots.length;
        if (length != 0) {
            _nonce = uint32(length - 1);
            _root = historicalRoots[_nonce];
        }
    }

    // ============ Public Functions  ============

    /**
     * @notice Check if an Attestation is an Improper Attestation;
     * if so, slash the Notary and set the contract to FAILED state.
     *
     * An Improper Attestation is a (_nonce, _root) attestation that doesn't correspond with
     * the historical state of Origin contract. Either of those needs to be true:
     * - _nonce is higher than current nonce (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce
     * This would mean that message(s) that were not truly
     * dispatched on Origin were falsely included in the signed root.
     *
     * An Improper Attestation will only be accepted as valid by the Mirror
     * If an Improper Attestation is attempted on Origin,
     * the Notary will be slashed immediately.
     * If an Improper Attestation is submitted to the Mirror,
     * it should be relayed to the Origin contract using this function
     * in order to slash the Notary with an Improper Attestation.
     *
     * @dev Reverts (and doesn't slash notary) if signature is invalid
     * @param _attestation  Attestation data and signature
     * @return TRUE if attestation was an Improper Attestation (implying Notary was slashed)
     */
    function improperAttestation(bytes memory _attestation) public notFailed returns (bool) {
        // This will revert if signature is not valid
        (address _notary, bytes29 _view) = _checkNotaryAuth(_attestation);
        uint32 _nonce = _view.attestationNonce();
        bytes32 _root = _view.attestationRoot();
        // Check if nonce is valid, if not => attestation is fraud
        if (_nonce < historicalRoots.length) {
            if (_root == historicalRoots[_nonce]) {
                // Signed (nonce, root) attestation is valid
                return false;
            }
            // Signed root is not the same as the historical one => attestation is fraud
        }
        _fail(_notary);
        emit ImproperAttestation(_notary, _attestation);
        return true;
    }

    // ============ Internal Functions  ============

    /**
     * @notice Set the NotaryManager
     * @param _notaryManager Address of the NotaryManager
     */
    function _setNotaryManager(INotaryManager _notaryManager) internal {
        require(Address.isContract(address(_notaryManager)), "!contract notaryManager");
        notaryManager = INotaryManager(_notaryManager);
        emit NewNotaryManager(address(_notaryManager));
    }

    /**
     * @notice Slash the Notary and set contract state to FAILED
     * @dev Called when fraud is proven (Improper Attestation)
     */
    function _fail(address _notary) internal {
        /**
         * TODO: remove Failed state
         * @dev With the asynchronous updates Origin is never in the FAILED state.
         * It's rather some Destinations might end up with a corrupted merkle state
         * (upon receiving a fraud attestation). It's a Destination job to get this conflict fixed.
         * As long as there's more than one active Notary, new messages on Origin could be verified
         * by an honest Notary signing fresh valid attestations.
         * Meaning Origin should not be halted upon discovering a fraud attestation.
         */
        // set contract to FAILED
        state = States.Failed;
        // slash Notary
        notaryManager.slashNotary(payable(msg.sender));
        emit NotarySlashed(_notary, msg.sender);
    }

    /**
     * @notice Internal utility function that combines
     * `_destination` and `_nonce`.
     * @dev Both destination and nonce should be less than 2^32 - 1
     * @param _destination Domain of destination chain
     * @param _nonce Current nonce for given destination chain
     * @return Returns (`_destination` << 32) & `_nonce`
     */
    function _destinationAndNonce(uint32 _destination, uint32 _nonce)
        internal
        pure
        returns (uint64)
    {
        return (uint64(_destination) << 32) | _nonce;
    }

    /**
     * @notice  Returns "adjusted" sender address.
     * @dev     By default, "sender address" is msg.sender.
     *          However, if SystemMessenger sends a message, specifying SYSTEM_SENDER as the recipient,
     *          SYSTEM_SENDER is used as "sender address" on origin chain.
     *          Note that transaction will revert if anyone but SystemMessenger uses SYSTEM_SENDER as the recipient.
     */
    function _checkForSystemMessage(bytes32 _recipientAddress)
        internal
        view
        returns (bytes32 sender)
    {
        if (_recipientAddress != SystemMessage.SYSTEM_SENDER) {
            sender = TypeCasts.addressToBytes32(msg.sender);
            /**
             * @dev Note: SYSTEM_SENDER has highest 12 bytes set,
             *      whereas TypeCasts.addressToBytes32 sets only the lowest 20 bytes.
             *      Thus, in this branch: sender != SystemMessage.SYSTEM_SENDER
             */
        } else {
            // Check that SystemMessenger specified SYSTEM_SENDER as recipient, revert otherwise.
            _assertSystemMessenger();
            // Adjust "sender address" for correct processing on remote chain.
            sender = SystemMessage.SYSTEM_SENDER;
        }
    }
}
