import { Request, Response } from 'express'

import { db } from '../db'
import { qDisputes } from '../queries'
import { nest_results } from '../utils/nestResults'

export const disputesController = async (req: Request, res: Response) => {
  try {
    const query = db
      .with('disputes', () => qDisputes({activeOnly: true}))
      .selectFrom('disputes')
      .selectAll()
      .orderBy('blockTimestamp_dispute', 'desc')

    const results = await query.execute()
    const disputes = nest_results(results)

    if (disputes && disputes.length > 0) {
      res.json(disputes)
    } else {
      res.status(200).json({ message: 'No active disputes found' })
    }
  } catch (error) {
    console.error('Error fetching active disputes:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
