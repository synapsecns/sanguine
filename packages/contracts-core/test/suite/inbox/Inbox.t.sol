// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {CallerNotDestination, MustBeSynapseDomain, NotaryInDispute} from "../../../contracts/libs/Errors.sol";
import {InterfaceSummit} from "../../../contracts/interfaces/InterfaceSummit.sol";

import {StatementInboxTest} from "./StatementInbox.t.sol";

import {BaseMock} from "../../mocks/base/BaseMock.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {
    RawExecReceipt,
    RawTips,
    RawTipsProof,
    RawReceiptTips,
    RawState,
    RawStateIndex
} from "../../utils/libs/SynapseStructs.t.sol";

import {Inbox, SynapseTest} from "../../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract InboxTest is StatementInboxTest {
    // Deploy mocks for everything except BondingManager and Inbox
    constructor() SynapseTest(0) {}

    // ═══════════════════════════════════════════════ TESTS: SETUP ════════════════════════════════════════════════════

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = DOMAIN_SYNAPSE;
        address caller = random.nextAddress();
        address agentManager = random.nextAddress();
        address origin_ = random.nextAddress();
        address destination_ = random.nextAddress();
        address summit_ = random.nextAddress();
        Inbox inbox_ = new Inbox(domain);
        vm.prank(caller);
        inbox_.initialize(agentManager, origin_, destination_, summit_);
        assertEq(inbox_.owner(), caller);
        assertEq(inbox_.localDomain(), domain);
        assertEq(inbox_.origin(), origin_);
        assertEq(inbox_.destination(), destination_);
        assertEq(inbox_.summit(), summit_);
        assertEq(inbox_.agentManager(), agentManager);
    }

    function test_setup() public override {
        super.test_setup();
        assertEq(inbox.summit(), summit);
        assertEq(inbox.version(), LATEST_VERSION);
    }

    function test_constructor_revert_notOnSynapseChain(uint32 domain) public {
        vm.assume(domain != DOMAIN_SYNAPSE);
        vm.expectRevert(MustBeSynapseDomain.selector);
        new Inbox(domain);
    }

    function initializeLocalContract() public override {
        Inbox(localContract()).initialize(address(0), address(0), address(0), address(0));
    }

    // ══════════════════════════════════════════ TEST: SUBMIT STATEMENTS ══════════════════════════════════════════════

    function test_submitSnapshot_guard(uint256 agentId, RawState memory rs, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        address guard = getGuard(agentId);
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(guard, rs, rsi);
        vm.expectCall(
            summit,
            abi.encodeWithSelector(
                InterfaceSummit.acceptGuardSnapshot.selector, agentIndex[guard], nextSignatureIndex(), snapPayload
            )
        );
        inbox.submitSnapshot(snapPayload, snapSig);
    }

    function test_submitSnapshot_notary(uint256 domainId, uint256 agentId, RawState memory rs, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        address notary = getNotary(domainId, agentId);
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, rs, rsi);
        vm.expectCall(
            summit,
            abi.encodeWithSelector(
                InterfaceSummit.acceptNotarySnapshot.selector,
                agentIndex[notary],
                nextSignatureIndex(),
                getAgentRoot(),
                snapPayload
            )
        );
        inbox.submitSnapshot(snapPayload, snapSig);
    }

    function test_submitReceipt(
        uint256 domainId,
        uint256 agentId,
        uint256 attNotaryId,
        RawReceiptTips memory receipt,
        uint256 attNonce
    ) public {
        address rcptNotary = getNotary(domainId, agentId);
        receipt.re.destination = DOMAIN_REMOTE;
        receipt.re.attNotary = domains[DOMAIN_REMOTE].agents[attNotaryId % DOMAIN_AGENTS];
        receipt.re.messageHash = receipt.tips.getMessageHash(receipt.rtp);
        (bytes memory receiptPayload, bytes memory receiptSig) = signReceipt(rcptNotary, receipt.re);
        // Set value for getAttestationNonce call
        attNonce = bound(attNonce, 1, type(uint32).max);
        BaseMock(localDestination()).setMockReturnValue(attNonce);
        vm.expectCall(
            summit,
            abi.encodeWithSelector(
                InterfaceSummit.acceptReceipt.selector,
                agentIndex[rcptNotary],
                agentIndex[receipt.re.attNotary],
                nextSignatureIndex(),
                attNonce,
                receipt.tips.encodeTips(),
                receipt.re.formatReceipt()
            )
        );
        inbox.submitReceipt(
            receiptPayload, receiptSig, receipt.tips.encodeTips(), receipt.rtp.headerHash, receipt.rtp.bodyHash
        );
    }

    function test_passReceipt_revert_notDestination(address caller) public {
        vm.assume(caller != localDestination());
        vm.expectRevert(CallerNotDestination.selector);
        vm.prank(caller);
        inbox.passReceipt(0, 0, 0, "");
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
