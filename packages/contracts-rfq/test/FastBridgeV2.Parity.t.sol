// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridgeTest, SenderIncorrect} from "./FastBridge.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2ParityTest is FastBridgeTest {
    address public anotherRelayer = makeAddr("Another Relayer");

    function deployFastBridge() internal virtual override returns (address) {
        // Use the cheatcode to deploy 0.8.24 contract within a 0.8.20 test
        return deployCode({what: "FastBridgeV2", args: abi.encode(owner)});
    }

    /// @notice Relay function is no longer permissioned, so we skip this test
    function test_failedRelayNotRelayer() public virtual override {
        vm.skip(true);
    }

    /// @notice Claim function is no longer permissioned by the role (but still by proven address),
    /// so we skip this test
    function test_failedClaimNotRelayer() public virtual override {
        vm.skip(true);
    }

    /// @notice Claim function is no longer permissioned by the role (but still by proven address),
    /// so we modify the parent test by removing the role assignment.
    function test_failedClaimNotOldRelayer() public virtual override {
        setUpRoles();
        test_successfulBridge();
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);
        vm.warp(block.timestamp + 31 minutes);
        vm.prank(relayer);
        fastBridge.prove(request, bytes32("0x04"));
        // TODO: uncomment
        // vm.expectRevert(abi.encodeWithSelector(SenderIncorrect.selector));
        vm.prank(anotherRelayer);
        fastBridge.claim(request, relayer);
    }
}
