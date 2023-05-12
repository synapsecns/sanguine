// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, AgentStatus} from "../../contracts/libs/Structures.sol";

import {AttestationProofGenerator} from "./proof/AttestationProofGenerator.t.sol";
import {DynamicProofGenerator} from "./proof/DynamicProofGenerator.t.sol";
import {HistoricalProofGenerator} from "./proof/HistoricalProofGenerator.t.sol";

import {RawSnapshot} from "./libs/SynapseStructs.t.sol";

// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract SynapseProofs {
    HistoricalProofGenerator internal originGen;
    AttestationProofGenerator internal summitGen;
    DynamicProofGenerator internal agentGen;

    mapping(address => uint32) internal agentIndex;
    mapping(address => uint32) internal agentDomain;
    mapping(address => AgentFlag) internal agentFlag;
    uint32 internal totalAgents;

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

    // ══════════════════════════════════════════════ MESSAGE PROOFS ═══════════════════════════════════════════════════

    function insertMessage(bytes32 msgHash) public {
        originGen.insert(msgHash);
    }

    function getLatestProof(uint256 index) public view returns (bytes32[] memory proof) {
        return originGen.getLatestProof(index);
    }

    function getProof(uint256 index, uint256 count) public view returns (bytes32[] memory) {
        return originGen.getProof(index, count);
    }

    function getLeaf(uint256 index) public view returns (bytes32) {
        return originGen.getLeaf(index);
    }

    function getRoot(uint256 count) public view returns (bytes32) {
        return originGen.getRoot(count);
    }

    // ══════════════════════════════════════════════ SNAPSHOT PROOFS ══════════════════════════════════════════════════

    function acceptSnapshot(RawSnapshot memory rs) public {
        summitGen.acceptSnapshot(rs.formatStates());
    }

    function genSnapshotProof(uint256 index) public view returns (bytes32[] memory) {
        return summitGen.generateProof(index);
    }

    function getSnapshotRoot() public view returns (bytes32) {
        return summitGen.root();
    }

    // ═══════════════════════════════════════════════ AGENT PROOFS ════════════════════════════════════════════════════

    function addNewAgent(uint32 domain, address agent) public returns (bytes32 newRoot) {
        require(agentIndex[agent] == 0, "Already added");
        uint32 index = ++totalAgents;
        agentIndex[agent] = index;
        agentDomain[agent] = domain;
        agentFlag[agent] = AgentFlag.Active;
        agentGen.update(index, getAgentLeaf(AgentFlag.Active, domain, agent));
        return agentGen.getRoot();
    }

    function updateAgent(AgentFlag flag, address agent) public returns (bytes32 newRoot) {
        uint32 index = agentIndex[agent];
        require(index != 0, "Unknown agent");
        agentFlag[agent] = flag;
        agentGen.update(index, getAgentLeaf(flag, agentDomain[agent], agent));
        return agentGen.getRoot();
    }

    function getAgentRoot() public view returns (bytes32) {
        return agentGen.getRoot();
    }

    function getAgentProof(address agent) public view returns (bytes32[] memory) {
        require(agentIndex[agent] != 0, "Unknown agent");
        return agentGen.getProof(agentIndex[agent]);
    }

    function getZeroProof() public view returns (bytes32[] memory) {
        // Proof for zero value after the latest added agent
        return agentGen.getProof(totalAgents + 1);
    }

    function getAgentStatus(address agent) public view returns (AgentStatus memory) {
        require(agentIndex[agent] != 0, "Unknown agent");
        return AgentStatus({flag: agentFlag[agent], domain: agentDomain[agent], index: agentIndex[agent]});
    }

    function getAgentLeaf(uint256 index) public view returns (bytes32) {
        return agentGen.getLeaf(index);
    }

    function getAgentLeaf(AgentFlag flag, uint32 domain, address agent) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(flag, domain, agent));
    }
}
