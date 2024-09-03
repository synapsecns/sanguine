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
}: {
  balance?: string
  maxBridgeableBalance?: number
  gasCost?: string
  isGasToken: boolean
  isGasEstimateLoading: boolean
  isDisabled: boolean
}) => {
  const labelClassNames = {
    space: 'block',
    text: 'text-xxs md:text-xs',
    cursor: 'cursor-default',
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
      <label className={joinClassNames(labelClassNames)} htmlFor="inputRow">
        <span className="text-zinc-500 dark:text-zinc-400">
          {t('Available')}:{' '}
        </span>
        {formatAmount(maxBridgeableBalance?.toString()) ?? balance ?? '0.0'}
      </label>
    )
  }
}
