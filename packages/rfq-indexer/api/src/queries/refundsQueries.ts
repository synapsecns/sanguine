import { db } from '../db'

export const qRefunds = () => {
  return db
    .selectFrom('BridgeDepositRefundedEvents')
    .select([
      'BridgeDepositRefundedEvents.transactionId as transactionId_refund',
      'BridgeDepositRefundedEvents.blockNumber as blockNumber_refund',
      'BridgeDepositRefundedEvents.blockTimestamp as blockTimestamp_refund',
      'BridgeDepositRefundedEvents.transactionHash as transactionHash_refund',

      'BridgeDepositRefundedEvents.to as to_refund',
      'BridgeDepositRefundedEvents.amountFormatted as amountFormatted_refund',
    ])
}
