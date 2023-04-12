// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceSummit} from "../../contracts/Summit.sol";

import {AgentFlag, ISystemContract, SynapseTest} from "../utils/SynapseTest.t.sol";
import {IDisputeHub, DisputeHubTest} from "./hubs/DisputeHub.t.sol";

import {fakeState} from "../utils/libs/FakeIt.t.sol";
import {RawExecReceipt, RawState, RawStateIndex, RawSnapshot} from "../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract SummitTipsTest is DisputeHubTest {
    RawState internal state0;
    address internal guard0;
    uint32 internal origin0;

    RawState internal state1;
    address internal guard1;
    uint32 internal origin1;

    RawSnapshot internal snapshot;
    address internal snapNotary;
    bytes32 internal snapRoot;

    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

    function setUp() public override {
        super.setUp();
        guard0 = domains[0].agents[0];
        guard1 = domains[0].agents[1];
        snapNotary = domains[DOMAIN_LOCAL].agent;
        // Prepare test snapshot data
        origin0 = 1;
        state0 = fakeState(origin0);
        origin1 = 2;
        state1 = fakeState(origin1);
        snapshot.states.push(state0);
        snapshot.states.push(state1);
        // Submit snapshots to Summit
        submitGuardSnapshot(guard0, state0);
        submitGuardSnapshot(guard1, state1);
        submitSnapshot(snapNotary, snapshot);
        // Extract snapshot root
        acceptSnapshot(snapshot.formatStates());
        snapRoot = getSnapshotRoot();
    }

    // ══════════════════════════════════════════ TESTS: SUBMIT RECEIPTS ═══════════════════════════════════════════════

    function test_submitReceipt(RawExecReceipt memory re, bool originZero, uint256 attNotaryIndex, bool isSuccess)
        public
    {
        prepareReceipt(re, originZero, attNotaryIndex, isSuccess);
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectEmit();
        emit ReceiptAccepted(DOMAIN_REMOTE, notary, rcptPayload, rcptSignature);
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_signedByGuard(RawExecReceipt memory re) public {
        prepareReceipt(re, false, 0, false);
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(guard0, re);
        vm.expectRevert("Signer is not a Notary");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_wrongNotaryDomain(RawExecReceipt memory re) public {
        // TODO: remove when Notary restrictions are revisited
        prepareReceipt(re, false, 0, false);
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert("Wrong Notary domain");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_notaryInDispute(RawExecReceipt memory re) public {
        prepareReceipt(re, false, 0, false);
        // Put DOMAIN_REMOTE notary in Dispute
        check_submitStateReport(summit, DOMAIN_REMOTE, state0, RawStateIndex(0, 1));
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert("Notary is in dispute");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    function test_submitReceipt_revert_unknownSnapRoot(RawExecReceipt memory re) public {
        vm.assume(re.snapshotRoot != snapRoot);
        bytes32 oldValue = re.snapshotRoot;
        prepareReceipt(re, false, 0, false);
        re.snapshotRoot = oldValue;
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory rcptPayload, bytes memory rcptSignature) = signReceipt(notary, re);
        vm.expectRevert("Unknown snapshot root");
        InterfaceSummit(summit).submitReceipt(rcptPayload, rcptSignature);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function prepareReceipt(RawExecReceipt memory re, bool originZero, uint256 attNotaryIndex, bool isSuccess)
        public
        view
    {
        re.origin = originZero ? origin0 : origin1;
        re.destination = DOMAIN_REMOTE;
        re.snapshotRoot = snapRoot;
        re.attNotary = domains[DOMAIN_REMOTE].agents[attNotaryIndex % DOMAIN_AGENTS];
        vm.assume(re.firstExecutor != address(0));
        if (isSuccess) {
            vm.assume(re.finalExecutor != address(0));
        } else {
            re.finalExecutor = address(0);
        }
        // Make every tip component non-zero
        re.tips.floorTips(1);
    }

    function submitGuardSnapshot(address guard, RawState memory rs) public {
        RawSnapshot memory rawSnap;
        rawSnap.states = new RawState[](1);
        rawSnap.states[0] = rs;
        submitSnapshot(guard, rawSnap);
    }

    function submitSnapshot(address agent, RawSnapshot memory rawSnap) public {
        (bytes memory snapPayload, bytes memory snapSignature) = signSnapshot(agent, rawSnap);
        InterfaceSummit(summit).submitSnapshot(snapPayload, snapSignature);
    }

    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested system contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
