// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModule} from "./InterchainModule.sol";

import {SynapseModuleEvents} from "../events/SynapseModuleEvents.sol";
import {ISynapseGasOracle} from "../interfaces/ISynapseGasOracle.sol";
import {ISynapseModule} from "../interfaces/ISynapseModule.sol";
import {ThresholdECDSA} from "../libs/ThresholdECDSA.sol";

import {ClaimableFees} from "../fees/ClaimableFees.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract SynapseModule is InterchainModule, ClaimableFees, Ownable, SynapseModuleEvents, ISynapseModule {
    // TODO: make sure this is a good enough default value
    uint256 public constant DEFAULT_VERIFY_GAS_LIMIT = 100_000;

    /// @dev Struct to hold the verifiers and the threshold for the module.
    ThresholdECDSA internal _verifiers;

    /// @dev Gas limit for the verifyBatch function on the remote chain.
    mapping(uint64 chainId => uint256 gasLimit) internal _verifyGasLimit;
    /// @dev Hash of the last gas data sent to the remote chain.
    mapping(uint64 chainId => bytes32 gasDataHash) internal _lastGasDataHash;
    /// @dev Nonce of the last gas data received from the remote chain.
    mapping(uint64 chainId => uint64 gasDataNonce) internal _lastGasDataNonce;

    /// @dev Fraction of the fees to be paid to the claimer (100% = 1e18).
    uint256 internal _claimerFraction;
    /// @dev Recipient of the fees collected by the module.
    address internal _feeRecipient;

    /// @notice Address of the gas oracle used for estimating the verification fees.
    address public gasOracle;

    constructor(address interchainDB, address owner_) InterchainModule(interchainDB) Ownable(owner_) {
        // ThresholdECDSA throws an explicit error if threshold is not set, so default value is not needed
    }

    // ═══════════════════════════════════════════════ PERMISSIONED ════════════════════════════════════════════════════

    /// @notice Adds a new verifier to the module.
    /// @dev Could be only called by the owner. Will revert if the verifier is already added.
    function addVerifier(address verifier) external onlyOwner {
        _addVerifier(verifier);
    }

    /// @notice Adds a list of new verifiers to the module.
    /// @dev Could be only called by the owner. Will revert if any of the verifiers is already added.
    function addVerifiers(address[] calldata verifiers) external onlyOwner {
        uint256 length = verifiers.length;
        for (uint256 i = 0; i < length; ++i) {
            _addVerifier(verifiers[i]);
        }
    }

    /// @notice Removes a verifier from the module.
    /// @dev Could be only called by the owner. Will revert if the verifier is not added.
    function removeVerifier(address verifier) external onlyOwner {
        _removeVerifier(verifier);
    }

    /// @notice Removes a list of verifiers from the module.
    /// @dev Could be only called by the owner. Will revert if any of the verifiers is not added.
    function removeVerifiers(address[] calldata verifiers) external onlyOwner {
        uint256 length = verifiers.length;
        for (uint256 i = 0; i < length; ++i) {
            _removeVerifier(verifiers[i]);
        }
    }

    /// @notice Sets the threshold of the module.
    /// @dev Could be only called by the owner. Will revert if the threshold is zero.
    function setThreshold(uint256 threshold) external onlyOwner {
        _verifiers.modifyThreshold(threshold);
        emit ThresholdSet(threshold);
    }

    /// @notice Sets the address of the fee collector, which will have the verification fees forwarded to it.
    /// @dev Could be only called by the owner.
    function setFeeRecipient(address feeRecipient) external onlyOwner {
        if (feeRecipient == address(0)) {
            revert SynapseModule__FeeRecipientZeroAddress();
        }
        _feeRecipient = feeRecipient;
        emit FeeRecipientSet(feeRecipient);
    }

    /// @notice Sets the fraction of the accumulated fees to be paid to caller of `claimFees`.
    /// This encourages rational actors to call the function as soon as claim fee is higher than the gas cost.
    /// @dev Could be only called by the owner. Could not exceed 1% (1e16).
    function setClaimerFraction(uint256 claimerFraction) external onlyOwner {
        if (claimerFraction > MAX_CLAIMER_FRACTION) {
            revert ClaimableFees__ClaimerFractionAboveMax(claimerFraction, MAX_CLAIMER_FRACTION);
        }
        _claimerFraction = claimerFraction;
        emit ClaimerFractionSet(claimerFraction);
    }

    /// @notice Sets the address of the gas oracle to be used for estimating the verification fees.
    /// @dev Could be only called by the owner. Will revert if the gas oracle is not a contract.
    function setGasOracle(address gasOracle_) external onlyOwner {
        if (gasOracle_.code.length == 0) {
            revert SynapseModule__GasOracleNotContract(gasOracle_);
        }
        gasOracle = gasOracle_;
        emit GasOracleSet(gasOracle_);
    }

    /// @notice Sets the estimated gas limit for verifying a batch on the given chain.
    /// @dev Could be only called by the owner.
    /// @param chainId      The chain ID for which to set the gas limit
    /// @param gasLimit     The new gas limit for the verification on the specified chain
    function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) external onlyOwner {
        _verifyGasLimit[chainId] = gasLimit;
        emit VerifyGasLimitSet(chainId, gasLimit);
    }

    // ══════════════════════════════════════════════ PERMISSIONLESS ═══════════════════════════════════════════════════

    /// @notice Verifies a batch from the remote chain using a set of verifier signatures.
    /// If the threshold is met, the batch will be marked as verified in the Interchain DataBase.
    /// @dev List of recovered signers from the signatures must be sorted in the ascending order.
    /// @param encodedBatch The encoded batch to verify
    /// @param signatures   Signatures used to verify the batch, concatenated
    function verifyRemoteBatch(bytes calldata encodedBatch, bytes calldata signatures) external {
        bytes32 ethSignedHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedBatch));
        _verifiers.verifySignedHash(ethSignedHash, signatures);
        _verifyBatch(encodedBatch);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the list of verifiers for the module.
    function getVerifiers() external view returns (address[] memory) {
        return _verifiers.getSigners();
    }

    /// @notice Checks if the given account is a verifier for the module.
    function isVerifier(address account) external view returns (bool) {
        return _verifiers.isSigner(account);
    }

    /// @notice Gets the threshold of the module. This is the minimum number of signatures required for verification.
    function getThreshold() public view returns (uint256) {
        return _verifiers.getThreshold();
    }

    /// @notice Returns the estimated gas limit for verifying a batch on the given chain.
    /// Note: this defaults to DEFAULT_VERIFY_GAS_LIMIT if not set.
    function getVerifyGasLimit(uint64 chainId) public view override returns (uint256 gasLimit) {
        gasLimit = _verifyGasLimit[chainId];
        if (gasLimit == 0) {
            gasLimit = DEFAULT_VERIFY_GAS_LIMIT;
        }
    }

    /// @notice Returns the amount of fees that can be claimed.
    function getClaimableAmount() public view override returns (uint256) {
        return address(this).balance;
    }

    /// @notice Returns the fraction of the fees that the claimer will receive.
    /// The result is in the range [0, 1e18], where 1e18 is 100%.
    function getClaimerFraction() public view override returns (uint256) {
        return _claimerFraction;
    }

    /// @notice Returns the address that will receive the claimed fees.
    function getFeeRecipient() public view override returns (address) {
        return _feeRecipient;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Adds a verifier to the module. Permissions should be checked in the calling function.
    function _addVerifier(address verifier) internal {
        _verifiers.addSigner(verifier);
        emit VerifierAdded(verifier);
    }

    /// @dev Removes a verifier from the module. Permissions should be checked in the calling function.
    function _removeVerifier(address verifier) internal {
        _verifiers.removeSigner(verifier);
        emit VerifierRemoved(verifier);
    }

    /// @dev Hook that is called before the fees are claimed.
    /// Useful if the inheriting contract needs to manage the state when the fees are claimed.
    // solhint-disable-next-line no-empty-blocks
    function _beforeFeesClaimed(uint256, uint256) internal override {
        // No op, as the claimable amount is tracked as the contract balance
    }

    /// @dev Internal logic to fill the module data for the specified destination chain.
    function _fillModuleData(
        uint64 dstChainId,
        uint64 // dbNonce
    )
        internal
        override
        returns (bytes memory moduleData)
    {
        moduleData = _getSynapseGasOracle().getLocalGasData();
        // Exit early if data is empty
        if (moduleData.length == 0) {
            return moduleData;
        }
        bytes32 dataHash = keccak256(moduleData);
        // Don't send the same data twice
        if (dataHash == _lastGasDataHash[dstChainId]) {
            moduleData = "";
        } else {
            _lastGasDataHash[dstChainId] = dataHash;
            emit GasDataSent(dstChainId, moduleData);
        }
    }

    /// @dev Internal logic to handle the auxiliary module data relayed from the remote chain.
    function _receiveModuleData(uint64 srcChainId, uint64 dbNonce, bytes memory moduleData) internal override {
        // Exit early if data is empty
        if (moduleData.length == 0) {
            return;
        }
        // Don't process outdated data
        uint64 lastNonce = _lastGasDataNonce[srcChainId];
        if (lastNonce == 0 || lastNonce < dbNonce) {
            _lastGasDataNonce[srcChainId] = dbNonce;
            _getSynapseGasOracle().receiveRemoteGasData(srcChainId, moduleData);
            emit GasDataReceived(srcChainId, moduleData);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Internal logic to get the module fee for verifying an batch on the specified destination chain.
    function _getModuleFee(
        uint64 dstChainId,
        uint64 // dbNonce
    )
        internal
        view
        override
        returns (uint256)
    {
        // On the remote chain the verifyRemoteBatch(batch, signatures) function will be called.
        // We need to figure out the calldata size for the remote call.
        // selector (4 bytes) + batch + signatures
        // batch is 32 (length) + 32*3 (fields) = 128
        // signatures: 32 (length) + 65*threshold (padded up to be a multiple of 32 bytes)
        // Total formula is: 4 + 32 (batch offset) + 32 (signatures offset) + 128 + 32
        return _getSynapseGasOracle().estimateTxCostInLocalUnits({
            remoteChainId: dstChainId,
            gasLimit: getVerifyGasLimit(dstChainId),
            calldataSize: 260 + 64 * getThreshold()
        });
    }

    /// @dev Internal logic to get the Synapse Gas Oracle. Reverts if the gas oracle is not set.
    function _getSynapseGasOracle() internal view returns (ISynapseGasOracle synapseGasOracle) {
        synapseGasOracle = ISynapseGasOracle(gasOracle);
        if (address(synapseGasOracle) == address(0)) {
            revert SynapseModule__GasOracleZeroAddress();
        }
    }
}
