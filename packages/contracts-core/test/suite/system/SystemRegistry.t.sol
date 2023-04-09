// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";
import {ISystemRegistry} from "../../../contracts/interfaces/ISystemRegistry.sol";
import {SystemContractTest} from "./SystemContract.t.sol";
import {SynapseTest} from "../../utils/SynapseTest.t.sol";

abstract contract SystemRegistryTest is SystemContractTest {
    function expectAgentSlashed(uint32 domain, address agent, address prover) public {
        vm.expectEmit();
        emit AgentSlashed(domain, agent, prover);
        vm.expectCall(localAgentManager(), abi.encodeWithSelector(IAgentManager.registrySlash.selector, domain, agent));
    }
}
