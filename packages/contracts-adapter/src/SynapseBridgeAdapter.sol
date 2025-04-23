// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IBurnableToken} from "./interfaces/IBurnableToken.sol";
import {ISynapseBridge} from "./interfaces/ISynapseBridge.sol";
import {ISynapseBridgeAdapter} from "./interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "./interfaces/ISynapseBridgeAdapterErrors.sol";
import {BridgeMessage} from "./libs/BridgeMessage.sol";
import {ReadableSymbol} from "./libs/ReadableSymbol.sol";

import {MessagingFee, OApp, Origin} from "@layerzerolabs/oapp-evm/contracts/oapp/OApp.sol";
import {OptionsBuilder} from "@layerzerolabs/oapp-evm/contracts/oapp/libs/OptionsBuilder.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract SynapseBridgeAdapter is OApp, ISynapseBridgeAdapter, ISynapseBridgeAdapterErrors {
    using OptionsBuilder for bytes;
    using SafeERC20 for IERC20;

    struct TokenAddress {
        TokenType tokenType;
        address token;
    }

    struct TokenSymbol {
        TokenType tokenType;
        bytes31 symbol;
    }

    uint64 public constant MIN_GAS_LIMIT = 100_000;

    address public bridge;

    mapping(address => TokenSymbol) internal _symbolByAddress;
    mapping(bytes31 => TokenAddress) internal _addressBySymbol;

    event BridgeSet(address bridge);
    event TokenAdded(address token, TokenType tokenType, bytes31 symbol);
    event TokenSent(uint32 indexed dstEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);
    event TokenReceived(uint32 indexed srcEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);

    constructor(address endpoint_, address owner_) OApp(endpoint_, owner_) Ownable(owner_) {}

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function addToken(address token, TokenType tokenType, bytes31 symbol) external onlyOwner {
        // Check: new parameters
        if (token == address(0)) revert SBA__ZeroAddress();
        if (symbol == bytes31(0)) revert SBA__ZeroSymbol();
        // Check: existing state
        if (_symbolByAddress[token].symbol != bytes31(0)) revert SBA__TokenAlreadyAdded(token);
        if (_addressBySymbol[symbol].token != address(0)) revert SBA__SymbolAlreadyAdded(symbol);
        // Store
        _symbolByAddress[token] = TokenSymbol(tokenType, symbol);
        _addressBySymbol[symbol] = TokenAddress(tokenType, token);
        emit TokenAdded(token, tokenType, symbol);
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function setBridge(address newBridge) external onlyOwner {
        // Check: new parameters
        if (newBridge == address(0)) revert SBA__ZeroAddress();
        // Check: existing state
        if (bridge != address(0)) revert SBA__BridgeAlreadySet();
        // Store
        bridge = newBridge;
        emit BridgeSet(newBridge);
    }

    // ════════════════════════════════════════════════ USER FACING ════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function bridgeERC20(uint32 dstEid, address to, address token, uint256 amount, uint64 gasLimit) external payable {
        // Verify the parameters
        if (to == address(0)) revert SBA__ZeroAddress();
        if (amount == 0) revert SBA__ZeroAmount();
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        // Cache bridge address
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        // Cache token type and symbol (note: this reverts if token is unknown)
        (TokenType tokenType, bytes31 symbol) = getSymbolByAddress(token);
        // Burn tokens from sender or deposit them into the bridge as prerequisite
        if (tokenType == TokenType.MintBurn) {
            IBurnableToken(token).burnFrom(msg.sender, amount);
        } else {
            IERC20(token).transferFrom(msg.sender, cachedBridge, amount);
        }
        // Send the bridge message
        bytes32 guid = _lzSend({
            _dstEid: dstEid,
            _message: BridgeMessage.encodeBridgeMessage(to, symbol, amount),
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _fee: MessagingFee({nativeFee: msg.value, lzTokenFee: 0}),
            _refundAddress: msg.sender
        }).guid;
        emit TokenSent(dstEid, to, token, amount, guid);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function getNativeFee(uint32 dstEid, uint64 gasLimit) external view returns (uint256 nativeFee) {
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        // Since all the messages have the same length, we can use arbitrary data for fee estimation
        bytes memory message = BridgeMessage.encodeBridgeMessage(address(0), 0, 0);
        return _quote({
            _dstEid: dstEid,
            _message: message,
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _payInLzToken: false
        }).nativeFee;
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function getReadableSymbolByAddress(address token)
        external
        view
        returns (TokenType tokenType, string memory readableSymbol)
    {
        bytes31 symbol;
        (tokenType, symbol) = getSymbolByAddress(token);
        readableSymbol = ReadableSymbol.toString(symbol);
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function getAddressByReadableSymbol(string memory readableSymbol)
        external
        view
        returns (TokenType tokenType, address token)
    {
        bytes31 symbol = ReadableSymbol.toBytes31(readableSymbol);
        return getAddressBySymbol(symbol);
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function getSymbolByAddress(address token) public view returns (TokenType tokenType, bytes31 symbol) {
        tokenType = _symbolByAddress[token].tokenType;
        symbol = _symbolByAddress[token].symbol;
        if (symbol == bytes31(0)) revert SBA__TokenUnknown(token);
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function getAddressBySymbol(bytes31 symbol) public view returns (TokenType tokenType, address token) {
        tokenType = _addressBySymbol[symbol].tokenType;
        token = _addressBySymbol[symbol].token;
        if (token == address(0)) revert SBA__SymbolUnknown(symbol);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Handles the received message from origin's Adapter counterpart.
    /// All validation checks for integrity of the message have been performed at this point.
    function _lzReceive(
        Origin calldata _origin,
        bytes32 _guid,
        bytes calldata _message,
        address, // _executor,
        bytes calldata // _extraData
    )
        internal
        override
    {
        (address to, bytes31 symbol, uint256 amount) = BridgeMessage.decodeBridgeMessage(_message);
        // Cache guid to avoid stack too deep error
        bytes32 guid = _guid;
        // Cache bridge address
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        // Cache token type and address (note: this reverts if symbol is unknown)
        (TokenType tokenType, address token) = getAddressBySymbol(symbol);
        // Mint or withdraw tokens from the bridge as the result of the message
        // Note: the fees are set to 0 to enable 1:1 bridging
        // Note: guid is used as "kappa" (which is SynapseBridge's own global unique identifier),
        // clashes between two different systems are not possible due to being different keccak256 hashes.
        if (tokenType == TokenType.MintBurn) {
            ISynapseBridge(cachedBridge).mint(to, token, amount, 0, guid);
        } else {
            ISynapseBridge(cachedBridge).withdraw(to, token, amount, 0, guid);
        }
        emit TokenReceived(_origin.srcEid, to, token, amount, guid);
    }
}
