// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../../contracts/system/SystemContract.sol";

abstract contract SystemContractTools {
    SystemContract internal systemContract;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function infoToArray(SystemContract.AgentInfo memory info)
        public
        pure
        returns (SystemContract.AgentInfo[] memory infos)
    {
        infos = new SystemContract.AgentInfo[](1);
        infos[0] = info;
    }

    function guardInfo(address guard, bool bonded)
        public
        pure
        returns (SystemContract.AgentInfo memory)
    {
        return agentInfo(0, guard, bonded);
    }

    function agentInfo(
        uint32 domain,
        address account,
        bool bonded
    ) public pure returns (SystemContract.AgentInfo memory) {
        return
            SystemContract.AgentInfo({
                agent: domain == 0 ? SystemContract.Agent.Guard : SystemContract.Agent.Notary,
                bonded: bonded,
                domain: domain,
                account: account
            });
    }
}
