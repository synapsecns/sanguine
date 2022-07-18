// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "../../contracts/UpdaterManager.sol";
import { Attestation } from "../../contracts/libs/Attestation.sol";
import { Tips } from "../../contracts/libs/Tips.sol";

contract SynapseTest is Test {
    using Attestation for bytes;

    uint256 updaterPK = 1;
    uint256 fakeUpdaterPK = 2;
    address updater = vm.addr(updaterPK);
    address fakeUpdater = vm.addr(fakeUpdaterPK);
    address signer = vm.addr(3);
    address fakeSigner = vm.addr(4);

    uint32 localDomain = 1500;
    uint32 remoteDomain = 1000;

    uint96 internal constant UPDATER_TIP = 1234;
    uint96 internal constant RELAYER_TIP = 3456;
    uint96 internal constant PROVER_TIP = 5678;
    uint96 internal constant PROCESSOR_TIP = 7890;
    uint96 internal constant TOTAL_TIPS = UPDATER_TIP + RELAYER_TIP + PROVER_TIP + PROCESSOR_TIP;

    function setUp() public virtual {
        vm.label(updater, "updater");
        vm.label(fakeUpdater, "fake updater");
        vm.label(signer, "signer");
        vm.label(fakeSigner, "fake signer");
    }

    function getDefaultTips() internal pure returns (bytes memory) {
        return Tips.formatTips(UPDATER_TIP, RELAYER_TIP, PROVER_TIP, PROCESSOR_TIP);
    }

    function getEmptyTips() internal pure returns (bytes memory) {
        return Tips.emptyTips();
    }

    function signHomeAttestation(
        uint256 privKey,
        uint32 nonce,
        bytes32 root
    ) public returns (bytes memory attestation, bytes memory signature) {
        bytes memory data = Attestation.formatAttestationData(localDomain, nonce, root);
        signature = signMessage(privKey, data);
        attestation = Attestation.formatAttestation(data, signature);
    }

    function signRemoteAttestation(
        uint256 privKey,
        uint32 nonce,
        bytes32 root
    ) public returns (bytes memory attestation, bytes memory signature) {
        bytes memory data = Attestation.formatAttestationData(remoteDomain, nonce, root);
        signature = signMessage(privKey, data);
        attestation = Attestation.formatAttestation(data, signature);
    }

    function signMessage(uint256 privKey, bytes memory message)
        public
        returns (bytes memory signature)
    {
        bytes32 digest = keccak256(message);
        digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", digest));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, digest);
        signature = abi.encodePacked(r, s, v);
    }

    function stringToBytes32(string memory source) public pure returns (bytes32 result) {
        bytes memory tempEmptyStringTest = bytes(source);
        if (tempEmptyStringTest.length == 0) {
            return 0x0;
        }
        assembly {
            result := mload(add(source, 32))
        }
    }

    function addressToBytes32(address addr) public pure returns (bytes32 result) {
        return bytes32(uint256(uint160(addr)));
    }
}

contract SynapseTestWithUpdaterManager is SynapseTest {
    UpdaterManager updaterManager;

    function setUp() public virtual override {
        super.setUp();
        updaterManager = new UpdaterManager(updater);
    }
}
