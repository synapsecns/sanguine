import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useAccount } from 'wagmi'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../../store/store'
import toast from 'react-hot-toast'
import { useSpring, animated } from 'react-spring'

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
  setSupportedFromTokenBalances,
  setDeadlineMinutes,
  setDestinationAddress,
} from '@/slices/bridgeSlice'

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
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { getCurrentTokenAllowance } from '../../actions/getCurrentTokenAllowance'
import { subtractSlippage } from '@/utils/slippage'
import { commify } from '@ethersproject/units'
import { formatBNToString } from '@/utils/bignumber/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { useEffect, useRef, useState } from 'react'
import { Token } from '@/utils/types'
import { fetchSigner } from '@wagmi/core'
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
import ExchangeRateInfo from '@/components/ExchangeRateInfo'
import { Transition } from '@headlessui/react'
import {
  SECTION_TRANSITION_PROPS,
  TRANSITION_PROPS,
} from '@/styles/transitions'
import { TokenSlideOver } from '@/components/StateManagedBridge/TokenSlideOver'
import { InputContainer } from '@/components/StateManagedBridge/InputContainer'
import { OutputContainer } from '@/components/StateManagedBridge/OutputContainer'
import { sortByTokenBalance, sortByVisibilityRank } from '@/utils/sortTokens'
import { ChainSlideOver } from '@/components/StateManagedBridge/ChainSlideOver'
import SettingsSlideOver from '@/components/StateManagedBridge/SettingsSlideOver'
import Button from '@/components/ui/tailwind/Button'
import { SettingsIcon } from '@/components/icons/SettingsIcon'
import { DestinationAddressInput } from '@/components/StateManagedBridge/DestinationAddressInput'
import { isAddress } from '@ethersproject/address'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { BridgeTransactionButton } from '@/components/StateManagedBridge/BridgeTransactionButton'

// NOTE: These are idle utility functions that will be re-written to
// support sorting by desired mechanism
// We want to keep them separate as to not overload Component and UI logic
// i.e., call when needed

// Function to sort the tokens by priorityRank and alphabetically
function sortTokensArray(arr: Token[]): Token[] {
  // Create a copy of the array to prevent modifying the original one
  const sortedArr = [...arr]

  return sortedArr.sort((a, b) => {
    // Sort by priorityRank first
    if (a.priorityRank !== b.priorityRank) {
      return a.priorityRank - b.priorityRank
    }

    // If priorityRank is the same, sort by symbol
    return a.symbol.localeCompare(b.symbol)
  })
}


const sortFromChainIds = (chainIds: number[]) => {
  return chainIds
}

const sortToChainIds = (chainIds: number[]) => {
  return chainIds
}

const sortFromTokens = (tokens: Token[]) => {
  return sortTokensArray(tokens)
}

const sortToTokens = (tokens: Token[]) => {
  return sortTokensArray(tokens)
}

// Need to update url params

const StateManagedBridge = () => {
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const bridgeDisplayRef = useRef(null)

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
  } = useSelector((state: RootState) => state.bridge)

  const {
    showFromTokenSlideOver,
    showToTokenSlideOver,
    showFromChainSlideOver,
    showToChainSlideOver,
    showSettingsSlideOver,
    showDestinationAddress,
  } = useSelector((state: RootState) => state.bridgeDisplay)

  const [isApproved, setIsApproved] = useState(false)

  const dispatch = useDispatch()

  const fromChainIds = Object.keys(CHAINS_BY_ID).map((id) => Number(id))
  const toChainIds = Object.keys(CHAINS_BY_ID).map((id) => Number(id))

  // Commenting out for a bit to debug, but basic issue is we need
  // a mapping for allowable routes/tokens, and how we set them on
  // init and state changes

  // const toChainIds = BRIDGE_CHAINS_BY_TYPE[fromToken.swapableType]
  //   .filter((chainId) => Number(chainId) !== fromChainId)
  //   .map((chainId) => Number(chainId))

  // Can be smarter about breaking out which calls happen assoc with which
  // dependencies (like some stuff should only change on fromChainId changes)
  useEffect(() => {
    const fromTokens = BRIDGABLE_TOKENS[fromChainId]
    const toTokens = BRIDGABLE_TOKENS[toChainId]

    const { bridgeableChainIds, bridgeableTokens, bridgeableToken } =
      findSupportedChainsAndTokens(
        fromToken,
        toChainId,
        toToken.symbol,
        fromChainId
      )

      let bridgeableToChainId
       // Check if toChainId is in the bridgeableChainIds
       if (!bridgeableChainIds.includes(toChainId)) {
        // Assuming you have an array or object with all chains,
        // sort bridgeableChainIds based on the priorityRank of the corresponding chains
        // TODO: This can be refactored using the sortChains functions defined in constants/chains/index.tsx
        const sortedChainIds = bridgeableChainIds.sort((a, b) => {
          const chainA = CHAINS_ARR.find(chain => chain.id === a);           // Get chain object corresponding to ID a
          const chainB = CHAINS_ARR.find(chain => chain.id === b); // Get chain object corresponding to ID b

          return chainB.priorityRank - chainA.priorityRank; // Sort in descending order
        });

        // Set toChainId to the chain with the highest priorityRank
        bridgeableToChainId = sortedChainIds[0];
      }


    // when any of those changes happen,
    dispatch(setSupportedToTokens(sortToTokens(bridgeableTokens)))
    dispatch(setToToken(bridgeableToken))

    sortByTokenBalance(fromTokens, fromChainId, address).then((res) => {
      const t = res.map((tokenAndBalances) => tokenAndBalances.token)

      dispatch(setSupportedFromTokenBalances(res))
      dispatch(setSupportedFromTokens(sortFromTokens(t)))
    })

    dispatch(setFromChainIds(fromChainIds))
    dispatch(setToChainIds(bridgeableChainIds))
    if (bridgeableToChainId && bridgeableToChainId !== toChainId) {
      dispatch(setToChainId(bridgeableToChainId))  // Dispatch the updated toChainId
  }
    /// maybe you need to wrap this in a then/finally so it only happens
    // after the dispatches happen
    console.log(`[useEffect] fromToken`, fromToken.symbol)
    console.log(`[useEffect] toToken`, toToken.symbol)
    if (fromValue.gt(0)) {
      getAndSetBridgeQuote()
    } else {
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setIsLoading(false))
      console.log(bridgeQuote)
    }
  }, [fromChainId, toChainId, fromToken, toToken, fromValue])

  // don't like this, rewrite: could be custom hook
  useEffect(() => {
    if (fromToken?.addresses[fromChainId] === AddressZero) {
      setIsApproved(true)
    } else {
      if (bridgeQuote?.allowance && fromValue.lte(bridgeQuote.allowance)) {
        setIsApproved(true)
      } else {
        setIsApproved(false)
      }
    }
  }, [bridgeQuote, fromToken, fromValue, fromChainId, toChainId])

  // Would like to move this into function outside of this component
  const getAndSetBridgeQuote = async () => {
    // will have to handle deadlineMinutes here at later time, gets passed as optional last arg in .bridgeQuote()
    try {
      dispatch(setIsLoading(true))

      const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
        await synapseSDK.bridgeQuote(
          fromChainId,
          toChainId,
          fromToken.addresses[fromChainId],
          toToken.addresses[toChainId],
          fromValue
        )

      console.log(`[getAndSetQuote] fromChainId`, fromChainId)
      console.log(`[getAndSetQuote] toChainId`, toChainId)
      console.log(`[getAndSetQuote] fromToken.symbol`, fromToken.symbol)
      console.log(`[getAndSetQuote] toToken.symbol`, toToken.symbol)
      console.log(`[getAndSetQuote] fromValue`, fromValue)

      console.log(`[getAndSetQuote] maxAmountOut`, maxAmountOut)

      if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setIsLoading(false))
        return
      }

      const toValueBigNum = maxAmountOut ?? Zero
      const originTokenDecimals = fromToken.decimals[fromChainId]
      const adjustedFeeAmount = feeAmount.lt(fromValue)
        ? feeAmount
        : feeAmount.div(BigNumber.from(10).pow(18 - originTokenDecimals))

      const isUnsupported = AcceptedChainId[fromChainId] ? false : true

      const allowance =
        fromToken.addresses[fromChainId] === AddressZero ||
        address === undefined ||
        isUnsupported
          ? Zero
          : await getCurrentTokenAllowance(
              address,
              fromChainId,
              fromToken,
              routerAddress
            )

      const originMinWithSlippage = subtractSlippage(
        originQuery?.minAmountOut ?? Zero,
        'ONE_TENTH',
        null
      )
      const destMinWithSlippage = subtractSlippage(
        destQuery?.minAmountOut ?? Zero,
        'ONE_TENTH',
        null
      )

      let newOriginQuery = { ...originQuery }
      newOriginQuery.minAmountOut = originMinWithSlippage

      let newDestQuery = { ...destQuery }
      newDestQuery.minAmountOut = destMinWithSlippage

      dispatch(
        setBridgeQuote({
          outputAmount: toValueBigNum,
          outputAmountString: commify(
            formatBNToString(toValueBigNum, toToken.decimals[toChainId], 8)
          ),
          routerAddress,
          allowance,
          exchangeRate: calculateExchangeRate(
            fromValue.sub(adjustedFeeAmount),
            fromToken.decimals[fromChainId],
            toValueBigNum,
            toToken.decimals[toChainId]
          ),
          feeAmount,
          delta: maxAmountOut,
          quotes: {
            originQuery: newOriginQuery,
            destQuery: newDestQuery,
          },
        })
      )
      return
    } catch {
      const str = formatBNToString(
        fromValue,
        fromToken.decimals[fromChainId],
        4
      )
      const message = `No route found for bridging ${str} ${fromToken.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
      console.log(message)
      toast(message)

      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      return
    } finally {
      dispatch(setIsLoading(false))
    }
  }

  const approveTxn = async () => {
    approveToken(
      bridgeQuote?.routerAddress,
      fromChainId,
      fromToken?.addresses[fromChainId]
    ).then(() => setIsApproved(true)).catch((err) => {
      console.log(err);
    })
  }

  const executeBridge = async () => {
    try {
      const wallet = await fetchSigner({
        chainId: fromChainId,
      })

      const toAddress =
        destinationAddress && isAddress(destinationAddress)
          ? destinationAddress
          : address

      const data = await synapseSDK.bridge(
        toAddress,
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId as keyof Token['addresses']],
        fromValue,
        bridgeQuote.quotes.originQuery,
        bridgeQuote.quotes.destQuery
      )
      const payload =
        fromToken.addresses[fromChainId as keyof Token['addresses']] ===
          AddressZero ||
        fromToken.addresses[fromChainId as keyof Token['addresses']] === ''
          ? { data: data.data, to: data.to, value: fromValue }
          : data
      const tx = await wallet.sendTransaction(payload)

      try {
        await tx.wait()

        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setDestinationAddress(null))
        dispatch(setShowDestinationAddress(false))
        dispatch(updateFromValue(Zero))

        return tx
      } catch (error) {
        console.log(`Transaction failed with error: ${error}`)
      }
    } catch (error) {
      console.log('Error executing bridge', error)
      return txErrorHandler(error)
    }
  }

  const springClass = 'fixed z-50 w-full h-full bg-opacity-50'

  return (
    <LandingPageWrapper>
      <div className="flex flex-col items-center justify-center">
        <div className="flex items-center space-x-20">
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
            max-w-lg px-1 pb-0 mb-3 overflow-hidden
            transition-all duration-100 transform rounded-xl
            bg-bgBase md:px-6 lg:px-6 mt-5
          `}
        >
          <div ref={bridgeDisplayRef}>
            <Transition show={showFromTokenSlideOver} {...TRANSITION_PROPS}>
              <animated.div  className={springClass}>
                <TokenSlideOver
                  key="fromBlock"
                  isOrigin={true}
                  tokens={supportedFromTokens}
                  chainId={fromChainId}
                  selectedToken={fromToken}
                />{' '}
              </animated.div>
            </Transition>
            <Transition show={showToTokenSlideOver} {...TRANSITION_PROPS}>
              <animated.div  className={springClass}>
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
                  isOrigin={true}
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
            <Transition
              appear={true}
              unmount={false}
              show={!fromValue.eq(0)}
              {...SECTION_TRANSITION_PROPS}
            >
              <ExchangeRateInfo
                fromAmount={fromValue}
                toToken={toToken}
                exchangeRate={bridgeQuote?.exchangeRate}
                toChainId={toChainId}
                showGasDrop={true}
              />
            </Transition>
            {showDestinationAddress && (
              <DestinationAddressInput
                toChainId={toChainId}
                destinationAddress={destinationAddress}
              />
            )}
            <div className="mt-3 mb-3">
              <BridgeTransactionButton
                isApproved={isApproved}
                approveTxn={approveTxn}
                executeBridge={executeBridge}
              />
            </div>
          </div>
        </Card>
        <div className="text-left text-white max-w-1/4">
          <div className="underline">Your bridge quote</div>
          <div>
            {Object.entries(bridgeQuote).map(([key, value]) => (
              <div key={key}>{`${key}: ${value}`}</div>
            ))}
          </div>
        </div>
      </div>
    </LandingPageWrapper>
  )
}

// TODO: Refactor
// would like to refactor this as a function that
// takes fromChainId, fromToken only and returns rest

const findSupportedChainsAndTokens = (
  token: Token,
  positedToChain: number | undefined,
  positedToSymbol: string | undefined,
  fromChainId: number
) => {
  let newToChain =
    positedToChain && positedToChain !== fromChainId
      ? Number(positedToChain)
      : DEFAULT_TO_CHAIN
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
  let bridgeableToken: Token = positedToToken
  if (!bridgeableTokens.includes(positedToToken)) {
    bridgeableToken = bridgeableTokens[0]
  }

  return {
    bridgeableChainIds: bridgeableChains.map((chainId: string) =>
      Number(chainId)
    ),
    bridgeableTokens,
    bridgeableToken,
  }
}

export default StateManagedBridge
