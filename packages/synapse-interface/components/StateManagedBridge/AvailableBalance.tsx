import React from 'react'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { joinClassNames } from '@/utils/joinClassNames'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { HoverTooltip } from '../HoverTooltip'

export const AvailableBalance = ({
  fromChainId,
  fromValue,
  fromToken,
  balance,
  parsedBalance,
  maxBridgeableBalance,
  isGasToken = false,
  parsedGasCost,
  onMaxBalance,
  hasMounted,
  isConnected,
  disabled = false,
}: {
  fromChainId: number | null
  fromValue: string
  fromToken: Token | null
  balance?: bigint
  parsedBalance?: string
  maxBridgeableBalance?: string
  isGasToken?: boolean
  parsedGasCost?: string
  onMaxBalance?: () => void
  hasMounted: boolean
  isConnected: boolean
  disabled?: boolean
}) => {
  const parsedBalanceFull = formatBigIntToString(
    balance,
    fromToken?.decimals[fromChainId]
  )

  const isTraceBalance = () => {
    return Boolean(
      balance &&
        hasOnlyZeroes(parsedBalance) &&
        !hasOnlyZeroes(parsedBalanceFull)
    )
  }

  const tooltipContent = (
    <div className="flex flex-col space-y-2 whitespace-nowrap">
      <span>
        {parsedBalanceFull} {fromToken?.symbol}
      </span>
    </div>
  )

  const labelClassName = joinClassNames({
    space: 'block',
    textColor: 'text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  return (
    <HoverTooltip isActive={true} hoverContent={tooltipContent}>
      <label
        onClick={onMaxBalance}
        className={labelClassName}
        htmlFor="inputRow"
      >
        {isTraceBalance() ? '<0.001' : parsedBalance ?? '0.0'}
        <span className="text-zinc-500 dark:text-zinc-400"> available</span>
      </label>
    </HoverTooltip>
  )
}
