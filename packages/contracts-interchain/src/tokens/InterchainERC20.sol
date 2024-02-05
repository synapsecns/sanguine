// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {ICERC20} from "../interfaces/ICERC20.sol";
import {RateLimit} from "../libs/RateLimit.sol";

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

contract InterchainERC20 is ERC20, AccessControl, Pausable, ICERC20 {
    bytes32 public constant EMERGENCY_PAUSER_ROLE = keccak256("EMERGENCY_PAUSER_ROLE");
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    /// @notice Address of the processor contract, responsible for Token<>InterchainToken conversions
    /// Note: this should be left as zero address, if the InterchainERC20 is used as a "native token deployment"
    /// on the chain rather than an interchain representation of an existing token.
    address public immutable PROCESSOR;

    /// @dev Rate Limit for Bridge's burn operations
    mapping(address bridge => RateLimit) internal _burnLimits;
    /// @dev Rate Limit for Bridge's mint operations
    mapping(address bridge => RateLimit) internal _mintLimits;

    constructor(
        string memory name_,
        string memory symbol_,
        address initialAdmin_,
        address processor_
    )
        ERC20(name_, symbol_)
    {
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin_);
        PROCESSOR = processor_;
    }

    // ══════════════════════════════════════════════ ADMIN FUNCTIONS ══════════════════════════════════════════════════

    /// @notice Sets a new total burn limit for the given bridge. This also updates the current burn limit,
    /// based on the amount spent from the previous limit.
    /// Note: setting the limit to 0 effectively disables the bridge's burn operations.
    /// @dev Could only be called by the Governor.
    /// @param bridge   The address of the bridge
    /// @param limit    The new total limit
    function setTotalBurnLimit(address bridge, uint256 limit) external onlyRole(GOVERNOR_ROLE) {
        _burnLimits[bridge].setTotalLimit(limit);
        emit BurnLimitSet(bridge, limit);
    }

    /// @notice Sets a new total mint limit for the given bridge. This also updates the current mint limit,
    /// based on the amount spent from the previous limit.
    /// Note: setting the limit to 0 effectively disables the bridge's mint operations.
    /// @dev Could only be called by the Governor.
    /// @param bridge   The address of the bridge
    /// @param limit    The new total limit
    function setTotalMintLimit(address bridge, uint256 limit) external onlyRole(GOVERNOR_ROLE) {
        _mintLimits[bridge].setTotalLimit(limit);
        emit MintLimitSet(bridge, limit);
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
    function mint(address account, uint256 amount) external whenNotPaused {
        // Spend from the Bridge's mint limit (will revert if the limit is exceeded)
        _mintLimits[msg.sender].spendLimit(amount);
        _mint(account, amount);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ICERC20
    function getCurrentBurnLimit(address bridge) external view returns (uint256) {
        return _burnLimits[bridge].getCurrentLimit();
    }

    /// @inheritdoc ICERC20
    function getCurrentMintLimit(address bridge) external view returns (uint256) {
        return _mintLimits[bridge].getCurrentLimit();
    }

    /// @inheritdoc ICERC20
    function getTotalBurnLimit(address bridge) external view override returns (uint256) {
        return _burnLimits[bridge].getTotalLimit();
    }

    /// @inheritdoc ICERC20
    function getTotalMintLimit(address bridge) external view override returns (uint256) {
        return _mintLimits[bridge].getTotalLimit();
    }
}
