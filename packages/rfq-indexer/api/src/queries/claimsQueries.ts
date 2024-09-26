import { db } from '../db'

// typical fields to return for a BridgeDepositClaimed event when it is joined to a BridgeRequest
export const qClaims = () => {
  return db
    .selectFrom('BridgeDepositClaimedEvents')
    .select([
      'BridgeDepositClaimedEvents.transactionId as transactionId_claim',
      'BridgeDepositClaimedEvents.blockNumber as blockNumber_claim',
      'BridgeDepositClaimedEvents.blockTimestamp as blockTimestamp_claim',
      'BridgeDepositClaimedEvents.transactionHash as transactionHash_claim',

      'BridgeDepositClaimedEvents.to as to_claim',
      'BridgeDepositClaimedEvents.relayer as relayer_claim',
      'BridgeDepositClaimedEvents.amountFormatted as amountFormatted_claim',
    ])
}
