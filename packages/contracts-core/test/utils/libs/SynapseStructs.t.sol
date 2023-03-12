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

import { Snapshot, SnapshotLib, State, StateLib } from "../../../contracts/libs/Snapshot.sol";

import { Attestation, AttestationLib } from "../../../contracts/libs/Attestation.sol";

import {
    AttestationFlag,
    AttestationReport,
    AttestationReportLib
} from "../../../contracts/libs/AttestationReport.sol";

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

struct RawState {
    bytes32 root;
    uint32 origin;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}
using { CastLib.castToState } for RawState global;

struct RawSnapshot {
    RawState[] states;
}
using {
    CastLib.castToStateList,
    CastLib.castToRawAttestation,
    CastLib.castToSnapshot
} for RawSnapshot global;

struct RawAttestation {
    bytes32 root;
    uint8 height;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}
using { CastLib.castToAttestation } for RawAttestation global;

struct RawAttestationReport {
    uint8 flag;
    RawAttestation attestation;
}
using { CastLib.castToAttestationReport } for RawAttestationReport global;

library CastLib {
    using AttestationLib for bytes;
    using HeaderLib for bytes;
    using MessageLib for bytes;
    using AttestationReportLib for bytes;
    using SnapshotLib for bytes;
    using StateLib for bytes;
    using TipsLib for bytes;
    using TypedMemView for bytes29;

    /// @notice Prevents this contract from being included in the coverage report
    function testCastLib() external {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               MESSAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                STATE                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToState(RawState memory rs) internal pure returns (bytes memory state, State ptr) {
        state = StateLib.formatState({
            _root: rs.root,
            _origin: rs.origin,
            _nonce: rs.nonce,
            _blockNumber: rs.blockNumber,
            _timestamp: rs.timestamp
        });
        ptr = state.castToState();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               SNAPSHOT                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToStateList(RawSnapshot memory rawSnap)
        internal
        pure
        returns (bytes[] memory states)
    {
        states = new bytes[](rawSnap.states.length);
        for (uint256 i = 0; i < rawSnap.states.length; ++i) {
            (states[i], ) = rawSnap.states[i].castToState();
        }
    }

    function castToRawAttestation(
        RawSnapshot memory rawSnap,
        uint32 nonce,
        uint40 blockNumber,
        uint40 timestamp
    ) internal view returns (RawAttestation memory ra) {
        (, Snapshot snapshot) = rawSnap.castToSnapshot();
        ra.root = snapshot.root();
        ra.height = snapshot.height();
        ra.nonce = nonce;
        ra.blockNumber = blockNumber;
        ra.timestamp = timestamp;
    }

    function castToSnapshot(RawSnapshot memory rawSnap)
        internal
        view
        returns (bytes memory snapshot, Snapshot ptr)
    {
        State[] memory states = new State[](rawSnap.states.length);
        for (uint256 i = 0; i < rawSnap.states.length; ++i) {
            (, states[i]) = rawSnap.states[i].castToState();
        }
        snapshot = SnapshotLib.formatSnapshot(states);
        ptr = snapshot.castToSnapshot();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ATTESTATION                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToAttestation(RawAttestation memory ra)
        internal
        pure
        returns (bytes memory attestation, Attestation ptr)
    {
        attestation = AttestationLib.formatAttestation({
            _root: ra.root,
            _height: ra.height,
            _nonce: ra.nonce,
            _blockNumber: ra.blockNumber,
            _timestamp: ra.timestamp
        });
        ptr = attestation.castToAttestation();
    }

    function castToAttestationReport(RawAttestationReport memory rawAR)
        internal
        pure
        returns (bytes memory attestationReport, AttestationReport ptr)
    {
        // Explicit revert when out of range
        require(rawAR.flag <= uint8(type(AttestationFlag).max), "Flag out of range");
        (bytes memory attestation, ) = rawAR.attestation.castToAttestation();
        attestationReport = AttestationFlag(rawAR.flag).formatAttestationReport(attestation);
        ptr = attestationReport.castToAttestationReport();
    }
}
