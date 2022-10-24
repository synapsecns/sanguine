// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";
import { MiniMerkleLib } from "../../contracts/libs/MiniMerkle.sol";

contract MiniMerkleTest is Test {
    // libZeroHashes contains zeroHashes produced by merkle lib
    bytes32[] internal libZeroHashes;
    MiniMerkleLib.Tree internal tree;

    constructor() {
        libZeroHashes = MiniMerkleLib.zeroHashes();
    }

    function test_zeroHashes() public {
        // Compute hashes in empty sparse Merkle tree
        bytes32[4] memory zero_hashes;

        for (uint256 height = 0; height < MiniMerkleLib.TREE_DEPTH - 1; height++) {
            zero_hashes[height + 1] = keccak256(
                abi.encodePacked(zero_hashes[height], zero_hashes[height])
            );
        }

        for (uint256 height = 0; height < MinilibZeroHashes[height]);
        }
    }

    function test_insertOverMax() public {
        bytes32 node = keccak256(abi.encodePacked("0"));
        vm.expectRevert("merkle tree full");
        MiniMerkleLib.insert(tree, MiniMerkleLib.MAX_LEAVES + 1, node);
    }
}
