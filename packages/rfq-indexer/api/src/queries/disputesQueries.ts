import { db } from '../db'

export const qDisputes = () => {
  return db
    .selectFrom('BridgeProofDisputedEvents')
    .select([
      'BridgeProofDisputedEvents.transactionId as transactionId_dispute',
      'BridgeProofDisputedEvents.blockNumber as blockNumber_dispute',
      'BridgeProofDisputedEvents.blockTimestamp as blockTimestamp_dispute',
      'BridgeProofDisputedEvents.transactionHash as transactionHash_dispute',
      'BridgeProofDisputedEvents.originChainId as originChainId_dispute',
      'BridgeProofDisputedEvents.originChain as originChain_dispute',
    ])
}
