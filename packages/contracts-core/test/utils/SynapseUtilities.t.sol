// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract SynapseUtilities is Test {
    bytes internal constant REVERT_ALREADY_INITIALIZED =
        "Initializable: contract is already initialized";
    bytes internal constant REVERT_NOT_OWNER = "Ownable: caller is not the owner";

    function expectRevertAlreadyInitialized() public {
        vm.expectRevert(REVERT_ALREADY_INITIALIZED);
    }

    function expectRevertNotOwner() public {
        vm.expectRevert(REVERT_NOT_OWNER);
    }

    function addressToBytes32(address addr) public pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    function bytes32ToAddress(bytes32 buf) public pure returns (address) {
        return address(uint160(uint256(buf)));
    }

    function generateAddress(bytes memory salt) public pure returns (address) {
        return bytes32ToAddress(keccak256(salt));
    }

    function generatePrivateKey(bytes memory salt) public pure returns (uint256) {
        return uint256(keccak256(salt));
    }

    function getActorSuffix(uint256 actorIndex) public pure returns (string memory) {
        return actorIndex == 0 ? "" : string.concat("[", Strings.toString(actorIndex), "]");
    }
}
