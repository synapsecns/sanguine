// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {FastBridgeInterceptor} from "../contracts/FastBridgeInterceptor.sol";
import {SynapseIntentPreviewer} from "../contracts/router/SynapseIntentPreviewer.sol";
import {SynapseIntentRouter} from "../contracts/router/SynapseIntentRouter.sol";
import {TokenZapV1} from "../contracts/zaps/TokenZapV1.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks, ordering
contract DeploySIR is SynapseScript {
    using stdJson for string;

    string public constant LATEST_SIR = "SynapseIntentRouter";
    string public constant LATEST_SIP = "SynapseIntentPreviewer";
    string public constant LATEST_ZAP = "TokenZapV1";

    string public constant LATEST_FBI = "FastBridgeInterceptor";

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDeploySIR() external {}

    function run() external broadcastWithHooks {
        address sir = deployAndSaveWithSalt(LATEST_SIR, type(SynapseIntentRouter).creationCode, "");
        address sip = deployAndSaveWithSalt(LATEST_SIP, type(SynapseIntentPreviewer).creationCode, "");
        address zap = deployAndSaveWithSalt(LATEST_ZAP, type(TokenZapV1).creationCode, "");
        address fbi = deployAndSaveWithSalt(LATEST_FBI, type(FastBridgeInterceptor).creationCode, "");
        printLog("Checking deployments");
        checkDeployment(LATEST_SIR, sir, type(SynapseIntentRouter).runtimeCode);
        checkDeployment(LATEST_SIP, sip, type(SynapseIntentPreviewer).runtimeCode);
        checkDeployment(LATEST_ZAP, zap, type(TokenZapV1).runtimeCode);
        checkDeployment(LATEST_FBI, fbi, type(FastBridgeInterceptor).runtimeCode);
    }

    // TODO: upstream to solidity-devops
    function deployAndSaveWithSalt(
        string memory contractName,
        bytes memory creationCode,
        bytes memory constructorArgs
    )
        internal
        returns (address deployedAt)
    {
        string memory config = readGlobalDeployProdConfig({contractName: "Create2Factory", revertIfNotFound: true});
        string memory keyPrefix = string.concat(".salts.", contractName);
        address predictedAddress;
        if (vm.keyExistsJson(config, keyPrefix)) {
            bytes32 salt = config.readBytes32({key: string.concat(keyPrefix, ".salt")});
            predictedAddress = config.readAddress({key: string.concat(keyPrefix, ".predictedAddress")});
            setNextDeploymentSalt(salt);
        } else {
            printInfo(string.concat("No salt found for ", contractName, ". Using zero salt."));
        }
        deployedAt = deployAndSave({
            contractName: contractName,
            constructorArgs: constructorArgs,
            deployCodeFunc: cbDeployCreate2
        });
        if (predictedAddress != address(0)) {
            checkCreationCodeHash({
                creationCode: creationCode,
                expectedCreationCodeHash: config.readBytes32({key: string.concat(keyPrefix, ".creationCodeHash")})
            });
            if (deployedAt != predictedAddress) {
                printFailWithIndent("Deployed address does not match predicted address");
                assert(false);
            }
        }
    }

    // TODO: upstream to solidity-devops
    function checkDeployment(string memory contractName, address deployedAt, bytes memory runtimeCode) internal view {
        if (keccak256(deployedAt.code) != keccak256(runtimeCode)) {
            printFailWithIndent(
                string.concat(contractName, " deployment at ", vm.toString(deployedAt), " is not the latest version")
            );
            assert(false);
        }
        printSuccessWithIndent(string.concat(contractName, " latest version deployed at ", vm.toString(deployedAt)));
    }

    // TODO: upstream to solidity-devops
    function checkCreationCodeHash(bytes memory creationCode, bytes32 expectedCreationCodeHash) internal view {
        bytes32 creationCodeHash = keccak256(creationCode);
        if (creationCodeHash != expectedCreationCodeHash) {
            printFailWithIndent("Creation code hash does not match");
            printFailWithIndent(string.concat("  Expected: ", vm.toString(expectedCreationCodeHash)));
            printFailWithIndent(string.concat("    Actual: ", vm.toString(creationCodeHash)));
            assert(false);
        }
    }
}
