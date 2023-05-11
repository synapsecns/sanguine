// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {TestClient} from "../../../contracts/client/TestClient.sol";
import {Request, RequestLib} from "../../../contracts/libs/stack/Request.sol";
import {InterfaceOrigin} from "../../../contracts/interfaces/InterfaceOrigin.sol";

import {OriginMock} from "../../mocks/OriginMock.t.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract TestClientTest is Test {
    TestClient internal client;

    address internal originMock;
    address internal destinationMock;

    event MessageReceived(
        uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, uint32 version, bytes content
    );

    event MessageSent(uint32 destination, uint32 nonce, bytes32 sender, bytes32 recipient, bytes content);

    function setUp() public {
        originMock = address(new OriginMock());
        destinationMock = makeAddr("Destination Mock");
        client = new TestClient(originMock, destinationMock);
    }

    function test_sendMessage(
        uint32 destination_,
        uint32 nonce,
        address recipient,
        uint32 optimisticPeriod,
        uint64 gasLimit,
        uint32 version,
        bytes memory content
    ) public {
        vm.assume(recipient != address(0));
        vm.assume(nonce != 0);
        Request request = RequestLib.encodeRequest({gasDrop_: 0, gasLimit_: gasLimit, version_: version});
        bytes32 clientBytes32 = bytes32(uint256(uint160(address(client))));
        bytes32 recipientBytes32 = bytes32(uint256(uint160(recipient)));
        // Mock returned values for sendBaseMessage call
        vm.mockCall(originMock, abi.encodeWithSelector(InterfaceOrigin.sendBaseMessage.selector), abi.encode(nonce, 0));
        vm.expectCall(
            originMock,
            abi.encodeWithSelector(
                InterfaceOrigin.sendBaseMessage.selector, destination_, recipient, optimisticPeriod, request, content
            )
        );
        vm.expectEmit();
        emit MessageSent(destination_, nonce, clientBytes32, recipientBytes32, content);
        client.sendMessage(destination_, recipient, optimisticPeriod, gasLimit, version, content);
    }

    function test_receiveMessage(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) public {
        vm.assume(nonce != 0);
        vm.assume(sender != 0);
        vm.assume(proofMaturity != 0);
        vm.expectEmit();
        emit MessageReceived(origin, nonce, sender, proofMaturity, version, content);
        vm.prank(destinationMock);
        client.receiveBaseMessage(origin, nonce, sender, proofMaturity, version, content);
    }
}
