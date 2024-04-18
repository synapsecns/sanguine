import { ponder } from '@/generated'
import { statusToString } from './status'

ponder.on('messageBus:MessageSent', async ({ event, context }) => {
  const { Message } = context.db
  const srcOptionalFields = {
    srcTxHash: event.transaction.hash,
    srcTimestamp: new Date(Number(event.block.timestamp) * 1000).toISOString(),
  }
  await Message.upsert({
    id: event.args.messageId,
    // Create a new message with Null status
    create: {
      srcChainId: event.args.srcChainID,
      dstChainId: event.args.dstChainId,
      status: 'Null',
      ...srcOptionalFields,
    },
    // Only populate the source chain optional fields if message is already known
    update: {
      ...srcOptionalFields,
    },
  })
})

ponder.on('messageBus:Executed', async ({ event, context }) => {
  const { Message } = context.db
  const dstOptionalFields = {
    dstTxHash: event.transaction.hash,
    dstTimestamp: new Date(Number(event.block.timestamp) * 1000).toISOString(),
  }
  const status = statusToString(event.args.status)
  await Message.upsert({
    id: event.args.messageId,
    // Create a new message with the status from the event
    create: {
      srcChainId: event.args.srcChainId,
      dstChainId: BigInt(context.network.chainId),
      status,
      ...dstOptionalFields,
    },
    // Update in following cases:
    // Null -> Anything
    // Anything -> Success
    update: ({ current }) =>
      current.status === 'Null' || status === 'Success'
        ? {
            status,
            ...dstOptionalFields,
          }
        : {},
  })
})
