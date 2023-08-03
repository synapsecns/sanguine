import React, { useEffect, useState, useRef } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'

import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { useAppSelector } from '@/store/hooks'
import { ConnectedIndicator } from './ConnectedIndicator'
import { FromChainSelector } from './FromChainSelector'
import { FromTokenSelector } from './FromTokenSelector'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const { fromChainId, fromToken, fromValue, bridgeTxHashes } = useAppSelector(
    (state) => state.bridge
  )
  const [showValue, setShowValue] = useState('')

  const [hasMounted, setHasMounted] = useState(false)
  const previousBridgeTxHashesRef = useRef<string[]>([])

  const { balancesAndAllowances } = useAppSelector((state) => state.portfolio)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    const previousBridgeTxHashes = previousBridgeTxHashesRef.current

    if (bridgeTxHashes.length !== previousBridgeTxHashes.length) {
      setShowValue('')
    }

    previousBridgeTxHashesRef.current = bridgeTxHashes
  }, [bridgeTxHashes])

  const { isConnected } = useAccount()

  const dispatch = useDispatch()

  const parsedBalance = balancesAndAllowances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )?.parsedBalance

  useEffect(() => {
    if (
      fromToken &&
      fromToken.decimals[fromChainId] &&
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) !== 0n
      // stringToBigInt(fromValue, fromToken.decimals[fromChainId]) ===
      //   fromTokenBalance
    ) {
      setShowValue(fromValue)
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

  const onClickBalance = () => {
    dispatch(updateFromValue(parsedBalance))
    setShowValue(parsedBalance)
  }

  return (
    <div
      data-test-id="input-container"
      className={`
        text-left px-2 sm:px-4 pt-4 mt-2 pb-1 rounded-xl
        bg-bgLight
      `}
    >
      <div className="flex justify-between mb-3">
        <div className="flex items-center space-x-2">
          <FromChainSelector />
          {(fromChainId || fromToken) && (
            <button
              className="bg-bgLight text-primaryTextColor border-[1px] p-1 rounded-md text-xxs"
              onClick={() => {
                dispatch(setFromChainId(null))
                dispatch(setFromToken(null))
              }}
            >
              clear
            </button>
          )}
        </div>
        {hasMounted && isConnected && <ConnectedIndicator />}
      </div>
      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-3 sm:pl-4
            w-full h-16
            rounded-xl
            border border-white border-opacity-20
          `}
        >
          <FromTokenSelector />
          <div className="flex flex-col ml-4">
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
                max-w-[100px]
                md:max-w-[160px]
                placeholder:text-[#88818C]
                text-white text-opacity-80 text-lg md:text-2xl font-medium
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
                className="hidden ml-3 -mt-2 text-xs text-white transition-all duration-150 md:block transform-gpu hover:text-opacity-70 hover:cursor-pointer"
                onClick={onClickBalance}
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
            <div className="m-auto">
              <MiniMaxButton
                disabled={parsedBalance === '0.0'}
                onClickBalance={onClickBalance}
              />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
