import { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Zero } from '@ethersproject/constants'
import { RootState } from '@/store/store'

import { updateFromValue } from '@/slices/bridgeSlice'
import { setShowFromTokenSlideOver } from '@/slices/bridgeDisplaySlice'
import { stringToBigNum } from '@/utils/stringToBigNum'
import SelectTokenDropdown from '@/components/input/TokenAmountInput/SelectTokenDropdown'
import { useAccount } from 'wagmi'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBNToString } from '@/utils/bignumber/format'
import { OriginChainLabel } from './OriginChainLabel'

export const InputContainer = () => {
  const { fromChainId, fromToken, fromChainIds, supportedFromTokenBalances } =
    useSelector((state: RootState) => state.bridge)
  const [showValue, setShowValue] = useState('')

  const [hasMounted, setHasMounted] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const { isConnected } = useAccount()

  const dispatch = useDispatch()

  const hasBalances = Object.keys(supportedFromTokenBalances).length > 0

  const fromTokenBalance =
    (hasBalances &&
      supportedFromTokenBalances.filter((token) => token.token === fromToken)[0]
        ?.balance) ??
    Zero

  const formattedBalance = hasBalances
    ? formatBNToString(fromTokenBalance, fromToken.decimals[fromChainId], 4)
    : '0'

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
      setShowValue(fromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigNumber')
    }
  }

  const onClickBalance = () => {
    dispatch(updateFromValue(fromTokenBalance))
    setShowValue(
      formatBNToString(fromTokenBalance, fromToken.decimals[fromChainId])
    )
  }

  return (
    <div
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
          <input
            pattern="[0-9.]+"
            disabled={false}
            className={`
              ml-4
              ${isConnected ? '-mt-0 md:-mt-4' : '-mt-0'}
              focus:outline-none
              bg-transparent
              pr-4
              w-2/3
              placeholder:text-[#88818C]
              text-white text-opacity-80 text-lg md:text-2xl lg:text-2xl font-medium
            `}
            placeholder="0.0000"
            onChange={handleFromValueChange}
            value={showValue}
            name="inputRow"
            autoComplete="off"
          />
          {hasMounted && isConnected && (
            <label
              htmlFor="inputRow"
              className="absolute hidden pt-1 mt-10 ml-40 text-xs text-white transition-all duration-150 md:block transform-gpu hover:text-opacity-70 hover:cursor-pointer"
              onClick={onClickBalance}
            >
              {formattedBalance}
              <span className="text-opacity-50 text-secondaryTextColor">
                {' '}
                available
              </span>
            </label>
          )}
          {hasMounted && isConnected && (
            <div className="hidden mr-2 sm:inline-block">
              <MiniMaxButton
                disabled={fromTokenBalance && fromTokenBalance.eq(Zero)}
                onClickBalance={onClickBalance}
              />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
