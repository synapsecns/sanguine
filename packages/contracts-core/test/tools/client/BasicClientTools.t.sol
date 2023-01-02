// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../DestinationTools.t.sol";
import { BasicClient } from "../../../contracts/client/BasicClient.sol";
import { BasicClientHarness } from "../../harnesses/client/BasicClientHarness.t.sol";

abstract contract BasicClientTools is DestinationTools {
    mapping(uint32 => BasicClient) internal basicClients;
    MessageContext internal clientLocalToRemote;
    MessageContext internal mockSenderRemoteToLocal;

    function setupBasicClients() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            basicClients[domain] = deployBasicClient(domain, APP_OPTIMISTIC_SECONDS);
        }
        // Context for sending tests
        clientLocalToRemote = MessageContext({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            sender: address(suiteBasicClient(DOMAIN_LOCAL))
        });
        // Context for receiving tests
        mockSenderRemoteToLocal = MessageContext({
            origin: DOMAIN_REMOTE,
            destination: DOMAIN_LOCAL,
            sender: bytes32ToAddress(mockTrustedSender(DOMAIN_REMOTE))
        });
    }

    function deployBasicClient(uint32 domain, uint32 optimisticSeconds)
        public
        virtual
        returns (BasicClient)
    {
        return
            new BasicClientHarness(
                address(suiteOrigin(domain)),
                address(suiteDestination(domain)),
                optimisticSeconds
            );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Creates a test message sent by BasicClient to another chain
    function createClientSentMessage(MessageContext memory context, bool mockTips) public {
        skipBlock();
        createDispatchedMessage({
            context: context,
            mockTips: mockTips,
            body: MOCK_BODY,
            recipient: mockTrustedSender(context.destination),
            optimisticSeconds: APP_OPTIMISTIC_SECONDS
        });
    }

    // Creates test messages sent from another chain for BasisClient to receive.
    // Submits attestation on destination chain to prepare for the messages execution.
    function createMessagesSubmitAttestation(MessageContext memory context) public {
        createMessages({
            context: context,
            recipient: address(suiteBasicClient(context.destination))
        });
        destinationSubmitAttestationSuggested({ context: context, returnValue: true });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectLogBasicClientMessage() public {
        emit LogBasicClientMessage(messageOrigin, messageNonce, rootSubmittedAt, messageBody);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function basicClientSendMessage(bytes memory revertMessage) public {
        vm.expectRevert(revertMessage);
        suiteBasicClient(messageOrigin).sendMessage{ value: tipsTotal }(
            messageDestination,
            tipsRaw,
            messageBody
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function basicClientSendMessage() public {
        suiteBasicClient(messageOrigin).sendMessage{ value: tipsTotal }(
            messageDestination,
            tipsRaw,
            messageBody
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            CHAIN GETTERS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function suiteBasicClient(uint32 domain) public view returns (BasicClientHarness) {
        return BasicClientHarness(address(basicClients[domain]));
    }

    function mockTrustedSender(uint32 domain) public pure returns (bytes32) {
        return bytes32(uint256(domain));
    }
}
