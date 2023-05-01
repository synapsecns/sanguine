// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString, CallData, MemView, MemViewLib} from "../../../contracts/libs/ByteString.sol";

import {BaseMessage, BaseMessageLib, Tips, TipsLib} from "../../../contracts/libs/BaseMessage.sol";
import {ChainGas, GasData, GasDataLib} from "../../../contracts/libs/GasData.sol";
import {Header, HeaderLib, Message, MessageFlag, MessageLib} from "../../../contracts/libs/Message.sol";
import {Number, NumberLib} from "../../../contracts/libs/Number.sol";
import {Receipt, ReceiptBody, ReceiptLib} from "../../../contracts/libs/Receipt.sol";
import {ReceiptFlag, ReceiptReport, ReceiptReportLib} from "../../../contracts/libs/ReceiptReport.sol";
import {Request, RequestLib} from "../../../contracts/libs/Request.sol";

import {Snapshot, SnapshotLib, SNAPSHOT_MAX_STATES, State, StateLib} from "../../../contracts/libs/Snapshot.sol";

import {Attestation, AttestationLib} from "../../../contracts/libs/Attestation.sol";

import {AttestationFlag, AttestationReport, AttestationReportLib} from "../../../contracts/libs/AttestationReport.sol";

import {StateFlag, StateReport, StateReportLib} from "../../../contracts/libs/StateReport.sol";

struct RawHeader {
    uint32 origin;
    uint32 nonce;
    uint32 destination;
    uint32 optimisticPeriod;
}

using CastLib for RawHeader global;

struct RawRequest {
    uint96 gasDrop;
    uint64 gasLimit;
}

using CastLib for RawRequest global;

struct RawTips {
    uint64 summitTip;
    uint64 attestationTip;
    uint64 executionTip;
    uint64 deliveryTip;
}

using CastLib for RawTips global;

struct RawReceiptBody {
    uint32 origin;
    uint32 destination;
    bytes32 messageHash;
    bytes32 snapshotRoot;
    uint8 stateIndex;
    address attNotary;
    address firstExecutor;
    address finalExecutor;
}

using CastLib for RawReceiptBody global;

// RawReceipt name is already taken in forge-std
struct RawExecReceipt {
    RawReceiptBody body;
    RawTips tips;
}

using CastLib for RawExecReceipt global;

struct RawReceiptReport {
    uint8 flag;
    RawReceiptBody body;
}

using CastLib for RawReceiptReport global;

struct RawCallData {
    bytes4 selector;
    bytes args;
}

using CastLib for RawCallData global;

struct RawManagerCall {
    uint32 origin;
    uint256 proofMaturity;
    RawCallData callData;
}

using CastLib for RawManagerCall global;

struct RawBaseMessage {
    bytes32 sender;
    bytes32 recipient;
    RawTips tips;
    RawRequest request;
    bytes content;
}

using CastLib for RawBaseMessage global;

struct RawMessage {
    uint8 flag;
    RawHeader header;
    bytes body;
}

using CastLib for RawMessage global;

struct RawNumber {
    uint16 number;
}

using CastLib for RawNumber global;

struct RawGasData {
    RawNumber gasPrice;
    RawNumber dataPrice;
    RawNumber execBuffer;
    RawNumber amortAttCost;
    RawNumber etherPrice;
    RawNumber markup;
}

using CastLib for RawGasData global;

struct RawGasData256 {
    uint256 gasPrice;
    uint256 dataPrice;
    uint256 execBuffer;
    uint256 amortAttCost;
    uint256 etherPrice;
    uint256 markup;
}

using CastLib for RawGasData256 global;

struct RawChainGas {
    uint32 domain;
    RawGasData gasData;
}

using CastLib for RawChainGas global;

struct RawState {
    bytes32 root;
    uint32 origin;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
    RawGasData gasData;
}

using CastLib for RawState global;

struct RawStateIndex {
    uint256 stateIndex;
    uint256 statesAmount;
}

using CastLib for RawStateIndex global;

struct RawSnapshot {
    RawState[] states;
}

using CastLib for RawSnapshot global;

struct RawAttestation {
    bytes32 snapRoot;
    bytes32 agentRoot;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}

using CastLib for RawAttestation global;

struct RawAttestationReport {
    uint8 flag;
    RawAttestation attestation;
}

using CastLib for RawAttestationReport global;

struct RawStateReport {
    uint8 flag;
    RawState state;
}

using CastLib for RawStateReport global;

// solhint-disable no-empty-blocks
// solhint-disable ordering
library CastLib {
    using AttestationLib for bytes;
    using AttestationReportLib for bytes;
    using ByteString for bytes;
    using BaseMessageLib for bytes;
    using MessageLib for bytes;
    using ReceiptLib for bytes;
    using ReceiptReportLib for bytes;
    using SnapshotLib for bytes;
    using StateLib for bytes;
    using StateReportLib for bytes;

    /// @notice Prevents this contract from being included in the coverage report
    function testCastLib() external {}

    // ══════════════════════════════════════════════════ MESSAGE ══════════════════════════════════════════════════════

    function formatMessage(RawMessage memory rm) internal pure returns (bytes memory msgPayload) {
        // Explicit revert when out of range
        require(rm.flag <= uint8(type(MessageFlag).max), "Flag out of range");
        return MessageLib.formatMessage(MessageFlag(rm.flag), rm.header.castToHeader(), rm.body);
    }

    function castToMessage(RawMessage memory rm) internal pure returns (Message ptr) {
        ptr = rm.formatMessage().castToMessage();
    }

    function encodeHeader(RawHeader memory rh) internal pure returns (uint128 encodedHeader) {
        encodedHeader = Header.unwrap(rh.castToHeader());
    }

    function castToHeader(RawHeader memory rh) internal pure returns (Header header) {
        header = HeaderLib.encodeHeader({
            origin_: rh.origin,
            nonce_: rh.nonce,
            destination_: rh.destination,
            optimisticPeriod_: rh.optimisticPeriod
        });
    }

    function encodeRequest(RawRequest memory rr) internal pure returns (uint160 encodedReq) {
        encodedReq = Request.unwrap(rr.castToRequest());
    }

    function castToRequest(RawRequest memory rr) internal pure returns (Request request) {
        request = RequestLib.encodeRequest({gasDrop_: rr.gasDrop, gasLimit_: rr.gasLimit});
    }

    function encodeTips(RawTips memory rt) internal pure returns (uint256 encodedTips) {
        encodedTips = Tips.unwrap(rt.castToTips());
    }

    function boundTips(RawTips memory rt, uint64 maxTipValue) internal pure {
        rt.summitTip = rt.summitTip % maxTipValue;
        rt.attestationTip = rt.attestationTip % maxTipValue;
        rt.executionTip = rt.executionTip % maxTipValue;
        rt.deliveryTip = rt.deliveryTip % maxTipValue;
    }

    function floorTips(RawTips memory rt, uint64 minTipValue) internal pure {
        rt.summitTip = minTipValue + uint64(rt.summitTip % (2 ** 64 - minTipValue));
        rt.attestationTip = minTipValue + uint64(rt.attestationTip % (2 ** 64 - minTipValue));
        rt.executionTip = minTipValue + uint64(rt.executionTip % (2 ** 64 - minTipValue));
        rt.deliveryTip = minTipValue + uint64(rt.deliveryTip % (2 ** 64 - minTipValue));
    }

    function cloneTips(RawTips memory rt) internal pure returns (RawTips memory crt) {
        crt.summitTip = rt.summitTip;
        crt.attestationTip = rt.attestationTip;
        crt.executionTip = rt.executionTip;
        crt.deliveryTip = rt.deliveryTip;
    }

    function castToTips(RawTips memory rt) internal pure returns (Tips tips) {
        tips = TipsLib.encodeTips({
            summitTip_: rt.summitTip,
            attestationTip_: rt.attestationTip,
            executionTip_: rt.executionTip,
            deliveryTip_: rt.deliveryTip
        });
    }

    function formatBaseMessage(RawBaseMessage memory rbm) internal pure returns (bytes memory bmPayload) {
        bmPayload = BaseMessageLib.formatBaseMessage({
            sender_: rbm.sender,
            recipient_: rbm.recipient,
            tips_: rbm.tips.castToTips(),
            request_: rbm.request.castToRequest(),
            content_: rbm.content
        });
    }

    function castToBaseMessage(RawBaseMessage memory rbm) internal pure returns (BaseMessage ptr) {
        ptr = rbm.formatBaseMessage().castToBaseMessage();
    }

    // ══════════════════════════════════════════════ MANAGER MESSAGE ══════════════════════════════════════════════════

    function formatCallData(RawCallData memory rcd) internal pure returns (bytes memory cdPayload) {
        // Explicit revert when args are not taking whole amount of words
        require(rcd.args.length % 32 == 0, "Args don't take exactly N words");
        cdPayload = abi.encodePacked(rcd.selector, rcd.args);
    }

    function castToCallData(RawCallData memory rcd) internal pure returns (CallData ptr) {
        ptr = rcd.formatCallData().castToCallData();
    }

    function callPayload(RawManagerCall memory rsc) internal view returns (bytes memory scPayload) {
        // Add (msgOrigin, proofMaturity) as the first two args
        scPayload = rsc.callData.castToCallData().addPrefix(abi.encode(rsc.origin, rsc.proofMaturity));
    }

    // ═════════════════════════════════════════════════ RECEIPT ═════════════════════════════════════════════════════

    function formatReceiptBody(RawReceiptBody memory rrb) internal pure returns (bytes memory) {
        return ReceiptLib.formatReceiptBody(
            rrb.origin,
            rrb.destination,
            rrb.messageHash,
            rrb.snapshotRoot,
            rrb.stateIndex,
            rrb.attNotary,
            rrb.firstExecutor,
            rrb.finalExecutor
        );
    }

    function castToReceiptBody(RawReceiptBody memory rrb) internal pure returns (ReceiptBody) {
        return rrb.formatReceiptBody().castToReceiptBody();
    }

    function formatReceipt(RawExecReceipt memory re) internal pure returns (bytes memory) {
        return ReceiptLib.formatReceipt(re.body.formatReceiptBody(), re.tips.castToTips());
    }

    function castToReceipt(RawExecReceipt memory re) internal pure returns (Receipt) {
        return re.formatReceipt().castToReceipt();
    }

    function modifyReceiptBody(RawReceiptBody memory rrb, uint256 mask)
        internal
        pure
        returns (RawReceiptBody memory mrb)
    {
        // Don't modify the destination, message hash
        mrb.destination = rrb.destination;
        mrb.messageHash = rrb.messageHash;
        // Make sure at least one value is modified, valid mask values are [1 .. 63]
        mask = 1 + (mask % 63);
        mrb.origin = rrb.origin ^ uint32(mask & 1);
        mrb.snapshotRoot = rrb.snapshotRoot ^ bytes32(mask & 2);
        mrb.stateIndex = rrb.stateIndex ^ uint8(mask & 4);
        mrb.attNotary = address(uint160(rrb.attNotary) ^ uint160(mask & 8));
        mrb.firstExecutor = address(uint160(rrb.firstExecutor) ^ uint160(mask & 16));
        mrb.finalExecutor = address(uint160(rrb.finalExecutor) ^ uint160(mask & 32));
    }

    function formatReceiptReport(RawReceiptReport memory rawRR) internal pure returns (bytes memory) {
        // Explicit revert when flag out of range
        require(rawRR.flag <= uint8(type(ReceiptFlag).max), "Flag out of range");
        return ReceiptFlag(rawRR.flag).formatReceiptReport(rawRR.body.formatReceiptBody());
    }

    function castToReceiptReport(RawReceiptReport memory rawRR) internal pure returns (ReceiptReport) {
        return rawRR.formatReceiptReport().castToReceiptReport();
    }

    // ═════════════════════════════════════════════════ GAS DATA ══════════════════════════════════════════════════════

    function encodeNumber(RawNumber memory rn) internal pure returns (Number) {
        return Number.wrap(rn.number);
    }

    function round(uint256 num) internal pure returns (uint256) {
        return NumberLib.decompress(NumberLib.compress(num));
    }

    function round(RawGasData256 memory rgd256) internal pure {
        rgd256.gasPrice = round(rgd256.gasPrice);
        rgd256.dataPrice = round(rgd256.dataPrice);
        rgd256.execBuffer = round(rgd256.execBuffer);
        rgd256.amortAttCost = round(rgd256.amortAttCost);
        rgd256.etherPrice = round(rgd256.etherPrice);
        rgd256.markup = round(rgd256.markup);
    }

    function compress(RawGasData256 memory rdg256) internal pure returns (RawGasData memory rgd) {
        rgd.gasPrice.number = Number.unwrap(NumberLib.compress(rdg256.gasPrice));
        rgd.dataPrice.number = Number.unwrap(NumberLib.compress(rdg256.dataPrice));
        rgd.execBuffer.number = Number.unwrap(NumberLib.compress(rdg256.execBuffer));
        rgd.amortAttCost.number = Number.unwrap(NumberLib.compress(rdg256.amortAttCost));
        rgd.etherPrice.number = Number.unwrap(NumberLib.compress(rdg256.etherPrice));
        rgd.markup.number = Number.unwrap(NumberLib.compress(rdg256.markup));
    }

    function encodeGasData(RawGasData memory rgd) internal pure returns (uint96 encodedGasData) {
        return GasData.unwrap(rgd.castToGasData());
    }

    function castToGasData(RawGasData memory rgd) internal pure returns (GasData) {
        return GasDataLib.encodeGasData({
            gasPrice_: rgd.gasPrice.encodeNumber(),
            dataPrice_: rgd.dataPrice.encodeNumber(),
            execBuffer_: rgd.execBuffer.encodeNumber(),
            amortAttCost_: rgd.amortAttCost.encodeNumber(),
            etherPrice_: rgd.etherPrice.encodeNumber(),
            markup_: rgd.markup.encodeNumber()
        });
    }

    function encodeChainGas(RawChainGas memory rcg) internal pure returns (uint128 encodedChainGas) {
        return ChainGas.unwrap(rcg.castToChainGas());
    }

    function castToChainGas(RawChainGas memory rcg) internal pure returns (ChainGas) {
        return GasDataLib.encodeChainGas({domain_: rcg.domain, gasData_: rcg.gasData.castToGasData()});
    }

    // ═══════════════════════════════════════════════════ STATE ═══════════════════════════════════════════════════════

    function formatState(RawState memory rs) internal pure returns (bytes memory state) {
        state = StateLib.formatState({
            root_: rs.root,
            origin_: rs.origin,
            nonce_: rs.nonce,
            blockNumber_: rs.blockNumber,
            timestamp_: rs.timestamp,
            gasData_: rs.gasData.castToGasData()
        });
    }

    function castToState(RawState memory rs) internal pure returns (State ptr) {
        ptr = rs.formatState().castToState();
    }

    function modifyState(RawState memory rs, uint256 mask) internal pure returns (RawState memory mrs) {
        // Make sure at least one value is modified, valid mask values are [1 .. 2047]
        mask = 1 + (mask % 2047);
        mrs.root = rs.root ^ bytes32(mask & 1);
        mrs.origin = rs.origin ^ uint32(mask & 2);
        mrs.nonce = rs.nonce ^ uint32(mask & 4);
        mrs.blockNumber = rs.blockNumber ^ uint32(mask & 8);
        mrs.timestamp = rs.timestamp ^ uint32(mask & 16);
        mrs.gasData.gasPrice.number = rs.gasData.gasPrice.number ^ uint16(mask & 32);
        mrs.gasData.dataPrice.number = rs.gasData.dataPrice.number ^ uint16(mask & 64);
        mrs.gasData.execBuffer.number = rs.gasData.execBuffer.number ^ uint16(mask & 128);
        mrs.gasData.amortAttCost.number = rs.gasData.amortAttCost.number ^ uint16(mask & 256);
        mrs.gasData.etherPrice.number = rs.gasData.etherPrice.number ^ uint16(mask & 512);
        mrs.gasData.markup.number = rs.gasData.markup.number ^ uint16(mask & 1024);
    }

    function formatStateReport(RawStateReport memory rawSR) internal pure returns (bytes memory stateReport) {
        // Explicit revert when flag out of range
        require(rawSR.flag <= uint8(type(StateFlag).max), "Flag out of range");
        bytes memory state = rawSR.state.formatState();
        stateReport = StateFlag(rawSR.flag).formatStateReport(state);
    }

    function castToStateReport(RawStateReport memory rawSR) internal pure returns (StateReport ptr) {
        ptr = rawSR.formatStateReport().castToStateReport();
    }

    function boundStateIndex(RawStateIndex memory rsi) internal pure {
        // [1 .. SNAPSHOT_MAX_STATES] range
        rsi.statesAmount = 1 + rsi.statesAmount % SNAPSHOT_MAX_STATES;
        // [0 .. statesAmount) range
        rsi.stateIndex = rsi.stateIndex % rsi.statesAmount;
    }

    // ═════════════════════════════════════════════════ SNAPSHOT ══════════════════════════════════════════════════════

    function formatStates(RawSnapshot memory rawSnap) internal pure returns (bytes[] memory states) {
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
        ra.snapRoot = snapshot.calculateRoot();
        ra.agentRoot = agentRoot;
        ra.nonce = nonce;
        ra.blockNumber = blockNumber;
        ra.timestamp = timestamp;
    }

    function formatSnapshot(RawSnapshot memory rawSnap) internal view returns (bytes memory snapshot) {
        State[] memory states = new State[](rawSnap.states.length);
        for (uint256 i = 0; i < rawSnap.states.length; ++i) {
            states[i] = rawSnap.states[i].castToState();
        }
        snapshot = SnapshotLib.formatSnapshot(states);
    }

    function castToSnapshot(RawSnapshot memory rawSnap) internal view returns (Snapshot ptr) {
        ptr = rawSnap.formatSnapshot().castToSnapshot();
    }

    // ════════════════════════════════════════════════ ATTESTATION ════════════════════════════════════════════════════

    function formatAttestation(RawAttestation memory ra) internal pure returns (bytes memory attestation) {
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

    function castToAttestationReport(RawAttestationReport memory rawAR) internal pure returns (AttestationReport ptr) {
        ptr = rawAR.formatAttestationReport().castToAttestationReport();
    }
}
