// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { SynapseClientUpgradeableHarness } from "./harnesses/SynapseClientUpgradeableHarness.sol";
import { OriginHarness } from "./harnesses/OriginHarness.sol";

import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";

import {
    TransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract SynapseClientTest is SynapseTestWithNotaryManager {
    SynapseClientUpgradeableHarness public client;
    OriginHarness public origin;

    address public constant destination = address(1234567890);
    address public constant owner = address(9876543210);
    bytes32 public constant trustedSender = bytes32(uint256(1234554321));

    function setUp() public override {
        super.setUp();

        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));

        vm.label(destination, "mirror");
        vm.label(owner, "owner");

        SynapseClientUpgradeableHarness impl = new SynapseClientUpgradeableHarness(
            address(origin),
            destination
        );

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(impl),
            address(1337),
            bytes("")
        );
        client = SynapseClientUpgradeableHarness(address(proxy));
        client.initialize();
        client.transferOwnership(owner);
    }

    function test_constructor() public {
        assertEq(client.origin(), address(origin));
        assertEq(client.destination(), destination);
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        client.initialize();
    }

    function test_setTrustedSender() public {
        vm.prank(owner);
        client.setTrustedSender(remoteDomain, trustedSender);
        assertEq(client.trustedSender(remoteDomain), trustedSender);
    }

    function test_setTrustedSenderAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != owner);
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(_notOwner);
        client.setTrustedSender(remoteDomain, trustedSender);
    }

    function test_setTrustedSenderEmptyDomain() public {
        vm.prank(owner);
        vm.expectRevert("!domain");
        client.setTrustedSender(0, trustedSender);
    }

    function test_setTrustedSenderEmptySender() public {
        vm.prank(owner);
        vm.expectRevert("!sender");
        client.setTrustedSender(remoteDomain, bytes32(0));
    }

    function test_setTrustedSenders() public {
        uint256 amount = 5;
        uint32[] memory domains = new uint32[](amount);
        bytes32[] memory senders = new bytes32[](amount);
        for (uint256 i = 0; i < amount; i++) {
            domains[i] = uint32(remoteDomain + i);
            senders[i] = bytes32(uint256(trustedSender) + i);
        }
        vm.prank(owner);
        client.setTrustedSenders(domains, senders);
        for (uint256 i = 0; i < amount; i++) {
            assertEq(client.trustedSender(domains[i]), senders[i]);
        }
    }

    function test_setTrustedSendersAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != owner);
        uint32[] memory domains = new uint32[](1);
        bytes32[] memory senders = new bytes32[](1);
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(_notOwner);
        client.setTrustedSenders(domains, senders);
    }

    function test_setTrustedSendersBadArrays() public {
        uint32[] memory domains = new uint32[](1);
        bytes32[] memory senders = new bytes32[](2);
        vm.expectRevert("!arrays");
        vm.prank(owner);
        client.setTrustedSenders(domains, senders);
    }

    function test_handle() public {
        test_setTrustedSender();

        vm.prank(destination);
        client.handle(remoteDomain, 0, trustedSender, block.timestamp, bytes(""));
    }

    function test_handleNotDestination(address _notDestination) public {
        vm.assume(_notDestination != destination);
        test_setTrustedSender();

        vm.prank(_notDestination);
        vm.expectRevert("BasicClient: !destination");
        client.handle(remoteDomain, 0, trustedSender, block.timestamp, bytes(""));
    }

    function test_handleFakeDomain(uint32 _notRemote) public {
        vm.assume(_notRemote != remoteDomain);

        test_setTrustedSender();

        vm.prank(destination);
        vm.expectRevert("BasicClient: !trustedSender");
        client.handle(_notRemote, 0, trustedSender, block.timestamp, bytes(""));
    }

    function test_handleFakeSender(bytes32 _notSender) public {
        vm.assume(_notSender != trustedSender);

        test_setTrustedSender();

        vm.prank(destination);
        vm.expectRevert("BasicClient: !trustedSender");
        client.handle(remoteDomain, 0, _notSender, block.timestamp, bytes(""));
    }

    function test_handleFakeDomainAndSender(uint32 _notRemote) public {
        vm.assume(_notRemote != remoteDomain);

        test_setTrustedSender();

        vm.prank(destination);
        vm.expectRevert("BasicClient: !trustedSender");
        // trustedSender for unknown remote is bytes32(0),
        // but this still has to revert
        client.handle(_notRemote, 0, bytes32(0), block.timestamp, bytes(""));
    }

    function test_handleOptimisticSecondsNotPassed() public {
        test_setTrustedSender();

        vm.prank(destination);
        vm.expectRevert("Client: !optimisticSeconds");
        client.handle(remoteDomain, 0, trustedSender, block.timestamp + 1, bytes(""));
    }

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes tips,
        bytes message
    );

    function test_send() public {
        test_setTrustedSender();
        bytes memory messageBody = hex"01030307";
        bytes memory _header = Header.formatHeader(
            localDomain,
            bytes32(uint256(uint160(address(client)))),
            1,
            remoteDomain,
            trustedSender,
            0
        );
        bytes memory _tips = getDefaultTips();
        bytes memory message = Message.formatMessage(_header, _tips, messageBody);
        vm.expectEmit(true, true, true, true);
        emit Dispatch(keccak256(message), 0, (uint64(remoteDomain) << 32) | 1, _tips, message);
        deal(address(this), TOTAL_TIPS);
        client.send{ value: TOTAL_TIPS }(remoteDomain, _tips, messageBody);
    }

    function test_sendNoRecipient() public {
        bytes memory messageBody = hex"01030307";
        vm.expectRevert("BasicClient: !recipient");
        client.send(remoteDomain, getEmptyTips(), messageBody);
    }
}
