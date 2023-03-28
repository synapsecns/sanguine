import _ from 'lodash'
import Grid from '@tw/Grid'
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
// import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { fetchBalance } from '@wagmi/core'

import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { ChainId } from '@constants/networks'
import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import {
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_TYPES_BY_CHAIN,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  BRIDGABLE_TOKENS,
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
export default function BridgePage() {
  // move to utils
  const sortByVisibilityRank = (tokens: Token[]) => {
    return Object.values(tokens).sort(
      (a, b) => b.visibilityRank - a.visibilityRank
    )
  }
  // move to utils
  const sortByTokenBalance = async (
    tokens: Token[],
    chainId: number,
    address: any
  ) => {
    let i = 0
    let tokensWithBalances: any[] = []
    let zeroTokensWithBalances: any[] = []
    // go through all tokens and retrieve token balances
    while (i < tokens.length) {
      if (chainId === undefined || address === undefined) {
        tokensWithBalances.push({
          token: tokens[i],
          balance: Zero,
        })
        i++
        continue
      }
      let tokenAddr = tokens[i].addresses[chainId as keyof Token['addresses']]

      let rawTokenBalance: any
      // Check for native token
      if (tokenAddr === '') {
        const data = await fetchBalance({
          address: address,
          chainId: chainId,
        })
        rawTokenBalance = data
      } else if (tokenAddr?.length > 0) {
        const data = await fetchBalance({
          address: address,
          token: `0x${tokenAddr.slice(2)}`,
          chainId: chainId,
        })
        rawTokenBalance = data
      }

      // manages two the array of tokens with zero balances and non-zero balances
      if (rawTokenBalance) {
        if (rawTokenBalance?.value._hex !== '0x00') {
          zeroTokensWithBalances.push({
            token: tokens[i],
            balance: rawTokenBalance.value,
          })
        } else {
          tokensWithBalances.push({
            token: tokens[i],
            balance: rawTokenBalance.value,
          })
        }
      }
      i++
    }
    console.log(
      'zeroTokensWithBalances',
      zeroTokensWithBalances,
      'tokensWithBalances',
      tokensWithBalances
    )
    let tokenList = sortByVisibilityRank(zeroTokensWithBalances).concat(
      sortByVisibilityRank(tokensWithBalances)
    )
    console.log('tokenBalances', tokenList)
    return tokenList
  }

  // move to utils
  const tokenSymbolToToken = (chainId: number, symbol: string) => {
    const token = _.find(BRIDGABLE_TOKENS[chainId], (token) => {
      return token.symbol === symbol
    })
    return token
  }
  const router = useRouter()

  // Get data from wagmi.
  const { address } = useAccount()

  // Get SynapseSDK
  const SynapseSDK = useSynapseContext()

  // Set current chain
  const { chain: fromChainIdRaw } = useNetwork()
  const fromChainId = fromChainIdRaw?.id
    ? fromChainIdRaw?.id
    : DEFAULT_FROM_CHAIN
  console.log('CHAIN', fromChainId)
  console.log('RELOAD')

  const { error: networkSwitchError, switchNetwork } = useSwitchNetwork()

  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)

  // TODO need to implemement last chain id to handle case when user cancels switching the chain.
  // probably best to make tochainId a dictionary that holds the current and last chain id.
  const [lastToChainId, setLastToChainId] = useState(DEFAULT_TO_CHAIN)

  // Init token
  const [fromTokenIdx, setFromTokenIdx] = useState(0)
  const [toTokenIdx, setToTokenIdx] = useState(0)

  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)

  const [fromTokens, setFromTokens] = useState([])
  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')
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

  // Upon update of connected chain or wallet address, the bridgeable tokens from the connected
  // chain are updated along with the user's wallet balance for each token.
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
      setFromTokenIdx(0)
    })
  }, [fromChainId, address])

  // Upon update from the url query, updates to according states
  // will only execute on initial load of the page
  useEffect(() => {
    console.log('dslajldkjasal', router.isReady)
    if (!router.isReady) {
      // if the router is not ready, exit
      return
    }

    const {
      outputChain: toChainIdUrl,
      inputCurrency: fromTokenSymbolUrl,
      outputCurrency: toTokenSymbolUrl,
    } = router.query

    let tempFromToken: Token = getMostCommonSwapableType(fromChainId)
    console.log('tempFromToken', tempFromToken)
    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        tempFromToken = token
      }
    }
    console.log('tempFromToken', tempFromToken)
    const { bridgeableToken, newToChain, bridgeableTokens, bridgeableChains } =
      handleNewFromTokenNew(
        tempFromToken,
        toChainIdUrl ? Number(toChainIdUrl) : undefined,
        toTokenSymbolUrl ? String(toTokenSymbolUrl) : undefined
      )
    setFromToken(tempFromToken)
    setToBridgeableChains(bridgeableChains)
    setToBridgeableTokens(bridgeableTokens)
    setToToken(bridgeableToken)
    setToChainId(newToChain)

    // Update url params if any passed params were invalid
    if (
      Number(toChainIdUrl) !== newToChain ||
      fromTokenSymbolUrl !== tempFromToken.symbol ||
      toTokenSymbolUrl !== bridgeableToken.symbol
    ) {
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: tempFromToken.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    }
  }, [router.isReady])

  // Listens for every time the source chan is changed and ensures
  // that there is not a clash between the source and destination chain.

  // check if resetting from token and such is necceary.
  useEffect(() => {
    resetRates()
    if (fromChainId === toChainId) {
      setToChainId(lastToChainId)
      updateUrlParams({
        outputChain: lastToChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: toToken.symbol,
      })
    }
  }, [fromChainId])

  // useEffect(() => {
  //   if (fromToken === undefined) {
  //     return BRIDGABLE_TOKENS[fromChainId]
  //   }

  //   let positedToChain = toChainId
  //   let token = fromToken
  //   let positedToSymbol = undefined
  //   let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN

  //   let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
  //     String(token.swapableType)
  //   ].filter((chainId) => Number(chainId) !== fromChainId)
  //   let swapExceptionsArr: number[] =
  //     token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
  //   if (swapExceptionsArr?.length > 0) {
  //     bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
  //   }

  //   // TODO filter above
  //   if (!bridgeableChains.includes(String(newToChain))) {
  //     newToChain =
  //       Number(bridgeableChains[0]) === fromChainId
  //         ? Number(bridgeableChains[1])
  //         : Number(bridgeableChains[0])
  //   }
  //   let positedToToken = positedToSymbol
  //     ? tokenSymbolToToken(newToChain, positedToSymbol)
  //     : undefined

  //   let bridgeableTokens: Token[] =
  //     BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain][String(token.swapableType)]

  //   if (swapExceptionsArr?.length > 0) {
  //     bridgeableTokens = bridgeableTokens.filter(
  //       (toToken) => toToken.symbol === token.symbol
  //     )
  //   }

  //   let bridgeableToken: Token =
  //     positedToToken &&
  //     !(swapExceptionsArr?.length > 0) &&
  //     token.swapableType === positedToToken.swapableType
  //       ? positedToToken
  //       : bridgeableTokens[0]
  //   return bridgeableTokens
  // }, [fromTokens, fromTokenIdx, toChainId])
  // Generates list of tokens for the current chain, while sorting them by balance.
  // useMemo is used to prevent the function from being called on every render (no need to re-request balances if fromChainId or address hasn't changed)
  // let fromTokens = useMemo(() => {
  //   console.log('MEMO fromTokens')
  //   sortByTokenBalance(
  //     BRIDGABLE_TOKENS[fromChainId],
  //     fromChainId,
  //     address
  //   ).then((tokens) => {
  //     console.log('tokensadhasjkhdakjshskjhs', tokens)
  //     return tokens
  //   })
  // }, [fromChainId, address])
  // console.log('FROM TOKENS', fromTokens)

  // let toTokens = useMemo(() => {
  //   console.log(
  //     'fromToken, toChainId, fromToken',
  //     fromToken,
  //     toChainId,
  //     fromToken
  //   )

  // console.log('TO TOKENS', toTokens)
  // Handle entry/amounts in the card

  // const [swapableType, setSwapableType] = useState(DEFAULT_SWAPABLE_TYPE)

  // Handle wagmi changes
  // useEffect(() => {
  //   if (address === undefined) {
  //     setFromChainId(DEFAULT_FROM_CHAIN)
  //   } else {
  //     setFromChainId(Number(fromChainIdRaw?.id))
  //   }
  // }, [fromChainIdRaw])
  // Get data from query url.

  // Handle url changes

  const getMostCommonSwapableType = (chainId: number) => {
    let fromChainTokensByType = Object.values(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId]
    )
    let maxTokenLength = 0

    let mostCommonSwapableType: Token[] = fromChainTokensByType[0]
    fromChainTokensByType.map((tokenArr, i) => {
      if (tokenArr.length > maxTokenLength) {
        console.log('UR A FGS')
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

  // //Keeps the url in sync with the state
  // useEffect(() => {
  //   // console.log("SLKdhsdhashdCUMMM")
  //   router.push({
  //     pathname: BRIDGE_PATH,

  //     query: {
  //       outputChain: toChainId,
  //       inputCurrency: fromToken.symbol,
  //       outputCurrency: toToken.symbol,
  //     },
  //   })
  // }, [toChainId, fromToken, toToken])

  const updateUrlParams = ({
    outputChain,
    inputCurrency,
    outputCurrency,
  }: {
    outputChain: any
    inputCurrency: any
    outputCurrency: any
  }) => {
    router.push({
      pathname: BRIDGE_PATH,
      query: {
        outputChain: outputChain,
        inputCurrency: inputCurrency,
        outputCurrency: outputCurrency,
      },
    })
  }

  const handleNewFromTokenNew = (
    token: Token,
    positedToChain: number | undefined,
    positedToSymbol: string | undefined
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

  const handleNewFromToken = (
    token: Token,
    positedToChain: number | undefined,
    positedToSymbol: string | undefined,
    updateUrl: boolean
  ) => {
    let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN

    console.log('NEW TO CHAIN', token)
    let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
      String(token.swapableType)
    ].filter((chainId) => Number(chainId) !== fromChainId)
    let swapExceptionsArr: number[] =
      token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
    if (swapExceptionsArr?.length > 0) {
      bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
    }

    // TODO filter above
    if (!bridgeableChains.includes(String(newToChain))) {
      newToChain =
        Number(bridgeableChains[0]) === fromChainId
          ? Number(bridgeableChains[1])
          : Number(bridgeableChains[0])
    }
    let positedToToken = positedToSymbol
      ? tokenSymbolToToken(newToChain, positedToSymbol)
      : tokenSymbolToToken(newToChain, token.symbol)

    console.log('positedToTokenpositedToTokenpositedToToken', positedToToken)
    let bridgeableTokens: Token[] =
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain][String(token.swapableType)]

    if (swapExceptionsArr?.length > 0) {
      bridgeableTokens = bridgeableTokens.filter(
        (toToken) => toToken.symbol === token.symbol
      )
    }
    let bridgeableToken: Token = positedToToken
    if (!bridgeableTokens.includes(positedToToken)) {
      bridgeableToken = bridgeableTokens[0]
    }

    // let bridgeableToken: Token =
    //   positedToToken &&
    //   !(swapExceptionsArr?.length > 0) &&
    //   token.swapableType === positedToToken.swapableType
    //     ? positedToToken
    //     : bridgeableTokens[0]

    setToToken(bridgeableToken)
    setToBridgeableTokens(bridgeableTokens)
    setToBridgeableChains(bridgeableChains)
    if (
      updateUrl ||
      newToChain !== positedToChain ||
      bridgeableToken.symbol !== positedToSymbol
    ) {
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: token.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    } else {
      setToChainId(Number(newToChain))
    }
  }

  // useEffect(() => {
  //   handleNewFromToken(fromToken, toChainId, toToken, true)
  // }, [fromToken])

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = async () => {
    // let oldFromChainId = fromChainId
    // let oldToChainId = toChainId
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      switchNetwork?.(toChainId)
    }
    resetRates()

    // setToChainId(fromChainId)
  }

  // Changes destination change when the user changes the toChainId
  const handleFromChainChange = (chainId: number) => {
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      switchNetwork?.(chainId)
      let newChainTokenList: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId][String(fromToken.swapableType)]
      )
      let newToken =
        newChainTokenList.length > 0
          ? fromToken
          : getMostCommonSwapableType(chainId)

      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromTokenNew(newToken, toChainId, toToken.symbol)
      setFromToken(newToken)
      setToBridgeableChains(bridgeableChains)
      setToBridgeableTokens(bridgeableTokens)
      setToToken(bridgeableToken)
      setToChainId(newToChain)
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
    /*
1. set the new token
2. if setting the origin, set the swapable types for the destination
3. set last switch?
4. calculate the bridge amount for destination/do a quote call to the sdk
5. update the url

*/
    // console.log('start start start start start start start start start start')
    // set the new token
    if (type == 'from') {
      console.log('from token change', token, token.swapableType, token.symbol)

      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromTokenNew(token, toChainId, token.symbol)
      setFromToken(token)
      setToBridgeableChains(bridgeableChains)
      setToBridgeableTokens(bridgeableTokens)
      setToToken(bridgeableToken)
      setToChainId(newToChain)

      // Update url params if any passed params were invalid

      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: token.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
      // setToToken(bridgeableToken)
      // setToChainId(newToChain)
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
  // REDO WITH SDK
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

  // //kasdhkajhajkshdksajhdjksahdkjashdoasdpoasid[asd[asdpasdasjdaskndas]]
  // const fromTokens = useMemo(() => {
  //   console.log('address, fromChainId', address, fromChainId)

  //   if (address === undefined || fromChainId === undefined) {
  //     return BRIDGABLE_TOKENS[fromChainId]
  //   }

  //   let tokens = BRIDGABLE_TOKENS[fromChainId]
  //   let chainId = fromChainId

  //   let i = 0
  //   let tokensWithBalances: any[] = []
  //   let zeroTokensWithBalances: any[] = []
  //   console.log('Sdskladalskjdklsajdlksajdskla  tok', tokens)
  //   // go through all tokens and retrieve token balances
  //   while (i < tokens?.length) {
  //     let tokenAddr = tokens[i].addresses[chainId as keyof Token['addresses']]

  //     let rawTokenBalance: any

  //     // Check for native token
  //     if (tokenAddr === '') {
  //       const { data } = useBalance({
  //         address: address,
  //         chainId: chainId,
  //       })
  //       rawTokenBalance = data
  //     } else if (tokenAddr?.length > 0) {
  //       const { data } = useBalance({
  //         address: address,
  //         token: `0x${tokenAddr.slice(2)}`,
  //         chainId: chainId,
  //       })
  //       rawTokenBalance = data
  //     }

  //     // manages two the array of tokens with zero balances and non-zero balances
  //     if (rawTokenBalance) {
  //       if (rawTokenBalance?.value._hex !== '0x00') {
  //         zeroTokensWithBalances.push({
  //           token: tokens[i],
  //           balance: rawTokenBalance.value,
  //         })
  //       } else {
  //         tokensWithBalances.push({
  //           token: tokens[i],
  //           balance: rawTokenBalance.value,
  //         })
  //       }
  //     }
  //     i++
  //   }
  //   let tokenList = zeroTokensWithBalances.concat(tokensWithBalances)
  //   console.log('tokenBalances', tokenList)
  //   return tokenList
  //   // return sortByTokenBalance(
  //   //   BRIDGABLE_TOKENS[fromChainId],
  //   //   fromChainId,
  //   //   address
  //   // )
  // }, [address, fromChainId])
  // console.log('fromTokens', fromTokens)

  // ///____Sd-sad-0as9d-0as9d-as0
  // const toTokens = useMemo(() => {
  //   console.log(
  //     'fromToken, toChainId, fromToken',
  //     fromToken,
  //     toChainId,
  //     fromToken
  //   )
  //   if (fromToken === undefined) {
  //     return BRIDGABLE_TOKENS[fromChainId]
  //   }
  //   let positedToChain = toChainId
  //   let token = fromToken
  //   let positedToSymbol = undefined
  //   let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN

  //   let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
  //     String(token.swapableType)
  //   ].filter((chainId) => Number(chainId) !== fromChainId)
  //   let swapExceptionsArr: number[] =
  //     token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
  //   if (swapExceptionsArr?.length > 0) {
  //     bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
  //   }

  //   // TODO filter above
  //   if (!bridgeableChains.includes(String(newToChain))) {
  //     newToChain =
  //       Number(bridgeableChains[0]) === fromChainId
  //         ? Number(bridgeableChains[1])
  //         : Number(bridgeableChains[0])
  //   }
  //   let positedToToken = positedToSymbol
  //     ? tokenSymbolToToken(newToChain, positedToSymbol)
  //     : undefined

  //   let bridgeableTokens: Token[] =
  //     BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain][String(token.swapableType)]

  //   if (swapExceptionsArr?.length > 0) {
  //     bridgeableTokens = bridgeableTokens.filter(
  //       (toToken) => toToken.symbol === token.symbol
  //     )
  //   }

  //   let bridgeableToken: Token =
  //     positedToToken &&
  //     !(swapExceptionsArr?.length > 0) &&
  //     token.swapableType === positedToToken.swapableType
  //       ? positedToToken
  //       : bridgeableTokens[0]
  //   return bridgeableTokens
  // }, [fromToken, toChainId, fromToken])
  // console.log('to', toTokens)
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
              {/* <button
                onClick={() => {
                  console.log(BRIDGABLE_TOKENS[toChainQuery], toChainQuery)
                  setFromCoin(BRIDGABLE_TOKENS[toChainQuery][2])
                }}
              >
                asdsadsa
              </button> */}
              {/* <HarmonyCheck fromChainId={fromChainId} toChainId={toChainId} /> */}
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
