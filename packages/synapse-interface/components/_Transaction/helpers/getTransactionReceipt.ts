import {
  createPublicClient,
  parseAbiItem,
  http,
  type Address,
  type Chain as ViemChain,
} from 'viem'

import { supportedChains } from '@/constants/chains/supportedChains'

export const getTransactionReceipt = async (
  txHash: Address,
  chainId: number
) => {
  const viemChain = supportedChains.find((c) => c.id === chainId)

  const publicClient = createPublicClient({
    chain: viemChain as ViemChain,
    transport: http(),
  })

  try {
    const receipt = await publicClient.getTransactionReceipt({
      hash: txHash,
    })
    return receipt
  } catch (error) {
    console.error('Error in getTransactionReceipt: ', error)
    return null
  }
}

export const getTransactionRefundLogs = async (
  bridgeContract: Address,
  chainId: number
) => {
  const viemChain = supportedChains.find((c) => c.id === chainId)

  const publicClient = createPublicClient({
    chain: viemChain as ViemChain,
    transport: http(),
  })

  try {
    const logs = await publicClient.getLogs({
      address: bridgeContract,
      // event: {
      //   name: 'BridgeDepositRefunded',
      //   type: 'function',
      //   inputs: [
      //     { type: 'bytes32', indexed: true, name: 'transactionId' },
      //     { type: 'address', indexed: true, name: 'to' },
      //     { type: 'address', indexed: false, name: 'token' },
      //     { type: 'uint256', indexed: false, name: 'amount' },
      //   ],
      // },
      // event: parseAbiItem(
      //   'event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)'
      // ),
      fromBlock: 15975605n,
      toBlock: 16009420n,
    })
    console.log('logs: ', logs)
  } catch (error) {
    console.error('Error in getTransactionRefundLogs: ', error)
  }
}

// ;[
//   {
//     name: 'transactionId',
//     type: 'bytes32',
//     indexed: true,
//     internalType: 'bytes32',
//   },
//   { name: 'to', type: 'address', indexed: true, internalType: 'address' },
//   { name: 'token', type: 'address', indexed: false, internalType: 'address' },
//   { name: 'amount', type: 'uint256', indexed: false, internalType: 'uint256' },
// ]
