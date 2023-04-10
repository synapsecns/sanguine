// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseTestConstants} from "./SynapseTestConstants.t.sol";

import {Test} from "forge-std/Test.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable no-empty-blocks
contract SynapseUtilities is SynapseTestConstants, Test {
    bytes internal constant REVERT_ALREADY_INITIALIZED = "Initializable: contract is already initialized";
    bytes internal constant REVERT_NOT_OWNER = "Ownable: caller is not the owner";

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseUtilities() external {}

    function expectRevertAlreadyInitialized() public {
        vm.expectRevert(REVERT_ALREADY_INITIALIZED);
    }

    function expectRevertNotOwner() public {
        vm.expectRevert(REVERT_NOT_OWNER);
    }

    function skipBlock() public {
        skipBlocks(1);
    }

    function skipBlocks(uint256 blocks) public {
        vm.roll(block.number + blocks);
        skip(blocks * BLOCK_TIME);
    }

    function addressToBytes32(address addr) public pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    function bytes32ToAddress(bytes32 buf) public pure returns (address) {
        return address(uint160(uint256(buf)));
    }

    function signMessage(uint256 privKey, bytes memory message) public pure returns (bytes memory signature) {
        return signMessage(privKey, keccak256(message));
    }

    function signMessage(uint256 privKey, bytes32 hashedMsg) public pure returns (bytes memory signature) {
        bytes32 digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hashedMsg));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, digest);
        signature = abi.encodePacked(r, s, v);
    }
}
