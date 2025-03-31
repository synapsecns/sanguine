import { Request, Response } from 'express'

import { jsonToHtmlTable } from '../utils/json_formatter'
import { db } from '../db'
import { addSenderStatus, addTokenSymbols, addUsdPricesCurrent } from '../utils/enrichResults';

export const recentInvalidRelaysController = async (
  req: Request,
  res: Response
) => {
  const flags = req.query.flags as string | undefined;
  const format = req.query.format as string | undefined;
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

    await addTokenSymbols(results);
  
    if (flags?.includes("synapse")) {
      await addSenderStatus(results);
      await addUsdPricesCurrent(results);
    }

    if (results && results.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(results));
      } else {
        res.json(results);
      }
    } else {
      res.status(200).json({ message: 'No recent invalid relays found' })
    }
  } catch (error) {
    console.error('Error fetching recent invalid relays:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
