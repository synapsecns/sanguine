import _ from 'lodash'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useRef, useState, useEffect, useMemo } from 'react'
import { Zero } from '@ethersproject/constants'
import { Token } from '@utils/classes/Token'
import { BigNumber } from '@ethersproject/bignumber'
import { BigintIsh } from '@synapsecns/sdk-router'
import { useSynapseContext } from '@/utils/SynapseProvider'
import { parseUnits, formatUnits } from '@ethersproject/units'
import { checksumAddress } from '@utils/checksum'
import { PageHeader } from '@components/PageHeader'
import Image from 'next/image'
import synLogo from '@assets/icons/synapse.svg'

// import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import {
  fetchBalance,
  switchNetwork,
  getNetwork,
  getAccount,
} from '@wagmi/core'

import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { ChainId } from '@constants/networks'
import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import {
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_TYPES_BY_CHAIN,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  BRIDGABLE_TOKENS,
  tokenSymbolToToken,
  // BRIDGE_SWAPABLE_TOKENS_BY_CHAIN,
} from '@constants/tokens'

import {
  DEFAULT_SWAPABLE_TYPE,
  DEFAULT_FROM_CHAIN,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
} from '@/constants/bridge'

import { checkCleanedValue } from '@utils/checkCleanedValue'
import { sanitizeValue } from '@utils/sanitizeValue'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'

console.log('BRIDGE_CHAINS_BY_TYPE', BRIDGE_CHAINS_BY_TYPE)
console.log('BRIDGABLE_TOKENS', BRIDGABLE_TOKENS)
console.log('BRIDGE_SWAPABLE_TOKENS_BY_TYPE', BRIDGE_SWAPABLE_TOKENS_BY_TYPE)
const bridgeFee = BigNumber.from('10000')
export default function BridgePage({ address }: { address: `0x${string}` }) {
  // move to utils

  const router = useRouter()

  // Get SynapseSDK
  const SynapseSDK = useSynapseContext()

  // Set current chain
  // const getCurrentNetwork = () => {
  //   const { chain: fromChainIdRaw } = getNetwork()
  //   return fromChainIdRaw ? fromChainIdRaw?.id : DEFAULT_FROM_CHAIN
  // }
  useEffect(() => {
    const { chain: fromChainIdRaw } = getNetwork()
    setFromChainId(fromChainIdRaw ? fromChainIdRaw?.id : DEFAULT_FROM_CHAIN)
  }, [])
  const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const chainIdsRef = useRef({ fromChainId, toChainId })

  // Init token

  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)

  const [fromTokens, setFromTokens] = useState([])
  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')
  const [init, setInit] = useState(false)
  // Auxiliary data
  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)
  const [feeAmount, setFeeAmount] = useState(BigNumber.from('10000'))
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const [toBridgeableTokens, setToBridgeableTokens] = useState(
    BRIDGABLE_TOKENS[DEFAULT_TO_CHAIN]
  )

  // TODO set to chain as a idx
  const [toBridgeableChains, setToBridgeableChains] = useState(
    BRIDGE_CHAINS_BY_TYPE[String(DEFAULT_FROM_TOKEN.swapableType)].filter(
      (chain) => Number(chain) !== DEFAULT_FROM_CHAIN
    )
  )

  const [bridgeQueries, setBridgeQueries] = useState({
    originQuery: null,
    destQuery: null,
  })

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
    console.log(
      'dslajldkjasal',
      router.isReady,
      toChainIdUrl,
      fromTokenSymbolUrl,
      toTokenSymbolUrl
    )

    let tempFromToken: Token = getMostCommonSwapableType(fromChainId)

    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        tempFromToken = token
      }
    }
    const { bridgeableToken, newToChain, bridgeableTokens, bridgeableChains } =
      handleNewFromTokenNew(
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
    let fromChainTokensByType = Object.values(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId]
    )
    let maxTokenLength = 0
    console.log('chainIdchainIdchainIdchainIdchainId', chainId)
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
    setPriceImpact(Zero)
    setExchangeRate(Zero)
  }
  const onChangeFromAmount = (value: string) => {
    if (
      !(
        value.split('.')[1]?.length >
        fromToken[fromChainId as keyof Token['decimals']]
      )
    ) {
      setFromValue(value)
    }
  }

  const onChangeToAmount = (value: string) => {
    if (
      !(
        value.split('.')[1]?.length >
        toToken[toChainId as keyof Token['decimals']]
      )
    ) {
      setToValue(value)
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
          outputChain: outputChain,
          inputCurrency: inputCurrency,
          outputCurrency: outputCurrency,
        },
      },
      undefined,
      { shallow: true }
    )
  }

  const handleNewFromTokenNew = (
    token: Token,
    positedToChain: number | undefined,
    positedToSymbol: string | undefined,
    fromChainId: number
  ) => {
    console.log('HIIEEIEI', token)
    let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN
    let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
      String(token.swapableType)
    ].filter((chainId) => Number(chainId) !== fromChainId)
    let swapExceptionsArr: number[] =
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

    console.log('NEW TO CHAIN', newToChain)
    let positedToToken = positedToSymbol
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
      bridgeableToken: bridgeableToken,
      newToChain: newToChain,
      bridgeableTokens: bridgeableTokens,
      bridgeableChains: bridgeableChains,
    }
  }

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = async () => {
    // let oldFromChainId = fromChainId
    // let oldToChainId = toChainId
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      let oldFromChain = fromChainId
      let res = switchNetwork({ chainId: toChainId })
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
      // try {
      //   let res = await switchNetwork({ chainId: toChainId })
      //   if (res === undefined) {
      //     console.log("can't switch network", toChainId)
      //     return
      //   }
      // } catch (err) {
      //   console.log("can't switch network sir", err)
      //   return
      // }
      let bridgeableFromTokens: Token[] = sortByVisibilityRank(
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
      } = handleNewFromTokenNew(
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
      let res = switchNetwork({ chainId: chainId })
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

      let bridgeableFromTokens: Token[] = sortByVisibilityRank(
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
      } = handleNewFromTokenNew(tempFromToken, chainId, toToken.symbol, chainId)
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
      } = handleNewFromTokenNew(token, toChainId, toToken.symbol, fromChainId)
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
      console.log('to token change', token)
      setToToken(token)
      updateUrlParams({
        outputChain: toChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: token.symbol,
      })
    }
  }

  const triggerRateAndImpact = ({
    amountToGive,
    amountToReceive,
    bridgeFee,
  }: {
    amountToGive: BigNumber
    amountToReceive: BigNumber
    bridgeFee: BigNumber
  }) => {
    setFeeAmount(bridgeFee)
    let umom = calculateExchangeRate(
      amountToGive.sub(
        feeAmount.div(
          BigNumber.from(10).pow(18 - fromToken.decimals[fromChainId])
        )
      ),
      fromToken.decimals[fromChainId],
      amountToReceive,
      toToken.decimals[toChainId]
    )
    console.log('umom', umom)
    setExchangeRate(umom)
  }
  const getQuote = async () => {
    let amount: BigintIsh = BigNumber.from(Number(fromValue) * 1000000)
    const quotes = await SynapseSDK.bridgeQuote(
      fromChainId, // From Chain
      toChainId, // To Chain
      fromToken.addresses[fromChainId].toLowerCase(), // From token Address
      toToken.addresses[toChainId].toLowerCase(), // To token Address
      amount // Amount in
    )
    console.log('quotesquotesquotesquotes', quotes)

    let maxAmountOut = quotes?.destQuery?.minAmountOut
      ? quotes.destQuery.minAmountOut
      : Zero
    // setToValue(maxAmountOut.div(100000000).toString())
    // console.log(maxAmountOut.toString(), quotes.originQuery, quotes.destQuery)
    return {
      amountToReceive: maxAmountOut,
      originQuery: quotes.originQuery,
      destQuery: quotes.destQuery,
    }
    // return {quotes.maxAmountOut, quotes.originQuery, quotes.destQuery}
  }
  const calculateBridgeAmount = async () => {
    let cleanedFromValue = sanitizeValue(fromValue)
    if (checkCleanedValue(cleanedFromValue)) {
      setToValue('')
      return
    }

    if (
      !(
        fromChainId &&
        toChainId &&
        String(fromToken.addresses[fromChainId]) &&
        String(toToken.addresses[toChainId]) &&
        fromValue
      )
    ) {
      return
    }
    const amountToGive = parseUnits(
      cleanedFromValue,
      fromToken.decimals[fromChainId]
    )

    const { amountToReceive, originQuery, destQuery } = await getQuote()
    console.log(
      ':SDLKDKSJGDKJHSDKJ',
      amountToGive.toString(),
      amountToReceive.toString(),
      typeof amountToGive
    )
    // // setToValue(amountToReceive.toString())
    // if (sanitizeValue(fromRef.current?.value) == sanitizeValue(fromValue)) {
    //   setToValue(formatUnits(amountToReceive, toToken.decimals[toChainId]))
    //   triggerRateAndImpact({ amountToGive, amountToReceive, bridgeFee })
    // }
    // setBridgeQueries({ originQuery, destQuery })
  }

  useEffect(() => {
    if (fromToken && toToken) {
      calculateBridgeAmount()
    }
  }, [fromToken, toToken, fromValue, fromChainId, toChainId, feeAmount])

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
                    address={address}
                    fromChainId={fromChainId}
                    toChainId={toChainId}
                    onSelectFromChain={handleFromChainChange}
                    onSelectToChain={handleToChainChange}
                    swapFromToChains={handleChainFlip}
                    onChangeFromAmount={onChangeFromAmount}
                    onChangeToAmount={onChangeToAmount}
                    fromCoin={fromToken}
                    toCoin={toToken}
                    possibleChains={toBridgeableChains}
                    handleTokenChange={handleTokenChange}
                    toBridgeableTokens={toBridgeableTokens}
                    quotes={bridgeQueries}
                    fromTokens={fromTokens}
                    {...{
                      fromValue,
                      toValue,
                      error,
                      priceImpact,
                      exchangeRate,
                      feeAmount,
                      destinationAddress,
                      setDestinationAddress,
                    }}
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
