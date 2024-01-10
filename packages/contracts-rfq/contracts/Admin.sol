// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

import {UniversalTokenLib} from "./libs/UniversalToken.sol";
import {IAdmin} from "./interfaces/IAdmin.sol";

contract Admin is IAdmin, AccessControl {
    using UniversalTokenLib for address;

    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");
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

    modifier onlyGuard() {
        require(hasRole(GUARD_ROLE, msg.sender), "Caller is not a guard");
        _;
    }

    modifier onlyRelayer() {
        require(hasRole(RELAYER_ROLE, msg.sender), "Caller is not a relayer");
        _;
    }

    modifier onlyGovernor() {
        require(hasRole(GOVERNOR_ROLE, msg.sender), "Caller is not a governor");
        _;
    }

    constructor(address _owner) {
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);
    }

    function addRelayer(address _relayer) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _grantRole(RELAYER_ROLE, _relayer);
        emit RelayerAdded(_relayer);
    }

    function removeRelayer(address _relayer) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _revokeRole(RELAYER_ROLE, _relayer);
        emit RelayerRemoved(_relayer);
    }

    function addGuard(address _guard) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _grantRole(GUARD_ROLE, _guard);
        emit GuardAdded(_guard);
    }

    function removeGuard(address _guard) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _revokeRole(GUARD_ROLE, _guard);
        emit GuardRemoved(_guard);
    }

    function addGovernor(address _governor) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _grantRole(GOVERNOR_ROLE, _governor);
        emit GovernorAdded(_governor);
    }

    function removeGovernor(address _governor) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _revokeRole(GOVERNOR_ROLE, _governor);
        emit GovernorRemoved(_governor);
    }

    function setProtocolFeeRate(uint256 newFeeRate) external onlyGovernor {
        require(newFeeRate <= FEE_RATE_MAX, "newFeeRate > max");
        uint256 oldFeeRate = protocolFeeRate;
        protocolFeeRate = newFeeRate;
        emit FeeRateUpdated(oldFeeRate, newFeeRate);
    }

    function sweepProtocolFees(address token, address recipient) external onlyGovernor {
        uint256 feeAmount = protocolFees[token];
        if (feeAmount == 0) return; // skip if no accumulated fees

        protocolFees[token] = 0;
        token.universalTransfer(recipient, feeAmount);
        emit FeesSwept(token, recipient, feeAmount);
    }

    function setChainGasAmount(uint256 newChainGasAmount) external onlyGovernor {
        uint256 oldChainGasAmount = chainGasAmount;
        chainGasAmount = newChainGasAmount;
        emit ChainGasAmountUpdated(oldChainGasAmount, newChainGasAmount);
    }
}
