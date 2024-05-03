import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'
import { HoverTooltip } from '../HoverTooltip'
import { Token } from '@/utils/types'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'

export const AvailableBalance = ({
  balance,
  maxBridgeableBalance,
  isGasToken,
  isGasEstimateLoading,
  isDisabled,
  onMaxBalance,
}: {
  balance?: string
  maxBridgeableBalance?: number
  gasCost?: string
  isGasToken: boolean
  isGasEstimateLoading: boolean
  isDisabled: boolean
  onMaxBalance?: () => void
}) => {
  const labelClassName = joinClassNames({
    space: 'block',
    textColor: 'text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  if (isDisabled) {
    return null
  } else if (isGasToken && isGasEstimateLoading) {
    return (
      <label className={labelClassName} htmlFor="inputRow">
        <span className="animate-pulse text-zinc-500 dark:text-zinc-400">
          calculating gas...
        </span>
      </label>
    )
  } else {
    return (
      <label
        onClick={onMaxBalance}
        className={labelClassName}
        htmlFor="inputRow"
      >
        {maxBridgeableBalance?.toFixed(4) ?? balance ?? '0.0'}
        <span className="text-zinc-500 dark:text-zinc-400"> available</span>
      </label>
    )
  }
}
