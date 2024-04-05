import toast from 'react-hot-toast'
import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { isNull } from 'lodash'
import { useAppSelector } from '@/store/hooks'
import { useDispatch } from 'react-redux'
import { zeroAddress } from 'viem'
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
import { calculateGasCost } from '../../utils/calculateGasCost'

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

  const shortenedParsedBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.parsedBalance

  const fullParsedBalance = formatBigIntToString(
    balance,
    fromToken?.decimals[fromChainId]
  )

  const { gasData } = useAppSelector((state) => state.gasData)
  const { gasPrice, maxFeePerGas } = gasData?.formatted

  const { rawGasCost, formattedGasCost } = calculateGasCost(
    maxFeePerGas,
    200_000
  )

  const isNativeToken = fromToken?.addresses[fromChainId] === zeroAddress

  console.log('fullParsedBalance:', fullParsedBalance)
  console.log('formattedGasCost: ', formattedGasCost)

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

  const calculateMaxBridgeableGas = (parsedGasBalance) => {}

  const onMaxBridgeableBalance = useCallback(() => {
    if (formattedGasCost && isNativeToken) {
      const maxBalance = Number(fullParsedBalance) - formattedGasCost

      if (maxBalance < 0) {
        toast.error(`Balance is less than estimated gas fee.`, {
          id: 'not-enough-balance-popup',
          duration: 5000,
        })

        dispatch(
          updateFromValue(
            formatBigIntToString(0n, fromToken?.decimals[fromChainId])
          )
        )
      } else {
        dispatch(updateFromValue(maxBalance.toString()))
      }
    } else {
      dispatch(
        updateFromValue(
          formatBigIntToString(balance, fromToken?.decimals[fromChainId])
        )
      )
    }
  }, [
    shortenedParsedBalance,
    balance,
    fromChainId,
    fromToken,
    formattedGasCost,
    isNativeToken,
  ])

  const connectedStatus = useMemo(() => {
    if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    } else if (hasMounted && isConnected && fromChainId === chain.id) {
      return <ConnectedIndicator />
    } else if (hasMounted && isConnected && fromChainId !== chain.id) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }, [chain, fromChainId, isConnected, hasMounted])

  const isGasBalanceLessThanFees = () => {
    if (isNativeToken && formattedGasCost && fullParsedBalance) {
      const gasBalance = fullParsedBalance
      const gasFees = formattedGasCost

      return gasFees > parseFloat(gasBalance)
    } else {
      return false
    }
  }

  const showMaxButton = () => {
    if (!hasMounted || !isConnected) return false
    if (isNativeToken && isNull(formattedGasCost)) return false
    return true
  }

  console.log('showMaxOption(): ', showMaxButton())
  console.log('isGasBalanceLessThanFees: ', isGasBalanceLessThanFees())

  return (
    <div
      data-test-id="input-container"
      className="text-left rounded-md p-md bg-bgLight"
    >
      <div className="flex items-center justify-between mb-3">
        <FromChainSelector />
        {connectedStatus}
      </div>
      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center justify-between pl-md w-full h-16
            rounded-md border border-white border-opacity-20
          `}
        >
          <div className="flex items-center">
            <FromTokenSelector />
            <div className="flex flex-col justify-between ml-4">
              <div style={{ display: 'table' }}>
                <input
                  ref={inputRef}
                  value={showValue}
                  placeholder="0.0000"
                  pattern="^[0-9]*[.,]?[0-9]*$"
                  onChange={handleFromValueChange}
                  className={`
                    border-none bg-transparent max-w-[190px] p-0 font-medium text-opacity-80
                    placeholder:text-[#88818C]text-white text-xl md:text-2xl
                    focus:outline-none focus:ring-0 focus:border-none
                  `}
                  style={{ display: 'table-cell', width: '100%' }}
                  name="inputRow"
                  autoComplete="off"
                  minLength={1}
                  maxLength={79}
                  disabled={false}
                />
              </div>
              {hasMounted && isConnected && (
                <label
                  htmlFor="inputRow"
                  onClick={onMaxBalance}
                  className={`
                    text-xs text-white transition-all duration-150 transform-gpu
                    hover:text-opacity-70 hover:cursor-pointer
                  `}
                >
                  {shortenedParsedBalance ?? '0.0'}
                  <span className="text-opacity-50 text-secondaryTextColor">
                    {' '}
                    available
                  </span>
                </label>
              )}
            </div>
          </div>
          <div>
            {showMaxButton() && (
              <div className="m">
                <MiniMaxButton
                  disabled={
                    !balance || balance === 0n
                      ? true
                      : false || isGasBalanceLessThanFees()
                  }
                  onClickBalance={onMaxBridgeableBalance}
                />
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

const AvailableBalance = (
  parsedBalance: number,
  estimatedGasCost: number,
  isGasToken: boolean,
  onMaxAvailableBalance: () => void,
  onMaxBridgeableBalance: () => void
) => {
  if (isGasToken) {
    const hasEnoughGas = parsedBalance - estimatedGasCost > 0

    return (
      <label
        htmlFor="inputRow"
        onClick={onMaxBridgeableBalance}
        className={`
        text-xs text-white transition-all duration-150 transform-gpu
        hover:text-opacity-70 hover:cursor-pointer
      `}
      >
        {parsedBalance ?? '0.0'}
        <span className="text-opacity-50 text-secondaryTextColor">
          {' '}
          available
        </span>
      </label>
    )
  } else {
    return (
      <label
        htmlFor="inputRow"
        onClick={onMaxAvailableBalance}
        className={`
        text-xs text-white transition-all duration-150 transform-gpu
        hover:text-opacity-70 hover:cursor-pointer
      `}
      >
        {parsedBalance ?? '0.0'}
        <span className="text-opacity-50 text-secondaryTextColor">
          {' '}
          available
        </span>
      </label>
    )
  }
}
