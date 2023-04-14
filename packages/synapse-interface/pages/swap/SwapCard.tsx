import debounce from 'lodash.debounce'
import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useNetwork } from 'wagmi'
import { useEffect, useState, useMemo } from 'react'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { fetchSigner, getNetwork, switchNetwork } from '@wagmi/core'
import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import {
  BRIDGABLE_TOKENS,
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  tokenSymbolToToken,
} from '@constants/tokens'
import { formatBNToString } from '@utils/bignumber/format'
import { commify } from '@ethersproject/units'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'

import { Token } from '@/utils/types'
import { SWAP_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { useSynapseContext } from '@/utils/SynapseProvider'

import { parseUnits } from '@ethersproject/units'

import { Transition } from '@headlessui/react'

import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'

import Card from '@tw/Card'

import { SwapQuote } from '@types'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
  EMPTY_SWAP_QUOTE,
  QUOTE_POLLING_INTERVAL,
} from '@/constants/swap'
import { cleanNumberInput } from '@utils/cleanNumberInput'

import { POOL_PRIORITY_RANKING } from '@constants/tokens'
import NoSwapCard from './NoSwapCard'

const SwapCard = ({ address }: { address: string }) => {
  const router = useRouter()
  const SynapseSDK = useSynapseContext()
  const { chain: connectedChain } = useNetwork()
  const [time, setTime] = useState(Date.now())
  const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [fromTokens, setFromTokens] = useState([])
  const [fromInput, setFromInput] = useState({ string: '', bigNum: Zero })
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [toTokens, setToTokens] = useState([]) //add default

  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')

  const [swapQuote, setSwapQuote] = useState<SwapQuote>(EMPTY_SWAP_QUOTE)
  const [displayType, setDisplayType] = useState(undefined)

  /*
  useEffect Trigger: onMount
  - Gets current network connected and sets it as the state.
  - Initializes polling (setInterval) func to re-retrieve quotes.
  */
  useEffect(() => {
    const { chain: fromChainIdRaw } = getNetwork()
    setFromChainId(fromChainIdRaw ? fromChainIdRaw?.id : DEFAULT_FROM_CHAIN)
    const interval = setInterval(
      () => setTime(Date.now()),
      QUOTE_POLLING_INTERVAL
    )
    return () => {
      clearInterval(interval)
    }
  }, [])

  useEffect(() => {
    if (!router.isReady) {
      return
    }
    const {
      inputCurrency: fromTokenSymbolUrl,
      outputCurrency: toTokenSymbolUrl,
    } = router.query

    let tempFromToken: Token = getMostCommonSwapableType(fromChainId)

    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        tempFromToken = token
      }
    }
    const { swapableToken, swapableTokens } = handleNewFromToken(
      tempFromToken,
      toTokenSymbolUrl ? String(toTokenSymbolUrl) : undefined,
      fromChainId
    )
    resetTokenPermutation(
      tempFromToken,
      swapableToken,
      swapableTokens,
      tempFromToken.symbol,
      swapableToken.symbol
    )
    updateUrlParams({
      inputCurrency: fromToken.symbol,
      outputCurrency: swapableToken.symbol,
    })
  }, [router.isReady])

  /*
  useEffect Trigger: connectedChain
  - when the connected chain changes (wagmi hook), update the state
  */
  useEffect(() => {
    if (connectedChain?.id) {
      if (address === undefined) {
        return
      }
      setFromChainId(connectedChain?.id)
      handleChainChange(connectedChain?.id, false, 'from')
      sortByTokenBalance(
        BRIDGABLE_TOKENS[connectedChain?.id],
        connectedChain?.id,
        address
      ).then((tokens) => {
        setFromTokens(tokens)
      })
      return
    }
  }, [connectedChain?.id])

  /*
  useEffect Triggers: toToken, fromInput, toChainId, time
  - Gets a quote when the polling function is executed or any of the bridge attributes are altered.
  */
  useEffect(() => {
    if (
      fromChainId &&
      String(fromToken.addresses[fromChainId]) &&
      fromInput &&
      fromInput.bigNum.gt(Zero)
    ) {
      // TODO this needs to be debounced or throttled somehow to prevent spam and lag in the ui
      getQuote()
    } else {
      setSwapQuote(EMPTY_SWAP_QUOTE)
    }
  }, [toToken, fromInput, time])

  /*
  Helper Function: resetTokenPermutation
  - Handles when theres a new from token/chain and all other parts of the bridge arrangement needs to be updated
  - Updates url params.
  */
  const resetTokenPermutation = (
    newFromToken: Token,
    newToToken: Token,
    newBridgeableTokens: Token[],
    newFromTokenSymbol: string,
    newBridgeableTokenSymbol: string
  ) => {
    setFromToken(newFromToken)
    setToToken(newToToken)
    setToTokens(newBridgeableTokens)
    resetRates()
    updateUrlParams({
      inputCurrency: newFromTokenSymbol,
      outputCurrency: newBridgeableTokenSymbol,
    })
  }

  /*
  Helper Function: resetRates
  - Called when switching from chain/token so that the from input isn't populated with stale data.
  */
  const resetRates = () => {
    setSwapQuote(EMPTY_SWAP_QUOTE)
    setFromInput({ string: '', bigNum: Zero })
  }

  /*
  Helper Function: onChangeFromAmount
  - Ensures inputted data isn't too long and then sets state with the input.
  - Calculates BigNum from the input and stores in state as well (for quotes)
  */
  const onChangeFromAmount = (value: string) => {
    if (
      !(
        value.split('.')[1]?.length >
        fromToken[fromChainId as keyof Token['decimals']]
      )
    ) {
      let bigNum =
        stringToBigNum(value, fromToken.decimals[fromChainId]) ?? Zero
      setFromInput({
        string: value,
        bigNum: bigNum,
      })
    }
  }

  /*
  Helper Function: getMostCommonSwapableType
  - Returns the default token to display when switching chains. Usually returns stables or eth/wrapped eth.
  */
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

  /*
  Helper Function: updateUrlParams
  - Pushes chain and token changes to url
  NOTE: did not alter any variable names in case previous users have saved links of different bridging permutations.
  */
  const updateUrlParams = ({
    inputCurrency,
    outputCurrency,
  }: {
    inputCurrency: string
    outputCurrency: string
  }) => {
    router.replace(
      {
        pathname: SWAP_PATH,
        query: {
          inputCurrency,
          outputCurrency,
        },
      },
      undefined,
      { shallow: true }
    )
  }

  /*
   Helper Function: getCurrentTokenAllowance
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  TODO store this erc20 and signer retrieval in a state in a parent component? add to utils + use memo?
  */
  const getCurrentTokenAllowance = async (routerAddress: string) => {
    const wallet = await fetchSigner({
      chainId: fromChainId,
    })

    const erc20 = new Contract(
      fromToken.addresses[fromChainId],
      erc20ABI,
      wallet
    )
    const allowance = await erc20.allowance(address, routerAddress)
    return allowance
  }

  /*
  Function: handleNewFromToken
  - Handles all the changes that occur when selecting a new "from token", such as generating lists of potential chains/tokens
   to bridge to and handling if the current "to chain/token" are incompatible.
  */
  const handleNewFromToken = (
    token: Token,
    positedToSymbol: string | undefined,
    fromChainId: number
  ) => {
    const swapExceptionsArr: number[] =
      token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]

    const positedToToken = positedToSymbol
      ? tokenSymbolToToken(fromChainId, positedToSymbol)
      : tokenSymbolToToken(fromChainId, token.symbol)

    let swapableTokens: Token[] = sortByVisibilityRank(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][String(token.swapableType)]
    )

    if (swapExceptionsArr?.length > 0) {
      swapableTokens = swapableTokens.filter(
        (toToken) => toToken.symbol === token.symbol
      )
    }
    let swapableToken: Token = positedToToken
    if (!swapableTokens.includes(positedToToken)) {
      swapableToken = swapableTokens[0]
    }

    return {
      swapableToken,
      swapableTokens,
    }
  }

  /*
    Function:handleTokenChange
  - Handles when the user selects a new token from either the origin or destination
  */
  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    switch (type) {
      case 'from':
        const { swapableToken, swapableTokens } = handleNewFromToken(
          token,
          toToken.symbol,
          fromChainId
        )
        resetTokenPermutation(
          token,
          swapableToken,
          swapableTokens,
          token.symbol,
          swapableToken.symbol
        )
        return
      case 'to':
        resetRates()
        setToToken(token)
        updateUrlParams({
          inputCurrency: fromToken.symbol,
          outputCurrency: token.symbol,
        })
        return
    }
  }

  /*
   Function: getQuote
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  */
  const getQuote = async () => {
    const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
      await SynapseSDK.bridgeQuote(
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId],
        toToken.addresses[toChainId],
        fromInput.bigNum
      )
    if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
      setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO)
      return
    }
    const toValueBigNum = maxAmountOut ?? Zero
    const adjustedFeeAmount = feeAmount.lt(fromInput.bigNum)
      ? feeAmount
      : feeAmount.div(BigNumber.from(10).pow(18 - toToken.decimals[toChainId]))

    const allowance =
      fromToken.addresses[fromChainId] === AddressZero
        ? Zero
        : await getCurrentTokenAllowance(routerAddress)
    setBridgeQuote({
      outputAmount: toValueBigNum,
      outputAmountString: commify(
        formatBNToString(toValueBigNum, toToken.decimals[toChainId], 8)
      ),
      routerAddress,
      allowance,
      exchangeRate: calculateExchangeRate(
        fromInput.bigNum.sub(adjustedFeeAmount),
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
    return
  }

  /*
  Function: approveToken
  - Gets raw unsigned tx data from sdk and then execute it with ethers.
  - Only executes if token has already been approved.
   */
  const executeBridge = async () => {
    const wallet = await fetchSigner({
      chainId: fromChainId,
    })

    const data = await SynapseSDK.bridge(
      address,
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId as keyof Token['addresses']],
      fromInput.bigNum,
      bridgeQuote.quotes.originQuery,
      bridgeQuote.quotes.destQuery
    )
    const tx = await wallet.sendTransaction(data)
    try {
      await tx.wait()
      console.log(`Transaction mined successfully: ${tx.hash}`)
      return tx
    } catch (error) {
      console.log(`Transaction failed with error: ${error}`)
    }
  }

  // const fromArgs = {
  //   isSwapFrom: true,
  //   selected: fromCoin,
  //   onChangeSelected: onSelectFromCoin,
  //   onChangeAmount: onChangeFromAmount,
  //   inputValue: fromValue,
  //   inputRef: fromRef,
  //   tokens: swapableTokens,
  //   chainId,
  //   setDisplayType,
  //   onChangeChain,
  //   selectedChainId: chainId,
  // }

  // const toArgs = {
  //   isSwapFrom: false,
  //   selected: toCoin,
  //   onChangeSelected: onSelectToCoin,
  //   onChangeAmount: onChangeToAmount,
  //   inputValue: toValue,
  //   swapFromToCoins: swapFromToCoins,
  //   inputRef: toRef,
  //   tokens: swapableTokens,
  //   chainId,
  //   setDisplayType,
  //   onChangeChain,
  // }

  const approvalBtn = (
    <TransactionButton
      onClick={approveToken}
      label={`Approve ${displaySymbol(chainId, fromCoin)}`}
      pendingLabel={`Approving ${displaySymbol(chainId, fromCoin)}  `}
    />
  )

  let swapButtonLabel
  if (error) {
    swapButtonLabel = error
  } else {
    swapButtonLabel = fromAmount.eq(0)
      ? 'Enter amount to swap'
      : 'Swap your funds'
  }
  const swapBtn = (
    <TransactionButton
      disabled={toAmount.eq(0) || error}
      onClick={() => {
        return approveAndSwap({
          fromAmount: fromAmount,
          fromCoin,
          toAmount: toAmount,
          toCoin,
        })
      }}
      onSuccess={() => {
        onChangeFromAmount('')
      }}
      label={swapButtonLabel}
      pendingLabel={`Swapping...`}
    />
  )

  let actionBtn
  if (
    approvalState === APPROVAL_STATE.NOT_APPROVED &&
    fromCoin.symbol != WETH.symbol
  ) {
    actionBtn = approvalBtn
  } else {
    actionBtn = swapBtn
  }

  const transitionProps = {
    ...COIN_SLIDE_OVER_PROPS,
    className: `
      origin-bottom absolute
      w-full h-full
      md:w-[95%] md:h-[95%]
      -ml-0 md:-ml-3
      md:mt-3
      bg-bgBase
      z-20 rounded-3xl
    `,
  }

  return (
    <Card
      divider={false}
      className="max-w-lg px-1 pb-0 -mb-3 transition-all duration-100 transform rounded-xl bg-bgBase md:px-6 lg:px-6"
    >
      <div className="mb-8">
        <Transition show={displayType === 'from'} {...transitionProps}>
          <CoinSlideOver key="fromBlock" {...fromArgs} />
        </Transition>
        <Transition show={displayType === 'to'} {...transitionProps}>
          <CoinSlideOver key="toBlock" {...toArgs} />{' '}
        </Transition>
        <Transition show={displayType === 'fromChain'} {...transitionProps}>
          <NetworkSlideOver key="fromChainBlock" {...fromArgs} />{' '}
        </Transition>
        <Grid cols={{ xs: 1 }} gap={4} className="place-content-center">
          <div className="pt-3 pb-3 pl-4 pr-4 mt-2 border-none bg-bgLight rounded-xl">
            <ChainLabel
              isOrigin={true}
              chainId={chainId}
              setDisplayType={setDisplayType}
              onChangeChain={onChangeChain}
              titleText="Chain"
            />
          </div>
          <CoreSwapContainer {...fromArgs} />
          <CoreSwapContainer {...toArgs} />
        </Grid>

        <ExchangeRateInfo
          fromAmount={fromAmount}
          fromToken={fromCoin}
          toCoin={toCoin}
          exchangeRate={exchangeRate}
          priceImpact={priceImpact}
          toChainId={chainId}
        />
        <div className="px-2 py-2 md:px-0 md:py-4">{actionBtn}</div>
      </div>
    </Card>
  )
}

export default SwapCard
