// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts (last updated v4.5.0) (proxy/utils/Initializable.sol)

pragma solidity ^0.8.0;


// OpenZeppelin Contracts (last updated v4.5.0) (utils/Address.sol)




library AddressUpgradeable {

    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }


    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }


    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCall(target, data, "Address: low-level call failed");
    }

    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }


    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }


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


    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }


    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        require(isContract(target), "Address: static call to non-contract");

        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }


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

abstract contract Initializable {

    bool private _initialized;


    bool private _initializing;


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


    modifier onlyInitializing() {
        require(_initializing, "Initializable: contract is not initializing");
        _;
    }

    function _isConstructor() private view returns (bool) {
        return !AddressUpgradeable.isContract(address(this));
    }
}








// OpenZeppelin Contracts v4.4.1 (access/Ownable.sol)




// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)





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


    uint256[50] private __gap;
}




abstract contract OwnableUpgradeable is Initializable, ContextUpgradeable {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    function __Ownable_init() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    function __Ownable_init_unchained() internal onlyInitializing {
        _transferOwnership(_msgSender());
    }


    function owner() public view virtual returns (address) {
        return _owner;
    }


    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }


    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }


    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }


    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }


    uint256[49] private __gap;
}


// OpenZeppelin Contracts v4.4.1 (security/Pausable.sol)






abstract contract PausableUpgradeable is Initializable, ContextUpgradeable {

    event Paused(address account);

    event Unpaused(address account);

    bool private _paused;


    function __Pausable_init() internal onlyInitializing {
        __Pausable_init_unchained();
    }

    function __Pausable_init_unchained() internal onlyInitializing {
        _paused = false;
    }


    function paused() public view virtual returns (bool) {
        return _paused;
    }


    modifier whenNotPaused() {
        require(!paused(), "Pausable: paused");
        _;
    }


    modifier whenPaused() {
        require(paused(), "Pausable: not paused");
        _;
    }


    function _pause() internal virtual whenNotPaused {
        _paused = true;
        emit Paused(_msgSender());
    }


    function _unpause() internal virtual whenPaused {
        _paused = false;
        emit Unpaused(_msgSender());
    }

    uint256[49] private __gap;
}






interface IGasFeePricing {

    function setCostPerChain(
        uint256 _dstChainId,
        uint256 _gasUnitPrice,
        uint256 _gasTokenPriceRatio
    ) external;


    function estimateGasFee(uint256 _dstChainId, bytes calldata _options) external returns (uint256);
}







abstract contract ContextChainIdUpgradeable is Initializable {
    function __ContextChainId_init() internal onlyInitializing {}

    function __ContextChainId_init_unchained() internal onlyInitializing {}

    function _chainId() internal view virtual returns (uint256) {
        return block.chainid;
    }


    uint256[50] private __gap;
}


contract MessageBusSenderUpgradeable is OwnableUpgradeable, PausableUpgradeable, ContextChainIdUpgradeable {
    address public gasFeePricing;
    uint64 public nonce;
    uint256 public fees;

    function __MessageBusSender_init(address _gasFeePricing) internal onlyInitializing {
        __Ownable_init_unchained();
        __Pausable_init_unchained();
        __MessageBusSender_init_unchained(_gasFeePricing);
    }

    function __MessageBusSender_init_unchained(address _gasFeePricing) internal onlyInitializing {
        gasFeePricing = _gasFeePricing;
    }

    event MessageSent(
        address indexed sender,
        uint256 srcChainID,
        bytes32 receiver,
        uint256 indexed dstChainId,
        bytes message,
        uint64 nonce,
        bytes options,
        uint256 fee,
        bytes32 indexed messageId
    );

    function computeMessageId(
        address _srcAddress,
        uint256 _srcChainId,
        bytes32 _dstAddress,
        uint256 _dstChainId,
        uint256 _srcNonce,
        bytes calldata _message
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(_srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message));
    }

    function estimateFee(uint256 _dstChainId, bytes calldata _options) public returns (uint256) {
        uint256 fee = IGasFeePricing(gasFeePricing).estimateGasFee(_dstChainId, _options);
        require(fee != 0, "Fee not set");
        return fee;
    }

    function sendMessage(
        bytes32 _receiver,
        uint256 _dstChainId,
        bytes calldata _message,
        bytes calldata _options
    ) external payable whenNotPaused {
        // use tx.origin for gas refund by default, so that older contracts,
        // interacting with MessageBus that don't have a fallback/receive
        // (i.e. not able to receive gas), will continue to work
        _sendMessage(_receiver, _dstChainId, _message, _options, payable(tx.origin));
    }

    function sendMessage(
        bytes32 _receiver,
        uint256 _dstChainId,
        bytes calldata _message,
        bytes calldata _options,
        address payable _refundAddress
    ) external payable {
        _sendMessage(_receiver, _dstChainId, _message, _options, _refundAddress);
    }

    function _sendMessage(
        bytes32 _receiver,
        uint256 _dstChainId,
        bytes calldata _message,
        bytes calldata _options,
        address payable _refundAddress
    ) internal {
        uint256 srcChainId = _chainId();
        require(_dstChainId != srcChainId, "Invalid chainId");
        uint256 fee = estimateFee(_dstChainId, _options);
        require(msg.value >= fee, "Insufficient gas fee");
        bytes32 msgId = computeMessageId(msg.sender, srcChainId, _receiver, _dstChainId, nonce, _message);
        emit MessageSent(msg.sender, srcChainId, _receiver, _dstChainId, _message, nonce, _options, fee, msgId);
        fees += fee;
        ++nonce;
        // refund gas fees in case of overpayment
        if (msg.value > fee) {
            _refundAddress.transfer(msg.value - fee);
        }
    }


    function withdrawGasFees(address payable to) external onlyOwner {
        uint256 withdrawAmount = fees;
        // Reset fees to 0
        to.transfer(withdrawAmount);
        delete fees;
    }


    function rescueGas(address payable to) external onlyOwner {
        uint256 withdrawAmount = address(this).balance - fees;
        to.transfer(withdrawAmount);
    }

    function updateGasFeePricing(address _gasFeePricing) external onlyOwner {
        require(_gasFeePricing != address(0), "Cannot set to 0");
        gasFeePricing = _gasFeePricing;
    }


    uint256[47] private __gap;
}












interface IAuthVerifier {

    function msgAuth(bytes calldata _authData) external view returns (bool authenticated);


    function setNodeGroup(address _nodegroup) external;
}





interface ISynMessagingReceiver {
    // Maps chain ID to the bytes32 trusted addresses allowed to be source senders
    // mapping(uint256 => bytes32) internal trustedRemoteLookup;


    function executeMessage(
        bytes32 _srcAddress,
        uint256 _srcChainId,
        bytes calldata _message,
        address _executor
    ) external;
}


contract MessageBusReceiverUpgradeable is OwnableUpgradeable, PausableUpgradeable {
    enum TxStatus {
        Null,
        Success,
        Fail
    }

    // TODO: Rename to follow one standard convention -> Send -> Receive?
    event Executed(
        bytes32 indexed messageId,
        TxStatus status,
        address indexed _dstAddress,
        uint64 srcChainId,
        uint64 srcNonce
    );
    event CallReverted(string reason);

    address public authVerifier;

    // Store all successfully executed messages
    mapping(bytes32 => TxStatus) internal executedMessages;

    function __MessageBusReceiver_init(address _authVerifier) internal {
        __Ownable_init_unchained();
        __Pausable_init_unchained();
        __MessageBusReceiver_init_unchained(_authVerifier);
    }

    function __MessageBusReceiver_init_unchained(address _authVerifier) internal {
        authVerifier = _authVerifier;
    }

    function getExecutedMessage(bytes32 _messageId) external view returns (TxStatus) {
        return executedMessages[_messageId];
    }


    function executeMessage(
        uint256 _srcChainId,
        bytes32 _srcAddress,
        address _dstAddress,
        uint256 _gasLimit,
        uint256 _nonce,
        bytes calldata _message,
        bytes32 _messageId
    ) external whenNotPaused {
        // In order to guarantee that an individual message is only executed once, a messageId is passed
        // enforce that this message ID hasn't already been tried ever
        require(executedMessages[_messageId] == TxStatus.Null, "Message already executed");
        // Authenticate executeMessage, will revert if not authenticated
        IAuthVerifier(authVerifier).msgAuth(abi.encode(msg.sender));

        TxStatus status;
        try
            ISynMessagingReceiver(_dstAddress).executeMessage{gas: _gasLimit}(
                _srcAddress,
                _srcChainId,
                _message,
                msg.sender
            )
        {
            // Assuming success state if no revert
            status = TxStatus.Success;
        } catch (bytes memory reason) {
            // call hard reverted & failed
            emit CallReverted(_getRevertMsg(reason));
            status = TxStatus.Fail;
        }

        executedMessages[_messageId] = status;
        emit Executed(_messageId, status, _dstAddress, uint64(_srcChainId), uint64(_nonce));
    }


    // https://ethereum.stackexchange.com/a/83577
    // https://github.com/Uniswap/v3-periphery/blob/v1.0.0/contracts/base/Multicall.sol
    function _getRevertMsg(bytes memory _returnData) internal pure returns (string memory) {
        // If the _res length is less than 68, then the transaction failed silently (without a revert message)
        if (_returnData.length < 68) return "Transaction reverted silently";
        // solhint-disable-next-line
        assembly {
            // Slice the sighash.
            _returnData := add(_returnData, 0x04)
        }
        return abi.decode(_returnData, (string)); // All that remains is the revert string
    }



    function updateMessageStatus(bytes32 _messageId, TxStatus _status) external onlyOwner {
        executedMessages[_messageId] = _status;
    }

    function updateAuthVerifier(address _authVerifier) external onlyOwner {
        require(_authVerifier != address(0), "Cannot set to 0");
        authVerifier = _authVerifier;
    }

    uint256[48] private __gap;
}


contract MessageBusUpgradeable is MessageBusSenderUpgradeable, MessageBusReceiverUpgradeable {
    function initialize(address _gasFeePricing, address _authVerifier) external initializer {
        __Ownable_init_unchained();
        __Pausable_init_unchained();
        __MessageBusSender_init_unchained(_gasFeePricing);
        __MessageBusReceiver_init_unchained(_authVerifier);
    }

    // PAUSABLE FUNCTIONS ***/
    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }
}





contract TestMessageBusUpgradeable is MessageBusUpgradeable {

    function testExecuted(
        bytes32 messageId,
        TxStatus status,
        address  _dstAddress,
        uint64 srcChainId,
        uint64 srcNonce
    ) external {
        emit Executed(
            messageId,
            status,
            _dstAddress,
            srcChainId,
            srcNonce
        );
    }
    function testMessageSent(
        address sender,
        uint256 srcChainID,
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        uint64 nonce,
        bytes calldata options,
        uint256 fee,
        bytes32 messageId
    ) external {
        emit MessageSent(
            sender,
            srcChainID,
            receiver,
            dstChainId,
            message,
            nonce,
            options,
            fee,
            messageId
        );
    }
    function testCallReverted(
        string calldata reason
    ) external {
        emit CallReverted(
            reason
        );
    }
}



