// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridge} from "../contracts/FastBridge.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks
contract ConfigureFastBridge is SynapseScript {
    using stdJson for string;

    string public constant NAME = "FastBridge";

    FastBridge public fastBridge;
    string public config;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testConfigureFastBridge() external {}

    function run() external broadcastWithHooks {
        configureFB(ENVIRONMENT_PROD);
    }

    function runTestnet() external broadcastWithHooks {
        configureFB("testnet");
    }

    function configureFB(string memory environment) internal {
        loadConfig(environment);
        syncRole("governor", fastBridge.GOVERNOR_ROLE());
        setChainGasAmount();
        setProtocolFeeRate();
        syncRole("guard", fastBridge.GUARD_ROLE());
        syncRole("relayer", fastBridge.RELAYER_ROLE());
        syncRole("refunder", fastBridge.REFUNDER_ROLE());
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, environment: environment, revertIfNotFound: true});
        fastBridge = FastBridge(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
    }

    function setChainGasAmount() internal {
        printLog("Setting chain gas amount");
        uint256 chainGasAmount = config.readUint(".chainGasAmount");
        string memory action = string.concat("set to ", vm.toString(chainGasAmount));
        if (fastBridge.chainGasAmount() != chainGasAmount) {
            fastBridge.setChainGasAmount(chainGasAmount);
            printSuccessWithIndent(string.concat("Chain gas amount ", action));
        } else {
            printSkipWithIndent(string.concat("already ", action));
        }
    }

    function setProtocolFeeRate() internal {
        printLog("Setting protocol fee rate");
        uint256 protocolFeeRate = config.readUint(".protocolFeeRate");
        string memory action = string.concat("set to ", vm.toString(protocolFeeRate));
        if (fastBridge.protocolFeeRate() != protocolFeeRate) {
            fastBridge.setProtocolFeeRate(protocolFeeRate);
            printSuccessWithIndent(string.concat("Protocol fee rate ", action));
        } else {
            printSkipWithIndent(string.concat("already ", action));
        }
    }

    function syncRole(string memory roleName, bytes32 role) internal {
        string memory roleNamePlural = string.concat(roleName, "s");
        printLog(string.concat("Syncing ", roleNamePlural));
        address[] memory members = config.readAddressArray(string.concat(".accounts.", roleNamePlural));
        address[] memory existingMembers = getMembers(role);
        // Remove members that are not in the config
        uint256 removed = 0;
        for (uint256 i = 0; i < existingMembers.length; i++) {
            if (!contains(members, existingMembers[i])) {
                fastBridge.revokeRole(role, existingMembers[i]);
                printSuccessWithIndent(string.concat("Removed ", roleName, " ", vm.toString(existingMembers[i])));
                ++removed;
            }
        }
        // Add members that are in the config but not in the contract
        uint256 added = 0;
        for (uint256 i = 0; i < members.length; i++) {
            if (!contains(existingMembers, members[i])) {
                fastBridge.grantRole(role, members[i]);
                printSuccessWithIndent(string.concat("Added ", roleName, " ", vm.toString(members[i])));
                ++added;
            }
        }
        if (added + removed == 0) {
            printSkipWithIndent(string.concat(roleNamePlural, " are up to date"));
        } else {
            printLog(
                string.concat("Added ", vm.toString(added), " and removed ", vm.toString(removed), " ", roleNamePlural)
            );
        }
    }

    function getMembers(bytes32 role) internal view returns (address[] memory members) {
        uint256 count = fastBridge.getRoleMemberCount(role);
        members = new address[](count);
        for (uint256 i = 0; i < count; i++) {
            members[i] = fastBridge.getRoleMember(role, i);
        }
    }

    function contains(address[] memory array, address value) internal pure returns (bool) {
        for (uint256 i = 0; i < array.length; i++) {
            if (array[i] == value) return true;
        }
        return false;
    }
}
