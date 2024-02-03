// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {IDefaultPool} from "../interfaces/IDefaultPool.sol";
import {AbstractProcessor} from "./AbstractProcessor.sol";

/// @notice BurningProcessor is a contract that enables the conversion between
/// the ERC20 token (underlying) and its InterchainERC20 counterpart by using the mint-burn
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
}
