import React, { useEffect, useState, useRef } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'

import { updateFromValue } from '@/slices/bridge/reducer'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import FromChainSelect from './FromChainSelect'
import FromTokenSelect from './FromTokenSelect'
import { useAppSelector } from '@/store/hooks'

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
        text-left px-2 sm:px-4 pt-4 pb-1 rounded-xl
        bg-bgLight
      `}
    >
      <FromChainSelect />
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
          <FromTokenSelect />
          <div className="flex flex-col pt-2 ml-4">
            <input
              ref={inputRef}
              pattern="^[0-9]*[.,]?[0-9]*$"
              disabled={false}
              className={`
              focus:outline-none
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
                className="hidden text-xs text-white transition-all duration-150 md:block transform-gpu hover:text-opacity-70 hover:cursor-pointer"
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
