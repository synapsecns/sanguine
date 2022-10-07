// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract DestinationHarnessEvents {
    event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip);
}
