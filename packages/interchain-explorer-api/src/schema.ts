import { createSchema } from 'graphql-yoga'
import { PrismaClient } from '@prisma/client'
import { InterchainBatch, InterchainTransaction } from '@/types'

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
    count: Int!
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
    count: Int!
    interchainTransaction: [InterchainTransaction!]!
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
  }
`

interface InterchainBatchQueryFilter {
  srcChainId?: number
  dstChainId?: number
  status?: string
}

interface InterchainTransactionQueryFilter {
  srcChainId?: number
  dstChainId?: number
  status?: string
  dstReceiver?: string
  id?: string
}

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
  },
  InterchainBatch: {
    interchainTransactions: async (parent: InterchainBatch) => {
      return await prisma.interchainTransaction.findMany({
        where: { interchainBatchId: parent.id },
      })
    },
    appConfigV1: async (parent: InterchainBatch) => {
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
      return await prisma.interchainTransactionSent.findUnique({
        where: { id: parent.interchainTransactionSentId },
      })
    },
    interchainTransactionReceived: async (parent: InterchainTransaction) => {
      return await prisma.interchainTransactionReceived.findUnique({
        where: { id: parent.interchainTransactionReceivedId },
      })
    },
  },
}

export const schema = createSchema({
  typeDefs,
  resolvers,
})
