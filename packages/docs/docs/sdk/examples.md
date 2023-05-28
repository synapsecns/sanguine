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

## Setting up SDK

using next.js, wagmi, and rainbowkit

```typescript
// _app.tsx
import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
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
import { SynapseProvider } from '@/utils/providers/SynapseProvider'

const allChains = [
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
]

const { chains, provider } = configureChains(allChains, [
  jsonRpcProvider({
    rpc: (chain) => ({
      http: chain.rpcUrls.default.http[0],
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
          <Component {...pageProps} />
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}

export default App
```

```typescript
// SynapseProvider.tsx
import { SynapseSDK } from '@synapsecns/sdk-router'
import { Provider } from '@ethersproject/abstract-provider'
import { createContext, useContext, memo, useMemo } from 'react'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { Provider as EthersProvider } from '@ethersproject/abstract-provider'

export const SynapseContext = createContext(null)

export const SynapseProvider = memo(
  ({ children, chains }: { children: React.ReactNode; chains: any[] }) => {
    const synapseProviders = useMemo(() => {
      return chains.map(
        (chain) => new StaticJsonRpcProvider(chain.configRpc, chain.id)
      )
    }, [chains])

    const providerMap = useMemo(() => {
      return chains.reduce((map, chain) => {
        map[chain.id] = synapseProviders.find(
          (provider) => provider.connection.url === chain.configRpc
        )
        return map
      }, {})
    }, [chains, synapseProviders])

    const chainIds = chains.map((chain) => chain.id)
    const synapseSDK = useMemo(
      () => new SynapseSDK(chainIds, synapseProviders),
      [chainIds, synapseProviders]
    )

    return (
      <SynapseContext.Provider value={{ synapseSDK, providerMap }}>
        {children}
      </SynapseContext.Provider>
    )
  }
)

export const useSynapseContext = () => useContext(SynapseContext)
```

## Bridging Example

Assumes that the sdk is set up in manner shown in the previous example.

```typescript
//homepage/index.tsx
import { useSynapseContext } from '@/utils/SynapseProvider'
import { Zero } from '@ethersproject/constants'
import { useState, useEffect } from 'react'
import { fetchSigner } from '@wagmi/core'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
} from '@/constants/bridge'

export default function HomePage({
  address,
  fromChainId,
}: {
  address: `0x${string}`
  fromChainId: number
}) {
  // Get the synapse sdk from the provider (see previous example)
  const SynapseSDK = useSynapseContext()

  // Example state hooks
  const [fromTokenAddress, setFromTokenAddress] = useState(DEFAULT_FROM_TOKEN)
  const [toTokenAddress, setToTokenAddress] = useState(DEFAULT_TO_TOKEN)
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [amount, setAmount] = useState(Zero)
  const [bridgeQuote, setBridgeQuote] = useState({
    outputAmountString: '',
    quotes: { originQuery: null, destQuery: null },
  })

  // Get Quote function
  // Suggestion: this function should be triggered from an useEffect when amount or to/from token/chain is altered
  const getQuote = async () => {
    SynapseSDK.bridgeQuote(
      fromChainId,
      toChainId,
      fromTokenAddress,
      toTokenAddress,
      amount
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
  const executeBridge = async () => {
    try {
      const wallet = await fetchSigner({
        chainId: fromChainId,
      })

      // const adjustedFrom = subtractSlippage(fromInput.bigNum, 'ONE_TENTH', null)
      const data = await synapseSDK.bridge(
        toAddress, // To Address
        fromChainId, // From Chain
        toChainId, // To Chain
        fromTokenAddress, // From token Address
        amount, // Amount to bridge
        quotes.originQuery, // Origin query from bridgeQuote()
        quotes.destQuery // Origin query from bridgeQuote()
      )

      // Add the value to the transaction payload if the from token is the native token. (SDK does not automatically add as of now)
      const payload =
        fromTokenAddress === AddressZero || fromTokenAddress === ''
          ? { data: data.data, to: data.to, value: fromInput.bigNum }
          : data

      const tx = await wallet.sendTransaction(payload)
      try {
        await tx.wait()
        console.log(`Transaction mined successfully: ${tx.hash}`)
        setBridgeTxHash(tx.hash)
        return tx
      } catch (error) {
        console.log(`Transaction failed with error: ${error}`)
      }
    } catch (error) {
      console.log('Error executing bridge', error)
      return
    }
  }

  // ...
}
```

## Swap Example

```typescript
//swap/index.tsx
import { useSynapseContext } from '@/utils/SynapseProvider'
import { Zero } from '@ethersproject/constants'
import { useState, useEffect } from 'react'
import { fetchSigner } from '@wagmi/core'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
} from '@/constants/swap'

export default function SwapPage({ address, fromChainId }: { address: `0x${string}`, fromChainId: number }) {
	// Get the synapse sdk from the provider (see previous example)
	const SynapseSDK = useSynapseContext()

	// Example state hooks
	const [fromTokenAddress, setFromTokenAddress] = useState(DEFAULT_FROM_TOKEN)
	const [toTokenAddress, setToTokenAddress] = useState(DEFAULT_TO_TOKEN)
	const [amount, setAmount] = useState(Zero)
	const [swapQuote, setSwapQuote] = useState({
    outputAmountString: '',
    quote: null,
  })


	// Get swap quote function
	// Suggestion: this function should be triggered from an useEffect when amount or to/from token or the connected chain is altered
	const getSwapQuote = async () = {
		SynapseSDK.bridgeQuote(
		      fromChainId,
          fromTokenAddress,
          toTokenAddress,
          amount
		    ).then(
		        ({ outerAddress, maxAmountOut, query }) => {
		          let toValueBigNum = maxAmountOut ?? Zero
		          const allowance =
                fromToken.addresses[connectedChainId] === AddressZero ||address === undefined
                  ? Zero
                  : await getCurrentTokenAllowance(routerAddress)


		          setSwapQuote({
		            outputAmountString: formatBNToString(toValueBigNum,toToken.decimals[connectedChainId], 8),
		            quote: query,
		          })
				// do something
		        }
		      )
		      .catch((err) => {
		        alert('error getting quote', err)
			  // do something
		      })

	}


	// Execute swap function
  const executeSwap = async () => {
    try {
      const wallet =  await fetchSigner({
        chainId: fromChainId,
      })

      const data = await synapseSDK.swap(
        connectedChainId,
        address,
        fromTokenAddress,
        amount,
        quote
      )

      // Add the value to the transaction payload if the from token is the native token. (SDK does not automatically add as of now)
      const payload =
        fromTokenAddress === AddressZero || fromTokenAddress === ''
          ? { data: data.data, to: data.to, value: fromInput.bigNum }
          : data
      const tx = await wallet.sendTransaction(payload)

      try {
        await tx.wait()
        console.log(`Transaction mined successfully: ${tx.hash}`)
        return tx
      } catch (error) {
        console.log(`Transaction failed with error: ${error}`)
      }
    } catch (error) {
      console.log(`Swap Execution failed with error: ${error}`)
    }
  }


// ...

}

```

## Helpful Utils

For the examples above

### formatBNToString()

```typescript
import { formatUnits, commify } from '@ethersproject/units'

export const formatBNToString = (
  bn: BigNumber,
  nativePrecison: number,
  decimalPlaces?: number
) => {
  const fullPrecision = formatUnits(bn, nativePrecison)
  const decimalIdx = fullPrecision.indexOf('.')

  if (decimalPlaces === undefined || decimalIdx === -1) {
    return fullPrecision
  } else {
    const rawNumber = Number(fullPrecision)

    if (rawNumber === 0) {
      return rawNumber.toFixed(1)
    } else {
      return rawNumber.toFixed(decimalPlaces)
    }
  }
}
```

### getCurrentTokenAllowance()

```typescript
import { erc20ABI } from 'wagmi'
import { fetchSigner } from '@wagmi/core'
import { Contract } from 'ethers'
const getCurrentTokenAllowance = async (
  routerAddress: string,
  connectedChainId: number
) => {
  const wallet = await fetchSigner({
    chainId: connectedChainId,
  })

  const erc20 = new Contract(fromTokenAddress, erc20ABI, wallet)
  const allowance = await erc20.allowance(address, routerAddress)
  console.log('allowance from getCurrentTokenAllowance: ', allowance)
  return allowance
}
```

### approveToken()

```typescript
import { erc20ABI } from 'wagmi'
import { fetchSigner } from '@wagmi/core'
import { Contract } from 'ethers'
import { MaxInt256 } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
export const approveToken = async (
  address: string,
  chainId: number,
  tokenAddress: string,
  amount?: BigNumber
) => {
  const signer = await fetchSigner({
    chainId,
  })

  const erc20 = new Contract(tokenAddress, erc20ABI, signer)
  try {
    const approveTx = await erc20.approve(address, amount ?? MaxInt256)

    await approveTx.wait()
    return approveTx
  } catch (error) {
    console.log(`Transaction failed with error: ${error}`)
  }
}
```
