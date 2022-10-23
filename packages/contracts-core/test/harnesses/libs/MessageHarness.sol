// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Message } from "../../../contracts/libs/Message.sol";
import { Header } from "../../../contracts/libs/Header.sol";
import { Tips } from "../../../contracts/libs/Tips.sol";

/**
 * @notice Exposes Message methods for testing against golang.
 */
contract MessageHarness {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatMessage(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        // tips params
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        bytes memory _tips = Tips.formatTips(_notaryTip, _broadcasterTip, _proverTip, _executorTip);

        bytes memory _header = Header.formatHeader(
            _origin,
            _sender,
            _nonce,
            _destination,
            _recipient,
            _optimisticSeconds
        );
        return formatMessage(_header, _tips, _messageBody);
    }

    function formatMessage(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        return Message.formatMessage(_header, _tips, _messageBody);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function messageVersion() public pure returns (uint16) {
        return Message.MESSAGE_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return Message.OFFSET_VERSION;
    }

    function offsetHeader() public pure returns (uint256) {
        return Message.OFFSET_HEADER;
    }
}
