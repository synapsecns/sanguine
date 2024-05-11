import { createSchema } from 'graphql-yoga'
import { PrismaClient } from '@prisma/client'
import {
  CountsFilter,
  type InterchainBatch,
  type InterchainBatchQueryFilter,
  type InterchainTransaction,
  type InterchainTransactionQueryFilter,
} from '@/types'
import { InterchainClientV1Abi } from '@/abis/InterchainClientV1Abi'
import { getTransactionStatusInfo } from '@/utils/getTransactionStatusInfo'
import { networkConfig } from '@/networkConfig'

const prisma = new PrismaClient()

const typeDefs = `
  scalar BigInt

  type AppConfigV1 {
    id: String!
    requiredResponses: Int!
    optimisticPeriod: Int!
    modules: [String!]!
  }

  type InterchainBatch {
    id: String!
    batchRoot: String!
    srcDbNonce: BigInt!
    srcChainId: Int
    dstDbNonce: BigInt!
    dstChainId: Int
    status: String!
    verifiedAt: BigInt
    appConfigV1Id: String
    appConfigV1: AppConfigV1
    interchainTransactions: [InterchainTransaction!]!
  }

  type InterchainTransactionSent {
    id: String!
    srcChainId: Int!
    name: String!
    transactionId: String!
    dbNonce: BigInt!
    entryIndex: BigInt!
    dstChainId: Int!
    srcSender: String!
    dstReceiver: String!
    verificationFee: BigInt!
    executionFee: BigInt!
    options: String!
    message: String!
    address: String!
    blockNumber: BigInt!
    transactionHash: String!
    timestamp: BigInt!
    interchainTransaction: [InterchainTransaction!]!
  }

  type InterchainTransactionReceived {
    id: String!
    dstChainId: Int!
    name: String!
    transactionId: String!
    dbNonce: BigInt!
    entryIndex: BigInt!
    srcChainId: Int!
    srcSender: String!
    dstReceiver: String!
    address: String!
    blockNumber: BigInt!
    transactionHash: String!
    timestamp: BigInt!
    interchainTransaction: [InterchainTransaction!]!
  }

  type TxReadinessType {
    code: String
    notes: [String]
    status: String
    firstArg: String
    secondArg: String
  }

  type InterchainTransaction {
    id: String!
    srcChainId: Int!
    dstChainId: Int!
    srcSender: String!
    dstReceiver: String!
    sentAt: BigInt
    receivedAt: BigInt
    createdAt: BigInt
    updatedAt: BigInt
    interchainTransactionSentId: String
    interchainTransactionSent: InterchainTransactionSent
    interchainTransactionReceivedId: String
    interchainTransactionReceived: InterchainTransactionReceived
    status: String
    interchainBatchId: String
    interchainBatch: InterchainBatch
    txReadiness: TxReadinessType
    timeElapsed: Int
  }

  type Counts {
    total: Int!
    sent: Int!
    received: Int!
  }

  type Query {
    interchainBatches(
      srcChainId: Int,
      dstChainId: Int,
      status: String,
      limit: Int
    ): [InterchainBatch!]!
  
    interchainTransactions(
      srcChainId: Int,
      dstChainId: Int,
      status: String,
      limit: Int,
      dstReceiver: String,
      transactionHash: String
    ): [InterchainTransaction!]!
    
    appConfigV1s: [AppConfigV1!]!
    
    counts(srcChainId: Int, dstChainId: Int): Counts!
  }
`

const resolvers = {
  Query: {
    interchainBatches: async (
      _: any,
      args: {
        srcChainId?: number
        dstChainId?: number
        limit?: number
        status?: string
      }
    ) => {
      const { srcChainId, dstChainId, limit = 100, status } = args

      const where: InterchainBatchQueryFilter = {}

      if (srcChainId !== undefined) where.srcChainId = srcChainId
      if (dstChainId !== undefined) where.dstChainId = dstChainId
      if (status !== undefined) where.status = status

      return await prisma.interchainBatch.findMany({
        where,
        include: {
          interchainTransactions: true,
        },
        take: limit,
      })
    },
    interchainTransactions: async (
      _: any,
      args: {
        srcChainId?: number
        dstChainId?: number
        limit?: number
        status?: string
        dstReceiver?: string
        transactionHash?: string
      }
    ) => {
      const {
        srcChainId,
        dstChainId,
        limit = 100,
        status,
        dstReceiver,
        transactionHash,
      } = args

      const where: InterchainTransactionQueryFilter = {}

      if (srcChainId !== undefined) where.srcChainId = srcChainId
      if (dstChainId !== undefined) where.dstChainId = dstChainId
      if (status !== undefined) where.status = status
      if (dstReceiver !== undefined) where.dstReceiver = dstReceiver
      if (transactionHash !== undefined) where.id = transactionHash

      return await prisma.interchainTransaction.findMany({
        where,
        include: {
          interchainBatch: true,
          interchainTransactionSent: true,
          interchainTransactionReceived: true,
        },
        orderBy: {
          sentAt: 'desc',
        },
        take: limit,
      })
    },
    appConfigV1s: async () => {
      const appConfigs = await prisma.appConfigV1.findMany({
        include: {
          InterchainBatches: true,
        },
      })

      const decodedAppConfigs = appConfigs.map((config) => {
        if (typeof config.modules === 'string') {
          try {
            config.modules = JSON.parse(config.modules)
          } catch (error) {
            console.error(
              `Error parsing 'modules' for AppConfigV1 ID ${config.id}:`,
              error
            )
            config.modules = []
          }
        }

        return config
      })

      return decodedAppConfigs
    },
    counts: async (_: any, args: CountsFilter) => {
      const { srcChainId, dstChainId } = args

      const where: CountsFilter = {}

      if (srcChainId !== undefined) {
        where.srcChainId = srcChainId
      }

      if (dstChainId !== undefined) {
        where.dstChainId = dstChainId
      }

      const total = await prisma.interchainTransaction.count({
        where,
      })
      const sent = await prisma.interchainTransactionSent.count({
        where,
      })
      const received = await prisma.interchainTransactionReceived.count({
        where,
      })

      return {
        total,
        sent,
        received,
      }
    },
  },
  InterchainBatch: {
    interchainTransactions: async (parent: InterchainBatch) => {
      return await prisma.interchainTransaction.findMany({
        where: { interchainBatchId: parent.id },
      })
    },
    appConfigV1: async (parent: InterchainBatch) => {
      if (!parent.appConfigV1Id) {
        return null
      }
      const appConfig = await prisma.appConfigV1.findUnique({
        where: { id: parent.appConfigV1Id },
      })

      if (appConfig && typeof appConfig.modules === 'string') {
        appConfig.modules = JSON.parse(appConfig.modules)
      }

      return appConfig
    },
  },
  InterchainTransaction: {
    interchainTransactionSent: async (parent: InterchainTransaction) => {
      if (!parent.interchainTransactionSentId) {
        return null
      }
      return await prisma.interchainTransactionSent.findUnique({
        where: { id: parent.interchainTransactionSentId },
      })
    },
    interchainTransactionReceived: async (parent: InterchainTransaction) => {
      if (!parent.interchainTransactionReceivedId) {
        return null
      }
      return await prisma.interchainTransactionReceived.findUnique({
        where: { id: parent.interchainTransactionReceivedId },
      })
    },
    timeElapsed: async (parent: InterchainTransaction) => {
      const sentAt = Number(parent.sentAt)
      const now = Math.floor(Date.now() / 1000)

      return now - sentAt
    },
    txReadiness: async (parent: InterchainTransaction) => {
      const sentTxn = await prisma.interchainTransactionSent.findUnique({
        where: { id: parent.interchainTransactionSentId },
      })

      const { srcChainId, dstChainId, srcSender, dstReceiver } = parent
      const chainConfig = networkConfig[dstChainId]

      if (chainConfig && sentTxn) {
        const { dbNonce, entryIndex, options, message } = sentTxn

        const viemClient = chainConfig.client

        const data = await viemClient.readContract({
          address: chainConfig.InterchainClientV1.address,
          abi: InterchainClientV1Abi,
          functionName: 'getTxReadinessV1',
          args: [
            [
              srcChainId,
              dstChainId,
              dbNonce,
              entryIndex,
              srcSender,
              dstReceiver,
              options, // from interchainTxn sent
              message, // from interchainTxn sent
            ],
            [],
          ],
        })

        const statusCode = data[0]
        const statusInfo = getTransactionStatusInfo(statusCode)

        return {
          ...statusInfo,
          status: statusCode,
          firstArg: data[1],
          secondArg: data[2],
        }
      }
    },
  },
}

export const schema = createSchema({
  typeDefs,
  resolvers,
})
