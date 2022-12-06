// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/AuthHarness.t.sol";

// solhint-disable func-name-mixedcase
contract AuthLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint256 internal constant MAX_PRIV_KEY =
        115792089237316195423570985008687907852837564279074904382605163141518161494337;
    bytes internal constant REVERT_WRONG_SIG_LENGTH = "Not a signature";
    bytes internal constant TEST_MESSAGE = "Nothing to see here, please disperse";
    AuthHarness internal libHarness;

    function setUp() public override {
        super.setUp();
        libHarness = new AuthHarness();
    }

    function test_toEthSignedMessageHash() public {
        assertEq(
            libHarness.toEthSignedMessageHash(TEST_MESSAGE),
            _createDigest(),
            "!toEthSignedMessageHash"
        );
    }

    function test_recoverSigner(uint256 privKey) public {
        vm.assume(privKey != 0);
        vm.assume(privKey < MAX_PRIV_KEY);
        address signer = vm.addr(privKey);
        bytes memory signature = signMessage(privKey, TEST_MESSAGE);
        assertEq(libHarness.recoverSigner(_createDigest(), signature), signer, "!recoverSigner");
    }

    function test_recoverSigner_revert_shortSignature() public {
        bytes memory signature = signMessage(1, TEST_MESSAGE);
        // Cut 1 byte from signature
        signature = signature
            .ref(0)
            .slice({ _index: 0, _len: signature.length - 1, newType: 0 })
            .clone();
        vm.expectRevert(REVERT_WRONG_SIG_LENGTH);
        libHarness.recoverSigner(_createDigest(), signature);
    }

    function test_recoverSigner_revert_longSignature() public {
        bytes memory signature = signMessage(1, TEST_MESSAGE);
        // Add 1 byte to signature
        signature = bytes.concat(signature, bytes1(0));
        vm.expectRevert(REVERT_WRONG_SIG_LENGTH);
        libHarness.recoverSigner(_createDigest(), signature);
    }

    function _createDigest() internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(TEST_MESSAGE))
            );
    }
}
