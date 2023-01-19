// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/TipsHarness.t.sol";

import "../../../contracts/libs/Tips.sol";

// solhint-disable func-name-mixedcase
contract TipsLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    TipsHarness internal libHarness;
    // First element is (uint16 tipsVersion)
    uint8 internal constant FIRST_ELEMENT_BYTES = 16 / 8;

    function setUp() public override {
        super.setUp();
        libHarness = new TipsHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formatTips(
        uint96 notaryTip,
        uint96 broadcasterTip,
        uint96 proverTip,
        uint96 executorTip
    ) public {
        // TODO: Determine if we actually need uint96 for storing tips / totalTips
        uint256 totalTips = uint256(notaryTip) + broadcasterTip + proverTip + executorTip;
        vm.assume(totalTips <= type(uint96).max);
        // Test formatting
        bytes memory payload = libHarness.formatTips(
            notaryTip,
            broadcasterTip,
            proverTip,
            executorTip
        );
        assertEq(
            payload,
            abi.encodePacked(
                TipsLib.TIPS_VERSION,
                notaryTip,
                broadcasterTip,
                proverTip,
                executorTip
            ),
            "!formatTips"
        );
        // Test formatting checker
        checkCastToTips({ payload: payload, isTips: true });
        // Test getters
        assertEq(libHarness.version(payload), TipsLib.TIPS_VERSION, "!tipsVersion");
        assertEq(libHarness.notaryTip(payload), notaryTip, "!notaryTip");
        assertEq(libHarness.broadcasterTip(payload), broadcasterTip, "!broadcasterTip");
        assertEq(libHarness.proverTip(payload), proverTip, "!proverTip");
        assertEq(libHarness.executorTip(payload), executorTip, "!executorTip");
        assertEq(libHarness.totalTips(payload), totalTips, "!totalTips");
    }

    function test_emptyTips() public {
        bytes memory payload = libHarness.emptyTips();
        assertEq(payload, createTestPayload(), "!formatTips");
        // Check formatting
        test_formatTips(0, 0, 0, 0);
    }

    function test_isTips_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Tips' first element (uint16 tipsVersion)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data);
        checkCastToTips({ payload: payload, isTips: false });
    }

    function test_isTips_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        bytes memory payload = createTestPayload();
        checkCastToTips({ payload: payload, isTips: true });
    }

    function test_isTips_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        bytes memory payload = cutLastByte(createTestPayload());
        checkCastToTips({ payload: payload, isTips: false });
    }

    function test_isTips_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        assertFalse(libHarness.isTips(addLastByte(createTestPayload())), "!isTips: 1 byte longer");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToTips(bytes memory payload, bool isTips) public {
        if (isTips) {
            assertTrue(libHarness.isTips(payload), "!isTips: when valid");
            assertEq(libHarness.castToTips(payload), payload, "!castToTips: when valid");
        } else {
            assertFalse(libHarness.isTips(payload), "!isTips: when valid");
            vm.expectRevert("Not a tips payload");
            libHarness.castToTips(payload);
        }
    }

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatTips(0, 0, 0, 0);
    }
}
