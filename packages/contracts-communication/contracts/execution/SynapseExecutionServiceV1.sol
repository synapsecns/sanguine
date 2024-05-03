// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceEvents} from "../events/SynapseExecutionServiceEvents.sol";
import {ClaimableFees} from "../fees/ClaimableFees.sol";
import {IExecutionService, ISynapseExecutionServiceV1} from "../interfaces/ISynapseExecutionServiceV1.sol";
import {IGasOracle} from "../interfaces/IGasOracle.sol";
import {OptionsLib, OptionsV1} from "../libs/Options.sol";
import {VersionedPayloadLib} from "../libs/VersionedPayload.sol";

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract SynapseExecutionServiceV1 is
    AccessControlUpgradeable,
    ClaimableFees,
    SynapseExecutionServiceEvents,
    ISynapseExecutionServiceV1
{
    /// @custom:storage-location erc7201:Synapse.ExecutionService.V1
    struct SynapseExecutionServiceV1Storage {
        address executorEOA;
        address gasOracle;
        uint256 globalMarkup;
        uint256 claimerFraction;
    }

    // keccak256(abi.encode(uint256(keccak256("Synapse.ExecutionService.V1")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant SYNAPSE_EXECUTION_SERVICE_V1_STORAGE_LOCATION =
        0xabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100;
    uint256 private constant WAD = 10 ** 18;

    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");
    bytes32 public constant IC_CLIENT_ROLE = keccak256("IC_CLIENT_ROLE");

    constructor() {
        // Ensure that the implementation contract could not be initialized
        _disableInitializers();
    }

    function initialize(address admin) external virtual initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setClaimerFraction(uint256 claimerFraction_) external virtual onlyRole(GOVERNOR_ROLE) {
        if (claimerFraction_ > MAX_CLAIMER_FRACTION) {
            revert ClaimableFees__ClaimerFractionExceedsMax(claimerFraction_);
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.claimerFraction = claimerFraction_;
        emit ClaimerFractionSet(claimerFraction_);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setExecutorEOA(address executorEOA_) external virtual onlyRole(GOVERNOR_ROLE) {
        if (executorEOA_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.executorEOA = executorEOA_;
        emit ExecutorEOASet(executorEOA_);
        emit FeeRecipientSet(executorEOA_);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setGasOracle(address gasOracle_) external virtual onlyRole(GOVERNOR_ROLE) {
        if (gasOracle_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.gasOracle = gasOracle_;
        emit GasOracleSet(gasOracle_);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setGlobalMarkup(uint256 globalMarkup_) external virtual onlyRole(GOVERNOR_ROLE) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.globalMarkup = globalMarkup_;
        emit GlobalMarkupSet(globalMarkup_);
    }

    /// @inheritdoc IExecutionService
    function requestTxExecution(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        bytes calldata options
    )
        external
        payable
        virtual
        onlyRole(IC_CLIENT_ROLE)
    {
        uint256 requiredFee = getExecutionFee(dstChainId, txPayloadSize, options);
        if (msg.value < requiredFee) {
            revert SynapseExecutionService__FeeAmountTooLow({actual: msg.value, required: requiredFee});
        }
        emit ExecutionRequested({transactionId: transactionId, client: msg.sender, executionFee: msg.value});
    }

    /// @inheritdoc IExecutionService
    function getExecutionFee(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes calldata options
    )
        public
        view
        virtual
        returns (uint256 executionFee)
    {
        address cachedGasOracle = gasOracle();
        if (cachedGasOracle == address(0)) {
            revert SynapseExecutionService__GasOracleNotSet();
        }
        // ExecutionServiceV1 implementation only supports Options V1.
        // Following versions will be supported by the future implementations.
        uint16 version = VersionedPayloadLib.getVersion(options);
        if (version > OptionsLib.OPTIONS_V1) {
            revert SynapseExecutionService__OptionsVersionNotSupported(version);
        }
        OptionsV1 memory optionsV1 = OptionsLib.decodeOptionsV1(options);
        executionFee = IGasOracle(cachedGasOracle).estimateTxCostInLocalUnits({
            remoteChainId: dstChainId,
            gasLimit: optionsV1.gasLimit,
            calldataSize: txPayloadSize
        });
        if (optionsV1.gasAirdrop > 0) {
            executionFee += IGasOracle(cachedGasOracle).convertRemoteValueToLocalUnits({
                remoteChainId: dstChainId,
                value: optionsV1.gasAirdrop
            });
        }
        executionFee += executionFee * globalMarkup() / WAD;
    }

    /// @inheritdoc IExecutionService
    function executorEOA() public view virtual returns (address) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.executorEOA;
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function gasOracle() public view virtual returns (address) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.gasOracle;
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function globalMarkup() public view virtual returns (uint256) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.globalMarkup;
    }

    /// @notice Returns the amount of fees that can be claimed.
    function getClaimableAmount() public view virtual override returns (uint256) {
        return address(this).balance;
    }

    /// @notice Returns the fraction of the fees that the claimer will receive.
    /// The result is in the range [0, 1e18], where 1e18 is 100%.
    function getClaimerFraction() public view virtual override returns (uint256) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.claimerFraction;
    }

    /// @notice Returns the address that will receive the claimed fees.
    function getFeeRecipient() public view virtual override returns (address) {
        return executorEOA();
    }

    /// @dev Hook that is called before the fees are claimed.
    /// Useful if the inheriting contract needs to manage the state when the fees are claimed.
    // solhint-disable-next-line no-empty-blocks
    function _beforeFeesClaimed(uint256, uint256) internal override {
        // No op, as the claimable amount is tracked as the contract balance
    }

    /// @dev ERC-7201 slot accessor
    function _getSynapseExecutionServiceV1Storage() private pure returns (SynapseExecutionServiceV1Storage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := SYNAPSE_EXECUTION_SERVICE_V1_STORAGE_LOCATION
        }
    }
}
