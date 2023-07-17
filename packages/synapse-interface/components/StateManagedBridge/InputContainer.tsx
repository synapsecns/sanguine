import React, { useEffect, useState, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useAccount } from 'wagmi'
import { RootState } from '@/store/store'

import { updateFromValue } from '@/slices/bridgeSlice'
import { setShowFromTokenSlideOver } from '@/slices/bridgeDisplaySlice'
import SelectTokenDropdown from '@/components/input/TokenAmountInput/SelectTokenDropdown'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { OriginChainLabel } from './OriginChainLabel'
import { cleanNumberInput } from '@/utils/cleanNumberInput'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const {
    fromChainId,
    fromToken,
    fromChainIds,
    supportedFromTokenBalances,
    fromValue,
    bridgeTxHashes,
  } = useSelector((state: RootState) => state.bridge)
  const [showValue, setShowValue] = useState('')

  const [hasMounted, setHasMounted] = useState(false)
  const previousBridgeTxHashesRef = useRef<string[]>([])

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

  const hasBalances = Object.keys(supportedFromTokenBalances).length > 0

  const fromTokenBalance: bigint =
    (hasBalances &&
      supportedFromTokenBalances.filter((token) => token.token === fromToken)[0]
        ?.balance) ??
    0n

  const formattedBalance = hasBalances
    ? formatBigIntToString(fromTokenBalance, fromToken.decimals[fromChainId], 4)
    : '0'

  useEffect(() => {
    if (
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) !== 0n &&
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) ===
        fromTokenBalance
    ) {
      setShowValue(fromValue)
    }
  }, [fromValue, inputRef, fromChainId, fromToken, fromTokenBalance])

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
    const str = formatBigIntToString(
      fromTokenBalance,
      fromToken.decimals[fromChainId],
      4
    )
    dispatch(updateFromValue(str))
    setShowValue(
      formatBigIntToString(fromTokenBalance, fromToken.decimals[fromChainId])
    )
  }

  return (
    <div
      data-test-id="input-container"
      className={`
        text-left px-2 sm:px-4 pt-2 pb-4 rounded-xl
        bg-bgLight
      `}
    >
      <div>
        <div className="pt-1 pb-3">
          <OriginChainLabel
            chainId={fromChainId}
            chains={fromChainIds}
            connectedChainId={fromChainId}
          />
        </div>
      </div>
      <div className="flex h-16 mb-4 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-3 sm:pl-4
            w-full h-20
            rounded-xl
            border border-white border-opacity-20
          `}
        >
          <SelectTokenDropdown
            chainId={fromChainId}
            selectedToken={fromToken}
            isOrigin={true}
            onClick={() => dispatch(setShowFromTokenSlideOver(true))}
          />
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
                {formattedBalance}
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
                disabled={fromTokenBalance && fromTokenBalance === 0n}
                onClickBalance={onClickBalance}
              />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
