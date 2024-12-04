// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {IAdminV2} from "./interfaces/IAdminV2.sol";
import {IAdminV2Errors} from "./interfaces/IAdminV2Errors.sol";

// ═════════════════════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════════════════════════

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @title AdminV2
/// @notice Provides administrative functions and controls for managing the FastBridgeV2 contract,
/// including access control and configuration settings.
contract AdminV2 is AccessControlEnumerable, IAdminV2, IAdminV2Errors {
    using SafeERC20 for IERC20;

    /// @notice Struct for storing information about a prover.
    /// @param id                   The ID of the prover: its position in `_allProvers` plus one,
    ///                             or zero if the prover has never been added.
    /// @param activeFromTimestamp  The timestamp at which the prover becomes active,
    ///                             or zero if the prover has never been added or is no longer active.
    struct ProverInfo {
        uint16 id;
        uint240 activeFromTimestamp;
    }

    /// @notice The address reserved for the native gas token (ETH on Ethereum and most L2s, AVAX on Avalanche, etc.).
    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @notice The role identifier for the Quoter API's off-chain authentication.
    /// @dev Only addresses with this role can post FastBridge quotes to the API.
    bytes32 public constant QUOTER_ROLE = keccak256("QUOTER_ROLE");

    /// @notice The role identifier for the Guard's on-chain authentication in FastBridge.
    /// @dev Only addresses with this role can dispute submitted relay proofs during the dispute period.
    bytes32 public constant GUARD_ROLE = keccak256("GUARD_ROLE");

    /// @notice The role identifier for the Canceler's on-chain authentication in FastBridge.
    /// @dev Only addresses with this role can cancel a FastBridge transaction without the cancel delay.
    bytes32 public constant CANCELER_ROLE = keccak256("CANCELER_ROLE");

    /// @notice The role identifier for the Governor's on-chain administrative authority.
    /// @dev Only addresses with this role can perform administrative tasks within the contract.
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    /// @notice The denominator for fee rates, representing 100%.
    uint256 public constant FEE_BPS = 1e6;
    /// @notice The maximum protocol fee rate: 1% of the origin amount.
    uint256 public constant FEE_RATE_MAX = 0.01e6;

    /// @notice The minimum cancel delay that can be set by the governor.
    uint256 public constant MIN_CANCEL_DELAY = 1 hours;
    /// @notice The default cancel delay set during contract deployment.
    uint256 public constant DEFAULT_CANCEL_DELAY = 1 days;

    /// @notice The minimum prover timeout that can be set by the governor.
    uint256 public constant MIN_PROVER_TIMEOUT = 1 minutes;
    /// @notice The default prover timeout set during contract deployment.
    uint256 public constant DEFAULT_PROVER_TIMEOUT = 30 minutes;

    /// @notice The protocol fee rate taken on the origin amount deposited in the origin chain.
    uint256 public protocolFeeRate;

    /// @notice The accumulated protocol fee amounts.
    mapping(address => uint256) public protocolFees;

    /// @notice The delay period after which a transaction can be permissionlessly cancelled.
    uint256 public cancelDelay;

    /// @notice The timeout period that is used to temporarily disactivate a disputed prover.
    uint256 public proverTimeout;

    /// @notice A list of all provers ever added to the contract. Can hold up to 2^16-1 provers.
    address[] private _allProvers;
    /// @notice A mapping of provers to their information: id and activeFromTimestamp.
    mapping(address => ProverInfo) private _proverInfos;

    /// @notice This variable is deprecated and should not be used.
    /// @dev Use ZapNative V2 requests instead.
    uint256 public immutable chainGasAmount = 0;

    constructor(address defaultAdmin) {
        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _setCancelDelay(DEFAULT_CANCEL_DELAY);
        _setProverTimeout(DEFAULT_PROVER_TIMEOUT);
    }

    /// @inheritdoc IAdminV2
    function addProver(address prover) external onlyRole(GOVERNOR_ROLE) {
        if (getActiveProverID(prover) != 0) revert ProverAlreadyActive();
        ProverInfo storage $ = _proverInfos[prover];
        // Add the prover to the list of all provers and record its id (its position + 1),
        // if this has not already been done.
        if ($.id == 0) {
            _allProvers.push(prover);
            uint256 id = _allProvers.length;
            if (id > type(uint16).max) revert ProverCapacityExceeded();
            // Note: this is a storage write.
            $.id = uint16(id);
        }
        // Update the activeFrom timestamp.
        // Note: this is a storage write.
        $.activeFromTimestamp = uint240(block.timestamp);
        emit ProverAdded(prover);
    }

    /// @inheritdoc IAdminV2
    function removeProver(address prover) external onlyRole(GOVERNOR_ROLE) {
        if (getActiveProverID(prover) == 0) revert ProverNotActive();
        // We never remove provers from the list of all provers to preserve their IDs,
        // so we just need to reset the activeFrom timestamp.
        _proverInfos[prover].activeFromTimestamp = 0;
        emit ProverRemoved(prover);
    }

    /// @inheritdoc IAdminV2
    function setCancelDelay(uint256 newCancelDelay) external onlyRole(GOVERNOR_ROLE) {
        _setCancelDelay(newCancelDelay);
    }

    /// @inheritdoc IAdminV2
    function setProverTimeout(uint256 newProverTimeout) external onlyRole(GOVERNOR_ROLE) {
        _setProverTimeout(newProverTimeout);
    }

    /// @inheritdoc IAdminV2
    function setProtocolFeeRate(uint256 newFeeRate) external onlyRole(GOVERNOR_ROLE) {
        if (newFeeRate > FEE_RATE_MAX) revert FeeRateAboveMax();
        uint256 oldFeeRate = protocolFeeRate;
        protocolFeeRate = newFeeRate;
        emit FeeRateUpdated(oldFeeRate, newFeeRate);
    }

    /// @inheritdoc IAdminV2
    function sweepProtocolFees(address token, address recipient) external onlyRole(GOVERNOR_ROLE) {
        // Early exit if no accumulated fees.
        uint256 feeAmount = protocolFees[token];
        if (feeAmount == 0) return;
        // Reset the accumulated fees first.
        protocolFees[token] = 0;
        emit FeesSwept(token, recipient, feeAmount);
        // Sweep the fees as the last transaction action.
        if (token == NATIVE_GAS_TOKEN) {
            Address.sendValue(payable(recipient), feeAmount);
        } else {
            IERC20(token).safeTransfer(recipient, feeAmount);
        }
    }

    /// @inheritdoc IAdminV2
    function getProvers() external view returns (address[] memory provers) {
        uint256 length = _allProvers.length;
        // Calculate the number of active provers.
        uint256 activeProversCount = 0;
        for (uint256 i = 0; i < length; i++) {
            if (getActiveProverID(_allProvers[i]) != 0) {
                activeProversCount++;
            }
        }
        // Do the second pass to populate the provers array.
        provers = new address[](activeProversCount);
        uint256 activeProversIndex = 0;
        for (uint256 i = 0; i < length; i++) {
            address prover = _allProvers[i];
            if (getActiveProverID(prover) != 0) {
                provers[activeProversIndex++] = prover;
            }
        }
    }

    /// @inheritdoc IAdminV2
    function getProverInfo(address prover) external view returns (uint16 proverID, uint256 activeFromTimestamp) {
        proverID = _proverInfos[prover].id;
        activeFromTimestamp = _proverInfos[prover].activeFromTimestamp;
    }

    /// @inheritdoc IAdminV2
    function getProverInfoByID(uint16 proverID) external view returns (address prover, uint256 activeFromTimestamp) {
        if (proverID == 0 || proverID > _allProvers.length) return (address(0), 0);
        prover = _allProvers[proverID - 1];
        activeFromTimestamp = _proverInfos[prover].activeFromTimestamp;
    }

    /// @inheritdoc IAdminV2
    function getActiveProverID(address prover) public view returns (uint16) {
        // Aggregate the read operations from the same storage slot.
        uint16 id = _proverInfos[prover].id;
        uint256 activeFromTimestamp = _proverInfos[prover].activeFromTimestamp;
        // Return zero if the prover has never been added or is no longer active.
        if (activeFromTimestamp == 0 || activeFromTimestamp > block.timestamp) return 0;
        return id;
    }

    /// @notice Internal logic to set the cancel delay. Security checks are performed outside of this function.
    /// @dev This function is marked as private to prevent child contracts from calling it directly.
    function _setCancelDelay(uint256 newCancelDelay) private {
        if (newCancelDelay < MIN_CANCEL_DELAY) revert CancelDelayBelowMin();
        uint256 oldCancelDelay = cancelDelay;
        cancelDelay = newCancelDelay;
        emit CancelDelayUpdated(oldCancelDelay, newCancelDelay);
    }

    /// @notice Internal logic to set the prover timeout. Security checks are performed outside of this function.
    /// @dev This function is marked as private to prevent child contracts from calling it directly.
    function _setProverTimeout(uint256 newProverTimeout) private {
        if (newProverTimeout < MIN_PROVER_TIMEOUT) revert ProverTimeoutBelowMin();
        uint256 oldProverTimeout = proverTimeout;
        proverTimeout = newProverTimeout;
        emit ProverTimeoutUpdated(oldProverTimeout, newProverTimeout);
    }
}
