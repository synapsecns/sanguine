// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainModule} from "./IInterchainModule.sol";

interface ISynapseModule is IInterchainModule {
    error SynapseModule__GasOracleNotContract(address gasOracle);
    error SynapseModule__GasOracleZeroAddress();
    error SynapseModule__FeeRecipientZeroAddress();

    function addVerifier(address verifier) external;
    function addVerifiers(address[] calldata verifiers) external;
    function removeVerifier(address verifier) external;
    function removeVerifiers(address[] calldata verifiers) external;
    function setThreshold(uint256 threshold) external;

    function setFeeRecipient(address feeRecipient) external;
    function setClaimerFraction(uint256 claimerFraction) external;
    function setGasOracle(address gasOracle_) external;
    function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) external;

    function verifyRemoteBatch(bytes calldata encodedBatch, bytes calldata signatures) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function gasOracle() external view returns (address);
    function getVerifiers() external view returns (address[] memory);
    function getThreshold() external view returns (uint256);
    function isVerifier(address account) external view returns (bool);
    function getVerifyGasLimit(uint64 chainId) external view returns (uint256);
}
