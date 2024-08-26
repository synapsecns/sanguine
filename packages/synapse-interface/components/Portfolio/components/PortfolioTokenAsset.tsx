import React, { useCallback } from 'react'
import { zeroAddress } from 'viem'
import { isNumber } from 'lodash'
import Image from 'next/image'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId, setFromToken } from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { inputRef } from '../../StateManagedBridge/InputContainer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { PortfolioAssetActionButton } from './PortfolioAssetActionButton'
import { HoverTooltip } from '@/components/HoverTooltip'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { useGasEstimator } from '@/utils/hooks/useGasEstimator'
import GasIcon from '@/components/icons/GasIcon'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { formatAmount } from '@/utils/formatAmount'
import { useWalletState } from '@/slices/wallet/hooks'

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
  const { isWalletPending } = useWalletState()
  const { icon, symbol, decimals, addresses } = token
  const tokenAddress = addresses[portfolioChainId]
  const tokenDecimals = isNumber(decimals)
    ? decimals
    : decimals[portfolioChainId]

  const parsedBalance = getParsedBalance(balance, tokenDecimals)
  const formattedBalance = formatAmount(parsedBalance)

  const isDisabled = false
  const isPortfolioChainSelected = fromChainId === portfolioChainId
  const isTokenSelected = isPortfolioChainSelected && fromToken === token
  const isGasToken = tokenAddress === zeroAddress

  const { maxBridgeableGas } = useGasEstimator()

  const handleFromSelectionCallback = useCallback(async () => {
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
      <div className="flex items-center gap-2 py-2 pl-2 pr-4 rounded cursor-default">
        <Image
          loading="lazy"
          alt={`${symbol} img`}
          className="w-6 h-6 rounded-md"
          src={icon}
        />
        <HoverTooltip
          hoverContent={
            isPortfolioChainSelected && isGasToken && maxBridgeableGas ? (
              <div className="whitespace-nowrap">
                Available:{' '}
                {trimTrailingZeroesAfterDecimal(maxBridgeableGas.toFixed(8))}{' '}
                {symbol}
              </div>
            ) : (
              <div className="whitespace-nowrap">
                {parsedBalance} {symbol}
              </div>
            )
          }
        >
          <div>
            {formattedBalance} {symbol}
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
        isDisabled={isDisabled || isTokenSelected || isWalletPending}
        isSelected={isTokenSelected}
      />
    </div>
  )
}
