// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";
import {Versioned} from "../../contracts/base/Version.sol";

contract VersionImpl is Versioned {
    // solhint-disable-next-line no-empty-blocks
    constructor(string memory version_) Versioned(version_) {}
}

contract VersionTest is Test {
    function testVersion(string memory version) public {
        vm.assume(bytes(version).length <= 32);
        Versioned v = new VersionImpl(version);
        assertEq(v.version(), version, "Version mismatch");
    }
}
