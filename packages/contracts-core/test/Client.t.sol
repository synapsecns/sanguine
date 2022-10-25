// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { Message } from "../contracts/libs/Message.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Tips } from "../contracts/libs/Tips.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";
import { SystemRouterHarness } from "./harnesses/SystemRouterHarness.sol";
import { ClientHarness } from "./harnesses/ClientHarness.sol";
import { ProofGenerator } from "./utils/ProofGenerator.sol";

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

    bytes32 internal constant ROOT = "root";

    ProofGenerator internal proofGen;
    uint32 internal constant NONCE = 1;
    uint32 internal constant INDEX = 0;
    bytes32[32] internal proof;
    bytes32 internal root;

    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
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
        destination.initialize();
        destination.addNotary(remoteDomain, notary);

        client = new ClientHarness(address(origin), address(destination), optimisticSeconds);
        tips = Tips.emptyTips();

        proofGen = new ProofGenerator();
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
        destination.execute(message, proof, INDEX);
    }

    function test_handle_optimisticPeriodNotPassed() public {
        bytes memory message = _prepareReceiveTest(false, remoteDomain);
        skip(optimisticSeconds - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(message, proof, INDEX);
    }

    function test_handle_optimisticPeriodForged() public {
        optimisticSeconds = 1;
        bytes memory message = _prepareReceiveTest(false, remoteDomain);
        skip(optimisticSeconds);
        vm.expectRevert("Client: !optimisticSeconds");
        destination.execute(message, proof, INDEX);
    }

    function test_handle_unknownSender() public {
        remoteAddress = "some other guy";
        bytes memory message = _prepareReceiveTest(false, remoteDomain);
        skip(optimisticSeconds);
        vm.expectRevert("BasicClient: !trustedSender");
        destination.execute(message, proof, INDEX);
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
        emit Dispatch(keccak256(message), nonce, remoteDomain, tips, message);
        deal(address(this), tipsValue);
        client.sendMessage{ value: tipsValue }(remoteDomain, tips, messageBody);
    }

    function _prepareReceiveTest(bool _success, uint32 _originDomain)
        internal
        returns (bytes memory message)
    {
        message = _createReceivedMessage();
        // Create a merkle tree with a lonely message
        bytes32[] memory leafs = new bytes32[](1);
        leafs[0] = keccak256(message);
        proofGen.createTree(leafs);
        root = proofGen.getRoot();
        proof = proofGen.getProof(INDEX);
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, root);
        destination.submitAttestation(attestation);
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
