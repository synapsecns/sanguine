// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { ByteString, SystemEntity, SystemRouter } from "../../../contracts/system/SystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterHarness is SystemRouter {
    using ByteString for bytes;

    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRouterHarness() external {}

    constructor(
        uint32 _domain,
        address _origin,
        address _destination,
        address _agentManager
    ) SystemRouter(_domain, _origin, _destination, _agentManager) {}

    /**
     * @notice Pranks a system call: calls a local system recipient, as if the system call
     * was initiated by the given caller on the given origin chain.
     */
    function systemPrank(
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
