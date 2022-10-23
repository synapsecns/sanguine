// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "../../contracts/NotaryManager.sol";
import { Attestation } from "../../contracts/libs/Attestation.sol";
import { Report } from "../../contracts/libs/Report.sol";
import { Tips } from "../../contracts/libs/Tips.sol";

contract SynapseTest is Test {
    using Attestation for bytes;

    uint256 internal notaryPK = 1;
    uint256 internal fakeNotaryPK = 2;
    uint256 internal guardPK = 3;
    uint256 internal fakeGuardPK = 4;
    address internal notary = vm.addr(notaryPK);
    address internal fakeNotary = vm.addr(fakeNotaryPK);
    address internal guard = vm.addr(guardPK);
    address internal fakeGuard = vm.addr(fakeGuardPK);

    uint32 internal localDomain = 1500;
    uint32 internal remoteDomain = 1000;

    uint96 internal constant NOTARY_TIP = 1234;
    uint96 internal constant BROADCASTER_TIP = 3456;
    uint96 internal constant PROVER_TIP = 5678;
    uint96 internal constant EXECUTOR_TIP = 7890;
    uint96 internal constant TOTAL_TIPS = NOTARY_TIP + BROADCASTER_TIP + PROVER_TIP + EXECUTOR_TIP;

    function setUp() public virtual {
        vm.label(notary, "notary");
        vm.label(fakeNotary, "fake notary");
        vm.label(guard, "guard");
        vm.label(fakeGuard, "fake guard");
    }

    function getDefaultTips() internal pure returns (bytes memory) {
        return Tips.formatTips(NOTARY_TIP, BROADCASTER_TIP, PROVER_TIP, EXECUTOR_TIP);
    }

    function getFormattedTips(
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip
    ) internal pure returns (bytes memory) {
        return Tips.formatTips(_notaryTip, _broadcasterTip, _proverTip, _executorTip);
    }

    function getEmptyTips() internal pure returns (bytes memory) {
        return Tips.emptyTips();
    }

    // solhint-disable-next-line ordering
    function signOriginAttestation(
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

    function signReport(
        uint256 privKey,
        Report.Flag flag,
        bytes memory attestation
    ) public returns (bytes memory report, bytes memory signature) {
        bytes memory data = Report.formatReportData(flag, attestation);
        signature = signMessage(privKey, data);
        report = Report.formatReport(flag, attestation, signature);
    }

    function signFraudReport(uint256 privKey, bytes memory attestation)
        public
        returns (bytes memory report, bytes memory signature)
    {
        return signReport(privKey, Report.Flag.Fraud, attestation);
    }

    function signValidReport(uint256 privKey, bytes memory attestation)
        public
        returns (bytes memory report, bytes memory signature)
    {
        return signReport(privKey, Report.Flag.Valid, attestation);
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
        // solhint-disable-next-line no-inline-assembly
        assembly {
            result := mload(add(source, 32))
        }
    }

    function addressToBytes32(address addr) public pure returns (bytes32 result) {
        return bytes32(uint256(uint160(addr)));
    }
}

contract SynapseTestWithNotaryManager is SynapseTest {
    NotaryManager internal notaryManager;

    function setUp() public virtual override {
        super.setUp();
        notaryManager = new NotaryManager(notary);
    }
}
