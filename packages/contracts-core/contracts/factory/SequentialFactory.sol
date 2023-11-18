// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {ISequentialFactory} from "../interfaces/ISequentialFactory.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SequentialFactory is Ownable, ISequentialFactory {
    error SequentialFactory__DeploymentFailed(uint256 nonce);
    error SequentialFactory__InvalidNonce(uint256 nonce, uint256 expectedNonce);

    constructor(address owner_) {
        transferOwnership(owner_);
    }

    /// @inheritdoc ISequentialFactory
    function deploy(uint256 nonce, bytes memory code) external onlyOwner returns (address deployedAt) {}

    /// @inheritdoc ISequentialFactory
    function getNonce() external view returns (uint256 nonce) {}

    /// @inheritdoc ISequentialFactory
    function predictDeployment(uint256 nonce) external view returns (address deployedAt) {}

    /// @inheritdoc ISequentialFactory
    function predictNextDeployment() external view returns (address deployedAt) {}
}
