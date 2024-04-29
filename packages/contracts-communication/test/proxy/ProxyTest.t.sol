// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Test} from "forge-std/Test.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract ProxyTest is Test {
    bytes32 public constant ADMIN_SLOT = bytes32(uint256(keccak256("eip1967.proxy.admin")) - 1);

    function deployProxy(address implementation) internal returns (address proxy) {
        proxy = deployProxy(implementation, "");
    }

    function deployProxy(address implementation, bytes memory initData) internal returns (address proxy) {
        proxy = address(new TransparentUpgradeableProxy(implementation, address(this), initData));
    }

    function expectRevertInvalidInitialization() internal {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
    }

    function assertStorageAddress(address target, bytes32 slot, address expected) internal {
        bytes32 actual = vm.load(target, slot);
        bytes32 expectedBytes32 = bytes32(uint256(uint160(expected)));
        assertEq(actual, expectedBytes32);
    }

    function assertStorageUint(address target, bytes32 slot, uint256 expected) internal {
        bytes32 actual = vm.load(target, slot);
        assertEq(uint256(actual), expected);
    }

    /// @dev Function to use in the fuzz tests to assume that the caller is not the proxy admin
    /// Calls from the proxy admin are not directed to the implementation contract, so they should be
    /// treated differently in the tests.
    function assumeNotProxyAdmin(address target, address caller) internal view {
        vm.assume(caller != getProxyAdmin(target));
    }

    function getProxyAdmin(address target) internal view returns (address admin) {
        bytes32 adminBytes32 = vm.load(target, ADMIN_SLOT);
        admin = address(uint160(uint256(adminBytes32)));
    }

    function getExpectedLocationERC7201(
        string memory namespaceId,
        uint256 stolOffset
    )
        internal
        pure
        returns (bytes32 slot)
    {
        slot = keccak256(abi.encode(uint256(keccak256(bytes(namespaceId))) - 1)) & ~bytes32(uint256(0xff));
        slot = bytes32(uint256(slot) + stolOffset);
    }
}
