// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    Header,
    HeaderLib,
    Message,
    MessageLib,
    Tips,
    TipsLib,
    TypedMemView
} from "../../../contracts/libs/Message.sol";

struct RawHeader {
    uint32 origin;
    bytes32 sender;
    uint32 nonce;
    uint32 destination;
    bytes32 recipient;
    uint32 optimisticSeconds;
}
using { CastLib.castToHeader } for RawHeader global;

struct RawTips {
    uint96 notaryTip;
    uint96 broadcasterTip;
    uint96 proverTip;
    uint96 executorTip;
}
using { CastLib.castToTips } for RawTips global;

struct RawMessage {
    RawHeader header;
    RawTips tips;
    bytes body;
}
using { CastLib.castToMessage } for RawMessage global;

library CastLib {
    using HeaderLib for bytes;
    using MessageLib for bytes;
    using TipsLib for bytes;
    using TypedMemView for bytes29;

    /// @notice Prevents this contract from being included in the coverage report
    function testCastLib() external {}

    function castToMessage(RawMessage memory rm)
        internal
        pure
        returns (bytes memory message, Message ptr)
    {
        (bytes memory header, ) = rm.header.castToHeader();
        (bytes memory tips, ) = rm.tips.castToTips();
        message = MessageLib.formatMessage(header, tips, rm.body);
        ptr = message.castToMessage();
    }

    function castToHeader(RawHeader memory rh)
        internal
        pure
        returns (bytes memory header, Header ptr)
    {
        header = HeaderLib.formatHeader({
            _origin: rh.origin,
            _sender: rh.sender,
            _nonce: rh.nonce,
            _destination: rh.destination,
            _recipient: rh.recipient,
            _optimisticSeconds: rh.optimisticSeconds
        });
        ptr = header.castToHeader();
    }

    function castToTips(RawTips memory rt) internal pure returns (bytes memory tips, Tips ptr) {
        tips = TipsLib.formatTips({
            _notaryTip: rt.notaryTip,
            _broadcasterTip: rt.broadcasterTip,
            _proverTip: rt.proverTip,
            _executorTip: rt.executorTip
        });
        ptr = tips.castToTips();
    }
}
