import { createPublicClient, http, Address, Log, Chain } from 'viem'

export const getErc20TokenTransfers = async (
  tokenAddress: Address,
  fromAddress: Address,
  toAddress: Address,
  startBlock: bigint,
  chain: Chain
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
    startBlock,
    chain
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
  startBlock: bigint,
  chain: Chain
) => {
  // const publicClient = getPublicClient(wagmiConfig[chainId])

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
