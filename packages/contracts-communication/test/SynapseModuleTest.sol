pragma solidity 0.8.20;

import {Test} from "forge-std/Test.sol";

import {InterchainDB, InterchainEntry, IInterchainDB, InterchainDBEvents} from "../contracts/InterchainDB.sol";

import {SynapseModule, SynapseModuleEvents} from "../contracts/modules/SynapseModule.sol";

contract SynapseModuleTest is Test, SynapseModuleEvents {
    IInterchainDB icDB;
    SynapseModule synapseModule;

    uint256 public constant SRC_CHAIN_ID = 1337;
    uint256 public constant DST_CHAIN_ID = 7331;

    uint256 public constant INITIAL_WRITER_F_NONCE = 1;
    uint256 public constant INITIAL_WRITER_S_NONCE = 2;

    address public contractOwner = makeAddr("Contract Owner");
    address public requestCaller = makeAddr("Request Caller");
    address public writerF = makeAddr("First Writer");
    address public writerS = makeAddr("Second Writer");
    address public notWriter = makeAddr("Not a Writer");

    Account verifierA = makeAccount("Verifier A");
    Account verifierB = makeAccount("Verifier B");
    Account verifierC = makeAccount("Verifier C");
    address[] verifiers = [verifierA.addr, verifierB.addr, verifierC.addr];
    uint256[] verifiersPrivateKeys = [verifierA.key, verifierB.key, verifierC.key];

    function setUp() public {
        vm.startPrank(contractOwner);
        icDB = new InterchainDB();
        synapseModule = new SynapseModule();
        synapseModule.setInterchainDB(address(icDB));
        synapseModule.setVerifiers(verifiers);
        synapseModule.setRequiredThreshold(2);
        vm.stopPrank();

        setupWriterNonce(writerF, INITIAL_WRITER_F_NONCE);
        setupWriterNonce(writerS, INITIAL_WRITER_S_NONCE);
    }

    function setupWriterNonce(address writer, uint256 nonce) internal {
        for (uint256 i = 0; i < nonce; i++) {
            writeEntry(writer, getMockDataHash(writer, i));
        }
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertEq(InterchainEntry memory entry, InterchainEntry memory expected) internal {
        assertEq(entry.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(entry.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(entry.writerNonce, expected.writerNonce, "!writerNonce");
        assertEq(entry.dataHash, expected.dataHash, "!dataHash");
    }

    function getMockDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    function getMockEntry(address writer, uint256 nonce) internal pure returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            srcWriter: addressToBytes32(writer),
            writerNonce: nonce,
            dataHash: getMockDataHash(writer, nonce)
        });
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    function writeEntry(address writer, bytes32 dataHash) internal returns (uint256 writerNonce) {
        vm.prank(writer);
        writerNonce = icDB.writeEntry(dataHash);
    }

    // ═══════════════════════════════════════════════ TESTS: PERMISSIONING ════════════════════════════════════════════════════

    function test_setInterchainDB_notOwner() public {
        vm.expectRevert();
        vm.prank(requestCaller);
        synapseModule.setInterchainDB(address(icDB));
    }

    function test_setInterchainDB_owner() public {
        vm.prank(contractOwner);
        synapseModule.setInterchainDB(address(icDB));
        assertEq(synapseModule.interchainDB(), address(icDB));
    }

    function test_setRequiredThreshold_notOwner() public {
        vm.expectRevert();
        vm.prank(requestCaller);
        synapseModule.setRequiredThreshold(2);
    }

    function test_setRequiredThreshold_owner() public {
        vm.prank(contractOwner);
        synapseModule.setRequiredThreshold(2);
        assertEq(synapseModule.requiredThreshold(), 2);
    }

    function test_setVerifiers_notOwner() public {
        vm.expectRevert();
        vm.prank(requestCaller);
        synapseModule.setVerifiers(verifiers);
    }

    function test_setVerifiers_owner() public {
        vm.prank(contractOwner);
        synapseModule.setVerifiers(verifiers);
        for (uint256 i = 0; i < verifiers.length; i++) {
            assertEq(synapseModule.verifiers(i), verifiers[i]);
        }
    }

    // ═══════════════════════════════════════════════ TESTS: REQUEST VERIFICATION ════════════════════════════════════════════════════
    function test_requestVerification_notInterchainDB() public {
        vm.expectRevert();
        synapseModule.requestVerification(DST_CHAIN_ID, getMockEntry(writerF, 1));
    }

    function test_requestVerification_insufficientFee() public {
        uint256 expectedFee = synapseModule.getModuleFee(DST_CHAIN_ID);
        vm.expectRevert();
        vm.prank(address(icDB));
        synapseModule.requestVerification{value: expectedFee - 1}(DST_CHAIN_ID, getMockEntry(writerF, 1));
    }

    function test_requestVerification_feesPaid() public {
        address executor = address(synapseModule.executor());
        uint256 prevBalance = executor.balance;
        uint256 expectedFee = synapseModule.getModuleFee(DST_CHAIN_ID);
        vm.deal(address(icDB), expectedFee);
        vm.prank(address(icDB));
        synapseModule.requestVerification{value: expectedFee}(DST_CHAIN_ID, getMockEntry(writerF, 1));
        assertEq(executor.balance, prevBalance + expectedFee);
    }

    function test_requestVerification_eventEmit() public {
        uint256 expectedFee = synapseModule.getModuleFee(DST_CHAIN_ID);
        vm.deal(address(icDB), expectedFee);
        vm.prank(address(icDB));
        InterchainEntry memory mockEntry = getMockEntry(writerF, 1);
        vm.expectEmit();
        emit VerificationRequested(DST_CHAIN_ID, mockEntry, keccak256(abi.encode(mockEntry)));
        synapseModule.requestVerification{value: expectedFee}(DST_CHAIN_ID, getMockEntry(writerF, 1));
    }

    // ═══════════════════════════════════════════════ TESTS: VERIFY ENTRY ════════════════════════════════════════════════════

    function test_verifyEntry_notEnoughSignatures() public {
        InterchainEntry memory entry = getMockEntry(writerF, 1);
        vm.expectRevert();
        synapseModule.verifyEntry(entry, new bytes[](0));
    }

    function test_verifyEntry_notEnoughValidSignatures() public {
        InterchainEntry memory entry = getMockEntry(writerF, 1);
        bytes[] memory signatures = new bytes[](verifiers.length);
        uint256 expectedThreshold = synapseModule.requiredThreshold();
        uint256 plannedValidSignatures;
        for (uint256 i = 0; i < verifiers.length; i++) {
            if (plannedValidSignatures < expectedThreshold - 1) {
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(verifiersPrivateKeys[i], keccak256(abi.encode(entry)));
                signatures[i] = abi.encodePacked(r, s, v);
                plannedValidSignatures++;
            } else {
                (uint8 v, bytes32 r, bytes32 s) =
                    vm.sign(verifiersPrivateKeys[i], keccak256(abi.encode("invalid message")));
                signatures[i] = abi.encodePacked(r, s, v);
            }
        }
        vm.expectRevert();
        synapseModule.verifyEntry(entry, signatures);
    }

    function test_verifyEntry_validSignatures() public {
        InterchainEntry memory entry = getMockEntry(writerF, 1);
        bytes[] memory signatures = new bytes[](verifiers.length);
        for (uint256 i = 0; i < verifiers.length; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(verifiersPrivateKeys[i], keccak256(abi.encode(entry)));
            signatures[i] = abi.encodePacked(r, s, v);
        }
        vm.expectEmit();
        emit EntryVerified(entry, keccak256(abi.encode(entry)));
        // set block time to 2 to verify entry at block.timestamp = 2
        vm.warp(2);
        synapseModule.verifyEntry(entry, signatures);
        uint256 moduleVerifiedAt = icDB.readEntry(address(synapseModule), entry);
        assertEq(moduleVerifiedAt, 2);
    }
}
