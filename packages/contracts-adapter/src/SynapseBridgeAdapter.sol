// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter} from "./interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "./interfaces/ISynapseBridgeAdapterErrors.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SynapseBridgeAdapter is Ownable, ISynapseBridgeAdapter, ISynapseBridgeAdapterErrors {
    address public bridge;

    event BridgeSet(address bridge);
    event TokenAdded(address token, TokenType tokenType, bytes31 symbol);

    constructor(address owner) Ownable(owner) {}

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function addToken(address token, TokenType tokenType, bytes31 symbol) external onlyOwner {
        // TODO: implement
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
    function getSymbolByAddress(address token) external view returns (TokenType tokenType, bytes31 symbol) {
        // TODO: implement
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function getAddressBySymbol(bytes31 symbol) external view returns (TokenType tokenType, address token) {
        // TODO: implement
    }
}
