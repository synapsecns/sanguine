// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Origin} from "../../contracts/Origin.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract OriginHarness is Origin {
    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 domain, address agentManager_, address gasOracle_) Origin(domain, agentManager_, gasOracle_) {}

    // TODO: add / remove Agents in standalone Go tests
}
