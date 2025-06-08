// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IBurnableToken} from "../../src/interfaces/IBurnableToken.sol";

import {ERC20, ERC20Burnable} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

// solhint-disable no-empty-blocks
/// @notice Obviously, do NOT use this token in production. It's only for testing purposes.
contract TestToken is ERC20Burnable, IBurnableToken {
    constructor() ERC20("TestTokenMock", "TLM") {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testTestToken() external {}

    function mintTestTokens(address to, uint256 amount) public {
        _mint(to, amount);
    }

    function burn(uint256 amount) public override(ERC20Burnable, IBurnableToken) {
        super.burn(amount);
    }

    function burnFrom(address from, uint256 amount) public override(ERC20Burnable, IBurnableToken) {
        super.burnFrom(from, amount);
    }
}
