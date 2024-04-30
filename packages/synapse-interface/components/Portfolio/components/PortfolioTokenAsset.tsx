import React, { useCallback } from 'react'
import _ from 'lodash'
import { useAppDispatch } from '@/store/hooks'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { inputRef } from '../../StateManagedBridge/InputContainer'
import Image from 'next/image'
import { useBridgeState } from '@/slices/bridge/hooks'
import { PortfolioAssetActionButton } from './PortfolioAssetActionButton'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { zeroAddress } from 'viem'
import GasIcon from '@/components/icons/GasIcon'
import { HoverTooltip } from '../../HoverTooltip'
import { getParsedBalance } from '@/utils/getParsedBalance'

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
  const { icon, symbol, decimals, addresses } = token

  const tokenAddress = addresses[portfolioChainId]
  const tokenDecimals = _.isNumber(decimals)
    ? decimals
    : decimals[portfolioChainId]

  const parsedBalance = getParsedBalance(balance, tokenDecimals, 3)
  const parsedBalanceLong = getParsedBalance(balance, tokenDecimals, 8)

  const isDisabled = false
  const isTokenSelected =
    fromToken === token && fromChainId === portfolioChainId
  const isGasToken = tokenAddress === zeroAddress

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
      >
        <Image
          loading="lazy"
          alt={`${symbol} img`}
          className="w-6 h-6 rounded-md"
          src={icon}
        />
        <HoverTooltip
          hoverContent={
            <div className="whitespace-nowrap">
              {parsedBalanceLong} {symbol}
            </div>
          }
        >
          <div>
            {parsedBalance} {symbol}
          </div>
        </HoverTooltip>

        {isGasToken ? (
          <HoverTooltip
            hoverContent={<div className="whitespace-nowrap">Gas token</div>}
          >
            <GasIcon className="pt-0.5 m-auto fill-secondary" />
          </HoverTooltip>
        ) : null}
      </div>
      <PortfolioAssetActionButton
        selectCallback={handleFromSelectionCallback}
        isDisabled={isDisabled || isTokenSelected}
        isSelected={isTokenSelected}
      />
    </div>
  )
}
