import { Request, Response } from 'express'
import { sql } from 'kysely'

import { db } from '../db'
import { qDeposits, qRelays, qProofs } from '../queries'
import { nest_results } from '../utils/nestResults'

export const conflictingProofsController = async (
  req: Request,
  res: Response
) => {
  try {
    const query = db
      .with('deposits', () => qDeposits())
      .with('relays', () => qRelays())
      .with('proofs', () => qProofs())
      .with('combined', (qb) =>
        qb
          .selectFrom('deposits')
          .leftJoin('relays', 'transactionId_deposit', 'transactionId_relay')
          .leftJoin('proofs', 'transactionId_deposit', 'transactionId_proof')
          .selectAll('deposits')
          .selectAll('relays')
          .selectAll('proofs')
      )
      .selectFrom('combined')
      .selectAll()
      .where('relayer_proof', 'is not', null)
      .where('relayer_relay', 'is not', null)
      .where(
        (eb) =>
          sql<boolean>`LOWER(${eb.ref('relayer_relay')}) != LOWER(${eb.ref(
            'relayer_proof'
          )})`
      )
      .orderBy('blockTimestamp_proof', 'desc')

    const results = await query.execute()
    const conflictingProofs = nest_results(results)

    if (conflictingProofs && conflictingProofs.length > 0) {
      res.json(conflictingProofs)
    } else {
      res.status(200).json({ message: 'No conflicting proofs found' })
    }
  } catch (error) {
    console.error('Error fetching conflicting proofs:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
