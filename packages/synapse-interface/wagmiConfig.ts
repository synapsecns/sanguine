import '@/patch'
import { type Chain } from 'viem'
import { createConfig } from '@wagmi/core'
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
import binanceWallet from '@binance/w3w-rainbow-connector-v2'

import { createTransports } from '@/utils/createTransports'
import { supportedChains } from '@/constants/chains/supportedChains'

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
        binanceWallet,
      ],
    },
  ],
  {
    projectId,
    appName,
  }
)

const transports = createTransports(supportedChains as Chain[])

export const wagmiConfig = createConfig({
  connectors,
  chains: supportedChains as unknown as readonly [Chain, ...Chain[]],
  transports,
  ssr: true,
})
