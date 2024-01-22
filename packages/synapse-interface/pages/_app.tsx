import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import '@/patch'
import { Analytics } from '@vercel/analytics/react'
import { PersistGate } from 'redux-persist/integration/react'
import LogRocket from 'logrocket'
import setupLogRocketReact from 'logrocket-react'

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
import { rabbyWallet } from '@rainbow-me/rainbowkit/wallets'
import { JsonRpcProvider } from '@ethersproject/providers'
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc'
import { publicProvider } from 'wagmi/providers/public'
import * as CHAINS from '@constants/chains/master'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import CustomToaster from '@/components/toast'
import { SegmentAnalyticsProvider } from '@/contexts/SegmentAnalyticsProvider'

import { Provider } from 'react-redux'
import { store, persistor } from '@/store/store'
import { UserProvider } from '@/contexts/UserProvider'

import ApplicationUpdater from '@/slices/application/updater'
import BridgeUpdater from '@/slices/bridge/updater'
import PortfolioUpdater from '@/slices/portfolio/updater'
import TransactionsUpdater from '@/slices/transactions/updater'
import _TransactionsUpdater from '@/slices/_transactions/updater'

import HotJar from '@/components/HotJar/HotJar'

const rawChains = [
  mainnet,
  arbitrum,
  aurora,
  avalanche,
  base,
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

// only initialize when in the browser
if (
  typeof window !== 'undefined' &&
  !location.hostname.match('synapseprotocol.com')
) {
  LogRocket.init('npdhrc/synapse-staging', {
    mergeIframes: true,
  })
  // plugins should also only be initialized when in the browser
  setupLogRocketReact(LogRocket)

  LogRocket.getSessionURL((sessionURL) => {
    console.log('session url for debugging ' + sessionURL)
  })
}

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

export const wagmiConfig = createConfig({
  autoConnect: true,
  connectors,
  publicClient,
  webSocketPublicClient,
})

function Updaters() {
  return (
    <>
      <ApplicationUpdater />
      <PortfolioUpdater />
      <TransactionsUpdater />
      <_TransactionsUpdater />
      <BridgeUpdater />
    </>
  )
}

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Head>
        <title>Synapse Protocol</title>
      </Head>
      <HotJar />
      <WagmiConfig config={wagmiConfig}>
        <RainbowKitProvider chains={chains} theme={darkTheme()}>
          <SynapseProvider chains={chains}>
            <Provider store={store}>
              <PersistGate loading={null} persistor={persistor}>
                <SegmentAnalyticsProvider>
                  <UserProvider>
                    <Updaters />
                    <Component {...pageProps} />
                    <Analytics />
                    <CustomToaster />
                  </UserProvider>
                </SegmentAnalyticsProvider>
              </PersistGate>
            </Provider>
          </SynapseProvider>
        </RainbowKitProvider>
      </WagmiConfig>
    </>
  )
}

export default App
