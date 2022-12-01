// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SystemContract } from "../../../contracts/system/SystemContract.sol";

abstract contract SystemContractMockEvents {
    event SlashAgentCall(SystemContract.AgentInfo info);

    event SyncAgentsCall(uint256 requestID, bool removeExisting, SystemContract.AgentInfo[] infos);
}
