// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentFlag } from "../libs/Structures.sol";

abstract contract BondingManagerEvents {
    event StatusUpdated(
        AgentFlag flag,
        uint32 indexed domain,
        address indexed agent,
        bytes32 newRoot
    );
}
