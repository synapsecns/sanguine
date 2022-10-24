// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import { MiniMerkleLib } from "../../contracts/libs/MiniMerkle.sol";

contract MiniMerkleTest is Test {
    using MiniMerkleLib for MiniMerkleLib.Tree;
    // libZeroHashes contains zeroHashes produced by merkle lib
    bytes32[] internal libZeroHashes;
    MiniMerkleLib.Tree internal tree;

    // keccak256 zero hashes
    bytes32 internal constant Z_0 =
        hex"0000000000000000000000000000000000000000000000000000000000000000";
    bytes32 internal constant Z_1 =
        hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    bytes32 internal constant Z_2 =
        hex"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30";
    bytes32 internal constant Z_3 =
        hex"21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85";

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

        for (uint256 height = 0; height < MiniMerkleLib.TREE_DEPTH; height++) {
            assertEq(zero_hashes[height], libZeroHashes[height]);
        }
    }

    function test_successfulAddMaxMessages() public {
        bytes32 node = keccak256(abi.encodePacked("hello"));
        for (uint256 i = 0; i < MiniMerkleLib.MAX_LEAVES; i++) {
            tree.insert(i + 1, node);
        }
    }

    function test_branchRootManualProof() public {
        // Insert hello as first leaf
        bytes32 hello = keccak256(abi.encodePacked("hello"));
        tree.insert(1, hello);
        // Check current root
        bytes32 rootCtxHello = tree.rootWithCtx(1, MiniMerkleLib.zeroHashes());
        // Check branch root based on manually created proof
        bytes32 branchRootHello = MiniMerkleLib.branchRoot(hello, MiniMerkleLib.zeroHashes(), 0);
        assertEq(rootCtxHello, branchRootHello, "!rootHello");

        // Insert bye as second leaf
        bytes32 bye = keccak256(abi.encodePacked("bye"));
        tree.insert(2, bye);
        // Check current root
        bytes32 rootCtxBye = tree.rootWithCtx(2, MiniMerkleLib.zeroHashes());
        // Check branch root based on manually created proof
        bytes32 branchRootBye = MiniMerkleLib.branchRoot(bye, [hello, Z_1, Z_2, Z_3], 1);
        assertEq(rootCtxBye, branchRootBye, "!rootBye");

        // Insert world as third leaf
        bytes32 world = keccak256(abi.encodePacked("world"));
        tree.insert(3, world);
        // Check current root
        bytes32 rootCtxWorld = tree.rootWithCtx(3, MiniMerkleLib.zeroHashes());
        // Check branch root based on manually created proof
        bytes32 branchRootWorld = MiniMerkleLib.branchRoot(
            // node
            world,
            // proof
            [Z_0, keccak256(abi.encodePacked(hello, bye)), Z_2, Z_3],
            // index
            2
        );
        assertEq(rootCtxWorld, branchRootWorld, "!rootWorld");

        // new manual proof for hello, which is the first message inserted
        bytes32 branchRootHelloProofTwo = MiniMerkleLib.branchRoot(
            hello,
            [
                bye,
                keccak256(abi.encodePacked(world, Z_0)),
                bytes32(hex"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30"),
                bytes32(hex"21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85")
            ],
            0
        );
        // checks branch root for hello against latest message insertion rootWithCtx
        assertEq(branchRootHelloProofTwo, rootCtxWorld, "!rootHelloVsWorld");
    }

    function test_insertOverMax() public {
        bytes32 node = keccak256(abi.encodePacked("0"));
        vm.expectRevert("merkle tree full");
        tree.insert(MiniMerkleLib.MAX_LEAVES + 1, node);
    }
}
