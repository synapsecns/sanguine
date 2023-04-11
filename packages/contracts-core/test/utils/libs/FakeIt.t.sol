// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {State, StateLib, STATE_LENGTH, SummitState} from "../../../contracts/libs/State.sol";

import {RawSnapshot, RawState, RawStateIndex} from "./SynapseStructs.t.sol";

// solhint-disable func-visibility
// Collection of free functions to generate test data

/// @notice Returns RawState struct filled with fake data.
function fakeState(uint256 fakeValue) pure returns (RawState memory state) {
    state.root = bytes32(fakeValue);
    state.origin = uint32(fakeValue);
    state.nonce = uint32(fakeValue);
    state.blockNumber = uint40(fakeValue);
    state.timestamp = uint40(fakeValue);
}

/// @notice Returns RawSnapshot struct with given state on given position,
/// and with fake states for everything else.
function fakeSnapshot(RawState memory state, RawStateIndex memory rsi) pure returns (RawSnapshot memory rawSnap) {
    rawSnap.states = new RawState[](rsi.statesAmount);
    for (uint256 i = 0; i < rsi.statesAmount; ++i) {
        // Create different non-zero garbage values
        rawSnap.states[i] = i == rsi.stateIndex ? state : fakeState(i + 1);
    }
}

/// @notice Returns RawSnapshot struct with fake states.
function fakeSnapshot(uint256 statesAmount) pure returns (RawSnapshot memory rawSnap) {
    RawState memory state;
    return fakeSnapshot(state, RawStateIndex(statesAmount, statesAmount));
}
