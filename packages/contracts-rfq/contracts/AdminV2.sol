// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IAdminV2} from "./interfaces/IAdminV2.sol";
import {IAdminV2Errors} from "./interfaces/IAdminV2Errors.sol";

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

contract AdminV2 is AccessControlEnumerable, IAdminV2, IAdminV2Errors {
    using SafeERC20 for IERC20;

    /// @notice Address reserved for native gas token (ETH on Ethereum and most L2s, AVAX on Avalanche, etc)
    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    bytes32 public constant PROVER_ROLE = keccak256("PROVER_ROLE");
    bytes32 public constant CANCELER_ROLE = keccak256("CANCELER_ROLE");
    bytes32 public constant GUARD_ROLE = keccak256("GUARD_ROLE");
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    uint256 public constant FEE_BPS = 1e6;
    uint256 public constant FEE_RATE_MAX = 0.01e6; // max 1% on origin amount
    uint256 public constant MIN_CANCEL_DELAY = 1 hours;
    uint256 public constant DEFAULT_CANCEL_DELAY = 1 days;

    /// @notice Protocol fee rate taken on origin amount deposited in origin chain
    uint256 public protocolFeeRate;

    /// @notice Protocol fee amounts accumulated
    mapping(address => uint256) public protocolFees;

    /// @notice Delay for a transaction after which it could be permisionlessly cancelled
    uint256 public cancelDelay;

    /// @notice This is deprecated and should not be used.
    /// @dev Use ZapNative V2 requests instead.
    uint256 public immutable chainGasAmount = 0;

    constructor(address _owner) {
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);
        _setCancelDelay(DEFAULT_CANCEL_DELAY);
    }

    /// @notice Allows the contract governor to set the cancel delay. The cancel delay is the time after the transaction
    /// deadline after which it can be permissionlessly cancelled, if it hasn't been proven by any of the Relayers.
    function setCancelDelay(uint256 newCancelDelay) external onlyRole(GOVERNOR_ROLE) {
        _setCancelDelay(newCancelDelay);
    }

    /// @notice Allows the contract governor to set the protocol fee rate. The protocol fee is taken from the origin
    /// amount only for completed and claimed transactions.
    /// @dev The protocol fee is abstracted away from the relayers, they always operate using the amounts after fees:
    /// what they see as the origin amount emitted in the log is what they get credited with.
    function setProtocolFeeRate(uint256 newFeeRate) external onlyRole(GOVERNOR_ROLE) {
        if (newFeeRate > FEE_RATE_MAX) revert FeeRateAboveMax();
        uint256 oldFeeRate = protocolFeeRate;
        protocolFeeRate = newFeeRate;
        emit FeeRateUpdated(oldFeeRate, newFeeRate);
    }

    /// @notice Allows the contract governor to sweep the accumulated protocol fees in the contract.
    function sweepProtocolFees(address token, address recipient) external onlyRole(GOVERNOR_ROLE) {
        uint256 feeAmount = protocolFees[token];
        if (feeAmount == 0) return; // skip if no accumulated fees

        protocolFees[token] = 0;
        emit FeesSwept(token, recipient, feeAmount);
        /// Sweep the fees as the last transaction action
        if (token == NATIVE_GAS_TOKEN) {
            Address.sendValue(payable(recipient), feeAmount);
        } else {
            IERC20(token).safeTransfer(recipient, feeAmount);
        }
    }

    /// @notice Internal function to set the cancel delay. Security checks are performed outside of this function.
    function _setCancelDelay(uint256 newCancelDelay) private {
        if (newCancelDelay < MIN_CANCEL_DELAY) revert CancelDelayBelowMin();
        uint256 oldCancelDelay = cancelDelay;
        cancelDelay = newCancelDelay;
        emit CancelDelayUpdated(oldCancelDelay, newCancelDelay);
    }
}
