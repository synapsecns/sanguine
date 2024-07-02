// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract Counter {
    // this is used for testing account impersonation
    address constant VITALIK = address(0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045);

    event Incremented(int count);
    event Decremented(int count);
    event IncrementedByUser(address user, int count);

    int private count = 0;
    int private vitaikCount = 0;

    // @dev the block the contract was deployed at
    uint256 public immutable deployBlock;

    constructor()  {
        deployBlock = block.number;
    }


    function incrementCounter() public {
        count += 1;
        emit Incremented(count);
    }
    function decrementCounter() public {
        count -= 1;
        emit Decremented(count);
    }

    function vitalikIncrement() public {
        require(msg.sender == VITALIK, "Only Vitalik can count by 10");
        vitaikCount += 10;
        emit IncrementedByUser(msg.sender, vitaikCount);
    }

    function getCount() public view returns (int) {
        return count;
    }

    function getVitalikCount() public view returns (int) {
        return vitaikCount;
    }
}
