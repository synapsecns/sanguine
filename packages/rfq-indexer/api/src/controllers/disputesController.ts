import { Request, Response } from 'express'
import { jsonToHtmlTable } from '../utils/json_formatter'

import { db } from '../db'
import { qDisputes } from '../queries'
import { nest_results } from '../utils/nestResults'
import { addSenderStatus, addTokenSymbols, addUsdPricesCurrent } from '../utils/enrichResults'

export const disputesController = async (req: Request, res: Response) => {
  const flags = req.query.flags as string | undefined;
  const format = req.query.format as string | undefined;
  try {
    const query = db
      .with('disputes', () => qDisputes({ activeOnly: true }))
      .selectFrom('disputes')
      .selectAll()
      .orderBy('blockTimestamp_dispute', 'desc')

    const results = await query.execute()

    await addTokenSymbols(results);
  
    if (flags?.includes("synapse")) {
      await addSenderStatus(results);
      await addUsdPricesCurrent(results);
    }
    
    
    const disputes = nest_results(results)

    if (disputes && disputes.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(disputes));
      } else {
        res.json(disputes);
      }
    } else {
      res.status(200).json({ message: 'No active disputes found' })
    }
  } catch (error) {
    console.error('Error fetching active disputes:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
