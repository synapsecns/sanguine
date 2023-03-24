import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'

import type { AppProps } from 'next/app'
import {
  klaytn,
  boba,
  cronos,
  dfk,
  dogechain,
} from '@constants/extraWagmiChains'
import { Header } from '../components/layouts/Header'
import { Footer } from '../components/layouts/Footer'

import { configureChains, createClient, WagmiConfig } from 'wagmi'
import { Chain } from 'wagmi/chains'
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
  return (
    <WagmiConfig client={wagmiClient}>
      <RainbowKitProvider chains={chains} theme={darkTheme()}>
        <Component {...pageProps} />
        <Footer />
      </RainbowKitProvider>
    </WagmiConfig>
  )
}
