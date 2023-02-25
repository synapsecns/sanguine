// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../../contracts/client/PingPongClient.sol";

import { Test } from "forge-std/Test.sol";

contract OriginMock is IOrigin {
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external returns (uint32 messageNonce, bytes32 messageHash) {}
}

// solhint-disable func-name-mixedcase
contract PingPongTest is Test {
    PingPongClient internal client;

    address internal originMock;
    address internal destinationMock;

    // Emitted when "Ping" is sent
    event Ping(uint256 indexed pingId, uint16 pongsLeft);

    // Emitted when "Ping" is received
    event Pong(uint256 indexed pingId, uint16 pongsLeft);

    function setUp() public {
        originMock = address(new OriginMock());
        destinationMock = makeAddr("Destination Mock");
        client = new PingPongClient(originMock, destinationMock);
    }

    function test_ping(
        uint32 destination,
        address recipient,
        uint16 pongsTotal
    ) public {
        uint256 pingId = client.totalSent();
        uint32 nextPeriod = client.nextOptimisticPeriod();
        // Should call Origin
        _expectOriginCall(destination, recipient, nextPeriod, pingId, pongsTotal);
        // Should emit Ping
        vm.expectEmit(true, true, true, true);
        emit Ping(pingId, pongsTotal);
        client.doPing(destination, recipient, pongsTotal);
        assertEq(client.totalSent(), 1);
    }

    function test_pings(
        uint8 pingCount,
        uint32 destination,
        address recipient,
        uint16 pongsTotal
    ) public {
        uint256 pingId = client.totalSent();
        uint256 random = client.random();
        uint32[] memory periods = new uint32[](pingCount);
        for (uint256 i = 0; i < pingCount; ++i) {
            periods[i] = uint32(random % 60);
            random = uint256(keccak256(abi.encode(random)));
        }
        for (uint256 i = 0; i < pingCount; ++i) {
            // Should call Origin
            _expectOriginCall(destination, recipient, periods[i], pingId + i, pongsTotal);
            // Should emit Ping
            vm.expectEmit(true, true, true, true);
            emit Ping(pingId + i, pongsTotal);
        }
        client.doPings(pingCount, destination, recipient, pongsTotal);
        assertEq(client.totalSent(), pingCount);
    }

    function test_pong(
        uint32 origin,
        address sender,
        uint256 pingId,
        uint16 pongsLeft
    ) public {
        uint32 nextPeriod = client.nextOptimisticPeriod();
        // Should emit Pong
        vm.expectEmit(true, true, true, true);
        emit Pong(pingId, pongsLeft);
        if (pongsLeft != 0) {
            // Should call Origin if amount of pongs > 0
            _expectOriginCall(origin, sender, nextPeriod, pingId, pongsLeft - 1);
        }
        vm.prank(destinationMock);
        client.handle(
            origin,
            0,
            bytes32(uint256(uint160(sender))),
            0,
            _messageBody(pingId, pongsLeft)
        );
        assertEq(client.totalReceived(), 1);
    }

    function _expectOriginCall(
        uint32 destination,
        address recipient,
        uint32 optimisticPeriod,
        uint256 pingId,
        uint16 pongsLeft
    ) internal {
        bytes memory tips = Tips.emptyTips();
        bytes memory body = _messageBody(pingId, pongsLeft);
        vm.expectCall(
            originMock,
            abi.encodeWithSelector(
                IOrigin.dispatch.selector,
                destination,
                recipient,
                optimisticPeriod,
                tips,
                body
            )
        );
    }

    function _messageBody(uint256 pingId, uint16 pongsLeft) internal pure returns (bytes memory) {
        return abi.encode(pingId, pongsLeft);
    }
}
