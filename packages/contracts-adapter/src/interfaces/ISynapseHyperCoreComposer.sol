// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseHyperCoreComposer {
    function bridgeToHyperCore(address recipient, uint256 amount) external;
    function adapter() external view returns (address);
    function token() external view returns (address);
}
