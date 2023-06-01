import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import { useEffect } from 'react'
import type { AppProps } from 'next/app'
import { useRouter } from 'next/router'
import {
  boba,
  cronos,
  dfk,
  dogechain,
  klaytn,
} from '@constants/extraWagmiChains'
import { WagmiConfig, configureChains, createClient } from 'wagmi'
import {
  arbitrum,
  aurora,
  avalanche,
  bsc,
  canto,
  fantom,
  harmonyOne,
  mainnet,
  metis,
  moonbeam,
  moonriver,
  optimism,
  polygon,
} from 'wagmi/chains'
import {
  RainbowKitProvider,
  darkTheme,
  getDefaultWallets,
} from '@rainbow-me/rainbowkit'
import { JsonRpcProvider } from '@ethersproject/providers'
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc'
import * as CHAINS from '@constants/chains/master'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import * as amplitude from '@amplitude/analytics-browser'
import CustomToaster from '@/components/toast'

const AMPLITUDE_API_KEY = process.env.NEXT_PUBLIC_AMPLITUDE_KEY

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

const { chains, provider } = configureChains(chainsMatured, [
  jsonRpcProvider({
    rpc: (chain) => ({
      http: chain['configRpc'],
    }),
  }),
])

const { connectors } = getDefaultWallets({
  appName: 'Synapse',
  chains,
})

export const wagmiClient = createClient({
  autoConnect: true,
  connectors,
  provider,
})

const App = ({ Component, pageProps }: AppProps) => {
  const router = useRouter()

  useEffect(() => {
    if (router.isReady) {
      try {
        amplitude.init(AMPLITUDE_API_KEY, 'test@test.com', {
          defaultTracking: {
            sessions: true,
            pageViews: true,
            formInteractions: true,
            fileDownloads: true,
          },
        })
        console.log('amplitude initialized')
      } catch (error) {
        console.error('Error initializing Amplitude: ', error)
      }
    }
  }, [router.isReady])

  return (
    <WagmiConfig client={wagmiClient}>
      <RainbowKitProvider chains={chains} theme={darkTheme()}>
        <SynapseProvider chains={chains}>
          <Component {...pageProps} />
          <CustomToaster />
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}

export default App
