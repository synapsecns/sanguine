import { Transport, type Chain } from 'viem'
import { fallback, http } from '@wagmi/core'

import { CHAINS_BY_ID } from '@/constants/chains'

type Transports = Record<Chain['id'], Transport>

export const createTransports = (chains: Chain[]): Transports => {
  return chains.reduce<Transports>((acc, chain) => {
    const synapseChain = CHAINS_BY_ID[chain.id]

    acc[chain.id] = fallback([
      http(synapseChain.rpcUrls.primary),
      http(synapseChain.rpcUrls.fallback),
    ])
    return acc
  }, {})
}
