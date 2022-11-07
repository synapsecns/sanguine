// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/client/ClientTools.t.sol";

// solhint-disable func-name-mixedcase
contract ClientTest is ClientTools {
    function setUp() public virtual override {
        super.setUp();
        setupBasicClients();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setup_constructor() public {
        // Check Client storage
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            ClientHarness client = suiteClient(domain);
            assertEq(client.destination(), address(suiteDestination(domain)), "!destination");
            assertEq(client.origin(), address(suiteOrigin(domain)), "!origin");
        }
    }

    function test_setup_harnessMocks() public {
        // Check ClientHarness mocks
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            ClientHarness client = suiteClient(domain);
            assertEq(client.optimisticSeconds(), APP_OPTIMISTIC_SECONDS, "!optimisticSeconds");
            // Use another domain for testing mock
            uint32 foreign = foreignDomain(domain);
            assertEq(client.trustedSender(foreign), mockTrustedSender(foreign), "!trustedSender");
            assertEq(client.trustedSender(0), bytes32(0), "!trustedSender(0)");
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        TESTS: SENDING MESSAGE                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_sendMessage_noTips() public {
        createEmptyTips();
        createClientSentMessage({ context: clientLocalToRemote, mockTips: false });
        // Should be able to send a message without tips
        expectDispatch();
        clientSendMessage();
    }

    function test_sendMessage_withTips() public {
        createClientSentMessage({ context: clientLocalToRemote, mockTips: true });
        // Should be able to send a message with tips
        expectDispatch();
        clientSendMessage();
    }

    function test_sendMessage_revert_noRecipient() public {
        messageOrigin = DOMAIN_LOCAL;
        messageDestination = 0;
        // trustedSender(0) = 0, and you should not be able to send a message without recipient
        clientSendMessage({ revertMessage: "BasicClient: !recipient" });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: RECEIVING MESSAGE                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_handle() public {
        createMessagesSubmitAttestation({ context: mockSenderRemoteToLocal });
        skip(APP_OPTIMISTIC_SECONDS);
        expectLogClientMessage();
        destinationExecute({ domain: DOMAIN_LOCAL, index: 0 });
    }

    function test_handle_revert_notDestination() public {
        ClientHarness client = suiteClient(DOMAIN_LOCAL);
        vm.expectRevert("BasicClient: !destination");
        vm.prank(attacker);
        // Only Destination should be able to call client.handle()
        client.handle(0, 0, bytes32(0), 0, "");
    }

    function test_handle_revert_optimisticPeriodNotPassed() public {
        createMessagesSubmitAttestation({ context: mockSenderRemoteToLocal });
        skip(APP_OPTIMISTIC_SECONDS - 1);
        // When optimistic period is not over, executing should not be possible
        destinationExecute({ domain: DOMAIN_LOCAL, index: 0, revertMessage: "!optimisticSeconds" });
    }

    function test_handle_revert_optimisticPeriodTooSmall() public {
        basicClients[DOMAIN_LOCAL] = deployBasicClient({
            domain: DOMAIN_LOCAL,
            optimisticSeconds: APP_OPTIMISTIC_SECONDS + 1
        });
        // Client optimistic period is higher than one specified in the message
        createMessagesSubmitAttestation({ context: mockSenderRemoteToLocal });
        skip(APP_OPTIMISTIC_SECONDS);
        // Message's optimistic period is over, but client specifies a longer period
        // Client enforces optimistic period by default, so tx will revert
        destinationExecute({
            domain: DOMAIN_LOCAL,
            index: 0,
            revertMessage: "Client: !optimisticSeconds"
        });
    }

    function test_handle_revert_unknownSender() public {
        createMessagesSubmitAttestation({ context: userRemoteToLocal });
        skip(APP_OPTIMISTIC_SECONDS);
        // Message sent by anyone other than trustedSender should be rejected
        destinationExecute({
            domain: DOMAIN_LOCAL,
            index: 0,
            revertMessage: "BasicClient: !trustedSender"
        });
    }
}
