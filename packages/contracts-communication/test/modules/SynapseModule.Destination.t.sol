// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";
import {SynapseModuleEvents} from "../../contracts/events/SynapseModuleEvents.sol";
import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {InterchainBatch} from "../../contracts/libs/InterchainBatch.sol";
import {ThresholdECDSALib} from "../../contracts/libs/ThresholdECDSA.sol";
import {SynapseModule} from "../../contracts/modules/SynapseModule.sol";

import {InterchainBatchLibHarness} from "../harnesses/InterchainBatchLibHarness.sol";
import {VersionedPayloadLibHarness} from "../harnesses/VersionedPayloadLibHarness.sol";

import {SynapseGasOracleMock} from "../mocks/SynapseGasOracleMock.sol";
import {InterchainDBMock, IInterchainDB} from "../mocks/InterchainDBMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseModuleDestinationTest is Test, InterchainModuleEvents, SynapseModuleEvents {
    InterchainBatchLibHarness public batchLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    SynapseModule public module;
    SynapseGasOracleMock public gasOracle;
    InterchainDBMock public interchainDB;

    address public feeRecipient = makeAddr("FeeRecipient");
    address public owner = makeAddr("Owner");

    uint64 public constant SRC_CHAIN_ID = 1337;
    uint64 public constant DST_CHAIN_ID = 7331;

    uint16 public constant MOCK_DB_VERSION = 42;

    InterchainBatch public mockBatch =
        InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: 2, batchRoot: bytes32(uint256(3))});
    bytes public mockVersionedBatch;
    bytes public mockModuleData = "";

    uint256 public constant PK_0 = 1000;
    uint256 public constant PK_1 = 2000;
    uint256 public constant PK_2 = 3000;
    uint256 public constant PK_3 = 4000;
    uint256 public constant PK_4 = 5000;

    address public constant SIGNER_0 = 0x7F1d642DbfD62aD4A8fA9810eA619707d09825D0;
    address public constant SIGNER_1 = 0x5793e629c061e7FD642ab6A1b4d552CeC0e2D606;
    address public constant SIGNER_2 = 0xf6c0eB696e44d15E8dceb3B63A6535e469Be6C62;
    address public constant SIGNER_3 = 0xAD1117CAB797E37CAB0Eee8Ca7C30bD2452Ef2a3;
    address public constant SIGNER_4 = 0x725D327818161E0B4C6cCA5b8b1567d2A40b5B86;

    function setUp() public {
        vm.chainId(DST_CHAIN_ID);
        interchainDB = new InterchainDBMock();
        module = new SynapseModule(address(interchainDB), owner);
        gasOracle = new SynapseGasOracleMock();
        batchLibHarness = new InterchainBatchLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
        mockVersionedBatch = getVersionedBatch(mockBatch);
        vm.startPrank(owner);
        module.setGasOracle(address(gasOracle));
        module.setFeeRecipient(feeRecipient);
        module.addVerifier(SIGNER_0);
        module.addVerifier(SIGNER_1);
        module.addVerifier(SIGNER_2);
        module.setThreshold(2);
        vm.stopPrank();
    }

    function test_pks() public pure {
        assertEq(SIGNER_0, vm.addr(PK_0));
        assertEq(SIGNER_1, vm.addr(PK_1));
        assertEq(SIGNER_2, vm.addr(PK_2));
        assertEq(SIGNER_3, vm.addr(PK_3));
        assertEq(SIGNER_4, vm.addr(PK_4));
    }

    function toArray(uint256 a) internal pure returns (uint256[] memory arr) {
        arr = new uint256[](1);
        arr[0] = a;
    }

    function toArray(uint256 a, uint256 b) internal pure returns (uint256[] memory arr) {
        arr = new uint256[](2);
        arr[0] = a;
        arr[1] = b;
    }

    function toArray(uint256 a, uint256 b, uint256 c) internal pure returns (uint256[] memory arr) {
        arr = new uint256[](3);
        arr[0] = a;
        arr[1] = b;
        arr[2] = c;
    }

    function encodeAndHashBatch(InterchainBatch memory batch)
        internal
        view
        returns (bytes memory encodedBatch, bytes32 ethSignedHash)
    {
        bytes memory versionedBatch = getVersionedBatch(batch);
        encodedBatch = abi.encode(versionedBatch, mockModuleData);
        ethSignedHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(encodedBatch)));
    }

    function signBatch(
        InterchainBatch memory batch,
        uint256[] memory pks
    )
        internal
        view
        returns (bytes memory signatures)
    {
        (, bytes32 ethSignedHash) = encodeAndHashBatch(batch);
        signatures = "";
        for (uint256 i = 0; i < pks.length; ++i) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(pks[i], ethSignedHash);
            signatures = bytes.concat(signatures, abi.encodePacked(r, s, v));
        }
    }

    function getVersionedBatch(InterchainBatch memory batch) internal view returns (bytes memory versionedBatch) {
        versionedBatch = payloadLibHarness.encodeVersionedPayload(MOCK_DB_VERSION, batchLibHarness.encodeBatch(batch));
    }

    function verifyRemoteBatch(bytes memory versionedBatch, bytes memory signatures) internal {
        module.verifyRemoteBatch(abi.encode(versionedBatch, mockModuleData), signatures);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function expectCallVerifyRemoteBatch(InterchainBatch memory batch) internal {
        bytes memory versionedBatch =
            payloadLibHarness.encodeVersionedPayload(MOCK_DB_VERSION, batchLibHarness.encodeBatch(batch));
        bytes memory expectedCallData = abi.encodeCall(IInterchainDB.verifyRemoteBatch, (versionedBatch));
        // expectCall(address callee, bytes calldata data, uint64 count)
        vm.expectCall(address(interchainDB), expectedCallData, 1);
    }

    function expectEventBatchVerified(InterchainBatch memory batch) internal {
        (bytes memory encodedBatch, bytes32 ethSignedHash) = encodeAndHashBatch(batch);
        vm.expectEmit(address(module));
        emit BatchVerified(batch.srcChainId, encodedBatch, ethSignedHash);
    }

    function expectRevertIncorrectSignaturesLength(uint256 length) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__IncorrectSignaturesLength.selector, length)
        );
    }

    function expectRevertNotEnoughSignatures(uint256 provided, uint256 threshold) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__NotEnoughSignatures.selector, provided, threshold)
        );
    }

    function expectRevertRecoveredSignersNotSorted() internal {
        vm.expectRevert(ThresholdECDSALib.ThresholdECDSA__RecoveredSignersNotSorted.selector);
    }

    function expectRevertSameChainId(uint64 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainModule.InterchainModule__SameChainId.selector, chainId));
    }

    // ════════════════════════════════════════════ TESTS: VERIFY BATCH ════════════════════════════════════════════════

    // Signers order sorted by their address:
    // SIGNER_1, SIGNER_4, SIGNER_0, SIGNER_3, SIGNER_2

    // Should be verified if the enough valid signatures, which match the signers ascending order

    function test_verifyRemoteBatch_zeroSignatures_revertNotEnoughSignatures() public {
        expectRevertNotEnoughSignatures(0, 2);
        verifyRemoteBatch(mockVersionedBatch, "");
    }

    function test_verifyRemoteBatch_oneSignature_valid_revertNotEnoughSignatures() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_0));
        expectRevertNotEnoughSignatures(1, 2);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_oneSignature_invalid_revertNotEnoughSignatures() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_3));
        expectRevertNotEnoughSignatures(0, 2);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_valid_sorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_0));
        expectCallVerifyRemoteBatch(mockBatch);
        expectEventBatchVerified(mockBatch);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_valid_duplicate_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_valid_unsorted_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_0, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_invalidOne_sorted_revertNotEnoughSignatures() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_3));
        expectRevertNotEnoughSignatures(1, 2);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_invalidOne_unsorted_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_3, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_invalidTwo_sorted_revertNotEnoughSignatures() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_3));
        expectRevertNotEnoughSignatures(0, 2);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_twoSignatures_invalidTwo_duplicate_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_4));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_valid_sorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_0, PK_2));
        expectCallVerifyRemoteBatch(mockBatch);
        expectEventBatchVerified(mockBatch);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_valid_duplicate_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_0, PK_0));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_valid_unsorted_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_0, PK_2, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidOne_sorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_0, PK_2));
        expectCallVerifyRemoteBatch(mockBatch);
        expectEventBatchVerified(mockBatch);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidOne_duplicate_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_0, PK_0));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidOne_unsorted_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_0, PK_4, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidTwo_sorted_revertNotEnoughSignatures() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_3, PK_2));
        expectRevertNotEnoughSignatures(1, 2);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidTwo_duplicate_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_0, PK_3, PK_3));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidTwo_unsorted_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_3, PK_0, PK_4));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidThree_sorted_revertNotEnoughSignatures() public {
        vm.prank(owner);
        module.removeVerifier(SIGNER_0);
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_0, PK_3));
        expectRevertNotEnoughSignatures(0, 2);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidThree_duplicate_revertNotSorted() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_4, PK_3, PK_3));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_threeSignatures_invalidThree_unsorted_revertNotSorted() public {
        vm.prank(owner);
        module.removeVerifier(SIGNER_0);
        bytes memory signatures = signBatch(mockBatch, toArray(PK_3, PK_4, PK_0));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_revertSameChainId() public {
        InterchainBatch memory batch = mockBatch;
        batch.srcChainId = DST_CHAIN_ID;
        bytes memory versionedBatch = getVersionedBatch(batch);
        bytes memory signatures = signBatch(batch, toArray(PK_1, PK_0));
        expectRevertSameChainId(DST_CHAIN_ID);
        verifyRemoteBatch(versionedBatch, signatures);
    }

    function test_verifyRemoteBatch_revertZeroThreshold() public {
        // Deploy a module without setting up the threshold
        module = new SynapseModule(address(interchainDB), owner);
        vm.startPrank(owner);
        module.addVerifier(SIGNER_0);
        module.addVerifier(SIGNER_1);
        vm.stopPrank();
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_0));
        vm.expectRevert(ThresholdECDSALib.ThresholdECDSA__ZeroThreshold.selector);
        verifyRemoteBatch(mockVersionedBatch, signatures);
    }

    function test_verifyRemoteBatch_revertIncorrectSignaturesLengthTooShort() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_0));
        bytes memory signaturesShort = new bytes(signatures.length - 1);
        for (uint256 i = 0; i < signaturesShort.length; ++i) {
            signaturesShort[i] = signatures[i];
        }
        expectRevertIncorrectSignaturesLength(signaturesShort.length);
        verifyRemoteBatch(mockVersionedBatch, signaturesShort);
    }

    function test_verifyRemoteBatch_revertIncorrectSignaturesLengthTooLong() public {
        bytes memory signatures = signBatch(mockBatch, toArray(PK_1, PK_0));
        bytes memory signaturesLong = bytes.concat(signatures, bytes1(0x2A));
        expectRevertIncorrectSignaturesLength(signaturesLong.length);
        verifyRemoteBatch(mockVersionedBatch, signaturesLong);
    }
}
