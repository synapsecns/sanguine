import { db } from '../db'

export const qDeposits = () => {
  return db
    .selectFrom('BridgeRequestEvents')
    .select([
      'BridgeRequestEvents.transactionId as transactionId_deposit',
      'BridgeRequestEvents.blockNumber as blockNumber_deposit',
      'BridgeRequestEvents.blockTimestamp as blockTimestamp_deposit',
      'BridgeRequestEvents.transactionHash as transactionHash_deposit',
      'BridgeRequestEvents.originChain',
      'BridgeRequestEvents.destChain',
      'BridgeRequestEvents.originChainId',
      'BridgeRequestEvents.destChainId',
      'BridgeRequestEvents.originToken',
      'BridgeRequestEvents.destToken',
      'BridgeRequestEvents.originAmountFormatted',
      'BridgeRequestEvents.destAmountFormatted',
      'BridgeRequestEvents.sender',
      'BridgeRequestEvents.request',
      'BridgeRequestEvents.sendChainGas',
    ])
    .where('BridgeRequestEvents.blockTimestamp', '>', 1722729600)
  // if index is partially loaded, we must limit lookback or will have various data issues from relays
  // that happened to be in flight at the point of the index's start.
  // may also improve query performance
}
