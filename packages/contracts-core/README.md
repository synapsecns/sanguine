<br/>
<p align="center">
<a href="https://interchain.synapseprotocol.com/" target="_blank">
<img src="https://raw.githubusercontent.com/synapsecns/sanguine/master/assets/interchain-logo.svg" width="225" alt="Synapse Interchain logo">
</a>
</p>
<br/>

[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/synapsecns/sanguine/foundry-tests.yml?style=flat-square&label=Forge%20Tests)](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/synapsecns/sanguine/solidity.yml?style=flat-square&label=Slither)](https://github.com/synapsecns/sanguine/actions/workflows/solidity.yml)
[![Static Badge](https://img.shields.io/badge/Forge-Docs-f?style=flat-square&logo=gitbook)](https://interchain-contracts.synapseprotocol.com/)
[![npm](https://img.shields.io/npm/v/%40synapsecns%2Fcontracts-core?style=flat-square)](https://www.npmjs.com/package/@synapsecns/contracts-core)


`@synapsecns/contracts-core` contain the Solidity contracts used within the Synapse Interchain Network messaging system.

# Usage
These contracts can be installed with:

`npm i @synapsecns/contracts-core`

Please refer to our [usage guide](https://docs.synapseprotocol.com/synapse-interchain-network-sin/build-on-the-synapse-interchain-network) or [examples](contracts/client/TestClient.sol)

## Directory Structure

<pre>
root
├── <a href="./contracts">contracts</a>: Contains core contracts
│   ├── <a href="./contracts/base">base</a>: Base contracts of the protocol
│   ├── <a href="./contracts/client">client</a>: Client contracts for callers of the messaging system.
│   ├── <a href="./contracts/events">events</a>: Event types
│   ├── <a href="./contracts/hubs">hubs</a>: Hubs
│   ├── <a href="./contracts/inbox">inbox</a>: Inbox contracts
│   ├── <a href="./contracts/interfaces">interfaces</a>: Interfaces
│   ├── <a href="./contracts/libs">libs</a>: Library contracts
│   ├── <a href="./contracts/manager">Manager</a>: Manager contracts
├── <a href="./deployments">deployments</a>: Non-devnet deployments of the contracts
├── <a href="./lib">lib</a>: Git-module based dependencies
├── <a href="./script">script</a>: Scripts for deploying + interacting with contracts
├── <a href="./test">test</a>: Test contracts
</pre>
