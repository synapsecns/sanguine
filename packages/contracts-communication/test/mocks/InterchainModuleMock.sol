// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";
import {InterchainBatch, IInterchainDB, IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";

// solhint-disable ordering
contract InterchainModuleMock is IInterchainModule {
    function requestBatchVerification(uint256 dstChainId, InterchainBatch memory batch) external payable {}

    function getModuleFee(uint256 dstChainId, uint256 dbNonce) external view returns (uint256) {}

    function mockVerifyBatch(address interchainDB, InterchainBatch memory batch) external {
        IInterchainDB(interchainDB).verifyRemoteBatch(batch);
    }

    function mockVerifyEntry(address interchainDB, InterchainEntry memory entry) external {
        // TODO: deprecated
    }
}
