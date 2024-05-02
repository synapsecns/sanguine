import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'

export const AvailableBalance = ({
  balance,
  maxBalanceBridgeable,
  onMaxBalance,
  disabled,
  isGasEstimateLoading,
}: {
  balance?: string
  maxBalanceBridgeable?: string
  onMaxBalance?: () => void
  disabled: boolean
  isGasEstimateLoading: boolean
}) => {
  const labelClassName = joinClassNames({
    space: 'block',
    textColor: 'text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

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
      <label
        onClick={onMaxBalance}
        className={labelClassName}
        htmlFor="inputRow"
      >
        {maxBalanceBridgeable ?? balance ?? '0.0'}
        <span className="text-zinc-500 dark:text-zinc-400"> available</span>
      </label>
    )
  }
}
