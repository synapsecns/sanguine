// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";
import {SynapseModuleEvents} from "../../contracts/events/SynapseModuleEvents.sol";
import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";
import {ThresholdECDSALib} from "../../contracts/libs/ThresholdECDSA.sol";
import {SynapseModule} from "../../contracts/modules/SynapseModule.sol";

import {InterchainEntryLibHarness} from "../harnesses/InterchainEntryLibHarness.sol";
import {VersionedPayloadLibHarness} from "../harnesses/VersionedPayloadLibHarness.sol";

import {SynapseGasOracleMock} from "../mocks/SynapseGasOracleMock.sol";
import {InterchainDBMock, IInterchainDB} from "../mocks/InterchainDBMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseModuleDestinationTest is Test, InterchainModuleEvents, SynapseModuleEvents {
    InterchainEntryLibHarness public entryLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    SynapseModule public module;
    SynapseGasOracleMock public gasOracle;
    InterchainDBMock public interchainDB;

    address public feeRecipient = makeAddr("FeeRecipient");
    address public owner = makeAddr("Owner");

    uint64 public constant SRC_CHAIN_ID = 1337;
    uint64 public constant DST_CHAIN_ID = 7331;

    uint16 public constant MOCK_DB_VERSION = 42;

    InterchainEntry public mockEntry =
        InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: 2, entryValue: bytes32(uint256(3))});
    bytes public mockVersionedEntry;
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
        entryLibHarness = new InterchainEntryLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
        mockVersionedEntry = getVersionedEntry(mockEntry);
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

    function encodeAndHashEntry(InterchainEntry memory entry)
        internal
        view
        returns (bytes memory encodedEntry, bytes32 ethSignedHash)
    {
        bytes memory versionedEntry = getVersionedEntry(entry);
        encodedEntry = abi.encode(versionedEntry, mockModuleData);
        ethSignedHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(encodedEntry)));
    }

    function signEntry(
        InterchainEntry memory entry,
        uint256[] memory pks
    )
        internal
        view
        returns (bytes memory signatures)
    {
        (, bytes32 ethSignedHash) = encodeAndHashEntry(entry);
        signatures = "";
        for (uint256 i = 0; i < pks.length; ++i) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(pks[i], ethSignedHash);
            signatures = bytes.concat(signatures, abi.encodePacked(r, s, v));
        }
    }

    function getVersionedEntry(InterchainEntry memory entry) internal view returns (bytes memory versionedEntry) {
        versionedEntry = payloadLibHarness.encodeVersionedPayload(MOCK_DB_VERSION, entryLibHarness.encodeEntry(entry));
    }

    function verifyRemoteEntry(bytes memory versionedEntry, bytes memory signatures) internal {
        module.verifyRemoteEntry(abi.encode(versionedEntry, mockModuleData), signatures);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function expectCallVerifyRemoteEntry(InterchainEntry memory entry) internal {
        bytes memory versionedEntry =
            payloadLibHarness.encodeVersionedPayload(MOCK_DB_VERSION, entryLibHarness.encodeEntry(entry));
        bytes memory expectedCallData = abi.encodeCall(IInterchainDB.verifyRemoteEntry, (versionedEntry));
        // expectCall(address callee, bytes calldata data, uint64 count)
        vm.expectCall(address(interchainDB), expectedCallData, 1);
    }

    function expectEventEntryVerified(InterchainEntry memory entry) internal {
        (bytes memory encodedEntry, bytes32 ethSignedHash) = encodeAndHashEntry(entry);
        vm.expectEmit(address(module));
        emit EntryVerified(entry.srcChainId, encodedEntry, ethSignedHash);
    }

    function expectRevertSignaturesPayloadLengthInvalid(uint256 length) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__SignaturesPayloadLengthInvalid.selector, length)
        );
    }

    function expectRevertSignaturesAmountBelowThreshold(uint256 signaturesAmount, uint256 threshold) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                ThresholdECDSALib.ThresholdECDSA__SignaturesAmountBelowThreshold.selector, signaturesAmount, threshold
            )
        );
    }

    function expectRevertRecoveredSignersNotSorted() internal {
        vm.expectRevert(ThresholdECDSALib.ThresholdECDSA__RecoveredSignersNotSorted.selector);
    }

    function expectRevertChainIdNotRemote(uint64 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainModule.InterchainModule__ChainIdNotRemote.selector, chainId));
    }

    // ════════════════════════════════════════════ TESTS: VERIFY ENTRY ════════════════════════════════════════════════

    // Signers order sorted by their address:
    // SIGNER_1, SIGNER_4, SIGNER_0, SIGNER_3, SIGNER_2

    // Should be verified if the enough valid signatures, which match the signers ascending order

    function test_verifyRemoteEntry_zeroSignatures_revertSignaturesAmountBelowThreshold() public {
        expectRevertSignaturesAmountBelowThreshold(0, 2);
        verifyRemoteEntry(mockVersionedEntry, "");
    }

    function test_verifyRemoteEntry_oneSignature_valid_revertSignaturesAmountBelowThreshold() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_0));
        expectRevertSignaturesAmountBelowThreshold(1, 2);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_oneSignature_invalid_revertSignaturesAmountBelowThreshold() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_3));
        expectRevertSignaturesAmountBelowThreshold(0, 2);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_valid_sorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_0));
        expectCallVerifyRemoteEntry(mockEntry);
        expectEventEntryVerified(mockEntry);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_valid_duplicate_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_valid_unsorted_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_0, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_invalidOne_sorted_revertSignaturesAmountBelowThreshold() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_3));
        expectRevertSignaturesAmountBelowThreshold(1, 2);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_invalidOne_unsorted_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_3, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_invalidTwo_sorted_revertSignaturesAmountBelowThreshold() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_3));
        expectRevertSignaturesAmountBelowThreshold(0, 2);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_twoSignatures_invalidTwo_duplicate_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_4));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_valid_sorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_0, PK_2));
        expectCallVerifyRemoteEntry(mockEntry);
        expectEventEntryVerified(mockEntry);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_valid_duplicate_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_0, PK_0));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_valid_unsorted_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_0, PK_2, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidOne_sorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_0, PK_2));
        expectCallVerifyRemoteEntry(mockEntry);
        expectEventEntryVerified(mockEntry);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidOne_duplicate_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_0, PK_0));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidOne_unsorted_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_0, PK_4, PK_1));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidTwo_sorted_revertSignaturesAmountBelowThreshold() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_3, PK_2));
        expectRevertSignaturesAmountBelowThreshold(1, 2);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidTwo_duplicate_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_0, PK_3, PK_3));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidTwo_unsorted_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_3, PK_0, PK_4));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidThree_sorted_revertSignaturesAmountBelowThreshold() public {
        vm.prank(owner);
        module.removeVerifier(SIGNER_0);
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_0, PK_3));
        expectRevertSignaturesAmountBelowThreshold(0, 2);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidThree_duplicate_revertNotSorted() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_4, PK_3, PK_3));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_threeSignatures_invalidThree_unsorted_revertNotSorted() public {
        vm.prank(owner);
        module.removeVerifier(SIGNER_0);
        bytes memory signatures = signEntry(mockEntry, toArray(PK_3, PK_4, PK_0));
        expectRevertRecoveredSignersNotSorted();
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_revertChainIdNotRemote() public {
        InterchainEntry memory entry = mockEntry;
        entry.srcChainId = DST_CHAIN_ID;
        bytes memory versionedEntry = getVersionedEntry(entry);
        bytes memory signatures = signEntry(entry, toArray(PK_1, PK_0));
        expectRevertChainIdNotRemote(DST_CHAIN_ID);
        verifyRemoteEntry(versionedEntry, signatures);
    }

    function test_verifyRemoteEntry_revertThresholdZero() public {
        // Deploy a module without setting up the threshold
        module = new SynapseModule(address(interchainDB), owner);
        vm.startPrank(owner);
        module.addVerifier(SIGNER_0);
        module.addVerifier(SIGNER_1);
        vm.stopPrank();
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_0));
        vm.expectRevert(ThresholdECDSALib.ThresholdECDSA__ThresholdZero.selector);
        verifyRemoteEntry(mockVersionedEntry, signatures);
    }

    function test_verifyRemoteEntry_revertSignaturesPayloadLengthInvalidTooShort() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_0));
        bytes memory signaturesShort = new bytes(signatures.length - 1);
        for (uint256 i = 0; i < signaturesShort.length; ++i) {
            signaturesShort[i] = signatures[i];
        }
        expectRevertSignaturesPayloadLengthInvalid(signaturesShort.length);
        verifyRemoteEntry(mockVersionedEntry, signaturesShort);
    }

    function test_verifyRemoteEntry_revertSignaturesPayloadLengthInvalidTooLong() public {
        bytes memory signatures = signEntry(mockEntry, toArray(PK_1, PK_0));
        bytes memory signaturesLong = bytes.concat(signatures, bytes1(0x2A));
        expectRevertSignaturesPayloadLengthInvalid(signaturesLong.length);
        verifyRemoteEntry(mockVersionedEntry, signaturesLong);
    }
}
