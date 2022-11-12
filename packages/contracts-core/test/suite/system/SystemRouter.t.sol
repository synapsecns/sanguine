// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/system/SystemRouterTools.t.sol";

// solhint-disable func-name-mixedcase
contract SystemRouterTest is SystemRouterTools {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function setUp() public override {
        super.setUp();
        saveSystemContracts();
    }

    function test_constructor() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            assertEq(suiteSystemRouter(domain).origin(), address(suiteOrigin(domain)), "!origin");
            assertEq(
                suiteSystemRouter(domain).destination(),
                address(suiteDestination(domain)),
                "!destination"
            );
        }
    }

    // TODO: System Router can consume a message "sent" from a non-existing chain
    // Make sure that this kind of attack is prevented:
    // Attestations and messages from "unknown" domains should be rejected on Destination level
    function test_trustedSender(uint32 remoteDomain) public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            assertEq(
                suiteSystemRouter(domain).trustedSender(remoteDomain),
                SystemCall.SYSTEM_ROUTER,
                "!trustedSender"
            );
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TESTS: SYSTEM CALL (REVERTS)                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_systemCall_revert_local_notConfigured() public {
        prepareMisconfiguredTest({ origin: DOMAIN_LOCAL, destination: DOMAIN_LOCAL });
        // System Call should revert, when a system contract is not specified in SystemRouter
        systemRouterCall({ index: 0, revertMessage: "System Contract not set" });
        // System MultiCall should revert, when a system contract is not specified in SystemRouter
        systemRouterMultiCall({ revertMessage: "System Contract not set" });
    }

    function test_systemCall_revert_remote_notConfigured() public {
        prepareMisconfiguredTest({ origin: DOMAIN_REMOTE, destination: DOMAIN_LOCAL });
        vm.expectRevert("System Contract not set");
        // Mock executing of the message on Destination
        vm.prank(address(suiteDestination(DOMAIN_LOCAL)));
        suiteSystemRouter(DOMAIN_LOCAL).handle({
            _origin: DOMAIN_REMOTE,
            _nonce: 1,
            _sender: SystemCall.SYSTEM_ROUTER,
            _rootSubmittedAt: block.timestamp,
            _message: abi.encode(formattedSystemCalls)
        });
    }

    function test_systemCall_revert_notSystemContract(address sender) public {
        for (uint256 o = 0; o < DOMAINS; ++o) {
            uint32 origin = domains[o];
            // Restrict to calls from NOT a system contract only
            if (_isSystemEntity({ domain: origin, addr: sender })) continue;
            for (uint256 d = 0; d < DOMAINS; ++d) {
                uint32 destination = domains[d];
                MessageContext memory context = MessageContext(origin, sender, destination);
                createSystemCall({
                    context: context,
                    recipient: address(suiteOrigin(destination))
                });
                // System Call (to any destination) fails, if caller is not a system contract
                systemRouterCall({ index: 0, revertMessage: "Unauthorized caller" });
                // System MultiCall (to any destination) fails, if caller is not a system contract
                systemRouterMultiCall({ revertMessage: "Unauthorized caller" });
            }
        }
    }

    function test_systemCall_revert_notSystemCall() public {
        formattedSystemCalls = new bytes[](1);
        // Try passing empty payload as "system call"
        vm.expectRevert("Not a system call");
        // Mock executing of the message on Destination
        vm.prank(address(suiteDestination(DOMAIN_LOCAL)));
        suiteSystemRouter(DOMAIN_LOCAL).handle({
            _origin: DOMAIN_REMOTE,
            _nonce: 1,
            _sender: SystemCall.SYSTEM_ROUTER,
            _rootSubmittedAt: block.timestamp,
            _message: abi.encode(formattedSystemCalls)
        });
    }

    function test_systemCall_revert_unknownRecipient() public {
        MessageContext memory context = MessageContext({
            origin: DOMAIN_REMOTE,
            sender: address(suiteOrigin(DOMAIN_REMOTE)),
            destination: DOMAIN_LOCAL
        });
        // Create a valid system call payload
        createSystemCall({ context: context, recipient: address(suiteOrigin(DOMAIN_LOCAL)) });
        // Change first byte (uint8 recipient) to an invalid value
        bytes memory systemCall = formattedSystemCalls[0];
        // First byte is set to value that doesn't match any enum SystemEntity value
        // All other bytes remain the same
        formattedSystemCalls[0] = abi.encodePacked(
            uint8(type(ISystemRouter.SystemEntity).max) + 1,
            systemCall.ref(0).sliceFrom({ _index: 1, newType: 0 }).clone()
        );
        // Sanity check
        assert(formattedSystemCalls[0].length == systemCall.length);
        vm.expectRevert("Unknown recipient");
        // Mock executing of the message on Destination
        vm.prank(address(suiteDestination(DOMAIN_LOCAL)));
        suiteSystemRouter(DOMAIN_LOCAL).handle({
            _origin: DOMAIN_REMOTE,
            _nonce: 1,
            _sender: SystemCall.SYSTEM_ROUTER,
            _rootSubmittedAt: block.timestamp,
            _message: abi.encode(formattedSystemCalls)
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: SYSTEM CALL                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_systemCall_local() public {
        bytes4 selector = SystemContractHarness.setSensitiveValue.selector;
        // Expect a bunch of "system call succeeded" events
        expectSystemCalls({ domain: DOMAIN_LOCAL, selector: selector });
        // Send system calls (LOCAL -> LOCAL) one by one between all system contracts
        // and check that every system call was successful
        triggerTestSystemCalls({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_LOCAL,
            selector: selector,
            isMultiCall: false
        });
    }

    function test_systemCall_remote() public {
        // Send system calls (REMOTE -> LOCAL) one by one between all system contracts
        // and check that corresponding messages were dispatched
        bytes4 selector = SystemContractHarness.setSensitiveValue.selector;
        triggerTestSystemCalls({
            origin: DOMAIN_REMOTE,
            destination: DOMAIN_LOCAL,
            selector: selector,
            isMultiCall: false
        });
        // Prepare and submit attestation for the dispatched messages
        destinationSubmitAttestationSuggested({
            origin: DOMAIN_REMOTE,
            destination: DOMAIN_LOCAL,
            returnValue: true
        });
        // Expect a bunch of "system call succeeded" events
        expectSystemCalls({ domain: DOMAIN_LOCAL, selector: selector });
        // Execute all dispatched messages on destination chain
        // and check that every system call was successful
        destinationExecuteAll({ domain: DOMAIN_LOCAL });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: SYSTEM MULTI CALL                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_systemMultiCall_local() public {
        bytes4 selector = SystemContractHarness.setSensitiveValue.selector;
        // Expect a bunch of "system call succeeded" events
        expectSystemCalls({ domain: DOMAIN_LOCAL, selector: selector });
        // Send system multicalls (LOCAL -> LOCAL) one by one between all system contracts
        // and check that every system call was successful
        triggerTestSystemCalls({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_LOCAL,
            selector: selector,
            isMultiCall: true
        });
    }

    function test_systemMultiCall_sameRecipient() public {
        bytes4 selector = SystemContractHarness.setSensitiveValue.selector;
        // Expect a bunch of "system call succeeded" events
        expectSystemCallsWrapperTests({
            domain: DOMAIN_LOCAL,
            selector: selector,
            sameRecipient: true
        });
        triggerTestWrapperMultiCall({
            domain: DOMAIN_LOCAL,
            selector: selector,
            sameRecipient: true
        });
    }

    function test_systemMultiCall_sameData() public {
        bytes4 selector = SystemContractHarness.setSensitiveValue.selector;
        // Expect a bunch of "system call succeeded" events
        expectSystemCallsWrapperTests({
            domain: DOMAIN_LOCAL,
            selector: selector,
            sameRecipient: false
        });
        triggerTestWrapperMultiCall({
            domain: DOMAIN_LOCAL,
            selector: selector,
            sameRecipient: false
        });
    }

    function test_systemMultiCall_remote() public {
        // Send system multicalls (REMOTE -> LOCAL) one by one between all system contracts
        // and check that corresponding messages were dispatched
        bytes4 selector = SystemContractHarness.setSensitiveValue.selector;
        triggerTestSystemCalls({
            origin: DOMAIN_REMOTE,
            destination: DOMAIN_LOCAL,
            selector: selector,
            isMultiCall: true
        });
        // Prepare and submit attestation for the dispatched messages
        destinationSubmitAttestationSuggested({
            origin: DOMAIN_REMOTE,
            destination: DOMAIN_LOCAL,
            returnValue: true
        });
        // Expect a bunch of "system call succeeded" events
        expectSystemCalls({ domain: DOMAIN_LOCAL, selector: selector });
        // Execute all dispatched messages on destination chain
        // and check that every system call was successful
        destinationExecuteAll({ domain: DOMAIN_LOCAL });
    }
}
