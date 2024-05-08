// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IGasOracle} from "./IGasOracle.sol";

interface ISynapseGasOracle is IGasOracle {
    function receiveRemoteGasData(uint64 srcChainId, bytes calldata data) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function getLocalGasData() external view returns (bytes memory);
}
