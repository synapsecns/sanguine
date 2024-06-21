import { createConfig } from '@ponder/core'
import { http } from 'viem'

import { messageBusAbi } from './abis/messageBusAbi'

type MessageBusChain = 'dfk' | 'harmony' | 'klaytn'
// Different start dates for testing purposes:
// - deploy: block number when the contract was deployed
// - '2024-01-01': to be used for indexing the 2024 messages
// - '2024-04-01': to be used for initial testing of the 2024 messages
const startDates = ['deploy', '2024-01-01', '2024-04-01'] as const
type StartDateType = (typeof startDates)[number]

// Fetch the start date from the environment variables, use the latest date as fallback
const startDate: StartDateType = startDates.includes(
  process.env.START_DATE as StartDateType
)
  ? (process.env.START_DATE as StartDateType)
  : '2024-04-01'

console.log(`Using start date: ${startDate}`)

const startBlock: Record<MessageBusChain, Record<StartDateType, number>> = {
  dfk: {
    deploy: 1513462,
    '2024-01-01': 26749099,
    '2024-04-01': 30461868,
  },
  harmony: {
    deploy: 26108972,
    '2024-01-01': 51743502,
    '2024-04-01': 55625555,
  },
  klaytn: {
    deploy: 105416660,
    '2024-01-01': 142088644,
    '2024-04-01': 149946308,
  },
}

export default createConfig({
  networks: {
    dfk: {
      chainId: 53935,
      transport: http(process.env.PONDER_RPC_URL_53935),
    },
    harmony: {
      chainId: 1666600000,
      transport: http(process.env.PONDER_RPC_URL_1666600000),
    },
    klaytn: {
      chainId: 8217,
      transport: http(process.env.PONDER_RPC_URL_8217),
      maxRequestsPerSecond: 20,
    },
  },
  contracts: {
    messageBus: {
      abi: messageBusAbi,
      network: {
        dfk: {
          address: '0x7bc5fD6b80067d6052A4550c69f152877bF7C748',
          startBlock: startBlock.dfk[startDate],
          maxBlockRange: 2048,
        },
        harmony: {
          address: '0x4F437be4A3448fCf394e513FC1A8EF92E190D1ba',
          startBlock: startBlock.harmony[startDate],
          maxBlockRange: 1024,
        },
        klaytn: {
          address: '0xaEe80e4B92Ba497aF1378Bc799687FBF816Ab87b',
          startBlock: startBlock.klaytn[startDate],
          maxBlockRange: 5000,
        },
      },
    },
  },
})
