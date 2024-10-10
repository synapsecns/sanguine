import { Request, Response } from 'express'

import { db } from '../db'
import { qDeposits, qRelays, qProofs, qClaims, qRefunds, qDisputes } from '../queries'
import { nest_results } from '../utils/nestResults'

export const getTransactionById = async (req: Request, res: Response) => {
  const { transactionId } = req.params

  try {
    const query = db
      .with('deposits', () =>
        qDeposits().where('transactionId', '=', transactionId as string)
      )
      .with('relays', () => qRelays())
      .with('proofs', () => qProofs({ activeOnly: false })) // display proofs even if they have been invalidated/replaced by a dispute
      .with('disputes', () => qDisputes({ activeOnly: true })) // do not show disputes that have been invalidated/replaced by a proof
      .with('claims', () => qClaims())
      .with('refunds', () => qRefunds())
      .with('combined', (qb) =>
        qb
          .selectFrom('deposits')
          .leftJoin('relays', 'transactionId_deposit', 'transactionId_relay')
          .leftJoin('proofs', 'transactionId_deposit', 'transactionId_proof')
          .leftJoin('disputes', 'transactionId_deposit', 'transactionId_dispute')
          .leftJoin('claims', 'transactionId_deposit', 'transactionId_claim')
          .leftJoin('refunds', 'transactionId_deposit', 'transactionId_refund')
          .selectAll('deposits')
          .selectAll('relays')
          .selectAll('proofs')
          .selectAll('disputes')
          .selectAll('claims')
          .selectAll('refunds')
      )
      .selectFrom('combined')
      .selectAll()

    const results = await query.execute()
    const nestedResult = nest_results(results)[0] || null

    if (nestedResult) {
      const filteredResult = Object.fromEntries(
        Object.entries(nestedResult).filter(([_, value]) => {
          if (value === null) {
            return false
          }
          if (typeof value !== 'object') {
            return true
          }
          return Object.values(value).some((v) => v !== null)
        })
      )
      res.json(filteredResult)
    } else {
      res.status(200).json({ message: 'Transaction not found' })
    }
  } catch (error) {
    console.error(error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
