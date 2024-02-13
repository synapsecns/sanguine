// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainDBEvents} from "./interfaces/IInterchainDBEvents.sol";

contract InterchainDB is IInterchainDB, IInterchainDBEvents {
    // ════════════════════════════════════════════════ APP-FACING ═════════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function writeEntry(bytes32 dataHash, uint256 destChainId, address[] calldata srcModules) external payable {}

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function confirmEntry(InterchainEntry memory entry) external {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function getInterchainFee(uint256 destChainId, address[] calldata srcModules) external view returns (uint256) {}

    /// @inheritdoc IInterchainDB
    function getWriterNonce(address writer) external view returns (uint256) {}

    /// @inheritdoc IInterchainDB
    function readEntry(
        InterchainEntry memory entry,
        address[] memory dstModules
    )
        external
        view
        returns (uint256[] memory moduleConfirmedAt)
    {}
}
