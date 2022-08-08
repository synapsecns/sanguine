// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Version0 } from "./Version0.sol";
import { UpdaterStorage } from "./UpdaterStorage.sol";
import { ReportHub } from "./auth/ReportHub.sol";
import { Attestation } from "./libs/Attestation.sol";
import { Report } from "./libs/Report.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";
import { QueueLib } from "./libs/Queue.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Header } from "./libs/Header.sol";
import { Message } from "./libs/Message.sol";
import { Tips } from "./libs/Tips.sol";
import { SystemMessage } from "./system/SystemMessage.sol";
import { MerkleTreeManager } from "./Merkle.sol";
import { IUpdaterManager } from "./interfaces/IUpdaterManager.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
// ============ External Imports ============
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Home
 * @author Illusory Systems Inc.
 * @notice Accepts messages to be dispatched to remote chains,
 * constructs a Merkle tree of the messages,
 * and accepts signatures from a bonded Updater
 * which notarize the Merkle tree roots.
 * Accepts submissions of fraudulent signatures
 * by the Updater and slashes the Updater in this case.
 */
contract Home is Version0, MerkleTreeManager, UpdaterStorage, ReportHub {
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
    //   contract will no longer accept updates or new messages
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
    // contract responsible for Updater bonding, slashing and rotation
    IUpdaterManager public updaterManager;
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
     * @notice Emitted when proof of an invalid attestation is submitted,
     * which sets the contract to FAILED state
     * @param updater       Updater who signed invalid attestation
     * @param attestation   Attestation data and signature
     */
    event InvalidAttestation(address updater, bytes attestation);

    /**
     * @notice Emitted when proof of an invalid fraud report is submitted
     * @param watchtower    Watchtower who signed invalid fraud report
     * @param report        Report data and signature
     */
    event InvalidReport(address watchtower, bytes report);

    /**
     * @notice Emitted when the Updater is slashed
     * (should be paired with InvalidAttestation event)
     * @param updater       The address of the updater
     * @param reporter      The address of the entity that reported the updater misbehavior
     * @param watchtower    The address of watchtower that signed the fraud report
     */
    event UpdaterSlashed(
        address indexed updater,
        address indexed reporter,
        address indexed watchtower
    );

    /**
     * @notice Emitted when the UpdaterManager contract is changed
     * @param updaterManager The address of the new updaterManager
     */
    event NewUpdaterManager(address updaterManager);

    // ============ Constructor ============

    constructor(uint32 _localDomain) UpdaterStorage(_localDomain) {} // solhint-disable-line no-empty-blocks

    // ============ Initializer ============

    function initialize(IUpdaterManager _updaterManager, address _watchtower) public initializer {
        // initialize queue, set Updater Manager, and initialize
        _setUpdaterManager(_updaterManager);
        __SynapseBase_initialize(updaterManager.updater(), _watchtower);
        state = States.Active;
    }

    // ============ Modifiers ============

    /**
     * @notice Ensures that function is called by the UpdaterManager contract
     */
    modifier onlyUpdaterManager() {
        require(msg.sender == address(updaterManager), "!updaterManager");
        _;
    }

    /**
     * @notice Ensures that contract state != FAILED when the function is called
     */
    modifier notFailed() {
        require(state != States.Failed, "failed state");
        _;
    }

    // ============ External: Updater & UpdaterManager Configuration  ============

    /**
     * @notice Set a new Updater
     * @dev To be set when rotating Updater after Fraud
     * @param _updater the new Updater
     */
    function setUpdater(address _updater) external onlyUpdaterManager {
        _setUpdater(_updater);
        // set the Home state to Active
        // now that Updater has been rotated
        state = States.Active;
    }

    /**
     * @notice Set a new UpdaterManager contract
     * @dev Home(s) will initially be initialized using a trusted UpdaterManager contract;
     * we will progressively decentralize by swapping the trusted contract with a new implementation
     * that implements Updater bonding & slashing, and rules for Updater selection & rotation
     * @param _updaterManager the new UpdaterManager contract
     */
    function setUpdaterManager(address _updaterManager) external onlyOwner {
        _setUpdaterManager(IUpdaterManager(_updaterManager));
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
     * @notice Suggest an update for the Updater to sign and submit.
     * @dev If no messages have been sent, null bytes returned for both
     * @return _nonce Current nonce
     * @return _root Current merkle root
     */
    function suggestUpdate() external view returns (uint32 _nonce, bytes32 _root) {
        uint256 length = historicalRoots.length;
        if (length != 0) {
            _nonce = uint32(length - 1);
            _root = historicalRoots[_nonce];
        }
    }

    /**
     * @notice Hash of Home domain concatenated with "SYN"
     */
    function homeDomainHash() external view returns (bytes32) {
        return _domainHash(localDomain);
    }

    // ============ Internal Functions  ============

    /**
     * @notice Check if a reported Attestation is an Invalid Attestation;
     * if so, slash the Notary and set the contract to FAILED state.
     *
     * An Invalid Attestation is a (_nonce, _root) attestation that doesn't correspond with
     * the historical state of Home contract. Either of those needs to be true:
     * - _nonce is higher than current nonce (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce
     * This would mean that message(s) that were not truly
     * dispatched on Home were falsely included in the signed root.
     *
     * An Invalid Attestation will only be accepted as valid by the Replica
     * If an Invalid Attestation is attempted on Home, the Notary will be slashed immediately.
     * If an Invalid Attestation is submitted to the Replica, a Guard should generate a Report.
     * This Report should be submitted to the Home contract using this function
     * in order to slash the Notary with an Invalid Attestation.
     *
     * @dev Both Notary and Guard signatures
     * have been checked at this point (see ReportHub.sol).
     *
     * @param _guard            Guard address
     * @param _notary           Notary address
     * @param _attestationView  Memory view over reported Attestation
     * @param _report           Payload with Report data and signature
     * @return TRUE if attestation was an Invalid Attestation (implying Notary was slashed)
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes memory _report
    ) internal override notFailed returns (bool) {
        // Get merkle state from the attestation
        uint32 _nonce = _attestationView.attestationNonce();
        bytes32 _root = _attestationView.attestationRoot();
        // Check if `_nonce` exists, if not => attestation is fraud
        if (_nonce < historicalRoots.length) {
            if (_root == historicalRoots[_nonce]) {
                // (nonce, root) corresponds with the historical merkle state of Home.
                // Means Notary attestation is valid, while Guard report is invalid.
                // TODO: slash Guard for signing an invalid fraud report
                emit InvalidReport(_guard, _report);
                return false;
            }
            // `_root` doesn't match historical root for `_nonce` => attestation is fraud
        }
        _fail(_guard);
        emit InvalidAttestation(_notary, _attestationView.clone());
        return true;
    }

    /**
     * @notice Set the UpdaterManager
     * @param _updaterManager Address of the UpdaterManager
     */
    function _setUpdaterManager(IUpdaterManager _updaterManager) internal {
        require(Address.isContract(address(_updaterManager)), "!contract updaterManager");
        updaterManager = IUpdaterManager(_updaterManager);
        emit NewUpdaterManager(address(_updaterManager));
    }

    /**
     * @notice Slash the Updater and set contract state to FAILED
     * @dev Called when fraud is proven (Invalid Attestation)
     */
    function _fail(address _watchtower) internal {
        // set contract to FAILED
        state = States.Failed;
        // slash Updater
        updaterManager.slashUpdater(payable(msg.sender));
        emit UpdaterSlashed(updater, msg.sender, _watchtower);
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

    function _isUpdater(uint32 _homeDomain, address _updater)
        internal
        view
        override
        returns (bool)
    {
        require(_homeDomain == localDomain, "Wrong domain");
        return _updater == updater;
    }

    function _isWatchtower(address _watchtower) internal view override returns (bool) {
        return _watchtower == watchtower;
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
