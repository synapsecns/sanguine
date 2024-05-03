// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Every Module may opt a different method to confirm the verified entries on destination chain,
/// therefore this is not a part of a common interface.
interface IInterchainModule {
    error InterchainModule__NotInterchainDB(address caller);
    error InterchainModule__IncorrectSourceChainId(uint64 chainId);
    error InterchainModule__FeeAmountBelowMin(uint256 feeAmount, uint256 minRequired);
    error InterchainModule__SameChainId(uint64 chainId);

    /// @notice Request the verification of a batch from the Interchain DataBase by the module.
    /// If the batch is not yet finalized, the verification on destination chain will be delayed until
    /// the finalization is done and batch root is saved on the source chain.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.verifyRemoteBatch(batch)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param dstChainId       The chain id of the destination chain
    /// @param versionedBatch   The versioned batch to verify
    function requestBatchVerification(uint64 dstChainId, bytes memory versionedBatch) external payable;

    /// @notice Get the Module fee for verifying a batch on the specified destination chain.
    /// @param dstChainId   The chain id of the destination chain
    /// @param dbNonce      The database nonce of the batch on the source chain
    function getModuleFee(uint64 dstChainId, uint64 dbNonce) external view returns (uint256);
}
