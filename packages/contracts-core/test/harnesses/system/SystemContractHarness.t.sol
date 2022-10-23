// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SystemContract } from "../../../contracts/system/SystemContract.sol";
import { ISystemRouter } from "../../../contracts/interfaces/ISystemRouter.sol";
import { LocalDomainContext } from "../../../contracts/context/LocalDomainContext.sol";
import { SystemContractHarnessEvents } from "../events/SystemContractHarnessEvents.sol";

abstract contract SystemContractHarness is SystemContractHarnessEvents, SystemContract {
    uint256 public sensitiveValue;

    function setSensitiveValue(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) external onlySystemRouter {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit UsualCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyLocal(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) external onlySystemRouter onlyLocalDomain(_origin) {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit OnlyLocalCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyOrigin(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) external onlySystemRouter onlyCallers(ORIGIN, ISystemRouter.SystemEntity(_caller)) {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit OnlyOriginCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyDestination(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) external onlySystemRouter onlyCallers(DESTINATION, ISystemRouter.SystemEntity(_caller)) {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit OnlyDestinationCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyOriginDestination(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    )
        external
        onlySystemRouter
        onlyCallers(ORIGIN | DESTINATION, ISystemRouter.SystemEntity(_caller))
    {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit OnlyOriginDestinationCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyTwoHours(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) external onlySystemRouter onlyOptimisticPeriodOver(_rootSubmittedAt, 2 hours) {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit OnlyTwoHoursCall(address(this), _newValue);
    }

    function setSensitiveValueOnlySynapseChain(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) external onlySystemRouter onlySynapseChain(_origin) {
        _setSensitiveValue(_newValue, _origin, _caller, _rootSubmittedAt);
        emit OnlySynapseChainCall(address(this), _newValue);
    }

    function _setSensitiveValue(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller,
        uint256 _rootSubmittedAt
    ) internal {
        sensitiveValue = _newValue;
        emit LogSystemCall(_origin, _caller, _rootSubmittedAt);
    }
}

// solhint-disable no-empty-blocks
contract SystemContractMock is LocalDomainContext, SystemContractHarness {
    // Expose internal constants for tests
    uint256 public constant ORIGIN_MASK = ORIGIN;
    uint256 public constant DESTINATION_MASK = DESTINATION;

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
}
