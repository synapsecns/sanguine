// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TipsLib } from "../../../contracts/libs/Tips.sol";

abstract contract TipsTools {
    // Mock values
    uint96 internal constant MOCK_NOTARY_TIP = 1e2;
    uint96 internal constant MOCK_BROADCASTER_TIP = 1e3;
    uint96 internal constant MOCK_PROVER_TIP = 1e4;
    uint96 internal constant MOCK_EXECUTOR_TIP = 1e5;

    uint96 internal tipNotary;
    uint96 internal tipBroadcaster;
    uint96 internal tipProver;
    uint96 internal tipExecutor;

    bytes internal tipsRaw;
    uint256 internal tipsTotal;

    /// @notice Prevents this contract from being included in the coverage report
    function testTipsTools() external {}

    // Create the tip payload and calculates the total tip using the saved data
    function createTips() public {
        tipsRaw = TipsLib.formatTips(tipNotary, tipBroadcaster, tipProver, tipExecutor);
        calcTotalTips();
    }

    // Create mock tips with given nonce. Useful for tests with different tip values
    function createMockTips(uint32 nonce) public {
        tipNotary = MOCK_NOTARY_TIP + nonce;
        tipBroadcaster = MOCK_BROADCASTER_TIP + nonce;
        tipProver = MOCK_PROVER_TIP + nonce;
        tipExecutor = MOCK_EXECUTOR_TIP + nonce;
        createTips();
    }

    // Create empty tips
    function createEmptyTips() public {
        tipNotary = tipBroadcaster = tipProver = tipExecutor = 0;
        createTips();
    }

    // Populate total tips value
    function calcTotalTips() public {
        tipsTotal = tipNotary + tipBroadcaster + tipProver + tipExecutor;
    }
}
