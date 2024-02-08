// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {ICERC20} from "../interfaces/ICERC20.sol";
import {IInterchainFactory} from "../interfaces/IInterchainFactory.sol";
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

    /// @dev Number of decimals for the token, setup during the contract deployment
    uint8 internal immutable _DECIMALS;

    /// @dev Rate Limit for Bridge's burn operations
    mapping(address bridge => RateLimit) internal _burnLimits;
    /// @dev Rate Limit for Bridge's mint operations
    mapping(address bridge => RateLimit) internal _mintLimits;

    error InterchainERC20__AdminZero();

    constructor(string memory name_, string memory symbol_, uint8 decimals_) ERC20(name_, symbol_) {
        (address initialAdmin, address processor) = IInterchainFactory(msg.sender).getInterchainTokenDeployParameters();
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin);
        _DECIMALS = decimals_;
        PROCESSOR = processor;
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
    function burn(uint256 amount) external whenNotPaused {
        // Processor is allowed to burn tokens without any rate limits
        if (msg.sender != PROCESSOR) {
            // Spend from the Bridge's burn limit (will revert if the limit is exceeded)
            _burnLimits[msg.sender].spendLimit(amount);
        }
        _burn(msg.sender, amount);
    }

    /// @inheritdoc ICERC20
    function burnFrom(address account, uint256 amount) external whenNotPaused {
        // Spend token transfer allowance regardless of the caller
        _spendAllowance(account, msg.sender, amount);
        // Processor is allowed to burn tokens without any rate limits
        if (msg.sender != PROCESSOR) {
            // Spend from the Bridge's burn limit (will revert if the limit is exceeded)
            _burnLimits[msg.sender].spendLimit(amount);
        }
        _burn(account, amount);
    }

    /// @inheritdoc ICERC20
    function mint(address account, uint256 amount) external whenNotPaused {
        // Processor is allowed to mint tokens without any rate limits
        if (msg.sender != PROCESSOR) {
            // Spend from the Bridge's mint limit (will revert if the limit is exceeded)
            _mintLimits[msg.sender].spendLimit(amount);
        }
        _mint(account, amount);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ICERC20
    function getCurrentBurnLimit(address burner) external view returns (uint256) {
        return burner == PROCESSOR ? type(uint256).max : _burnLimits[burner].getCurrentLimit();
    }

    /// @inheritdoc ICERC20
    function getCurrentMintLimit(address minter) external view returns (uint256) {
        return minter == PROCESSOR ? type(uint256).max : _mintLimits[minter].getCurrentLimit();
    }

    /// @inheritdoc ICERC20
    function getTotalBurnLimit(address burner) external view override returns (uint256) {
        return burner == PROCESSOR ? type(uint256).max : _burnLimits[burner].getTotalLimit();
    }

    /// @inheritdoc ICERC20
    function getTotalMintLimit(address minter) external view override returns (uint256) {
        return minter == PROCESSOR ? type(uint256).max : _mintLimits[minter].getTotalLimit();
    }

    /// @inheritdoc ERC20
    function decimals() public view override returns (uint8) {
        return _DECIMALS;
    }
}
