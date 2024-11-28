// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

// solhint-disable no-empty-blocks
/// @notice WETH mock for testing purposes. DO NOT USE IN PRODUCTION.
contract WETHMock is ERC20 {
    constructor() ERC20("Mock Wrapped Ether", "Mock WETH") {}

    receive() external payable {
        deposit();
    }

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testWETHMock() external {}

    function withdraw(uint256 amount) external {
        _burn(msg.sender, amount);
        Address.sendValue(payable(msg.sender), amount);
    }

    function deposit() public payable {
        _mint(msg.sender, msg.value);
    }
}
