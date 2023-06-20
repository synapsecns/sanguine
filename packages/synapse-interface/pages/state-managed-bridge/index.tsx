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
  setShowFromTokenSlideOver,
  setShowToTokenSlideOver,
} from '../../slices/bridgeSlice'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'

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
import { BRIDGABLE_TOKENS, BRIDGE_CHAINS_BY_TYPE } from '@/constants/tokens'
import { CHAINS_BY_ID } from '@/constants/chains'
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
import { sortByTokenBalance } from '@/utils/sortTokens'

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
  return tokens
}

const sortToTokens = (tokens: Token[]) => {
  return tokens
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
    showFromTokenSlideOver,
    showToTokenSlideOver,
    supportedFromTokens,
    supportedToTokens,
  } = useSelector((state: RootState) => state.bridge)

  const [isApproved, setIsApproved] = useState(false)

  const dispatch = useDispatch()

  const fromChainIds = Object.keys(CHAINS_BY_ID).map((id) => Number(id))
  const toChainIds = BRIDGE_CHAINS_BY_TYPE[fromToken.swapableType]
    .filter((chainId) => Number(chainId) !== fromChainId)
    .map((chainId) => Number(chainId))

  // Can be smarter about breaking out which calls happen assoc with which
  // dependencies (like some stuff should only change on fromChainId changes)
  useEffect(() => {
    const fromTokens = BRIDGABLE_TOKENS[fromChainId]
    const toTokens = BRIDGABLE_TOKENS[toChainId]

    dispatch(setSupportedFromTokens(fromTokens))
    dispatch(setSupportedToTokens(toTokens))

    sortByTokenBalance(fromTokens, fromChainId, address).then((res) => {
      dispatch(setSupportedFromTokenBalances(res))
    })

    dispatch(setFromChainIds(fromChainIds))
    dispatch(setToChainIds(toChainIds))
    getAndSetBridgeQuote()
  }, [fromChainId, toChainId, fromToken, toToken, fromValue])

  // don't like this, rewrite: could be custom hook
  useEffect(() => {
    if (fromToken?.addresses[fromChainId] === AddressZero) {
      setIsApproved(true)
    } else {
      if (bridgeQuote?.allowance && fromValue.lt(bridgeQuote.allowance)) {
        setIsApproved(true)
      } else {
        setIsApproved(false)
      }
    }
  }, [bridgeQuote, fromToken, fromValue, fromChainId, toChainId])

  const handleFromTokenChange = (
    event: React.ChangeEvent<HTMLSelectElement>
  ) => {
    const selectedToken = supportedFromTokens.find(
      (token) => token.name === event.target.value
    )
    dispatch(setFromToken(selectedToken))
  }

  const handleToTokenChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedToken = supportedToTokens.find(
      (token) => token.name === event.target.value
    )
    dispatch(setToToken(selectedToken))
  }

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    let fromValueString = event.target.value
    try {
      let fromValueBigNumber = stringToBigNum(
        fromValueString,
        fromToken.decimals[fromChainId]
      )
      dispatch(updateFromValue(fromValueBigNumber))
    } catch (error) {
      console.error('Invalid value for conversion to BigNumber')
    }
  }

  const handleFromChainChange = (
    event: React.ChangeEvent<HTMLSelectElement>
  ) => {
    let fromChainId = Number(event.target.value)
    try {
      dispatch(setFromChainId(fromChainId))
    } catch (error) {
      console.log(`error`, error)
    }
  }

  const handleToChainChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    let toChainId = Number(event.target.value)

    try {
      dispatch(setToChainId(toChainId))
    } catch (error) {
      console.log(`error`, error)
    }
  }

  // Would like to move this into function outside of this component
  const getAndSetBridgeQuote = async () => {
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

      const allowance =
        fromToken.addresses[fromChainId] === AddressZero ||
        address === undefined
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
      dispatch(setIsLoading(false))
      return
    } catch {
      const str = formatBNToString(
        fromValue,
        fromToken.decimals[fromChainId],
        4
      )
      const message = `No route found for bridging ${fromToken.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name} for ${str}`
      toast(message)
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setIsLoading(false))
      return
    }
  }

  // Would like to move this function outside of this component

  const approveTxn = async () => {
    approveToken(
      bridgeQuote?.routerAddress,
      fromChainId,
      fromToken?.addresses[fromChainId]
    ).then(() => setIsApproved(true))
  }

  // Would like to move this into function outside of this component
  const executeBridge = async () => {
    try {
      const wallet = await fetchSigner({
        chainId: fromChainId,
      })
      const data = await synapseSDK.bridge(
        address,
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
        return tx
      } catch (error) {
        console.log(`Transaction failed with error: ${error}`)
      }
    } catch (error) {
      console.log('Error executing bridge', error)
      return txErrorHandler(error)
    }
  }

  return (
    <LandingPageWrapper>
      <div className="flex justify-center">
        <div className="text-white">
          <PageHeader
            title="Redux State Managed Bridge"
            subtitle="Send your assets across chains."
          />
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
                <animated.div>
                  <TokenSlideOver
                    key="fromBlock"
                    isOrigin={true}
                    tokens={supportedFromTokens}
                    chainId={fromChainId}
                    selectedToken={fromToken}
                    setToken={setFromToken}
                    setShowSlideOver={setShowFromTokenSlideOver}
                  />{' '}
                </animated.div>
              </Transition>
              <Transition show={showToTokenSlideOver} {...TRANSITION_PROPS}>
                <animated.div>
                  <TokenSlideOver
                    key="toBlock"
                    isOrigin={false}
                    tokens={supportedToTokens}
                    chainId={toChainId}
                    selectedToken={toToken}
                    setToken={setToToken}
                    setShowSlideOver={setShowToTokenSlideOver}
                  />{' '}
                </animated.div>
              </Transition>
              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <div>fromChain</div>
                  <div>
                    <select
                      className="text-black"
                      onChange={handleFromChainChange}
                      value={fromChainId}
                    >
                      {sortFromChainIds(fromChainIds).map((chainId) => (
                        <option key={chainId} value={chainId}>
                          {CHAINS_BY_ID[chainId]?.name}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>
                <div className="flex items-center justify-between">
                  <div>fromToken</div>
                  <select
                    className="text-black"
                    onChange={handleFromTokenChange}
                    value={fromToken?.name}
                  >
                    {sortFromTokens(supportedFromTokens).map((token) => (
                      <option key={token.name} value={token.name}>
                        {token.symbol}
                      </option>
                    ))}
                  </select>
                </div>
                <div className="flex items-center justify-between">
                  <div>from amount</div>
                  <input
                    type="text"
                    onChange={handleFromValueChange}
                    className="text-black"
                    placeholder="Enter value"
                  />
                </div>
                <div className="flex items-center justify-between">
                  <div>toChain</div>
                  <div>
                    <select
                      className="text-black"
                      onChange={handleToChainChange}
                      value={toChainId}
                    >
                      {sortToChainIds(toChainIds).map((chainId) => (
                        <option key={chainId} value={chainId}>
                          {CHAINS_BY_ID[chainId]?.name}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>
                <div className="flex items-center justify-between">
                  <div>toToken</div>
                  <select
                    className="text-black"
                    onChange={handleToTokenChange}
                    value={toToken?.name}
                  >
                    {sortToTokens(supportedToTokens).map((token) => (
                      <option key={token.name} value={token.name}>
                        {token.symbol}
                      </option>
                    ))}
                  </select>
                </div>
                <div className="flex items-center justify-between">
                  <div>Output amount</div>
                  <div>{bridgeQuote?.outputAmountString}</div>
                </div>
                <h1 className="text-2xl">UI experimentation below</h1>
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
                <div>
                  {!isApproved ? (
                    <button
                      className="p-2 bg-blue-500 disabled:opacity-50"
                      onClick={approveTxn}
                      disabled={
                        isLoading ||
                        bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
                        bridgeQuote === EMPTY_BRIDGE_QUOTE
                      }
                    >
                      Approve
                    </button>
                  ) : (
                    <button
                      className="p-2 bg-blue-500 disabled:opacity-50"
                      onClick={executeBridge}
                      disabled={
                        isLoading ||
                        bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
                        bridgeQuote === EMPTY_BRIDGE_QUOTE
                      }
                    >
                      Bridge
                    </button>
                  )}
                </div>
              </div>
            </div>
          </Card>
          <div className="max-w-1/4">
            <div className="underline">Your bridge quote</div>
            <div>
              {Object.entries(bridgeQuote).map(([key, value]) => (
                <div key={key}>{`${key}: ${value}`}</div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </LandingPageWrapper>
  )
}

export default StateManagedBridge
