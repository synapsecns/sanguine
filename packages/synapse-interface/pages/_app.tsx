import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import '@/patch'
import { Analytics } from '@vercel/analytics/react'
import { PersistGate } from 'redux-persist/integration/react'
import LogRocket from 'logrocket'
import setupLogRocketReact from 'logrocket-react'

import { WagmiConfig } from 'wagmi'
import { RainbowKitProvider, darkTheme } from '@rainbow-me/rainbowkit'
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
import { wagmiChains, wagmiConfig } from '@/wagmiConfig'

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
      <WagmiConfig config={wagmiConfig}>
        <RainbowKitProvider chains={wagmiChains} theme={darkTheme()}>
          <SynapseProvider chains={wagmiChains}>
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
