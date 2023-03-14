// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentInfo } from "../../../contracts/libs/Structures.sol";

import {
    SystemEntity,
    SystemRouterHarness
} from "../../harnesses/system/SystemRouterHarness.t.sol";

import { ISystemContract, SynapseTest } from "../../utils/SynapseTest.t.sol";

abstract contract BondingManagerTest is SynapseTest {
    uint256 internal rootSubmittedAt;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _localDomain() internal view virtual returns (uint32);

    function _skipBondingOptimisticPeriod() internal {
        _skipPeriod(BONDING_OPTIMISTIC_PERIOD);
    }

    function _skipPeriod(uint256 period) internal {
        rootSubmittedAt = block.timestamp;
        skip(period);
    }

    function _systemPrank(
        SystemRouterHarness router,
        uint32 callOrigin,
        SystemEntity systemCaller,
        bytes memory data
    ) internal {
        router.mockSystemCall({
            _recipient: SystemEntity.BondingManager,
            _rootSubmittedAt: callOrigin == _localDomain() ? block.timestamp : rootSubmittedAt,
            _callOrigin: callOrigin,
            _systemCaller: systemCaller,
            _data: data
        });
    }

    function _expectedCall(bytes4 selector, AgentInfo memory info)
        internal
        view
        returns (bytes memory)
    {
        return
            abi.encodeWithSelector(
                selector,
                block.timestamp,
                _localDomain(),
                SystemEntity.BondingManager,
                info
            );
    }

    function _dataSyncAgentCall(AgentInfo memory info) internal pure returns (bytes memory) {
        return
            abi.encodeWithSelector(
                ISystemContract.syncAgent.selector,
                0, // rootSubmittedAt
                0, // callOrigin
                0, // systemCaller
                info
            );
    }

    function _dataSlashAgentCall(AgentInfo memory info) internal pure returns (bytes memory) {
        return
            abi.encodeWithSelector(
                ISystemContract.slashAgent.selector,
                0, // rootSubmittedAt
                0, // callOrigin
                0, // systemCaller
                info
            );
    }
}
