import React, { useCallback } from 'react'
import _ from 'lodash'
import { useAppDispatch } from '@/store/hooks'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { inputRef } from '../../StateManagedBridge/InputContainer'
import Image from 'next/image'
import { useBridgeState } from '@/slices/bridge/hooks'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { PortfolioAssetActionButton } from './PortfolioAssetActionButton'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'

const handleFocusOnBridgeInput = () => {
  inputRef.current?.focus()
}

type PortfolioTokenAssetProps = {
  token: Token
  balance: bigint
  portfolioChainId: number
  connectedChainId: number
}

export const PortfolioTokenAsset = ({
  token,
  balance,
  portfolioChainId,
  connectedChainId,
}: PortfolioTokenAssetProps) => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken } = useBridgeState()
  const { icon, symbol, decimals } = token

  const tokenDecimals = _.isNumber(decimals)
    ? decimals
    : decimals[portfolioChainId]

  const parsedBalance = getParsedBalance(balance, tokenDecimals, 3)
  const parsedBalanceLong = getParsedBalance(balance, tokenDecimals, 8)

  const isDisabled = false
  const isTokenSelected =
    fromToken === token && fromChainId === portfolioChainId

  const handleFromSelectionCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId))
    dispatch(setFromToken(token))
    handleFocusOnBridgeInput()
    dispatch(
      updateFromValue(
        trimTrailingZeroesAfterDecimal(getParsedBalance(balance, tokenDecimals))
      )
    )
  }, [token, balance, portfolioChainId])

  return (
    <div
      id="portfolio-token-asset"
      className={`
        p-2 flex items-center border-y text-white justify-between last:rounded-b-md
        ${isTokenSelected ? 'bg-tint border-surface' : 'border-transparent'}
      `}
    >
      <div
        onClick={handleFromSelectionCallback}
        className={`
          flex items-center gap-2
          pl-2 pr-4 py-2 cursor-pointer rounded
          hover:bg-surface active:opacity-70
        `}
        title={`${parsedBalanceLong} ${symbol}`}
      >
        <Image
          loading="lazy"
          alt={`${symbol} img`}
          className="w-6 h-6 rounded-md"
          src={icon}
        />
        {parsedBalance} {symbol}
      </div>
      <PortfolioAssetActionButton
        selectCallback={handleFromSelectionCallback}
        isDisabled={isDisabled || isTokenSelected}
        isSelected={isTokenSelected}
      />
    </div>
  )
}

const getParsedBalance = (
  balance: bigint,
  decimals: number,
  places?: number
) => {
  const formattedBalance = formatBigIntToString(balance, decimals, places)
  const verySmallBalance = balance > 0n && hasOnlyZeroes(formattedBalance)

  return verySmallBalance ? '< 0.001' : formattedBalance
}
