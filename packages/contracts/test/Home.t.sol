// SPDX-License-Identifier: MIT

import "forge-std/Test.sol";
import "./utils/Utilities.sol";

import "../contracts/Home.sol";

import "../contracts/UpdaterManager.sol";

pragma solidity 0.8.13;

contract HomeTest is Test {
    address[] users = createUsers(10);
    Home home;
    address signer;
    address fakeSigner;
    address updater;
    address fakeUpdater;
    UpdaterManager updaterManager;

    function setUp() public {

    }

    function testSetup() public {
        assertEq(true, true);
    }
}