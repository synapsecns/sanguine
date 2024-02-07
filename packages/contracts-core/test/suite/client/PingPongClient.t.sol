// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {PingPongClient} from "../../../contracts/client/PingPongClient.sol";
import {Request, RequestLib} from "../../../contracts/libs/stack/Request.sol";
import {InterfaceOrigin} from "../../../contracts/interfaces/InterfaceOrigin.sol";

import {OriginMock} from "../../mocks/OriginMock.t.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract PingPongTest is Test {
    PingPongClient internal client;

    address internal originMock;
    address internal destinationMock;

    event PingSent(uint256 pingId);

    event PingReceived(uint256 pingId);

    event PongSent(uint256 pingId);

    event PongReceived(uint256 pingId);

    function setUp() public {
        originMock = address(new OriginMock());
        destinationMock = makeAddr("Destination Mock");
        client = new PingPongClient(originMock, destinationMock);
    }

    function test_ping(uint32 destination, address recipient, uint16 counter) public {
        vm.assume(recipient != address(0));
        uint256 pingId = client.pingsSent();
        uint32 nextPeriod = client.nextOptimisticPeriod();
        // Should call Origin
        _expectOriginCall(destination, recipient, nextPeriod, pingId, true, counter);
        // Should emit PingSent
        vm.expectEmit(true, true, true, true);
        emit PingSent(pingId);
        client.doPing(destination, recipient, counter);
        // Pings sent: 1
        assertEq(client.pingsSent(), 1);
        // Received pings: 0, pongs: 0
        assertEq(client.pingsReceived(), 0);
        assertEq(client.pongsReceived(), 0);
    }

    function test_pings(uint8 pingCount, uint32 destination, address recipient, uint16 counter) public {
        vm.assume(recipient != address(0));
        uint256 pingId = client.pingsSent();
        uint256 random = client.random();
        uint32[] memory periods = new uint32[](pingCount);
        for (uint256 i = 0; i < pingCount; ++i) {
            periods[i] = uint32(random % 60);
            random = uint256(keccak256(abi.encode(random)));
        }
        for (uint256 i = 0; i < pingCount; ++i) {
            // Should call Origin
            _expectOriginCall(destination, recipient, periods[i], pingId + i, true, counter);
            // Should emit PingSent
            vm.expectEmit(true, true, true, true);
            emit PingSent(pingId + i);
        }
        client.doPings(pingCount, destination, recipient, counter);
        // Pings sent: pingCount
        assertEq(client.pingsSent(), pingCount);
        // Received pings: 0, pongs: 0
        assertEq(client.pingsReceived(), 0);
        assertEq(client.pongsReceived(), 0);
    }

    function test_receivePing(uint32 origin, address sender, uint256 pingId, uint16 counter) public {
        vm.assume(sender != address(0));
        uint32 nextPeriod = client.nextOptimisticPeriod();
        // Should emit PingReceived
        vm.expectEmit(true, true, true, true);
        emit PingReceived(pingId);
        // Should emit PongSent
        emit PongSent(pingId);
        _expectOriginCall(origin, sender, nextPeriod, pingId, false, counter);
        vm.prank(destinationMock);
        client.receiveBaseMessage(origin, 1, bytes32(uint256(uint160(sender))), 1, 0, _content(pingId, true, counter));
        // Pings sent: 0
        assertEq(client.pingsSent(), 0);
        // Received pings: 1, pongs: 0
        assertEq(client.pingsReceived(), 1);
        assertEq(client.pongsReceived(), 0);
    }

    function test_receivePong(uint32 origin, address sender, uint256 pingId, uint16 counter) public {
        vm.assume(sender != address(0));
        uint256 localPingId = client.pingsSent();
        uint32 nextPeriod = client.nextOptimisticPeriod();
        // Should emit PongReceived
        vm.expectEmit(true, true, true, true);
        emit PongReceived(pingId);
        if (counter > 0) {
            // Should emit PingSent
            vm.expectEmit(true, true, true, true);
            emit PingSent(localPingId);
            _expectOriginCall(origin, sender, nextPeriod, localPingId, true, counter - 1);
        }
        vm.prank(destinationMock);
        client.receiveBaseMessage(origin, 1, bytes32(uint256(uint160(sender))), 1, 0, _content(pingId, false, counter));
        // Pings sent: 0/1 (based on counter being zero / non-zero)
        assertEq(client.pingsSent(), counter == 0 ? 0 : 1);
        // Received pings: 0, pongs: 1
        assertEq(client.pingsReceived(), 0);
        assertEq(client.pongsReceived(), 1);
    }

    function _expectOriginCall(
        uint32 destination,
        address recipient,
        uint32 optimisticPeriod,
        uint256 pingId,
        bool isPing,
        uint16 counter
    ) internal {
        Request request = RequestLib.encodeRequest({gasDrop_: 0, gasLimit_: 500_000, version_: 0});
        bytes memory content = _content(pingId, isPing, counter);
        vm.expectCall(
            originMock,
            abi.encodeWithSelector(
                InterfaceOrigin.sendBaseMessage.selector, destination, recipient, optimisticPeriod, request, content
            )
        );
    }

    function _content(uint256 pingId, bool isPing, uint16 counter) internal pure returns (bytes memory) {
        return abi.encode(pingId, isPing, counter);
    }
}
