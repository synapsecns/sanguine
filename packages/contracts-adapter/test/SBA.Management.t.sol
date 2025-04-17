// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter} from "../src/interfaces/ISynapseBridgeAdapter.sol";
import {SynapseBridgeAdapter, SynapseBridgeAdapterTest} from "./SBA.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseBridgeAdapterManagementTest is SynapseBridgeAdapterTest {
    address internal bridge = makeAddr("Bridge");
    address internal owner = makeAddr("Owner");
    address internal token = makeAddr("Token");
    address internal anotherToken = makeAddr("AnotherToken");
    bytes31 internal symbol = "SYMBOL";
    bytes31 internal anotherSymbol = "ANOTHERSYMBOL";

    function deployAdapter() internal virtual override returns (SynapseBridgeAdapter) {
        return new SynapseBridgeAdapter(owner);
    }

    function checkTokenAdded(
        address token_,
        ISynapseBridgeAdapter.TokenType tokenType_,
        bytes31 symbol_
    )
        internal
        view
    {
        ISynapseBridgeAdapter.TokenType adapterTokenType;
        bytes31 adapterSymbol;
        address adapterToken;
        // Check token by address
        (adapterTokenType, adapterSymbol) = adapter.getSymbolByAddress(token_);
        assertEq(uint8(adapterTokenType), uint8(tokenType_));
        assertEq(symbol_, symbol);
        // Check token by symbol
        (adapterTokenType, adapterToken) = adapter.getAddressBySymbol(symbol_);
        assertEq(uint8(adapterTokenType), uint8(tokenType_));
        assertEq(adapterToken, token_);
    }

    // ═════════════════════════════════════════════════ ADD TOKEN ═════════════════════════════════════════════════════

    function addToken(address token_, ISynapseBridgeAdapter.TokenType tokenType_, bytes31 symbol_) internal {
        vm.prank(owner);
        adapter.addToken(token_, tokenType_, symbol_);
    }

    function test_addToken_mintBurn() public {
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
    }

    function test_addToken_withdrawDeposit() public {
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, symbol);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, symbol);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, symbol);
    }

    function test_addToken_revert_tokenSymbolAlreadyAdded() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        // Same token type
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        // Different token type
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, symbol);
    }

    function test_addToken_revert_tokenAlreadyAdded() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        // Same token type
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, anotherSymbol);
        // Different token type
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, anotherSymbol);
    }

    function test_addToken_revert_symbolAlreadyAdded() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        // Same token type
        expectRevertSymbolAlreadyAdded(symbol);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        // Different token type
        expectRevertSymbolAlreadyAdded(symbol);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, symbol);
    }

    function test_addToken_revert_zeroAddress() public {
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
    }

    function test_addToken_revert_zeroSymbol() public {
        expectRevertZeroSymbol();
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, bytes31(0));
    }

    function test_addToken_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertCallerNotOwner(caller);
        vm.prank(caller);
        adapter.addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
    }

    function test_getSymbolByAddress_revert_tokenUnknown() public {
        expectRevertTokenUnknown(token);
        adapter.getSymbolByAddress(token);
    }

    function test_getAddressBySymbol_revert_symbolUnknown() public {
        expectRevertSymbolUnknown(symbol);
        adapter.getAddressBySymbol(symbol);
    }

    // ════════════════════════════════════════════════ SET BRIDGE ═════════════════════════════════════════════════════

    function test_setBridge() public {
        expectEventBridgeSet(bridge);
        vm.prank(owner);
        adapter.setBridge(bridge);
        assertEq(adapter.bridge(), bridge);
    }

    function test_setBridge_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertCallerNotOwner(caller);
        vm.prank(caller);
        adapter.setBridge(bridge);
    }

    function test_setBridge_revert_bridgeAlreadySet() public {
        vm.prank(owner);
        adapter.setBridge(bridge);
        expectRevertBridgeAlreadySet();
        vm.prank(owner);
        adapter.setBridge(bridge);
    }

    function test_setBridge_revert_zeroAddress() public {
        expectRevertZeroAddress();
        vm.prank(owner);
        adapter.setBridge(address(0));
    }
}
