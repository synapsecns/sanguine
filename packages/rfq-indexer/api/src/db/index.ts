import { Kysely, PostgresDialect } from 'kysely'
import { Pool } from 'pg'

import type {
  BridgeRequestEvents,
  BridgeRelayedEvents,
  BridgeProofProvidedEvents,
  BridgeDepositRefundedEvents,
  BridgeDepositClaimedEvents,
  BridgeProofDisputedEvents,
} from '../types'

const { DATABASE_URL } = process.env

const pool = new Pool({ connectionString: DATABASE_URL })

const dialect = new PostgresDialect({ pool })

export interface Database {
  BridgeRequestEvents: BridgeRequestEvents
  BridgeRelayedEvents: BridgeRelayedEvents
  BridgeProofProvidedEvents: BridgeProofProvidedEvents
  BridgeDepositRefundedEvents: BridgeDepositRefundedEvents
  BridgeDepositClaimedEvents: BridgeDepositClaimedEvents
  BridgeProofDisputedEvents: BridgeProofDisputedEvents
}

export const db = new Kysely<Database>({ dialect })
