// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";

import {IAdmin} from "./interfaces/IAdmin.sol";
import {UniversalTokenLib} from "UniversalToken.sol";

contract Admin is IAdmin, AccessControlEnumerable {
    using UniversalTokenLib for address;

    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");
    bytes32 public constant REFUNDER_ROLE = keccak256("REFUNDER_ROLE");
    bytes32 public constant GUARD_ROLE = keccak256("GUARD_ROLE");
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    uint256 public constant FEE_BPS = 1e6;
    uint256 public constant FEE_RATE_MAX = 0.01e6; // max 1% on origin amount

    /// @notice Protocol fee rate taken on origin amount deposited in origin chain
    uint256 public protocolFeeRate;

    /// @notice Protocol fee amounts accumulated
    mapping(address => uint256) public protocolFees;

    /// @notice Chain gas amount to forward as rebate if requested
    uint256 public chainGasAmount;

    constructor(address _owner) {
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);
    }

    function setProtocolFeeRate(uint256 newFeeRate) external onlyRole(GOVERNOR_ROLE) {
        require(newFeeRate <= FEE_RATE_MAX, "newFeeRate > max");
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

    function setChainGasAmount(uint256 newChainGasAmount) external onlyRole(GOVERNOR_ROLE) {
        uint256 oldChainGasAmount = chainGasAmount;
        chainGasAmount = newChainGasAmount;
        emit ChainGasAmountUpdated(oldChainGasAmount, newChainGasAmount);
    }
}
