// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseScript, StringUtils} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

abstract contract DeployProxy is SynapseScript {
    using StringUtils for *;

    bytes32 public constant ADMIN_SLOT = bytes32(uint256(keccak256("eip1967.proxy.admin")) - 1);

    function deployAndSaveProxy(
        string memory contractName,
        address implementation,
        address proxyAdminOwner,
        bytes memory initData
    )
        internal
        returns (address deployedAt)
    {
        // Deploy the proxy and save the artifact with Proxy ABI as ContractName.Proxy
        deployedAt = deployAndSaveAs({
            contractName: "TransparentUpgradeableProxy",
            contractAlias: contractName.concat(".Proxy"),
            constructorArgs: abi.encode(implementation, proxyAdminOwner, initData),
            deployCodeFunc: cbDeploy
        });
        // Save the proxy artifact with implementation ABI as ContractName
        if (!isDeployed(contractName)) {
            saveDeployment({
                contractName: contractName,
                contractAlias: contractName,
                deployedAt: deployedAt,
                constructorArgs: ""
            });
        }
        // Save the ProxyAdmin artifact as ContractName.ProxyAdmin
        string memory proxyAdminAlias = contractName.concat(".ProxyAdmin");
        if (!isDeployed(proxyAdminAlias)) {
            saveDeployment({
                contractName: "ProxyAdmin",
                contractAlias: proxyAdminAlias,
                deployedAt: getProxyAdmin(deployedAt),
                constructorArgs: abi.encode(proxyAdminOwner)
            });
        }
    }

    function getProxyAdmin(address target) internal view returns (address admin) {
        bytes32 adminBytes32 = vm.load(target, ADMIN_SLOT);
        admin = address(uint160(uint256(adminBytes32)));
    }
}
