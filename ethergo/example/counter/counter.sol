// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract Counter {
    // this is used for testing account impersonation
    address constant VITALIK = address(0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045);

    int private count = 0;
    int private vitaikCount = 0;

    function incrementCounter() public {
        count += 1;
    }
    function decrementCounter() public {
        count -= 1;
    }

    function vitalikIncrement() public {
        require(msg.sender == VITALIK, "Only Vitalik can count by 10");
        vitaikCount += 10;
    }

    function getCount() public view returns (int) {
        return count;
    }

    function getVitalikCount() public view returns (int) {
        return vitaikCount;
    }
}
