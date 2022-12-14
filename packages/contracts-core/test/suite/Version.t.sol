// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";
import "../../contracts/Version.sol";

contract VersionImpl is Versioned {
    // solhint-disable-next-line no-empty-blocks
    constructor(string memory _version) Versioned(_version) {}
}

contract VersionTest is Test {
    function testVersion(string memory version) public {
        vm.assume(bytes(version).length <= 32);
        Versioned v = new VersionImpl(version);
        assertEq(v.version(), version, "Version mismatch");
    }
}
