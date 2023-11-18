// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract A {
    uint256 public value;

    constructor() {
        value = 1337;
    }
}

contract B {
    uint256 public value;

    constructor(uint256 value_) {
        value = value_;
    }
}

contract UnDeployable {
    constructor() {
        revert("GM, I don't feel so good...");
    }
}
