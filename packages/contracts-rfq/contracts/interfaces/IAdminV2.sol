// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IAdminV2 {
    event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay);
    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);

    event ProverAdded(address prover);
    event ProverRemoved(address prover);

    /// @notice Allows the governor to add a new prover to the contract.
    function addProver(address prover) external;

    /// @notice Allows the governor to remove a prover from the contract.
    function removeProver(address prover) external;

    /// @notice Allows the governor to set the cancel delay. The cancel delay is the time period after the transaction
    /// deadline during which a transaction can be permissionlessly cancelled if it hasn't been proven by any Relayer.
    function setCancelDelay(uint256 newCancelDelay) external;

    /// @notice Allows the governor to set the protocol fee rate. The protocol fee is taken from the origin
    /// amount and is only applied to completed and claimed transactions.
    /// @dev The protocol fee is abstracted away from the relayers; they always operate using the amounts after fees.
    /// The origin amount they see in the emitted log is what they get credited with.
    function setProtocolFeeRate(uint256 newFeeRate) external;

    /// @notice Allows the governor to withdraw the accumulated protocol fees from the contract.
    function sweepProtocolFees(address token, address recipient) external;

    /// @notice Returns the ID of the active prover, or zero if the prover is not currently active.
    function getActiveProverID(address prover) external view returns (uint16);

    /// @notice Returns the information about the prover with the provided address.
    /// @return proverID            The ID of the prover if it has been added before, or zero otherwise.
    /// @return activeFromTimestamp The timestamp when the prover becomes active, or zero if the prover isn't active.
    function getProverInfo(address prover) external view returns (uint16 proverID, uint256 activeFromTimestamp);

    /// @notice Returns the information about the prover with the provided ID.
    /// @return prover              The address of the prover with the provided ID, or zero the ID does not exist.
    /// @return activeFromTimestamp The timestamp when the prover becomes active, or zero if the prover isn't active.
    function getProverInfoByID(uint16 proverID) external view returns (address prover, uint256 activeFromTimestamp);

    /// @notice Returns the list of the active provers.
    function getProvers() external view returns (address[] memory);
}
