# @synapsecns/contracts

`@synapsencs/contracts` contain the Solidity contracts used within the Synapse Optimistic messaging system

# Setup

Foundry is used to test the contracts. To run the tests, make sure Foundry is installed, by running `curl https://sh.rustup.rs -sSf | sh`, then `. ~/.<zshrc/bashrc>`. More information about installation is here: https://book.getfoundry.sh/getting-started/installation

After, ensure all node packages are installed by running `yarn` or `npm i` in `packages/contracts`. Running `git submodule update --init --recursive` will populate the `forge-std`. From here, running `forge test` will execute the tests in `packages/contracts/test`, with each file `<ContractName>.t.sol` representing a test for the file `<ContractName>.sol`.

## Gotchas:

[`sol-merger`](https://github.com/RyuuGan/sol-merger) [cannot recursively resolve dependencies in node_modules](https://github.com/RyuuGan/sol-merger/issues/58). As a result, you might see an error like this:

``
Error: ENOENT: no such file or directory, stat '/path/to/sanguine/packages/contracts/node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol' +0ms
``

If this happens, verify the dependency is in the `workspaces.nohoist` section of the package.json in the root of this repository.

