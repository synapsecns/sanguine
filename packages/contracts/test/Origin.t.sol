// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/console2.sol";
import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";
import { Report } from "../contracts/libs/Report.sol";
import { ISystemRouter } from "../contracts/interfaces/ISystemRouter.sol";
import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

// solhint-disable func-name-mixedcase
contract OriginTest is SynapseTestWithNotaryManager {
    OriginHarness origin;
    uint32 optimisticSeconds;

    ISystemRouter internal systemRouter;

    function setUp() public override {
        super.setUp();
        optimisticSeconds = 10;
        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));
        systemRouter = ISystemRouter(address(1234567890));
        origin.setSystemRouter(systemRouter);
        origin.addGuard(guard);
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

    // ============ REPORTS ============
    event CorrectFraudReport(address indexed guard, bytes report);
    event IncorrectReport(address indexed guard, bytes report);
    event FraudAttestation(address indexed notary, bytes attestation);
    event GuardSlashed(address indexed guard, address indexed reporter);
    event NotarySlashed(address indexed notary, address indexed guard, address indexed reporter);

    function test_submitReport_wrongDomain() public {
        uint32 nonce = 42;
        bytes32 root = "very real much wow";
        // Any signed attestation from another chain should be rejected
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, root);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectRevert("Wrong domain");
        origin.submitReport(report);
    }

    function test_submitReport_notNotary() public {
        uint32 nonce = 42;
        bytes32 root = "very real much wow";
        (bytes memory attestation, ) = signOriginAttestation(fakeNotaryPK, nonce, root);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectRevert("Signer is not a notary");
        origin.submitReport(report);
    }

    function test_submitReport_notGuard() public {
        uint32 nonce = 42;
        bytes32 root = "very real much wow";
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        (bytes memory report, ) = signFraudReport(fakeGuardPK, attestation);
        vm.expectRevert("Signer is not a guard");
        origin.submitReport(report);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CORRECT FRAUD REPORT                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitReport_fraud_bigNonce() public {
        test_dispatch();
        uint32 nonce = 2;
        bytes32 root = origin.root();
        // This root exists, but with nonce = 1
        // Nonce = 2 doesn't exist yet
        _checkFraudAttestation(nonce, root);
    }

    function test_submitReport_fraud_incorrectNonce() public {
        test_dispatch();
        test_dispatch();
        uint32 nonce = 1;
        bytes32 root = origin.root();
        // This root exists, but with nonce = 2
        // nonce = 1 exists, with a different Merkle root
        _checkFraudAttestation(nonce, root);
    }

    function test_submitReport_fraud_incorrectRoot() public {
        test_dispatch();
        uint32 nonce = 1;
        bytes32 root = "this is clearly fraud";
        // nonce = 1 exists, with a different Merkle root
        _checkFraudAttestation(nonce, root);
    }

    /// @dev Signs fraud report on fraud (nonce, root) attestation and presents it to Origin.
    function _checkFraudAttestation(uint32 nonce, bytes32 root) internal {
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        vm.expectEmit(true, true, true, true);
        emit CorrectFraudReport(guard, report);
        vm.expectEmit(true, true, true, true);
        emit FraudAttestation(notary, attestation);
        vm.expectEmit(true, true, true, true);
        emit NotarySlashed(notary, guard, address(this));
        // Origin should recognize this as a correct report on fraud attestation
        assertTrue(origin.submitReport(report));
        // Origin should be in Failed state
        assertEq(uint256(origin.state()), 2);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        INCORRECT FRAUD REPORT                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitReport_fraud_incorrectReport() public {
        test_dispatch();
        uint32 nonce = 1;
        bytes32 root = origin.root();
        _checkIncorrectReport(Report.Flag.Fraud, nonce, root);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CORRECT VALID REPORT                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitReport_valid() public {
        test_dispatch();
        uint32 nonce = 1;
        bytes32 root = origin.root();
        // valid attestation
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        // this makes the report incorrect
        (bytes memory report, ) = signValidReport(guardPK, attestation);
        // Origin should recognize this as a correct report on valid attestation
        assertTrue(origin.submitReport(report));
        // Origin should be in Active state
        assertEq(uint256(origin.state()), 1);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        INCORRECT VALID REPORT                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitReport_valid_bigNonce() public {
        test_dispatch();
        uint32 nonce = 2;
        bytes32 root = origin.root();
        // This root exists, but with nonce = 1
        // Nonce = 2 doesn't exist yet
        _checkIncorrectReport(Report.Flag.Valid, nonce, root);
    }

    function test_submitReport_valid_incorrectNonce() public {
        test_dispatch();
        test_dispatch();
        uint32 nonce = 1;
        bytes32 root = origin.root();
        // This root exists, but with nonce = 2
        // nonce = 1 exists, with a different Merkle root
        _checkIncorrectReport(Report.Flag.Valid, nonce, root);
    }

    function test_submitReport_valid_incorrectRoot() public {
        test_dispatch();
        uint32 nonce = 1;
        bytes32 root = "this is clearly fraud";
        // nonce = 1 exists, with a different Merkle root
        _checkIncorrectReport(Report.Flag.Valid, nonce, root);
    }

    /// @dev Signs incorrect report on (nonce, root) attestation and presents it to Origin.
    function _checkIncorrectReport(
        Report.Flag flag,
        uint32 nonce,
        bytes32 root
    ) internal {
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        (bytes memory report, ) = signReport(guardPK, flag, attestation);
        vm.expectEmit(true, true, true, true);
        emit IncorrectReport(guard, report);
        vm.expectEmit(true, true, true, true);
        emit GuardSlashed(guard, address(this));
        if (flag == Report.Flag.Valid) {
            // Incorrect Valid Report means reported attestation is in fact fraud
            vm.expectEmit(true, true, true, true);
            emit FraudAttestation(notary, attestation);
            vm.expectEmit(true, true, true, true);
            // Guard doesn't get a reward for incorrect report
            emit NotarySlashed(notary, address(0), address(this));
        }
        // Origin should recognize this as an incorrect report on the attestation
        assertFalse(origin.submitReport(report));
        if (flag == Report.Flag.Valid) {
            // Incorrect Valid Report means reported attestation is in fact fraud
            // Origin should be in Failed state
            assertEq(uint256(origin.state()), 2);
        } else {
            // Incorrect Fraud Report means reported attestation is in fact valid
            // Origin should be in Active state
            assertEq(uint256(origin.state()), 1);
        }
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
        (bytes memory report, ) = signFraudReport(guardPK, attestation);
        // Should not be an improper attestation
        assertFalse(origin.submitReport(report));
        assertEq(uint256(origin.state()), 1);
    }

    function test_onlySystemRouter() public {
        vm.prank(address(systemRouter));
        origin.setSensitiveValue(1337);
        assertEq(origin.sensitiveValue(), 1337);
    }

    function test_onlySystemRouter_rejectOthers() public {
        vm.expectRevert("!systemRouter");
        origin.setSensitiveValue(1337);
    }
}
