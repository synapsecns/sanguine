// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";
import { MerkleLib } from "../../contracts/libs/Merkle.sol";

contract MerkleTest is Test {
    // libZeroHashes contains zeroHashes produced by merkle lib
    bytes32[] internal libZeroHashes;
    MerkleLib.Tree internal tree;

    constructor() {
        libZeroHashes = MerkleLib.zeroHashes();
    }

    function test_zeroHashes() public {
        // Compute hashes in empty sparse Merkle tree
        bytes32[32] memory zero_hashes;

        for (uint256 height = 0; height < MerkleLib.TREE_DEPTH - 1; height++) {
            zero_hashes[height + 1] = keccak256(
                abi.encodePacked(zero_hashes[height], zero_hashes[height])
            );
        }

        for (uint256 height = 0; height < MerkleLib.TREE_DEPTH; height++) {
            assertEq(zero_hashes[height], libZeroHashes[height]);
        }
    }

    function test_insertOverMax() public {
        bytes32 node = keccak256(abi.encodePacked("0"));
        vm.expectRevert("merkle tree full");
        MerkleLib.insert(tree, MerkleLib.MAX_LEAVES + 1, node);
    }
}
