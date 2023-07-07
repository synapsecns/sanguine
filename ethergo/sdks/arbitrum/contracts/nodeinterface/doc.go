// Package nodeinterface contains the interface to the pre-compiled gas estimate info contract  for the Arbitrum chain.
// NodeInterface contains the interface for a pre-compile used by arbitrum for l2 gas estimation. Documentation is here: https://developer.arbitrum.io/arbos/gas#nodeinterfacesol
// the important thing to understand is in the following excerpt from the arbitrum docs:
//
//	To avoid creating new RPC methods for client-side tooling, nitro Geth's InterceptRPCMessage hook provides an opportunity to swap out the message its handling before deriving a transaction from it. The node uses this hook to detect messages sent to the address 0xc8, the location of the fictional NodeInterface contract specified in NodeInterface.sol.
//
// therefore: this cannot be deployed w/ deployers, but merely mocked or called via rpc. For this same reason, we do not need to return Filterers. As this contract does not actually exist.
package nodeinterface
