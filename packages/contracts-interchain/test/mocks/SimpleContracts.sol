// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract OneArgContract {
    uint256 public value;

    constructor(uint256 value_) payable {
        value = value_;
    }

    function testSimpleContract() external pure {
        // This function is only used to remove SimpleContract from coverage reports
    }
}

contract RevertingContract {
    error GM();

    constructor() payable {
        revert GM();
    }
}
