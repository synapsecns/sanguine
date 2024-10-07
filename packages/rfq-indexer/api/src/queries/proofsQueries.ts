import { db } from '../db'

// typical fields to return for a BridgeProofProvided event when it is joined to a BridgeRequest
export const qProofs = ({ activeOnly}: { activeOnly: boolean} = { activeOnly: false}) => {
  let query = db
    .selectFrom('BridgeProofProvidedEvents')
    .leftJoin('BridgeProofDisputedEvents', (join) =>
      join
        .onRef('BridgeProofDisputedEvents.transactionId', '=', 'BridgeProofProvidedEvents.transactionId')
        .onRef('BridgeProofDisputedEvents.blockTimestamp', '>', 'BridgeProofProvidedEvents.blockTimestamp')
    )
    .select([
      'BridgeProofProvidedEvents.transactionId as transactionId_proof',
      'BridgeProofProvidedEvents.blockNumber as blockNumber_proof',
      'BridgeProofProvidedEvents.blockTimestamp as blockTimestamp_proof',
      'BridgeProofProvidedEvents.transactionHash as transactionHash_proof',
      'BridgeProofProvidedEvents.relayer as relayer_proof',
    ]);

  if (activeOnly) {
    query = query.where('BridgeProofDisputedEvents.transactionId', 'is', null);
  }

  return query;
}