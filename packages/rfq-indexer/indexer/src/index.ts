import { trim } from 'viem'

import { ponder } from '@/generated'
import { formatAmount } from './utils/formatAmount'
import { getChainName } from './utils/chains'
import {
  contractNetworks_FastBridgeV1,
  contractNetworks_FastBridgeV2,
} from '@/ponder.config'

// ponder doesnt seem to handle a situation where two contracts on the same chain share the same topic.
// it seems instead to process the topic under *both* contracts, so this extra step is necessary to ensure
// it only performs indexing functions on the handler respective to the contract. duplicate errors will occur otherwise
const validContractAddresses = {
  FastBridgeV1: Object.values(contractNetworks_FastBridgeV1).map(
    (network: any) => network.contractAddr
  ),
  FastBridgeV2: Object.values(contractNetworks_FastBridgeV2).map(
    (network: any) => network.contractAddr
  ),
}

//* ############ FastBridgeV2 ########### *//

/* ORIGIN CHAIN EVENTS */

ponder.on('v2:BridgeQuoteDetails', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    db: { BridgeQuoteDetails },
    network: { chainId },
  } = context

  const {
    args: { transactionId, quoteId },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  await BridgeQuoteDetails.create({
    id: transactionId,
    data: {
      transactionId,
      quoteId,
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v2:BridgeRequested', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    db: { BridgeRequestEvents },
    network: { chainId },
  } = context

  const {
    args: {
      transactionId,
      sender,
      request,
      destChainId,
      originToken,
      destToken,
      originAmount,
      destAmount,
      sendChainGas,
    },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  await BridgeRequestEvents.create({
    id: transactionId,
    data: {
      transactionId,
      sender: trim(sender),
      request,
      originChainId: Number(chainId),
      originChain: getChainName(Number(chainId)),
      destChainId: Number(destChainId),
      destChain: getChainName(Number(destChainId)),
      originToken: trim(originToken),
      destToken: trim(destToken),
      originAmount,
      originAmountFormatted: formatAmount(originAmount, originToken),
      destAmount,
      destAmountFormatted: formatAmount(destAmount, destToken),
      sendChainGas,
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v2:BridgeDepositRefunded', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, to, token, amount },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeDepositRefundedEvents },
    network: { chainId },
  } = context

  await BridgeDepositRefundedEvents.create({
    id: transactionId,
    data: {
      transactionId,
      to: trim(to),
      token: trim(token),
      amount,
      amountFormatted: formatAmount(amount, token),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
      originChainId: chainId,
      originChain: getChainName(Number(chainId)),
    },
  })
})

ponder.on('v2:BridgeProofProvided', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, relayer },
    block: { timestamp },
    transaction: { hash },
    log: { address, blockNumber }, // may want to add address here eventually
  } = event

  const {
    db: { BridgeProofProvidedEvents },
    network: { chainId },
  } = context

  await BridgeProofProvidedEvents.upsert({
    id: transactionId,
    // Save the full data first time we index this event
    create: {
      transactionId,
      relayer: trim(relayer),
      originChainId: chainId,
      originChain: getChainName(Number(chainId)),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
    // Update the data with the latest event data on subsequent indexes
    update: {
      relayer: trim(relayer),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v2:BridgeDepositClaimed', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, relayer, to, token, amount },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeDepositClaimedEvents },
    network: { chainId },
  } = context

  await BridgeDepositClaimedEvents.create({
    id: transactionId,
    data: {
      transactionId,
      relayer: trim(relayer),
      to: trim(to),
      token: trim(token),
      amount,
      amountFormatted: formatAmount(amount, token),
      originChainId: chainId,
      originChain: getChainName(Number(chainId)),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v2:BridgeProofDisputed', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, relayer },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeProofDisputedEvents },
    network: { chainId },
  } = context

  await BridgeProofDisputedEvents.upsert({
    id: transactionId,
    create: {
      transactionId,
      relayer: trim(relayer),
      originChainId: Number(chainId),
      originChain: getChainName(Number(chainId)),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
    update: {
      relayer: trim(relayer),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

/* DESTINATION CHAIN EVENTS */

ponder.on('v2:BridgeRelayed', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV2.includes(event.log.address)) {
    return
  }

  const {
    args: {
      transactionId,
      relayer,
      to,
      originChainId,
      originToken,
      destToken,
      originAmount,
      destAmount,
    },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeRelayedEvents },
    network: { chainId },
  } = context

  await BridgeRelayedEvents.create({
    id: transactionId,
    data: {
      transactionId,
      relayer: trim(relayer),
      to: trim(to),
      originChainId: Number(originChainId),
      originChain: getChainName(Number(originChainId)),
      destChainId: Number(chainId),
      destChain: getChainName(Number(chainId)),
      originToken: trim(originToken),
      destToken: trim(destToken),
      originAmount,
      originAmountFormatted: formatAmount(originAmount, originToken),
      destAmount,
      destAmountFormatted: formatAmount(destAmount, destToken),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

//* ############ FastBridgeV1 ########### *//

/* ORIGIN CHAIN EVENTS */

ponder.on('v1:BridgeRequested', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV1.includes(event.log.address)) {
    return
  }

  const {
    db: { BridgeRequestEvents },
    network: { chainId },
  } = context

  const {
    args: {
      transactionId,
      sender,
      request,
      destChainId,
      originToken,
      destToken,
      originAmount,
      destAmount,
      sendChainGas,
    },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  await BridgeRequestEvents.create({
    id: transactionId,
    data: {
      transactionId,
      sender: trim(sender),
      request,
      originChainId: Number(chainId),
      originChain: getChainName(Number(chainId)),
      destChainId: Number(destChainId),
      destChain: getChainName(Number(destChainId)),
      originToken: trim(originToken),
      destToken: trim(destToken),
      originAmount,
      originAmountFormatted: formatAmount(originAmount, originToken),
      destAmount,
      destAmountFormatted: formatAmount(destAmount, destToken),
      sendChainGas,
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v1:BridgeDepositRefunded', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV1.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, to, token, amount },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeDepositRefundedEvents },
    network: { chainId },
  } = context

  await BridgeDepositRefundedEvents.create({
    id: transactionId,
    data: {
      transactionId,
      to: trim(to),
      token: trim(token),
      amount,
      amountFormatted: formatAmount(amount, token),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
      originChainId: chainId,
      originChain: getChainName(Number(chainId)),
    },
  })
})

ponder.on('v1:BridgeProofProvided', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV1.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, relayer },
    block: { timestamp },
    transaction: { hash },
    log: { address, blockNumber }, // may want to add address here eventually
  } = event

  const {
    db: { BridgeProofProvidedEvents },
    network: { chainId },
  } = context

  await BridgeProofProvidedEvents.upsert({
    id: transactionId,
    // Save the full data first time we index this event
    create: {
      transactionId,
      relayer: trim(relayer),
      originChainId: chainId,
      originChain: getChainName(Number(chainId)),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
    // Update the data with the latest event data on subsequent indexes
    update: {
      relayer: trim(relayer),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v1:BridgeDepositClaimed', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV1.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, relayer, to, token, amount },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeDepositClaimedEvents },
    network: { chainId },
  } = context

  await BridgeDepositClaimedEvents.create({
    id: transactionId,
    data: {
      transactionId,
      relayer: trim(relayer),
      to: trim(to),
      token: trim(token),
      amount,
      amountFormatted: formatAmount(amount, token),
      originChainId: chainId,
      originChain: getChainName(Number(chainId)),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

ponder.on('v1:BridgeProofDisputed', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV1.includes(event.log.address)) {
    return
  }

  const {
    args: { transactionId, relayer },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeProofDisputedEvents },
    network: { chainId },
  } = context

  await BridgeProofDisputedEvents.upsert({
    id: transactionId,
    create: {
      transactionId,
      relayer: trim(relayer),
      originChainId: Number(chainId),
      originChain: getChainName(Number(chainId)),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
    update: {
      relayer: trim(relayer),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})

/* DESTINATION CHAIN EVENTS */

ponder.on('v1:BridgeRelayed', async ({ event, context }) => {
  if (!validContractAddresses.FastBridgeV1.includes(event.log.address)) {
    return
  }

  const {
    args: {
      transactionId,
      relayer,
      to,
      originChainId,
      originToken,
      destToken,
      originAmount,
      destAmount,
    },
    block: { timestamp },
    transaction: { hash },
    log: { blockNumber },
  } = event

  const {
    db: { BridgeRelayedEvents },
    network: { chainId },
  } = context

  await BridgeRelayedEvents.create({
    id: transactionId,
    data: {
      transactionId,
      relayer: trim(relayer),
      to: trim(to),
      originChainId: Number(originChainId),
      originChain: getChainName(Number(originChainId)),
      destChainId: Number(chainId),
      destChain: getChainName(Number(chainId)),
      originToken: trim(originToken),
      destToken: trim(destToken),
      originAmount,
      originAmountFormatted: formatAmount(originAmount, originToken),
      destAmount,
      destAmountFormatted: formatAmount(destAmount, destToken),
      blockNumber: BigInt(blockNumber),
      blockTimestamp: Number(timestamp),
      transactionHash: hash,
    },
  })
})
