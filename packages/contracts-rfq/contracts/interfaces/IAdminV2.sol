// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IAdminV2 {
    event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay);
    event DeployBlockSet(uint256 blockNumber);
    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);

    /// @notice Allows the governor to set the cancel delay. The cancel delay is the time period after the transaction
    /// deadline during which a transaction can be permissionlessly cancelled if it hasn't been proven by any Relayer.
    function setCancelDelay(uint256 newCancelDelay) external;

    /// @notice Allows the default admin to set the deploy block.
    /// @dev This is only relevant for chains like Arbitrum that implement the `block.number` as the underlying L1
    /// block number rather than the chain's native block number.
    function setDeployBlock(uint256 blockNumber) external;

    /// @notice Allows the governor to set the protocol fee rate. The protocol fee is taken from the origin
    /// amount and is only applied to completed and claimed transactions.
    /// @dev The protocol fee is abstracted away from the relayers; they always operate using the amounts after fees.
    /// The origin amount they see in the emitted log is what they get credited with.
    function setProtocolFeeRate(uint256 newFeeRate) external;

    /// @notice Allows the governor to withdraw the accumulated protocol fees from the contract.
    function sweepProtocolFees(address token, address recipient) external;
}
