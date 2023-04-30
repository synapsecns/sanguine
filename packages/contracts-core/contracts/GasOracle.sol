// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessagingBase} from "./base/MessagingBase.sol";
import {InterfaceGasOracle} from "./interfaces/InterfaceGasOracle.sol";

/**
 * @notice `GasOracle` contract is responsible for tracking the gas data for both local and remote chains.
 * ## Local gas data tracking
 * - `GasOracle` is using the available tools such as `tx.gasprice` to track the time-averaged values
 * for different "gas statistics".
 * - These values are cached, so that the reported values are only changed when a big enough change is detected.
 * - The reported values are included in Origin's State, whenever a new message is sent.
 * > This leads to cached "chain gas data" being included in the Guard and Notary snapshots.
 * ## Remote gas data tracking
 * - To track gas data for the remote chains, GasOracle relies on the Notaries to pass the gas data alongside
 * their attestations.
 * - As the gas data is cached, this leads to a storage write only when the gas data
 * for the remote chain changes significantly.
 * - GasOracle is in charge of enforcing the optimistic periods for the gas data it gets from `Destination`.
 * - The optimistic period is smaller when the "gas statistics" are increasing, and bigger when they are decreasing.
 * > Reason for that is that the decrease of the gas price leads to lower execution/delivery tips, and we want the
 * > Executors to be protected against that.
 */
contract GasOracle is MessagingBase, InterfaceGasOracle {
    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 domain) MessagingBase("0.0.3", domain) {}

    /// @notice Initializes GasOracle contract:
    /// - msg.sender is set as contract owner
    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
    }

    /// @inheritdoc InterfaceGasOracle
    function getGasData() external view returns (uint256 paddedGasData) {}

    /// @inheritdoc InterfaceGasOracle
    function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength)
        external
        view
        returns (uint256 paddedTips)
    {}
}
