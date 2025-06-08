// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridgeAdapter {
    enum TokenType {
        Unknown,
        MintBurn,
        WithdrawDeposit
    }

    struct RemoteToken {
        uint32 eid;
        address addr;
    }

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @notice Allows the contract owner to add a new token to the adapter,
    /// or add new remote tokens to an existing token.
    /// @dev Will revert in the following cases:
    /// - `token` or any of the `remoteTokens.addr` is a zero address
    /// - `tokenType` is `TokenType.Unknown`
    /// - `token` has already been added with a different `tokenType`.
    /// - `token` has already a remote token assigned for any of the `eid`.
    /// - Any of the remote tokens has already been used for its `eid` with any of the local tokens.
    function addToken(address token, TokenType tokenType, RemoteToken[] memory remoteTokens) external;

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

    /// @notice Returns the local address for a given eid and remote address.
    /// @dev Will return `address(0)` if the remote address has not been added for the given eid.
    function getLocalAddress(uint32 eid, address remoteAddr) external view returns (address localAddr);

    /// @notice Returns the remote address for a given local token and eid.
    /// @dev Will return `address(0)` if the token has not been added for the given eid.
    function getRemoteAddress(uint32 eid, address localAddr) external view returns (address remoteAddr);

    /// @notice Returns the token type for a given local address.
    /// @dev Will return `TokenType.Unknown` if the token has not been added.
    function getTokenType(address localAddr) external view returns (TokenType tokenType);
}
