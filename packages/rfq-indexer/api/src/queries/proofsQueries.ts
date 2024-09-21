import { db } from '../db'

// typical fields to return for a BridgeProofProvided event when it is joined to a BridgeRequest
export const qProofs = () => {
  return db
    .selectFrom('BridgeProofProvidedEvents')
    .select([
      'BridgeProofProvidedEvents.transactionId as transactionId_proof',
      'BridgeProofProvidedEvents.blockNumber as blockNumber_proof',
      'BridgeProofProvidedEvents.blockTimestamp as blockTimestamp_proof',
      'BridgeProofProvidedEvents.transactionHash as transactionHash_proof',

      'BridgeProofProvidedEvents.relayer as relayer_proof',
    ])
}
