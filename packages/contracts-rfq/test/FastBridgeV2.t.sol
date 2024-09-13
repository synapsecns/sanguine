// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2} from "../contracts/FastBridgeV2.sol";

import {MockERC20} from "./MockERC20.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Test} from "forge-std/Test.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";

// solhint-disable no-empty-blocks
abstract contract FastBridgeV2Test is Test {
    using stdStorage for StdStorage;

    address public constant ETH_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    FastBridgeV2 public fastBridge;
    MockERC20 public srcToken;
    MockERC20 public dstToken;

    address public relayerA = makeAddr("Relayer A");
    address public relayerB = makeAddr("Relayer B");
    address public guard = makeAddr("Guard");
    address public userA = makeAddr("User A");
    address public userB = makeAddr("User B");
    address public governor = makeAddr("Governor");
    address public refunder = makeAddr("Refunder");

    function setUp() public virtual {
        srcToken = new MockERC20("SrcToken", 6);
        dstToken = new MockERC20("DstToken", 6);
        fastBridge = deployFastBridge();
        configureFastBridge();
        mintTokens();
    }

    function deployFastBridge() public virtual returns (FastBridgeV2);

    function configureFastBridge() public virtual {}

    function mintTokens() public virtual {}

    function expectUnauthorized(address caller, bytes32 role) public {
        vm.expectRevert(abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, role));
    }

    function cheatCollectedProtocolFees(address token, uint256 amount) public {
        stdstore.target(address(fastBridge)).sig("protocolFees(address)").with_key(token).checked_write(amount);
    }
}
