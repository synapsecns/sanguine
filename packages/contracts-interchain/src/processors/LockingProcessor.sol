// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {IDefaultPool} from "../interfaces/IDefaultPool.sol";
import {AbstractProcessor} from "./AbstractProcessor.sol";

/// @notice LockingProcessor is a contract that enables the conversion between
/// the ERC20 token (underlying) and its InterchainERC20 counterpart by using the lock-unlock
/// mechanism.
/// - Interchain token is minted when the ERC20 token is locked.
/// - ERC20 token is unlocked when the Interchain token is burned.
/// See AbstractProcessor.sol for more details.
contract LockingProcessor is AbstractProcessor {
    constructor(
        address interchainToken_,
        address underlyingToken_
    )
        AbstractProcessor(interchainToken_, underlyingToken_)
    {}

    /// @dev Burns the InterchainERC20 token taken from `msg.sender`, then
    /// transfers the same amount of the underlying token to `msg.sender`.
    function _burnInterchainToken(uint256 amount) internal override {
        // TODO: implement
    }

    /// @dev Handles the underlying token taken from `msg.sender`, then
    /// mints the same amount of the InterchainERC20 token to `msg.sender`.
    function _mintInterchainToken(uint256 amount) internal override {
        // TODO: implement
    }
}
