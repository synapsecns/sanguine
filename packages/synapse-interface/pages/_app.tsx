import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import { SynapseProvider } from '@/utils/SynapseProvider'
import type { AppProps } from 'next/app'
import { Provider as EthersProvider } from '@ethersproject/abstract-provider'
import { JsonRpcProvider } from '@ethersproject/providers'
import {
  klaytn,
  boba,
  cronos,
  dfk,
  dogechain,
} from '@constants/extraWagmiChains'

import { configureChains, createClient, WagmiConfig } from 'wagmi'
import {
  mainnet,
  arbitrum,
  aurora,
  avalanche,
  bsc,
  canto,
  celo,
  fantom,
  harmonyOne,
  metis,
  moonbeam,
  moonriver,
  optimism,
  polygon,
} from 'wagmi/chains'

import {
  getDefaultWallets,
  RainbowKitProvider,
  darkTheme,
} from '@rainbow-me/rainbowkit'
import { alchemyProvider } from 'wagmi/providers/alchemy'
import { publicProvider } from 'wagmi/providers/public'
import { CHAIN_INFO_MAP } from '@constants/networks'
export default function App({ Component, pageProps }: AppProps) {
  let rawChains = [
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
  let chainsWithIcons: any[] = []
  for (let chain of rawChains) {
    chainsWithIcons.push({
      ...chain,
      iconUrl: CHAIN_INFO_MAP[chain.id].chainImg.src,
    })
  }
  const { chains, provider } = configureChains(chainsWithIcons, [
    alchemyProvider({ apiKey: '_UFN4P3jhI9zYma6APzoKX5aqKKadp2V' }),
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
  let synapseProviders: EthersProvider[] = []
  chains.map((chain) => {
    let rpc: EthersProvider = new JsonRpcProvider(chain.rpcUrls.default.http[0])
    synapseProviders.push(rpc)
  })
  return (
    <WagmiConfig client={wagmiClient}>
      <RainbowKitProvider chains={chains} theme={darkTheme()}>
        <SynapseProvider
          chainIds={chains.map((chain) => chain.id)}
          providers={synapseProviders}
        >
          <Component {...pageProps} />
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}
