import { createPublicClient, http, Address, Log } from 'viem'
import { arbitrum } from 'viem/chains'

export const publicClient = createPublicClient({
  chain: arbitrum,
  transport: http(),
})

export const getErc20TokenTransferData = async (
  tokenAddress: Address,
  fromAddress: Address,
  toAddress: Address,
  startBlock: bigint
) => {
  if (!tokenAddress || !fromAddress || !toAddress) {
    console.error('Missing required address')
    return null
  }
  if (!startBlock && typeof startBlock !== 'bigint') {
    console.error('Invalid start block')
    return null
  }

  const logs = await getErc20TokenTransferLogs(
    tokenAddress,
    fromAddress,
    toAddress,
    startBlock
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
  startBlock: bigint
) => {
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
