// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyMessageLib} from "../../../contracts/legacy/libs/LegacyMessage.sol";

contract LegacyMessageLibHarness {
    function encodeLegacyMessage(
        address srcSender,
        address dstReceiver,
        uint64 srcNonce,
        bytes memory message
    )
        public
        pure
        returns (bytes memory legacyMsg)
    {
        return LegacyMessageLib.encodeLegacyMessage(srcSender, dstReceiver, srcNonce, message);
    }

    function decodeLegacyMessage(bytes calldata legacyMsg)
        public
        pure
        returns (address srcSender, address dstReceiver, uint64 srcNonce, bytes memory message)
    {
        return LegacyMessageLib.decodeLegacyMessage(legacyMsg);
    }
}
