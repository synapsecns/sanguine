// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter} from "./interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "./interfaces/ISynapseBridgeAdapterErrors.sol";
import {ReadableSymbol} from "./libs/ReadableSymbol.sol";

import {OApp, Origin} from "@layerzerolabs/oapp-evm/contracts/oapp/OApp.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SynapseBridgeAdapter is OApp, ISynapseBridgeAdapter, ISynapseBridgeAdapterErrors {
    struct TokenAddress {
        TokenType tokenType;
        address token;
    }

    struct TokenSymbol {
        TokenType tokenType;
        bytes31 symbol;
    }

    address public bridge;

    mapping(address => TokenSymbol) internal _symbolByAddress;
    mapping(bytes31 => TokenAddress) internal _addressBySymbol;

    event BridgeSet(address bridge);
    event TokenAdded(address token, TokenType tokenType, bytes31 symbol);

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
        // TODO: implement
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function getNativeFee(uint32 dstEid, uint64 gasLimit) external view returns (uint256 nativeFee) {
        // TODO: implement
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
        address _executor,
        bytes calldata _extraData
    )
        internal
        override
    {
        // TODO: implement
    }
}
