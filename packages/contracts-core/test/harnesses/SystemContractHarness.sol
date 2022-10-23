// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SystemContract } from "../../contracts/system/SystemContract.sol";
import { ISystemRouter } from "../../contracts/interfaces/ISystemRouter.sol";

abstract contract SystemContractHarness is SystemContract {
    uint256 public sensitiveValue;

    // events used for testing
    event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt);

    event UsualCall(address recipient, uint256 newValue);
    event OnlyLocalCall(address recipient, uint256 newValue);
    event OnlyOriginCall(address recipient, uint256 newValue);
    event OnlyDestinationCall(address recipient, uint256 newValue);
    event OnlyTwoHoursCall(address recipient, uint256 newValue);
    event OnlySynapseChainCall(address recipient, uint256 newValue);

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
        emit OnlyDestinationCall(address(this), _newValue);
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
