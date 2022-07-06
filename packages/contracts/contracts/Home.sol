// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Version0 } from "./Version0.sol";
import { UpdaterStorage } from "./UpdaterStorage.sol";
import { AuthManager } from "./auth/AuthManager.sol";
import { HomeUpdate } from "./libs/HomeUpdate.sol";
import { QueueLib } from "./libs/Queue.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
import { MerkleTreeManager } from "./Merkle.sol";
import { IUpdaterManager } from "./interfaces/IUpdaterManager.sol";
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
contract Home is Version0, MerkleTreeManager, UpdaterStorage, AuthManager {
    // ============ Libraries ============

    using HomeUpdate for bytes29;
    using MerkleLib for MerkleLib.Tree;

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
    mapping(uint32 => uint32) public nonces;
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
     * @param message Raw bytes of message
     */
    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes message
    );

    /**
     * @notice Emitted when proof of an improper update is submitted,
     * which sets the contract to FAILED state
     * @param nonce Nonce of the improper update
     * @param root Root of the improper update
     * @param signature Signature on `nonce` and `root`
     */
    event ImproperUpdate(uint32 nonce, bytes32 root, bytes signature);

    /**
     * @notice Emitted when the Updater is slashed
     * (should be paired with ImproperUpdater or DoubleUpdate event)
     * @param updater The address of the updater
     * @param reporter The address of the entity that reported the updater misbehavior
     */
    event UpdaterSlashed(address indexed updater, address indexed reporter);

    /**
     * @notice Emitted when the UpdaterManager contract is changed
     * @param updaterManager The address of the new updaterManager
     */
    event NewUpdaterManager(address updaterManager);

    // ============ Constructor ============

    constructor(uint32 _localDomain) UpdaterStorage(_localDomain) {} // solhint-disable-line no-empty-blocks

    // ============ Initializer ============

    function initialize(IUpdaterManager _updaterManager) public initializer {
        // initialize queue, set Updater Manager, and initialize
        _setUpdaterManager(_updaterManager);
        __SynapseBase_initialize(updaterManager.updater());
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
        bytes memory _messageBody
    ) external notFailed {
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        // get the next nonce for the destination domain, then increment it
        uint32 _nonce = nonces[_destinationDomain];
        nonces[_destinationDomain] = _nonce + 1;
        // format the message into packed bytes
        bytes memory _message = Message.formatMessage(
            localDomain,
            bytes32(uint256(uint160(msg.sender))),
            _nonce,
            _destinationDomain,
            _recipientAddress,
            _optimisticSeconds,
            _messageBody
        );
        // insert the hashed message into the Merkle tree
        bytes32 _messageHash = keccak256(_message);
        // new root is added to the historical roots
        _insertHash(_messageHash);
        // Emit Dispatch event with message information
        // note: leafIndex is count() - 1 since new leaf has already been inserted
        emit Dispatch(
            _messageHash,
            count() - 1,
            _destinationAndNonce(_destinationDomain, _nonce),
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
        // TODO: rework for the new update scheme
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

    // ============ Public Functions  ============

    /**
     * @notice Check if an Update is an Improper Update;
     * if so, slash the Updater and set the contract to FAILED state.
     *
     * An Improper Update is a (_nonce, _root) update that doesn't correspond with
     * the historical state of Home contract. Either of those needs to be true:
     * - _nonce is higher than current nonce (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce
     * This would mean that message(s) that were not truly
     * dispatched on Home were falsely included in the signed root.
     *
     * An Improper Update will only be accepted as valid by the Replica
     * If an Improper Update is attempted on Home,
     * the Updater will be slashed immediately.
     * If an Improper Update is submitted to the Replica,
     * it should be relayed to the Home contract using this function
     * in order to slash the Updater with an Improper Update.
     *
     * @dev Reverts (and doesn't slash updater) if signature is invalid or
     * update not current
     * @param _updater      Updater who signed the update
     * @param _update       Update message
     * @param _signature    Updater signature on `_update`
     * @return TRUE if update was an Improper Update (implying Updater was slashed)
     */
    function improperUpdate(
        address _updater,
        bytes memory _update,
        bytes memory _signature
    ) public notFailed returns (bool) {
        // This will revert if signature is not valid
        bytes29 homeUpdate = _checkUpdaterAuth(_updater, _update, _signature);
        require(homeUpdate.homeDomain() == localDomain, "Wrong domain");
        uint32 _nonce = homeUpdate.nonce();
        bytes32 _root = homeUpdate.root();
        // Check if nonce is valid, if not => update is fraud
        if (_nonce < historicalRoots.length) {
            if (_root == historicalRoots[_nonce]) {
                // Signed (nonce, root) update is valid
                return false;
            }
            // Signed root is not the same as the historical one => update is fraud
        }
        _fail();
        emit ImproperUpdate(_nonce, _root, _signature);
        return true;
    }

    // ============ Internal Functions  ============

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
     * @dev Called when fraud is proven (Improper Update or Double Update)
     */
    function _fail() internal {
        // set contract to FAILED
        state = States.Failed;
        // slash Updater
        updaterManager.slashUpdater(payable(msg.sender));
        emit UpdaterSlashed(updater, msg.sender);
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

    function _isUpdater(uint32, address _updater) internal view override returns (bool) {
        return _updater == updater;
    }

    function _isWatchtower(address) internal pure override returns (bool) {
        return false;
    }
}
