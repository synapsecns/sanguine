import { createSchema } from 'graphql-yoga'
import { PrismaClient } from '@prisma/client'
import {
  type InterchainBatch,
  type InterchainBatchQueryFilter,
  type InterchainTransaction,
  type InterchainTransactionQueryFilter,
} from '@/types'
import { publicClient } from '@/utils/publicClient'
import { InterchainClientV1Abi } from '@/abis/InterchainClientV1Abi'

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
    txReadiness: async (parent: InterchainTransaction) => {
      const sentTxn = await prisma.interchainTransactionSent.findUnique({
        where: { id: parent.interchainTransactionSentId },
      })

      if (sentTxn) {
        const { srcChainId, dstChainId, srcSender, dstReceiver } = parent
        const { dbNonce, entryIndex, options, message } = sentTxn

        const viemClient = publicClient[dstChainId]

        const data = await viemClient.readContract({
          address: networkDetails[dstChainId].InterchainClientV1.address,
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
          secondArt: data[2],
        }
      }
    },
  },
}

export const schema = createSchema({
  typeDefs,
  resolvers,
})

export const networkDetails: any = {
  11155111: {
    name: 'ethSepolia',
    InterchainClientV1: {
      address: '0x6bAb7426099ba52ac37F309903169C4c0A5f7534',
      abi: InterchainClientV1Abi,
    },
  },
  421614: {
    name: 'arbSepolia',
    InterchainClientV1: {
      address: '0x15ACDFd1F2027aE084B4d92da20D22cc945d07Ec',
      abi: InterchainClientV1Abi,
    },
  },
}

function getTransactionStatusInfo(statusCode: number) {
  const status = statusCode as TransactionStatus

  const statusLabel = TransactionStatus[
    status
  ] as keyof typeof TransactionStatus
  const notes = TransactionStatusNotes[status] || []
  return { code: statusLabel, notes }
}

enum TransactionStatus {
  Ready = 0,
  AlreadyExecuted = 1,
  EntryAwaitingResponses = 2,
  EntryConflict = 3,
  ReceiverNotICApp = 4,
  ReceiverZeroRequiredResponses = 5,
  TxWrongDstChainId = 6,
  UndeterminedRevert = 7,
}

const TransactionStatusNotes: { [key in TransactionStatus]: string[] } = {
  [TransactionStatus.Ready]: [],
  [TransactionStatus.AlreadyExecuted]: ['`firstArg` is the transaction ID'],
  [TransactionStatus.EntryAwaitingResponses]: [
    '`firstArg` is the number of responses received',
    '`secondArg` is the number of responses required',
  ],
  [TransactionStatus.EntryConflict]: [
    '`firstArg` is the address of the module. This is either one of the modules that the app trusts, or the Guard module used by the app',
  ],
  [TransactionStatus.ReceiverNotICApp]: ['`firstArg` is the receiver address'],
  [TransactionStatus.ReceiverZeroRequiredResponses]: [
    'the app config requires zero responses for the transaction',
  ],
  [TransactionStatus.TxWrongDstChainId]: [
    '`firstArg` is the destination chain ID',
  ],
  [TransactionStatus.UndeterminedRevert]: [
    'the transaction will revert for another reason',
  ],
}
