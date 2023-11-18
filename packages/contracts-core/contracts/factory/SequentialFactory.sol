// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {ISequentialFactory} from "../interfaces/ISequentialFactory.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract SequentialFactory is Ownable, ISequentialFactory {
    error SequentialFactory__DeploymentFailed(uint256 nonce);
    error SequentialFactory__InvalidNonce(uint256 nonce, uint256 expectedNonce);

    /// @dev Tracks the nonce of the factory. This is required because address(this).nonce does not exist.
    uint256 internal _nonce;

    constructor(address owner_) {
        transferOwnership(owner_);
        // Nonce of the smart contract starts at 1, see EIP-161
        // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
        _nonce = 1;
    }

    /// @inheritdoc ISequentialFactory
    function deploy(uint256 nonce, bytes memory code) external onlyOwner returns (address deployedAt) {
        uint256 expectedNonce = getNonce();
        if (nonce != expectedNonce) revert SequentialFactory__InvalidNonce(nonce, expectedNonce);
        ++_nonce;
        // Use assembly to deploy the contract
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Add 0x20 to skip the length field of code
            deployedAt := create(0, add(code, 0x20), mload(code))
        }
        // CREATE opcode returns 0 on deployment failure
        if (deployedAt == address(0)) revert SequentialFactory__DeploymentFailed(nonce);
    }

    /// @inheritdoc ISequentialFactory
    function getNonce() public view returns (uint256 nonce) {
        return _nonce;
    }

    /// @inheritdoc ISequentialFactory
    function predictDeployment(uint256 nonce) public view returns (address deployedAt) {
        // TODO: complete
    }

    /// @inheritdoc ISequentialFactory
    function predictNextDeployment() external view returns (address deployedAt) {
        return predictDeployment(getNonce());
    }
}
