// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { Message } from "../contracts/libs/Message.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Tips } from "../contracts/libs/Tips.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";
import { SystemRouterHarness } from "./harnesses/SystemRouterHarness.sol";
import { ClientHarness } from "./harnesses/ClientHarness.sol";

// solhint-disable func-name-mixedcase
contract ClientTest is SynapseTestWithNotaryManager {
    SystemRouterHarness internal systemRouter;
    OriginHarness internal origin;
    DestinationHarness internal destination;
    ClientHarness internal client;

    bytes32 internal remoteAddress = keccak256(abi.encode(remoteDomain));
    uint32 internal optimisticSeconds = 420;
    uint32 internal nonce = 1;
    uint256 internal tipsValue = 0;
    bytes internal tips;
    bytes internal messageBody = "such body much wow";

    uint32 internal constant NONCE = 420;
    bytes32 internal constant ROOT = "root";

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes tips,
        bytes message
    );

    event LogMessage(uint32 origin, uint32 nonce, bytes message);

    function setUp() public override {
        super.setUp();

        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));

        destination = new DestinationHarness(localDomain);
        destination.initialize(remoteDomain, notary);

        client = new ClientHarness(address(origin), address(destination), optimisticSeconds);
        tips = Tips.emptyTips();
    }

    function test_constructor() public {
        assertEq(client.origin(), address(origin), "!origin");
        assertEq(client.destination(), address(destination), "!destination");
        assertEq(client.optimisticSeconds(), optimisticSeconds, "!optimisticSeconds");
        assertEq(client.trustedSender(remoteDomain), remoteAddress, "!remoteAddress");
        assertEq(client.trustedSender(0), bytes32(0), "!zero");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        TEST: SENDING MESSAGE                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_send_noTips() public {
        _checkSendMessage();
    }

    function test_send_tips() public {
        tips = getDefaultTips();
        tipsValue = TOTAL_TIPS;
        _checkSendMessage();
    }

    function test_send_noRecipient() public {
        vm.expectRevert("BasicClient: !recipient");
        client.sendMessage(0, tips, messageBody);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TEST: RECEIVING MESSAGE                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_handle() public {
        bytes memory message = _prepareReceiveTest(true, remoteDomain);
        skip(optimisticSeconds);
        destination.execute(message);
    }

    function test_handle_optimisticPeriodNotPassed() public {
        bytes memory message = _prepareReceiveTest(false, remoteDomain);
        skip(optimisticSeconds - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(message);
    }

    function test_handle_optimisticPeriodForged() public {
        optimisticSeconds = 1;
        bytes memory message = _prepareReceiveTest(false, remoteDomain);
        skip(optimisticSeconds);
        vm.expectRevert("Client: !optimisticSeconds");
        destination.execute(message);
    }

    function test_handle_unknownSender() public {
        remoteAddress = "some other guy";
        bytes memory message = _prepareReceiveTest(false, remoteDomain);
        skip(optimisticSeconds);
        vm.expectRevert("BasicClient: !trustedSender");
        destination.execute(message);
    }

    function test_handle_notDestination() public {
        vm.expectRevert("BasicClient: !destination");
        client.handle(0, 0, bytes32(0), 0, "");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _checkSendMessage() internal {
        bytes memory message = _createSentMessage();
        vm.expectEmit(true, true, true, true);
        emit Dispatch(
            keccak256(message),
            nonce - 1,
            (uint64(remoteDomain) << 32) | nonce,
            tips,
            message
        );
        deal(address(this), tipsValue);
        client.sendMessage{ value: tipsValue }(remoteDomain, tips, messageBody);
    }

    function _prepareReceiveTest(bool _success, uint32 _originDomain)
        internal
        returns (bytes memory message)
    {
        message = _createReceivedMessage();
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, ROOT);
        destination.submitAttestation(attestation);
        destination.setMessageStatus(_originDomain, keccak256(message), ROOT);
        if (_success) {
            vm.expectEmit(true, true, true, true);
            emit LogMessage(_originDomain, nonce, messageBody);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _createReceivedMessage() internal view returns (bytes memory) {
        return
            Message.formatMessage(
                Header.formatHeader(
                    remoteDomain,
                    remoteAddress,
                    nonce,
                    localDomain,
                    addressToBytes32(address(client)),
                    optimisticSeconds
                ),
                tips,
                messageBody
            );
    }

    function _createSentMessage() internal view returns (bytes memory) {
        return
            Message.formatMessage(
                Header.formatHeader(
                    localDomain,
                    addressToBytes32(address(client)),
                    nonce,
                    remoteDomain,
                    remoteAddress,
                    optimisticSeconds
                ),
                tips,
                messageBody
            );
    }
}
