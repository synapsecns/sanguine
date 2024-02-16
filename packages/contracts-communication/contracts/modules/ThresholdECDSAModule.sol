// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModule} from "./InterchainModule.sol";

import {ThresholdECDSAModuleEvents} from "../events/ThresholdECDSAModuleEvents.sol";
import {IGasOracle} from "../interfaces/IGasOracle.sol";
import {IThresholdECDSAModule} from "../interfaces/IThresholdECDSAModule.sol";

import {InterchainEntry} from "../libs/InterchainEntry.sol";
import {ThresholdECDSA} from "../libs/ThresholdECDSA.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

contract ThresholdECDSAModule is InterchainModule, Ownable, ThresholdECDSAModuleEvents, IThresholdECDSAModule {
    uint256 public constant VERIFY_GAS_LIMIT = 100_000;

    /// @dev Struct to hold the verifiers and the threshold for the module.
    ThresholdECDSA internal _verifiers;

    /// @inheritdoc IThresholdECDSAModule
    address public feeCollector;

    /// @inheritdoc IThresholdECDSAModule
    address public gasOracle;

    constructor(address interchainDB, address initialOwner) InterchainModule(interchainDB) Ownable(initialOwner) {
        // Explicitly disable the module functionality until the threshold is set.
        _setThreshold(type(uint256).max);
    }

    // ═══════════════════════════════════════════════ PERMISSIONED ════════════════════════════════════════════════════

    /// @inheritdoc IThresholdECDSAModule
    function addVerifier(address verifier) external onlyOwner {
        _verifiers.addSigner(verifier);
        emit VerifierAdded(verifier);
    }

    /// @inheritdoc IThresholdECDSAModule
    function removeVerifier(address verifier) external onlyOwner {
        _verifiers.removeSigner(verifier);
        emit VerifierRemoved(verifier);
    }

    /// @inheritdoc IThresholdECDSAModule
    function setThreshold(uint256 threshold) external onlyOwner {
        _setThreshold(threshold);
    }

    /// @inheritdoc IThresholdECDSAModule
    function setFeeCollector(address feeCollector_) external onlyOwner {
        _setFeeCollector(feeCollector_);
    }

    /// @inheritdoc IThresholdECDSAModule
    function setGasOracle(address gasOracle_) external onlyOwner {
        if (gasOracle_.code.length == 0) {
            revert ThresholdECDSAModule__GasOracleNotContract(gasOracle_);
        }
        gasOracle = gasOracle_;
        emit GasOracleChanged(gasOracle_);
    }

    // ══════════════════════════════════════════════ PERMISSIONLESS ═══════════════════════════════════════════════════

    /// @inheritdoc IThresholdECDSAModule
    function verifyEntry(bytes calldata encodedEntry, bytes calldata signatures) external {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IThresholdECDSAModule
    function getVerifiers() external view returns (address[] memory) {
        return _verifiers.getSigners();
    }

    /// @inheritdoc IThresholdECDSAModule
    function isVerifier(address account) external view returns (bool) {
        return _verifiers.isSigner(account);
    }

    /// @inheritdoc IThresholdECDSAModule
    function getThreshold() public view returns (uint256) {
        return _verifiers.getThreshold();
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Sets the threshold for the module. Permissions should be checked in the calling function.
    function _setThreshold(uint256 threshold) internal {
        _verifiers.modifyThreshold(threshold);
        emit ThresholdChanged(threshold);
    }

    /// @dev Internal logic to set the address of the fee collector.
    /// Permissions should be checked in the calling function.
    function _setFeeCollector(address feeCollector_) internal {
        feeCollector = feeCollector_;
        emit FeeCollectorChanged(feeCollector_);
    }

    /// @dev Internal logic to request the verification of an entry on the destination chain.
    function _requestVerification(
        uint256, // destChainId
        bytes memory // encodedEntry
    )
        internal
        override
    {
        // All the hark work has been done in InterchainModule.requestVerification
        Address.sendValue(payable(feeCollector), msg.value);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Internal logic to get the module fee for verifying an entry on the specified destination chain.
    function _getModuleFee(uint256 destChainId) internal view override returns (uint256) {
        // On the remote chain the verifyEntry(entry, signatures) function will be called.
        // We need to figure out the calldata size for the remote call.
        // selector (4 bytes) + entry + signatures
        // entry is 32 (length) + 32*4 (fields) = 160
        // signatures: 32 (length) + 65*threshold (padded up to be a multiple of 32 bytes)
        // Total formula is: 4 + 32 (entry offset) + 32 (signatures offset) + 160 + 32
        return IGasOracle(gasOracle).estimateTxCostInLocalUnits({
            remoteChainId: destChainId,
            gasLimit: VERIFY_GAS_LIMIT,
            calldataSize: 260 + (2 * getThreshold() + 1) * 32
        });
    }
}
