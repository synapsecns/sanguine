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
    uint32 internal constant SRC_EID = 1;
    uint32 internal constant DST_EID = 2;
    uint32 internal constant OTHER_DST_EID = 3;
    uint32 internal constant UNKNOWN_EID = 1337;
    bytes32 internal constant REMOTE_ADAPTER = keccak256("Dest Adapter");
    bytes32 internal constant MOCK_GUID = keccak256("mockGuid");

    address internal remoteToken = makeAddr("Remote Token");

    SynapseBridgeAdapter internal adapter;
    address internal endpoint;

    BridgeMessageHarness internal bridgeMessageLib;

    event BridgeSet(address bridge);
    event TokenAdded(
        address token, ISynapseBridgeAdapter.TokenType tokenType, ISynapseBridgeAdapter.RemoteToken[] remoteTokens
    );
    event TokenSent(uint32 indexed dstEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);
    event TokenReceived(uint32 indexed srcEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);

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

    function expectEventTokenAdded(
        address token,
        ISynapseBridgeAdapter.TokenType tokenType,
        ISynapseBridgeAdapter.RemoteToken[] memory remoteTokens
    )
        internal
    {
        vm.expectEmit(address(adapter));
        emit TokenAdded(token, tokenType, remoteTokens);
    }

    function expectEventTokenSent(uint32 dstEid, address to, address token, uint256 amount, bytes32 guid) internal {
        vm.expectEmit(address(adapter));
        emit TokenSent(dstEid, to, token, amount, guid);
    }

    function expectEventTokenReceived(
        uint32 srcEid,
        address to,
        address token,
        uint256 amount,
        bytes32 guid
    )
        internal
    {
        vm.expectEmit(address(adapter));
        emit TokenReceived(srcEid, to, token, amount, guid);
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

    function expectRevertLocalPairAlreadyExists(uint32 eid, address remoteAddr) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__LocalPairAlreadyExists.selector, eid, remoteAddr));
    }

    function expectRevertLocalPairNotFound(uint32 eid, address remoteAddr) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__LocalPairNotFound.selector, eid, remoteAddr));
    }

    function expectRevertRemotePairAlreadySet(uint32 eid, address localAddr) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__RemotePairAlreadySet.selector, eid, localAddr));
    }

    function expectRevertRemotePairNotSet(uint32 eid, address localAddr) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__RemotePairNotSet.selector, eid, localAddr));
    }

    function expectRevertTokenAlreadyAdded(address token) internal {
        vm.expectRevert(abi.encodeWithSelector(SBA__TokenAlreadyAdded.selector, token));
    }

    function expectRevertTokenTypeUnknown() internal {
        vm.expectRevert(SBA__TokenTypeUnknown.selector);
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

    function toArray(ISynapseBridgeAdapter.RemoteToken memory a)
        internal
        pure
        returns (ISynapseBridgeAdapter.RemoteToken[] memory arr)
    {
        arr = new ISynapseBridgeAdapter.RemoteToken[](1);
        arr[0] = a;
    }

    function toArray(
        ISynapseBridgeAdapter.RemoteToken memory a,
        ISynapseBridgeAdapter.RemoteToken memory b
    )
        internal
        pure
        returns (ISynapseBridgeAdapter.RemoteToken[] memory arr)
    {
        arr = new ISynapseBridgeAdapter.RemoteToken[](2);
        arr[0] = a;
        arr[1] = b;
    }
}
