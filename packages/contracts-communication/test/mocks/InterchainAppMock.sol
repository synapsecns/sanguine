// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";

// solhint-disable no-empty-blocks
contract InterchainAppMock is IInterchainApp {
    function appReceive(
        uint64 srcChainId,
        bytes32 sender,
        uint64 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable
    {}

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {}
}
