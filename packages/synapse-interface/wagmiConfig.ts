import '@/patch'

import {
  boba,
  cronos,
  dfk,
  dogechain,
  klaytn,
  metis,
  aurora,
  canto,
  base,
  blast,
} from '@constants/extraWagmiChains'
import { configureChains, createConfig } from 'wagmi'
import {
  arbitrum,
  avalanche,
  bsc,
  fantom,
  harmonyOne,
  mainnet,
  moonbeam,
  moonriver,
  optimism,
  polygon,
} from 'wagmi/chains'
import { getDefaultWallets, connectorsForWallets } from '@rainbow-me/rainbowkit'
import { rabbyWallet } from '@rainbow-me/rainbowkit/wallets'
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc'
import { publicProvider } from 'wagmi/providers/public'
import * as CHAINS from '@constants/chains/master'

const rawChains = [
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
]

// Add custom icons
const chainsMatured = []
for (const chain of rawChains) {
  const configChain = Object.values(CHAINS).filter(
    (chainObj) => chainObj.id === chain.id
  )[0]

  chainsMatured.push({
    ...chain,
    iconUrl: configChain.chainImg.src,
    configRpc: configChain.rpcUrls.primary,
    fallbackRpc: configChain.rpcUrls.fallback,
  })
}

const { chains, publicClient, webSocketPublicClient } = configureChains(
  chainsMatured,
  [
    jsonRpcProvider({
      rpc: (chain) => ({
        http: chain['configRpc'],
      }),
    }),
    jsonRpcProvider({
      rpc: (chain) => ({
        http: chain['fallbackRpc'],
      }),
    }),
    publicProvider(),
  ]
)

const projectId = 'ab0a846bc693996606734d788cb6561d'

const { wallets } = getDefaultWallets({
  appName: 'Synapse',
  projectId,
  chains,
})

const connectors = connectorsForWallets([
  ...wallets,
  {
    groupName: 'Other',
    wallets: [rabbyWallet({ chains })],
  },
])

export const wagmiChains = chains

export const wagmiConfig = createConfig({
  autoConnect: true,
  connectors,
  publicClient,
  webSocketPublicClient,
})
