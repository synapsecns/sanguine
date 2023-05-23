---
sidebar_position: 2
---

# Examples

## Hello World

```typescript
import { JsonRpcProvider } from '@ethersproject/providers'
import { Provider } from '@ethersproject/abstract-provider'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { BigNumber } from '@ethersproject/bignumber'

//Set up providers (RPCs) for each chain desired**
const arbitrumProvider: Provider = new JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)
const avalancheProvider: Provider = new JsonRpcProvider(
  'https://api.avax.network/ext/bc/C/rpc'
)

//Structure arguments properly
const chainIds = [42161, 43114]
const providers = [arbitrumProvider, avalancheProvider]

//Set up a SynapseSDK instance
const Synapse = new SynapseSDK(chainIds, providers)

// quote
const quotes = await Synapse.bridgeQuote(
  42161, // From Chain
  43114, // To Chain
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // From Token
  '0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664', // To Token
  BigNumber.from('20000000') // Amount
)

// bridge
const data = await Synapse.bridge(
  '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9', // To Address
  42161, // From Chain
  43114, // To Chain
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // To token Address
  BigNumber.from('20000000'), // Amount
  quotes.originQuery, // Origin query from bridgeQuote()
  quotes.destQuery // Origin query from bridgeQuote()
)

// execute the transaction (will trigger the wallet connected)
const transactionResponse = await arbitrumProvider.sendTransaction(data)
try {
  await tx.wait()
  console.log(`Transaction mined successfully: ${tx.hash}`)
  return tx
} catch (error) {
  console.log(`Transaction failed with error: ${error}`)
}
```

## Real world example for setting up SDK

using next.js, wagmi, and rainbowkit

```typescript
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
  publicProvider({ stallTimeout: 1_000 }),
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

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <WagmiConfig client={wagmiClient}>
      <RainbowKitProvider chains={chains} theme={darkTheme()}>
        <SynapseProvider
          chainIds={chains.map((chain) => chain.id)}
          providers={synapseProviders}
        >
          <Component {...pageProps} />
          <CustomToaster />
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}

export default App
```
