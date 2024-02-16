// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModule} from "./InterchainModule.sol";

import {IThresholdECDSAModule} from "../interfaces/IThresholdECDSAModule.sol";

import {InterchainEntry} from "../libs/InterchainEntry.sol";
import {ThresholdECDSA} from "../libs/ThresholdECDSA.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract ThresholdECDSAModule is InterchainModule, Ownable, IThresholdECDSAModule {
    /// @inheritdoc IThresholdECDSAModule
    address public gasOracle;

    constructor(address interchainDB, address initialOwner) InterchainModule(interchainDB) Ownable(initialOwner) {}

    // ═══════════════════════════════════════════════ PERMISSIONED ════════════════════════════════════════════════════

    /// @inheritdoc IThresholdECDSAModule
    function addVerifier(address verifier) external onlyOwner {}

    /// @inheritdoc IThresholdECDSAModule
    function removeVerifier(address verifier) external onlyOwner {}

    /// @inheritdoc IThresholdECDSAModule
    function setThreshold(uint256 threshold) external onlyOwner {}

    /// @inheritdoc IThresholdECDSAModule
    function setGasOracle(address gasOracle_) external onlyOwner {}

    // ══════════════════════════════════════════════ PERMISSIONLESS ═══════════════════════════════════════════════════

    /// @inheritdoc IThresholdECDSAModule
    function verifyEntry(bytes calldata encodedEntry, bytes[] calldata signatures) external {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IThresholdECDSAModule
    function getThreshold() external view returns (uint256) {}

    /// @inheritdoc IThresholdECDSAModule
    function getVerifiers() external view returns (address[] memory) {}

    /// @inheritdoc IThresholdECDSAModule
    function isVerifier(address account) external view returns (bool) {}

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Internal logic to request the verification of an entry on the destination chain.
    function _requestVerification(uint256 destChainId, bytes memory encodedEntry) internal override {}

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Internal logic to get the module fee for verifying an entry on the specified destination chain.
    function _getModuleFee(uint256 destChainId) internal view override returns (uint256) {}
}
