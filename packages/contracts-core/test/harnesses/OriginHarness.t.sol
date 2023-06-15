// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Origin} from "../../contracts/Origin.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract OriginHarness is Origin {
    constructor(uint32 domain, address agentManager_, address inbox_, address gasOracle_)
        Origin(domain, agentManager_, inbox_, gasOracle_)
    {} // solhint-disable-line no-empty-blocks

    // TODO: add / remove Agents in standalone Go tests
}
