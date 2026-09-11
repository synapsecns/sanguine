// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "./SynapseOFTAdapter.sol";
import {ISynapseOFTAdapterFactory} from "./interfaces/ISynapseOFTAdapterFactory.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @notice Deploys SynapseERC20 OFT adapters at deterministic addresses independent of token and endpoint addresses.
/// @dev Deploy with CREATE2 using the same initial owner and salt, then initialize each chain's endpoint.
contract SynapseOFTAdapterFactory is ISynapseOFTAdapterFactory, Ownable {
    address public endpoint;

    ConstructorParams private constructorParams;

    event EndpointInitialized(address indexed endpoint);

    error AlreadyInitialized();
    error InvalidEndpoint();
    error NotInitialized();

    /// @param initialOwner The account authorized to deploy adapters.
    constructor(address initialOwner) Ownable(initialOwner) {}

    /// @notice Sets the LayerZero V2 endpoint once. Only callable by the factory owner.
    /// @param lzEndpoint The endpoint used by all adapters deployed by this factory.
    function initialize(address lzEndpoint) external onlyOwner {
        if (endpoint != address(0)) revert AlreadyInitialized();
        if (lzEndpoint == address(0)) revert InvalidEndpoint();
        endpoint = lzEndpoint;
        emit EndpointInitialized(lzEndpoint);
    }

    /// @inheritdoc ISynapseOFTAdapterFactory
    function deploy(address token, bytes32 salt) external onlyOwner returns (address adapter) {
        if (endpoint == address(0)) revert NotInitialized();
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
