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
import { type Chain as SynapseChain } from '@types'

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

type Transports = Record<SynapseChain['id'], Transport>

const createTransports = (chains: SynapseChain[]): Transports => {
  return chains.reduce<Transports>((acc, chain) => {
    acc[chain.id] = fallback([
      http(chain.rpcUrls.primary),
      http(chain.rpcUrls.fallback),
    ])
    return acc
  }, {})
}

const synapseChains = Object.values(CHAINS_BY_ID)

const transports = createTransports(synapseChains)

export const wagmiConfig = createConfig({
  connectors,
  chains: supportedChains as unknown as readonly [Chain, ...Chain[]],
  transports,
  ssr: true,
})
