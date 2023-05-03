import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
import { Provider as EthersProvider } from '@ethersproject/abstract-provider'
import { JsonRpcProvider } from '@ethersproject/providers'
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
    const iconUrl = Object.values(CHAINS).filter(
      (chainObj) => chainObj.id === chain.id
    )[0].chainImg.src
    chainsWithIcons.push({
      ...chain,
      iconUrl,
    })
  }
  const { chains, provider } = configureChains(chainsWithIcons, [
    alchemyProvider({ apiKey: process.env.NEXT_PUBLIC_ALCHEMY_KEY }),
    publicProvider(),
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

  // Synapse client
  const synapseProviders: EthersProvider[] = []
  chains.map((chain) => {
    const rpc: EthersProvider = new JsonRpcProvider(
      chain.id === 7700
        ? 'https://mainnode.plexnode.org:8545'
        : chain.rpcUrls.default.http[0]
    )
    rpc['projectId'] = chain.id
    synapseProviders.push(rpc)
  })
  return (
    <WagmiConfig client={wagmiClient}>
      <RainbowKitProvider chains={chains} theme={darkTheme()}>
        <SynapseProvider
          chainIds={chains.map((chain) => chain.id)}
          providers={synapseProviders}
        >
          {process.env.NEXT_PUBLIC_ALCHEMY_KEY ? (
            <Component {...pageProps} />
          ) : (
            <div>Alchemy key not set{process.env.NEXT_PUBLIC_ALCHEMY_KEY}</div>
          )}
          <CustomToaster />
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}

export default App
