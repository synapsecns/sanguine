// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {AbstractProcessor} from "./AbstractProcessor.sol";

import {IBurnableGMX} from "../interfaces/IBurnableGMX.sol";
import {ICERC20} from "../interfaces/ICERC20.sol";

/// @notice BurningProcessor is a contract that enables the conversion between
/// the ERC20 token (underlying) and its ICERC20 counterpart (interchain) by using the mint-burn
/// mechanism.
/// - Interchain token is minted when the ERC20 token is burned.
/// - ERC20 token is minted when the Interchain token is burned.
/// See AbstractProcessor.sol for more details.
contract BurningProcessor is AbstractProcessor {
    constructor(
        address interchainToken_,
        address underlyingToken_
    )
        AbstractProcessor(interchainToken_, underlyingToken_)
    {}

    /// @dev Burns the ICERC20 token taken from `msg.sender`, then
    /// mints the same amount of the underlying token to `msg.sender`.
    function _burnInterchainToken(uint256 amount) internal override {
        ICERC20(INTERCHAIN_TOKEN).burn(amount);
        IBurnableGMX(UNDERLYING_TOKEN).mint(msg.sender, amount);
    }

    /// @dev Burns the underlying token taken from `msg.sender`, then
    /// mints the same amount of the ICERC20 token to `msg.sender`.
    function _mintInterchainToken(uint256 amount) internal override {
        IBurnableGMX(UNDERLYING_TOKEN).burn(address(this), amount);
        ICERC20(INTERCHAIN_TOKEN).mint(msg.sender, amount);
    }
}
