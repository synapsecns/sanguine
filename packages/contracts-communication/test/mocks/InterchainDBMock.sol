// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, InterchainEntry} from "../../contracts/interfaces/IInterchainDB.sol";

contract InterchainDBMock is IInterchainDB {
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce) {}

    function requestVerification(
        uint256 destChainId,
        address writer,
        uint256 writerNonce,
        address[] memory srcModules
    )
        external
        payable
    {}

    function writeEntryWithVerification(
        uint256 destChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint256 writerNonce)
    {}

    function verifyEntry(InterchainEntry memory entry) external {}

    function getInterchainFee(uint256 destChainId, address[] memory srcModules) external view returns (uint256) {}

    function getEntry(address writer, uint256 writerNonce) external view returns (InterchainEntry memory) {}

    function getWriterNonce(address writer) external view returns (uint256) {}

    function readEntry(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        returns (uint256 moduleVerifiedAt)
    {}
}
