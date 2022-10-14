// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { LocalDomainContext } from "./context/LocalDomainContext.sol";
import { Version0 } from "./Version0.sol";
import { OriginHub } from "./hubs/OriginHub.sol";
import { Header } from "./libs/Header.sol";
import { Message } from "./libs/Message.sol";
import { Tips } from "./libs/Tips.sol";
import { SystemMessage } from "./libs/SystemMessage.sol";
import { SystemContract } from "./system/SystemContract.sol";
import { INotaryManager } from "./interfaces/INotaryManager.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
// ============ External Imports ============
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Origin
 * @author Illusory Systems Inc.
 * @notice Accepts messages to be dispatched to remote chains and
 * constructs a Merkle tree of the messages.
 * Notaries are signing the attestations of the Merkle tree's root state (aka merkle state),
 * which are broadcasted to Destination, where the merkle root is used for proving that
 * the message has been indeed dispatched on Origin.
 * Origin accepts submissions of fraudulent signatures by the Notary,
 * directly or in the form of a Guard's Fraud report on such an attestation,
 * and slashes the Notary in this case.
 * Origin accepts submissions of fraudulent signatures by the Guard in the form
 * of a Guard's report with said signature and slashes Guard in that case.
 */
contract Origin is Version0, SystemContract, LocalDomainContext, OriginHub {
    using Tips for bytes;
    using Tips for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Maximum bytes per message = 2 KiB
    // (somewhat arbitrarily set to begin)
    uint256 public constant MAX_MESSAGE_BODY_BYTES = 2 * 2**10;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // contract responsible for Notary bonding, slashing and rotation
    // TODO: use "bonding manager" instead when implemented
    INotaryManager public notaryManager;

    // gap for upgrade safety
    uint256[49] private __GAP; //solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Emitted when a new message is dispatched
     * @param messageHash Hash of message; the leaf inserted to the Merkle tree
     *        for the message
     * @param nonce Nonce of sent message (starts from 1)
     * @param destination Destination domain
     * @param tips Tips paid for the remote off-chain agents
     * @param message Raw bytes of message
     */
    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    /**
     * @notice Emitted when the Guard is slashed
     * (should be paired with IncorrectReport event)
     * @param guard     The address of the guard that signed the incorrect report
     * @param reporter  The address of the entity that reported the guard misbehavior
     */
    event GuardSlashed(address indexed guard, address indexed reporter);

    /**
     * @notice Emitted when the Notary is slashed
     * (should be paired with FraudAttestation event)
     * @param notary    The address of the notary
     * @param guard     The address of the guard that signed the fraud report
     * @param reporter  The address of the entity that reported the notary misbehavior
     */
    event NotarySlashed(address indexed notary, address indexed guard, address indexed reporter);

    /**
     * @notice Emitted when the NotaryManager contract is changed
     * @param notaryManager The address of the new notaryManager
     */
    event NewNotaryManager(address notaryManager);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Ensures that function is called by the NotaryManager contract
     */
    modifier onlyNotaryManager() {
        require(msg.sender == address(notaryManager), "!notaryManager");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) LocalDomainContext(_domain) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize(INotaryManager _notaryManager) external initializer {
        __SystemContract_initialize();
        _initializeHistoricalRoots();
        _setNotaryManager(_notaryManager);
        _addNotary(notaryManager.notary());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    EXTERNAL FUNCTIONS: RESTRICTED                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
         *      1a. onlyNotaryManager -> onlyBondingManager (or w/e the name would be)
         *      2. There is supposed to be more than one active Notary
         *      2a. setNotary() -> addNotary()
         */
        _addNotary(_notary);
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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Dispatch the message to the destination domain & recipient
     * @dev Format the message, insert its hash into Merkle tree,
     * enqueue the new Merkle root, and emit `Dispatch` event with message information.
     * @param _destination      Domain of destination chain
     * @param _recipientAddress Address of recipient on destination chain as bytes32
     * @param _messageBody      Raw bytes content of message
     */
    function dispatch(
        uint32 _destination,
        bytes32 _recipientAddress,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable haveActiveNotary returns (uint32 messageNonce, bytes32 messageHash) {
        // TODO: add unit tests covering return values
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        require(_tips.castToTips().totalTips() == msg.value, "!tips");
        // Latest nonce (i.e. "last message" nonce) is current amount of leaves in the tree.
        // Message nonce is the amount of leaves after the new leaf insertion
        messageNonce = nonce() + 1;
        // format the message into packed bytes
        bytes memory message = Message.formatMessage({
            _origin: _localDomain(),
            _sender: _getSender(_recipientAddress),
            _nonce: messageNonce,
            _destination: _destination,
            _recipient: _recipientAddress,
            _optimisticSeconds: _optimisticSeconds,
            _tips: _tips,
            _messageBody: _messageBody
        });
        messageHash = keccak256(message);
        // insert the hashed message into the Merkle tree
        _insertMessage(messageNonce, messageHash);
        // Emit Dispatch event with message information
        // note: leaf index in the tree is messageNonce - 1, meaning we don't need to emit that
        emit Dispatch(messageHash, messageNonce, _destination, _tips, message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
     * @notice Slash the Notary.
     * @dev Called when fraud is proven (Fraud Attestation).
     * @param _notary   Notary to slash
     * @param _guard    Guard who reported fraudulent Notary [address(0) if not a Guard report]
     */
    function _slashNotary(address _notary, address _guard) internal override {
        // _notary is always an active Notary at this point
        _removeNotary(_notary);
        notaryManager.slashNotary(payable(msg.sender));
        emit NotarySlashed(_notary, _guard, msg.sender);
    }

    /**
     * @notice Slash the Guard.
     * @dev Called when guard misbehavior is proven (Incorrect Report).
     * @param _guard    Guard to slash
     */
    function _slashGuard(address _guard) internal override {
        // _guard is always an active Guard at this point
        _removeGuard(_guard);
        emit GuardSlashed(_guard, msg.sender);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns "adjusted" sender address.
     * @dev By default, "sender address" is msg.sender.
     * However, if SystemRouter sends a message, specifying SYSTEM_ROUTER as the recipient,
     * SYSTEM_ROUTER is used as "sender address" on origin chain.
     * Note: tx will revert if anyone but SystemRouter uses SYSTEM_ROUTER as the recipient.
     */
    function _getSender(bytes32 _recipientAddress) internal view returns (bytes32 sender) {
        if (_recipientAddress != SystemMessage.SYSTEM_ROUTER) {
            sender = TypeCasts.addressToBytes32(msg.sender);
            /**
             * @dev Note: SYSTEM_ROUTER has highest 12 bytes set,
             *      whereas TypeCasts.addressToBytes32 sets only the lowest 20 bytes.
             *      Thus, in this branch: sender != SystemMessage.SYSTEM_ROUTER
             */
        } else {
            // Check that SystemRouter specified SYSTEM_ROUTER as recipient, revert otherwise.
            _assertSystemRouter();
            // Adjust "sender address" for correct processing on remote chain.
            sender = SystemMessage.SYSTEM_ROUTER;
        }
    }
}
