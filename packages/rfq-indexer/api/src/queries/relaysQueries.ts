import { db } from '../db'

// typical fields to return for a BridgeRelayed event when it is joined to a BridgeRequest
export const qRelays = () => {
  return db
    .selectFrom('BridgeRelayedEvents')
    .select([
      'BridgeRelayedEvents.transactionId as transactionId_relay',
      'BridgeRelayedEvents.blockNumber as blockNumber_relay',
      'BridgeRelayedEvents.blockTimestamp as blockTimestamp_relay',
      'BridgeRelayedEvents.transactionHash as transactionHash_relay',

      'BridgeRelayedEvents.relayer as relayer_relay',
      'BridgeRelayedEvents.to as to_relay',
    ])
}
