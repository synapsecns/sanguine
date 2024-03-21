// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {
    ProxyAdmin,
    TransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Test} from "forge-std/Test.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract ProxyTest is Test {
    ProxyAdmin public proxyAdmin;

    function deployProxy(address implementation) internal returns (address proxy) {
        proxy = deployProxy(implementation, "");
    }

    function deployProxy(address implementation, bytes memory initData) internal returns (address proxy) {
        if (address(proxyAdmin) == address(0)) {
            // Use a single proxy admin owned by this contract for tests simplicity
            proxyAdmin = new ProxyAdmin(address(this));
        }
        proxy = address(new TransparentUpgradeableProxy(implementation, address(proxyAdmin), initData));
    }

    function expectRevertInvalidInitialization() internal {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
    }
}
