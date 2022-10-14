// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/console2.sol";
import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";
import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

// solhint-disable func-name-mixedcase
contract OriginGasGolfTest is SynapseTestWithNotaryManager {
    OriginHarness internal origin;

    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    function setUp() public override {
        super.setUp();
        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));
    }

    function test_dispatch_30() public {
        uint256 amount = 30;
        bytes32[] memory roots = new bytes32[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            roots[i] = _dispatch();
        }
        for (uint256 i = 0; i < amount; ++i) {
            assertEq(origin.historicalRoots(i + 1), roots[i]);
        }
    }

    function _dispatch() internal returns (bytes32 newRoot) {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = bytes("message");
        uint32 nonce = origin.nonce() + 1;
        bytes memory _header = Header.formatHeader(
            localDomain,
            addressToBytes32(sender),
            nonce,
            remoteDomain,
            recipient,
            0
        );
        bytes memory _tips = getDefaultTips();
        bytes memory message = Message.formatMessage(_header, _tips, messageBody);
        bytes32 messageHash = keccak256(message);
        vm.expectEmit(true, true, true, true);
        emit Dispatch(messageHash, nonce, remoteDomain, _tips, message);
        hoax(sender);
        origin.dispatch{ value: TOTAL_TIPS }(remoteDomain, recipient, 0, _tips, messageBody);
        newRoot = origin.root();
    }
}
