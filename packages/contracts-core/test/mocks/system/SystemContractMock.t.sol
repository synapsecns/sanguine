// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../../../contracts/interfaces/ISystemRouter.sol";
import { SystemContract } from "../../../contracts/system/SystemContract.sol";
import { LocalDomainContext } from "../../../contracts/context/LocalDomainContext.sol";
import "../events/SystemContractMockEvents.sol";

// solhint-disable no-empty-blocks
contract SystemContractMock is SystemContractMockEvents, SystemContract, LocalDomainContext {
    // Expose internal constants for tests
    uint256 public constant ORIGIN_MASK = ORIGIN;
    uint256 public constant DESTINATION_MASK = DESTINATION;
    uint256 public constant BONDING_MANAGER_MASK = BONDING_MANAGER;

    uint256 public constant BONDING_OPTIMISTIC_PERIOD_PUB = BONDING_OPTIMISTIC_PERIOD;

    constructor(uint32 _domain) LocalDomainContext(_domain) {}

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    // Expose modifiers for tests
    function mockOnlySystemRouter() external onlySystemRouter {}

    function mockOnlySynapseChain(uint32 domain) external onlySynapseChain(domain) {}

    function mockOnlyCallers(uint256 mask, ISystemRouter.SystemEntity caller)
        external
        onlyCallers(mask, caller)
    {}

    function mockOnlyOptimisticPeriodOver(uint256 rootSubmittedAt, uint256 optimisticSeconds)
        external
        onlyOptimisticPeriodOver(rootSubmittedAt, optimisticSeconds)
    {}

    function slashAgent(
        uint256,
        uint32,
        ISystemRouter.SystemEntity,
        AgentInfo memory _info
    ) external override {
        emit SlashAgentCall(_info);
    }

    function syncAgents(
        uint256,
        uint32,
        ISystemRouter.SystemEntity,
        uint256 _requestID,
        bool _removeExisting,
        AgentInfo[] memory _infos
    ) external override {
        emit SyncAgentsCall(_requestID, _removeExisting, _infos);
    }
}
