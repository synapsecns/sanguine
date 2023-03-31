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

import { StateFlag, StateReport, StateReportLib } from "../../../contracts/libs/StateReport.sol";

struct RawHeader {
    uint32 origin;
    bytes32 sender;
    uint32 nonce;
    uint32 destination;
    bytes32 recipient;
    uint32 optimisticSeconds;
}
using { CastLib.castToHeader, CastLib.formatHeader } for RawHeader global;

struct RawTips {
    uint96 notaryTip;
    uint96 broadcasterTip;
    uint96 proverTip;
    uint96 executorTip;
}
using { CastLib.castToTips, CastLib.formatTips } for RawTips global;

struct RawMessage {
    RawHeader header;
    RawTips tips;
    bytes body;
}
using { CastLib.castToMessage, CastLib.formatMessage } for RawMessage global;

struct RawState {
    bytes32 root;
    uint32 origin;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}
using { CastLib.castToState, CastLib.formatState } for RawState global;

struct RawSnapshot {
    RawState[] states;
}
using {
    CastLib.formatStates,
    CastLib.castToRawAttestation,
    CastLib.castToSnapshot,
    CastLib.formatSnapshot
} for RawSnapshot global;

struct RawAttestation {
    bytes32 snapRoot;
    bytes32 agentRoot;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}
using {
    CastLib.castToAttestation,
    CastLib.formatAttestation,
    CastLib.modifyAttestation
} for RawAttestation global;

struct RawAttestationReport {
    uint8 flag;
    RawAttestation attestation;
}
using {
    CastLib.castToAttestationReport,
    CastLib.formatAttestationReport
} for RawAttestationReport global;

struct RawStateReport {
    uint8 flag;
    RawState state;
}
using { CastLib.castToStateReport, CastLib.formatStateReport } for RawStateReport global;

// solhint-disable no-empty-blocks
// solhint-disable ordering
library CastLib {
    using AttestationLib for bytes;
    using AttestationReportLib for bytes;
    using HeaderLib for bytes;
    using MessageLib for bytes;
    using SnapshotLib for bytes;
    using StateLib for bytes;
    using StateReportLib for bytes;
    using TipsLib for bytes;
    using TypedMemView for bytes29;

    /// @notice Prevents this contract from being included in the coverage report
    function testCastLib() external {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               MESSAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatMessage(RawMessage memory rm) internal pure returns (bytes memory msgPayload) {
        bytes memory header = rm.header.formatHeader();
        bytes memory tipsPayload = rm.tips.formatTips();
        return MessageLib.formatMessage(header, tipsPayload, rm.body);
    }

    function castToMessage(RawMessage memory rm) internal pure returns (Message ptr) {
        ptr = rm.formatMessage().castToMessage();
    }

    function formatHeader(RawHeader memory rh) internal pure returns (bytes memory header) {
        header = HeaderLib.formatHeader({
            origin_: rh.origin,
            sender_: rh.sender,
            nonce_: rh.nonce,
            destination_: rh.destination,
            recipient_: rh.recipient,
            optimisticSeconds_: rh.optimisticSeconds
        });
    }

    function castToHeader(RawHeader memory rh) internal pure returns (Header ptr) {
        ptr = rh.formatHeader().castToHeader();
    }

    function formatTips(RawTips memory rt) internal pure returns (bytes memory tipsPayload) {
        tipsPayload = TipsLib.formatTips({
            notaryTip_: rt.notaryTip,
            broadcasterTip_: rt.broadcasterTip,
            proverTip_: rt.proverTip,
            executorTip_: rt.executorTip
        });
    }

    function castToTips(RawTips memory rt) internal pure returns (Tips ptr) {
        ptr = rt.formatTips().castToTips();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                STATE                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatState(RawState memory rs) internal pure returns (bytes memory state) {
        state = StateLib.formatState({
            root_: rs.root,
            origin_: rs.origin,
            nonce_: rs.nonce,
            blockNumber_: rs.blockNumber,
            timestamp_: rs.timestamp
        });
    }

    function castToState(RawState memory rs) internal pure returns (State ptr) {
        ptr = rs.formatState().castToState();
    }

    function formatStateReport(RawStateReport memory rawSR)
        internal
        pure
        returns (bytes memory stateReport)
    {
        // Explicit revert when flag out of range
        require(rawSR.flag <= uint8(type(StateFlag).max), "Flag out of range");
        bytes memory state = rawSR.state.formatState();
        stateReport = StateFlag(rawSR.flag).formatStateReport(state);
    }

    function castToStateReport(RawStateReport memory rawSR)
        internal
        pure
        returns (StateReport ptr)
    {
        ptr = rawSR.formatStateReport().castToStateReport();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               SNAPSHOT                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatStates(RawSnapshot memory rawSnap)
        internal
        pure
        returns (bytes[] memory states)
    {
        states = new bytes[](rawSnap.states.length);
        for (uint256 i = 0; i < rawSnap.states.length; ++i) {
            states[i] = rawSnap.states[i].formatState();
        }
    }

    function castToRawAttestation(
        RawSnapshot memory rawSnap,
        bytes32 agentRoot,
        uint32 nonce,
        uint40 blockNumber,
        uint40 timestamp
    ) internal view returns (RawAttestation memory ra) {
        Snapshot snapshot = rawSnap.castToSnapshot();
        ra.snapRoot = snapshot.root();
        ra.agentRoot = agentRoot;
        ra.nonce = nonce;
        ra.blockNumber = blockNumber;
        ra.timestamp = timestamp;
    }

    function formatSnapshot(RawSnapshot memory rawSnap)
        internal
        view
        returns (bytes memory snapshot)
    {
        State[] memory states = new State[](rawSnap.states.length);
        for (uint256 i = 0; i < rawSnap.states.length; ++i) {
            states[i] = rawSnap.states[i].castToState();
        }
        snapshot = SnapshotLib.formatSnapshot(states);
    }

    function castToSnapshot(RawSnapshot memory rawSnap) internal view returns (Snapshot ptr) {
        ptr = rawSnap.formatSnapshot().castToSnapshot();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ATTESTATION                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestation(RawAttestation memory ra)
        internal
        pure
        returns (bytes memory attestation)
    {
        attestation = AttestationLib.formatAttestation({
            snapRoot_: ra.snapRoot,
            agentRoot_: ra.agentRoot,
            nonce_: ra.nonce,
            blockNumber_: ra.blockNumber,
            timestamp_: ra.timestamp
        });
    }

    function castToAttestation(RawAttestation memory ra) internal pure returns (Attestation ptr) {
        ptr = ra.formatAttestation().castToAttestation();
    }

    function modifyAttestation(RawAttestation memory ra, uint256 mask)
        internal
        pure
        returns (bool isEqual, RawAttestation memory mra)
    {
        // Don't modify the nonce
        mra.nonce = ra.nonce;
        // Check if at least one value was modified by checking last 4 bits
        isEqual = mask & 15 == 0;
        mra.snapRoot = ra.snapRoot ^ bytes32(mask & 1);
        mra.agentRoot = ra.agentRoot ^ bytes32(mask & 2);
        mra.blockNumber = ra.blockNumber ^ uint40(mask & 4);
        mra.timestamp = ra.timestamp ^ uint40(mask & 8);
    }

    function formatAttestationReport(RawAttestationReport memory rawAR)
        internal
        pure
        returns (bytes memory attestationReport)
    {
        // Explicit revert when out of range
        require(rawAR.flag <= uint8(type(AttestationFlag).max), "Flag out of range");
        bytes memory attestation = rawAR.attestation.formatAttestation();
        attestationReport = AttestationFlag(rawAR.flag).formatAttestationReport(attestation);
    }

    function castToAttestationReport(RawAttestationReport memory rawAR)
        internal
        pure
        returns (AttestationReport ptr)
    {
        ptr = rawAR.formatAttestationReport().castToAttestationReport();
    }
}
