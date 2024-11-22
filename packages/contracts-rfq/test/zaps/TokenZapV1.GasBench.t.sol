// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {TokenZapV1} from "../../contracts/zaps/TokenZapV1.sol";

import {SimpleVaultMock} from "../mocks/SimpleVaultMock.sol";
import {MockERC20} from "../MockERC20.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract TokenZapV1GasBenchmarkTest is Test {
    uint256 internal constant AMOUNT = 0.1337 ether;

    SimpleVaultMock internal vault;
    TokenZapV1 internal tokenZap;
    MockERC20 internal erc20;

    address internal user;
    address internal nativeGasToken = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    function setUp() public {
        tokenZap = new TokenZapV1();
        vault = new SimpleVaultMock();
        erc20 = new MockERC20("TKN", 18);

        user = makeAddr("user");

        erc20.mint(address(this), AMOUNT);
        deal(address(this), AMOUNT);
        // To simulate an average case we assume that the Vault contract has already other deposited funds.
        erc20.mint(address(vault), 1000 * AMOUNT);
        deal(address(vault), 1000 * AMOUNT);
        // We also assume that this is not the first tx through the Zap, so the infinite approval has already been set.
        vm.prank(address(tokenZap));
        erc20.approve(address(vault), type(uint256).max);
    }

    function getVaultPayload(address token, uint256 amount) public view returns (bytes memory) {
        return abi.encodeCall(vault.deposit, (token, amount, user));
    }

    function getZapData(bytes memory originalPayload) public view returns (bytes memory) {
        // Amount is the second argument of the deposit function.
        return tokenZap.encodeZapData(address(vault), originalPayload, 4 + 32);
    }

    function test_deposit_erc20() public {
        bytes memory depositPayload = getVaultPayload(address(erc20), AMOUNT);
        bytes memory zapData = getZapData(depositPayload);
        // Transfer tokens to the zap contract first.
        erc20.transfer(address(tokenZap), AMOUNT);
        tokenZap.zap(address(erc20), AMOUNT, zapData);
        // Check that the vault registered the deposit.
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    function test_deposit_native() public {
        bytes memory depositPayload = getVaultPayload(nativeGasToken, AMOUNT);
        bytes memory zapData = getZapData(depositPayload);
        tokenZap.zap{value: AMOUNT}(nativeGasToken, AMOUNT, zapData);
        // Check that the vault registered the deposit.
        assertEq(vault.balanceOf(user, nativeGasToken), AMOUNT);
    }
}
