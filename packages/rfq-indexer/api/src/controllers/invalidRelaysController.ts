import { Request, Response } from 'express'

import { db } from '../db'

export const recentInvalidRelaysController = async (
  req: Request,
  res: Response
) => {
  try {
    const query = db
      .selectFrom('BridgeRelayedEvents')
      .leftJoin(
        'BridgeRequestEvents',
        'BridgeRelayedEvents.transactionId',
        'BridgeRequestEvents.transactionId'
      )
      .select([
        'BridgeRelayedEvents.transactionId',
        'BridgeRelayedEvents.blockNumber',
        'BridgeRelayedEvents.blockTimestamp',
        'BridgeRelayedEvents.transactionHash',
        'BridgeRelayedEvents.originChain',
        'BridgeRelayedEvents.destChain',
        'BridgeRelayedEvents.originChainId',
        'BridgeRelayedEvents.destChainId',
        'BridgeRelayedEvents.originToken',
        'BridgeRelayedEvents.destToken',
        'BridgeRelayedEvents.originAmountFormatted',
        'BridgeRelayedEvents.destAmountFormatted',
        'BridgeRelayedEvents.to',
        'BridgeRelayedEvents.relayer',
      ])
      // lookback approx 2 weeks
      .where(
        'BridgeRelayedEvents.blockTimestamp',
        '>',
        Math.floor(Date.now() / 1000) - 2 * 7 * 24 * 60 * 60
      )
      .where('BridgeRequestEvents.transactionId', 'is', null)

    const results = await query.execute()

    if (results && results.length > 0) {
      res.json(results)
    } else {
      res.status(200).json({ message: 'No recent invalid relays found' })
    }
  } catch (error) {
    console.error('Error fetching recent invalid relays:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
