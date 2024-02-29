// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ThresholdECDSALib, ThresholdECDSALibHarness} from "../harnesses/ThresholdECDSALibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
// solhint-disable var-name-mixedcase
contract ThresholdECDSALibTest is Test {
    ThresholdECDSALibHarness public libHarness;

    uint256 public constant PK_0 = 1000;
    uint256 public constant PK_1 = 2000;
    uint256 public constant PK_2 = 3000;
    uint256 public constant PK_3 = 4000;

    address public constant SIGNER_0 = 0x7F1d642DbfD62aD4A8fA9810eA619707d09825D0;
    address public constant SIGNER_1 = 0x5793e629c061e7FD642ab6A1b4d552CeC0e2D606;
    address public constant SIGNER_2 = 0xf6c0eB696e44d15E8dceb3B63A6535e469Be6C62;
    address public constant SIGNER_3 = 0xAD1117CAB797E37CAB0Eee8Ca7C30bD2452Ef2a3;

    bytes32 public constant HASH_0 = keccak256("Some data");

    bytes public sig_0_0 = encodeSignature(PK_0, HASH_0);
    bytes public sig_1_0 = encodeSignature(PK_1, HASH_0);
    bytes public sig_2_0 = encodeSignature(PK_2, HASH_0);
    bytes public sig_3_0 = encodeSignature(PK_3, HASH_0);

    function setUp() public {
        libHarness = new ThresholdECDSALibHarness();
        libHarness.addSigner(SIGNER_0);
        libHarness.addSigner(SIGNER_1);
        libHarness.addSigner(SIGNER_2);
        // Set initial threshold
        libHarness.modifyThreshold(2);
    }

    function encodeSignature(uint256 pk, bytes32 digest) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        // 1-byte v is encoded last
        return abi.encodePacked(r, s, v);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function expectAlreadySignerError(address account) internal {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__AlreadySigner.selector, account));
    }

    function expectIncorrectSignaturesLengthError(uint256 length) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__IncorrectSignaturesLength.selector, length)
        );
    }

    function expectInvalidSignatureError(bytes memory signature) internal {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__InvalidSignature.selector, signature));
    }

    function expectNotEnoughSignaturesError(uint256 threshold) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__NotEnoughSignatures.selector, threshold)
        );
    }

    function expectNotSignerError(address account) internal {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__NotSigner.selector, account));
    }

    function expectRecoveredSignersNotSortedError() internal {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__RecoveredSignersNotSorted.selector));
    }

    function expectZeroAddressError() internal {
        vm.expectRevert(ThresholdECDSALib.ThresholdECDSA__ZeroAddress.selector);
    }

    function expectZeroThresholdError() internal {
        vm.expectRevert(abi.encodeWithSelector(ThresholdECDSALib.ThresholdECDSA__ZeroThreshold.selector));
    }

    // ═══════════════════════════════════════════════════ TESTS ═══════════════════════════════════════════════════════

    function test_pks() public {
        assertEq(SIGNER_0, vm.addr(PK_0));
        assertEq(SIGNER_1, vm.addr(PK_1));
        assertEq(SIGNER_2, vm.addr(PK_2));
        assertEq(SIGNER_3, vm.addr(PK_3));
    }

    function test_addSigner_addsSigner() public {
        libHarness.addSigner(SIGNER_3);
        assertTrue(libHarness.isSigner(SIGNER_3));
    }

    function test_addSigner_expandsList() public {
        libHarness.addSigner(SIGNER_3);
        address[] memory signers = libHarness.getSigners();
        assertEq(signers.length, 4);
        assertEq(signers[0], SIGNER_0);
        assertEq(signers[1], SIGNER_1);
        assertEq(signers[2], SIGNER_2);
        assertEq(signers[3], SIGNER_3);
    }

    function test_addSigner_doesNotModifyThreshold_wasOverThreshold() public {
        libHarness.addSigner(SIGNER_3);
        assertEq(libHarness.getThreshold(), 2);
    }

    function test_addSigner_doesNotModifyThreshold_wasUnderThreshold() public {
        libHarness.modifyThreshold(5);
        libHarness.addSigner(SIGNER_3);
        assertEq(libHarness.getThreshold(), 5);
    }

    function test_addSigner_doesNotModifyThreshold_wasAtThreshold() public {
        libHarness.modifyThreshold(3);
        libHarness.addSigner(SIGNER_3);
        assertEq(libHarness.getThreshold(), 3);
    }

    function test_removeSigner_removesSigner() public {
        libHarness.removeSigner(SIGNER_0);
        assertFalse(libHarness.isSigner(SIGNER_0));
    }

    function test_removeSigner_shrinksList() public {
        libHarness.removeSigner(SIGNER_0);
        address[] memory signers = libHarness.getSigners();
        // SIGNER_0 is removed, SIGNER_2 takes its place
        assertEq(signers.length, 2);
        assertEq(signers[0], SIGNER_2);
        assertEq(signers[1], SIGNER_1);
    }

    function test_removeSigner_doesNotModifyThreshold_wasOverThreshold() public {
        libHarness.removeSigner(SIGNER_0);
        assertEq(libHarness.getThreshold(), 2);
    }

    function test_removeSigner_doesNotModifyThreshold_wasUnderThreshold() public {
        libHarness.modifyThreshold(5);
        libHarness.removeSigner(SIGNER_0);
        assertEq(libHarness.getThreshold(), 5);
    }

    function test_removeSigner_doesNotModifyThreshold_wasAtThreshold() public {
        libHarness.modifyThreshold(3);
        libHarness.removeSigner(SIGNER_0);
        assertEq(libHarness.getThreshold(), 3);
    }

    function test_modifyThreshold_changesThreshold_moreThanSigners() public {
        libHarness.modifyThreshold(4);
        assertEq(libHarness.getThreshold(), 4);
    }

    function test_modifyThreshold_changesThreshold_lessThanSigners() public {
        libHarness.modifyThreshold(1);
        assertEq(libHarness.getThreshold(), 1);
    }

    function test_modifyThreshold_changesThreshold_sameAsSigners() public {
        libHarness.modifyThreshold(3);
        assertEq(libHarness.getThreshold(), 3);
    }

    function test_isSigner_existingSigner() public {
        assertTrue(libHarness.isSigner(SIGNER_0));
        assertTrue(libHarness.isSigner(SIGNER_1));
        assertTrue(libHarness.isSigner(SIGNER_2));
    }

    function test_isSigner_nonExistentSigner() public {
        assertFalse(libHarness.isSigner(SIGNER_3));
    }

    function test_getSigners() public {
        address[] memory signers = libHarness.getSigners();
        assertEq(signers.length, 3);
        assertEq(signers[0], SIGNER_0);
        assertEq(signers[1], SIGNER_1);
        assertEq(signers[2], SIGNER_2);
    }

    function test_getThreshold() public {
        assertEq(libHarness.getThreshold(), 2);
    }

    // ══════════════════════════════════════════════ TESTS: REVERTS ═══════════════════════════════════════════════════

    function test_addSigner_revert_alreadySigner() public {
        expectAlreadySignerError(SIGNER_0);
        libHarness.addSigner(SIGNER_0);
        expectAlreadySignerError(SIGNER_1);
        libHarness.addSigner(SIGNER_1);
        expectAlreadySignerError(SIGNER_2);
        libHarness.addSigner(SIGNER_2);
    }

    function test_addSigner_revert_zeroAddress() public {
        expectZeroAddressError();
        libHarness.addSigner(address(0));
    }

    function test_removeSigner_revert_notSigner() public {
        expectNotSignerError(SIGNER_3);
        libHarness.removeSigner(SIGNER_3);
    }

    function test_modifyThreshold_revert_zeroThreshold() public {
        expectZeroThresholdError();
        libHarness.modifyThreshold(0);
    }

    // ═════════════════════════════════════════ TESTS: VERIFY SIGNED HASH ═════════════════════════════════════════════

    // Signers order sorted by their address:
    // SIGNER_1, SIGNER_0, SIGNER_3, SIGNER_2

    // Should not revert if at least `threshold` signatures are found, and
    // the recovered signers are sorted in ascending order.

    function test_verifySignedHash_providedUnderThreshold_revert_sorted_allSigners() public {
        libHarness.modifyThreshold(3);
        expectNotEnoughSignaturesError(3);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0, sig_2_0));
    }

    function test_verifySignedHash_providedUnderThreshold_revert_sorted_hasNonSigners() public {
        libHarness.modifyThreshold(3);
        expectNotEnoughSignaturesError(3);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0, sig_3_0));
    }

    function test_verifySignedHash_providedUnderThreshold_revert_unsorted_allSigners() public {
        libHarness.modifyThreshold(3);
        expectNotEnoughSignaturesError(3);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_2_0, sig_0_0));
    }

    function test_verifySignedHash_providedUnderThreshold_revert_unsorted_hasNonSigners() public {
        libHarness.modifyThreshold(3);
        expectNotEnoughSignaturesError(3);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_3_0, sig_0_0));
    }

    function test_verifySignedHash_providedUnderThreshold_revert_unsorted_hasDuplicates() public {
        libHarness.modifyThreshold(3);
        expectNotEnoughSignaturesError(3);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0, sig_0_0));
    }

    function test_verifySignedHash_providedExactlyThreshold_sorted_allSigners() public view {
        // Should not revert
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_0_0));
    }

    function test_verifySignedHash_providedExactlyThreshold_revert_sorted_hasNonSigners() public {
        expectNotEnoughSignaturesError(2);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_3_0));
    }

    function test_verifySignedHash_providedExactlyThreshold_revert_unsorted_allSigners() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0, sig_1_0));
    }

    function test_verifySignedHash_providedExactlyThreshold_revert_unsorted_hasNonSigners() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_3_0, sig_1_0));
    }

    function test_verifySignedHash_providedExactlyThreshold_revert_unsorted_hasDuplicates() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0, sig_0_0));
    }

    function test_verifySignedHash_providedOverThreshold_sorted_allSigners() public view {
        // Should not revert
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_0_0, sig_2_0));
    }

    function test_verifySignedHash_providedOverThreshold_sorted_hasNonSigners_enoughSigners() public view {
        // Should not revert
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_3_0, sig_2_0));
    }

    function test_verifySignedHash_providedOverThreshold_revert_sorted_hasNonSigners_notEnoughSigners() public {
        libHarness.removeSigner(SIGNER_2);
        expectNotEnoughSignaturesError(2);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_3_0, sig_2_0));
    }

    function test_verifySignedHash_providedOverThreshold_revert_sorted_hasDuplicates() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_0_0, sig_0_0));
    }

    function test_verifySignedHash_providedOverThreshold_revert_unsorted_allSigners() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_2_0, sig_0_0));
    }

    function test_verifySignedHash_providedOverThreshold_revert_unsorted_hasNonSigners_enoughSigners() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0, sig_3_0, sig_1_0));
    }

    function test_verifySignedHash_providedOverThreshold_revert_unsorted_hasNonSigners_notEnoughSigners() public {
        libHarness.removeSigner(SIGNER_2);
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_2_0, sig_3_0, sig_1_0));
    }

    function test_verifySignedHash_providedOverThreshold_revert_unsorted_hasDuplicates() public {
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_0_0, sig_0_0));
        expectRecoveredSignersNotSortedError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, sig_1_0, sig_0_0));
    }

    function test_verifySignedHash_revert_zeroThreshold_signerSignature() public {
        // Set up a new harness without setting up the threshold
        libHarness = new ThresholdECDSALibHarness();
        libHarness.addSigner(SIGNER_0);
        expectZeroThresholdError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_0_0));
    }

    function test_verifySignedHash_revert_zeroThreshold_notSignerSignature() public {
        // Set up a new harness without setting up the threshold
        libHarness = new ThresholdECDSALibHarness();
        libHarness.addSigner(SIGNER_0);
        expectZeroThresholdError();
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0));
    }

    function test_verifySignedHash_revert_emptySignatures() public {
        expectIncorrectSignaturesLengthError(0);
        libHarness.verifySignedHash(HASH_0, new bytes(0));
    }

    function test_verifySignedHash_revert_incorrectSignaturesLength(uint256 length) public {
        length = bound(length, 0, 1000);
        vm.assume(length % 65 != 0);
        expectIncorrectSignaturesLengthError(length);
        libHarness.verifySignedHash(HASH_0, new bytes(length));
    }

    function test_verifySignedHash_revert_invalidSignature() public {
        bytes memory corruptSig0 = sig_0_0;
        corruptSig0[64] = 0xFF;
        expectInvalidSignatureError(corruptSig0);
        libHarness.verifySignedHash(HASH_0, bytes.concat(sig_1_0, corruptSig0));
    }
}
