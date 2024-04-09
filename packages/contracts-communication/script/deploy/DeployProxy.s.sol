// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseScript, StringUtils} from "@synapsecns/solidity-devops/src/SynapseScript.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// solhint-disable custom-errors
abstract contract DeployProxy is SynapseScript {
    using StringUtils for *;

    bytes32 public constant ADMIN_SLOT = bytes32(uint256(keccak256("eip1967.proxy.admin")) - 1);

    /// @notice Deploys a TransparentUpgradeableProxy contract, then saves the fresh deployment artifacts for:
    /// - The proxy as ContractName.Proxy (with Proxy ABI)
    /// - The proxy as ContractName (with implementation ABI)
    /// - The ProxyAdmin as ContractName.ProxyAdmin
    function deployAndSaveProxy(
        string memory contractName,
        address implementation,
        address proxyAdminOwner,
        bytes memory initData
    )
        internal
        returns (address deployedAt)
    {
        require(implementation != address(0), "No implementation provided");
        require(proxyAdminOwner != address(0), "No proxy admin owner provided");
        // Deploy the proxy and save the artifact with Proxy ABI as TransparentUpgradeableProxy.ContractName
        deployedAt = deployAndSaveAs({
            contractName: "TransparentUpgradeableProxy",
            contractAlias: string.concat("TransparentUpgradeableProxy.", contractName),
            constructorArgs: abi.encode(implementation, proxyAdminOwner, initData),
            deployCodeFunc: cdDeployTransparentUpgradeableProxy
        });
        // Save the proxy artifact with implementation ABI as ContractName
        if (!isDeployed(contractName)) {
            saveDeployment({contractAlias: contractName, deployedAt: deployedAt, constructorArgs: ""});
        }
        // Save the ProxyAdmin artifact as ProxyAdmin.ContractName
        string memory proxyAdminAlias = string.concat("ProxyAdmin.", contractName);
        if (!isDeployed(proxyAdminAlias)) {
            saveDeployment({
                contractAlias: proxyAdminAlias,
                deployedAt: getProxyAdmin(deployedAt),
                constructorArgs: abi.encode(proxyAdminOwner)
            });
        }
    }

    /// @notice Callback to deploy a TransparentUpgradeableProxy contract.
    /// Note: this version of TransparentUpgradeableProxy spawns its own ProxyAdmin contract,
    /// and therefore the ProxyAdmin contract doesn't need to be deployed separately.
    function cdDeployTransparentUpgradeableProxy(
        string memory contractName,
        bytes memory constructorArgs
    )
        internal
        returns (address deployedAt)
    {
        require(
            contractName.equals("TransparentUpgradeableProxy"), string.concat("Invalid contract name: ", contractName)
        );
        (address implementation, address proxyAdminOwner, bytes memory initData) =
            abi.decode(constructorArgs, (address, address, bytes));
        deployedAt = address(new TransparentUpgradeableProxy(implementation, proxyAdminOwner, initData));
    }

    /// @notice Retrieves the address from the ADMIN_SLOT of a TransparentUpgradeableProxy contract.
    /// This is the address that is able to upgrade the proxy.
    function getProxyAdmin(address target) internal view returns (address admin) {
        bytes32 adminBytes32 = vm.load(target, ADMIN_SLOT);
        admin = address(uint160(uint256(adminBytes32)));
    }
}
