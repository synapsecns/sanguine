import React, { useMemo, useCallback, useState } from 'react'
import { useAppDispatch } from '@/store/hooks'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { inputRef } from '../../../StateManagedBridge/InputContainer'
import Image from 'next/image'
import { useBridgeState } from '@/slices/bridge/hooks'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'

const handleFocusOnInput = () => {
  inputRef.current.focus()
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
  const { icon, symbol, decimals } = token as Token

  const parsedBalance: string = useMemo(() => {
    const formattedBalance = formatBigIntToString(
      balance,
      decimals[portfolioChainId],
      3
    )
    return balance > 0n && hasOnlyZeroes(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const parsedBalanceLong: string = useMemo(() => {
    const formattedBalance = formatBigIntToString(
      balance,
      decimals[portfolioChainId],
      8
    )
    return balance > 0n && hasOnlyZeroes(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const isTokenSelected: boolean = useMemo(() => {
    return fromToken === token && fromChainId === portfolioChainId
  }, [fromChainId, fromToken, token, portfolioChainId])

  const isDisabled: boolean = false

  const handleTotalBalanceInputCallback = useCallback(async () => {
    await dispatch(setFromChainId(portfolioChainId as number))
    await dispatch(setFromToken(token as Token))
    await dispatch(
      await updateFromValue(
        formatBigIntToString(
          balance,
          token.decimals[portfolioChainId]
        ) as string
      )
    )
    handleFocusOnInput()
  }, [isDisabled, token, balance, portfolioChainId])

  const handleSelectFromTokenCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId as number))
    dispatch(setFromToken(token as Token))
    handleFocusOnInput()
  }, [token, isDisabled, portfolioChainId])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className={`
        p-2 flex items-center border-y text-white justify-between last:rounded-b-md
        ${isTokenSelected ? 'bg-tint border-surface' : 'border-transparent'}
      `}
    >
      <div
        onClick={handleTotalBalanceInputCallback}
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
        selectCallback={handleSelectFromTokenCallback}
        isDisabled={isDisabled || isTokenSelected}
        isSelected={isTokenSelected}
      />
    </div>
  )
}

export const HoverClickableText = ({
  defaultText,
  hoverText,
  callback,
}: {
  defaultText: string
  hoverText: string
  callback: () => void
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      data-test-id="hover-clickable-text"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      onClick={callback}
      className={`
        group px-2
        text-[#A3A3C2]
        hover:text-[#75E6F0]
        hover:underline
        hover:cursor-pointer
        active:opacity-70
      `}
    >
      <div className="text-sm">{isHovered ? hoverText : defaultText}</div>
    </div>
  )
}

type PortfolioAssetActionButtonProps = {
  selectCallback: () => void
  isDisabled: boolean
  isSelected: boolean
}

const PortfolioAssetActionButton = ({
  selectCallback,
  isDisabled,
  isSelected,
}: PortfolioAssetActionButtonProps) => {
  return (
    <React.Fragment>
      <button
        data-test-id="portfolio-asset-action-button"
        className={`
          py-1 px-6 rounded-sm
          border border-synapsePurple
          ${!isDisabled && 'cursor-pointer hover:bg-surface active:opacity-70'}
        `}
        onClick={selectCallback}
        disabled={isDisabled}
      >
        Select{isSelected && 'ed'}
      </button>
    </React.Fragment>
  )
}
