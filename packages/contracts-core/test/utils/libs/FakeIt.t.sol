// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../../contracts/libs/State.sol";

// solhint-disable func-visibility
// Collection of free functions to generate test data

function fakeBytes(uint256 length, uint256 nonce) pure returns (bytes memory fake) {
    fake = new bytes(length);
    for (uint256 i = 0; i < length; ++i) {
        fake[i] = bytes1(uint8(nonce));
    }
}

/// @notice Returns a list of states with given state on given position
/// + fake data for everything else.
function fakeStates(
    SummitState memory state,
    uint256 statesAmount,
    uint256 stateIndex
) pure returns (bytes[] memory states, State[] memory ptrs) {
    states = new bytes[](statesAmount);
    ptrs = new State[](statesAmount);
    for (uint256 i = 0; i < statesAmount; ++i) {
        if (i == stateIndex) {
            states[i] = state.formatSummitState();
        } else {
            // Create different garbage values
            states[i] = fakeBytes(STATE_LENGTH, i + 1);
        }
        ptrs[i] = StateLib.castToState(states[i]);
    }
}
