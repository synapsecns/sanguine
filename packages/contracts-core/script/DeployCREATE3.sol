// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {console, stdJson} from "forge-std/Script.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DeployerUtils} from "./utils/DeployerUtils.sol";

import {CREATE3Factory} from "../contracts/create3/CREATE3Factory.sol";


// TODO: remove this script, I don't think we need it since we handle the devnet create3 deploy in setupDevnetIfEnabled();
contract DeployCREATE3 is DeployerUtils {
    using stdJson for string;
    using Strings for uint256;
    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
        setupDevnetIfEnabled();
    }

    function run() external {
        startBroadcast(true);
        CREATE3Factory NewFactory = new CREATE3Factory();
        saveDeployment("CREATE3", "CREATE3", address(NewFactory), "0x");
        stopBroadcast();
    }
}
