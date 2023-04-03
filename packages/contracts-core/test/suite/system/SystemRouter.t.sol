// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceOrigin, SystemContractMock} from "../../mocks/OriginMock.t.sol";
import {SynapseTest} from "../../utils/SynapseTest.t.sol";

import {SystemEntity, RawHeader, RawSystemCall, RawSystemMessage} from "../../utils/libs/SynapseStructs.t.sol";

import {Random} from "../../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract SystemRouterTest is SynapseTest {
    // Deploy mocks for every other contract except for the SynapseRouter
    constructor() SynapseTest(0) {}

    // ════════════════════════════════════════ TESTS: SENDING SYSTEM CALLS ════════════════════════════════════════════

    function test_sendSystemMessage(RawSystemMessage memory rsm, RawHeader memory rh, uint256 words) public {
        words = words % MAX_SYSTEM_CALL_WORDS;
        rsm.callData.args = Random("sendSystemMessage").nextBytesWords(words);
        // Make sure sender/recipient are valid SystemEntity values
        rsm.boundEntities();
        address sender = systemAddress(SystemEntity(rsm.sender));
        bytes memory expectedCall = abi.encodeWithSelector(
            InterfaceOrigin.sendSystemMessage.selector, rh.destination, rh.optimisticPeriod, rsm.formatSystemMessage()
        );
        vm.expectCall(origin, expectedCall);
        vm.prank(sender);
        systemRouter.systemCall(
            rh.destination, rh.optimisticPeriod, SystemEntity(rsm.recipient), rsm.callData.formatCallData()
        );
    }

    function test_sendSystemMessage_revert_notSystemContract(
        RawSystemMessage memory rsm,
        RawHeader memory rh,
        address caller
    ) public {
        vm.assume(caller != address(lightManager) && caller != destination && caller != origin);
        vm.assume(rh.destination != DOMAIN_LOCAL);
        rsm.callData.args = "";
        // Make sure sender/recipient are valid SystemEntity values
        rsm.boundEntities();
        vm.expectRevert("Unauthorized caller");
        vm.prank(caller);
        systemRouter.systemCall(
            rh.destination, rh.optimisticPeriod, SystemEntity(rsm.recipient), rsm.callData.formatCallData()
        );
    }

    // ═══════════════════════════════════════ TESTS: RECEIVING SYSTEM CALLS ═══════════════════════════════════════════

    function test_receiveSystemMessage(RawSystemCall memory rsc, bytes32 data) public {
        rsc.systemMessage.callData.selector = SystemContractMock.remoteMockFunc.selector;
        rsc.systemMessage.callData.args = bytes.concat(data);
        // Make sure sender/recipient are valid SystemEntity values
        rsc.systemMessage.boundEntities();
        address recipient = systemAddress(SystemEntity(rsc.systemMessage.recipient));
        bytes memory expectedCall = rsc.callPayload();
        vm.expectCall(recipient, expectedCall);
        vm.prank(destination);
        systemRouter.receiveSystemMessage(
            rsc.origin, rsc.nonce, rsc.rootSubmittedAt, rsc.systemMessage.formatSystemMessage()
        );
    }

    function test_receiveSystemMessage_revert_notDestination(RawSystemCall memory rsc, address caller) public {
        rsc.systemMessage.callData.args = "";
        vm.assume(caller != destination);
        // Make sure sender/recipient are valid SystemEntity values
        rsc.systemMessage.boundEntities();
        vm.expectRevert("SystemRouter: !destination");
        vm.prank(caller);
        systemRouter.receiveSystemMessage(
            rsc.origin, rsc.nonce, rsc.rootSubmittedAt, rsc.systemMessage.formatSystemMessage()
        );
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function systemAddress(SystemEntity entity) public view returns (address) {
        if (entity == SystemEntity.AgentManager) return address(lightManager);
        if (entity == SystemEntity.Destination) return destination;
        if (entity == SystemEntity.Origin) return origin;
        revert("Unsupported enum value");
    }
}
