// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";

import {UniversalTokenLib} from "./libs/UniversalToken.sol";
import {IAdminV2} from "./interfaces/IAdminV2.sol";
import {IAdminV2Errors} from "./interfaces/IAdminV2Errors.sol";

contract AdminV2 is AccessControlEnumerable, IAdminV2, IAdminV2Errors {
    using UniversalTokenLib for address;

    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");
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

    function setCancelDelay(uint256 newCancelDelay) external onlyRole(GOVERNOR_ROLE) {
        _setCancelDelay(newCancelDelay);
    }

    function setProtocolFeeRate(uint256 newFeeRate) external onlyRole(GOVERNOR_ROLE) {
        if (newFeeRate > FEE_RATE_MAX) revert FeeRateAboveMax();
        uint256 oldFeeRate = protocolFeeRate;
        protocolFeeRate = newFeeRate;
        emit FeeRateUpdated(oldFeeRate, newFeeRate);
    }

    function sweepProtocolFees(address token, address recipient) external onlyRole(GOVERNOR_ROLE) {
        uint256 feeAmount = protocolFees[token];
        if (feeAmount == 0) return; // skip if no accumulated fees

        protocolFees[token] = 0;
        token.universalTransfer(recipient, feeAmount);
        emit FeesSwept(token, recipient, feeAmount);
    }

    /// @notice Internal function to set the cancel delay. Security checks are performed outside of this function.
    function _setCancelDelay(uint256 newCancelDelay) private {
        if (newCancelDelay < MIN_CANCEL_DELAY) revert CancelDelayBelowMin();
        uint256 oldCancelDelay = cancelDelay;
        cancelDelay = newCancelDelay;
        emit CancelDelayUpdated(oldCancelDelay, newCancelDelay);
    }
}
