// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "./SynapseOFTAdapter.sol";
import {ISynapseOFTAdapterFactory} from "./interfaces/ISynapseOFTAdapterFactory.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @notice Deploys SynapseERC20 OFT adapters at deterministic addresses independent of the token address.
/// @dev Deploy this factory with CREATE2 using the same constructor arguments and salt on each chain.
contract SynapseOFTAdapterFactory is ISynapseOFTAdapterFactory, Ownable {
    address public immutable endpoint;

    ConstructorParams private constructorParams;

    /// @param initialOwner The account authorized to deploy adapters.
    /// @param lzEndpoint The LayerZero V2 endpoint used by all adapters deployed by this factory.
    constructor(address initialOwner, address lzEndpoint) Ownable(initialOwner) {
        endpoint = lzEndpoint;
    }

    /// @inheritdoc ISynapseOFTAdapterFactory
    function deploy(address token, bytes32 salt) external onlyOwner returns (address adapter) {
        constructorParams = ConstructorParams({token: token, lzEndpoint: endpoint, owner: owner()});
        adapter = address(new SynapseOFTAdapter{salt: salt}());
        delete constructorParams;
        emit AdapterDeployed(token, adapter, salt);
    }

    /// @inheritdoc ISynapseOFTAdapterFactory
    function getConstructorParams() external view returns (ConstructorParams memory) {
        return constructorParams;
    }
}
