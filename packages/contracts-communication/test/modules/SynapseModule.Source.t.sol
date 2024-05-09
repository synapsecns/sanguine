// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ClaimableFeesEvents} from "../../contracts/events/ClaimableFeesEvents.sol";
import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";
import {SynapseModuleEvents} from "../../contracts/events/SynapseModuleEvents.sol";
import {IClaimableFees} from "../../contracts/interfaces/IClaimableFees.sol";
import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";
import {SynapseModule} from "../../contracts/modules/SynapseModule.sol";

import {InterchainEntryLibHarness} from "../harnesses/InterchainEntryLibHarness.sol";
import {VersionedPayloadLibHarness} from "../harnesses/VersionedPayloadLibHarness.sol";
import {SynapseGasOracleMock} from "../mocks/SynapseGasOracleMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseModuleSourceTest is Test, ClaimableFeesEvents, InterchainModuleEvents, SynapseModuleEvents {
    InterchainEntryLibHarness public entryLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    SynapseModule public module;
    SynapseGasOracleMock public gasOracle;

    address public interchainDB = makeAddr("InterchainDB");
    address public feeRecipient = makeAddr("FeeRecipient");
    address public owner = makeAddr("Owner");
    address public claimer = makeAddr("Claimer");

    uint64 public constant SRC_CHAIN_ID = 1337;
    uint64 public constant DST_CHAIN_ID = 7331;

    uint16 public constant MOCK_DB_VERSION = 42;

    uint256 public constant DEFAULT_GAS_LIMIT = 100_000;

    uint256 public constant FEE = 100;

    InterchainEntry public mockEntry =
        InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: 2, entryValue: bytes32(uint256(3))});
    bytes public mockModuleData = "";

    function setUp() public {
        vm.chainId(SRC_CHAIN_ID);
        module = new SynapseModule(interchainDB, owner);
        gasOracle = new SynapseGasOracleMock();
        entryLibHarness = new InterchainEntryLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
        vm.startPrank(owner);
        module.setGasOracle(address(gasOracle));
        module.setFeeRecipient(feeRecipient);
        module.addVerifier(address(1));
        module.addVerifier(address(2));
        module.setThreshold(2);
        vm.stopPrank();
        // Mock: gasOracle.estimateTxCostInLocalUnits(DST_CHAIN_ID, *, *) to return FEE
        vm.mockCall(
            address(gasOracle),
            abi.encodeWithSelector(SynapseGasOracleMock.estimateTxCostInLocalUnits.selector, DST_CHAIN_ID),
            abi.encode(FEE)
        );
    }

    function getVersionedEntry(InterchainEntry memory entry) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(MOCK_DB_VERSION, entryLibHarness.encodeEntry(entry));
    }

    function requestEntryVerification(uint256 msgValue, uint64 dbNonce, bytes memory versionedEntry) internal {
        deal(interchainDB, msgValue);
        vm.prank(interchainDB);
        module.requestEntryVerification{value: msgValue}(DST_CHAIN_ID, dbNonce, versionedEntry);
    }

    function encodeAndHashEntry(InterchainEntry memory entry)
        internal
        view
        returns (bytes memory encodedEntry, bytes32 ethSignedHash)
    {
        bytes memory versionedEntry =
            payloadLibHarness.encodeVersionedPayload(MOCK_DB_VERSION, entryLibHarness.encodeEntry(entry));
        encodedEntry = abi.encode(versionedEntry, mockModuleData);
        ethSignedHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(encodedEntry)));
    }

    function test_setup() public view {
        assertEq(module.owner(), owner);
        assertEq(module.INTERCHAIN_DB(), interchainDB);
        assertTrue(module.isVerifier(address(1)));
        assertEq(module.getThreshold(), 2);
        assertEq(module.gasOracle(), address(gasOracle));
    }

    function test_requestVerification_emitsEvent() public {
        (bytes memory encodedEntry, bytes32 ethSignedHash) = encodeAndHashEntry(mockEntry);
        bytes memory versionedEntry = getVersionedEntry(mockEntry);
        vm.expectEmit(address(module));
        emit EntryVerificationRequested(DST_CHAIN_ID, encodedEntry, ethSignedHash);
        requestEntryVerification(FEE, mockEntry.dbNonce, versionedEntry);
    }

    function test_requestVerification_accumulatesFee() public {
        deal(address(module), 5 ether);
        bytes memory versionedEntry = getVersionedEntry(mockEntry);
        requestEntryVerification(FEE, mockEntry.dbNonce, versionedEntry);
        assertEq(address(module).balance, 5 ether + FEE);
    }

    function test_requestVerification_feeAboveRequired_emitsEvent() public {
        (bytes memory encodedEntry, bytes32 ethSignedHash) = encodeAndHashEntry(mockEntry);
        bytes memory versionedEntry = getVersionedEntry(mockEntry);
        vm.expectEmit(address(module));
        emit EntryVerificationRequested(DST_CHAIN_ID, encodedEntry, ethSignedHash);
        requestEntryVerification(FEE + 1, mockEntry.dbNonce, versionedEntry);
    }

    function test_requestVerification_feeAboveRequired_accumulatesFee() public {
        deal(address(module), 5 ether);
        bytes memory versionedEntry = getVersionedEntry(mockEntry);
        requestEntryVerification(FEE + 1, mockEntry.dbNonce, versionedEntry);
        assertEq(address(module).balance, 5 ether + FEE + 1);
    }

    function test_requestVerification_revert_feeBelowRequired() public {
        bytes memory versionedEntry = getVersionedEntry(mockEntry);
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainModule.InterchainModule__FeeAmountBelowMin.selector, FEE - 1, FEE)
        );
        requestEntryVerification(FEE - 1, mockEntry.dbNonce, versionedEntry);
    }

    function test_claimFees_zeroClaimFee_emitsEvent() public {
        deal(address(module), 5 ether);
        vm.expectEmit(address(module));
        emit FeesClaimed(feeRecipient, 5 ether, claimer, 0);
        vm.prank(claimer);
        module.claimFees();
    }

    function test_claimFees_zeroClaimFee_distributesFees() public {
        deal(address(module), 5 ether);
        vm.prank(claimer);
        module.claimFees();
        assertEq(feeRecipient.balance, 5 ether);
        assertEq(claimer.balance, 0);
    }

    function test_claimFees_zeroClaimFee_stateChanges() public {
        deal(address(module), 5 ether);
        assertEq(module.getClaimableAmount(), 5 ether);
        assertEq(module.getClaimerFraction(), 0);
        assertEq(module.getClaimerReward(), 0);
        assertEq(module.getFeeRecipient(), feeRecipient);
        vm.prank(claimer);
        module.claimFees();
        assertEq(module.getClaimableAmount(), 0);
        assertEq(module.getClaimerFraction(), 0);
        assertEq(module.getClaimerReward(), 0);
        assertEq(module.getFeeRecipient(), feeRecipient);
    }

    function test_claimFees_zeroClaimFee_revert_FeeRecipientZeroAddress() public {
        SynapseModule freshModule = new SynapseModule(interchainDB, address(this));
        deal(address(freshModule), 5 ether);
        vm.expectRevert(abi.encodeWithSelector(IClaimableFees.ClaimableFees__FeeRecipientZeroAddress.selector));
        vm.prank(claimer);
        freshModule.claimFees();
    }

    function test_claimFees_zeroClaimFee_revert_noFeesToClaim() public {
        vm.expectRevert(abi.encodeWithSelector(IClaimableFees.ClaimableFees__FeeAmountZero.selector));
        vm.prank(claimer);
        module.claimFees();
    }

    function test_claimFees_nonZeroClaimFee_emitsEvent() public {
        // Set claim fee to 0.1%
        vm.prank(owner);
        module.setClaimerFraction(0.001e18);
        deal(address(module), 5 ether);
        vm.expectEmit(address(module));
        emit FeesClaimed(feeRecipient, 4.995 ether, claimer, 0.005 ether);
        vm.prank(claimer);
        module.claimFees();
    }

    function test_claimFees_nonZeroClaimFee_distributesFees() public {
        // Set claim fee to 0.1%
        vm.prank(owner);
        module.setClaimerFraction(0.001e18);
        deal(address(module), 5 ether);
        vm.prank(claimer);
        module.claimFees();
        assertEq(feeRecipient.balance, 4.995 ether);
        assertEq(claimer.balance, 0.005 ether);
    }

    function test_claimFees_nonZeroClaimFee_stateChanges() public {
        // Set claim fee to 0.1%
        vm.prank(owner);
        module.setClaimerFraction(0.001e18);
        deal(address(module), 5 ether);
        assertEq(module.getClaimableAmount(), 5 ether);
        assertEq(module.getClaimerFraction(), 0.001e18);
        assertEq(module.getClaimerReward(), 0.005 ether);
        assertEq(module.getFeeRecipient(), feeRecipient);
        vm.prank(claimer);
        module.claimFees();
        assertEq(module.getClaimableAmount(), 0);
        assertEq(module.getClaimerFraction(), 0.001e18);
        assertEq(module.getClaimerReward(), 0);
        assertEq(module.getFeeRecipient(), feeRecipient);
    }

    function test_claimFees_nonZeroClaimFee_revert_FeeRecipientZeroAddress() public {
        SynapseModule freshModule = new SynapseModule(interchainDB, address(this));
        // Set claim fee to 0.1%
        freshModule.setClaimerFraction(0.001e18);
        deal(address(freshModule), 5 ether);
        vm.expectRevert(abi.encodeWithSelector(IClaimableFees.ClaimableFees__FeeRecipientZeroAddress.selector));
        vm.prank(claimer);
        freshModule.claimFees();
    }

    function test_claimFees_nonZeroClaimFee_revert_noFeesToClaim() public {
        // Set claim fee to 0.1%
        vm.prank(owner);
        module.setClaimerFraction(0.001e18);
        vm.expectRevert(abi.encodeWithSelector(IClaimableFees.ClaimableFees__FeeAmountZero.selector));
        vm.prank(claimer);
        module.claimFees();
    }

    function test_getModuleFee_thresholdTwo() public view {
        // TODO: dbNonce
        assertEq(module.getModuleFee(DST_CHAIN_ID, 0), FEE);
    }

    function test_getModuleFee_callsGasOracle_gasLimitDefault_twoSigners() public {
        bytes memory mockedSignatures = new bytes(2 * 65);
        bytes memory remoteCalldata =
            abi.encodeCall(module.verifyRemoteEntry, (abi.encode(mockEntry), mockedSignatures));
        bytes memory expectedCalldata = abi.encodeCall(
            gasOracle.estimateTxCostInLocalUnits, (DST_CHAIN_ID, DEFAULT_GAS_LIMIT, remoteCalldata.length)
        );
        vm.expectCall(address(gasOracle), expectedCalldata);
        // TODO: dbNonce
        module.getModuleFee(DST_CHAIN_ID, 0);
    }

    function test_getModuleFee_callsGasOracle_gasLimitDefault_threeSigners() public {
        vm.prank(owner);
        module.setThreshold(3);
        bytes memory mockedSignatures = new bytes(3 * 65);
        bytes memory remoteCalldata =
            abi.encodeCall(module.verifyRemoteEntry, (abi.encode(mockEntry), mockedSignatures));
        bytes memory expectedCalldata = abi.encodeCall(
            gasOracle.estimateTxCostInLocalUnits, (DST_CHAIN_ID, DEFAULT_GAS_LIMIT, remoteCalldata.length)
        );
        vm.expectCall(address(gasOracle), expectedCalldata);
        // TODO: dbNonce
        module.getModuleFee(DST_CHAIN_ID, 0);
    }

    function test_getModuleFee_callsGasOracle_gasLimitSet_twoSigners() public {
        vm.prank(owner);
        module.setVerifyGasLimit(DST_CHAIN_ID, 200_000);
        bytes memory mockedSignatures = new bytes(2 * 65);
        bytes memory remoteCalldata =
            abi.encodeCall(module.verifyRemoteEntry, (abi.encode(mockEntry), mockedSignatures));
        bytes memory expectedCalldata =
            abi.encodeCall(gasOracle.estimateTxCostInLocalUnits, (DST_CHAIN_ID, 200_000, remoteCalldata.length));
        vm.expectCall(address(gasOracle), expectedCalldata);
        // TODO: dbNonce
        module.getModuleFee(DST_CHAIN_ID, 0);
    }

    function test_getModuleFee_callsGasOracle_gasLimitSet_threeSigners() public {
        vm.prank(owner);
        module.setThreshold(3);
        vm.prank(owner);
        module.setVerifyGasLimit(DST_CHAIN_ID, 200_000);
        bytes memory mockedSignatures = new bytes(3 * 65);
        bytes memory remoteCalldata =
            abi.encodeCall(module.verifyRemoteEntry, (abi.encode(mockEntry), mockedSignatures));
        bytes memory expectedCalldata =
            abi.encodeCall(gasOracle.estimateTxCostInLocalUnits, (DST_CHAIN_ID, 200_000, remoteCalldata.length));
        vm.expectCall(address(gasOracle), expectedCalldata);
        // TODO: dbNonce
        module.getModuleFee(DST_CHAIN_ID, 0);
    }

    function test_getVerifyGasLimit_default() public view {
        assertEq(module.getVerifyGasLimit(DST_CHAIN_ID), DEFAULT_GAS_LIMIT);
    }
}
