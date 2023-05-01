// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RawAttestation, RawGasData, RawState, RawStateIndex, RawSnapshot} from "./SynapseStructs.t.sol";

struct Random {
    bytes32 seed;
}

using RandomLib for Random global;

// solhint-disable no-empty-blocks
// solhint-disable ordering
library RandomLib {
    /// @notice Prevents this contract from being included in the coverage report
    function testRandomLib() external {}

    // @notice Returns next "random" bytes32 value and updates the Random's seed.
    function next(Random memory r) internal pure returns (bytes32 value) {
        value = r.seed;
        r.seed = keccak256(bytes.concat(value));
    }

    // @notice Returns next "random" bytes value of given length and updates the Random's seed.
    function nextBytes(Random memory r, uint256 length) internal pure returns (bytes memory value) {
        value = new bytes(length);
        uint256 words = length / 32;
        for (uint256 i = 0; i < words; ++i) {
            bytes32 word = r.next();
            // TODO: This is probably not the best way to do this - rewrite in assembly
            for (uint256 j = 0; j < 32; ++j) {
                value[32 * i + j] = word[j];
            }
        }
        uint256 remainder = length % 32;
        if (remainder != 0) {
            bytes32 word = r.next();
            for (uint256 j = 0; j < remainder; ++j) {
                value[32 * words + j] = word[j];
            }
        }
    }

    // @notice Returns next "random" bytes value having N memory words and updates the Random's seed.
    function nextBytesWords(Random memory r, uint256 words) internal pure returns (bytes memory value) {
        bytes32[] memory args = new bytes32[](words);
        for (uint256 i = 0; i < words; ++i) {
            args[i] = r.next();
        }
        return abi.encodePacked(args);
    }

    // @notice Returns next "random" uint256 value and updates the Random's seed.
    function nextUint256(Random memory r) internal pure returns (uint256 value) {
        return uint256(r.next());
    }

    // @notice Returns next "random" uint160 value and updates the Random's seed.
    function nextUint160(Random memory r) internal pure returns (uint160 value) {
        return uint160(r.nextUint256());
    }

    // @notice Returns next "random" uint64 value and updates the Random's seed.
    function nextUint64(Random memory r) internal pure returns (uint64 value) {
        return uint64(r.nextUint256());
    }

    // @notice Returns next "random" uint40 value and updates the Random's seed.
    function nextUint40(Random memory r) internal pure returns (uint40 value) {
        return uint40(r.nextUint256());
    }

    // @notice Returns next "random" uint32 value and updates the Random's seed.
    function nextUint32(Random memory r) internal pure returns (uint32 value) {
        return uint32(r.nextUint256());
    }

    // @notice Returns next "random" uint16 value and updates the Random's seed.
    function nextUint16(Random memory r) internal pure returns (uint16 value) {
        return uint16(r.nextUint256());
    }

    // @notice Returns next "random" uint8 value and updates the Random's seed.
    function nextUint8(Random memory r) internal pure returns (uint8 value) {
        return uint8(r.nextUint256());
    }

    // @notice Returns next "random" address value and updates the Random's seed.
    function nextAddress(Random memory r) internal pure returns (address value) {
        return address(r.nextUint160());
    }

    function nextState(Random memory r, uint32 origin, uint32 nonce) internal pure returns (RawState memory state) {
        state.root = r.next();
        state.origin = origin;
        state.nonce = nonce;
        state.blockNumber = r.nextUint40();
        state.timestamp = r.nextUint40();
        state.gasData = r.nextGasData();
    }

    function nextState(Random memory r) internal pure returns (RawState memory state) {
        return r.nextState(r.nextUint32(), r.nextUint32());
    }

    function nextGasData(Random memory r) internal pure returns (RawGasData memory rgd) {
        rgd.gasPrice.number = r.nextUint16();
        rgd.dataPrice.number = r.nextUint16();
        rgd.execBuffer.number = r.nextUint16();
        rgd.amortAttCost.number = r.nextUint16();
        rgd.etherPrice.number = r.nextUint16();
        rgd.markup.number = r.nextUint16();
    }

    function nextStateIndex(Random memory r) internal pure returns (RawStateIndex memory rsi) {
        rsi.stateIndex = r.nextUint256();
        rsi.statesAmount = r.nextUint256();
        rsi.boundStateIndex();
    }

    function nextAttestation(Random memory r, uint32 nonce) internal pure returns (RawAttestation memory ra) {
        ra.snapRoot = r.next();
        ra.agentRoot = r.next();
        ra.nonce = nonce;
        ra.blockNumber = r.nextUint40();
        ra.timestamp = r.nextUint40();
    }

    function nextAttestation(Random memory r, RawSnapshot memory rawSnap, uint32 nonce)
        internal
        view
        returns (RawAttestation memory ra)
    {
        return rawSnap.castToRawAttestation(r.next(), nonce, r.nextUint40(), r.nextUint40());
    }
}
