import '@/patch'
import { Transport, type Chain } from 'viem'
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
} from '@wagmi/core/chains'
import { createConfig, fallback, http } from '@wagmi/core'
import {
  metaMaskWallet,
  rabbyWallet,
  coinbaseWallet,
  rainbowWallet,
  walletConnectWallet,
  trustWallet,
  ledgerWallet,
  frameWallet,
  safeWallet,
} from '@rainbow-me/rainbowkit/wallets'
import { connectorsForWallets } from '@rainbow-me/rainbowkit'

import { CHAINS_BY_ID } from '@/constants/chains'
import { dfk, dogechain } from '@/constants/extraWagmiChains'

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
].map((chain) => {
  return {
    ...chain,
    configRpc: CHAINS_BY_ID[chain.id]?.rpcUrls.primary,
    fallbackRpc: CHAINS_BY_ID[chain.id]?.rpcUrls.fallback,
    iconUrl: CHAINS_BY_ID[chain.id]?.chainImg.src,
  }
})
const appName = 'Synapse'
const projectId = 'ab0a846bc693996606734d788cb6561d'

const connectors = connectorsForWallets(
  [
    {
      groupName: 'Wallets',
      wallets: [
        metaMaskWallet,
        walletConnectWallet,
        coinbaseWallet,
        rainbowWallet,
        rabbyWallet,
        trustWallet,
        ledgerWallet,
        frameWallet,
        safeWallet,
      ],
    },
  ],
  {
    projectId,
    appName,
  }
)

type Transports = Record<Chain['id'], Transport>

const createTransports = (chains: Chain[]): Transports => {
  return chains.reduce<Transports>((acc, chain) => {
    const synapseChain = CHAINS_BY_ID[chain.id]

    acc[chain.id] = fallback([
      http(synapseChain.rpcUrls.primary),
      http(synapseChain.rpcUrls.fallback),
    ])
    return acc
  }, {})
}

const transports = createTransports(supportedChains as Chain[])

export const wagmiConfig = createConfig({
  connectors,
  chains: supportedChains as unknown as readonly [Chain, ...Chain[]],
  transports,
  ssr: true,
})
