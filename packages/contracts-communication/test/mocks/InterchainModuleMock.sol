// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {IInterchainDB} from "../../contracts/interfaces/IInterchainDB.sol";
import {InterchainBatch, InterchainBatchLib} from "../../contracts/libs/InterchainBatch.sol";
import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

// solhint-disable ordering
// solhint-disable no-empty-blocks
contract InterchainModuleMock is IInterchainModule {
    function requestBatchVerification(uint256 dstChainId, bytes calldata versionedBatch) external payable {}

    function getModuleFee(uint256 dstChainId, uint256 dbNonce) external view returns (uint256) {}

    function mockVerifyRemoteBatch(address interchainDB, bytes calldata versionedBatch) external {
        IInterchainDB(interchainDB).verifyRemoteBatch(versionedBatch);
    }

    // @notice This function is exposed for simplicity of Go tests. It uses the version signalled
    // by the InterchainDB contract to encode the versioned batch.
    function mockVerifyRemoteBatchStruct(address interchainDB, InterchainBatch memory batch) external {
        bytes memory versionedBatch = VersionedPayloadLib.encodeVersionedPayload({
            version: IInterchainDB(interchainDB).DB_VERSION(),
            payload: InterchainBatchLib.encodeBatch(batch)
        });
        IInterchainDB(interchainDB).verifyRemoteBatch(versionedBatch);
    }
}
