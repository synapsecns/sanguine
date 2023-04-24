// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";
import {IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";
import {MessagingBaseTest} from "./MessagingBase.t.sol";
import {AgentFlag, SynapseTest} from "../../utils/SynapseTest.t.sol";

abstract contract AgentSecuredTest is MessagingBaseTest {
    function expectAgentSlashed(uint32 domain, address agent, address prover) public {
        vm.expectEmit(localAgentManager());
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        vm.expectEmit(systemContract());
        emit AgentSlashed(domain, agent, prover);
    }
}
