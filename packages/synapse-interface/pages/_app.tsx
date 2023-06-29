import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
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
import { WagmiConfig, configureChains, createClient } from 'wagmi'
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
} from '@rainbow-me/rainbowkit'
import { JsonRpcProvider } from '@ethersproject/providers'
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc'
import * as CHAINS from '@constants/chains/master'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import CustomToaster from '@/components/toast'

import { Provider } from 'react-redux'
import { store } from '@/store/store'

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
  return (
    <WagmiConfig client={wagmiClient}>
      <RainbowKitProvider chains={chains} theme={darkTheme()}>
        <SynapseProvider chains={chains}>
          <Provider store={store}>
            <Component {...pageProps} />
            <CustomToaster />
          </Provider>
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}

export default App
