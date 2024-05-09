// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Every Module may opt a different method to confirm the verified entries on destination chain,
/// therefore this is not a part of a common interface.
interface IInterchainModule {
    error InterchainModule__CallerNotInterchainDB(address caller);
    error InterchainModule__ChainIdNotRemote(uint64 chainId);
    error InterchainModule__FeeAmountBelowMin(uint256 feeAmount, uint256 minRequired);

    function requestEntryVerification(uint64 dstChainId, bytes memory versionedEntry) external payable;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function getModuleFee(uint64 dstChainId) external view returns (uint256);
}
