# @synapsecns/contracts-core

`@synapsecns/contracts-core` contain the Solidity contracts used within the Synapse Optimistic messaging system

# Setup

TODO: update docs here.

## Gotchas:

[`sol-merger`](https://github.com/RyuuGan/sol-merger) [cannot recursively resolve dependencies in node_modules](https://github.com/RyuuGan/sol-merger/issues/58). As a result, you might see an error like this:

`Error: ENOENT: no such file or directory, stat '/path/to/sanguine/packages/contracts-core/node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol' +0ms`

If this happens, verify the dependency is in the `workspaces.nohoist` section of the package.json in the root of this repository.


