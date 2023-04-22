// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { console, stdJson } from "forge-std/Script.sol";

import { DeployerUtils } from "./utils/DeployerUtils.sol";

import { TestClient } from "../contracts/client/TestClient.sol";

contract DeployMessagingMVPScript is DeployerUtils {
    using stdJson for string;

    uint256 internal nonce;
    address internal origin;
    address internal destination;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
        nonce = vm.envUint("TARGET_NONCE");
    }

    /// @notice Main function with the deploy logic.
    function run() external {
        _deploy(true);
    }

    /// @notice Function to simulate the deployment procedure.
    function runDry() external {
        _deploy(false);
    }

    /// @dev Deploys second TestClient contract
    function _deploy(bool _isBroadcasted) internal {
        startBroadcast(_isBroadcasted);
        _skipToNonce();
        origin = loadDeployment("Origin");
        destination = loadDeployment("Destination");
        deployContract({
            contractName: "TestClient",
            saveAsName: "TestClientSecond",
            deployFunc: _deployTestClient
        });
        stopBroadcast();
    }

    function _deployTestClient() internal returns (address) {
        TestClient _testClient = new TestClient({ _origin: origin, _destination: destination });
        // TestClient is unowned
        return address(_testClient);
    }

    function _skipToNonce() internal {
        uint256 _nonce = vm.getNonce(broadcasterAddress);
        for (; _nonce < nonce; _nonce++) {
            console.log("Skipping nonce: %s", _nonce);
            payable(broadcasterAddress).transfer(0);
        }
    }
}
