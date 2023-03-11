// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract EnumerableSetTools {
    uint256 internal constant ELEMENTS = 4;
    uint256[ELEMENTS] internal removalOrder = [2, 0, 1, 3];
    uint256[][] internal expectedStates;

    /// @notice Prevents this contract from being included in the coverage report
    function testEnumerableSetTools() external {}

    function createExpectedStates() public {
        // Test example for checking the removal function
        // On every step element removalOrder[i] is removed
        // Expected state for Enumerable set is saved
        expectedStates = new uint256[][](ELEMENTS);
        expectedStates[0] = [0, 1, 2, 3]; // (2) is removed
        expectedStates[1] = [0, 1, 3]; // (0) is removed
        expectedStates[2] = [3, 1]; // (1) is removed
        expectedStates[3] = [3]; // (3) is removed
    }
}
