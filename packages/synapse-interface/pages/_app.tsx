import '@/styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import { Provider } from 'react-redux'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { PersistGate } from 'redux-persist/integration/react'
import { RainbowKitProvider, darkTheme } from '@rainbow-me/rainbowkit'
import { store, persistor } from '@/store/store'
import { WagmiProvider } from 'wagmi'
import LogRocket from 'logrocket'
import setupLogRocketReact from 'logrocket-react'
import { NextIntlClientProvider } from 'next-intl'
import { useRouter } from 'next/router'

import { SegmentAnalyticsProvider } from '@/contexts/SegmentAnalyticsProvider'
import { UserProvider } from '@/contexts/UserProvider'
import { BackgroundListenerProvider } from '@/contexts/BackgroundListenerProvider'
import CustomToaster from '@/components/toast'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import { wagmiConfig } from '@/wagmiConfig'
import { supportedChains } from '@/constants/chains/supportedChains'

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

const queryClient = new QueryClient()

function App({ Component, pageProps }: AppProps) {
  const router = useRouter()
  return (
    <>
      <Head>
        <title>Synapse Protocol</title>
      </Head>
      <NextIntlClientProvider
        locale={router.locale}
        timeZone="UTC"
        messages={pageProps.messages}
      >
        <WagmiProvider config={wagmiConfig}>
          <QueryClientProvider client={queryClient}>
            <RainbowKitProvider theme={darkTheme()}>
              <SynapseProvider chains={supportedChains}>
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
      </NextIntlClientProvider>
    </>
  )
}

export default App
