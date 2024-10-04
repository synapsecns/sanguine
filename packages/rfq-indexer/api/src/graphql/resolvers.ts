import { sql } from 'kysely'

import { db } from '../db'

// typical fields to return for a BridgeRequest
const qDeposits = () => {
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
      'BridgeRequestEvents.sendChainGas',
    ])
    .where('BridgeRequestEvents.blockTimestamp', '>', 1722729600)
  // if index is partially loaded, we must limit lookback or will have various data issues from relays
  // that happened to be in flight at the point of the index's start.
  // may also improve query performance
}

// typical fields to return for a BridgeRelayed event when it is joined to a BridgeRequest
const qRelays = () => {
  return db
    .selectFrom('BridgeRelayedEvents')
    .select([
      'BridgeRelayedEvents.transactionId as transactionId_relay',
      'BridgeRelayedEvents.blockNumber as blockNumber_relay',
      'BridgeRelayedEvents.blockTimestamp as blockTimestamp_relay',
      'BridgeRelayedEvents.transactionHash as transactionHash_relay',

      'BridgeRelayedEvents.relayer as relayer_relay',
      'BridgeRelayedEvents.to as to_relay',
    ])
}

// typical fields to return for a BridgeProofProvided event when it is joined to a BridgeRequest
const qProofs = () => {
  return db
    .selectFrom('BridgeProofProvidedEvents')
    .select([
      'BridgeProofProvidedEvents.transactionId as transactionId_proof',
      'BridgeProofProvidedEvents.blockNumber as blockNumber_proof',
      'BridgeProofProvidedEvents.blockTimestamp as blockTimestamp_proof',
      'BridgeProofProvidedEvents.transactionHash as transactionHash_proof',

      'BridgeProofProvidedEvents.relayer as relayer_proof',
    ])
}

// typical fields to return for a BridgeDepositClaimed event when it is joined to a BridgeRequest
const qClaims = () => {
  return db
    .selectFrom('BridgeDepositClaimedEvents')
    .select([
      'BridgeDepositClaimedEvents.transactionId as transactionId_claim',
      'BridgeDepositClaimedEvents.blockNumber as blockNumber_claim',
      'BridgeDepositClaimedEvents.blockTimestamp as blockTimestamp_claim',
      'BridgeDepositClaimedEvents.transactionHash as transactionHash_claim',

      'BridgeDepositClaimedEvents.to as to_claim',
      'BridgeDepositClaimedEvents.relayer as relayer_claim',
      'BridgeDepositClaimedEvents.amountFormatted as amountFormatted_claim',
    ])
}

// typical fields to return for a BridgeDepositRefunded event when it is joined to a BridgeRequest
const qRefunds = () => {
  return db
    .selectFrom('BridgeDepositRefundedEvents')
    .select([
      'BridgeDepositRefundedEvents.transactionId as transactionId_refund',
      'BridgeDepositRefundedEvents.blockNumber as blockNumber_refund',
      'BridgeDepositRefundedEvents.blockTimestamp as blockTimestamp_refund',
      'BridgeDepositRefundedEvents.transactionHash as transactionHash_refund',

      'BridgeDepositRefundedEvents.to as to_refund',
      'BridgeDepositRefundedEvents.amountFormatted as amountFormatted_refund',
    ])
}

// typical fields to return for a BridgeProofDisputed event when it is joined to a BridgeRequest
const qDisputes = () => {
  return db
    .selectFrom('BridgeProofDisputedEvents')
    .select([
      'BridgeProofDisputedEvents.transactionId as transactionId_dispute',
      'BridgeProofDisputedEvents.blockNumber as blockNumber_dispute',
      'BridgeProofDisputedEvents.blockTimestamp as blockTimestamp_dispute',
      'BridgeProofDisputedEvents.transactionHash as transactionHash_dispute',
      'BridgeProofDisputedEvents.chainId as originChainId_dispute',
      'BridgeProofDisputedEvents.chain as originChain_dispute',
    ])
}

// using the suffix of a field, move it into a nested sub-object. This is a cleaner final resultset
// example: transactionHash_deposit:0xyz would get moved into BridgeRequest{transactionHash:0xyz}
//
// also note that transactionId_xxxx will just convert into a single "transactionId"
const nest_results = (sqlResults: any[]) => {
  return sqlResults.map((transaction: any) => {
    const bridgeRequest: { [key: string]: any } = {}
    const bridgeRelay: { [key: string]: any } = {}
    const bridgeProof: { [key: string]: any } = {}
    const bridgeClaim: { [key: string]: any } = {}
    const bridgeRefund: { [key: string]: any } = {}
    const bridgeDispute: { [key: string]: any } = {}
    const transactionFields: { [key: string]: any } = {}

    let transactionIdSet = false

    for (const [key, value] of Object.entries(transaction)) {
      if (key.startsWith('transactionId')) {
        if (!transactionIdSet) {
          transactionFields[key.replace(/_.+$/, '')] = value
          transactionIdSet = true
        }
        // Ignore other transactionId fields
      } else if (key.endsWith('_deposit')) {
        bridgeRequest[key.replace('_deposit', '')] = value
      } else if (key.endsWith('_relay')) {
        bridgeRelay[key.replace('_relay', '')] = value
      } else if (key.endsWith('_proof')) {
        bridgeProof[key.replace('_proof', '')] = value
      } else if (key.endsWith('_claim')) {
        bridgeClaim[key.replace('_claim', '')] = value
      } else if (key.endsWith('_refund')) {
        bridgeRefund[key.replace('_refund', '')] = value
      } else if (key.endsWith('_dispute')) {
        bridgeDispute[key.replace('_dispute', '')] = value
      } else {
        transactionFields[key] = value
      }
    }

    const result: { [key: string]: any } = { Bridge: transactionFields }
    if (Object.keys(bridgeRequest).length) {
      result.BridgeRequest = bridgeRequest
    }
    if (Object.keys(bridgeRelay).length) {
      result.BridgeRelay = bridgeRelay
    }
    if (Object.keys(bridgeProof).length) {
      result.BridgeProof = bridgeProof
    }
    if (Object.keys(bridgeClaim).length) {
      result.BridgeClaim = bridgeClaim
    }
    if (Object.keys(bridgeRefund).length) {
      result.BridgeRefund = bridgeRefund
    }
    if (Object.keys(bridgeDispute).length) {
      result.BridgeDispute = bridgeDispute
    }
    return result
  })
}

const resolvers = {
  Query: {
    events: async (
      _: any,
      { first = 10, after, filter }: { first?: any; after?: any; filter?: any }
    ) => {
      let query = db
        .selectFrom('BridgeRequestEvents')
        .select([
          'BridgeRequestEvents.id',
          'BridgeRequestEvents.transactionId',
          'BridgeRequestEvents.blockNumber',
          'BridgeRequestEvents.blockTimestamp',
          'BridgeRequestEvents.transactionHash',
          'BridgeRequestEvents.originChainId',
          'BridgeRequestEvents.originChain',
        ])
        .unionAll(
          db
            .selectFrom('BridgeRelayedEvents')
            .select([
              'BridgeRelayedEvents.id',
              'BridgeRelayedEvents.transactionId',
              'BridgeRelayedEvents.blockNumber',
              'BridgeRelayedEvents.blockTimestamp',
              'BridgeRelayedEvents.transactionHash',
              'BridgeRelayedEvents.originChainId',
              'BridgeRelayedEvents.originChain',
            ])
        )
        .unionAll(
          db
            .selectFrom('BridgeProofProvidedEvents')
            .select([
              'BridgeProofProvidedEvents.id',
              'BridgeProofProvidedEvents.transactionId',
              'BridgeProofProvidedEvents.blockNumber',
              'BridgeProofProvidedEvents.blockTimestamp',
              'BridgeProofProvidedEvents.transactionHash',
              'BridgeProofProvidedEvents.originChainId',
              'BridgeProofProvidedEvents.originChain',
            ])
        )
        .unionAll(
          db
            .selectFrom('BridgeDepositRefundedEvents')
            .select([
              'BridgeDepositRefundedEvents.id',
              'BridgeDepositRefundedEvents.transactionId',
              'BridgeDepositRefundedEvents.blockNumber',
              'BridgeDepositRefundedEvents.blockTimestamp',
              'BridgeDepositRefundedEvents.transactionHash',
              'BridgeDepositRefundedEvents.originChainId',
              'BridgeDepositRefundedEvents.originChain',
            ])
        )
        .unionAll(
          db
            .selectFrom('BridgeDepositClaimedEvents')
            .select([
              'BridgeDepositClaimedEvents.id',
              'BridgeDepositClaimedEvents.transactionId',
              'BridgeDepositClaimedEvents.blockNumber',
              'BridgeDepositClaimedEvents.blockTimestamp',
              'BridgeDepositClaimedEvents.transactionHash',
              'BridgeDepositClaimedEvents.originChainId',
              'BridgeDepositClaimedEvents.originChain',
            ])
        )
        .unionAll(
          db
            .selectFrom('BridgeProofDisputedEvents')
            .select([
              'BridgeProofDisputedEvents.id',
              'BridgeProofDisputedEvents.transactionId',
              'BridgeProofDisputedEvents.blockNumber',
              'BridgeProofDisputedEvents.blockTimestamp',
              'BridgeProofDisputedEvents.transactionHash',
              'BridgeProofDisputedEvents.chainId',
              'BridgeProofDisputedEvents.chain',
            ])
        )

      if (filter) {
        if (filter.transactionId) {
          query = query.where('transactionId', '=', filter.transactionId)
        }
        if (filter.originChainId) {
          query = query.where('originChainId', '=', filter.originChainId)
        }
        // Add more filters as needed
      }

      if (after) {
        query = query.where('id', '>', after)
      }

      const events = await query
        .orderBy('blockTimestamp', 'desc')
        .limit(first + 1)
        .execute()

      const hasNextPage = events.length > first
      const edges = events.slice(0, first).map((event: any) => ({
        node: event,
        cursor: event.id,
      }))

      return {
        edges,
        pageInfo: {
          hasNextPage,
          endCursor: edges.length > 0 ? edges[edges.length - 1]?.cursor : null,
        },
      }
    },
    pendingTransactionsMissingRelay: async () => {
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
              .leftJoin(
                'relays',
                'transactionId_deposit',
                'transactionId_relay'
              )
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

      return nest_results(await query.execute())
    },
    pendingTransactionsMissingProof: async () => {
      const query = db
        .with('deposits', () => qDeposits())
        .with('relays', () => qRelays())
        .with('proofs', () => qProofs())
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

      return nest_results(await query.execute())
    },
    pendingTransactionsMissingClaim: async () => {
      const query = db
        .with('deposits', () => qDeposits())
        .with('relays', () => qRelays())
        .with('proofs', () => qProofs())
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

      return nest_results(await query.execute())
    },
    recentInvalidRelays: async () => {
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

      // intentionally do not nest - doesnt make sense w/ this dataset because the whole point is that no Deposit exists
      return query.execute()
    },
    refundedAndRelayedTransactions: async () => {
      const query = db
        .with('deposits', () => qDeposits())
        .with('relays', () => qRelays())
        .with('refunds', () => qRefunds())
        .with('combined', (qb) =>
          qb
            .selectFrom('deposits')
            .innerJoin('relays', 'transactionId_deposit', 'transactionId_relay')
            .innerJoin(
              'refunds',
              'transactionId_deposit',
              'transactionId_refund'
            )
            .selectAll('deposits')
            .selectAll('relays')
            .selectAll('refunds')
        )
        .selectFrom('combined')
        .selectAll()
        .orderBy('blockTimestamp_refund', 'desc')

      return nest_results(await query.execute())
    },
    transactionById: async (
      _: any,
      { transactionId }: { transactionId: string }
    ) => {
      const query = db
        .with('deposits', () =>
          qDeposits().where('transactionId', '=', transactionId)
        )
        .with('relays', () => qRelays())
        .with('proofs', () => qProofs())
        .with('claims', () => qClaims())
        .with('refunds', () => qRefunds())
        .with('combined', (qb) =>
          qb
            .selectFrom('deposits')
            .leftJoin('relays', 'transactionId_deposit', 'transactionId_relay')
            .leftJoin('proofs', 'transactionId_deposit', 'transactionId_proof')
            .leftJoin('claims', 'transactionId_deposit', 'transactionId_claim')
            .leftJoin(
              'refunds',
              'transactionId_deposit',
              'transactionId_refund'
            )
            .selectAll('deposits')
            .selectAll('relays')
            .selectAll('proofs')
            .selectAll('claims')
            .selectAll('refunds')
        )
        .selectFrom('combined')
        .selectAll()

      const nestedResult = nest_results(await query.execute())[0] || null

      if (nestedResult) {
        return Object.fromEntries(
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
      }

      return null
    },
    conflictingProofs: async () => {
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

      return nest_results(await query.execute())
    },
    disputedRelays: async () => {
      const query = db
        .with('deposits', () => qDeposits())
        .with('relays', () => qRelays())
        .with('proofs', () => qProofs())
        .with('disputes', () => qDisputes())
        .with('combined', (qb) =>
          qb
            .selectFrom('proofs')
            .leftJoin(
              'disputes',
              'transactionId_proof',
              'transactionId_dispute'
            )
            .selectAll('proofs')
            .selectAll('disputes')
        )
        .selectFrom('combined')
        .selectAll()
        .orderBy('blockTimestamp_proof', 'desc')

      return nest_results(await query.execute())
    },
  },
  BridgeEvent: {
    // eslint-disable-next-line prefer-arrow/prefer-arrow-functions
    __resolveType(obj: any) {
      // Implement logic to determine the event type based on the object properties
      // For example:
      if ('sender' in obj) {
        return 'BridgeRequestEvent'
      }
      if ('relayer' in obj && 'to' in obj) {
        return 'BridgeRelayedEvent'
      }
      if ('relayer' in obj && !('to' in obj)) {
        return 'BridgeProofProvidedEvent'
      }
      if ('to' in obj && 'token' in obj) {
        return 'BridgeDepositRefundedEvent'
      }
      if ('relayer' in obj && 'to' in obj && 'token' in obj) {
        return 'BridgeDepositClaimedEvent'
      }
      return null
    },
  },
}

export { resolvers }
