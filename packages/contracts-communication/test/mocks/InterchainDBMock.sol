// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, InterchainEntry} from "../../contracts/interfaces/IInterchainDB.sol";

// solhint-disable no-empty-blocks
contract InterchainDBMock is IInterchainDB {
    uint16 public constant DB_VERSION = 1;

    function writeEntry(bytes32 dataHash) external returns (uint64 writerNonce) {}

    function requestEntryVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] memory srcModules
    )
        external
        payable
    {}

    function writeEntryRequestVerification(
        uint64 dstChainId,
        bytes32 digest,
        address[] memory srcModules
    )
        external
        payable
        returns (uint64 dbNonce)
    {}

    function verifyRemoteEntry(bytes calldata encodedEntry) external {}

    function getInterchainFee(uint64 dstChainId, address[] memory srcModules) external view returns (uint256) {}

    function getEncodedEntry(uint64 dbNonce) external view returns (bytes memory) {}

    function getEntry(uint64 dbNonce) external view returns (InterchainEntry memory) {}

    function getEntryValue(uint64 dbNonce) external view returns (bytes32) {}

    function getDBNonce() external view returns (uint64) {}

    function checkEntryVerification(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        returns (uint256 moduleVerifiedAt)
    {}
}
