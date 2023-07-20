import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'
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
} from '@constants/extraWagmiChains'
import { WagmiConfig, configureChains, createConfig } from 'wagmi'
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
import {
  RainbowKitProvider,
  darkTheme,
  getDefaultWallets,
  connectorsForWallets,
} from '@rainbow-me/rainbowkit'
import { JsonRpcProvider } from '@ethersproject/providers'
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc'
import * as CHAINS from '@constants/chains/master'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import CustomToaster from '@/components/toast'
import { SegmentAnalyticsProvider } from '@/contexts/SegmentAnalyticsProvider'

import { Provider } from 'react-redux'
import { store } from '@/store/store'
import { WalletAnalyticsProvider } from '@/contexts/WalletAnalyticsProvider'

const rawChains = [
  mainnet,
  arbitrum,
  aurora,
  avalanche,
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
    configRpc: configChain.rpc,
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
  ]
)

const projectId = 'ab0a846bc693996606734d788cb6561d'

const { wallets } = getDefaultWallets({
  appName: 'Synapse',
  projectId,
  chains,
})

const connectors = connectorsForWallets([...wallets])

export const wagmiConfig = createConfig({
  autoConnect: true,
  connectors,
  publicClient,
  webSocketPublicClient,
})

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Head>
        <title>Synapse Protocol</title>
      </Head>
      <WagmiConfig config={wagmiConfig}>
        <RainbowKitProvider chains={chains} theme={darkTheme()}>
          <SynapseProvider chains={chains}>
            <SegmentAnalyticsProvider>
              <WalletAnalyticsProvider>
                <Provider store={store}>
                  <Component {...pageProps} />
                  <CustomToaster />
                </Provider>
              </WalletAnalyticsProvider>
            </SegmentAnalyticsProvider>
          </SynapseProvider>
        </RainbowKitProvider>
      </WagmiConfig>
    </>
  )
}

export default App
