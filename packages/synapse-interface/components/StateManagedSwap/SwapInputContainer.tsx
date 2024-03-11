import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount, useNetwork } from 'wagmi'

import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { SwapChainSelector } from './SwapChainSelector'
import { SwapFromTokenSelector } from './SwapFromTokenSelector'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { initialState, updateSwapFromValue } from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'
import { BridgeCardTokenInput, BridgeContainer } from '../ui/BridgeCard'

export const SwapInputContainer = () => {
  const inputRef = useRef<HTMLInputElement>(null)
  const { swapChainId, swapFromToken, swapToToken, swapFromValue } =
    useSwapState()
  const [showValue, setShowValue] = useState('')

  const [hasMounted, setHasMounted] = useState(false)

  const { balances } = usePortfolioState()

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const { isConnected } = useAccount()
  const { chain } = useNetwork()

  const dispatch = useDispatch()

  const tokenData = balances[swapChainId]?.find(
    (token) => token.tokenAddress === swapFromToken?.addresses[swapChainId]
  )

  const parsedBalance = tokenData?.parsedBalance

  const balance = tokenData?.balance

  useEffect(() => {
    if (
      swapFromToken &&
      swapFromToken.decimals[swapChainId] &&
      stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]) !== 0n
    ) {
      setShowValue(swapFromValue)
    }

    if (swapFromValue === initialState.swapFromValue) {
      setShowValue(initialState.swapFromValue)
    }
  }, [swapFromValue, swapChainId, swapFromToken])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const swapFromValueString: string = cleanNumberInput(event.target.value)
    try {
      dispatch(updateSwapFromValue(swapFromValueString))
      setShowValue(swapFromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        dispatch(updateSwapFromValue(''))
        setShowValue(inputValue)
      }
    }
  }

  const onMaxBalance = useCallback(() => {
    dispatch(
      updateSwapFromValue(
        formatBigIntToString(balance, swapFromToken.decimals[swapChainId])
      )
    )
  }, [balance, swapChainId, swapFromToken])

  const connectedStatus = useMemo(() => {
    if (hasMounted && isConnected) {
      if (swapChainId === chain.id) {
        return <ConnectedIndicator />
      } else if (swapChainId !== chain.id) {
        return <ConnectToNetworkButton chainId={swapChainId} />
      }
    } else if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    }
  }, [chain, swapChainId, isConnected, hasMounted])

  return (
    <BridgeContainer>
      <div className="flex items-center justify-between">
        <SwapChainSelector />
        {connectedStatus}
      </div>
      <BridgeCardTokenInput>
        <SwapFromTokenSelector />
        <div>
          <input
            ref={inputRef}
            pattern="^[0-9]*[.,]?[0-9]*$"
            disabled={false}
            className={`
              focus:outline-none focus:ring-0 focus:border-none
              border-none
              bg-transparent
              p-0 block
              placeholder:text-[#88818C]
              text-white text-opacity-80 text-xl md:text-2xl font-medium
            `}
            placeholder="0.0000"
            onChange={handleFromValueChange}
            value={showValue}
            name="inputRow"
            autoComplete="off"
            minLength={1}
            maxLength={79}
          />
          {hasMounted && isConnected && (
            <label
              htmlFor="inputRow"
              className="text-xs text-white transition-all duration-150 transform-gpu hover:text-opacity-70 hover:cursor-pointer block"
              onClick={onMaxBalance}
            >
              {parsedBalance ?? '0.0'}
              <span className="text-opacity-50 text-secondaryTextColor">
                {' '}
                available
              </span>
            </label>
          )}
        </div>
        {hasMounted && isConnected && (
          <MiniMaxButton
            disabled={!balance || balance === 0n ? true : false}
            onClickBalance={onMaxBalance}
          />
        )}
      </BridgeCardTokenInput>
    </BridgeContainer>
  )
}
