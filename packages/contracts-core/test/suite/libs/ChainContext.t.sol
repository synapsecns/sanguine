// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ChainContextHarness} from "../../harnesses/libs/ChainContextHarness.t.sol";

import {SynapseLibraryTest} from "../../utils/SynapseLibraryTest.t.sol";

// solhint-disable func-name-mixedcase
contract ChainContextLibraryTest is SynapseLibraryTest {
    ChainContextHarness public libHarness;

    function setUp() public {
        libHarness = new ChainContextHarness();
    }

    function test_blockNumber() public {
        uint256 expectedBN = 1337;
        vm.roll(expectedBN);
        assertEq(libHarness.blockNumber(), expectedBN);
    }

    function test_blockNumber_maxValue() public {
        uint256 expectedBN = 2 ** 40 - 1;
        vm.roll(expectedBN);
        assertEq(libHarness.blockNumber(), expectedBN);
    }

    function test_blockNumber_revert_blockNumberOverflow() public {
        uint256 overflowBN = 2 ** 40;
        vm.roll(overflowBN);
        vm.expectRevert("SafeCast: value doesn't fit in 40 bits");
        libHarness.blockNumber();
    }

    function test_blockTimestamp() public {
        uint256 expectedBT = 7331;
        vm.warp(expectedBT);
        assertEq(libHarness.blockTimestamp(), expectedBT);
    }

    function test_blockTimestamp_maxValue() public {
        uint256 expectedBT = 2 ** 40 - 1;
        vm.warp(expectedBT);
        assertEq(libHarness.blockTimestamp(), expectedBT);
    }

    function test_blockTimestamp_revert_blockTimestampOverflow() public {
        uint256 overflowBT = 2 ** 40;
        vm.warp(overflowBT);
        vm.expectRevert("SafeCast: value doesn't fit in 40 bits");
        libHarness.blockTimestamp();
    }

    function test_chainId() public {
        uint256 expectedChainId = 420;
        vm.chainId(expectedChainId);
        assertEq(libHarness.chainId(), expectedChainId);
    }

    function test_chainId_maxValue() public {
        uint256 expectedChainId = 2 ** 32 - 1;
        vm.chainId(expectedChainId);
        assertEq(libHarness.chainId(), expectedChainId);
    }

    function test_chainId_revert_chainIdOverflow() public {
        uint256 overflowChainId = 2 ** 32;
        vm.chainId(overflowChainId);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        libHarness.chainId();
    }
}
