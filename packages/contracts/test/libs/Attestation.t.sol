// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseTest } from "../utils/SynapseTest.sol";
import { Bytes29Test } from "../utils/Bytes29Test.sol";
import { Attestation } from "../../contracts/libs/Attestation.sol";
import { SynapseTypes } from "../../contracts/libs/SynapseTypes.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract AttestationTest is SynapseTest, Bytes29Test {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");

    uint256 internal constant SIGNER_PK = 1337;

    function test_formattedCorrectly() public {
        bytes29 _view = _createTestView();
        assertTrue(_view.isAttestation());

        assertEq(_view.attestedDomain(), domain);
        assertEq(_view.attestedNonce(), nonce);
        assertEq(_view.attestedRoot(), root);

        bytes memory data = abi.encodePacked(domain, nonce, root);
        assertEq(_view.attestationData().clone(), data);
        bytes memory sig = signMessage(SIGNER_PK, data);
        assertEq(_view.notarySignature().clone(), sig);
    }

    function test_isAttestation_tooShort() public {
        // no signature provided
        bytes memory _data = Attestation.formatAttestationData(domain, nonce, root);
        assertFalse(_data.ref(0).isAttestation());
    }

    function test_incorrectType_attestedDomain() public {
        _prepareMistypedTest(SynapseTypes.ATTESTATION).attestedDomain();
    }

    function test_incorrectType_attestedNonce() public {
        _prepareMistypedTest(SynapseTypes.ATTESTATION).attestedNonce();
    }

    function test_incorrectType_attestationData() public {
        _prepareMistypedTest(SynapseTypes.ATTESTATION).attestationData();
    }

    function test_incorrectType_notarySignature() public {
        _prepareMistypedTest(SynapseTypes.ATTESTATION).notarySignature();
    }

    function test_incorrectType_attestedRoot() public {
        _prepareMistypedTest(SynapseTypes.ATTESTATION).attestedRoot();
    }

    function _createTestView() internal override returns (bytes29 _view) {
        bytes memory data = Attestation.formatAttestationData(domain, nonce, root);
        bytes memory sig = signMessage(SIGNER_PK, data);
        bytes memory attestation = Attestation.formatAttestation(data, sig);
        _view = attestation.castToAttestation();
    }
}
