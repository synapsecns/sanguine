// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModule} from "./InterchainModule.sol";

import {SynapseModuleEvents} from "../events/SynapseModuleEvents.sol";
import {IGasOracle} from "../interfaces/IGasOracle.sol";
import {ISynapseModule} from "../interfaces/ISynapseModule.sol";

import {ThresholdECDSA} from "../libs/ThresholdECDSA.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract SynapseModule is InterchainModule, Ownable, SynapseModuleEvents, ISynapseModule {
    // TODO: make sure this is a good enough default value
    uint256 public constant DEFAULT_VERIFY_GAS_LIMIT = 100_000;

    uint256 internal constant MAX_CLAIM_FEE_FRACTION = 0.01e18; // 1%
    uint256 internal constant FEE_PRECISION = 1e18;

    /// @dev Struct to hold the verifiers and the threshold for the module.
    ThresholdECDSA internal _verifiers;
    /// @dev Claim fee fraction, 100% = 1e18
    uint256 internal _claimFeeFraction;
    /// @dev Gas limit for the verifyEntry function on the remote chain.
    mapping(uint256 chainId => uint256 gasLimit) internal _verifyGasLimit;

    /// @inheritdoc ISynapseModule
    address public feeCollector;
    /// @inheritdoc ISynapseModule
    address public gasOracle;

    constructor(address interchainDB, address owner_) InterchainModule(interchainDB) Ownable(owner_) {
        // ThresholdECDSA throws an explicit error if threshold is not set, so default value is not needed
    }

    // ═══════════════════════════════════════════════ PERMISSIONED ════════════════════════════════════════════════════

    /// @inheritdoc ISynapseModule
    function addVerifier(address verifier) external onlyOwner {
        _addVerifier(verifier);
    }

    /// @inheritdoc ISynapseModule
    function addVerifiers(address[] calldata verifiers) external onlyOwner {
        uint256 length = verifiers.length;
        for (uint256 i = 0; i < length; ++i) {
            _addVerifier(verifiers[i]);
        }
    }

    /// @inheritdoc ISynapseModule
    function removeVerifier(address verifier) external onlyOwner {
        _removeVerifier(verifier);
    }

    /// @inheritdoc ISynapseModule
    function removeVerifiers(address[] calldata verifiers) external onlyOwner {
        uint256 length = verifiers.length;
        for (uint256 i = 0; i < length; ++i) {
            _removeVerifier(verifiers[i]);
        }
    }

    /// @inheritdoc ISynapseModule
    function setThreshold(uint256 threshold) external onlyOwner {
        _verifiers.modifyThreshold(threshold);
        emit ThresholdChanged(threshold);
    }

    /// @inheritdoc ISynapseModule
    function setFeeCollector(address feeCollector_) external onlyOwner {
        feeCollector = feeCollector_;
        emit FeeCollectorChanged(feeCollector_);
    }

    /// @inheritdoc ISynapseModule
    function setClaimFeeFraction(uint256 claimFeeFraction) external onlyOwner {
        if (claimFeeFraction > MAX_CLAIM_FEE_FRACTION) {
            revert SynapseModule__ClaimFeeFractionExceedsMax(claimFeeFraction);
        }
        _claimFeeFraction = claimFeeFraction;
        emit ClaimFeeFractionChanged(claimFeeFraction);
    }

    /// @inheritdoc ISynapseModule
    function setGasOracle(address gasOracle_) external onlyOwner {
        if (gasOracle_.code.length == 0) {
            revert SynapseModule__GasOracleNotContract(gasOracle_);
        }
        gasOracle = gasOracle_;
        emit GasOracleChanged(gasOracle_);
    }

    /// @inheritdoc ISynapseModule
    function setVerifyGasLimit(uint256 chainId, uint256 gasLimit) external onlyOwner {
        _verifyGasLimit[chainId] = gasLimit;
        emit VerifyGasLimitChanged(chainId, gasLimit);
    }

    // ══════════════════════════════════════════════ PERMISSIONLESS ═══════════════════════════════════════════════════

    /// @inheritdoc ISynapseModule
    function claimFees() external {
        if (feeCollector == address(0)) {
            revert SynapseModule__FeeCollectorNotSet();
        }
        if (address(this).balance == 0) {
            revert SynapseModule__NoFeesToClaim();
        }
        uint256 claimFee = getClaimFeeAmount();
        uint256 collectedFee = address(this).balance - claimFee;
        Address.sendValue(payable(feeCollector), collectedFee);
        Address.sendValue(payable(msg.sender), claimFee);
        emit FeesClaimed(feeCollector, collectedFee, msg.sender, claimFee);
    }

    /// @inheritdoc ISynapseModule
    function verifyEntry(bytes calldata encodedEntry, bytes calldata signatures) external {
        bytes32 ethSignedHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedEntry));
        _verifiers.verifySignedHash(ethSignedHash, signatures);
        _verifyEntry(encodedEntry);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISynapseModule
    function getClaimFeeFraction() external view returns (uint256) {
        return _claimFeeFraction;
    }

    /// @inheritdoc ISynapseModule
    function getVerifiers() external view returns (address[] memory) {
        return _verifiers.getSigners();
    }

    /// @inheritdoc ISynapseModule
    function isVerifier(address account) external view returns (bool) {
        return _verifiers.isSigner(account);
    }

    /// @inheritdoc ISynapseModule
    function getClaimFeeAmount() public view returns (uint256) {
        return address(this).balance * _claimFeeFraction / FEE_PRECISION;
    }

    /// @inheritdoc ISynapseModule
    function getThreshold() public view returns (uint256) {
        return _verifiers.getThreshold();
    }

    /// @inheritdoc ISynapseModule
    function getVerifyGasLimit(uint256 chainId) public view override returns (uint256 gasLimit) {
        gasLimit = _verifyGasLimit[chainId];
        if (gasLimit == 0) {
            gasLimit = DEFAULT_VERIFY_GAS_LIMIT;
        }
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
            gasLimit: getVerifyGasLimit(destChainId),
            calldataSize: 292 + 64 * getThreshold()
        });
    }
}
