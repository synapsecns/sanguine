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
import { ChainSelector } from './ChainSelector'
import { FromTokenSelector } from './FromTokenSelector'
import { TokenSelector } from './TokenSelector'
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
      className="text-left rounded-md p-2 flex flex-col gap-2 bg-white dark:bg-zinc-700 border border-zinc-200 dark:border-transparent"
    >
      <div className="flex items-center justify-between">
        <ChainSelector side="from" />
        {connectedStatus}
      </div>
      <div
        className={`
          flex gap-2 items-center
          px-2 py-1 min-h-[4rem] rounded-md border
          border-zinc-200 dark:border-zinc-600
        `}
      >
        <TokenSelector side="from" />
        <div className="flex flex-col">
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
              w-full
              p-0
              placeholder:text-zinc-400 placeholder:dark:text-zinc-500
              text-xl xs:text-2xl font-medium block
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
              className="text-sm opacity-50 hover:opacity-100 hover:cursor-pointer"
              onClick={onMaxBalance}
            >
              {parsedBalance ?? '0.0'} available
            </label>
          )}
        </div>
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
  )
}
