// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../utils/SynapseTestSuite.t.sol";

import { ReportTools } from "./libs/ReportTools.t.sol";
import { MessageTools } from "./libs/MessageTools.t.sol";

abstract contract OriginTools is MessageTools, SynapseTestSuite, ReportTools {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Create a dispatched message: given {context} and mock {body, recipient, optimistic period}
    function createDispatchedMessage(MessageContext memory context, bool mockTips) public {
        createDispatchedMessage(
            context,
            mockTips,
            MOCK_BODY,
            MOCK_RECIPIENT,
            MOCK_OPTIMISTIC_SECONDS
        );
    }

    // Create a dispatched message: given {context, body, recipient, optimistic period}
    // pass MOCK_X constant to mock field X instead
    function createDispatchedMessage(
        MessageContext memory context,
        bool mockTips,
        bytes memory body,
        bytes32 recipient,
        uint32 optimisticSeconds
    ) public {
        createMessage({
            origin: context.origin,
            sender: _getSender(context, recipient),
            nonce: _nextOriginNonce(context.origin, context.destination),
            destination: context.destination,
            mockTips: mockTips,
            body: body,
            recipient: recipient,
            optimisticSeconds: optimisticSeconds
        });
        // Insert the message hash into the local Merkle Tree to predict the next root
        bytes32 expectedLeaf = keccak256(messageRaw);
        proofGen.insert(expectedLeaf);
    }

    // Chain's default Notary attestation for given domain's Origin: current root and current nonce
    function createSuggestedAttestation(uint32 origin, uint32 destination) public {
        (uint32 nonce, bytes32 root, uint40 blockNumber, uint40 timestamp) = suiteOrigin(origin)
            .suggestNonceRoot(destination);
        saveAttestationMetadata(blockNumber, timestamp);
        createAttestation(origin, destination, nonce, root);
    }

    // Chain's given Notary attestation for given domain's Origin: current root and current nonce
    function createSuggestedAttestation(
        uint32 origin,
        uint32 destination,
        uint256 notaryIndex
    ) public {
        (uint32 nonce, bytes32 root, uint40 blockNumber, uint40 timestamp) = suiteOrigin(origin)
            .suggestNonceRoot(destination);
        saveAttestationMetadata(blockNumber, timestamp);
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: nonce,
            root: root,
            guardIndex: 0,
            notaryIndex: notaryIndex
        });
    }

    // Chain's default Notary attestation for given domain's Origin: current root and fake nonce
    function createFraudAttestation(
        uint32 origin,
        uint32 destination,
        uint32 fakeNonce
    ) public {
        (uint32 nonce, bytes32 root, uint40 blockNumber, uint40 timestamp) = suiteOrigin(origin)
            .suggestNonceRoot(destination);
        saveAttestationMetadata(blockNumber, timestamp);
        require(nonce != fakeNonce, "Failed to provide wrong nonce");
        createAttestation(origin, destination, fakeNonce, root);
    }

    // Chain's given Notary attestation for given domain's Origin: current root and fake nonce
    function createFraudAttestation(
        uint32 origin,
        uint32 destination,
        uint32 fakeNonce,
        uint256 notaryIndex
    ) public {
        (uint32 nonce, bytes32 root, uint40 blockNumber, uint40 timestamp) = suiteOrigin(origin)
            .suggestNonceRoot(destination);
        saveAttestationMetadata(blockNumber, timestamp);
        require(nonce != fakeNonce, "Failed to provide wrong nonce");
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: fakeNonce,
            root: root,
            guardIndex: 0,
            notaryIndex: notaryIndex
        });
    }

    // Chain's default Notary attestation for given domain's Origin: fake root and current nonce
    function createFraudAttestation(
        uint32 origin,
        uint32 destination,
        bytes32 fakeRoot
    ) public {
        (uint32 nonce, bytes32 root, uint40 blockNumber, uint40 timestamp) = suiteOrigin(origin)
            .suggestNonceRoot(destination);
        saveAttestationMetadata(blockNumber, timestamp);
        require(root != fakeRoot, "Failed to provide wrong nonce");
        createAttestation(origin, destination, nonce, fakeRoot);
    }

    // Chain's given Notary attestation for given domain's Origin: fake root and current nonce
    function createFraudAttestation(
        uint32 origin,
        uint32 destination,
        bytes32 fakeRoot,
        uint256 notaryIndex
    ) public {
        (uint32 nonce, bytes32 root, uint40 blockNumber, uint40 timestamp) = suiteOrigin(origin)
            .suggestNonceRoot(destination);
        saveAttestationMetadata(blockNumber, timestamp);
        require(root != fakeRoot, "Failed to provide wrong nonce");
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: nonce,
            root: fakeRoot,
            guardIndex: 0,
            notaryIndex: notaryIndex
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectDispatch() public {
        bytes32 expectedRoot = proofGen.getRoot(messageNonce);
        vm.expectEmit(true, true, true, true);
        emit Dispatch({
            messageHash: keccak256(messageRaw),
            nonce: messageNonce,
            destination: messageDestination,
            root: expectedRoot,
            tips: tipsRaw,
            message: messageRaw
        });
    }

    // Events when a valid fraud report is presented
    function expectFraudAttestationEvents() public {
        vm.expectEmit(true, true, true, true);
        emit FraudAttestation(attestationGuards, attestationNotaries, attestationRaw);
        for (uint256 i = 0; i < attestationGuards.length; ++i) {
            vm.expectEmit(true, true, true, true);
            emit GuardSlashed(attestationGuards[i], broadcaster);
        }
        for (uint256 i = 0; i < attestationNotaries.length; ++i) {
            vm.expectEmit(true, true, true, true);
            emit NotarySlashed(attestationNotaries[i], reportGuard, broadcaster);
        }
    }

    // Events when Notary fraud is reported without a guard's fraud report
    function expectFraudAttestationEventsNoGuard() public {
        reportGuard = address(0);
        expectFraudAttestationEvents();
    }

    function expectIncorrectReportEvents() public {
        vm.expectEmit(true, true, true, true);
        emit IncorrectReport(reportGuard, reportRaw);
        vm.expectEmit(true, true, true, true);
        emit GuardSlashed(reportGuard, broadcaster);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger origin.dispatch() with saved data and expect a revert
    function originDispatch(bytes memory revertMessage) public {
        OriginHarness origin = suiteOrigin(messageOrigin);
        vm.expectRevert(revertMessage);
        _dispatch(origin);
    }

    // Trigger origin.submitAttestation() with saved data and expect a revert
    function originSubmitAttestation(uint32 domain, bytes memory revertMessage) public {
        OriginHarness origin = suiteOrigin(domain);
        vm.expectRevert(revertMessage);
        _submitAttestation(origin);
    }

    // Trigger origin.submitReport() with saved data and expect a revert
    function originSubmitReport(uint32 domain, bytes memory revertMessage) public {
        OriginHarness origin = suiteOrigin(domain);
        vm.expectRevert(revertMessage);
        _submitReport(origin);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger origin.dispatch() with saved data and check return values
    function originDispatch() public {
        OriginHarness origin = suiteOrigin(messageOrigin);
        (uint32 nonce, bytes32 hash) = _dispatch(origin);
        assertEq(nonce, messageNonce, "!return: messageNonce");
        assertEq(hash, keccak256(messageRaw), "!return: messageHash");
    }

    // Trigger origin.submitAttestation() with saved data and check return value
    function originSubmitAttestation(uint32 domain, bool returnValue) public {
        OriginHarness origin = suiteOrigin(domain);
        assertEq(_submitAttestation(origin), returnValue, "!returnValue");
    }

    // Trigger origin.submitReport() with saved data and check return value
    function originSubmitReport(uint32 domain, bool returnValue) public {
        OriginHarness origin = suiteOrigin(domain);
        assertEq(_submitReport(origin), returnValue, "!returnValue");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      INTERNAL HELPERS: TRIGGERS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger origin.dispatch(), using saved data for msg.sender
    function _dispatch(OriginHarness origin) internal returns (uint32, bytes32) {
        // Give gas to cover the tips
        hoax(messageSenderAddress);
        return
            origin.dispatch{ value: tipsTotal }(
                messageDestination,
                messageRecipient,
                messageOptimisticSeconds,
                tipsRaw,
                messageBody
            );
    }

    // Trigger origin.submitAttestation(), using Broadcaster for msg.sender
    function _submitAttestation(OriginHarness origin) internal returns (bool) {
        reportGuard = address(0);
        vm.prank(broadcaster);
        return origin.submitAttestation(attestationRaw);
    }

    // Trigger origin.submitReport(), using Broadcaster for msg.sender
    function _submitReport(OriginHarness origin) internal returns (bool) {
        vm.prank(broadcaster);
        return origin.submitReport(reportRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                 INTERNAL HELPERS: SUBMIT ATTESTATION                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Checks whether given Notary is active, submits attestation and does the final checks
    function _testSubmitAttestation(uint32 domain, bool isValidAttestation) internal {
        // Get initial state and check if Notary is active
        (
            OriginHarness origin,
            uint256 guardsAmount,
            uint256 notariesAmount
        ) = _preSubmitAttestation(domain);
        // If attestation is fraud, expect corresponding events
        if (!isValidAttestation) expectFraudAttestationEvents();
        // Should return true, if and only if the attestation is valid
        originSubmitAttestation(domain, isValidAttestation);
        // Check if the Notary was removed
        _postSubmitAttestation(origin, guardsAmount, notariesAmount, isValidAttestation);
    }

    // Checks after attestation was submitted
    // Notary was removed if and only if their attestation was fraud
    function _postSubmitAttestation(
        OriginHarness origin,
        uint256 _guardsAmount,
        uint256 notariesAmount,
        bool isValidAttestation
    ) internal {
        // Check if every Guard signer was removed
        for (uint256 i = 0; i < attestationGuards.length; ++i) {
            assertEq(
                origin.isActiveAgent({ _domain: 0, _account: attestationGuards[i] }),
                isValidAttestation,
                "Wrong Guard active status"
            );
        }
        // Check if amount of Guards changed
        assertEq(
            origin.amountAgents({ _domain: 0 }),
            _guardsAmount - (isValidAttestation ? 0 : attestationGuards.length),
            "Wrong amount of notaries"
        );
        // Check if every Notary signer was removed
        for (uint256 i = 0; i < attestationNotaries.length; ++i) {
            assertEq(
                origin.isActiveAgent(ra.destination, attestationNotaries[i]),
                isValidAttestation,
                "Wrong Notary active status"
            );
        }
        // Check if amount of Notaries changed
        assertEq(
            origin.amountAgents(ra.destination),
            notariesAmount - (isValidAttestation ? 0 : attestationNotaries.length),
            "Wrong amount of notaries"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   INTERNAL HELPERS: SUBMIT REPORT                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Checks whether given Guard and Notary are active, submits report and does the final checks
    function _testSubmitReport(
        uint32 domain,
        bool isValidAttestation,
        bool isCorrectReport
    ) internal {
        // Get initial state and check if Guard and Notary are active
        (OriginHarness origin, uint256 notariesAmount, uint256 guardsAmount) = _preSubmitReport(
            domain
        );
        // If report is incorrect, expect corresponding events
        if (!isCorrectReport) expectIncorrectReportEvents();
        if (!isValidAttestation) {
            // If attestation is fraud, expect corresponding events
            if (isCorrectReport) {
                // Guard should be mentioned in the events, as their report is correct
                expectFraudAttestationEvents();
            } else {
                // Guard should not be mentioned in the events, as their report is incorrect
                expectFraudAttestationEventsNoGuard();
            }
        }
        // Should return true, if and only if the report is correct (regardless of attestation flag)
        originSubmitReport(domain, isCorrectReport);
        // Check if the Notary and Guard were removed
        _postSubmitReport(
            origin,
            notariesAmount,
            isValidAttestation,
            guardsAmount,
            isCorrectReport
        );
    }

    // Checks after attestation was submitted
    // Notary was removed if and only if their attestation was fraud
    // Guard was removed if and only if their report was incorrect
    function _postSubmitReport(
        OriginHarness origin,
        uint256 notariesAmount,
        bool isValidAttestation,
        uint256 guardsAmount,
        bool isCorrectReport
    ) internal {
        // Do the checks for Notary
        _postSubmitAttestation(origin, guardsAmount, notariesAmount, isValidAttestation);
        // Check if given Guard was removed
        assertEq(
            origin.isActiveAgent({ _domain: 0, _account: reportGuard }),
            isCorrectReport,
            "Wrong Guard active status"
        );
        // Check if amount of Guards changed
        assertEq(
            origin.amountAgents({ _domain: 0 }),
            guardsAmount - (isCorrectReport ? 0 : 1),
            "Wrong amount of guards"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Returns state before submitting attestation and checks that Notary is active
    function _preSubmitAttestation(uint32 domain)
        internal
        view
        returns (
            OriginHarness origin,
            uint256 guardsAmount,
            uint256 notariesAmount
        )
    {
        origin = suiteOrigin(domain);
        guardsAmount = origin.amountAgents({ _domain: 0 });
        notariesAmount = origin.amountAgents(ra.destination);
        // Sanity check: guards were active
        for (uint256 i = 0; i < attestationGuards.length; ++i) {
            require(
                origin.isActiveAgent({ _domain: 0, _account: attestationGuards[i] }),
                "Guard wasn't active"
            );
        }
        // Sanity check: notaries were active
        for (uint256 i = 0; i < attestationNotaries.length; ++i) {
            require(
                origin.isActiveAgent(ra.destination, attestationNotaries[i]),
                "Notary wasn't active"
            );
        }
    }

    // Returns state before submitting report and checks that Guard and Notary are active
    function _preSubmitReport(uint32 domain)
        internal
        view
        returns (
            OriginHarness origin,
            uint256 notariesAmount,
            uint256 guardsAmount
        )
    {
        // Get Notary-related state and perform Notary check
        (origin, guardsAmount, notariesAmount) = _preSubmitAttestation(domain);
        // Sanity check: guard was active
        require(origin.isActiveAgent({ _domain: 0, _account: reportGuard }), "Guard wasn't active");
    }

    // Returns nonce for the next message dispatched via Origin on given domain
    function _nextOriginNonce(uint32 origin, uint32 destination) internal view returns (uint32) {
        return suiteOrigin(origin).nonce(destination) + 1;
    }

    // Returns "sender" field that Origin should be using when constructing a message.
    // A special value of SYSTEM_ROUTER (see SystemCall.sol) will be used if an only if:
    // - Message was sent by System Router, specifying SYSTEM_ROUTER as recipient
    // Otherwise, sender address will be used, casted to bytes32, preserving alignment
    function _getSender(MessageContext memory context, bytes32 recipient)
        internal
        view
        returns (bytes32)
    {
        if (
            context.sender == address(suiteSystemRouter({ domain: context.origin })) &&
            recipient == SystemCall.SYSTEM_ROUTER
        ) {
            return SystemCall.SYSTEM_ROUTER;
        } else {
            return addressToBytes32(context.sender);
        }
    }
}
