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

    function test_formattedCorrectly(
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
            abi.encodePacked(Tips.TIPS_VERSION, notaryTip, broadcasterTip, proverTip, executorTip),
            "!formatTips"
        );
        // Test formatting checker
        assertTrue(libHarness.isTips(payload), "!isTips");
        // Test getters
        assertEq(
            libHarness.tipsVersion(SynapseTypes.MESSAGE_TIPS, payload),
            Tips.TIPS_VERSION,
            "!tipsVersion"
        );
        assertEq(libHarness.notaryTip(SynapseTypes.MESSAGE_TIPS, payload), notaryTip, "!notaryTip");
        assertEq(
            libHarness.broadcasterTip(SynapseTypes.MESSAGE_TIPS, payload),
            broadcasterTip,
            "!broadcasterTip"
        );
        assertEq(libHarness.proverTip(SynapseTypes.MESSAGE_TIPS, payload), proverTip, "!proverTip");
        assertEq(
            libHarness.executorTip(SynapseTypes.MESSAGE_TIPS, payload),
            executorTip,
            "!executorTip"
        );
        assertEq(libHarness.totalTips(SynapseTypes.MESSAGE_TIPS, payload), totalTips, "!totalTips");
    }

    function test_emptyTips() public {
        bytes memory payload = libHarness.emptyTips();
        assertEq(payload, createTestPayload(), "!formatTips");
        // Check formatting
        test_formattedCorrectly(0, 0, 0, 0);
    }

    function test_isTips_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Tips' first element (uint16 tipsVersion)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isTips(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isTips: short payload"
        );
    }

    function test_isTips_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        assertTrue(libHarness.isTips(createTestPayload()), "!isTips: test payload");
    }

    function test_isTips_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        assertFalse(libHarness.isTips(cutLastByte(createTestPayload())), "!isTips: 1 byte shorter");
    }

    function test_isTips_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        assertFalse(libHarness.isTips(addLastByte(createTestPayload())), "!isTips: 1 byte longer");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_tipsVersion(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_TIPS });
        libHarness.tipsVersion(wrongType, payload);
    }

    function test_wrongTypeRevert_notaryTip(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_TIPS });
        libHarness.notaryTip(wrongType, payload);
    }

    function test_wrongTypeRevert_broadcasterTip(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_TIPS });
        libHarness.broadcasterTip(wrongType, payload);
    }

    function test_wrongTypeRevert_proverTip(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_TIPS });
        libHarness.proverTip(wrongType, payload);
    }

    function test_wrongTypeRevert_executorTip(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_TIPS });
        libHarness.executorTip(wrongType, payload);
    }

    function test_wrongTypeRevert_totalTips(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_TIPS });
        libHarness.totalTips(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatTips(0, 0, 0, 0);
    }
}
