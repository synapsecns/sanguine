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
  linea
} from '@wagmi/core/chains'

import { dfk, dogechain, worldchain, unichain, berachain } from '@/constants/chains/extraWagmiChains'
import { CHAINS_BY_ID } from '@/constants/chains'

export const supportedChains = [
  mainnet,
  unichain,
  berachain,
  worldchain,
  optimism,
  arbitrum,
  avalanche,
  base,
  scroll,
  linea,
  blast,
  metis,
  dfk,
  bsc,
  polygon,
  aurora,
  canto,
  klaytn,
  fantom,
  moonbeam,
  moonriver,
  cronos,
  dogechain,
  boba,
  harmonyOne,
].map((chain) => {
  return {
    ...chain,
    configRpc: CHAINS_BY_ID[chain.id]?.rpcUrls.primary,
    fallbackRpc: CHAINS_BY_ID[chain.id]?.rpcUrls.fallback,
    iconUrl: CHAINS_BY_ID[chain.id]?.chainImg.src,
  }
})
