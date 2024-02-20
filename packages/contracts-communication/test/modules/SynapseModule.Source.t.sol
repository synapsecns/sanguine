// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";
import {SynapseModuleEvents} from "../../contracts/events/SynapseModuleEvents.sol";
import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {SynapseModule, InterchainEntry, ISynapseModule} from "../../contracts/modules/SynapseModule.sol";

import {GasOracleMock} from "../mocks/GasOracleMock.sol";

import {Test} from "forge-std/Test.sol";

contract SynapseModuleSourceTest is Test, InterchainModuleEvents, SynapseModuleEvents {
    SynapseModule public module;
    GasOracleMock public gasOracle;

    address public interchainDB = makeAddr("InterchainDB");
    address public feeCollector = makeAddr("FeeCollector");
    address public owner = makeAddr("Owner");

    uint256 public constant SRC_CHAIN_ID = 1337;
    uint256 public constant DST_CHAIN_ID = 7331;

    // TODO: this should be configurable
    uint256 public constant EXPECTED_GAS_LIMIT = 100_000;

    uint256 public constant FEE = 100;

    InterchainEntry public mockEntry = InterchainEntry({
        srcChainId: SRC_CHAIN_ID,
        dbNonce: 2,
        srcWriter: bytes32(uint256(3)),
        dataHash: bytes32(uint256(4))
    });

    function setUp() public {
        vm.chainId(SRC_CHAIN_ID);
        module = new SynapseModule(interchainDB, owner);
        gasOracle = new GasOracleMock();
        vm.startPrank(owner);
        module.setGasOracle(address(gasOracle));
        module.setFeeCollector(feeCollector);
        module.addVerifier(address(1));
        module.addVerifier(address(2));
        module.setThreshold(2);
        vm.stopPrank();
        // Mock: gasOracle.estimateTxCostInLocalUnits(DST_CHAIN_ID, *, *) to return FEE
        vm.mockCall(
            address(gasOracle),
            abi.encodeWithSelector(GasOracleMock.estimateTxCostInLocalUnits.selector, DST_CHAIN_ID),
            abi.encode(FEE)
        );
    }

    function mockRequestVerification(uint256 msgValue, InterchainEntry memory entry) internal {
        deal(interchainDB, msgValue);
        vm.prank(interchainDB);
        module.requestVerification{value: msgValue}(DST_CHAIN_ID, entry);
    }

    function encodeAndHashEntry(InterchainEntry memory entry)
        internal
        pure
        returns (bytes memory encodedEntry, bytes32 ethSignedHash)
    {
        encodedEntry = abi.encode(entry);
        ethSignedHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(encodedEntry)));
    }

    function test_setup() public {
        assertEq(module.owner(), owner);
        assertEq(module.INTERCHAIN_DB(), interchainDB);
        assertTrue(module.isVerifier(address(1)));
        assertEq(module.getThreshold(), 2);
        assertEq(module.gasOracle(), address(gasOracle));
    }

    function test_requestVerification_emitsEvent() public {
        (bytes memory encodedEntry, bytes32 ethSignedHash) = encodeAndHashEntry(mockEntry);
        vm.expectEmit(address(module));
        emit VerificationRequested(DST_CHAIN_ID, encodedEntry, ethSignedHash);
        mockRequestVerification(FEE, mockEntry);
    }

    function test_requestVerification_transfersFeeToCollector() public {
        mockRequestVerification(FEE, mockEntry);
        assertEq(feeCollector.balance, FEE);
    }

    function test_requestVerification_feeAboveRequired_emitsEvent() public {
        (bytes memory encodedEntry, bytes32 ethSignedHash) = encodeAndHashEntry(mockEntry);
        vm.expectEmit(address(module));
        emit VerificationRequested(DST_CHAIN_ID, encodedEntry, ethSignedHash);
        mockRequestVerification(FEE + 1, mockEntry);
    }

    function test_requestVerification_feeAboveRequired_transfersFeeToCollector() public {
        mockRequestVerification(FEE + 1, mockEntry);
        assertEq(feeCollector.balance, FEE + 1);
    }

    function test_requestVerification_revert_feeBelowRequired() public {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainModule.InterchainModule__InsufficientFee.selector, FEE - 1, FEE)
        );
        mockRequestVerification(FEE - 1, mockEntry);
    }

    function test_getModuleFee_thresholdTwo() public {
        assertEq(module.getModuleFee(DST_CHAIN_ID), FEE);
    }

    function test_getModuleFee_callsGasOracle_twoSigners() public {
        bytes memory mockedSignatures = new bytes(2 * 65);
        bytes memory remoteCalldata = abi.encodeCall(module.verifyEntry, (abi.encode(mockEntry), mockedSignatures));
        bytes memory expectedCalldata = abi.encodeCall(
            gasOracle.estimateTxCostInLocalUnits, (DST_CHAIN_ID, EXPECTED_GAS_LIMIT, remoteCalldata.length)
        );
        vm.expectCall(address(gasOracle), expectedCalldata);
        module.getModuleFee(DST_CHAIN_ID);
    }

    function test_getModuleFee_callsGasOracle_threeSigners() public {
        vm.prank(owner);
        module.setThreshold(3);
        bytes memory mockedSignatures = new bytes(3 * 65);
        bytes memory remoteCalldata = abi.encodeCall(module.verifyEntry, (abi.encode(mockEntry), mockedSignatures));
        bytes memory expectedCalldata = abi.encodeCall(
            gasOracle.estimateTxCostInLocalUnits, (DST_CHAIN_ID, EXPECTED_GAS_LIMIT, remoteCalldata.length)
        );
        vm.expectCall(address(gasOracle), expectedCalldata);
        module.getModuleFee(DST_CHAIN_ID);
    }
}
