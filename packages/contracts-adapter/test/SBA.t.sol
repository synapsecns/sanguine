// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseBridgeAdapter} from "../src/SynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapter} from "../src/interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "../src/interfaces/ISynapseBridgeAdapterErrors.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable no-empty-blocks
abstract contract SynapseBridgeAdapterTest is Test, ISynapseBridgeAdapterErrors {
    SynapseBridgeAdapter internal adapter;

    event BridgeSet(address bridge);
    event TokenAdded(address token, ISynapseBridgeAdapter.TokenType tokenType, bytes31 symbol);

    function setUp() public virtual {
        adapter = deployAdapter();
        configureAdapter();
    }

    function deployAdapter() internal virtual returns (SynapseBridgeAdapter);
    function configureAdapter() internal virtual {}

    function expectEventBridgeSet(address bridge) internal {
        vm.expectEmit(address(adapter));
        emit BridgeSet(bridge);
    }

    function expectEventTokenAdded(address token, ISynapseBridgeAdapter.TokenType tokenType, bytes31 symbol) internal {
        vm.expectEmit(address(adapter));
        emit TokenAdded(token, tokenType, symbol);
    }

    function expectRevertCallerNotOwner(address caller) internal {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, caller));
    }

    function expectRevertBridgeAlreadySet() internal {
        vm.expectRevert(SBA__BridgeAlreadySet.selector);
    }

    function expectRevertBridgeNotSet() internal {
        vm.expectRevert(SBA__BridgeNotSet.selector);
    }

    function expectRevertRecipientZeroAddress() internal {
        vm.expectRevert(SBA__RecipientZeroAddress.selector);
    }

    function expectRevertSymbolAlreadyAdded(bytes31 symbol) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__SymbolAlreadyAdded.selector, symbol));
    }

    function expectRevertSymbolUnknown(bytes31 symbol) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__SymbolUnknown.selector, symbol));
    }

    function expectRevertTokenAlreadyAdded(address token) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__TokenAlreadyAdded.selector, token));
    }

    function expectRevertTokenUnknown(address token) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__TokenUnknown.selector, token));
    }

    function expectRevertZeroAddress() internal {
        vm.expectRevert(SBA__ZeroAddress.selector);
    }

    function expectRevertZeroAmount() internal {
        vm.expectRevert(SBA__ZeroAmount.selector);
    }

    function expectRevertZeroSymbol() internal {
        vm.expectRevert(SBA__ZeroSymbol.selector);
    }
}
