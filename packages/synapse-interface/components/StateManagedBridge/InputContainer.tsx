import React, { useEffect, useRef, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount, useNetwork } from 'wagmi'
import { zeroAddress } from 'viem'

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
import { useBridgeValidations } from '@/utils/hooks/useBridgeValidations'
import { useAppDispatch } from '@/store/hooks'

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

  const onMaxBalance = () => {
    dispatch(
      updateFromValue(
        formatBigIntToString(balance, fromToken?.decimals[fromChainId])
      )
    )
  }

  const connectedStatus = () => {
    if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    } else if (hasMounted && isConnected && fromChainId === chain.id) {
      return <ConnectedIndicator />
    } else if (hasMounted && isConnected && fromChainId !== chain.id) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }

  return (
    <div
      data-test-id="input-container"
      className="text-left rounded-md p-md bg-bgLight"
    >
      <div className="flex items-center justify-between mb-3">
        <FromChainSelector />
        {connectedStatus()}
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
              {hasMounted && isConnected && <ShowLabel />}
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

const ShowLabel = () => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken, bridgeQuote, isLoading } = useBridgeState()
  const { hasEnoughBalance, hasInputAmount, hasEnoughApproved, hasValidRoute } =
    useBridgeValidations()

  const { balances } = usePortfolioState()

  const labelContentRef = useRef('')

  const token = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )
  const parsedBalance = token?.parsedBalance
  const balance = token?.balance
  const allowance = bridgeQuote.allowance

  const onMaxBalance = () => {
    dispatch(
      updateFromValue(
        formatBigIntToString(balance, fromToken?.decimals[fromChainId])
      )
    )
  }

  const onMaxApproved = () => {
    dispatch(
      updateFromValue(
        formatBigIntToString(allowance, fromToken?.decimals[fromChainId])
      )
    )
  }

  let labelContent
  let onClick
  let hasError = false

  if (hasInputAmount && !hasValidRoute) {
    labelContent = `${parsedBalance ?? '0'} available`
    onClick = onMaxBalance
    hasError = true
  } else if (hasEnoughBalance && !hasEnoughApproved) {
    if (allowance === 0n) {
      labelContent = `0 approved`
      onClick = onMaxBalance
      hasError = true
    } else {
      labelContent = `${
        formatBigIntToString(allowance, fromToken?.decimals[fromChainId], 5) ??
        '0'
      } approved`
      onClick = onMaxApproved
      hasError = true
    }
  } else if (
    (hasInputAmount && !hasEnoughBalance) ||
    fromToken?.addresses[fromChainId] === zeroAddress
  ) {
    labelContent = `${parsedBalance ?? '0'} available`
    onClick = onMaxBalance
    hasError = true
  } else {
    labelContent = `${parsedBalance ?? '0'} available`
    onClick = onMaxBalance
    hasError = false
  }

  useEffect(() => {
    if (!isLoading) {
      labelContentRef.current = labelContent
    }
  }, [isLoading, labelContent])

  const labelClassName = `text-xs hover:text-opacity-70 hover:cursor-pointer ${
    hasError ? 'text-amber-200' : 'text-secondaryTextColor'
  }`

  if (isLoading) {
    return (
      <label htmlFor="inputRow" className={labelClassName}>
        {labelContentRef.current}
      </label>
    )
  }

  return (
    <label htmlFor="inputRow" onClick={onClick} className={labelClassName}>
      {labelContent}
    </label>
  )
}
