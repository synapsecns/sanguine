// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter} from "./interfaces/ISynapseBridgeAdapter.sol";

contract SynapseBridgeAdapter is ISynapseBridgeAdapter {
    address public bridge;

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function addToken(address token, TokenType tokenType, bytes31 symbol) external {
        // TODO: implement
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function setBridge(address newBridge) external {
        // TODO: implement
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
