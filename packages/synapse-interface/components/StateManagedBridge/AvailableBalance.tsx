import React from 'react'
import { useTranslations } from 'next-intl'

import { joinClassNames } from '@/utils/joinClassNames'
import { formatAmount } from '@/utils/formatAmount'

export const AvailableBalance = ({
  balance,
  maxBridgeableBalance,
  isGasToken,
  isGasEstimateLoading,
  isDisabled,
  onClick,
}: {
  balance?: string
  maxBridgeableBalance?: number
  gasCost?: string
  isGasToken: boolean
  isGasEstimateLoading: boolean
  isDisabled: boolean
  onClick?: () => void
}) => {
  const labelClassNames = {
    space: 'block',
    text: 'text-xxs md:text-xs',
    cursor: onClick ? 'cursor-pointer' : 'cursor-default',
    hover: onClick ? 'hover:opacity-70' : '',
    animation: onClick ? 'transition-all duration-150' : '',
  }

  const t = useTranslations('Bridge')

  if (isDisabled) {
    return null
  } else if (isGasToken && isGasEstimateLoading) {
    return (
      <label className={joinClassNames(labelClassNames)} htmlFor="inputRow">
        <span className="animate-pulse text-zinc-500 dark:text-zinc-400">
          {t('calculating gas')}...
        </span>
      </label>
    )
  } else {
    return (
      <label
        className={joinClassNames(labelClassNames)}
        htmlFor="inputRow"
        onClick={onClick}
      >
        <span className="text-zinc-500 dark:text-zinc-400">
          {t('Balance')}:{' '}
        </span>
        <span className={onClick ? 'text-fuchsia-400' : ''}>
          {maxBridgeableBalance != null
            ? formatAmount(maxBridgeableBalance.toString())
            : balance ?? '0.0'}
        </span>
      </label>
    )
  }
}
