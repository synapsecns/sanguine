import { Request, response, Response } from 'express'

import { jsonToHtmlTable } from '../utils/json_formatter'
import { addSenderStatus, addTokenSymbols, addUsdPricesCurrent } from '../utils/enrichResults'

import { db } from '../db'
import { qDeposits, qRelays, qProofs, qClaims, qRefunds } from '../queries'
import { nest_results } from '../utils/nestResults'




const sevenDaysAgo = Math.floor(Date.now() / 1000) - 7 * 24 * 60 * 60

export const pendingTransactionsMissingClaimController = async (
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
      .with('claims', () => qClaims())
      .with('combined', (qb) =>
        qb
          .selectFrom('deposits')
          .innerJoin('relays', 'transactionId_deposit', 'transactionId_relay')
          .innerJoin('proofs', 'transactionId_deposit', 'transactionId_proof')
          .leftJoin('claims', 'transactionId_deposit', 'transactionId_claim')
          .selectAll('deposits')
          .selectAll('relays')
          .selectAll('proofs')
          .where('transactionId_claim', 'is', null)
      )
      .selectFrom('combined')
      .selectAll()
      .orderBy('blockTimestamp_proof', 'desc')

    const results = await query.execute()

    await addTokenSymbols(results);
  
    if (flags?.includes("synapse")) {
      await addSenderStatus(results);
      await addUsdPricesCurrent(results);
    }
    
    
    const nestedResults = nest_results(results)

    if (nestedResults && nestedResults.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(nestedResults, flags));
      } else {
        res.json(nestedResults);
      }
    } else {
      res
        .status(200)
        .json({ message: 'No pending transactions missing claim found' })
    }
  } catch (error) {
    console.error('Error fetching pending transactions missing claim:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}

export const pendingTransactionsMissingProofController = async (
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
          .innerJoin('relays', 'transactionId_deposit', 'transactionId_relay')
          .leftJoin('proofs', 'transactionId_deposit', 'transactionId_proof')
          .selectAll('deposits')
          .selectAll('relays')
          .where('transactionId_proof', 'is', null)
      )
      .selectFrom('combined')
      .selectAll()
      .orderBy('blockTimestamp_relay', 'desc')

    const results = await query.execute()

    await addTokenSymbols(results);
  
    if (flags?.includes("synapse")) {
      await addSenderStatus(results);
      await addUsdPricesCurrent(results);
    }
    
    
    const nestedResults = nest_results(results)

    if (nestedResults && nestedResults.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(nestedResults, flags));
      } else {
        res.json(nestedResults);
      }
    } else {
      res
        .status(200)
        .json({ message: 'No pending transactions missing proof found' })
    }
  } catch (error) {
    console.error('Error fetching pending transactions missing proof:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}

export const pendingTransactionsMissingRelayController = async (
  req: Request,
  res: Response
) => {
  const flags = req.query.flags as string | undefined;
  const format = req.query.format as string | undefined;
  try {
    const query = db
      .with('deposits', () => qDeposits())
      .with('relays', () => qRelays())
      .with('refunds', () => qRefunds())
      .with(
        'combined',
        (qb) =>
          qb
            .selectFrom('deposits')
            .selectAll('deposits')
            .leftJoin('relays', 'transactionId_deposit', 'transactionId_relay')
            .leftJoin(
              'refunds',
              'transactionId_deposit',
              'transactionId_refund'
            )
            .where('transactionId_relay', 'is', null) // is not relayed
            .where('transactionId_refund', 'is', null) // is not refunded
      )
      .selectFrom('combined')
      .selectAll()
      .orderBy('blockTimestamp_deposit', 'desc')
      .where('blockTimestamp_deposit', '>', sevenDaysAgo)

      const results = await query.execute()

      await addTokenSymbols(results);
    
      if (flags?.includes("synapse")) {
        await addSenderStatus(results);
        await addUsdPricesCurrent(results);
      }

    const nestedResults = nest_results(results)

    if (nestedResults && nestedResults.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(nestedResults, flags));
      } else {
        res.json(nestedResults);
      }
    } else {
      res
        .status(200)
        .json({ message: 'No pending transactions missing relay found' })
    }
  } catch (error) {
    console.error('Error fetching pending transactions missing relay:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}

export const pendingTransactionsMissingRelayExceedDeadlineController = async (
  req: Request,
  res: Response
) => {
  const flags = req.query.flags as string | undefined;
  const format = req.query.format as string | undefined;
  try {
    const query = db
      .with('deposits', () => qDeposits())
      .with('relays', () => qRelays())
      .with('refunds', () => qRefunds())
      .with(
        'combined',
        (qb) =>
          qb
            .selectFrom('deposits')
            .selectAll('deposits')
            .leftJoin('relays', 'transactionId_deposit', 'transactionId_relay')
            .leftJoin(
              'refunds',
              'transactionId_deposit',
              'transactionId_refund'
            )
            .where('transactionId_relay', 'is', null) // is not relayed
            .where('transactionId_refund', 'is', null) // is not refunded
      )
      .selectFrom('combined')
      .selectAll()
      .orderBy('blockTimestamp_deposit', 'desc')
      .where('blockTimestamp_deposit', '<=', sevenDaysAgo)

    const results = await query.execute()

    await addTokenSymbols(results);
  
    if (flags?.includes("synapse")) {
      await addSenderStatus(results);
      await addUsdPricesCurrent(results);
    }

    const nestedResults = nest_results(results)

    if (nestedResults && nestedResults.length > 0) {
      if (format === 'html') {
        res.send(jsonToHtmlTable(nestedResults, flags));
      } else {
        res.json(nestedResults);
      }
    } else {
      res
        .status(200)
        .json({ message: 'No pending transactions missing relay found' })
    }
  } catch (error) {
    console.error('Error fetching pending transactions missing relay:', error)
    res.status(500).json({ message: 'Internal server error' })
  }
}
