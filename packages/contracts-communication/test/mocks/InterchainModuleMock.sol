// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, IInterchainDB, IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";

// solhint-disable ordering
// solhint-disable no-empty-blocks
contract InterchainModuleMock is IInterchainModule {
    function requestBatchVerification(uint256 dstChainId, InterchainBatch memory batch) external payable {}

    function getModuleFee(uint256 dstChainId, uint256 dbNonce) external view returns (uint256) {}

    function mockVerifyRemoteBatch(address interchainDB, InterchainBatch memory batch) external {
        IInterchainDB(interchainDB).verifyRemoteBatch(batch);
    }
}
