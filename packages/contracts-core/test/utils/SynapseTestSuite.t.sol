// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./SynapseTestStorage.t.sol";
import "./SynapseUtilities.t.sol";

contract SynapseTestSuite is SynapseUtilities, SynapseTestStorage {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               SIGNING                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function signMessage(uint256 privKey, bytes memory message)
        public
        returns (bytes memory signature)
    {
        bytes32 digest = keccak256(message);
        digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", digest));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, digest);
        signature = abi.encodePacked(r, s, v);
    }

    function signMessage(address signer, bytes memory message)
        public
        returns (bytes memory signature)
    {
        uint256 privKey = privKeys[signer];
        require(privKey != 0, "Unknown account");
        return signMessage(privKey, message);
    }

    function registerPK(uint256 privKey) public returns (address account) {
        account = vm.addr(privKey);
        // Save priv key for later usage
        privKeys[account] = privKey;
    }

    function registerActor(string memory actorName) public returns (address account) {
        account = registerPK(generatePrivateKey(abi.encode(actorName)));
        vm.label(account, actorName);
    }
}
