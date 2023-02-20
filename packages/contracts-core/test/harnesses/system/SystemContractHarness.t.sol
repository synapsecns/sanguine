// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SystemContract } from "../../../contracts/system/SystemContract.sol";
import { ISystemRouter } from "../../../contracts/interfaces/ISystemRouter.sol";
import { DomainContext } from "../../../contracts/context/DomainContext.sol";
import { SystemContractHarnessEvents } from "../events/SystemContractHarnessEvents.sol";

abstract contract SystemContractHarness is SystemContractHarnessEvents, SystemContract {
    uint256 public sensitiveValue;

    function setSensitiveValue(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    ) external onlySystemRouter {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit UsualCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyLocal(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    ) external onlySystemRouter onlyLocalDomain(_callOrigin) {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit OnlyLocalCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyOrigin(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    ) external onlySystemRouter onlyCallers(ORIGIN, ISystemRouter.SystemEntity(_systemCaller)) {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit OnlyOriginCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyDestination(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    )
        external
        onlySystemRouter
        onlyCallers(DESTINATION, ISystemRouter.SystemEntity(_systemCaller))
    {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit OnlyDestinationCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyOriginDestination(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    )
        external
        onlySystemRouter
        onlyCallers(ORIGIN | DESTINATION, ISystemRouter.SystemEntity(_systemCaller))
    {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit OnlyOriginDestinationCall(address(this), _newValue);
    }

    function setSensitiveValueOnlyTwoHours(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    ) external onlySystemRouter onlyOptimisticPeriodOver(_rootSubmittedAt, 2 hours) {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit OnlyTwoHoursCall(address(this), _newValue);
    }

    function setSensitiveValueOnlySynapseChain(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    ) external onlySystemRouter onlySynapseChain(_callOrigin) {
        _setSensitiveValue(_rootSubmittedAt, _callOrigin, _systemCaller, _newValue);
        emit OnlySynapseChainCall(address(this), _newValue);
    }

    function _setSensitiveValue(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        uint8 _systemCaller,
        uint256 _newValue
    ) internal {
        sensitiveValue = _newValue;
        emit LogSystemCall(_callOrigin, _systemCaller, _rootSubmittedAt);
    }
}
