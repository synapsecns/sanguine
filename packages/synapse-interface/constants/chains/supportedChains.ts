import {
  arbitrum,
  aurora,
  avalanche,
  base,
  bsc,
  blast,
  boba,
  canto,
  cronos,
  fantom,
  klaytn,
  harmonyOne,
  mainnet,
  metis,
  moonbeam,
  moonriver,
  optimism,
  polygon,
  scroll,
} from '@wagmi/core/chains'

import { dfk, dogechain } from '@/constants/chains/extraWagmiChains'
import { CHAINS_BY_ID } from '@/constants/chains'

export const supportedChains = [
  mainnet,
  arbitrum,
  aurora,
  avalanche,
  base,
  blast,
  bsc,
  canto,
  fantom,
  harmonyOne,
  metis,
  moonbeam,
  moonriver,
  optimism,
  polygon,
  klaytn,
  cronos,
  dfk,
  dogechain,
  boba,
  scroll,
].map((chain) => {
  return {
    ...chain,
    configRpc: CHAINS_BY_ID[chain.id]?.rpcUrls.primary,
    fallbackRpc: CHAINS_BY_ID[chain.id]?.rpcUrls.fallback,
    iconUrl: CHAINS_BY_ID[chain.id]?.chainImg.src,
  }
})
