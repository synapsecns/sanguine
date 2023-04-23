// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";
import {ISystemRegistry} from "../../../contracts/interfaces/ISystemRegistry.sol";
import {SystemContractTest} from "./SystemContract.t.sol";
import {AgentFlag, SynapseTest} from "../../utils/SynapseTest.t.sol";

abstract contract SystemRegistryTest is SystemContractTest {
    function expectAgentSlashed(uint32 domain, address agent, address prover) public {
        vm.expectEmit(localAgentManager());
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        vm.expectEmit(systemContract());
        emit AgentSlashed(domain, agent, prover);
    }
}
