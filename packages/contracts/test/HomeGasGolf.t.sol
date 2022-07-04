// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/console2.sol";
import { HomeHarness } from "./harnesses/HomeHarness.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";
import { IUpdaterManager } from "../contracts/interfaces/IUpdaterManager.sol";
import { SynapseTestWithUpdaterManager } from "./utils/SynapseTest.sol";

contract HomeGasGolfTest is SynapseTestWithUpdaterManager {
    HomeHarness home;

    function setUp() public override {
        super.setUp();
        home = new HomeHarness(localDomain);
        home.initialize(IUpdaterManager(updaterManager));
        updaterManager.setHome(address(home));
    }

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes32 committedRoot,
        bytes message
    );

    function _dispatch() internal returns (bytes32 newRoot) {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = bytes("message");
        uint32 nonce = home.nonces(remoteDomain);
        bytes memory _header = Header.formatHeader(
            localDomain,
            addressToBytes32(sender),
            nonce,
            remoteDomain,
            recipient,
            0
        );
        bytes memory message = Message.formatMessage(_header, messageBody);
        bytes32 messageHash = keccak256(message);
        vm.expectEmit(true, true, true, true);
        emit Dispatch(
            messageHash,
            home.count(),
            (uint64(remoteDomain) << 32) | nonce,
            home.committedRoot(),
            message
        );
        vm.prank(sender);
        home.dispatch(remoteDomain, recipient, 0, messageBody);
        newRoot = home.root();
    }

    function _update(bytes32 _newRoot, uint256 _newLength) internal {
        bytes memory sig = signHomeUpdate(updaterPK, home.committedRoot(), _newRoot);
        home.update(home.committedRoot(), _newRoot, sig);
        // Updater signs latest root
        assertEq(home.committedRoot(), _newRoot);
        // Queue is cleared, no messages left to sign
        assertEq(home.queueLength(), _newLength);
    }

    function test_contains() public {
        uint256 amount = 15;
        bytes32[] memory roots = new bytes32[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            roots[i] = _dispatch();
        }
        for (uint256 i = 0; i < amount; ++i) {
            assertTrue(home.queueContains(roots[i]));
        }
        _dispatch();
        _update(roots[0], amount);
        assertFalse(home.queueContains(roots[0]));
    }

    function test_D15_U15_D15_U15() public {
        bytes32 root1;
        for (uint256 i = 0; i < 15; ++i) {
            root1 = _dispatch();
        } // root1 refers to msg 15

        // use msgs [01-15] for the update
        // msgs left in queue: []
        _update(root1, 0);

        bytes32 root2;
        for (uint256 i = 0; i < 15; ++i) {
            root2 = _dispatch();
        } // root2 refers to msg 30

        // use msgs [16-30] for the update
        // msgs left in queue: []
        _update(root2, 0);
    }

    function test_D15_U10_D15_U10_U10() public {
        bytes32 root1;
        for (uint256 i = 0; i < 10; ++i) {
            root1 = _dispatch();
        } // root1 refers to msg 10

        for (uint256 i = 0; i < 5; ++i) {
            _dispatch();
        } // a total of 15 msgs has been sent

        // use msgs [01-10] for the update
        // msgs left in queue: [11-15]
        _update(root1, 5);

        bytes32 root2;
        for (uint256 i = 0; i < 5; ++i) {
            root2 = _dispatch();
        } // root2 refers to msg 20

        bytes32 root3;
        for (uint256 i = 0; i < 10; ++i) {
            root3 = _dispatch();
        } // root3 refers to msg 30

        // use msgs [11-20] for the update
        // msgs left in queue: [21-30]
        _update(root2, 10);

        // use msgs [11-20] for the update
        // msgs left in queue: []
        _update(root3, 0);
    }

    function test_D15_U05_D10_U10_D05_U15() public {
        bytes32 root1;
        for (uint256 i = 0; i < 5; ++i) {
            root1 = _dispatch();
        } // root1 refers to msg 05

        bytes32 root2;
        for (uint256 i = 0; i < 10; ++i) {
            root2 = _dispatch();
        } // root2 refers to msg 15

        // use msgs [01-05] for the update
        // msgs left in queue: [06-15]
        _update(root1, 10);
        for (uint256 i = 0; i < 10; ++i) {
            _dispatch();
        } // a total of 25 msgs has been sent

        // use msgs [06-15] for the update
        // msgs left in queue: [16-25]
        _update(root2, 10);

        bytes32 root3;
        for (uint256 i = 0; i < 5; ++i) {
            root3 = _dispatch();
        } // root3 refers to msg 30

        // use msgs [16-30] for the update
        // msgs left in queue: []
        _update(root3, 0);
    }
}
