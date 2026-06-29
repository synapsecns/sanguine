import { Transport, type Chain } from 'viem'
import { fallback, http } from '@wagmi/core'

import { CHAINS_BY_ID } from '@/constants/chains'

type Transports = Record<Chain['id'], Transport>

const DFK_CHAIN_ID = 53935

export const createTransports = (chains: Chain[]): Transports => {
  return chains.reduce<Transports>((acc, chain) => {
    const synapseChain = CHAINS_BY_ID[chain.id]

    const rpcUrls = [
      ...(chain.id === DFK_CHAIN_ID ? [] : [`/api/rpc/${chain.id}`]),
      synapseChain.rpcUrls.primary,
      synapseChain.rpcUrls.fallback,
    ]

    acc[chain.id] = fallback(rpcUrls.map((url) => http(url)))
    return acc
  }, {})
}
