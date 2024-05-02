import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'
import { HoverTooltip } from '../HoverTooltip'
import { Token } from '@/utils/types'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'

export const AvailableBalance = ({
  fromToken,
  balance,
  maxBalanceBridgeable,
  onMaxBalance,
  disabled,
  // estimatedGasLimit,
  isGasToken,
  isGasEstimateLoading,
}: {
  fromToken: Token
  balance?: string
  maxBalanceBridgeable?: number
  onMaxBalance?: () => void
  disabled: boolean
  // estimatedGasLimit: bigint
  isGasToken: boolean
  isGasEstimateLoading: boolean
}) => {
  const labelClassName = joinClassNames({
    space: 'block',
    textColor: 'text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  const tooltipContent = (
    <div className="flex flex-col space-y-2 whitespace-nowrap">
      <span>
        Available: {maxBalanceBridgeable?.toString()} {fromToken?.symbol}
      </span>
    </div>
  )

  if (disabled) {
    return null
  } else if (isGasEstimateLoading) {
    return (
      <label className={labelClassName} htmlFor="inputRow">
        <span className="animate-pulse text-zinc-500 dark:text-zinc-400">
          calculating gas...
        </span>
      </label>
    )
  } else {
    return (
      <HoverTooltip
        isActive={Boolean(isGasToken && maxBalanceBridgeable)}
        hoverContent={tooltipContent}
      >
        <label
          onClick={onMaxBalance}
          className={labelClassName}
          htmlFor="inputRow"
        >
          {maxBalanceBridgeable?.toFixed(4) ?? balance ?? '0.0'}
          <span className="text-zinc-500 dark:text-zinc-400"> available</span>
        </label>
      </HoverTooltip>
    )
  }
}
