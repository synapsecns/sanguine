import React from 'react'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { joinClassNames } from '@/utils/joinClassNames'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { HoverTooltip } from '../HoverTooltip'

export const AvailableBalance = ({
  fromChainId,
  fromToken,
  balance,
  parsedBalance,
  onMaxBalance,
  disabled,
  isGasEstimateLoading,
}: {
  fromChainId: number | null
  fromToken: Token | null
  balance?: bigint
  parsedBalance?: string
  onMaxBalance?: () => void
  disabled: boolean
  isGasEstimateLoading: boolean
}) => {
  const tooltipContent = (
    <div className="flex flex-col space-y-2 whitespace-nowrap">
      <span>
        {parsedBalance} {fromToken?.symbol}
      </span>
    </div>
  )

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
        <span className="text-zinc-500 dark:text-zinc-400">
          calculating gas...
        </span>
      </label>
    )
  } else {
    return (
      <HoverTooltip isActive={true} hoverContent={tooltipContent}>
        <label
          onClick={onMaxBalance}
          className={labelClassName}
          htmlFor="inputRow"
        >
          {parsedBalance ?? '0.0'}
          <span className="text-zinc-500 dark:text-zinc-400"> available</span>
        </label>
      </HoverTooltip>
    )
  }
}
