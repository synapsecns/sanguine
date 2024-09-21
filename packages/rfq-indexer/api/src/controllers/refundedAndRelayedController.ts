import { Request, Response } from 'express'

import { db } from '../db'
import { qDeposits, qRelays, qRefunds } from '../queries'
import { nest_results } from '../utils/nestResulsts'

export const refundedAndRelayedTransactionsController = async (
  req: Request,
  res: Response
) => {
  try {
    const query = db
      .with('deposits', () => qDeposits())
      .with('relays', () => qRelays())
      .with('refunds', () => qRefunds())
      .with('combined', (qb) =>
        qb
          .selectFrom('deposits')
          .innerJoin('relays', 'transactionId_deposit', 'transactionId_relay')
          .innerJoin('refunds', 'transactionId_deposit', 'transactionId_refund')
          .selectAll('deposits')
          .selectAll('relays')
          .selectAll('refunds')
      )
      .selectFrom('combined')
      .selectAll()
      .orderBy('blockTimestamp_refund', 'desc')

    const results = await query.execute()
    const nestedResults = nest_results(results)

    if (nestedResults && nestedResults.length > 0) {
      res.json(nestedResults)
    } else {
      res
        .status(200)
        .json({ message: 'No refunded and relayed transactions found' })
    }
  } catch (error) {
    console.error('Error fetching refunded and relayed transactions:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
