import { useAccount, useNetwork } from 'wagmi'
import { useSelector } from 'react-redux'
import { RootState } from '../../store/store'
import toast from 'react-hot-toast'
import { useSpring, animated } from 'react-spring'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import BridgeWatcher from '@/pages/bridge/BridgeWatcher'
import { useRouter } from 'next/router'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import {
  setFromToken,
  setToToken,
  updateFromValue,
  setBridgeQuote,
  setIsLoading,
  setFromChainId,
  setToChainId,
  setSupportedFromTokens,
  setSupportedToTokens,
  setFromChainIds,
  setToChainIds,
  setDeadlineMinutes,
  setDestinationAddress,
  addBridgeTxHash,
} from '@/slices/bridge/reducer'

import {
  setShowDestinationAddress,
  setShowFromChainSlideOver,
  setShowSettingsSlideOver,
  setShowToChainSlideOver,
} from '@/slices/bridgeDisplaySlice'

import {
  DEFAULT_TO_CHAIN,
  EMPTY_BRIDGE_QUOTE,
  EMPTY_BRIDGE_QUOTE_ZERO,
} from '@/constants/bridge'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { subtractSlippage } from '@/utils/slippage'
import { commify } from '@ethersproject/units'
import { formatBigIntToString, powBigInt } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { useEffect, useRef, useState, useMemo } from 'react'
import { Token } from '@/utils/types'
import { getWalletClient } from '@wagmi/core'
import { txErrorHandler } from '@/utils/txErrorHandler'
import {
  BRIDGABLE_TOKENS,
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  tokenSymbolToToken,
} from '@/constants/tokens'
import { AcceptedChainId, CHAINS_ARR, CHAINS_BY_ID } from '@/constants/chains'
import { approveToken } from '@/utils/approveToken'
import { PageHeader } from '@/components/PageHeader'
import Card from '@/components/ui/tailwind/Card'
import BridgeExchangeRateInfo from '@/components/StateManagedBridge/BridgeExchangeRateInfo'
import { Transition } from '@headlessui/react'
import {
  SECTION_TRANSITION_PROPS,
  TRANSITION_PROPS,
} from '@/styles/transitions'
import { TokenSlideOver } from '@/components/StateManagedBridge/TokenSlideOver'
import { InputContainer } from '@/components/StateManagedBridge/InputContainer'
import { OutputContainer } from '@/components/StateManagedBridge/OutputContainer'
import {
  sortByVisibilityRank,
  separateAndSortTokensWithBalances,
  sortTokensByPriorityRankAndAlpha,
} from '@/utils/sortTokens'
import { ChainSlideOver } from '@/components/StateManagedBridge/ChainSlideOver'
import SettingsSlideOver from '@/components/StateManagedBridge/SettingsSlideOver'
import Button from '@/components/ui/tailwind/Button'
import { SettingsIcon } from '@/components/icons/SettingsIcon'
import { DestinationAddressInput } from '@/components/StateManagedBridge/DestinationAddressInput'
import { isAddress } from '@ethersproject/address'
import { BridgeTransactionButton } from '@/components/StateManagedBridge/BridgeTransactionButton'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { Address, zeroAddress } from 'viem'
import { stringToBigInt } from '@/utils/bigint/format'
import { Warning } from '@/components/Warning'
import { useAppDispatch } from '@/store/hooks'
import {
  fetchAndStoreSingleTokenAllowance,
  fetchAndStoreSingleTokenBalance,
  useFetchPortfolioBalances,
} from '@/slices/portfolio/hooks'
import {
  FetchState,
  updateSingleTokenAllowance,
} from '@/slices/portfolio/actions'
import { addRecentBridgeTransaction } from '@/slices/bridge/actions'
import { getTimeMinutesFromNow } from '@/utils/time'

// NOTE: These are idle utility functions that will be re-written to
// support sorting by desired mechanism
// We want to keep them separate as to not overload Component and UI logic
// i.e., call when needed

const sortFromChainIds = (chainIds: number[]) => {
  return chainIds
}

const sortToChainIds = (chainIds: number[]) => {
  return chainIds
}

const sortFromTokens = (tokens: Token[]) => {
  return sortTokensByPriorityRankAndAlpha(tokens)
}

const sortToTokens = (tokens: Token[]) => {
  return sortTokensByPriorityRankAndAlpha(tokens)
}

// Need to update url params

const StateManagedBridge = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const { synapseSDK } = useSynapseContext()
  const bridgeDisplayRef = useRef(null)
  const currentSDKRequestID = useRef(0)
  const router = useRouter()
  const { query, pathname } = router

  const { balancesAndAllowances: portfolioBalances, status: portfolioStatus } =
    useFetchPortfolioBalances()

  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    bridgeQuote,
    fromValue,
    isLoading,
    supportedFromTokens,
    supportedToTokens,
    destinationAddress,
    bridgeTxHashes,
  } = useSelector((state: RootState) => state.bridge)

  const {
    showFromTokenSlideOver,
    showToTokenSlideOver,
    showFromChainSlideOver,
    showToChainSlideOver,
    showSettingsSlideOver,
    showDestinationAddress,
  } = useSelector((state: RootState) => state.bridgeDisplay)

  let pendingPopup
  let successPopup

  const [isApproved, setIsApproved] = useState(false)

  const dispatch = useAppDispatch()

  const fromTokens = BRIDGABLE_TOKENS[fromChainId]
  const fromChainIds = Object.keys(CHAINS_BY_ID).map((id) => Number(id))
  const toChainIds = Object.keys(CHAINS_BY_ID).map((id) => Number(id))

  useEffect(() => {
    segmentAnalyticsEvent(`[Bridge page] arrives`, {
      fromChainId,
      query,
      pathname,
    })
  }, [query])

  // Commenting out for a bit to debug, but basic issue is we need
  // a mapping for allowable routes/tokens, and how we set them on
  // init and state changes

  // const toChainIds = BRIDGE_CHAINS_BY_TYPE[fromToken.swapableType]
  //   .filter((chainId) => Number(chainId) !== fromChainId)
  //   .map((chainId) => Number(chainId))

  // Can be smarter about breaking out which calls happen assoc with which
  // dependencies (like some stuff should only change on fromChainId changes)
  useEffect(() => {
    let fromTokens = BRIDGABLE_TOKENS[fromChainId] ?? []
    const toTokens = BRIDGABLE_TOKENS[toChainId]

    // Checking whether the selected fromToken exists in the BRIDGABLE_TOKENS for the chosen chain
    if (!fromTokens.some((token) => token.symbol === fromToken?.symbol)) {
      // Sort the tokens based on priorityRank in ascending order
      const sortedTokens = fromTokens.sort(
        (a, b) => a.priorityRank - b.priorityRank
      )
      // Select the token with the highest priority rank
      dispatch(setFromToken(sortedTokens[0]))
      // Update fromTokens for the selected fromToken
      fromTokens = [fromToken]
    }

    let { bridgeableChainIds, bridgeableTokens, bridgeableToken } =
      findSupportedChainsAndTokens(
        fromToken,
        toChainId,
        toToken?.symbol,
        fromChainId
      )

    let bridgeableToChainId = toChainId
    if (!bridgeableChainIds?.includes(toChainId)) {
      const sortedChainIds = bridgeableChainIds?.sort((a, b) => {
        const chainA = CHAINS_ARR.find((chain) => chain.id === a)
        const chainB = CHAINS_ARR.find((chain) => chain.id === b)
        return chainB.priorityRank - chainA.priorityRank
      })
      bridgeableToChainId = sortedChainIds?.[0]
    }

    dispatch(setSupportedToTokens(sortToTokens(bridgeableTokens)))
    dispatch(setToToken(bridgeableToken))
    dispatch(setSupportedFromTokens(portfolioBalances[fromChainId] ?? []))
    dispatch(setFromChainIds(fromChainIds))
    dispatch(setToChainIds(bridgeableChainIds))

    if (bridgeableToChainId && bridgeableToChainId !== toChainId) {
      dispatch(setToChainId(bridgeableToChainId))
    }

    console.log(`[useEffect] fromToken`, fromToken?.symbol)
    console.log(`[useEffect] toToken`, toToken?.symbol)
    // TODO: Double serialization happening somewhere??
    if (
      fromToken?.decimals[fromChainId] &&
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) > 0n
    ) {
      console.log('trying to set bridge quote')
      getAndSetBridgeQuote()
    } else {
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setIsLoading(false))
    }
  }, [
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    fromValue,
    address,
    portfolioBalances,
  ])

  // don't like this, rewrite: could be custom hook
  useEffect(() => {
    if (fromToken?.addresses[fromChainId] === zeroAddress) {
      setIsApproved(true)
    } else {
      if (
        bridgeQuote?.allowance &&
        stringToBigInt(fromValue, fromToken.decimals[fromChainId]) <=
          bridgeQuote.allowance
      ) {
        setIsApproved(true)
      } else {
        setIsApproved(false)
      }
    }
  }, [bridgeQuote, fromToken, fromValue, fromChainId, toChainId])

  let quoteToast
  // Would like to move this into function outside of this component
  const getAndSetBridgeQuote = async () => {
    currentSDKRequestID.current += 1
    const thisRequestId = currentSDKRequestID.current
    // will have to handle deadlineMinutes here at later time, gets passed as optional last arg in .bridgeQuote()
    try {
      dispatch(setIsLoading(true))

      const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
        await synapseSDK.bridgeQuote(
          fromChainId,
          toChainId,
          fromToken.addresses[fromChainId],
          toToken.addresses[toChainId],
          stringToBigInt(fromValue, fromToken.decimals[fromChainId])
        )

      console.log(`[getAndSetQuote] fromChainId`, fromChainId)
      console.log(`[getAndSetQuote] toChainId`, toChainId)
      console.log(`[getAndSetQuote] fromToken.symbol`, fromToken.symbol)
      console.log(`[getAndSetQuote] toToken.symbol`, toToken.symbol)
      console.log(`[getAndSetQuote] fromValue`, fromValue)
      console.log('feeAmount', feeAmount)
      console.log(`[getAndSetQuote] maxAmountOut`, maxAmountOut)

      if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setIsLoading(false))
        return
      }

      const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

      const originTokenDecimals = fromToken.decimals[fromChainId]
      const adjustedFeeAmount =
        BigInt(feeAmount) <
        stringToBigInt(`${fromValue}`, fromToken.decimals[fromChainId])
          ? BigInt(feeAmount)
          : BigInt(feeAmount) / powBigInt(10n, BigInt(18 - originTokenDecimals))

      const isUnsupported = AcceptedChainId[fromChainId] ? false : true

      const allowance =
        fromToken.addresses[fromChainId] === zeroAddress ||
        address === undefined ||
        isUnsupported
          ? 0n
          : await getErc20TokenAllowance({
              address,
              chainId: fromChainId,
              tokenAddress: fromToken.addresses[fromChainId] as Address,
              spender: routerAddress,
            })

      if (fromToken.addresses[fromChainId] !== zeroAddress && address) {
        dispatch(
          updateSingleTokenAllowance({
            chainId: fromChainId,
            allowance: allowance,
            spender: routerAddress,
            token: fromToken,
          })
        )
      }

      const originMinWithSlippage = subtractSlippage(
        originQuery?.minAmountOut ?? 0n,
        'ONE_TENTH',
        null
      )
      const destMinWithSlippage = subtractSlippage(
        destQuery?.minAmountOut ?? 0n,
        'ONE_TENTH',
        null
      )

      let newOriginQuery = { ...originQuery }
      newOriginQuery.minAmountOut = originMinWithSlippage

      let newDestQuery = { ...destQuery }
      newDestQuery.minAmountOut = destMinWithSlippage
      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(
          setBridgeQuote({
            outputAmount: toValueBigInt,
            outputAmountString: commify(
              formatBigIntToString(
                toValueBigInt,
                toToken.decimals[toChainId],
                8
              )
            ),
            routerAddress,
            allowance,
            exchangeRate: calculateExchangeRate(
              stringToBigInt(fromValue, fromToken.decimals[fromChainId]) -
                BigInt(adjustedFeeAmount),
              fromToken.decimals[fromChainId],
              toValueBigInt,
              toToken.decimals[toChainId]
            ),
            feeAmount,
            delta: BigInt(maxAmountOut.toString()),
            quotes: {
              originQuery: newOriginQuery,
              destQuery: newDestQuery,
            },
          })
        )

        toast.dismiss(quoteToast)
        const message = `Route found for bridging ${fromValue} ${fromToken.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        console.log(message)
        quoteToast = toast(message, { duration: 3000 })
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        toast.dismiss(quoteToast)
        const message = `No route found for bridging ${fromValue} ${fromToken.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        console.log(message)
        quoteToast = toast(message, { duration: 3000 })

        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        return
      }
    } finally {
      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(setIsLoading(false))
      }
    }
  }

  const approveTxn = async () => {
    try {
      const tx = approveToken(
        bridgeQuote?.routerAddress,
        fromChainId,
        fromToken?.addresses[fromChainId]
      ).then(() => {
        dispatch(
          fetchAndStoreSingleTokenAllowance({
            routerAddress: bridgeQuote?.routerAddress as Address,
            tokenAddress: fromToken?.addresses[fromChainId] as Address,
            address: address,
            chainId: fromChainId,
          })
        )
      })

      try {
        await tx
        setIsApproved(true)
      } catch (error) {
        return txErrorHandler(error)
      }
    } catch (error) {
      return txErrorHandler(error)
    }
  }

  const executeBridge = async () => {
    segmentAnalyticsEvent(`[Bridge] initiates bridge`, {
      address,
      originChainId: fromChainId,
      destinationChainId: toChainId,
      inputAmount: fromValue,
      expectedReceivedAmount: bridgeQuote.outputAmountString,
      slippage: bridgeQuote.exchangeRate,
    })
    try {
      const wallet = await getWalletClient({
        chainId: fromChainId,
      })

      const toAddress =
        destinationAddress && isAddress(destinationAddress)
          ? destinationAddress
          : address

      const data = await synapseSDK.bridge(
        toAddress,
        bridgeQuote.routerAddress,
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId as keyof Token['addresses']],
        stringToBigInt(fromValue, fromToken.decimals[fromChainId]),
        bridgeQuote.quotes.originQuery,
        bridgeQuote.quotes.destQuery
      )

      const payload =
        fromToken.addresses[fromChainId as keyof Token['addresses']] ===
          zeroAddress ||
        fromToken.addresses[fromChainId as keyof Token['addresses']] === ''
          ? {
              data: data.data,
              to: data.to,
              value: stringToBigInt(fromValue, fromToken.decimals[fromChainId]),
            }
          : data

      const tx = await wallet.sendTransaction(payload)

      const originChainName = CHAINS_BY_ID[fromChainId]?.name
      const destinationChainName = CHAINS_BY_ID[toChainId]?.name
      pendingPopup = toast(
        `Bridging from ${fromToken.symbol} on ${originChainName} to ${toToken.symbol} on ${destinationChainName}`,
        { id: 'bridge-in-progress-popup', duration: Infinity }
      )

      try {
        segmentAnalyticsEvent(`[Bridge] bridges successfully`, {
          address,
          originChainId: fromChainId,
          destinationChainId: toChainId,
          inputAmount: fromValue,
          expectedReceivedAmount: bridgeQuote.outputAmountString,
          slippage: bridgeQuote.exchangeRate,
        })
        dispatch(
          addRecentBridgeTransaction({
            fromChainId: fromChainId,
            fromToken: fromToken,
            fromValue: fromValue,
            toChainId: toChainId,
            toToken: toToken,
            transactionHash: tx,
            timestamp: getTimeMinutesFromNow(0),
          })
        )
        dispatch(addBridgeTxHash(tx))
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setDestinationAddress(null))
        dispatch(setShowDestinationAddress(false))
        dispatch(updateFromValue(''))

        const successToastContent = (
          <div>
            <div>
              Successfully initiated bridge from {fromToken.symbol} on{' '}
              {originChainName} to {toToken.symbol} on {destinationChainName}
            </div>
            <ExplorerToastLink
              transactionHash={tx ?? zeroAddress}
              chainId={fromChainId}
            />
          </div>
        )

        successPopup = toast.success(successToastContent, {
          id: 'bridge-success-popup',
          duration: 10000,
        })

        toast.dismiss(pendingPopup)

        setTimeout(async () => {
          await dispatch(
            fetchAndStoreSingleTokenBalance({
              token: fromToken,
              routerAddress: bridgeQuote?.routerAddress as Address,
              address: address,
              chainId: fromChainId,
            })
          )
        }, 3000)

        console.log('tx:', tx)

        return tx
      } catch (error) {
        segmentAnalyticsEvent(`[Bridge] error bridging`, {
          address,
          errorCode: error.code,
        })
        console.log(`Transaction failed with error: ${error}`)
        toast.dismiss(pendingPopup)
      }
    } catch (error) {
      segmentAnalyticsEvent(`[Bridge]  error bridging`, {
        address,
        errorCode: error.code,
      })
      console.log('Error executing bridge', error)
      toast.dismiss(pendingPopup)
      return txErrorHandler(error)
    }
  }

  const springClass = 'fixed z-50 w-full h-full bg-opacity-50'

  return (
    <div className="flex flex-col w-full max-w-lg mx-auto lg:mx-0">
      <div className="flex flex-col">
        <div className="flex items-center justify-between">
          <PageHeader
            title="Bridge"
            subtitle="Send your assets across chains."
          />
          <div>
            <Button
              className="flex items-center p-3 text-opacity-75 bg-bgLight hover:bg-bgLighter text-secondaryTextColor hover:text-white"
              onClick={() => {
                if (showSettingsSlideOver === true) {
                  dispatch(setShowSettingsSlideOver(false))
                } else {
                  dispatch(setShowSettingsSlideOver(true))
                }
              }}
            >
              {!showSettingsSlideOver ? (
                <>
                  <SettingsIcon className="w-5 h-5 mr-2" />
                  <span>Settings</span>
                </>
              ) : (
                <span>Close</span>
              )}
            </Button>
          </div>
        </div>
        <Card
          divider={false}
          className={`
                px-2 pb-2 pt-2 md:px-6 md:pt-5 mt-5 overflow-hidden
                transition-all duration-100 transform rounded-xl
                bg-bgBase
              `}
        >
          <div ref={bridgeDisplayRef}>
            <Transition show={showFromTokenSlideOver} {...TRANSITION_PROPS}>
              <animated.div className={springClass}>
                <TokenSlideOver
                  key="fromBlock"
                  isOrigin={true}
                  tokens={
                    portfolioStatus === FetchState.VALID
                      ? separateAndSortTokensWithBalances(supportedFromTokens)
                      : sortByVisibilityRank(fromTokens)
                  }
                  chainId={fromChainId}
                  selectedToken={fromToken}
                />{' '}
              </animated.div>
            </Transition>
            <Transition show={showToTokenSlideOver} {...TRANSITION_PROPS}>
              <animated.div className={springClass}>
                <TokenSlideOver
                  key="toBlock"
                  isOrigin={false}
                  tokens={supportedToTokens}
                  chainId={toChainId}
                  selectedToken={toToken}
                />{' '}
              </animated.div>
            </Transition>
            <Transition show={showFromChainSlideOver} {...TRANSITION_PROPS}>
              <animated.div className={springClass}>
                <ChainSlideOver
                  key="fromChainBlock"
                  isOrigin={true}
                  chains={fromChainIds}
                  chainId={fromChainId}
                  setChain={setFromChainId}
                  setShowSlideOver={setShowFromChainSlideOver}
                />
              </animated.div>
            </Transition>
            <Transition show={showToChainSlideOver} {...TRANSITION_PROPS}>
              <animated.div className={springClass}>
                <ChainSlideOver
                  key="toChainBlock"
                  isOrigin={false}
                  chains={toChainIds}
                  chainId={toChainId}
                  setChain={setToChainId}
                  setShowSlideOver={setShowToChainSlideOver}
                />
              </animated.div>
            </Transition>
            <Transition show={showSettingsSlideOver} {...TRANSITION_PROPS}>
              <animated.div>
                <SettingsSlideOver key="settings" />
              </animated.div>
            </Transition>
            <InputContainer />
            <OutputContainer />
            <Warning
              originChainId={fromChainId}
              destinationChainId={toChainId}
              originToken={fromToken}
              destinationToken={toToken}
            />
            <Transition
              appear={true}
              unmount={false}
              show={true}
              {...SECTION_TRANSITION_PROPS}
            >
              <BridgeExchangeRateInfo showGasDrop={true} />
            </Transition>
            {showDestinationAddress && (
              <DestinationAddressInput
                toChainId={toChainId}
                destinationAddress={destinationAddress}
              />
            )}
            <div className="md:my-3">
              <BridgeTransactionButton
                isApproved={isApproved}
                approveTxn={approveTxn}
                executeBridge={executeBridge}
              />
            </div>
          </div>
        </Card>
        {/* <ActionCardFooter link={HOW_TO_BRIDGE_URL} /> */}
      </div>
      {/* <div className="mt-8">
        <BridgeWatcher
          fromChainId={fromChainId}
          toChainId={toChainId}
          address={address}
          destinationAddress={destinationAddress}
        />
      </div> */}
    </div>
  )
}

// TODO: Refactor
// would like to refactor this as a function that
// takes fromChainId, fromToken only and returns rest
// Determines the chain to be used for the token swap.
const getNewToChain = (positedToChain, fromChainId, bridgeableChains) => {
  // If positedToChain is defined and different from fromChainId, use it.
  // Otherwise, use a default chain.
  let newToChain =
    positedToChain && positedToChain !== fromChainId
      ? Number(positedToChain)
      : DEFAULT_TO_CHAIN
  // If newToChain is not a part of bridgeableChains, select a chain from bridgeableChains
  // that is different from fromChainId.
  if (!bridgeableChains?.includes(String(newToChain))) {
    newToChain =
      Number(bridgeableChains?.[0]) === fromChainId
        ? Number(bridgeableChains?.[1])
        : Number(bridgeableChains?.[0])
  }
  return newToChain
}

// Determines which chains are bridgeable based on the swapableType of the token.
const getBridgeableChains = (token, fromChainId, swapExceptionsArr) => {
  // Filter out chains that are not bridgeable for the given token type.
  let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
    String(token?.swapableType)
  ]?.filter((chainId) => Number(chainId) !== fromChainId)
  // If there are swap exceptions, replace bridgeableChains with the chains from exceptions.
  if (swapExceptionsArr?.length > 0) {
    bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
  }
  return bridgeableChains
}

// Determines which tokens are bridgeable on the new chain.
const getBridgeableTokens = (newToChain, token, swapExceptionsArr) => {
  // Get tokens that are bridgeable on the new chain and of the same type as the given token.
  let bridgeableTokens: Token[] = sortToTokens(
    BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain]?.[String(token?.swapableType)]
  )
  // If there are swap exceptions, filter out tokens that have a different symbol from the given token.
  if (swapExceptionsArr?.length > 0) {
    bridgeableTokens = bridgeableTokens.filter(
      (toToken) => toToken.symbol === token.symbol
    )
  }
  return bridgeableTokens
}

// Determines the token to be used for the swap.
const getBridgeableToken = (bridgeableTokens, positedToToken) => {
  // If positedToToken is a part of bridgeableTokens, use it.
  // Otherwise, use the first token from bridgeableTokens.
  let bridgeableToken: Token = positedToToken
  if (!bridgeableTokens?.includes(positedToToken)) {
    bridgeableToken = bridgeableTokens?.[0]
  }
  return bridgeableToken
}

// The main function to find bridgeable chains and tokens.
const findSupportedChainsAndTokens = (
  token: Token,
  positedToChain: number | undefined,
  positedToSymbol: string | undefined,
  fromChainId: number
) => {
  // Get the swap exceptions for the given fromChainId if any.
  const swapExceptionsArr: number[] =
    token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
  // Determine which chains are bridgeable.
  const bridgeableChains = getBridgeableChains(
    token,
    fromChainId,
    swapExceptionsArr
  )
  // Determine the new chain to be used for the swap.
  const newToChain = getNewToChain(
    positedToChain,
    fromChainId,
    bridgeableChains
  )
  // Determine the token to be used for the swap based on the posited symbol or the symbol of the given token.
  const positedToToken = positedToSymbol
    ? tokenSymbolToToken(newToChain, positedToSymbol)
    : tokenSymbolToToken(newToChain, token?.symbol)
  // Determine which tokens are bridgeable on the new chain.
  const bridgeableTokens = getBridgeableTokens(
    newToChain,
    token,
    swapExceptionsArr
  )
  // Determine the specific token to be used for the swap.
  const bridgeableToken = getBridgeableToken(bridgeableTokens, positedToToken)

  // Return the bridgeable chains, bridgeable tokens, and the specific bridgeable token.
  return {
    bridgeableChainIds: bridgeableChains?.map((chainId: string) =>
      Number(chainId)
    ),
    bridgeableTokens,
    bridgeableToken,
  }
}

export default StateManagedBridge
