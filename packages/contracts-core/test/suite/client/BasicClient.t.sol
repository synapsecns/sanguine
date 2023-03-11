// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/client/BasicClientTools.t.sol";

// solhint-disable func-name-mixedcase
contract BasicClientTest is BasicClientTools {
    function setUp() public override {
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
            BasicClientHarness client = suiteBasicClient(domain);
            assertEq(client.destination(), address(suiteDestination(domain)), "!destination");
            assertEq(client.origin(), address(suiteOrigin(domain)), "!origin");
        }
    }

    function test_setup_harnessMocks() public {
        // Check ClientHarness mocks
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            BasicClientHarness client = suiteBasicClient(domain);
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
        basicClientSendMessage();
    }

    function test_sendMessage_withTips() public {
        createClientSentMessage({ context: clientLocalToRemote, mockTips: true });
        // Should be able to send a message with tips
        expectDispatch();
        basicClientSendMessage();
    }

    function test_sendMessage_revert_noRecipient() public {
        messageOrigin = DOMAIN_LOCAL;
        messageDestination = 0;
        // trustedSender(0) = 0, and you should not be able to send a message without recipient
        basicClientSendMessage({ revertMessage: "BasicClient: !recipient" });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: RECEIVING MESSAGE                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_handle() public {
        createMessagesSubmitAttestation(mockSenderRemoteToLocal);
        skip(APP_OPTIMISTIC_SECONDS);
        expectLogBasicClientMessage();
        destinationExecute({ domain: DOMAIN_LOCAL, index: 0 });
    }

    function test_handle_optimisticPeriodTooSmall() public {
        basicClients[DOMAIN_LOCAL] = deployBasicClient({
            domain: DOMAIN_LOCAL,
            optimisticSeconds: APP_OPTIMISTIC_SECONDS + 1
        });
        // Client optimistic period is higher than period specified in the message
        createMessagesSubmitAttestation({ context: mockSenderRemoteToLocal });
        skip(APP_OPTIMISTIC_SECONDS);
        // Message's optimistic period is over, but client specifies a longer period
        // Basic Client does not enforce optimistic period by default, so message will be executed
        expectLogBasicClientMessage();
        destinationExecute(DOMAIN_LOCAL, 0);
    }

    function test_handle_revert_notDestination() public {
        BasicClientHarness client = suiteBasicClient(DOMAIN_LOCAL);
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
