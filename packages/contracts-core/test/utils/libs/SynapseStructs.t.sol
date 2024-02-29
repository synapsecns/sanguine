// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString, CallData, MemView, MemViewLib} from "../../../contracts/libs/memory/ByteString.sol";

import {BaseMessage, BaseMessageLib, Tips, TipsLib} from "../../../contracts/libs/memory/BaseMessage.sol";
import {ChainGas, GasData, GasDataLib} from "../../../contracts/libs/stack/GasData.sol";
import {Header, HeaderLib, MessageFlag} from "../../../contracts/libs/stack/Header.sol";
import {Message, MessageLib} from "../../../contracts/libs/memory/Message.sol";
import {Number, NumberLib} from "../../../contracts/libs/stack/Number.sol";
import {Receipt, ReceiptLib} from "../../../contracts/libs/memory/Receipt.sol";
import {Request, RequestLib} from "../../../contracts/libs/stack/Request.sol";

import {
    Snapshot, SnapshotLib, SNAPSHOT_MAX_STATES, State, StateLib
} from "../../../contracts/libs/memory/Snapshot.sol";

import {Attestation, AttestationLib} from "../../../contracts/libs/memory/Attestation.sol";

struct RawHeader {
    uint8 flag;
    uint32 origin;
    uint32 nonce;
    uint32 destination;
    uint32 optimisticPeriod;
}

using CastLib for RawHeader global;

struct RawRequest {
    uint96 gasDrop;
    uint64 gasLimit;
    uint32 version;
}

using CastLib for RawRequest global;

struct RawTips {
    uint64 summitTip;
    uint64 attestationTip;
    uint64 executionTip;
    uint64 deliveryTip;
}

using CastLib for RawTips global;

struct RawTipsProof {
    bytes32 headerHash;
    bytes32 bodyHash;
}

using CastLib for RawTips global;

// RawReceipt name is already taken in forge-std
struct RawExecReceipt {
    uint32 origin;
    uint32 destination;
    bytes32 messageHash;
    bytes32 snapshotRoot;
    uint8 stateIndex;
    address attNotary;
    address firstExecutor;
    address finalExecutor;
}

struct RawReceiptTips {
    RawExecReceipt re;
    RawTips tips;
    RawTipsProof rtp;
}

using CastLib for RawExecReceipt global;

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
    RawTips tips;
    bytes32 sender;
    bytes32 recipient;
    RawRequest request;
    bytes content;
}

using CastLib for RawBaseMessage global;

struct RawMessage {
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
    uint8 stateIndex;
    uint256 statesAmount;
}

using CastLib for RawStateIndex global;

struct RawSnapshot {
    RawState[] states;
}

using CastLib for RawSnapshot global;

struct RawAttestation {
    bytes32 snapRoot;
    bytes32 dataHash;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
    // Merged into dataHash
    bytes32 _agentRoot;
    bytes32 _snapGasHash;
}

using CastLib for RawAttestation global;

// solhint-disable no-empty-blocks
// solhint-disable ordering
library CastLib {
    using AttestationLib for bytes;
    using ByteString for bytes;
    using BaseMessageLib for bytes;
    using MessageLib for bytes;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;
    using StateLib for bytes;

    /// @notice Prevents this contract from being included in the coverage report
    function testCastLib() external {}

    // ══════════════════════════════════════════════════ MESSAGE ══════════════════════════════════════════════════════

    function formatMessage(RawMessage memory rm) internal pure returns (bytes memory msgPayload) {
        return MessageLib.formatMessage(rm.header.castToHeader(), rm.body);
    }

    function castToMessage(RawMessage memory rm) internal pure returns (Message ptr) {
        ptr = rm.formatMessage().castToMessage();
    }

    function encodeHeader(RawHeader memory rh) internal pure returns (uint136 encodedHeader) {
        encodedHeader = Header.unwrap(rh.castToHeader());
    }

    function castToHeader(RawHeader memory rh) internal pure returns (Header header) {
        // Explicit revert when out of range
        require(rh.flag <= uint8(type(MessageFlag).max), "Flag out of range");
        header = HeaderLib.encodeHeader({
            flag_: MessageFlag(rh.flag),
            origin_: rh.origin,
            nonce_: rh.nonce,
            destination_: rh.destination,
            optimisticPeriod_: rh.optimisticPeriod
        });
    }

    function boundFlag(RawHeader memory rh) internal pure {
        rh.flag = rh.flag % (uint8(type(MessageFlag).max) + 1);
    }

    function encodeRequest(RawRequest memory rr) internal pure returns (uint192 encodedReq) {
        encodedReq = Request.unwrap(rr.castToRequest());
    }

    function castToRequest(RawRequest memory rr) internal pure returns (Request request) {
        request = RequestLib.encodeRequest({gasDrop_: rr.gasDrop, gasLimit_: rr.gasLimit, version_: rr.version});
    }

    function boundRequest(RawRequest memory rr, uint96 maxGasDrop, uint64 maxGasLimit) internal pure {
        require(maxGasDrop != 0, "maxGasDrop can't be 0");
        require(maxGasLimit != 0, "maxGasLimit can't be 0");
        rr.gasDrop = rr.gasDrop % maxGasDrop;
        rr.gasLimit = rr.gasLimit % maxGasLimit;
    }

    function encodeTips(RawTips memory rt) internal pure returns (uint256 encodedTips) {
        encodedTips = Tips.unwrap(rt.castToTips());
    }

    function boundTips(RawTips memory rt, uint64 maxTipValue) internal pure {
        require(maxTipValue != 0, "maxTipValue can't be 0");
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

    function getMessageHash(RawTips memory rt, RawTipsProof memory rtp) internal pure returns (bytes32) {
        bytes32 baseMessageLeaf = keccak256(bytes.concat(rt.castToTips().leaf(), rtp.bodyHash));
        return keccak256(bytes.concat(rtp.headerHash, baseMessageLeaf));
    }

    function formatBaseMessage(RawBaseMessage memory rbm) internal pure returns (bytes memory bmPayload) {
        bmPayload = BaseMessageLib.formatBaseMessage({
            tips_: rbm.tips.castToTips(),
            sender_: rbm.sender,
            recipient_: rbm.recipient,
            request_: rbm.request.castToRequest(),
            content_: rbm.content
        });
    }

    function castToBaseMessage(RawBaseMessage memory rbm) internal pure returns (BaseMessage ptr) {
        ptr = rbm.formatBaseMessage().castToBaseMessage();
    }

    function getTipsProof(RawHeader memory rh, RawBaseMessage memory rbm)
        internal
        pure
        returns (RawTipsProof memory rtp)
    {
        rtp.headerHash = rh.castToHeader().leaf();
        rtp.bodyHash = rbm.castToBaseMessage().bodyLeaf();
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

    function formatReceipt(RawExecReceipt memory re) internal pure returns (bytes memory) {
        return ReceiptLib.formatReceipt(
            re.origin,
            re.destination,
            re.messageHash,
            re.snapshotRoot,
            re.stateIndex,
            re.attNotary,
            re.firstExecutor,
            re.finalExecutor
        );
    }

    function castToReceipt(RawExecReceipt memory re) internal pure returns (Receipt) {
        return re.formatReceipt().castToReceipt();
    }

    function modifyReceipt(RawExecReceipt memory re, uint256 mask) internal pure returns (RawExecReceipt memory mrb) {
        // Don't modify the destination, message hash
        mrb.destination = re.destination;
        mrb.messageHash = re.messageHash;
        // Make sure at least one value is modified, valid mask values are [1 .. 63]
        mask = 1 + (mask % 63);
        mrb.origin = re.origin ^ uint32(mask & 1);
        mrb.snapshotRoot = re.snapshotRoot ^ bytes32(mask & 2);
        mrb.stateIndex = re.stateIndex ^ uint8(mask & 4);
        mrb.attNotary = address(uint160(re.attNotary) ^ uint160(mask & 8));
        mrb.firstExecutor = address(uint160(re.firstExecutor) ^ uint160(mask & 16));
        mrb.finalExecutor = address(uint160(re.finalExecutor) ^ uint160(mask & 32));
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

    function decompress(RawGasData memory rgd) internal pure returns (RawGasData256 memory rdg256) {
        rdg256.gasPrice = Number.wrap(rgd.gasPrice.number).decompress();
        rdg256.dataPrice = Number.wrap(rgd.dataPrice.number).decompress();
        rdg256.execBuffer = Number.wrap(rgd.execBuffer.number).decompress();
        rdg256.amortAttCost = Number.wrap(rgd.amortAttCost.number).decompress();
        rdg256.etherPrice = Number.wrap(rgd.etherPrice.number).decompress();
        rdg256.markup = Number.wrap(rgd.markup.number).decompress();
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

    function boundStateIndex(RawStateIndex memory rsi) internal pure {
        // [1 .. SNAPSHOT_MAX_STATES] range
        rsi.statesAmount = 1 + rsi.statesAmount % SNAPSHOT_MAX_STATES;
        // [0 .. statesAmount) range
        rsi.stateIndex = uint8(rsi.stateIndex % rsi.statesAmount);
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
        ra._agentRoot = agentRoot;
        ra._snapGasHash = rawSnap.snapGasHash();
        ra.setDataHash();
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

    function snapGas(RawSnapshot memory rawSnap) internal view returns (uint256[] memory snapGas_) {
        ChainGas[] memory chainData = rawSnap.castToSnapshot().snapGas();
        snapGas_ = new uint256[](chainData.length);
        for (uint256 i = 0; i < snapGas_.length; ++i) {
            snapGas_[i] = ChainGas.unwrap(chainData[i]);
        }
    }

    function snapGasHash(RawSnapshot memory rawSnap) internal view returns (bytes32) {
        return GasDataLib.snapGasHash(rawSnap.castToSnapshot().snapGas());
    }

    // ════════════════════════════════════════════════ ATTESTATION ════════════════════════════════════════════════════

    function formatAttestation(RawAttestation memory ra) internal pure returns (bytes memory attestation) {
        attestation = AttestationLib.formatAttestation({
            snapRoot_: ra.snapRoot,
            dataHash_: ra.dataHash,
            nonce_: ra.nonce,
            blockNumber_: ra.blockNumber,
            timestamp_: ra.timestamp
        });
    }

    function castToAttestation(RawAttestation memory ra) internal pure returns (Attestation ptr) {
        ptr = ra.formatAttestation().castToAttestation();
    }

    function setDataHash(RawAttestation memory ra) internal pure {
        ra.dataHash = keccak256(bytes.concat(ra._agentRoot, ra._snapGasHash));
    }

    function modifyAttestation(RawAttestation memory ra, uint256 mask)
        internal
        pure
        returns (bool isEqual, RawAttestation memory mra)
    {
        // Don't modify the nonce
        mra.nonce = ra.nonce;
        // Check if at least one value was modified by checking last 5 bits
        isEqual = mask & 31 == 0;
        mra.snapRoot = ra.snapRoot ^ bytes32(mask & 1);
        mra._agentRoot = ra._agentRoot ^ bytes32(mask & 2);
        mra._snapGasHash = ra._snapGasHash ^ bytes32(mask & 4);
        mra.blockNumber = ra.blockNumber ^ uint40(mask & 8);
        mra.timestamp = ra.timestamp ^ uint40(mask & 16);
        mra.setDataHash();
    }
}
