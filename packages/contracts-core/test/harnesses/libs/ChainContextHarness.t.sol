// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ChainContext} from "../../../contracts/libs/ChainContext.sol";

contract ChainContextHarness {
    function blockNumber() public view returns (uint40) {
        return ChainContext.blockNumber();
    }

    function blockTimestamp() public view returns (uint40) {
        return ChainContext.blockTimestamp();
    }

    function chainId() public view returns (uint32) {
        return ChainContext.chainId();
    }
}
