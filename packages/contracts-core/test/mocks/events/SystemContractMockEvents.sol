// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../../contracts/system/SystemContract.sol";

abstract contract SystemContractMockEvents {
    event SlashAgentCall(AgentInfo info);

    event SyncAgentsCall(uint256 requestID, bool removeExisting, AgentInfo[] infos);
}
