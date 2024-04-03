// import '@/patch'
import { type Chain } from 'viem'
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
  dogechain,
  dfk,
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
import { createConfig, http } from '@wagmi/core'
import { connectorsForWallets } from '@rainbow-me/rainbowkit'
import {
  metaMaskWallet,
  rabbyWallet,
  coinbaseWallet,
  rainbowWallet,
  walletConnectWallet,
} from '@rainbow-me/rainbowkit/wallets'

import {
  ARBITRUM,
  AURORA,
  AVALANCHE,
  BASE,
  BLAST,
  BNB,
  BOBA,
  CANTO,
  CRONOS,
  DFK,
  DOGE,
  ETH,
  FANTOM,
  HARMONY,
  KLAYTN,
  METIS,
  MOONBEAM,
  MOONRIVER,
  OPTIMISM,
  POLYGON,
} from '@/constants/chains/master'

export const rawChains = [
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
      ],
    },
  ],
  {
    projectId,
    appName,
  }
)

export const wagmiConfig = createConfig({
  connectors,
  chains: [
    mainnet,
    arbitrum,
    aurora,
    avalanche,
    base as Chain,
    blast,
    bsc,
    canto,
    fantom,
    harmonyOne,
    metis,
    moonbeam,
    moonriver,
    optimism as Chain,
    polygon,
    klaytn,
    cronos,
    dfk,
    dogechain,
    boba,
  ],
  transports: {
    [mainnet.id]: http(ETH.rpcUrls.primary),
    [arbitrum.id]: http(ARBITRUM.rpcUrls.primary),
    [aurora.id]: http(AURORA.rpcUrls.primary),
    [avalanche.id]: http(AVALANCHE.rpcUrls.primary),
    [base.id]: http(BASE.rpcUrls.primary),
    [blast.id]: http(BLAST.rpcUrls.primary),
    [bsc.id]: http(BNB.rpcUrls.primary),
    [canto.id]: http(CANTO.rpcUrls.primary),
    [fantom.id]: http(FANTOM.rpcUrls.primary),
    [harmonyOne.id]: http(HARMONY.rpcUrls.primary),
    [metis.id]: http(METIS.rpcUrls.primary),
    [moonbeam.id]: http(MOONBEAM.rpcUrls.primary),
    [moonriver.id]: http(MOONRIVER.rpcUrls.primary),
    [optimism.id]: http(OPTIMISM.rpcUrls.primary),
    [polygon.id]: http(POLYGON.rpcUrls.primary),
    [klaytn.id]: http(KLAYTN.rpcUrls.primary),
    [cronos.id]: http(CRONOS.rpcUrls.primary),
    [dfk.id]: http(DFK.rpcUrls.primary),
    [dogechain.id]: http(DOGE.rpcUrls.primary),
    [boba.id]: http(BOBA.rpcUrls.primary),
  },
})
