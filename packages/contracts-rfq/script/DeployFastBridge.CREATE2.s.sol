// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridge} from "../contracts/FastBridge.sol";

// TODO: this should work without remapping
import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract DeployFastBridgeCREATE2 is SynapseScript {
    using stdJson for string;

    string public constant NAME = "FastBridge";

    address public admin;
    FastBridge public fastBridge;

    function run() external broadcastWithHooks {
        loadConfig();
        bytes memory constructorArgs = abi.encode(admin);
        // Use CREATE2 Factory to deploy the contract
        fastBridge = FastBridge(
            deployAndSave({contractName: NAME, constructorArgs: constructorArgs, deployCodeFunc: cbDeployCreate2})
        );
    }

    function afterExecution() internal override {
        printLog(string.concat("Checking: ", NAME));
        checkAdminCount();
        checkAdmin();
        super.afterExecution();
    }

    function loadConfig() internal {
        string memory config = readGlobalDeployConfig({contractName: NAME, globalProperty: "", revertIfNotFound: true});
        admin = config.readAddress(".accounts.admin");
    }

    function checkAdminCount() internal view {
        uint256 count = fastBridge.getRoleMemberCount(0);
        string memory statement = string.concat("Admin count is ", vm.toString(count));
        if (count != 1) {
            printFailWithIndent(string.concat(statement, " instead of 1"));
            assert(false);
        }
        printSuccessWithIndent(statement);
    }

    function checkAdmin() internal view {
        address adminAddress = fastBridge.getRoleMember(0, 0);
        string memory statement = string.concat("Admin address is ", vm.toString(adminAddress));
        if (adminAddress != admin) {
            printFailWithIndent(string.concat(statement, " instead of ", vm.toString(admin)));
            assert(false);
        }
        printSuccessWithIndent(statement);
    }
}
