// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/console2.sol";
import { HomeHarness } from "./harnesses/HomeHarness.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";
import { ISystemMessenger } from "../contracts/interfaces/ISystemMessenger.sol";
import { IUpdaterManager } from "../contracts/interfaces/IUpdaterManager.sol";
import { SynapseTestWithUpdaterManager } from "./utils/SynapseTest.sol";

contract HomeTest is SynapseTestWithUpdaterManager {
    HomeHarness home;
    uint32 optimisticSeconds;

    ISystemMessenger internal systemMessenger;

    function setUp() public override {
        super.setUp();
        optimisticSeconds = 10;
        home = new HomeHarness(localDomain);
        home.initialize(IUpdaterManager(updaterManager));
        updaterManager.setHome(address(home));
        systemMessenger = ISystemMessenger(address(1234567890));
        home.setSystemMessenger(systemMessenger);
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

    function test_haltsOnFail() public {
        home.setFailed();
        vm.expectRevert("failed state");
        home.dispatch(
            remoteDomain,
            addressToBytes32(address(1337)),
            optimisticSeconds,
            getEmptyTips(),
            bytes("")
        );
    }

    // ============ DISPATCHING MESSAGING ============

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes tips,
        bytes message
    );

    // Tests sending a message and adding it to queue
    function test_dispatch() public {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = bytes("message");
        uint32 nonce = home.nonce() + 1;
        bytes memory _header = Header.formatHeader(
            localDomain,
            addressToBytes32(sender),
            nonce,
            remoteDomain,
            recipient,
            optimisticSeconds
        );
        bytes memory _tips = getDefaultTips();
        bytes memory message = Message.formatMessage(_header, _tips, messageBody);
        bytes32 messageHash = keccak256(message);
        uint256 count = home.count();
        vm.expectEmit(true, true, true, true);
        emit Dispatch(
            messageHash,
            home.count(),
            (uint64(remoteDomain) << 32) | nonce,
            _tips,
            message
        );
        hoax(sender);
        home.dispatch{ value: TOTAL_TIPS }(
            remoteDomain,
            recipient,
            optimisticSeconds,
            _tips,
            messageBody
        );
        assert(home.historicalRoots(count + 1) == home.root());
    }

    // Rejects messages over a set size
    function test_dispatchRejectBigMessage() public {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = new bytes(2 * 2**10 + 1);
        vm.prank(sender);
        vm.expectRevert("msg too long");
        home.dispatch(remoteDomain, recipient, optimisticSeconds, getEmptyTips(), messageBody);
    }

    // ============ UPDATING MESSAGES ============
    event ImproperAttestation(address updater, bytes attestation);

    function test_improperAttestation_wrongDomain() public {
        uint32 nonce = 42;
        bytes32 root = "very real much wow";
        // Any signed attestation from another chain should be rejected
        (bytes memory attestation, ) = signRemoteAttestation(updaterPK, nonce, root);
        vm.expectRevert("Wrong domain");
        home.improperAttestation(updater, attestation);
    }

    function test_improperAttestation_fraud_invalidNonce() public {
        test_dispatch();
        uint32 nonce = 2;
        bytes32 root = home.root();
        // This root exists, but with nonce = 1
        // Nonce = 0 doesn't exists yet
        _checkImproperUpdate(nonce, root);
    }

    function test_improperAttestation_fraud_correctRootWrongNonce() public {
        test_dispatch();
        test_dispatch();
        uint32 nonce = 0;
        bytes32 root = home.root();
        // This root exists, but with nonce = 1
        // nonce = 0 exists, with a different Merkle root
        _checkImproperUpdate(nonce, root);
    }

    function test_improperAttestation_fraud_validNonceWrongRoot() public {
        test_dispatch();
        uint32 nonce = 0;
        bytes32 root = "this is clearly fraud";
        // nonce = 0 exists, with a different Merkle root
        _checkImproperUpdate(nonce, root);
    }

    /// @dev Signs improper (nonce, root) attestation and presents it to Home.
    function _checkImproperUpdate(uint32 nonce, bytes32 root) internal {
        (bytes memory attestation, ) = signHomeAttestation(updaterPK, nonce, root);
        vm.expectEmit(true, true, true, true);
        emit ImproperAttestation(updater, attestation);
        // Home should recognize this as improper attestation
        assertTrue(home.improperAttestation(updater, attestation));
        // Home should be in Failed state
        assertEq(uint256(home.state()), 2);
    }

    // Dispatches 4 messages, and then Updater signs latest new roots
    function test_suggestUpdate() public {
        test_dispatch();
        test_dispatch();
        test_dispatch();
        test_dispatch();
        (uint32 nonce, bytes32 root) = home.suggestUpdate();
        // sanity checks
        assertEq(nonce, 4);
        assertEq(root, home.historicalRoots(nonce));
        (bytes memory attestation, ) = signHomeAttestation(updaterPK, nonce, root);
        // Should not be an improper attestation
        assertFalse(home.improperAttestation(updater, attestation));
        assertEq(uint256(home.state()), 1);
    }

    function test_onlySystemMessenger() public {
        vm.prank(address(systemMessenger));
        home.setSensitiveValue(1337);
        assertEq(home.sensitiveValue(), 1337);
    }

    function test_onlySystemMessenger_rejectOthers() public {
        vm.expectRevert("!systemMessenger");
        home.setSensitiveValue(1337);
    }
}
