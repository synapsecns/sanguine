// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentInfo, SystemContract } from "../../../contracts/system/SystemContract.sol";

abstract contract SystemContractTools {
    SystemContract internal systemContract;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function infoToArray(AgentInfo memory info) public pure returns (AgentInfo[] memory infos) {
        infos = new AgentInfo[](1);
        infos[0] = info;
    }

    function guardInfo(address guard, bool bonded) public pure returns (AgentInfo memory) {
        return agentInfo(0, guard, bonded);
    }

    function agentInfo(
        uint32 domain,
        address account,
        bool bonded
    ) public pure returns (AgentInfo memory) {
        return AgentInfo({ bonded: bonded, domain: domain, account: account });
    }
}
