// SPDX-License-Identifier: UNLICENSED
// solhint-disable one-contract-per-file
pragma solidity ^0.8.0;

contract OneArgContract {
    uint256 public value;

    constructor(uint256 value_) payable {
        value = value_;
    }

    // solhint-disable-next-line no-empty-blocks
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
