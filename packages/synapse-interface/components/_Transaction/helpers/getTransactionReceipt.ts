import {
  createPublicClient,
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
