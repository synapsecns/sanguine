// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridgeAdapter {
    enum TokenType {
        MintBurn,
        WithdrawDeposit
    }

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @notice Allows the contract owner to add a new token to the adapter.
    /// @dev Will revert in the following cases:
    /// - `token` has already been added, or is a zero address
    /// - `symbol` has already been added, or is a zero bytes31
    function addToken(address token, TokenType tokenType, bytes31 symbol) external;

    /// @notice Allows the contract owner to set the SynapseBridge address.
    /// @dev Will revert in the following cases:
    /// - `bridge` has already been set
    /// - `newBridge` is a zero address
    function setBridge(address newBridge) external;

    // ════════════════════════════════════════════════ USER FACING ════════════════════════════════════════════════════

    /// @notice Allows a user to bridge an ERC20 token to another chain. Fee is paid in native token,
    /// which is supplied as `msg.value`.
    /// @dev Will revert in the following cases:
    /// - bridge address has not been set
    /// - adapter is not connected to chain `dstEid`
    /// - `to` is the zero address
    /// - `token` has not been added to the adapter
    /// - `amount` is zero
    /// - `gasLimit` is below the required minimum
    /// - `msg.value` is not enough to cover the native fee
    function bridgeERC20(uint32 dstEid, address to, address token, uint256 amount, uint64 gasLimit) external payable;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the SynapseBridge address.
    function bridge() external view returns (address);

    /// @notice Returns the native fee for bridging a token to another chain.
    /// @dev Will revert in the following cases:
    /// - adapter is not connected to chain `dstEid`
    /// - `gasLimit` is below the required minimum
    function getNativeFee(uint32 dstEid, uint64 gasLimit) external view returns (uint256 nativeFee);

    /// @notice Returns the token type and symbol for a given token address.
    /// @dev Will revert in the following cases:
    /// - `token` has not been added to the adapter
    function getSymbolByAddress(address token) external view returns (TokenType tokenType, bytes31 symbol);

    /// @notice Returns the token type and address for a given symbol.
    /// @dev Will revert in the following cases:
    /// - `symbol` has not been added to the adapter
    function getAddressBySymbol(bytes31 symbol) external view returns (TokenType tokenType, address token);
}
