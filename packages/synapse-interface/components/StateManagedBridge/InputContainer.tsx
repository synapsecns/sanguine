import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount, useNetwork } from 'wagmi'

import { initialState, updateFromValue } from '@/slices/bridge/reducer'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { FromChainSelector } from './FromChainSelector'
import { FromTokenSelector } from './FromTokenSelector'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const { fromChainId, fromToken, fromValue } = useBridgeState()
  const [showValue, setShowValue] = useState('')

  const [hasMounted, setHasMounted] = useState(false)

  const { balances } = usePortfolioState()

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const { isConnected } = useAccount()
  const { chain } = useNetwork()

  const dispatch = useDispatch()

  const parsedBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.parsedBalance

  const balance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.balance

  useEffect(() => {
    if (fromToken && fromToken?.decimals[fromChainId]) {
      setShowValue(fromValue)
    }

    if (fromValue === initialState.fromValue) {
      setShowValue(initialState.fromValue)
    }
  }, [fromValue, inputRef, fromChainId, fromToken])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const fromValueString: string = cleanNumberInput(event.target.value)
    try {
      dispatch(updateFromValue(fromValueString))
      setShowValue(fromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        dispatch(updateFromValue(inputValue))
        setShowValue(inputValue)
      }
    }
  }

  const onMaxBalance = useCallback(() => {
    dispatch(
      updateFromValue(
        formatBigIntToString(balance, fromToken?.decimals[fromChainId])
      )
    )
  }, [balance, fromChainId, fromToken])

  const connectedStatus = useMemo(() => {
    if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    } else if (hasMounted && isConnected && fromChainId === chain.id) {
      return <ConnectedIndicator />
    } else if (hasMounted && isConnected && fromChainId !== chain.id) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }, [chain, fromChainId, isConnected, hasMounted])

  return (
    <div
      data-test-id="input-container"
      className="text-left rounded-md p-md bg-bgBase/10"
    >
      <div className="flex items-center justify-between mb-3">
        <FromChainSelector />
        {connectedStatus}
      </div>
      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center justify-between
            pl-md
            w-full h-16
            rounded-md
            border border-white border-opacity-20
          `}
        >
          <div className="flex items-center">
            <FromTokenSelector />
            <div className="flex flex-col justify-between ml-4">
              <div style={{ display: 'table' }}>
                <input
                  ref={inputRef}
                  pattern="^[0-9]*[.,]?[0-9]*$"
                  disabled={false}
                  className={`
                    focus:outline-none
                    focus:ring-0
                    focus:border-none
                    border-none
                    bg-transparent
                    max-w-[190px]
                    p-0
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
                  style={{ display: 'table-cell', width: '100%' }}
                />
              </div>
              {hasMounted && isConnected && (
                <label
                  htmlFor="inputRow"
                  className="text-xs text-white transition-all duration-150 transform-gpu hover:text-opacity-70 hover:cursor-pointer"
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
          </div>
          <div>
            {hasMounted && isConnected && (
              <div className="m">
                <MiniMaxButton
                  disabled={!balance || balance === 0n ? true : false}
                  onClickBalance={onMaxBalance}
                />
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
