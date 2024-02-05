// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20, RateLimit} from "../../src/tokens/InterchainERC20.sol";

// solhint-disable func-name-mixedcase
contract InterchainERC20Harness is InterchainERC20 {
    constructor(
        string memory name_,
        string memory symbol_,
        address initialAdmin_,
        address processor_
    )
        InterchainERC20(name_, symbol_, initialAdmin_, processor_)
    {}

    function exposed__setBurnRateLimit(address bridge, RateLimit memory limit) external {
        _burnLimits[bridge] = limit;
    }

    function exposed__setMintRateLimit(address bridge, RateLimit memory limit) external {
        _mintLimits[bridge] = limit;
    }

    function exposed__getBurnRateLimit(address bridge) external view returns (RateLimit memory) {
        return _burnLimits[bridge];
    }

    function exposed__getMintRateLimit(address bridge) external view returns (RateLimit memory) {
        return _mintLimits[bridge];
    }
}
