// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentFlag, AgentStatus } from "../../contracts/libs/Structures.sol";

import { AttestationProofGenerator } from "./proof/AttestationProofGenerator.t.sol";
import { DynamicProofGenerator } from "./proof/DynamicProofGenerator.t.sol";
import { HistoricalProofGenerator } from "./proof/HistoricalProofGenerator.t.sol";

abstract contract SynapseProofs {
    HistoricalProofGenerator internal originGen;
    AttestationProofGenerator internal summitGen;
    DynamicProofGenerator internal agentGen;

    mapping(address => uint256) internal agentIndex;
    mapping(address => uint32) internal agentDomain;
    mapping(address => AgentFlag) internal agentFlag;
    uint256 internal totalAgents;

    constructor() {
        clear();
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseProofs() external {}

    /// @notice Clears proof generators
    function clear() public {
        originGen = new HistoricalProofGenerator();
        summitGen = new AttestationProofGenerator();
        agentGen = new DynamicProofGenerator();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            MESSAGE PROOFS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function insertMessage(bytes memory message) public {
        originGen.insert(keccak256(message));
    }

    function getLatestProof(uint256 index) public view returns (bytes32[] memory proof) {
        return originGen.getLatestProof(index);
    }

    function getProof(uint256 index, uint256 count) public view returns (bytes32[] memory) {
        return originGen.getProof(index, count);
    }

    function getRoot(uint256 count) public view returns (bytes32) {
        return originGen.getRoot(count);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SNAPSHOT PROOFS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function acceptSnapshot(bytes[] memory snapshotStates) public {
        summitGen.acceptSnapshot(snapshotStates);
    }

    function genSnapshotProof(uint256 index) public view returns (bytes32[] memory) {
        return summitGen.generateProof(index);
    }

    function getSnapshotRoot() public view returns (bytes32) {
        return summitGen.root();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             AGENT PROOFS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function addNewAgent(uint32 domain, address agent) public returns (bytes32 newRoot) {
        require(agentIndex[agent] == 0, "Already added");
        uint256 index = ++totalAgents;
        agentIndex[agent] = index;
        agentDomain[agent] = domain;
        agentFlag[agent] = AgentFlag.Active;
        agentGen.update(index, getAgentLeaf(AgentFlag.Active, domain, agent));
        return agentGen.getRoot();
    }

    function updateAgent(AgentFlag flag, address agent) public returns (bytes32 newRoot) {
        uint256 index = agentIndex[agent];
        require(index != 0, "Unknown agent");
        agentFlag[agent] = flag;
        agentGen.update(index, getAgentLeaf(flag, agentDomain[agent], agent));
        return agentGen.getRoot();
    }

    function getAgentRoot() public view returns (bytes32) {
        return agentGen.getRoot();
    }

    function getAgentProof(address agent) public view returns (bytes32[] memory) {
        uint256 index = agentIndex[agent];
        require(index != 0, "Unknown agent");
        return agentGen.getProof(index);
    }

    function getZeroProof() public view returns (bytes32[] memory) {
        // Proof for zero value after the latest added agent
        return agentGen.getProof(totalAgents + 1);
    }

    function getAgentStatus(address _agent) public view returns (AgentStatus memory) {
        uint32 index = uint32(agentIndex[_agent]);
        require(index != 0, "Unknown agent");
        return AgentStatus({ flag: agentFlag[_agent], domain: agentDomain[_agent], index: index });
    }

    function getAgentLeaf(uint256 index) public view returns (bytes32) {
        return agentGen.getLeaf(index);
    }

    function getAgentLeaf(
        AgentFlag _flag,
        uint32 _domain,
        address _agent
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_flag, _domain, _agent));
    }
}
