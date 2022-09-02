// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Version0 } from "./Version0.sol";
import { DomainNotaryRegistry } from "./registry/DomainNotaryRegistry.sol";
import { GuardRegistry } from "./registry/GuardRegistry.sol";
import { ReportHub } from "./hubs/ReportHub.sol";
import { Attestation } from "./libs/Attestation.sol";
import { Report } from "./libs/Report.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Header } from "./libs/Header.sol";
import { Message } from "./libs/Message.sol";
import { Tips } from "./libs/Tips.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";
import { SystemMessage } from "./libs/SystemMessage.sol";
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
    ReportHub,
    DomainNotaryRegistry,
    GuardRegistry
{
    // ============ Libraries ============

    using Attestation for bytes29;
    using Report for bytes29;
    using TypedMemView for bytes29;
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
     * @notice Emitted when a correct report on a fraud attestation is submitted.
     * @param guard     Guard who signed the fraud report
     * @param report    Report data and signature
     */
    event CorrectFraudReport(address indexed guard, bytes report);

    /**
     * @notice Emitted when proof of an incorrect report is submitted.
     * @param guard     Guard who signed the incorrect report
     * @param report    Report data and signature
     */
    event IncorrectReport(address indexed guard, bytes report);

    /**
     * @notice Emitted when proof of an fraud attestation is submitted,
     * which sets the contract to FAILED state
     * @param notary        Notary who signed fraud attestation
     * @param attestation   Attestation data and signature
     */
    event FraudAttestation(address indexed notary, bytes attestation);

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
         *      1a. onlyNotaryManager -> onlyBondingManager (or w/e the name would be)
         *      2. There is supposed to be more than one active Notary
         *      2a. setNotary() -> addNotary()
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
    ) external payable notFailed {
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        require(_tips.castToTips().totalTips() == msg.value, "!tips");
        // get the next nonce for the destination domain, then increment it
        nonce = nonce + 1;
        bytes32 _sender = _checkForSystemMessage(_recipientAddress);
        // format the message into packed bytes
        bytes memory _header = Header.formatHeader(
            localDomain,
            _sender,
            nonce,
            _destination,
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
            _destinationAndNonce(_destination, nonce),
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

    // ============ Internal Functions  ============

    // TODO: handleAttestation (needs AttestationHub) for reporting
    // a fraud attestation without a Guard signature

    /**
     * @notice Checks if a submitted Report is a correct Report. Reported Attestation
     * can be either valid or fraud. Report flag can also be either Valid or Fraud.
     * Report is correct if its flag matches the Attestation validity.
     * 1. Attestation: valid, Flag: Fraud.
     *      Report is deemed incorrect, Guard is slashed (if they haven't been already).
     * 2. Attestation: valid, Flag: Valid.
     *      Report is deemed correct, no action is done.
     * 3. Attestation: Fraud, Flag: Fraud.
     *      Report is deemed correct, Notary is slashed (if they haven't been already).
     * 4. Attestation: Fraud, Flag: Valid.
     *      Report is deemed incorrect, Guard is slashed (if they haven't been already).
     *      Notary is slashed (if they haven't been already), but Guard doesn't receive
     *      any rewards (as their report indicated that the attestation was valid).
     *
     * A Fraud Attestation is a (_nonce, _root) attestation that doesn't correspond with
     * the historical state of Origin contract. Either of those needs to be true:
     * - _nonce is higher than current nonce (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce
     * This would mean that message(s) that were not truly
     * dispatched on Origin were falsely included in the signed root.
     *
     * A Fraud Attestation will only be accepted as valid by the Mirror.
     * If a Fraud Attestation is submitted to the Mirror, a Guard should
     * submit a Fraud Report using Origin.submitReport()
     * in order to slash the Notary with a Fraud Attestation.
     *
     * @dev Both Notary and Guard signatures
     * have been checked at this point (see ReportHub.sol).
     *
     * @param _guard            Guard address
     * @param _notary           Notary address
     * @param _attestationView  Memory view over reported Attestation
     * @param _reportView       Memory view over Report
     * @param _report           Payload with Report data and signature
     * @return TRUE if Report was correct (implying Guard was not slashed)
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override notFailed returns (bool) {
        uint32 _nonce = _attestationView.attestedNonce();
        bytes32 _root = _attestationView.attestedRoot();
        if (_isValidAttestation(_nonce, _root)) {
            // Attestation: Valid
            if (_reportView.reportedFraud()) {
                // Flag: Fraud
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                return false;
            } else {
                // Flag: Valid
                // Report is correct, no action needed
                return true;
            }
        } else {
            // Attestation: Fraud
            if (_reportView.reportedFraud()) {
                // Flag: Fraud
                // Report is correct, slash the Notary
                emit CorrectFraudReport(_guard, _report);
                emit FraudAttestation(_notary, _attestationView.clone());
                _fail(_notary, _guard);
                return true;
            } else {
                // Flag: Valid
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                emit FraudAttestation(_notary, _attestationView.clone());
                // Guard doesn't receive anything due to Valid flag on the Report
                _fail(_notary, address(0));
                return false;
            }
        }
    }

    function _isValidAttestation(uint32 _nonce, bytes32 _root) internal view returns (bool) {
        // Check if nonce is valid, if not => attestation is fraud
        // Check if root the same as the historical one, if not => attestation is fraud
        return (_nonce < historicalRoots.length && _root == historicalRoots[_nonce]);
    }

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
     * @dev Called when fraud is proven (Fraud Attestation)
     */
    function _fail(address _notary, address _guard) internal {
        /**
         * TODO: remove Failed state
         * @dev With the asynchronous attestations Origin is never in the FAILED state.
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
        emit NotarySlashed(_notary, _guard, msg.sender);
    }

    /**
     * @notice Slash the Guard.
     * @dev Called when guard misbehavior is proven (Incorrect Report).
     */
    function _slashGuard(address _guard) internal {
        // TODO: implement actual slashing & slash only once
        emit GuardSlashed(_guard, msg.sender);
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
     *          However, if SystemRouter sends a message, specifying SYSTEM_ROUTER as the recipient,
     *          SYSTEM_ROUTER is used as "sender address" on origin chain.
     *          Note that transaction will revert if anyone but SystemRouter uses SYSTEM_ROUTER as the recipient.
     */
    function _checkForSystemMessage(bytes32 _recipientAddress)
        internal
        view
        returns (bytes32 sender)
    {
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
