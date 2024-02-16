import { createPublicClient, http, Address, Chain as ViemChain } from 'viem'

import { rawChains } from '@/wagmiConfig'
import { Chain } from '@/utils/types'

export const getTransactionReceipt = async (txHash: Address, chain: Chain) => {
  const viemChain = rawChains.find((rawChain) => chain.id === rawChain.id)

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
