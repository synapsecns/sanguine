// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract BasicClientHarnessEvents {
    event LogBasicClientMessage(
        uint32 origin,
        uint32 nonce,
        uint256 rootSubmittedAt,
        bytes content
    );
}
