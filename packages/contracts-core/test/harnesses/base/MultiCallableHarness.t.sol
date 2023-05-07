// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MultiCallable} from "../../../contracts/base/MultiCallable.sol";

contract MultiCallableHarness is MultiCallable {
    error GmError();

    bytes32[] private buffer;

    mapping(uint256 => bool) private willRevert;

    function addUint(uint256 value) external returns (uint256) {
        if (willRevert[value]) revert GmError();
        buffer.push(bytes32(value));
        return value * 2;
    }

    function toggleRevert(uint256 value, bool revertFlag) external {
        willRevert[value] = revertFlag;
    }

    function getBuffer() external view returns (bytes32[] memory) {
        return buffer;
    }
}
