import { createPublicClient, http, Address, Chain as ViemChain } from 'viem'
import { useEffect, useState } from 'react'

import { rawChains } from '@/wagmiConfig'
import { Chain } from '@/utils/types'

export const getTransactionReceipt = async (txHash: Address, chain: Chain) => {
  const viemChain = rawChains.find((rawChain) => chain.id === rawChain.id)

  const publicClient = createPublicClient({
    chain: viemChain as ViemChain,
    transport: http(),
  })

  const receipt = await publicClient.getTransactionReceipt({
    hash: txHash,
  })

  return receipt
}

export const useIsTxReverted = (txHash: Address, chain: Chain) => {
  const [isReverted, setIsReverted] = useState<boolean>(false)

  useEffect(() => {
    const fetchTransactionReceipt = async (txHash: Address, chain: Chain) => {
      const receipt = await getTransactionReceipt(txHash, chain)
      console.log('receipt:', receipt)
      if (receipt.status === 'reverted') {
        console.log('reverted')
        setIsReverted(true)
      }
    }
    fetchTransactionReceipt(txHash, chain)
  }, [])

  return [isReverted]
}
