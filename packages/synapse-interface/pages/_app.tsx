import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
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
import { alchemyProvider } from 'wagmi/providers/alchemy'
import { publicProvider } from 'wagmi/providers/public'
import { jsonRpcProvider } from 'wagmi/providers/jsonRpc'
import * as CHAINS from '@constants/chains/master'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import CustomToaster from '@/components/toast'
const App = ({ Component, pageProps }: AppProps) => {
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
  const chainsWithIcons = []
  for (const chain of rawChains) {
    const configChain = Object.values(CHAINS).filter(
      (chainObj) => chainObj.id === chain.id
    )[0]

    chainsWithIcons.push({
      ...chain,
      iconUrl: configChain.chainImg.src,
      configRpc: configChain.rpc,
    })
  }

  const { chains, provider } = configureChains(chainsWithIcons, [
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

  const wagmiClient = createClient({
    autoConnect: true,
    connectors,
    provider,
  })

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
