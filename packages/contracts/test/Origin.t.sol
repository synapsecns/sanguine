// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/console2.sol";
import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";
import { ISystemMessenger } from "../contracts/interfaces/ISystemMessenger.sol";
import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

contract OriginTest is SynapseTestWithNotaryManager {
    OriginHarness origin;
    uint32 optimisticSeconds;

    ISystemMessenger internal systemMessenger;

    function setUp() public override {
        super.setUp();
        optimisticSeconds = 10;
        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));
        systemMessenger = ISystemMessenger(address(1234567890));
        origin.setSystemMessenger(systemMessenger);
    }

    // ============ STATE AND PERMISSIONING ============
    function test_correctlyInitialized() public {
        assertEq(address(origin.notaryManager()), address(notaryManager));
        assertEq(origin.owner(), address(this));
        assertEq(uint256(origin.state()), 1);
        assertTrue(origin.isNotary(notary));
        // Root of an empty sparse Merkle tree should be stored with nonce=0
        assertEq(origin.historicalRoots(0), origin.root());
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        origin.initialize(notaryManager);
    }

    function test_cannotSetNotaryAsNotNotaryManager() public {
        vm.expectRevert("!notaryManager");
        origin.setNotary(address(0));
    }

    function test_setNotary() public {
        assertFalse(origin.isNotary(address(1337)));
        vm.prank(address(notaryManager));
        origin.setNotary(address(1337));
        assertTrue(origin.isNotary(address(1337)));
    }

    function test_cannotSetNotaryManagerAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != origin.owner());
        vm.startPrank(_notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        // Must pass in a contract to setNotaryManager, otherwise will revert with !contract notaryManger
        origin.setNotaryManager(address(origin));
    }

    function test_setNotaryManager() public {
        assertFalse(address(origin.notaryManager()) == address(origin));
        origin.setNotaryManager(address(origin));
        // Must pass in a contract to setNotaryManager, otherwise will revert with !contract notaryManger
        assertEq(address(origin.notaryManager()), address(origin));
    }

    function test_onlyContractCanBeNotaryManager() public {
        vm.expectRevert("!contract notaryManager");
        origin.setNotaryManager(address(1337));
    }

    function test_haltsOnFail() public {
        origin.setFailed(notary);
        vm.expectRevert("failed state");
        origin.dispatch(
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
        uint32 nonce = origin.nonce() + 1;
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
        uint256 count = origin.count();
        vm.expectEmit(true, true, true, true);
        emit Dispatch(
            messageHash,
            origin.count(),
            (uint64(remoteDomain) << 32) | nonce,
            _tips,
            message
        );
        hoax(sender);
        origin.dispatch{ value: TOTAL_TIPS }(
            remoteDomain,
            recipient,
            optimisticSeconds,
            _tips,
            messageBody
        );
        assert(origin.historicalRoots(count + 1) == origin.root());
    }

    // Rejects messages over a set size
    function test_dispatchRejectBigMessage() public {
        bytes32 recipient = addressToBytes32(vm.addr(1337));
        address sender = vm.addr(1555);
        bytes memory messageBody = new bytes(2 * 2**10 + 1);
        vm.prank(sender);
        vm.expectRevert("msg too long");
        origin.dispatch(remoteDomain, recipient, optimisticSeconds, getEmptyTips(), messageBody);
    }

    // ============ UPDATING MESSAGES ============
    event ImproperAttestation(address notary, bytes attestation);

    function test_improperAttestation_wrongDomain() public {
        uint32 nonce = 42;
        bytes32 root = "very real much wow";
        // Any signed attestation from another chain should be rejected
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, root);
        vm.expectRevert("Wrong domain");
        origin.improperAttestation(attestation);
    }

    function test_improperAttestation_fraud_invalidNonce() public {
        test_dispatch();
        uint32 nonce = 2;
        bytes32 root = origin.root();
        // This root exists, but with nonce = 1
        // Nonce = 0 doesn't exists yet
        _checkImproperAttestation(nonce, root);
    }

    function test_improperAttestation_fraud_correctRootWrongNonce() public {
        test_dispatch();
        test_dispatch();
        uint32 nonce = 0;
        bytes32 root = origin.root();
        // This root exists, but with nonce = 1
        // nonce = 2 exists, with a different Merkle root
        _checkImproperAttestation(nonce, root);
    }

    function test_improperAttestation_fraud_validNonceWrongRoot() public {
        test_dispatch();
        uint32 nonce = 0;
        bytes32 root = "this is clearly fraud";
        // nonce = 0 exists, with a different Merkle root
        _checkImproperAttestation(nonce, root);
    }

    /// @dev Signs improper (nonce, root) attestation and presents it to Origin.
    function _checkImproperAttestation(uint32 nonce, bytes32 root) internal {
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        vm.expectEmit(true, true, true, true);
        emit ImproperAttestation(notary, attestation);
        // Origin should recognize this as improper attestation
        assertTrue(origin.improperAttestation(attestation));
        // Origin should be in Failed state
        assertEq(uint256(origin.state()), 2);
    }

    // Dispatches 4 messages, and then Notary signs latest new roots
    function test_suggestAttestation() public {
        test_dispatch();
        test_dispatch();
        test_dispatch();
        test_dispatch();
        (uint32 nonce, bytes32 root) = origin.suggestAttestation();
        // sanity checks
        assertEq(nonce, 4);
        assertEq(root, origin.historicalRoots(nonce));
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        // Should not be an improper attestation
        assertFalse(origin.improperAttestation(attestation));
        assertEq(uint256(origin.state()), 1);
    }

    function test_onlySystemMessenger() public {
        vm.prank(address(systemMessenger));
        origin.setSensitiveValue(1337);
        assertEq(origin.sensitiveValue(), 1337);
    }

    function test_onlySystemMessenger_rejectOthers() public {
        vm.expectRevert("!systemMessenger");
        origin.setSensitiveValue(1337);
    }
}
