---
sidebar_label: Sample Code
---

# Sample Code

Example SDK & API implementations

## Basic Implementation

```js
// app.js

import { JsonRpcProvider } from '@ethersproject/providers'
import { Provider } from '@ethersproject/abstract-provider'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { BigNumber } from '@ethersproject/bignumber'

//Set up providers (RPCs) for each chain desired**
const arbitrumProvider: Provider = new JsonRpcProvider('https://arb1.arbitrum.io/rpc')
const avalancheProvider: Provider = new JsonRpcProvider('https://api.avax.network/ext/bc/C/rpc')

//Structure arguments properly
const chainIds = [42161, 43114]
const providers = [arbitrumProvider, avalancheProvider]

//Set up a SynapseSDK instance
const Synapse = new SynapseSDK(chainIds, providers)

// quote
const quote = await Synapse.bridgeQuote(
  42161, // From Chain
  43114, // To Chain
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // From Token
    '0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664', // To Token
    BigNumber.from('20000000'), // Amount
        {
          deadline: 1234567890,
          excludedModules: ['SynapseRFQ'],
          originUserAddress: '0x1234567890abcdef1234567890abcdef12345678',
        }
)

// bridge
await Synapse.bridge(
  '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9', // To Address
  quote.routerAddress, // bridge router contract address
  42161, // Origin Chain
  43114, // Destination Chain
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // Origin Token Address
  BigNumber.from('20000000'), // Amount
  quote.originQuery, // Origin query from bridgeQuote()
  quote.destQuery // Destination query from bridgeQuote()
)
```

## NextJS Implementation

:::note Dependencies

Wagmi, RainbowKit

:::

### App

```tsx
// app.tsx

import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import { SynapseProvider } from '@/utils/SynapseProvider'
import type { AppProps } from 'next/app'
import { Provider as EthersProvider } from '@ethersproject/abstract-provider'
import { JsonRpcProvider } from '@ethersproject/providers'
import { configureChains, createClient, WagmiConfig } from 'wagmi'
import {
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
} from 'wagmi/chains'

import {
  getDefaultWallets,
  RainbowKitProvider,
  darkTheme,
} from '@rainbow-me/rainbowkit'
import { alchemyProvider } from 'wagmi/providers/alchemy'
import { publicProvider } from 'wagmi/providers/public'

export default function App({ Component, pageProps }: AppProps) {
  const { chains, provider } = configureChains([mainnet,
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
  polygon], [
    alchemyProvider({ apiKey: 'API_KEY' }),
    publicProvider(),
  ])

  const { connectors } = getDefaultWallets({
    appName: 'ExampleApp',
    chains,
  })

  const wagmiClient = createClient({
    autoConnect: true,
    connectors,
    provider,
  })

  // Synapse client params setup example
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
```

### Provider

```tsx
// `@/utils/SynapseProvider.tsx`

import { createContext, useContext } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { Provider } from '@ethersproject/abstract-provider'

export const SynapseContext = createContext<SynapseSDK>(null)

export const SynapseProvider = ({
  children,
  chainIds,
  providers,
}: {
  children: React.ReactNode
  chainIds: number[]
  providers: Provider[]
}) => {
  const sdk = new SynapseSDK(chainIds, providers)
  return (
    <SynapseContext.Provider value={sdk}>{children}</SynapseContext.Provider>
  )
}

export const useSynapseContext = () => useContext(SynapseContext)
```

### Homepage

```tsx
// `/homepage/index.tsx`

import { useSynapseContext } from '@/utils/SynapseProvider'
import { Zero } from '@ethersproject/constants'
import { useState, useEffect } from 'react'
import { getNetwork } from '@wagmi/core'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
} from '@/constants/bridge'

export default function HomePage({ address }: { address: `0x${string}` }) {
  // Get the synapse sdk
  const SynapseSDK = useSynapseContext()
  
  // Get the current time
  const time = // add logic to get the current unix timestamp
  // Example state hooks
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [amount, setAmount] = useState(Zero)
  const [bridgeQuote, setBridgeQuote] = useState({
    outputAmountString: '',
    quotes: { originQuery: null, destQuery: null },
  })

  // Set connected network when component is mounted
  useEffect(() => {
      const { chain: fromChainIdRaw } = getNetwork()
      setFromChainId(fromChainIdRaw ? fromChainIdRaw?.id : DEFAULT_FROM_CHAIN)
    }, [])

  // Get Quote function
  // Suggestion: this function should be triggered from an useEffect when amount or to/from token/chain is altered
  const getQuote = async () = {
    SynapseSDK.bridgeQuote(
      fromChainId,
      toChainId,
      fromToken,
      toToken,
      amount,
      {
        deadline: time + 10000,
        excludedModules: [],
        originUserAddress: address,
      }
    )
    .then(
      ({ feeAmount, bridgeFee, maxAmountOut, originQuery, destQuery }) => {
        let toValueBigNum = maxAmountOut ?? Zero
        let toValueBase = toValueBigNum.div(toDecimals).toString()
        let toValueMantissa = toValueBigNum.mod(toDecimals).toString()

        setBridgeQuote({
          outputAmountString: toValueBase + '.' + toValueMantissa,
          quotes: {
            originQuery,
            destQuery,
          },
        })
        // do something
      }
    )
    .catch((err) => {
      alert('error getting quote', err)
      // do something
    })

  }

  // Execute bridge function
  const executeBridge = async () = {
    await Synapse.bridge(
      toAddress, // To Address
      bridgeQuote.routerAddress, // bridge router contract address
      fromChainId, // Origin Chain
      toChainId, // Destination Chain
      fromToken, // Origin Token Address
      amount, // Amount
      bridgeQuote.quotes.originQuery, // Origin query from bridgeQuote()
      bridgeQuote.quotes.destQuery // Destination query from bridgeQuote()
    ).then(({to, data}) => {
        // do something
      }
    )
    .catch((err) => {
      alert('error bridging', err)
      // do something
     })
  }

// ...

}
```

## API Functions

### Estimate bridge output

```js
async function estimateBridgeOutput(
  fromChain,
  toChain,
  fromToken,
  toToken,
  amountFrom
) {
  const query_string = `fromChain=${fromChain}&toChain=${toChain}&fromToken=${fromToken}&toToken=${toToken}&amountFrom=${amountFrom}`;
  const response = await fetch(
    `https://api.synapseprotocol.com/bridge?${query_string}`
  );

  const response_json = await response.json();
  // Check if the response is an array and has at least one item
  if (Array.isArray(response_json) && response_json.length > 0) {
    return response_json[0]; // Return the first item
  } else {
    throw new Error('No bridge quotes available');
  }
}

estimateBridgeOutput(
  1,     // Ethereum
  42161, // Arbitrum
  "USDC",
  "USDC",
  "1000"
).then(firstQuote => {
  console.log('First bridge quote:', firstQuote);
}).catch(error => {
  console.error('Error:', error.message);
});
```

### Generate unsigned bridge transaction

```js
async function generateUnsignedBridgeTxn(
  fromChain,
  toChain,
  fromToken,
  toToken,
  amountFrom,
  destAddress
) {
  const query_string = `fromChain=${fromChain}&toChain=${toChain}&fromToken=${fromToken}&toToken=${toToken}&amount=${amountFrom}&destAddress=${addressTo}`;
  const response = await fetch(
    `https://api.synapseprotocol.com/bridgeTxInfo?${query_string}`
  );
  const response_json = await response.json();
  return await response_json;
}

generateUnsignedBridgeTxn(
  1,     // Ethereum
  42161, // Arbitrum
  "USDC",
  "USDC",
  "1000"
  "0x2D2c027E0d1A899a1965910Dd272bcaE1cD03c22"
);
```
