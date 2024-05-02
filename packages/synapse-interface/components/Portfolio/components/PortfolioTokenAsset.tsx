import React, { useCallback } from 'react'
import { zeroAddress } from 'viem'
import { isNumber } from 'lodash'
import Image from 'next/image'
import { useAppDispatch } from '@/store/hooks'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { inputRef } from '../../StateManagedBridge/InputContainer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { PortfolioAssetActionButton } from './PortfolioAssetActionButton'
import { HoverTooltip } from '@/components/HoverTooltip'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { useGasEstimator } from '@/utils/hooks/useGasEstimator'
import GasIcon from '@/components/icons/GasIcon'
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
  const { icon, symbol, decimals, addresses } = token

  const tokenAddress = addresses[portfolioChainId]
  const tokenDecimals = isNumber(decimals)
    ? decimals
    : decimals[portfolioChainId]

  const parsedBalance = getParsedBalance(balance, tokenDecimals, 3)
  const parsedBalanceLong = getParsedBalance(balance, tokenDecimals, 8)

  const isDisabled = false
  const isPortfolioChainSelected = fromChainId === portfolioChainId
  const isTokenSelected = isPortfolioChainSelected && fromToken === token
  const isGasToken = tokenAddress === zeroAddress

  const handleFromSelectionCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId))
    dispatch(setFromToken(token))
    handleFocusOnBridgeInput()
  }, [token, portfolioChainId])

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

        {isGasToken && (
          <HoverTooltip
            hoverContent={<div className="whitespace-nowrap">Gas token</div>}
          >
            <GasIcon className="pt-0.5 m-auto fill-secondary" />
          </HoverTooltip>
        )}
      </div>
      <PortfolioAssetActionButton
        selectCallback={handleFromSelectionCallback}
        isDisabled={isDisabled || isTokenSelected}
        isSelected={isTokenSelected}
      />
    </div>
  )
}
