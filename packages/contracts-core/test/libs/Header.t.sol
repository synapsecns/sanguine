// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Bytes29Test } from "../utils/Bytes29Test.sol";
import { Header } from "../../contracts/libs/Header.sol";
import { SynapseTypes } from "../../contracts/libs/SynapseTypes.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract HeaderTest is Bytes29Test {
    using Header for bytes;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Header for bytes29;

    uint32 internal constant ORIGIN = 1234;
    bytes32 internal constant SENDER = bytes32("sender");
    uint32 internal constant NONCE = 3456;
    uint32 internal constant DESTINATION = 5678;
    bytes32 internal constant RECIPIENT = bytes32("recipient");
    uint32 internal constant OPTIMISTIC_SECONDS = 7890;

    function test_formattedCorrectly() public {
        bytes29 _view = _createTestView();

        assertTrue(_view.isHeader());

        assertEq(_view.headerVersion(), Header.HEADER_VERSION);
        assertEq(_view.origin(), ORIGIN);
        assertEq(_view.sender(), SENDER);
        assertEq(_view.nonce(), NONCE);
        assertEq(_view.destination(), DESTINATION);
        assertEq(_view.recipient(), RECIPIENT);
        assertEq(_view.optimisticSeconds(), OPTIMISTIC_SECONDS);
    }

    function test_incorrectType_headerVersion() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).headerVersion();
    }

    function test_incorrectType_origin() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).origin();
    }

    function test_incorrectType_sender() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).sender();
    }

    function test_incorrectType_nonce() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).nonce();
    }

    function test_incorrectType_destination() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).destination();
    }

    function test_incorrectType_recipient() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).recipient();
    }

    function test_incorrectType_recipientAddress() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).recipientAddress();
    }

    function test_incorrectType_optimisticSeconds() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_HEADER).optimisticSeconds();
    }

    function test_isHeader_incorrectVersion() public {
        bytes memory _header = abi.encodePacked(uint16(0), new bytes(Header.HEADER_LENGTH - 2));
        assert(_header.length == Header.HEADER_LENGTH);
        assertFalse(_header.castToHeader().isHeader());
    }

    function test_isHeader_emptyPayload() public {
        bytes memory _header = bytes("");
        assert(_header.length == 0);
        assertFalse(_header.castToHeader().isHeader());
    }

    function test_isHeader_tooShort() public {
        bytes memory _header = abi.encodePacked(uint16(1), new bytes(Header.HEADER_LENGTH - 3));
        assert(_header.length < Header.HEADER_LENGTH);
        assertFalse(_header.castToHeader().isHeader());
    }

    function test_isHeader_tooLong() public {
        bytes memory _header = abi.encodePacked(uint16(1), new bytes(Header.HEADER_LENGTH - 1));
        assert(_header.length > Header.HEADER_LENGTH);
        assertFalse(_header.castToHeader().isHeader());
    }

    function _createTestView() internal pure override returns (bytes29) {
        bytes memory _header = Header.formatHeader(
            ORIGIN,
            SENDER,
            NONCE,
            DESTINATION,
            RECIPIENT,
            OPTIMISTIC_SECONDS
        );
        return _header.castToHeader();
    }
}
