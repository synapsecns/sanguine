import { Request, Response } from 'express'
import { sql } from 'kysely'

import { jsonToHtmlTable } from '../utils/json_formatter'
import { db } from '../db'
import { qDeposits, qRelays, qProofs } from '../queries'
import { nest_results } from '../utils/nestResults'
import { addSenderStatus, addTokenSymbols, addUsdPricesCurrent } from '../utils/enrichResults'

export const conflictingProofsController = async (
  req: Request,
  res: Response
) => {
  const flags = req.query.flags as string | undefined;
  const format = req.query.format as string | undefined;
  try {
    const query = db
      .with('deposits', () => qDeposits())
      .with('relays', () => qRelays())
      .with('proofs', () => qProofs({ activeOnly: true }))
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

    await addTokenSymbols(results);
  
    if (flags?.includes("synapse")) {
      await addSenderStatus(results);
      await addUsdPricesCurrent(results);
    }
    
    const conflictingProofs = nest_results(results)

    if (conflictingProofs && conflictingProofs.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(conflictingProofs));
      } else {
        res.json(conflictingProofs);
      }
    } else {
      res.status(200).json({ message: 'No active conflicting proofs found' })
    }
  } catch (error) {
    console.error('Error fetching active conflicting proofs:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
