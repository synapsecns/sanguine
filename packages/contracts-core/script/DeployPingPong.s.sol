// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Script.sol";

import { DeployerUtils } from "./utils/DeployerUtils.sol";

import { PingPongClient } from "../contracts/client/PingPongClient.sol";

contract DeployPingPongScript is DeployerUtils {
    using stdJson for string;

    uint256 internal nonce;
    address internal origin;
    address internal destination;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
        nonce = vm.envUint("PING_PONG_NONCE");
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
            contractName: "PingPongClient",
            saveAsName: "PingPongA",
            deployFunc: _deployPingPong
        });
        deployContract({
            contractName: "PingPongClient",
            saveAsName: "PingPongB",
            deployFunc: _deployPingPong
        });
        stopBroadcast();
    }

    function _deployPingPong() internal returns (address) {
        PingPongClient _pingPong = new PingPongClient({
            _origin: origin,
            _destination: destination
        });
        // TestClient is unowned
        return address(_pingPong);
    }

    function _skipToNonce() internal {
        uint256 _nonce = vm.getNonce(broadcasterAddress);
        for (; _nonce < nonce; _nonce++) {
            console.log("Skipping nonce: %s", _nonce);
            payable(broadcasterAddress).transfer(0);
        }
    }
}
