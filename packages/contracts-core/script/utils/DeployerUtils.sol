// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {console, Script, stdJson} from "forge-std/Script.sol";

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {CREATE3Factory} from "create3/CREATE3Factory.sol";

interface ICreate3Factory {
    function deploy(bytes32 salt, bytes memory creationCode) external payable returns (address deployed);

    function getDeployed(address deployer, bytes32 salt) external view returns (address deployed);
}

// solhint-disable no-console
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract DeployerUtils is Script {
    using stdJson for string;

    /// @dev Path to artifacts, deployments and configs directories
    string private constant ARTIFACTS = "artifacts/";
    string private constant DEPLOYMENTS = "deployments/";
    string private deployConfigs = "script/configs/";

    // @dev wether or not devnet is enabled
    bool private devnetEnabled = false;
    // @dev env var for wether or not devnet is in use
    string private constant DEVNET_ENABLED_VAR = "DEVNET";
    // @dev artifacted path to use if devnet is enabled;
    string private constant DEPLOY_CONFIGS_DEVNET = "script/configs/devnet/";

    /// @dev Whether the script will be broadcasted or not
    bool internal isBroadcasted = false;
    /// @dev Current chain alias
    string internal chainAlias;

    /// @dev Private key and address for deploying contracts
    uint256 internal broadcasterPK;
    address internal broadcasterAddress;

    ICreate3Factory private factory = ICreate3Factory(0x9fBB3DF7C40Da2e5A0dE984fFE2CCB7C47cd0ABf);

    bytes32 internal deploymentSalt;

    /// @notice Prevents this contract from being included in the coverage report
    function testDeployerUtils() external {}

    // ═══════════════════════════════════════════════════ SETUP ═══════════════════════════════════════════════════════

    function stopBroadcast() public {
        vm.stopBroadcast();
        isBroadcasted = false;
    }

    function startBroadcast(bool isBroadcasted_) public {
        chainAlias = getChainAlias();
        if (isBroadcasted_) createDir(string.concat(DEPLOYMENTS, chainAlias));
        vm.startBroadcast(broadcasterPK);
        isBroadcasted = isBroadcasted_;
    }

    // @dev this is called just in time so we can make sure that startBroadcast() is called before this.
    // it's only overriden and deployed just-in-time in the case of devnet
    // TODO: this pattern sucks. It introduces potential unexpected behavior if the dev calls factory. directly.
    // It's also slow.
    function getFactory() internal returns (ICreate3Factory) {
        if (!devnetEnabled) {
            return factory;
        }

        address factoryDeployment = tryLoadDeployment("CREATE3Factory");
        if (factoryDeployment == address(0)) {
            if (broadcasterPK == 0) {
                console.log("please setup a private key before calling this function");
            }

            console.log("Create3Factory not deployed on devnet, deploying now");
            CREATE3Factory NewFactory = new CREATE3Factory();
            saveDeployment("Create3Factory", "Create3Factory", address(NewFactory), "0x");
            factoryDeployment = address(NewFactory);
        }
        factory = ICreate3Factory(factoryDeployment);
        return factory;
    }

    // @dev must be called after setupPK()
    function setupDevnetIfEnabled() internal {
        devnetEnabled = vm.envOr(DEVNET_ENABLED_VAR, false);

        if (devnetEnabled) {
            devnetEnabled = true;
            // setup the chains
            setChain("chain_a", Chain("chain_a", 42, "chain_a", "http://localhost:9001/rpc/42"));
            setChain("chain_b", Chain("chain_b", 43, "chain_b", "http://localhost:9001/rpc/43"));
            setChain("chain_c", Chain("chain_c", 44, "chain_c", "http://localhost:9001/rpc/44"));

            // override the configs path
            deployConfigs = DEPLOY_CONFIGS_DEVNET;

            chainAlias = getChainAlias();
        }
    }

    function setupDeployerPK() public {
        setupPK("DEPLOYER_PRIVATE_KEY");
    }

    function setupPK(string memory pkEnvKey) public {
        // Load deployer PK from .env
        broadcasterPK = vm.envOr(pkEnvKey, uint256(0));
        if (broadcasterPK == 0) {
            console.log("Key not specified in .env: %s", pkEnvKey);
        } else {
            // Derive deployer address
            broadcasterAddress = vm.addr(broadcasterPK);
            console.log("Deployer address: %s", broadcasterAddress);
            console.log("Deployer balance: %s", _fromWei(broadcasterAddress.balance));
        }
    }

    /// @notice Returns name of the current chain.
    function getChainAlias() public returns (string memory) {
        return getChain(block.chainid).chainAlias;
    }

    // ════════════════════════════════════════════════ DEPLOYMENTS ════════════════════════════════════════════════════

    /// @notice Deploys the contract using Create3Factory. Does not save anything.
    function factoryDeploy(string memory contractName, bytes memory creationCode, bytes memory constructorArgs)
        internal
        returns (address deployment)
    {
        require(Address.isContract(address(getFactory())), "Factory not deployed");
        deployment = getFactory().deploy(
            getDeploymentSalt(contractName), // salt
            abi.encodePacked(creationCode, constructorArgs) // creation code with appended constructor args
        );
        require(deployment != address(0), "Factory deployment failed");
    }

    /// @notice Gets the deployment salt for a given contract.
    function getDeploymentSalt(string memory contractName) internal view returns (bytes32) {
        return keccak256(bytes.concat(deploymentSalt, bytes(contractName)));
    }

    /// @notice Predicts the deployment address for a contract.
    function predictFactoryDeployment(string memory contractName) internal returns (address) {
        ICreate3Factory _factory = getFactory();
        require(Address.isContract(address(_factory)), "Factory not deployed");
        return _factory.getDeployed(broadcasterAddress, getDeploymentSalt(contractName));
    }

    /// @notice Deploys the contract and saves the deployment artifact
    /// @dev Will reuse existing deployment, if it exists
    /// @param contractName     Contract name to deploy
    /// @param deployFunc       Callback function to deploy a requested contract
    /// @return deployment  The deployment address
    function deployContract(
        string memory contractName,
        function() internal returns (address, bytes memory) deployFunc,
        function(address) internal initFunc
    ) internal returns (address deployment, bytes memory constructorArgs) {
        return deployContract(contractName, contractName, deployFunc, initFunc);
    }

    /// @notice Deploys the contract and saves the deployment artifact under specified name
    /// @dev Will reuse existing deployment, if it exists
    /// @param contractName     Contract name to deploy
    /// @param saveAsName       Name to use for saving the deployment artifact
    /// @param deployFunc       Callback function to deploy a requested contract
    /// @param initFunc         Callback function to initialize a deployed contract
    /// @return deployment  The deployment address
    function deployContract(
        string memory contractName,
        string memory saveAsName,
        function() internal returns (address, bytes memory) deployFunc,
        function(address) internal initFunc
    ) internal returns (address deployment, bytes memory constructorArgs) {
        deployment = tryLoadDeployment(saveAsName);
        if (deployment == address(0)) {
            (deployment, constructorArgs) = deployFunc();
            saveDeployment(contractName, saveAsName, deployment, constructorArgs);
        } else {
            console.log("Reusing existing deployment for %s: %s", contractName, deployment);
        }
        vm.label(deployment, contractName);
        initFunc(deployment);
    }

    /// @notice Returns the deployment for a contract on the current chain, if it exists.
    /// Reverts if it doesn't exist.
    function loadDeployment(string memory contractName) public returns (address deployment) {
        deployment = tryLoadDeployment(contractName);
        require(deployment != address(0), string.concat(contractName, " doesn't exist on ", chainAlias));
    }

    /// @notice Returns the deployment for a contract on the current chain, if it exists.
    /// Returns address(0), if it doesn't exist.
    function tryLoadDeployment(string memory contractName) public returns (address deployment) {
        try vm.readFile(deploymentPath(contractName)) returns (string memory json) {
            // We assume that if a deployment file exists, the contract is indeed deployed
            deployment = json.readAddress(".address");
        } catch {
            // Doesn't exist
            deployment = address(0);
        }
    }

    /// @notice Saves the deployment JSON for a deployed contract.
    function saveDeployment(
        string memory contractName,
        string memory saveAsName,
        address deployedAt,
        bytes memory constructorArgs
    ) public {
        console.log("Deployed: [%s] on [%s] at %s", contractName, chainAlias, deployedAt);
        // Do nothing if script isn't broadcasted
        if (!isBroadcasted) return;
        // Otherwise, save the deployment JSON
        string memory deployment = "deployment";
        // First, write only the deployment address and the constructor args
        deployment.serialize("address", deployedAt);
        deployment = deployment.serialize("args", constructorArgs);
        deployment.write(deploymentPath(saveAsName));
        // Then, initiate the jq command to add "abi" as the next key
        // This makes sure that "address" value is printed first later
        string[] memory inputs = new string[](6);
        inputs[0] = "jq";
        // Read the full artifact file into $artifact variable
        inputs[1] = "--argfile";
        inputs[2] = "artifact";
        inputs[3] = artifactPath(contractName);
        // Set value for ".abi" key to artifact's ABI
        inputs[4] = ".abi = $artifact.abi";
        inputs[5] = deploymentPath(saveAsName);
        bytes memory full = vm.ffi(inputs);
        // Finally, print the updated deployment JSON
        string(full).write(deploymentPath(saveAsName));
    }

    // ═══════════════════════════════════════════════ DEPLOY CONFIG ═══════════════════════════════════════════════════

    /// @notice Checks if deploy config exists for a given contract on the current chain.
    function deployConfigExists(string memory contractName) public returns (bool) {
        try vm.fsMetadata(deployConfigPath(contractName)) {
            return true;
        } catch {
            return false;
        }
    }

    /// @notice Loads deploy config for a given contract on the current chain.
    /// Will revert if config doesn't exist.
    function loadDeployConfig(string memory contractName) public returns (string memory json) {
        return vm.readFile(deployConfigPath(contractName));
    }

    /// @notice Saves deploy config for a given contract on the current chain.
    function saveDeployConfig(string memory contractName, string memory config) public {
        console.log("Saved: config for [%s] on [%s]", contractName, chainAlias);
        string memory path = deployConfigPath(contractName);
        vm.writeJson(config, path);
        // Sort keys in config JSON for consistency
        sortJSON(path);
    }

    /// @notice Loads deploy config for a given contract on the current chain.
    /// Will revert if config doesn't exist.
    function loadGlobalDeployConfig(string memory contractName) public returns (string memory json) {
        return vm.readFile(globalDeployConfigPath(contractName));
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns path to the contract artifact.
    function artifactPath(string memory contractName) public pure returns (string memory path) {
        return string.concat(ARTIFACTS, contractName, ".sol/", deploymentFn(contractName));
    }

    /// @notice Returns deployment filename for the contract.
    function deploymentFn(string memory contractName) public pure returns (string memory path) {
        return string.concat(contractName, ".json");
    }

    /// @notice Returns path to the contract deployment for the current chain.
    function deploymentPath(string memory contractName) public view returns (string memory path) {
        require(bytes(chainAlias).length != 0, "Chain not set");
        return string.concat(DEPLOYMENTS, chainAlias, "/", deploymentFn(contractName));
    }

    /// @notice Returns deploy config filename for the contract.
    function deployConfigFn(string memory contractName) public pure returns (string memory path) {
        return string.concat(contractName, ".dc.json");
    }

    /// @notice Returns path to the contract deploy config JSON on the current chain.
    function deployConfigPath(string memory contractName) internal returns (string memory path) {
        require(bytes(chainAlias).length != 0, "Chain not set");
        return string.concat(deployConfigs, chainAlias, "/", deployConfigFn(contractName));
    }

    /// @notice Returns path to the global contract deploy config JSON.
    function globalDeployConfigPath(string memory contractName) public returns (string memory path) {
        return string.concat(deployConfigs, deployConfigFn(contractName));
    }

    /// @notice Create directory if it not exists already
    function createDir(string memory dirPath) public {
        // solhint-disable-next-line no-empty-blocks
        try vm.fsMetadata(dirPath) {}
        catch {
            string[] memory inputs = new string[](3);
            inputs[0] = "mkdir";
            inputs[1] = "--p";
            inputs[2] = dirPath;
            vm.ffi(inputs);
        }
    }

    /// @dev Reads JSON from given path, sorts its keys and overwrites the file.
    function sortJSON(string memory path) public {
        string[] memory inputs = new string[](4);
        inputs[0] = "jq";
        // sort keys of objects on output
        inputs[1] = "-S";
        // The simplest filter is ., which copies jq's input to its output unmodified
        inputs[2] = ".";
        inputs[3] = path;
        bytes memory sorted = vm.ffi(inputs);
        string(sorted).write(path);
    }

    // ═════════════════════════════════════════════ INTERNAL HELPERS ══════════════════════════════════════════════════

    function _fromWei(uint256 amount) internal pure returns (string memory s) {
        string memory a = Strings.toString(amount / 10 ** 18);
        string memory b = Strings.toString(amount % 10 ** 18);
        // Add leading zeroes to the decimal part
        while (bytes(b).length < 18) {
            b = string.concat("0", b);
        }
        return string.concat(a, ".", b);
    }
}
