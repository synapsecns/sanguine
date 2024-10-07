import { db } from '../db'

export const qDisputes = ({ activeOnly }: { activeOnly: boolean } = { activeOnly: false}) => {
  let query = db
    .selectFrom('BridgeProofDisputedEvents')
    .leftJoin('BridgeProofProvidedEvents', (join) =>
      // if a proof occurred after this dispute, consider the dispute to be stale/invalid & ignore it
      join
        .onRef('BridgeProofProvidedEvents.transactionId', '=', 'BridgeProofDisputedEvents.transactionId')
        .onRef('BridgeProofProvidedEvents.blockTimestamp', '>', 'BridgeProofDisputedEvents.blockTimestamp')
    )
    .select([
      'BridgeProofDisputedEvents.transactionId as transactionId_dispute',
      'BridgeProofDisputedEvents.blockNumber as blockNumber_dispute',
      'BridgeProofDisputedEvents.blockTimestamp as blockTimestamp_dispute',
      'BridgeProofDisputedEvents.transactionHash as transactionHash_dispute',
      'BridgeProofDisputedEvents.originChainId as originChainId_dispute',
      'BridgeProofDisputedEvents.originChain as originChain_dispute',
    ]);

  if (activeOnly) {
    query = query.where('BridgeProofProvidedEvents.transactionId', 'is', null);
  }

  return query;
}
