import { Transport, type Chain } from 'viem'
import { fallback, http } from '@wagmi/core'

import { CHAINS_BY_ID } from '@/constants/chains'

type Transports = Record<Chain['id'], Transport>

export const createTransports = (chains: Chain[]): Transports => {
  return chains.reduce<Transports>((acc, chain) => {
    const synapseChain = CHAINS_BY_ID[chain.id]
    const proxyPath = `/api/rpc/${chain.id}`
    // WalletConnect extracts this URL from the transport and requires HTTP(S).
    // Resolve against the current origin so previews and localhost use their proxy.
    const proxyUrl =
      typeof window === 'undefined'
        ? proxyPath
        : new URL(proxyPath, window.location.origin).href

    const rpcUrls = [
      proxyUrl,
      synapseChain.rpcUrls.primary,
      synapseChain.rpcUrls.fallback,
    ]

    acc[chain.id] = fallback(rpcUrls.map((url) => http(url)))
    return acc
  }, {})
}
