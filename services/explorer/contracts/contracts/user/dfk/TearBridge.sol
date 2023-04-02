

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
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
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
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
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
}

//
abstract contract SynMessagingReceiver is ISynMessagingReceiver, Ownable {
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
}

//
// OpenZeppelin Contracts (last updated v4.5.0) (token/ERC20/IERC20.sol)
/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `from` to `to` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);

    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

//
/**
 * @dev Interface of Inventory Items.
 */
interface IInventoryItem is IERC20 {
    /**
     * @dev Burns tokens.
     */
    function burnFrom(address from, uint256 amount) external;

    function mint(address to, uint256 amount) external;
}

//
contract TearBridge is SynMessagingReceiver {
    address public immutable gaiaTears;
    uint256 public msgGasLimit;

    struct MessageFormat {
        address dstUser;
        uint256 dstTearAmount;
    }

    event GaiaSent(address indexed dstUser, uint256 arrivalChainId);
    event GaiaArrived(address indexed dstUser, uint256 arrivalChainId);

    constructor(address _messageBus, address _gaiaTear) {
        messageBus = _messageBus;
        gaiaTears = _gaiaTear;
    }

    function _createMessage(address _dstUserAddress, uint256 _dstTearAmount) internal pure returns (bytes memory) {
        // create the message here from the nested struct
        MessageFormat memory msgFormat = MessageFormat({dstUser: _dstUserAddress, dstTearAmount: _dstTearAmount});
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

    function sendTear(uint256 _tearsAmount, uint256 _dstChainId) external payable {
        uint256 tearsAmount = _tearsAmount;
        uint256 dstChainId = _dstChainId;
        // Tears now burnt, equivalent amount will be bridged to dstChainId
        IInventoryItem(gaiaTears).burnFrom(msg.sender, tearsAmount);

        bytes32 receiver = trustedRemoteLookup[dstChainId];
        bytes memory message = _createMessage(msg.sender, tearsAmount);
        bytes memory options = _createOptions();

        _send(receiver, dstChainId, message, options);
        emit GaiaSent(msg.sender, tearsAmount);
    }

    // Function called by executeMessage() - handleMessage will handle the gaia tear mint
    // executeMessage() handles permissioning checks
    function _handleMessage(
        bytes32 _srcAddress,
        uint256 _srcChainId,
        bytes memory _message,
        address _executor
    ) internal override {
        MessageFormat memory passedMsg = _decodeMessage(_message);
        address dstUser = passedMsg.dstUser;
        uint256 dstTearAmount = passedMsg.dstTearAmount;
        IInventoryItem(gaiaTears).mint(dstUser, dstTearAmount);
        emit GaiaArrived(dstUser, dstTearAmount);
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
}
