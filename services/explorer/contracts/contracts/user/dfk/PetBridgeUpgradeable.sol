pragma solidity 0.8.13;


//
interface ISynMessagingReceiver {
    // Maps chain ID to the bytes32 trusted addresses allowed to be source senders
    // mapping(uint256 => bytes32) internal trustedRemoteLookup;

    /**
     * @notice Called by MessageBus
     * @dev MUST be permissioned to trusted source apps via trustedRemote
     * @param _srcAddress The bytes32 address of the source app contract
     * @param _srcChainId The source chain ID where the transfer is originated from
     * @param _message Arbitrary message bytes originated from and encoded by the source app contract
     * @param _executor Address who called the MessageBus execution function
     */
    function executeMessage(
        bytes32 _srcAddress,
        uint256 _srcChainId,
        bytes calldata _message,
        address _executor
    ) external;
}

//
interface IMessageBus {
    /**
     * @notice Sends a message to a receiving contract address on another chain.
     * Sender must make sure that the message is unique and not a duplicate message.
     * @param _receiver The bytes32 address of the destination contract to be called
     * @param _dstChainId The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
     * @param _message The arbitrary payload to pass to the destination chain receiver
     * @param _options Versioned struct used to instruct relayer on how to proceed with gas limits
     */
    function sendMessage(
        bytes32 _receiver,
        uint256 _dstChainId,
        bytes calldata _message,
        bytes calldata _options
    ) external payable;

    /**
     * @notice Relayer executes messages through an authenticated method to the destination receiver based on the originating transaction on source chain
     * @param _srcChainId Originating chain ID - typically a standard EVM chain ID, but may refer to a Synapse-specific chain ID on nonEVM chains
     * @param _srcAddress Originating bytes address of the message sender on the srcChain
     * @param _dstAddress Destination address that the arbitrary message will be passed to
     * @param _gasLimit Gas limit to be passed alongside the message, depending on the fee paid on srcChain
     * @param _nonce Nonce from origin chain
     * @param _message Arbitrary message payload to pass to the destination chain receiver
     * @param _messageId MessageId for uniqueness of messages (alongisde nonce)
     */
    function executeMessage(
        uint256 _srcChainId,
        bytes calldata _srcAddress,
        address _dstAddress,
        uint256 _gasLimit,
        uint256 _nonce,
        bytes calldata _message,
        bytes32 _messageId
    ) external;

    /**
     * @notice Returns srcGasToken fee to charge in wei for the cross-chain message based on the gas limit
     * @param _options Versioned struct used to instruct relayer on how to proceed with gas limits. Contains data on gas limit to submit tx with.
     */
    function estimateFee(uint256 _dstChainId, bytes calldata _options) external returns (uint256);

    /**
     * @notice Withdraws message fee in the form of native gas token.
     * @param _account The address receiving the fee.
     */
    function withdrawFee(address _account) external;
}

//
// OpenZeppelin Contracts (last updated v4.5.0) (utils/Address.sol)
/**
 * @dev Collection of functions related to the address type
 */
library AddressUpgradeable {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://diligence.consensys.net/posts/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.5.11/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCall(target, data, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        require(isContract(target), "Address: call to non-contract");

        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        require(isContract(target), "Address: static call to non-contract");

        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verifies that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            // Look for revert reason and bubble it up if present
            if (returndata.length > 0) {
                // The easiest way to bubble the revert reason is using memory via assembly

                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}

//
// OpenZeppelin Contracts (last updated v4.5.0) (proxy/utils/Initializable.sol)
/**
 * @dev This is a base contract to aid in writing upgradeable contracts, or any kind of contract that will be deployed
 * behind a proxy. Since proxied contracts do not make use of a constructor, it's common to move constructor logic to an
 * external initializer function, usually called `initialize`. It then becomes necessary to protect this initializer
 * function so it can only be called once. The {initializer} modifier provided by this contract will have this effect.
 *
 * TIP: To avoid leaving the proxy in an uninitialized state, the initializer function should be called as early as
 * possible by providing the encoded function call as the `_data` argument to {ERC1967Proxy-constructor}.
 *
 * CAUTION: When used with inheritance, manual care must be taken to not invoke a parent initializer twice, or to ensure
 * that all initializers are idempotent. This is not verified automatically as constructors are by Solidity.
 *
 * [CAUTION]
 * ====
 * Avoid leaving a contract uninitialized.
 *
 * An uninitialized contract can be taken over by an attacker. This applies to both a proxy and its implementation
 * contract, which may impact the proxy. To initialize the implementation contract, you can either invoke the
 * initializer manually, or you can include a constructor to automatically mark it as initialized when it is deployed:
 *
 * [.hljs-theme-light.nopadding]
 * ```
 * /// @custom:oz-upgrades-unsafe-allow constructor
 * constructor() initializer {}
 * ```
 * ====
 */
abstract contract Initializable {
    /**
     * @dev Indicates that the contract has been initialized.
     */
    bool private _initialized;

    /**
     * @dev Indicates that the contract is in the process of being initialized.
     */
    bool private _initializing;

    /**
     * @dev Modifier to protect an initializer function from being invoked twice.
     */
    modifier initializer() {
        // If the contract is initializing we ignore whether _initialized is set in order to support multiple
        // inheritance patterns, but we only do this in the context of a constructor, because in other contexts the
        // contract may have been reentered.
        require(_initializing ? _isConstructor() : !_initialized, "Initializable: contract is already initialized");

        bool isTopLevelCall = !_initializing;
        if (isTopLevelCall) {
            _initializing = true;
            _initialized = true;
        }

        _;

        if (isTopLevelCall) {
            _initializing = false;
        }
    }

    /**
     * @dev Modifier to protect an initialization function so that it can only be invoked by functions with the
     * {initializer} modifier, directly or indirectly.
     */
    modifier onlyInitializing() {
        require(_initializing, "Initializable: contract is not initializing");
        _;
    }

    function _isConstructor() private view returns (bool) {
        return !AddressUpgradeable.isContract(address(this));
    }
}

//
// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)
/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract ContextUpgradeable is Initializable {
    function __Context_init() internal onlyInitializing {
    }

    function __Context_init_unchained() internal onlyInitializing {
    }
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }

    /**
     * This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}

//
// OpenZeppelin Contracts v4.4.1 (access/Ownable.sol)
/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract OwnableUpgradeable is Initializable, ContextUpgradeable {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    function __Ownable_init() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    function __Ownable_init_unchained() internal onlyInitializing {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }

    /**
     * This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}

//
abstract contract SynMessagingReceiverUpgradeable is ISynMessagingReceiver, OwnableUpgradeable {
    address public messageBus;

    // Maps chain ID to the bytes32 trusted addresses allowed to be source senders
    mapping(uint256 => bytes32) internal trustedRemoteLookup;

    event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress);

    /**
     * @notice Executes a message called by MessageBus (MessageBusReceiver)
     * @dev Must be called by MessageBug & sent from src chain by a trusted srcApp
     * @param _srcAddress The bytes32 address of the source app contract
     * @param _srcChainId The source chain ID where the transfer is originated from
     * @param _message Arbitrary message bytes originated from and encoded by the source app contract
     * @param _executor Address who called the MessageBus execution function
     */
    function executeMessage(
        bytes32 _srcAddress,
        uint256 _srcChainId,
        bytes calldata _message,
        address _executor
    ) external {
        // Must be called by the MessageBus/MessageBus for security
        require(msg.sender == messageBus, "caller is not message bus");
        // Must also be from a trusted source app
        require(_srcAddress == trustedRemoteLookup[_srcChainId], "Invalid source sending app");

        _handleMessage(_srcAddress, _srcChainId, _message, _executor);
    }

    // Logic here handling messsage contents
    function _handleMessage(
        bytes32 _srcAddress,
        uint256 _srcChainId,
        bytes memory _message,
        address _executor
    ) internal virtual;

    function _send(
        bytes32 _receiver,
        uint256 _dstChainId,
        bytes memory _message,
        bytes memory _options
    ) internal virtual {
        require(trustedRemoteLookup[_dstChainId] != bytes32(0), "Receiver not trusted remote");
        IMessageBus(messageBus).sendMessage{value: msg.value}(_receiver, _dstChainId, _message, _options);
    }

    //** Config Functions */
    function setMessageBus(address _messageBus) public onlyOwner {
        messageBus = _messageBus;
    }

    // allow owner to set trusted addresses allowed to be source senders
    function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) external onlyOwner {
        trustedRemoteLookup[_srcChainId] = _srcAddress;
        emit SetTrustedRemote(_srcChainId, _srcAddress);
    }

    //** View functions */
    function getTrustedRemote(uint256 _chainId) external view returns (bytes32 trustedRemote) {
        return trustedRemoteLookup[_chainId];
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}

struct PriceTier {
    uint8 jewelCost;
    uint16 goldCost;
    uint8 tearCost;
    uint32 incubationTime;
    uint16 shinyChance;
}

struct Pet {
    uint256 id;
    uint8 originId;
    string name;
    uint8 season;
    uint8 eggType; // 0 = blue, 1 = grey, 2 = green, 3 = yellow, 4 = gold
    uint8 rarity;
    uint8 element;
    uint8 bonusCount;
    uint8 profBonus;
    uint8 profBonusScalar;
    uint8 craftBonus;
    uint8 craftBonusScalar;
    uint8 combatBonus;
    uint8 combatBonusScalar;
    uint16 appearance;
    uint8 background;
    uint8 shiny;
    uint64 hungryAt;
    uint64 equippableAt;
    uint256 equippedTo;
}

struct PetOptions {
    uint8 originId;
    string name;
    uint8 season;
    uint8 eggType;
    uint8 rarity;
    uint8 element;
    uint8 bonusCount;
    uint8 profBonus;
    uint8 profBonusScalar;
    uint8 craftBonus;
    uint8 craftBonusScalar;
    uint8 combatBonus;
    uint8 combatBonusScalar;
    uint16 appearance;
    uint8 background;
    uint8 shiny;
}

struct UnhatchedEgg {
    uint256 id;
    uint256 petId;
    address owner;
    uint8 eggType;
    uint256 seedblock;
    uint256 finishTime;
    uint8 tier; // 0 = Small, 1 = Medium, 2 = Large
}

//
interface IPetCoreUpgradeable {
    function getUserPets(address _address) external view returns (Pet[] memory);

    function getPet(uint256 _id) external view returns (Pet memory);

    function hatchPet(PetOptions memory _petOptions, address owner) external returns (uint256);

    function updatePet(Pet memory _pet) external;

    function bridgeMint(uint256 _id, address _to) external;

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory _data
    ) external;

    function ownerOf(uint256 tokenId) external view returns (address);

    function approve(address to, uint256 tokenId) external;
}

//
/** @title Core app for handling cross chain messaging passing to bridge Pet NFTs
 */
contract PetBridgeUpgradeable is Initializable, SynMessagingReceiverUpgradeable {
    address public pets;
    uint256 public msgGasLimit;

    struct MessageFormat {
        Pet dstPet;
        address dstUser;
        uint256 dstPetId;
    }

    function initialize(address _messageBus, address _pets) external initializer {
        __Ownable_init_unchained();
        messageBus = _messageBus;
        pets = _pets;
    }

    event PetSent(uint256 indexed petId, uint256 arrivalChainId);
    event PetArrived(uint256 indexed petId, uint256 arrivalChainId);

    function _createMessage(
        uint256 _petId,
        address _dstUserAddress,
        Pet memory _petToBridge
    ) internal pure returns (bytes memory) {
        // create the message here from the nested struct
        MessageFormat memory msgFormat = MessageFormat({
            dstPetId: _petId,
            dstPet: _petToBridge,
            dstUser: _dstUserAddress
        });
        return abi.encode(msgFormat);
    }

    function _decodeMessage(bytes memory _message) internal pure returns (MessageFormat memory) {
        MessageFormat memory decodedMessage = abi.decode(_message, (MessageFormat));
        return decodedMessage;
    }

    function decodeMessage(bytes memory _message) external pure returns (MessageFormat memory) {
        return _decodeMessage(_message);
    }

    function _createOptions() internal view returns (bytes memory) {
        return abi.encodePacked(uint16(1), msgGasLimit);
    }

    /**
     * @notice User must have an existing pet minted to bridge it.
     * @dev This function enforces the caller to receive the Pet being bridged to the same address on another chain.
     * @dev Do NOT call this from other contracts, unless the contract is deployed on another chain to the same address,
     * @dev and can receive ERC721s.
     * @param _petId specifics which pet msg.sender already holds and will transfer to the bridge contract
     * @param _dstChainId The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
     */
    function sendPet(uint256 _petId, uint256 _dstChainId) external payable {
        uint256 petId = _petId;
        uint256 dstChainId = _dstChainId;
        Pet memory petToBridge = IPetCoreUpgradeable(pets).getPet(petId);
        // revert if the pet is equipped
        require(petToBridge.equippedTo == 0, "pet is equipped");

        bytes32 receiver = trustedRemoteLookup[dstChainId];
        // _createMessage(petId, dstUserAddress, Pet);
        // Only bridgeable directly to the caller of this contract
        // @dev do not call this function from other contracts
        bytes memory msgToPass = _createMessage(petId, msg.sender, petToBridge);
        // Create _options
        bytes memory options = _createOptions();

        IPetCoreUpgradeable(pets).transferFrom(msg.sender, address(this), petId);
        require(IPetCoreUpgradeable(pets).ownerOf(petId) == address(this), "Failed to lock Pet");
        // Pet now locked, message can be safely emitted

        _send(receiver, dstChainId, msgToPass, options);
        emit PetSent(petId, dstChainId);
    }

    // Function called by executeMessage() - handleMessage will handle the pet bridge mint
    // executeMessage() handles permissioning checks
    function _handleMessage(
        bytes32,
        uint256,
        bytes memory _message,
        address
    ) internal override {
        // Decode _message, depending on exactly how the originating message is structured
        /**
            Message data:
                Pet memory petToBridge = IPetCoreUpgradeable(pets).getPet(_petId);
                address dstUserAddress = msg.sender;
                uint256 dstPetId = _petId;
             */
        MessageFormat memory passedMsg = _decodeMessage(_message);

        Pet memory dstPet = passedMsg.dstPet;
        address dstUser = passedMsg.dstUser;
        uint256 dstPetId = passedMsg.dstPetId;

        // will revert if non-existent Pet
        try IPetCoreUpgradeable(pets).ownerOf(dstPetId) returns (address petOwner) {
            /**
                If petId does exist (which means it should be locked on this contract), as it was bridged before.
                Transfer it to message.dstUserAddress
                */

            if (petOwner == address(this)) {
                IPetCoreUpgradeable(pets).safeTransferFrom(address(this), dstUser, dstPetId);
            }
        } catch {
            /**
                If pet ID doesn't exist:
                Mint a pet to msg.dstUserAddress
                */
            IPetCoreUpgradeable(pets).bridgeMint(dstPetId, dstUser);
        }

        // update the pet attributes based on the attributes in the message (Assumes the message has more recent attributes)
        IPetCoreUpgradeable(pets).updatePet(dstPet);
        // Tx completed, emit success
        emit PetArrived(dstPetId, block.chainid);
    }

    function _send(
        bytes32 _receiver,
        uint256 _dstChainId,
        bytes memory _message,
        bytes memory _options
    ) internal override {
        bytes32 trustedRemote = trustedRemoteLookup[_dstChainId];
        require(trustedRemote != bytes32(0), "No remote app set for dst chain");
        require(trustedRemote == _receiver, "Receiver is not in trusted remote apps");
        IMessageBus(messageBus).sendMessage{value: msg.value}(_receiver, _dstChainId, _message, _options);
    }

    function setMsgGasLimit(uint256 _msgGasLimit) external onlyOwner {
        msgGasLimit = _msgGasLimit;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
