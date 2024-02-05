// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {ICERC20} from "../interfaces/ICERC20.sol";

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

contract InterchainERC20 is ERC20, AccessControl, Pausable, ICERC20 {
    bytes32 public constant EMERGENCY_PAUSER_ROLE = keccak256("EMERGENCY_PAUSER_ROLE");
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    constructor(string memory name_, string memory symbol_, address initialAdmin_) ERC20(name_, symbol_) {
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin_);
    }

    // ══════════════════════════════════════════════ ADMIN FUNCTIONS ══════════════════════════════════════════════════

    function setTotalBurnLimit(address bridge, uint256 limit) external onlyRole(GOVERNOR_ROLE) {
        // TODO: Implement
    }

    function setTotalMintLimit(address bridge, uint256 limit) external onlyRole(GOVERNOR_ROLE) {
        // TODO: Implement
    }

    function pause() external onlyRole(EMERGENCY_PAUSER_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(EMERGENCY_PAUSER_ROLE) {
        _unpause();
    }

    // ══════════════════════════════════════════════ TOKEN FUNCTIONS ══════════════════════════════════════════════════

    /// @inheritdoc ICERC20
    function burn(uint256 amount) external {
        // TODO: Implement
    }

    /// @inheritdoc ICERC20
    function burnFrom(address account, uint256 amount) external {
        // TODO: Implement
    }

    /// @inheritdoc ICERC20
    function mint(address account, uint256 amount) external {
        // TODO: Implement
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ICERC20
    function getCurrentBurnLimit(address bridge) external view returns (uint256) {
        // TODO: Implement
    }

    /// @inheritdoc ICERC20
    function getCurrentMintLimit(address bridge) external view returns (uint256) {
        // TODO: Implement
    }

    /// @inheritdoc ICERC20
    function getTotalBurnLimit(address bridge) external view override returns (uint256) {
        // TODO: Implement
    }

    /// @inheritdoc ICERC20
    function getTotalMintLimit(address bridge) external view override returns (uint256) {
        // TODO: Implement
    }
}
