import _ from 'lodash'
import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { Zero } from '@ethersproject/constants'
import { Token } from '@/utils/types'
import { BigNumber } from '@ethersproject/bignumber'
import { BridgeQuote } from '@/utils/types'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { getNetwork, switchNetwork } from '@wagmi/core'
import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import {
  BRIDGABLE_TOKENS,
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  tokenSymbolToToken,
} from '@constants/tokens'

import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { stringToBigNum } from '@/utils/stringToBigNum'
import BridgeCard from './BridgeCard'
import { useSynapseContext } from '@/utils/SynapseProvider'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_CHAIN,
  DEFAULT_TO_TOKEN,
} from '@/constants/bridge'
// import BridgeWatcher from './BridgeWatcher'

export default function BridgePage({ address }: { address: `0x${string}` }) {
  const router = useRouter()
  const SynapseSDK = useSynapseContext()

  const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [fromTokens, setFromTokens] = useState([])
  const [fromInput, setFromInput] = useState({ string: '', bigNum: Zero })
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const [toBridgeableTokens, setToBridgeableTokens] = useState(
    BRIDGABLE_TOKENS[DEFAULT_TO_CHAIN]
  )
  const [toBridgeableChains, setToBridgeableChains] = useState(
    BRIDGE_CHAINS_BY_TYPE[String(DEFAULT_FROM_TOKEN.swapableType)].filter(
      (chain) => Number(chain) !== DEFAULT_FROM_CHAIN
    )
  )
  const [bridgeQuote, setBridgeQuote] = useState<BridgeQuote>({
    outputAmount: Zero,
    outputAmountString: '',
    exchangeRate: Zero,
    feeAmount: Zero,
    delta: Zero,
    quotes: { originQuery: null, destQuery: null },
  })
  useEffect(() => {
    const { chain: fromChainIdRaw } = getNetwork()
    setFromChainId(fromChainIdRaw ? fromChainIdRaw?.id : DEFAULT_FROM_CHAIN)
  }, [])

  // Upon update from the url query, updates to according states
  // will only execute on initial load of the page
  useEffect(() => {
    if (!router.isReady) {
      return
    }
    const {
      outputChain: toChainIdUrl,
      inputCurrency: fromTokenSymbolUrl,
      outputCurrency: toTokenSymbolUrl,
    } = router.query

    let tempFromToken: Token = getMostCommonSwapableType(fromChainId)

    if (fromTokenSymbolUrl) {
      const token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        tempFromToken = token
      }
    }
    const { bridgeableToken, newToChain, bridgeableTokens, bridgeableChains } =
      handleNewFromToken(
        tempFromToken,
        toChainIdUrl ? Number(toChainIdUrl) : undefined,
        toTokenSymbolUrl ? String(toTokenSymbolUrl) : undefined,
        fromChainId
      )
    setToChainId(newToChain)
    setFromToken(tempFromToken)
    setToToken(bridgeableToken)
    setToBridgeableChains(bridgeableChains)
    setToBridgeableTokens(bridgeableTokens)
    updateUrlParams({
      outputChain: newToChain,
      inputCurrency: fromToken.symbol,
      outputCurrency: bridgeableToken.symbol,
    })
  }, [router.isReady])

  // Listens for every time the source chan is changed and ensures
  // that there is not a clash between the source and destination chain.
  useEffect(() => {
    if (fromChainId === undefined || address === undefined) {
      return
    }
    sortByTokenBalance(
      BRIDGABLE_TOKENS[fromChainId],
      fromChainId,
      address
    ).then((tokens) => {
      setFromTokens(tokens)
    })
  }, [fromChainId])

  const getMostCommonSwapableType = (chainId: number) => {
    const fromChainTokensByType = Object.values(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId]
    )
    let maxTokenLength = 0
    let mostCommonSwapableType: Token[] = fromChainTokensByType[0]
    fromChainTokensByType.map((tokenArr, i) => {
      if (tokenArr.length > maxTokenLength) {
        maxTokenLength = tokenArr.length
        mostCommonSwapableType = tokenArr
      }
    })

    return sortByVisibilityRank(mostCommonSwapableType)[0]
  }

  // Helpers
  const resetRates = () => {
    setBridgeQuote({
      outputAmount: Zero,
      outputAmountString: '',
      exchangeRate: Zero,
      feeAmount: Zero,
      delta: Zero,
      quotes: { originQuery: null, destQuery: null },
    })
  }
  const onChangeFromAmount = (value: string) => {
    if (
      !(
        value.split('.')[1]?.length >
        fromToken[fromChainId as keyof Token['decimals']]
      )
    ) {
      setFromInput({
        string: value,
        bigNum: stringToBigNum(value, fromToken.decimals[fromChainId]),
      })
    }
  }

  const updateUrlParams = ({
    outputChain,
    inputCurrency,
    outputCurrency,
  }: {
    outputChain: any
    inputCurrency: any
    outputCurrency: any
  }) => {
    router.replace(
      {
        pathname: BRIDGE_PATH,
        query: {
          outputChain,
          inputCurrency,
          outputCurrency,
        },
      },
      undefined,
      { shallow: true }
    )
  }

  const handleNewFromToken = (
    token: Token,
    positedToChain: number | undefined,
    positedToSymbol: string | undefined,
    fromChainId: number
  ) => {
    let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN
    let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
      String(token.swapableType)
    ].filter((chainId) => Number(chainId) !== fromChainId)
    const swapExceptionsArr: number[] =
      token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
    if (swapExceptionsArr?.length > 0) {
      bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
    }

    if (!bridgeableChains.includes(String(newToChain))) {
      newToChain =
        Number(bridgeableChains[0]) === fromChainId
          ? Number(bridgeableChains[1])
          : Number(bridgeableChains[0])
    }

    const positedToToken = positedToSymbol
      ? tokenSymbolToToken(newToChain, positedToSymbol)
      : tokenSymbolToToken(newToChain, token.symbol)

    let bridgeableTokens: Token[] = sortByVisibilityRank(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain][String(token.swapableType)]
    )

    if (swapExceptionsArr?.length > 0) {
      bridgeableTokens = bridgeableTokens.filter(
        (toToken) => toToken.symbol === token.symbol
      )
    }
    // bridgeableTokens = sortByVisibilityRank(bridgeableTokens)
    let bridgeableToken: Token = positedToToken
    if (!bridgeableTokens.includes(positedToToken)) {
      bridgeableToken = bridgeableTokens[0]
    }
    return {
      bridgeableToken,
      newToChain,
      bridgeableTokens,
      bridgeableChains,
    }
  }

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = async () => {
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      const oldFromChain = fromChainId
      const res = switchNetwork({ chainId: toChainId })
        .then((res) => {
          if (res === undefined) {
            console.log("can't switch network", toChainId)
            return
          }
          return res
        })
        .catch((err) => {
          console.log("can't switch network sir", err)
          return undefined
        })

      const bridgeableFromTokens: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][
          String(fromToken.swapableType)
        ]
      )
      let tempFromToken: Token = fromToken

      if (bridgeableFromTokens?.length > 0) {
        tempFromToken = getMostCommonSwapableType(fromChainId)
      }
      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromToken(
        tempFromToken,
        oldFromChain,
        toToken.symbol,
        toChainId
      )
      setFromChainId(toChainId)

      setToChainId(newToChain)
      setFromToken(tempFromToken)
      setToToken(bridgeableToken)
      setToBridgeableChains(bridgeableChains)
      setToBridgeableTokens(bridgeableTokens)
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: tempFromToken.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    }

    // resetRates()
  }

  // Changes destination change when the user changes the toChainId
  const handleFromChainChange = async (chainId: number) => {
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      const res = switchNetwork({ chainId })
        .then((res) => {
          if (res === undefined) {
            console.log("can't switch network", chainId)
            return
          }
          return res
        })
        .catch((err) => {
          console.log("can't switch network sir", err)
          return undefined
        })
      if (res === undefined) {
        console.log("can't switch network chainId", chainId)
        return
      }

      const bridgeableFromTokens: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId][String(fromToken.swapableType)]
      )
      let tempFromToken: Token = fromToken

      if (bridgeableFromTokens?.length > 0) {
        tempFromToken = getMostCommonSwapableType(chainId)
      }
      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromToken(tempFromToken, chainId, toToken.symbol, chainId)

      setFromChainId(chainId)
      setToChainId(newToChain)
      setFromToken(fromToken)
      setToToken(bridgeableToken)
      setToBridgeableChains(bridgeableChains)
      setToBridgeableTokens(bridgeableTokens)
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: tempFromToken.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    }
  }

  const handleToChainChange = (chainId: number) => {
    setToChainId(chainId)
    updateUrlParams({
      outputChain: chainId,
      inputCurrency: fromToken.symbol,
      outputCurrency: toToken.symbol,
    })
  }

  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    if (type == 'from') {
      console.log('from token change', token, token.swapableType, token.symbol)
      setFromToken(token)
      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromToken(token, toChainId, toToken.symbol, fromChainId)
      setToChainId(newToChain)
      setToToken(bridgeableToken)
      setToBridgeableChains(bridgeableChains)
      setToBridgeableTokens(bridgeableTokens)
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: fromToken.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    } else {
      setToToken(token)
      updateUrlParams({
        outputChain: toChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: token.symbol,
      })
    }
  }

  const getQuote = async () => {
    const toDecimals = BigNumber.from(10).pow(toToken.decimals[toChainId])

    SynapseSDK.bridgeQuote(
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId].toLowerCase(),
      toToken.addresses[toChainId].toLowerCase(),
      fromInput.bigNum
    )
      .then(
        ({ feeAmount, feeConfig, maxAmountOut, originQuery, destQuery }) => {
          const toValueBigNum = destQuery.minAmountOut
            ? destQuery.minAmountOut
            : Zero
          const toValueBase = toValueBigNum.div(toDecimals).toString()
          const toValueMantissa = toValueBigNum.mod(toDecimals).toString()

          setBridgeQuote({
            outputAmount: toValueBigNum,
            outputAmountString: toValueBase + '.' + toValueMantissa,
            exchangeRate: calculateExchangeRate(
              fromInput.bigNum,
              fromToken.decimals[fromChainId],
              toValueBigNum,
              toToken.decimals[toChainId]
            ),
            feeAmount,
            delta: maxAmountOut,
            quotes: {
              originQuery,
              destQuery,
            },
          })
        }
      )
      .catch((err) => {
        console.log('error getting quote', err)
      })
  }

  useEffect(() => {
    if (
      fromChainId &&
      toChainId &&
      String(fromToken.addresses[fromChainId]) &&
      String(toToken.addresses[toChainId]) &&
      fromInput
    ) {
      getQuote()
    }
  }, [fromToken, toToken, fromInput, fromChainId, toChainId])

  return (
    <LandingPageWrapper>
      <main className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none">
        <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <div>
            <Grid
              cols={{ xs: 1 }}
              gap={6}
              className="justify-center px-2 py-16 sm:px-6 md:px-8"
            >
              <div className="flex justify-center">
                <div className="pb-3 place-self-center">
                  <BridgeCard
                    error={error}
                    address={address}
                    bridgeQuote={bridgeQuote}
                    fromInput={fromInput}
                    fromToken={fromToken}
                    fromTokens={fromTokens}
                    fromChainId={fromChainId}
                    toToken={toToken}
                    toChainId={toChainId}
                    toBridgeableChains={toBridgeableChains}
                    toBridgeableTokens={toBridgeableTokens}
                    destinationAddress={destinationAddress}
                    handleChainFlip={handleChainFlip}
                    handleTokenChange={handleTokenChange}
                    onChangeFromAmount={onChangeFromAmount}
                    onSelectFromChain={handleFromChainChange}
                    onSelectToChain={handleToChainChange}
                    setDestinationAddress={setDestinationAddress}
                  />

                  <ActionCardFooter link={HOW_TO_BRIDGE_URL} />
                </div>
              </div>
              <div>{/* <BridgeWatcher /> */}</div>
            </Grid>
          </div>
        </div>
      </main>
    </LandingPageWrapper>
  )
}

// export function HarmonyCheck({ fromChainId, toChainId }: { fromChainId: number; toChainId: number }) {
//   return (
//     <>
//       {(toChainId === ChainId.HARMONY || fromChainId === ChainId.HARMONY) && (
//         <div
//           className={`bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-lg text-white`}
//         >
//           The native Harmony bridge has been exploited, which lead to a hard depeg of the following Harmony-specific tokens: 1DAI, 1USDC, 1USDT, 1ETH.
//           <br /> Please see the{' '}
//           <a
//             className="underline"
//             href="https://twitter.com/harmonyprotocol/status/1540110924400324608"
//           >
//             official Harmony Twitter
//           </a>{' '}
//           for status updates and exercise caution when interacting with Harmony.
//         </div>
//       )}
//     </>
//   )
// }
