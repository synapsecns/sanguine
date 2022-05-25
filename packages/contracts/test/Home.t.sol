// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/console2.sol";
import { Home } from "../contracts/Home.sol";
import { Message } from "../contracts/libs/Message.sol";
import { IUpdaterManager } from "../contracts/interfaces/IUpdaterManager.sol";
import { SynapseTestWithUpdaterManager } from "./utils/SynapseTest.sol";

contract HomeTest is SynapseTestWithUpdaterManager {
    Home home;

    function setUp() public override {
        super.setUp();
        home = new Home(localDomain);
        home.initialize(IUpdaterManager(updaterManager));
        updaterManager.setHome(address(home));
    }

    // ============ STATE AND PERMISSIONING ============
    function test_correctlyInitialized() public {
        assertEq(address(home.updaterManager()), address(updaterManager));
        assertEq(home.owner(), address(this));
        assertEq(uint256(home.state()), 1);
        assertEq(home.updater(), updater);
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        home.initialize(updaterManager);
    }

    function test_cannotSetUpdaterAsNotUpdaterManager() public {
        vm.expectRevert("!updaterManager");
        home.setUpdater(address(0));
    }

    function test_setUpdater() public {
        assertFalse(home.updater() == address(1337));
        vm.prank(address(updaterManager));
        home.setUpdater(address(1337));
        assertEq(home.updater(), address(1337));
    }

    function test_cannotSetUpdaterManagerAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != home.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        // Must pass in a contract to setUpdaterManager, otherwise will revert with !contract updaterManger
        home.setUpdaterManager(address(home));
    }

    function test_setUpdaterManager() public {
        assertFalse(address(home.updaterManager()) == address(home));
        home.setUpdaterManager(address(home));
        // Must pass in a contract to setUpdaterManager, otherwise will revert with !contract updaterManger
        assertEq(address(home.updaterManager()), address(home));
    }

    function test_onlyContractCanBeUpdaterManager() public {
        vm.expectRevert("!contract updaterManager");
        home.setUpdaterManager(address(1337));
    }

    // TODO: testHashDomain against Go generated domains
    // function test_homeDomainHash() public {}

    function test_committedRoot() public {
        bytes32 emptyRoot;
        assertEq(home.committedRoot(), emptyRoot);
    }

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes32 committedRoot,
        bytes message
    );

    function test_dispatch() public {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = bytes("message");
        uint32 nonce = home.nonces(remoteDomain);
        bytes memory message = Message.formatMessage(
            localDomain,
            addressToBytes32(sender),
            nonce,
            remoteDomain,
            recipient,
            messageBody
        );
        bytes32 messageHash = keccak256(message);
        vm.expectEmit(true, true, true, true);
        emit Dispatch(
            messageHash,
            0,
            (uint64(remoteDomain) << 32) | nonce,
            home.committedRoot(),
            message
        );
        vm.prank(sender);
        home.dispatch(remoteDomain, recipient, messageBody);
        assert(home.queueContains(home.root()));
    }

    function test_dispatchRejectBigMessage() public {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = new bytes(2 * 2**10 + 1);
        uint32 nonce = home.nonces(remoteDomain);
        bytes memory message = Message.formatMessage(
            localDomain,
            addressToBytes32(sender),
            nonce,
            remoteDomain,
            recipient,
            messageBody
        );
        vm.prank(sender);
        vm.expectRevert("msg too long");
        home.dispatch(remoteDomain, recipient, messageBody);
    }

    event ImproperUpdate(bytes32 oldRoot, bytes32 newRoot, bytes signature);

    function test_improperUpdateAndFailedState() public {
        assertEq(uint256(home.state()), 1);
        bytes32 newRoot = "new root";
        bytes32 oldRoot = home.committedRoot();
        bytes memory sig = signHomeUpdate(updaterPK, oldRoot, newRoot);
        vm.expectEmit(false, false, false, true);
        emit ImproperUpdate(oldRoot, newRoot, sig);
        home.improperUpdate(oldRoot, newRoot, sig);
        assertEq(uint256(home.state()), 2);
        vm.expectRevert("failed state");
        home.dispatch(0, bytes32(0), bytes(""));
    }

    function test_update() public {
        // Send message first, which will add a new root to merkle
        test_dispatch();
        bytes32 newRoot = home.queueEnd();
        // sign latest update for new root
        bytes memory sig = signHomeUpdate(updaterPK, home.committedRoot(), newRoot);
        home.update(home.committedRoot(), home.queueEnd(), sig);
        // Updater signs latest root
        assertEq(home.committedRoot(), newRoot);
        // Queue is cleared, no messages left to sign
        assert(!home.queueContains(newRoot));
        assertEq(home.queueLength(), 0);
    }

    function test_cannotUpdateAsFakeUpdater() public {
        // Send message first, which will add a new root to merkle
        test_dispatch();
        bytes32 newRoot = home.queueEnd();
        bytes32 comittedRoot = home.committedRoot();
        // fake updater sign new root
        bytes memory sig = signHomeUpdate(fakeUpdaterPK, comittedRoot, newRoot);
        vm.expectRevert("!updater sig");
        home.update(comittedRoot, newRoot, sig);
    }
}
