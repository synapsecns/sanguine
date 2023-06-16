import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useAccount } from 'wagmi'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../../store/store'

import {
  setFromToken,
  setToToken,
  updateFromValue,
  setBridgeQuote,
  setIsLoading,
} from '../../slices/bridgeSlice'
import { ETH } from '@/constants/tokens/master'
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
import { useEffect } from 'react'
import { Token } from '@/utils/types'
import { fetchSigner } from '@wagmi/core'
import { txErrorHandler } from '@/utils/txErrorHandler'

const StateManagedBridge = () => {
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()

  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    bridgeQuote,
    fromValue,
    isLoading,
  } = useSelector((state: RootState) => state.bridge)

  const dispatch = useDispatch()

  const tokens = [ETH]

  useEffect(() => {
    getAndSetBridgeQuote()
  }, [fromChainId, toChainId, fromToken, toToken, fromValue])

  const handleFromTokenChange = (
    event: React.ChangeEvent<HTMLSelectElement>
  ) => {
    const selectedToken = tokens.find(
      (token) => token.name === event.target.value
    )
    dispatch(setFromToken(selectedToken))
  }

  const handleToTokenChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedToken = tokens.find(
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
    } catch (error) {
      console.log(error)
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setIsLoading(false))
      return
    }
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
      <div className="flex justify-center text-white">
        <div className="space-y-1">
          <div className="mb-5 text-xl">Redux State Managed Bridge</div>
          <div className="flex items-center justify-between">
            <div>fromChainId</div>
            <div>{fromChainId}</div>
          </div>
          <div className="flex items-center justify-between">
            <div>fromToken</div>
            <select
              className="text-black"
              onChange={handleFromTokenChange}
              value={fromToken?.name}
            >
              {tokens.map((token) => (
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
            <div>toChainId</div>
            <div>{toChainId}</div>
          </div>
          <div className="flex items-center justify-between">
            <div>toToken</div>
            <select
              className="text-black"
              onChange={handleToTokenChange}
              value={toToken.name}
            >
              {tokens.map((token) => (
                <option key={token.name} value={token.name}>
                  {token.symbol}
                </option>
              ))}
            </select>
          </div>
          <div>
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
          </div>
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
