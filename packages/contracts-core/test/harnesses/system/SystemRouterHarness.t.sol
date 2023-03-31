// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { ByteString, SystemEntity, SystemRouter } from "../../../contracts/system/SystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterHarness is SystemRouter {
    using ByteString for bytes;

    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRouterHarness() external {}

    constructor(
        uint32 domain,
        address origin,
        address destination,
        address _agentManager
    ) SystemRouter(domain, origin, destination, _agentManager) {}

    /**
     * @notice Pranks a system call: calls a local system recipient, as if the system call
     * was initiated by the given caller on the given origin chain.
     */
    function systemPrank(
        SystemEntity recipient,
        uint256 rootSubmittedAt,
        uint32 callOrigin,
        SystemEntity systemCaller,
        bytes memory data
    ) public {
        bytes memory prefix = abi.encode(rootSubmittedAt, callOrigin, systemCaller);
        _localSystemCall({
            recipient: uint8(recipient),
            _callData: data.castToCallData(),
            _prefix: prefix.castToRawBytes()
        });
    }
}
