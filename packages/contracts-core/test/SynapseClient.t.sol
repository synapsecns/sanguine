// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { SynapseClientHarness } from "./harnesses/SynapseClientHarness.sol";
import { OriginHarness } from "./harnesses/OriginHarness.sol";

import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";

// solhint-disable func-name-mixedcase
contract SynapseClientTest is SynapseTestWithNotaryManager {
    SynapseClientHarness public client;
    OriginHarness public origin;

    address public constant DESTINATION = address(1234567890);
    address public constant OWNER = address(9876543210);
    bytes32 public constant TRUSTED_SENDER = bytes32(uint256(1234554321));

    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    function setUp() public override {
        super.setUp();

        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));

        vm.label(DESTINATION, "destination");
        vm.label(OWNER, "owner");

        client = new SynapseClientHarness(address(origin), DESTINATION);
        client.transferOwnership(OWNER);
    }

    function test_constructor() public {
        assertEq(client.origin(), address(origin));
        assertEq(client.destination(), DESTINATION);
    }

    function test_setTrustedSender() public {
        vm.prank(OWNER);
        client.setTrustedSender(remoteDomain, TRUSTED_SENDER);
        assertEq(client.trustedSender(remoteDomain), TRUSTED_SENDER);
    }

    function test_setTrustedSenderAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != OWNER);
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(_notOwner);
        client.setTrustedSender(remoteDomain, TRUSTED_SENDER);
    }

    function test_setTrustedSenderEmptyDomain() public {
        vm.prank(OWNER);
        vm.expectRevert("!domain");
        client.setTrustedSender(0, TRUSTED_SENDER);
    }

    function test_setTrustedSenderEmptySender() public {
        vm.prank(OWNER);
        vm.expectRevert("!sender");
        client.setTrustedSender(remoteDomain, bytes32(0));
    }

    function test_setTrustedSenders() public {
        uint256 amount = 5;
        uint32[] memory domains = new uint32[](amount);
        bytes32[] memory senders = new bytes32[](amount);
        for (uint256 i = 0; i < amount; i++) {
            domains[i] = uint32(remoteDomain + i);
            senders[i] = bytes32(uint256(TRUSTED_SENDER) + i);
        }
        vm.prank(OWNER);
        client.setTrustedSenders(domains, senders);
        for (uint256 i = 0; i < amount; i++) {
            assertEq(client.trustedSender(domains[i]), senders[i]);
        }
    }

    function test_setTrustedSendersAsNotOwner(address _notOwner) public {
        vm.assume(_notOwner != OWNER);
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
        vm.prank(OWNER);
        client.setTrustedSenders(domains, senders);
    }

    function test_handle() public {
        test_setTrustedSender();
        vm.prank(DESTINATION);
        client.handle(remoteDomain, 0, TRUSTED_SENDER, block.timestamp, bytes(""));
    }

    function test_handleNotDestination(address _notDestination) public {
        vm.assume(_notDestination != DESTINATION);
        test_setTrustedSender();
        vm.prank(_notDestination);
        vm.expectRevert("BasicClient: !destination");
        client.handle(remoteDomain, 0, TRUSTED_SENDER, block.timestamp, bytes(""));
    }

    function test_handleFakeDomain(uint32 _notRemote) public {
        vm.assume(_notRemote != remoteDomain);
        test_setTrustedSender();
        vm.prank(DESTINATION);
        vm.expectRevert("BasicClient: !trustedSender");
        client.handle(_notRemote, 0, TRUSTED_SENDER, block.timestamp, bytes(""));
    }

    function test_handleFakeSender(bytes32 _notSender) public {
        vm.assume(_notSender != TRUSTED_SENDER);
        test_setTrustedSender();
        vm.prank(DESTINATION);
        vm.expectRevert("BasicClient: !trustedSender");
        client.handle(remoteDomain, 0, _notSender, block.timestamp, bytes(""));
    }

    function test_handleFakeDomainAndSender(uint32 _notRemote) public {
        vm.assume(_notRemote != remoteDomain);
        test_setTrustedSender();
        vm.prank(DESTINATION);
        vm.expectRevert("BasicClient: !trustedSender");
        // TRUSTED_SENDER for unknown remote is bytes32(0),
        // but this still has to revert
        client.handle(_notRemote, 0, bytes32(0), block.timestamp, bytes(""));
    }

    function test_handleOptimisticSecondsNotPassed() public {
        test_setTrustedSender();
        vm.prank(DESTINATION);
        vm.expectRevert("Client: !optimisticSeconds");
        client.handle(remoteDomain, 0, TRUSTED_SENDER, block.timestamp + 1, bytes(""));
    }

    function test_send() public {
        test_setTrustedSender();
        bytes memory messageBody = hex"01030307";
        bytes memory _header = Header.formatHeader(
            localDomain,
            bytes32(uint256(uint160(address(client)))),
            1,
            remoteDomain,
            TRUSTED_SENDER,
            0
        );
        bytes memory _tips = getDefaultTips();
        bytes memory message = Message.formatMessage(_header, _tips, messageBody);
        vm.expectEmit(true, true, true, true);
        emit Dispatch(keccak256(message), 1, remoteDomain, _tips, message);
        deal(address(this), TOTAL_TIPS);
        client.sendMessage{ value: TOTAL_TIPS }(remoteDomain, _tips, messageBody);
    }

    function test_sendNoRecipient() public {
        bytes memory messageBody = hex"01030307";
        vm.expectRevert("BasicClient: !recipient");
        client.sendMessage(remoteDomain, getEmptyTips(), messageBody);
    }
}
