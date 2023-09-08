# @synapsecns/contracts-core

`@synapsecns/contracts-core` contain the Solidity contracts used within the Synapse Optimistic messaging system

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
