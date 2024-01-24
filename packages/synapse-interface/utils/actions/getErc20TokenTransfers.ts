import { createPublicClient, http, Address, Log, Chain } from 'viem'

/**
 * Wrapper function around getErc20TokenTransferLogs() and transformTransferLogsToData()
 *
 * @param tokenAddress token to query transfer logs
 * @param fromAddress transfer from address
 * @param toAddress transfer to address
 * @param startBlock block to start query logs
 * @param toBlock block to stop query logs
 * @param chain viem Chain where token resides
 * @returns Will return raw logs and parsed logs
 */
export const getErc20TokenTransfers = async (
  tokenAddress: Address,
  fromAddress: Address,
  toAddress: Address,
  chain: Chain,
  startBlock: bigint,
  toBlock?: bigint
) => {
  if (!tokenAddress || !fromAddress || !toAddress) {
    console.error('Invalid address')
    return null
  }
  if (!startBlock && typeof startBlock !== 'bigint') {
    console.error('Invalid start block')
    return null
  }
  if (!chain) {
    console.error('Invalid Viem Chain')
    return null
  }

  const logs = await getErc20TokenTransferLogs(
    tokenAddress,
    fromAddress,
    toAddress,
    chain,
    startBlock,
    toBlock
  )
  const data = transformTransferLogsToData(logs)

  return {
    logs,
    data,
  }
}

const getErc20TokenTransferLogs = async (
  tokenAddress: Address,
  fromAddress: Address,
  toAddress: Address,
  chain: Chain,
  startBlock: bigint,
  toBlock?: bigint
) => {
  /** Create public client to access logs while connected to any chain */
  const publicClient = createPublicClient({
    chain,
    transport: http(),
  })

  const logs: Log[] = await publicClient.getLogs({
    address: tokenAddress,
    event: {
      type: 'event',
      name: 'Transfer',
      inputs: [
        { type: 'address', indexed: true, name: 'from' },
        { type: 'address', indexed: true, name: 'to' },
        { type: 'uint256', indexed: false, name: 'value' },
      ],
    },
    args: {
      from: fromAddress,
      to: toAddress,
    },
    fromBlock: startBlock,
    toBlock: toBlock ?? null,
  })

  return logs
}

const transformTransferLogsToData = (logs: any[]) => {
  return logs.map((log) => {
    return {
      tokenAddress: log.address,
      fromAddress: log?.args?.from,
      toAddress: log?.args?.to,
      transferValue: log?.args?.value,
      blockNumber: log.blockNumber,
      transactionHash: log.transactionHash,
    }
  })
}
