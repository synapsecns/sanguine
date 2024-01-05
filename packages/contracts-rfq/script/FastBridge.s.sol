// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridge} from "../contracts/FastBridge.sol";
import {Script} from "forge-std/Script.sol";

contract DeployFastBridge is Script {
    FastBridge public bridge;

    /// e.g. forge script contracts/script/FastBridge.s.sol --sig "run(address, address[])" 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720 "[0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f]"
    function run(address owner, address[] memory relayers) external {
        vm.startBroadcast();

        // deploy bridge making script sender the owner
        bridge = new FastBridge(msg.sender);

        // add relayers
        for (uint256 i = 0; i < relayers.length; i++) {
            address relayer = relayers[i];
            bridge.addRelayer(relayer);
        }

        // set new default admin as owner then renounce if owner != msg.sender
        if (msg.sender != owner) {
            bytes32 adminRole = bridge.DEFAULT_ADMIN_ROLE();
            bridge.grantRole(adminRole, owner);
            bridge.revokeRole(adminRole, msg.sender);
        }

        vm.stopBroadcast();
    }
}
