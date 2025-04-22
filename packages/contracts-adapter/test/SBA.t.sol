// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseBridgeAdapter} from "../src/SynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapter} from "../src/interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "../src/interfaces/ISynapseBridgeAdapterErrors.sol";

import {BridgeMessageHarness} from "./harnesses/BridgeMessageHarness.sol";
import {EndpointMock} from "./mocks/EndpointMock.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable no-empty-blocks
abstract contract SynapseBridgeAdapterTest is Test, ISynapseBridgeAdapterErrors {
    uint32 internal constant ORIGIN_EID = 1;
    uint32 internal constant DST_EID = 2;
    uint32 internal constant UNKNOWN_EID = 3;
    bytes32 internal constant DEST_ADAPTER = keccak256("Dest Adapter");

    SynapseBridgeAdapter internal adapter;
    address internal endpoint;

    BridgeMessageHarness internal bridgeMessageLib;

    event BridgeSet(address bridge);
    event TokenAdded(address token, ISynapseBridgeAdapter.TokenType tokenType, bytes31 symbol);

    function setUp() public virtual {
        endpoint = deployEndpoint();
        adapter = deployAdapter();
        bridgeMessageLib = new BridgeMessageHarness();
        afterAdapterDeployed();
    }

    function deployAdapter() internal virtual returns (SynapseBridgeAdapter);
    function afterAdapterDeployed() internal virtual {}

    function deployEndpoint() internal virtual returns (address) {
        return address(new EndpointMock());
    }

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

    function expectRevertGasLimitBelowMinimum() internal {
        vm.expectRevert(SBA__GasLimitBelowMinimum.selector);
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
