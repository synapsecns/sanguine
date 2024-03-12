// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {ICERC20} from "../../src/interfaces/ICERC20.sol";
import {MockERC20} from "./MockERC20.sol";

/// @notice MockInterchainERC20 is a mock ERC20 token that follows the ICERC20 interface
/// @dev Make sure to set the burn and mint limits for Bridge, and set infinite limits for Processor
/// in the tests.
contract MockInterchainERC20 is MockERC20, ICERC20 {
    mapping(address => uint256) internal _burnLimits;
    mapping(address => uint256) internal _mintLimits;

    constructor(string memory name_) MockERC20(name_) {}

    // solhint-disable custom-errors
    function burn(uint256 amount) external {
        _spendLimit(_burnLimits, msg.sender, amount, "Burn limit exceeded");
        _burn(msg.sender, amount);
    }

    function burnFrom(address account, uint256 amount) external {
        _spendAllowance(account, msg.sender, amount);
        _spendLimit(_burnLimits, msg.sender, amount, "Burn limit exceeded");
        _burn(account, amount);
    }

    function mint(address account, uint256 amount) external {
        _spendLimit(_mintLimits, msg.sender, amount, "Mint limit exceeded");
        _mint(account, amount);
    }

    function setBurnLimit(address bridge, uint256 limit) external {
        _burnLimits[bridge] = limit;
    }

    function setMintLimit(address bridge, uint256 limit) external {
        _mintLimits[bridge] = limit;
    }

    function getCurrentBurnLimit(address bridge) external view returns (uint256) {
        return _burnLimits[bridge];
    }

    function getCurrentMintLimit(address bridge) external view returns (uint256) {
        return _mintLimits[bridge];
    }

    function getTotalBurnLimit(address bridge) external view override returns (uint256) {
        return _burnLimits[bridge];
    }

    function getTotalMintLimit(address bridge) external view override returns (uint256) {
        return _mintLimits[bridge];
    }

    // solhint-disable-next-line no-empty-blocks
    function testMockInterchainERC20() external pure {
        // This function is only used to remove MockInterchainERC20 from coverage reports
    }

    function _spendLimit(
        mapping(address => uint256) storage limits,
        address bridge,
        uint256 amount,
        string memory errorMsg
    )
        internal
    {
        require(limits[bridge] >= amount, errorMsg);
        limits[bridge] -= amount;
    }
}
