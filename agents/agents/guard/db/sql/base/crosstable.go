package base

import "context"

// GetValidAgentRootProof works in x parts:
// 1. Take an agent root that has been seen on a remote chain and finds the block number associated with that root
// on summit from the `AgentTree` model.
// 2. Take all agent roots with a block number greater than or equal to the block number found in step 1.
// 3. Compare all agent roots from step 2, and return the ____ one that has been seen on the chain ID specified.
func (s Store) GetValidAgentRootProof(
	ctx context.Context,
	agentRoot [32]byte,
	chainID uint32,
) (
	validAgentRoot [32]byte,
	validProof [][32]byte,
	err error,
) {

	return
}
