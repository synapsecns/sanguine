// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseHyperCoreComposer} from "../src/SynapseHyperCoreComposer.sol";
import {ICoreWriter} from "../src/interfaces/ICoreWriter.sol";

import {TestToken} from "./mocks/TestToken.sol";
import {TestTokenDecimals} from "./mocks/TestTokenDecimals.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering, max-line-length
contract SynapseHyperCoreComposerTest is Test {
    uint64 internal constant CORE_INDEX = 1337;
    uint64 internal constant CORE_AMOUNT = 123_456_789;
    uint256 internal constant EVM_AMOUNT = uint256(CORE_AMOUNT) * 1e10;

    address internal recipient = makeAddr("Recipient");
    TestToken internal token;
    SynapseHyperCoreComposer internal composer;

    event Transfer(address indexed from, address indexed to, uint256 value);

    function setUp() public {
        token = new TestToken();
        composer = new SynapseHyperCoreComposer(address(this), address(token), CORE_INDEX);
        token.mintTestTokens(address(composer), EVM_AMOUNT);
        vm.etch(composer.CORE_WRITER(), hex"00");
    }

    function test_constructor() public view {
        assertEq(composer.adapter(), address(this));
        assertEq(composer.token(), address(token));
        assertEq(composer.coreIndex(), CORE_INDEX);
        assertEq(composer.EVM_DECIMALS(), 18);
        assertEq(composer.CORE_DECIMALS(), 8);
        assertEq(composer.AMOUNT_SCALE(), 1e10);
        assertEq(composer.assetBridge(), address(uint160(uint256(uint160(composer.BASE_ASSET_BRIDGE())) + CORE_INDEX)));
    }

    function test_constructor_revert_zeroAdapter() public {
        vm.expectRevert(SynapseHyperCoreComposer.SHCC__ZeroAddress.selector);
        new SynapseHyperCoreComposer(address(0), address(token), CORE_INDEX);
    }

    function test_constructor_revert_zeroToken() public {
        vm.expectRevert(SynapseHyperCoreComposer.SHCC__ZeroAddress.selector);
        new SynapseHyperCoreComposer(address(this), address(0), CORE_INDEX);
    }

    function test_constructor_revert_invalidTokenDecimals() public {
        TestTokenDecimals invalidToken = new TestTokenDecimals(6);
        vm.expectRevert(abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__InvalidTokenDecimals.selector, 6, 18));
        new SynapseHyperCoreComposer(address(this), address(invalidToken), CORE_INDEX);
    }

    function test_bridgeToHyperCore() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(CORE_AMOUNT);

        bytes memory payload =
            abi.encodePacked(composer.SPOT_SEND_HEADER(), abi.encode(recipient, CORE_INDEX, CORE_AMOUNT));
        vm.expectCall(composer.CORE_WRITER(), abi.encodeCall(ICoreWriter.sendRawAction, (payload)));
        vm.expectEmit(address(token));
        emit Transfer(address(composer), composer.assetBridge(), EVM_AMOUNT);
        vm.expectEmit(address(composer));
        emit SynapseHyperCoreComposer.HyperCoreTransferQueued(recipient, EVM_AMOUNT, CORE_AMOUNT, CORE_INDEX);

        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);

        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), EVM_AMOUNT);
        assertEq(token.balanceOf(recipient), 0);
    }

    function test_bridgeToHyperCore_revert_onlyAdapter(address caller) public {
        vm.assume(caller != address(this));
        vm.expectRevert(abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__OnlyAdapter.selector, caller));
        vm.prank(caller);
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_revert_zeroRecipient() public {
        vm.expectRevert(SynapseHyperCoreComposer.SHCC__ZeroAddress.selector);
        composer.bridgeToHyperCore(address(0), EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_revert_zeroAmount() public {
        vm.expectRevert(SynapseHyperCoreComposer.SHCC__ZeroAmount.selector);
        composer.bridgeToHyperCore(recipient, 0);
    }

    function test_bridgeToHyperCore_revert_dust() public {
        vm.expectRevert(
            abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__AmountNotRepresentable.selector, EVM_AMOUNT + 1, 1e10)
        );
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT + 1);
    }

    function test_bridgeToHyperCore_revert_amountExceedsUint64() public {
        uint256 coreAmount = uint256(type(uint64).max) + 1;
        uint256 evmAmount = coreAmount * 1e10;
        vm.expectRevert(abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__AmountExceedsUint64.selector, coreAmount));
        composer.bridgeToHyperCore(recipient, evmAmount);
    }

    function test_bridgeToHyperCore_revert_composerNotActivated() public {
        mockCoreUser(address(composer), false);
        vm.expectRevert(
            abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__CoreUserNotActivated.selector, address(composer))
        );
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_revert_recipientNotActivated() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, false);
        vm.expectRevert(abi.encodeWithSelector(SynapseHyperCoreComposer.SHCC__CoreUserNotActivated.selector, recipient));
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_revert_capacityInsufficient() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(CORE_AMOUNT - 1);
        vm.expectRevert(
            abi.encodeWithSelector(
                SynapseHyperCoreComposer.SHCC__AssetBridgeCapacityInsufficient.selector, CORE_AMOUNT - 1, CORE_AMOUNT
            )
        );
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_revert_sameBlockAggregateCapacity() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        uint64 available = CORE_AMOUNT * 2 - 1;
        mockSpotBalance(available);

        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
        token.mintTestTokens(address(composer), EVM_AMOUNT);

        vm.expectRevert(
            abi.encodeWithSelector(
                SynapseHyperCoreComposer.SHCC__AssetBridgeCapacityInsufficient.selector,
                uint256(available),
                uint256(CORE_AMOUNT) * 2
            )
        );
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);

        assertEq(token.balanceOf(address(composer)), EVM_AMOUNT);
        assertEq(token.balanceOf(composer.assetBridge()), EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_sameBlockAggregateCapacity() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(CORE_AMOUNT * 2);

        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
        token.mintTestTokens(address(composer), EVM_AMOUNT);
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);

        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), EVM_AMOUNT * 2);
    }

    function test_bridgeToHyperCore_reservationResetsNextBlock() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(CORE_AMOUNT);

        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);

        vm.roll(block.number + 1);
        token.mintTestTokens(address(composer), EVM_AMOUNT);
        mockSpotBalance(CORE_AMOUNT);
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);

        assertEq(token.balanceOf(address(composer)), 0);
        assertEq(token.balanceOf(composer.assetBridge()), EVM_AMOUNT * 2);
    }

    function test_bridgeToHyperCore_revert_coreWriter_rollsBackTransfer() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        mockSpotBalance(CORE_AMOUNT);
        vm.mockCallRevert(composer.CORE_WRITER(), bytes(""), bytes("core writer failed"));

        vm.expectRevert(bytes("core writer failed"));
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);

        assertEq(token.balanceOf(address(composer)), EVM_AMOUNT);
        assertEq(token.balanceOf(composer.assetBridge()), 0);
    }

    function test_bridgeToHyperCore_revert_precompileReadFailed() public {
        vm.mockCall(
            composer.CORE_USER_EXISTS_PRECOMPILE(), abi.encode(address(composer)), abi.encodePacked(bytes1(0x01))
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                SynapseHyperCoreComposer.SHCC__PrecompileReadFailed.selector, composer.CORE_USER_EXISTS_PRECOMPILE()
            )
        );
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
    }

    function test_bridgeToHyperCore_revert_spotBalanceReadFailed() public {
        mockCoreUser(address(composer), true);
        mockCoreUser(recipient, true);
        vm.mockCall(composer.SPOT_BALANCE_PRECOMPILE(), bytes(""), abi.encodePacked(bytes1(0x01)));
        vm.expectRevert(
            abi.encodeWithSelector(
                SynapseHyperCoreComposer.SHCC__PrecompileReadFailed.selector, composer.SPOT_BALANCE_PRECOMPILE()
            )
        );
        composer.bridgeToHyperCore(recipient, EVM_AMOUNT);
    }

    function mockCoreUser(address user, bool exists) internal {
        vm.mockCall(composer.CORE_USER_EXISTS_PRECOMPILE(), abi.encode(user), abi.encode(exists));
    }

    function mockSpotBalance(uint64 total) internal {
        vm.mockCall(
            composer.SPOT_BALANCE_PRECOMPILE(),
            abi.encode(composer.assetBridge(), CORE_INDEX),
            abi.encode(total, uint64(0), uint64(0))
        );
    }
}
