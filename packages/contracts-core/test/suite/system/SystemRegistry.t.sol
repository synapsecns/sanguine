// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";
import {ISystemRegistry} from "../../../contracts/interfaces/ISystemRegistry.sol";
import {SynapseTest} from "../../utils/SynapseTest.t.sol";

abstract contract SystemRegistryTest is SynapseTest {
    function expectAgentSlashed(uint32 domain, address agent, address prover) public {
        vm.expectEmit();
        emit AgentSlashed(domain, agent, prover);
        vm.expectCall(localAgentManager(), abi.encodeWithSelector(IAgentManager.registrySlash.selector, domain, agent));
    }

    /// @notice Address of AgentManager on tested chain
    function localAgentManager() public view virtual returns (address);
}
