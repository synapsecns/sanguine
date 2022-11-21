// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../utils/SynapseTestSuite.t.sol";

import { OriginTools } from "./OriginTools.t.sol";
import { ReentrantApp } from "../harnesses/client/ReentrantApp.t.sol";

abstract contract DestinationTools is OriginTools {
    uint32 internal constant MESSAGES = 5;

    bytes[] internal rawMessages;
    bytes32[] internal messageHashes;
    uint256 internal rootSubmittedAt;

    ReentrantApp internal reentrantApp;

    // Creates test messages and prepares their merkle proofs for future execution
    function createMessages(MessageContext memory context, address recipient) public {
        bytes32 recipientBytes32 = addressToBytes32(recipient);
        rawMessages = new bytes[](MESSAGES);
        messageHashes = new bytes32[](MESSAGES);
        for (uint32 index = 0; index < MESSAGES; ++index) {
            // Construct a dispatched message
            createDispatchedMessage({
                context: context,
                mockTips: true,
                body: MOCK_BODY,
                recipient: recipientBytes32,
                optimisticSeconds: APP_OPTIMISTIC_SECONDS
            });
            // Save raw message and its hash for later use
            rawMessages[index] = messageRaw;
            messageHashes[index] = keccak256(messageRaw);
            // Dispatch message on remote Origin
            originDispatch();
        }
        // Create merkle proofs for dispatched messages
        proofGen.createTree(messageHashes);
    }

    // Prepare app to receive a message from Destination
    function prepareApp(
        MessageContext memory context,
        AppHarness app,
        uint32 nonce
    ) public {
        // App will revert if any of values passed over by Destination will differ (see AppHarness)
        app.prepare({
            _origin: context.origin,
            _nonce: nonce,
            _sender: addressToBytes32(context.sender),
            _message: _createMockBody(context.origin, context.destination, nonce)
        });
    }

    // Check given message execution
    function checkMessageExecution(
        MessageContext memory context,
        AppHarness app,
        uint32 index
    ) public {
        uint32 nonce = index + 1;
        // Save mock data in app to check against data passed by Destination
        prepareApp(context, app, nonce);
        // Recreate tips used for that message
        createMockTips(nonce);
        expectLogTips();
        expectExecuted({ domain: context.origin, index: index });
        // Trigger Destination.execute() on destination chain
        destinationExecute({ domain: context.destination, index: index });
        // Check executed message status
        assertEq(destinationMessageStatus(context, index), attestationRoot, "!messageStatus");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectAttestationAccepted() public {
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(
            attestationDomain,
            attestationNonce,
            attestationRoot,
            signatureNotary
        );
    }

    function expectExecuted(uint32 domain, uint256 index) public {
        vm.expectEmit(true, true, true, true);
        emit Executed(domain, messageHashes[index]);
    }

    function expectNotaryBlacklisted() public {
        vm.expectEmit(true, true, true, true);
        emit NotaryBlacklisted(attestationNotary, reportGuard, broadcaster, reportRaw);
    }

    function expectLogTips() public {
        vm.expectEmit(true, true, true, true);
        emit LogTips(tipNotary, tipBroadcaster, tipProver, tipExecutor);
    }

    function expectSetConfirmation(uint256 prevConfirmAt, uint256 newConfirmAt) public {
        vm.expectEmit(true, true, true, true);
        emit SetConfirmation(attestationDomain, attestationRoot, prevConfirmAt, newConfirmAt);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger destination.submitAttestation() with saved data and expect a revert
    function destinationSubmitAttestation(uint32 domain, bytes memory revertMessage) public {
        DestinationHarness destination = suiteDestination(domain);
        vm.expectRevert(revertMessage);
        vm.prank(broadcaster);
        destination.submitAttestation(attestationRaw);
    }

    // Trigger destination.submitReport() with saved data and expect a revert
    function destinationSubmitReport(uint32 domain, bytes memory revertMessage) public {
        DestinationHarness destination = suiteDestination(domain);
        vm.expectRevert(revertMessage);
        vm.prank(broadcaster);
        destination.submitReport(reportRaw);
    }

    // Trigger destination.execute() for given message with saved proof and expect a revert
    function destinationExecute(
        uint32 domain,
        uint256 index,
        bytes memory revertMessage
    ) public {
        bytes32[TREE_DEPTH] memory proof = proofGen.getProof(index);
        destinationExecute(domain, index, proof, revertMessage);
    }

    // Trigger destination.execute() for given message with given proof and expect a revert
    function destinationExecute(
        uint32 domain,
        uint256 index,
        bytes32[TREE_DEPTH] memory proof,
        bytes memory revertMessage
    ) public {
        DestinationHarness destination = suiteDestination(domain);
        vm.expectRevert(revertMessage);
        _execute(destination, proof, index);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger destination.setConfirmation() with given and saved data
    function destinationSetConfirmAt(uint32 domain, uint256 newConfirmAt) public {
        DestinationHarness destination = suiteDestination(domain);
        destination.setConfirmation(attestationDomain, attestationRoot, newConfirmAt);
    }

    // Trigger destination.submitAttestation() with saved data and check the return value
    function destinationSubmitAttestation(uint32 domain, bool returnValue) public {
        DestinationHarness destination = suiteDestination(domain);
        vm.prank(broadcaster);
        assertEq(destination.submitAttestation(attestationRaw), returnValue, "!returnValue");
        if (returnValue) {
            rootSubmittedAt = block.timestamp;
        }
    }

    // Trigger destination.submitAttestation() with latest Origin state and check the return value
    function destinationSubmitAttestationSuggested(MessageContext memory context, bool returnValue)
        public
    {
        destinationSubmitAttestationSuggested(context.origin, context.destination, returnValue);
    }

    // Trigger destination.submitAttestation() with latest Origin state and check the return value
    function destinationSubmitAttestationSuggested(
        uint32 origin,
        uint32 destination,
        bool returnValue
    ) public {
        createSuggestedAttestation({ domain: origin });
        destinationSubmitAttestation({ domain: destination, returnValue: returnValue });
    }

    // Trigger destination.submitAttestation() with saved data and check the return value
    function destinationSubmitReport(uint32 domain, bool returnValue) public {
        DestinationHarness destination = suiteDestination(domain);
        vm.prank(broadcaster);
        assertEq(destination.submitReport(reportRaw), returnValue, "!returnValue");
    }

    // Trigger destination.execute() for a saved message
    function destinationExecute(uint32 domain, uint256 index) public {
        DestinationHarness destination = suiteDestination(domain);
        bytes32[TREE_DEPTH] memory proof = proofGen.getProof(index);
        _execute(destination, proof, index);
    }

    // Trigger destination.execute() for all saved messages
    function destinationExecuteAll(uint32 domain) public {
        uint256 messages = rawMessages.length;
        for (uint256 index = 0; index < messages; ++index) {
            destinationExecute(domain, index);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TRIGGER VIEWS (REVERTS)                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger destination.acceptableRoot() for saved data and expect a revert
    function destinationAcceptableRoot(uint32 domain, bytes memory revertMessage) public {
        DestinationHarness destination = suiteDestination(domain);
        vm.expectRevert(revertMessage);
        destination.acceptableRoot(attestationDomain, APP_OPTIMISTIC_SECONDS, attestationRoot);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TRIGGER VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger destination.acceptableRoot() for saved data and pass over its return value
    function destinationAcceptableRoot(uint32 domain) public view returns (bool) {
        return
            suiteDestination(domain).acceptableRoot(
                attestationDomain,
                APP_OPTIMISTIC_SECONDS,
                attestationRoot
            );
    }

    // Trigger destination.messageStatus() for given root and pass over its return value
    function destinationMessageStatus(MessageContext memory context, uint256 index)
        public
        view
        returns (bytes32)
    {
        return
            suiteDestination(context.destination).messageStatus(
                context.origin,
                messageHashes[index]
            );
    }

    // Trigger destination.submittedAt() for saved data and pass over its return value
    function destinationSubmittedAt(uint32 domain) public view returns (uint256) {
        return suiteDestination(domain).submittedAt(attestationDomain, attestationRoot);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      INTERNAL HELPERS: TRIGGERS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger destination.execute() for given message and proof, using Broadcaster as msg.sender
    function _execute(
        DestinationHarness destination,
        bytes32[TREE_DEPTH] memory proof,
        uint256 index
    ) internal {
        vm.prank(broadcaster);
        destination.execute(rawMessages[index], proof, index);
    }
}
