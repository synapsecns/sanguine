import '@/styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
import { useRouter } from 'next/router'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import { PersistGate } from 'redux-persist/integration/react'
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

import { RainbowKitProvider } from '@rainbow-me/rainbowkit'
import {
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
} from '@wagmi/core/chains'
import { Provider } from 'react-redux'
import { store, persistor } from '@/store/store'
import { SegmentAnalyticsProvider } from '@/contexts/SegmentAnalyticsProvider'
import { UserProvider } from '@/contexts/UserProvider'
import { BackgroundListenerProvider } from '@/contexts/BackgroundListenerProvider'
import CustomToaster from '@/components/toast'
import Head from 'next/head'
import { WagmiProvider } from 'wagmi'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'

import LogRocket from 'logrocket'
import setupLogRocketReact from 'logrocket-react'
import { wagmiConfig } from '@/wagmiConfig'

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

const chains = [
  {
    ...mainnet,
    configRpc: ETH.rpcUrls.primary,
    fallbackRpc: ETH.rpcUrls.fallback,
  },
  {
    ...arbitrum,
    configRpc: ARBITRUM.rpcUrls.primary,
    fallbackRpc: ARBITRUM.rpcUrls.fallback,
  },
  {
    ...aurora,
    configRpc: AURORA.rpcUrls.primary,
    fallbackRpc: AURORA.rpcUrls.fallback,
  },
  {
    ...avalanche,
    configRpc: AVALANCHE.rpcUrls.primary,
    fallbackRpc: AVALANCHE.rpcUrls.fallback,
  },
  {
    ...base,
    configRpc: BASE.rpcUrls.primary,
    fallbackRpc: BASE.rpcUrls.fallback,
  },
  {
    ...blast,
    configRpc: BLAST.rpcUrls.primary,
    fallbackRpc: BLAST.rpcUrls.fallback,
  },
  {
    ...bsc,
    configRpc: BNB.rpcUrls.primary,
    fallbackRpc: BNB.rpcUrls.fallback,
  },
  {
    ...canto,
    configRpc: CANTO.rpcUrls.primary,
    fallbackRpc: CANTO.rpcUrls.fallback,
  },
  {
    ...fantom,
    configRpc: FANTOM.rpcUrls.primary,
    fallbackRpc: FANTOM.rpcUrls.fallback,
  },
  {
    ...harmonyOne,
    configRpc: HARMONY.rpcUrls.primary,
    fallbackRpc: HARMONY.rpcUrls.fallback,
  },
  {
    ...metis,
    configRpc: METIS.rpcUrls.primary,
    fallbackRpc: METIS.rpcUrls.fallback,
  },
  {
    ...moonbeam,
    configRpc: MOONBEAM.rpcUrls.primary,
    fallbackRpc: MOONBEAM.rpcUrls.fallback,
  },
  {
    ...moonriver,
    configRpc: MOONRIVER.rpcUrls.primary,
    fallbackRpc: MOONRIVER.rpcUrls.fallback,
  },
  {
    ...optimism,
    configRpc: OPTIMISM.rpcUrls.primary,
    fallbackRpc: OPTIMISM.rpcUrls.fallback,
  },
  {
    ...polygon,
    configRpc: POLYGON.rpcUrls.primary,
    fallbackRpc: POLYGON.rpcUrls.fallback,
  },
  {
    ...klaytn,
    configRpc: KLAYTN.rpcUrls.primary,
    fallbackRpc: KLAYTN.rpcUrls.fallback,
  },
  {
    ...cronos,
    configRpc: CRONOS.rpcUrls.primary,
    fallbackRpc: CRONOS.rpcUrls.fallback,
  },
  {
    ...dfk,
    configRpc: DFK.rpcUrls.primary,
    fallbackRpc: DFK.rpcUrls.fallback,
  },
  {
    ...dogechain,
    configRpc: DOGE.rpcUrls.primary,
    fallbackRpc: DOGE.rpcUrls.fallback,
  },
  {
    ...boba,
    configRpc: BOBA.rpcUrls.primary,
    fallbackRpc: BOBA.rpcUrls.fallback,
  },
]

const queryClient = new QueryClient()

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>Synapse Protocol</title>
      </Head>
      <WagmiProvider config={wagmiConfig}>
        <QueryClientProvider client={queryClient}>
          <RainbowKitProvider>
            <SynapseProvider chains={chains}>
              <Provider store={store}>
                <PersistGate loading={null} persistor={persistor}>
                  <SegmentAnalyticsProvider>
                    <UserProvider>
                      <BackgroundListenerProvider>
                        <Component {...pageProps} />
                      </BackgroundListenerProvider>
                      <CustomToaster />
                    </UserProvider>
                  </SegmentAnalyticsProvider>
                </PersistGate>
              </Provider>
            </SynapseProvider>
          </RainbowKitProvider>
        </QueryClientProvider>
      </WagmiProvider>
    </>
  )
}

export default MyApp
