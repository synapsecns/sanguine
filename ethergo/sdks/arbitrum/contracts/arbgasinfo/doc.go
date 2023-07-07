// Package arbgasinfo contains the interface to the pre-compiled gas price contract  for the Arbitrum chain.
//
// ArbGasInfo contains the interface for a pre-compile used by arbitrum for l2 gas estimation. Documentation is here: https://developer.arbitrum.io/devs-how-tos/how-to-estimate-gas#an-example-of-how-to-apply-this-formula-in-your-code
// therefore: this cannot be deployed w/ deployers, but merely mocked or called via rpc. For this same reason, we do not need to return Filterers. As this contract does not actually exist.
package arbgasinfo
