import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'

export const AvailableBalance = ({
  balance,
  maxBridgeableBalance,
  isGasToken,
  isGasEstimateLoading,
  isDisabled,
}: {
  balance?: string
  maxBridgeableBalance?: number
  gasCost?: string
  isGasToken: boolean
  isGasEstimateLoading: boolean
  isDisabled: boolean
}) => {
  const labelClassName = joinClassNames({
    space: 'block',
    text: 'text-xxs md:text-xs',
    cursor: 'cursor-default',
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
      <label className={labelClassName} htmlFor="inputRow">
        <span className="text-zinc-500 dark:text-zinc-400">Available: </span>
        {maxBridgeableBalance?.toFixed(4) ?? balance ?? '0.0'}
      </label>
    )
  }
}
