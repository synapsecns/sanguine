// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceEvents} from "../events/SynapseExecutionServiceEvents.sol";
import {ClaimableFees} from "../fees/ClaimableFees.sol";
import {ISynapseExecutionServiceV1} from "../interfaces/ISynapseExecutionServiceV1.sol";
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

    /// @dev The storage location of the SynapseExecutionServiceV1Storage struct as per ERC-7201.
    /// keccak256(abi.encode(uint256(keccak256("Synapse.ExecutionService.V1")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant SYNAPSE_EXECUTION_SERVICE_V1_STORAGE_LOCATION =
        0xabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100;
    /// @dev Precision for the markup math.
    uint256 private constant WAD = 10 ** 18;

    /// @notice Role responsible for managing the SynapseExecutionService contract.
    /// Can set all the parameters defined in the SynapseExecutionServiceV1Storage struct.
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");
    /// @notice Role to track the Interchain Client contracts.
    /// Can request the execution of transactions by calling the requestTxExecution function.
    bytes32 public constant IC_CLIENT_ROLE = keccak256("IC_CLIENT_ROLE");

    constructor() {
        // Ensure that the implementation contract could not be initialized
        _disableInitializers();
    }

    /// @notice Initializes the SynapseExecutionService contract by setting the initial admin.
    /// @dev Needs to be called atomically after the proxy is deployed.
    function initialize(address admin) external initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /// @notice Sets the fraction of the accumulated fees to be paid to caller of `claimFees`.
    /// This encourages rational actors to call the function as soon as claim fee is higher than the gas cost.
    /// @dev Could be only called by the governor. The fraction could not exceed 1% (1e16).
    /// @param claimerFraction_     The fraction of the fees to be paid to the claimer (100% = 1e18)
    function setClaimerFraction(uint256 claimerFraction_) external virtual onlyRole(GOVERNOR_ROLE) {
        if (claimerFraction_ > MAX_CLAIMER_FRACTION) {
            revert ClaimableFees__ClaimerFractionAboveMax(claimerFraction_, MAX_CLAIMER_FRACTION);
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.claimerFraction = claimerFraction_;
        emit ClaimerFractionSet(claimerFraction_);
    }

    /// @notice Allows the contract governor to set the address of the EOA account that will be used
    /// to execute transactions on the remote chains. This address will also be used as the recipient
    /// of the execution fees collected by the contract.
    /// @dev Could be only called by the governor. Will revert if the zero address is passed.
    function setExecutorEOA(address executorEOA_) external virtual onlyRole(GOVERNOR_ROLE) {
        if (executorEOA_ == address(0)) {
            revert SynapseExecutionService__ExecutorZeroAddress();
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.executorEOA = executorEOA_;
        emit ExecutorEOASet(executorEOA_);
        emit FeeRecipientSet(executorEOA_);
    }

    /// @notice Allows the contract governor to set the address of the gas oracle. The gas oracle
    /// is used for estimating the gas cost of the transactions.
    /// @dev Could be only called by the governor. Will revert if the passed address is not a contract.
    function setGasOracle(address gasOracle_) external virtual onlyRole(GOVERNOR_ROLE) {
        if (gasOracle_.code.length == 0) {
            revert SynapseExecutionService__GasOracleNotContract(gasOracle_);
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.gasOracle = gasOracle_;
        emit GasOracleSet(gasOracle_);
    }

    /// @notice Allows the contract governor to set the global markup that the Execution Service charges
    /// on top of the GasOracle's gas cost estimates.
    /// Zero markup means that the Execution Service charges the exact gas cost estimated by the GasOracle.
    /// The markup is denominated in Wei, 1e18 being 100%.
    /// @dev Could be only called by the governor.
    function setGlobalMarkup(uint256 globalMarkup_) external virtual onlyRole(GOVERNOR_ROLE) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.globalMarkup = globalMarkup_;
        emit GlobalMarkupSet(globalMarkup_);
    }

    /// @notice Request the execution of an Interchain Transaction on a remote chain in exchange for
    /// the execution fee, attached to the transaction as `msg.value`.
    /// Note: the off-chain actor needs to fetch the transaction payload from the InterchainClient
    /// event with the same transactionId, then execute the transaction on the remote chain:
    /// `dstInterchainClient.executeTransaction(transactionPayload)`
    /// @dev Could only be called by `InterchainClient` contracts.
    /// Will revert if the execution fee is not big enough.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param txPayloadSize        The size of the transaction payload to use for the execution.
    /// @param transactionId        The id of the transaction to execute.
    /// @param options              The options to use for the execution.
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
            revert SynapseExecutionService__FeeAmountBelowMin({feeAmount: msg.value, minRequired: requiredFee});
        }
        emit ExecutionRequested({transactionId: transactionId, client: msg.sender, executionFee: msg.value});
    }

    /// @notice Get the execution fee for executing an Interchain Transaction on a remote chain.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param txPayloadSize        The size of the transaction payload to use for the execution.
    /// @param options              The options to use for the execution.
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
            revert SynapseExecutionService__GasOracleZeroAddress();
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

    /// @notice Address of the EOA account that will be used to execute transactions on the remote chains.
    function executorEOA() public view virtual returns (address) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.executorEOA;
    }

    /// @notice Address of the gas oracle used for estimating the gas cost of the transactions.
    function gasOracle() public view virtual returns (address) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.gasOracle;
    }

    /// @notice The markup that the Execution Service charges on top of the GasOracle's gas cost estimates.
    /// Zero markup means that the Execution Service charges the exact gas cost estimated by the GasOracle.
    /// The markup is denominated in Wei, 1e18 being 100%.
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
