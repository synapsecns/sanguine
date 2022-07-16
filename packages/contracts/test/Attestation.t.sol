// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { AttestationHarness } from "./harnesses/AttestationHarness.sol";
import { SynapseTest } from "./utils/SynapseTest.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract AttestationTest is SynapseTest {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    AttestationHarness internal harness;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");

    function setUp() public override {
        super.setUp();
        harness = new AttestationHarness();
    }

    function test_formatAttestation() public {
        bytes memory _data = harness.formatAttestationData(domain, nonce, root);
        bytes memory _sig = signMessage(1337, _data);
        bytes memory _view = harness.formatAttestation(domain, nonce, root, _sig);

        assertTrue(harness.isAttestation(_view));

        assertEq(harness.domain(_view), domain);
        assertEq(harness.nonce(_view), nonce);
        assertEq(harness.root(_view), root);
        assertEq(harness.signature(_view), _sig);
    }

    function test_invalidAttestation_tooShort() public {
        // no signature provided
        bytes memory _data = harness.formatAttestationData(domain, nonce, root);
        assertFalse(harness.isAttestation(_data));
    }
}
