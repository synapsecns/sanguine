// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @notice Interface for the InterChain ERC20 token (ICERC20).
/// Burning/minting of this token is rate limited for Bridge and unlimited for Processor contracts.
/// - Bridge contracts are minting or burning interchain tokens to facilitate the transfer of tokens between chains.
/// - Processor contracts are turning the interchain tokens into the underlying token and vice versa.
/// NOTE: the Bridge rate limits for initial issuance of interchain tokens make it possible to
/// let the Processor contract mint or burn the tokens without rate limits.
interface ICERC20 is IERC20 {
    /// @notice Emitted when a new burn limit is set for the `bridge`.
    /// @param bridge       The bridge contract address
    /// @param limit        The new burn limit
    event BurnLimitSet(address indexed bridge, uint256 limit);

    /// @notice Emitted when a new mint limit is set for the `bridge`.
    /// @param bridge       The bridge contract address
    /// @param limit        The new mint limit
    event MintLimitSet(address indexed bridge, uint256 limit);

    /// @notice Burn `amount` tokens from `msg.sender`
    /// @dev Could be called by Bridge or Processor contracts.
    /// - Bridge's burning limit is applied: will revert if the limit is
    /// exceeded.
    /// - Processor could burn any amount of tokens.
    /// @param amount       The amount of tokens to burn
    function burn(uint256 amount) external;

    /// @notice Burn `amount` tokens from `account`
    /// @dev Could be called by Bridge or Processor contracts.
    /// - Bridge's burning limit is applied: will revert if the limit is
    /// exceeded.
    /// - Processor could burn any amount of tokens.
    /// NOTE: this spends transfer allowance from `account` to `msg.sender`.
    /// @param account      The account to burn tokens from
    /// @param amount       The amount of tokens to burn
    function burnFrom(address account, uint256 amount) external;

    /// @notice Mint `amount` tokens to the `account`
    /// @dev Could be called by Bridge or Processor contracts.
    /// - Bridge's minting limit is applied: will revert if the limit is
    /// exceeded.
    /// - Processor could mint any amount of tokens.
    /// @param account      The account to mint tokens to
    /// @param amount       The amount of tokens to mint
    function mint(address account, uint256 amount) external;

    /// @notice Returns the maximum amount of tokens that could be burned by `bridge` at the moment.
    /// Bridge has a total burn limit, which is spent by burning tokens. This limit is replenished
    /// at a constant rate, and can never surpass the total burn limit for the bridge.
    function getCurrentBurnLimit(address bridge) external view returns (uint256);

    /// @notice Returns the maximum amount of tokens that could be minted by `bridge` at the moment.
    /// Bridge has a total mint limit, which is spent by minting tokens. This limit is replenished
    /// at a constant rate, and can never surpass the total mint limit for the bridge.
    function getCurrentMintLimit(address bridge) external view returns (uint256);

    /// @notice Returns the maximum value that `bridge` burn limit could reach.
    /// Note: it takes 24 hours to fully replenish the burn limit (assuming no mints are performed during that time).
    function getTotalBurnLimit(address bridge) external view returns (uint256);

    /// @notice Returns the maximum value that `bridge` mint limit could reach.
    /// Note: it takes 24 hours to fully replenish the mint limit (assuming no burns are performed during that time).
    function getTotalMintLimit(address bridge) external view returns (uint256);
}
