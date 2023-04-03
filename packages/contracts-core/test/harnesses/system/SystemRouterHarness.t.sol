// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {ByteString, SystemEntity, SystemRouter} from "../../../contracts/system/SystemRouter.sol";

// solhint-disable no-empty-blocks
contract SystemRouterHarness is SystemRouter {
    using ByteString for bytes;

    constructor(uint32 domain, address origin_, address destination_, address agentManager_)
        SystemRouter(domain, origin_, destination_, agentManager_)
    {}

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the SystemRouter.

    /**
     * @notice Pranks a system call: calls a local system recipient, as if the system call
     * was initiated by the given caller on the given origin chain.
     */
    function systemPrank(
        SystemEntity recipient,
        uint256 rootSubmittedAt,
        uint32 callOrigin,
        SystemEntity systemCaller,
        bytes memory payload
    ) public {
        bytes memory prefix = abi.encode(rootSubmittedAt, callOrigin, systemCaller);
        _callSystemRecipient({systemRecipient: recipient, payload: payload.castToCallData().addPrefix(prefix)});
    }
}
