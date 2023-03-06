// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../contracts/libs/State.sol";

import "../utils/libs/SynapseStructs.t.sol";
import "../utils/SynapseProofs.t.sol";
import "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract OriginTest is SynapseTest, SynapseProofs {
    // Deploy Production version of Origin and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_ORIGIN) {}

    function test_setupCorrectly() public {
        // Check Messaging addresses
        assertEq(
            address(ISystemContract(origin).systemRouter()),
            address(systemRouter),
            "!systemRouter"
        );
        // Check Agents
        // Origin should know about agents from all domains, including Guards
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                assertTrue(IAgentRegistry(origin).isActiveAgent(domain, agent), "!agent");
            }
        }
    }

    function test_dispatch() public {
        address sender = makeAddr("Sender");
        address recipient = makeAddr("Recipient");
        uint32 period = 1 minutes;
        bytes memory tips = TipsLib.emptyTips();
        bytes memory body = "test body";

        RawMessage[] memory rawMessages = new RawMessage[](MESSAGES);
        bytes[] memory messages = new bytes[](MESSAGES);
        bytes32[] memory roots = new bytes32[](MESSAGES);
        for (uint32 i = 0; i < MESSAGES; ++i) {
            rawMessages[i] = RawMessage(
                RawHeader({
                    origin: DOMAIN_LOCAL,
                    sender: addressToBytes32(sender),
                    nonce: i + 1,
                    destination: DOMAIN_REMOTE,
                    recipient: addressToBytes32(recipient),
                    optimisticSeconds: period
                }),
                RawTips(0, 0, 0, 0),
                body
            );
            (messages[i], ) = rawMessages[i].castToMessage();
            insertMessage(messages[i]);
            roots[i] = getRoot(i + 1);
        }

        // Expect Origin Events
        for (uint32 i = 0; i < MESSAGES; ++i) {
            // 1 block is skipped after each dispatched message
            OriginState memory state = OriginState(
                roots[i],
                uint40(block.number + i),
                uint40(block.timestamp + i * BLOCK_TIME)
            );
            vm.expectEmit(true, true, true, true);
            emit StateSaved(state.formatOriginState(DOMAIN_LOCAL, i + 1));
            vm.expectEmit(true, true, true, true);
            emit Dispatched(keccak256(messages[i]), i + 1, DOMAIN_REMOTE, messages[i]);
        }

        for (uint32 i = 0; i < MESSAGES; ++i) {
            vm.prank(sender);
            (uint32 messageNonce, bytes32 messageHash) = InterfaceOrigin(origin).dispatch(
                DOMAIN_REMOTE,
                addressToBytes32(recipient),
                period,
                tips,
                body
            );
            // Check return values
            assertEq(messageNonce, i + 1, "!messageNonce");
            assertEq(messageHash, keccak256(messages[i]), "!messageHash");
            skipBlock();
        }
    }

    function test_states() public {
        IStateHub hub = IStateHub(origin);
        // Check initial States
        assertEq(hub.statesAmount(), 1, "!initial statesAmount");
        // Initial state was saved "1 block ago"
        OriginState memory state = OriginState(
            EMPTY_ROOT,
            uint40(block.number - 1),
            uint40(block.timestamp - BLOCK_TIME)
        );
        assertEq(hub.suggestState(0), state.formatOriginState(DOMAIN_LOCAL, 0), "!state: 0");
        assertEq(hub.suggestState(0), hub.suggestLatestState(), "!latest state: 0");
        uint40 initialBN = uint40(block.number);
        uint40 initialTS = uint40(block.timestamp);
        // Dispatch some messages
        test_dispatch();
        // Check saved States
        assertEq(hub.statesAmount(), MESSAGES + 1, "!statesAmount");
        assertEq(hub.suggestState(0), state.formatOriginState(DOMAIN_LOCAL, 0), "!suggestState: 0");
        for (uint32 i = 0; i < MESSAGES; ++i) {
            state = OriginState(getRoot(i + 1), initialBN + i, uint40(initialTS + i * BLOCK_TIME));
            assertEq(
                hub.suggestState(i + 1),
                state.formatOriginState(DOMAIN_LOCAL, i + 1),
                "!suggestState"
            );
        }
        assertEq(
            hub.suggestLatestState(),
            state.formatOriginState(DOMAIN_LOCAL, MESSAGES),
            "!suggestLatestState"
        );
    }
}
