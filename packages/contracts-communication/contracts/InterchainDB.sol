// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainDBEvents} from "./interfaces/IInterchainDBEvents.sol";

contract InterchainDB is IInterchainDB, IInterchainDBEvents {
    // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce) {}

    /// @inheritdoc IInterchainDB
    function requestVerification(
        uint256 destChainId,
        address writer,
        uint256 writerNonce,
        address[] memory srcModules
    )
        external
        payable
    {}

    /// @inheritdoc IInterchainDB
    function writeEntryWithVerification(
        uint256 destChainId,
        bytes32 dataHash,
        address[] calldata srcModules
    )
        external
        payable
        returns (uint256 writerNonce)
    {}

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function verifyEntry(InterchainEntry memory entry) external {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function getInterchainFee(uint256 destChainId, address[] calldata srcModules) external view returns (uint256) {}

    /// @inheritdoc IInterchainDB
    function getEntry(address writer, uint256 writerNonce) external view returns (InterchainEntry memory) {}

    /// @inheritdoc IInterchainDB
    function getWriterNonce(address writer) external view returns (uint256) {}

    /// @inheritdoc IInterchainDB
    function readEntry(
        InterchainEntry memory entry,
        address[] memory dstModules
    )
        external
        view
        returns (uint256[] memory moduleVerifiedAt)
    {}
}
