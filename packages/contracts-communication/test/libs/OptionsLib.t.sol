// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

import {OptionsLib, OptionsLibHarness, OptionsV1} from "../harnesses/OptionsLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract OptionsLibTest is Test {
    struct MockOptionsV2 {
        uint256 gasLimit;
        uint256 gasAirdrop;
        bytes32 newField;
    }

    OptionsLibHarness public libHarness;

    function setUp() public {
        libHarness = new OptionsLibHarness();
    }

    function test_encodeOptionsV1Roundtrip(OptionsV1 memory options) public {
        bytes memory encoded = libHarness.encodeOptionsV1(options);
        OptionsV1 memory decoded = libHarness.decodeOptionsV1(encoded);
        assertEq(decoded.gasLimit, options.gasLimit);
        assertEq(decoded.gasAirdrop, options.gasAirdrop);
    }

    function test_decodeOptionsV1_decodesV2(MockOptionsV2 memory options) public {
        bytes memory encoded =
            VersionedPayloadLib.encodeVersionedPayload(OptionsLib.OPTIONS_V1 + 1, abi.encode(options));
        OptionsV1 memory decoded = libHarness.decodeOptionsV1(encoded);
        assertEq(decoded.gasLimit, options.gasLimit);
        assertEq(decoded.gasAirdrop, options.gasAirdrop);
    }

    function test_decodeOptionsV1_revertLowerVersion() public {
        OptionsV1 memory options = OptionsV1(200_000, 100_000);
        uint16 incorrectVersion = OptionsLib.OPTIONS_V1 - 1;
        bytes memory encoded = VersionedPayloadLib.encodeVersionedPayload(incorrectVersion, abi.encode(options));
        vm.expectRevert(abi.encodeWithSelector(OptionsLib.OptionsLib__IncorrectVersion.selector, incorrectVersion));
        libHarness.decodeOptionsV1(encoded);
    }
}
