// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/system/SystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterHarness is SystemRouter {
    using ByteString for bytes;

    constructor(
        uint32 _domain,
        address _origin,
        address _destination,
        address _bondingManager
    ) SystemRouter(_domain, _origin, _destination, _bondingManager) {}

    /**
     * @notice Mocks a system call from the given caller on the given origin chain.
     */
    function mockSystemCall(
        SystemEntity _recipient,
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _systemCaller,
        bytes memory _data
    ) public {
        bytes memory prefix = abi.encode(_rootSubmittedAt, _callOrigin, _systemCaller);
        _localSystemCall({
            _recipient: uint8(_recipient),
            _callData: _data.castToCallData(),
            _prefix: prefix.castToRawBytes()
        });
    }
}
