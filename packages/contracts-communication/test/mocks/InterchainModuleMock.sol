// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";
import {IInterchainDB} from "../../contracts/interfaces/IInterchainDB.sol";
import {InterchainEntry, InterchainEntryLib} from "../../contracts/libs/InterchainEntry.sol";
import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

// solhint-disable ordering
// solhint-disable no-empty-blocks
contract InterchainModuleMock is IInterchainModule {
    function requestEntryVerification(uint64 dstChainId, bytes calldata versionedEntry) external payable {}

    function getModuleFee(uint64 dstChainId) external view returns (uint256) {}

    function mockVerifyRemoteEntry(address interchainDB, bytes calldata versionedEntry) external {
        IInterchainDB(interchainDB).verifyRemoteEntry(versionedEntry);
    }

    /// @notice This function is exposed for simplicity of Go tests. It uses the version signalled
    /// by the InterchainDB contract to encode the versioned entry.
    function mockVerifyRemoteEntryStruct(address interchainDB, InterchainEntry memory entry) external {
        bytes memory versionedEntry = VersionedPayloadLib.encodeVersionedPayload({
            version: IInterchainDB(interchainDB).DB_VERSION(),
            payload: InterchainEntryLib.encodeEntry(entry)
        });
        IInterchainDB(interchainDB).verifyRemoteEntry(versionedEntry);
    }
}
