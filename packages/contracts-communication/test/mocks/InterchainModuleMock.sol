// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {IInterchainDB} from "../../contracts/interfaces/IInterchainDB.sol";

// solhint-disable ordering
// solhint-disable no-empty-blocks
contract InterchainModuleMock is IInterchainModule {
    function requestBatchVerification(uint256 dstChainId, bytes calldata versionedBatch) external payable {}

    function getModuleFee(uint256 dstChainId, uint256 dbNonce) external view returns (uint256) {}

    function mockVerifyRemoteBatch(address interchainDB, bytes calldata versionedBatch) external {
        IInterchainDB(interchainDB).verifyRemoteBatch(versionedBatch);
    }
}
