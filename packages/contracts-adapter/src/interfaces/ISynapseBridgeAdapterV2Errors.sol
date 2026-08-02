// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridgeAdapterV2Errors {
    error SBAV2__HyperCoreAmountExceedsUint64(uint256 coreAmount);
    error SBAV2__HyperCoreAmountNotRepresentable(uint256 amount, uint256 scale);
    error SBAV2__HyperCoreComposerAlreadySet(address token);
    error SBAV2__HyperCoreComposerInvalid(address composer);
    error SBAV2__HyperCoreComposerNotSet(address token);
    error SBAV2__HyperCoreDecimalsAlreadySet(uint32 eid, address token);
    error SBAV2__HyperCoreDecimalsInvalid(uint8 tokenDecimals, uint8 coreDecimals);
    error SBAV2__HyperCoreDecimalsNotSet(uint32 eid, address token);
}
