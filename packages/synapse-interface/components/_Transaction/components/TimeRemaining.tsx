import { useMemo } from 'react'
import { useTranslations } from 'next-intl'

import ExclamationIcon from '@components/icons/ExclamationIcon'
import { formatCompactDuration } from '@/utils/time'

export const TimeRemaining = ({
  isDelayed,
  remainingTime,
  delayedTime,
  status,
}: {
  isDelayed: boolean
  remainingTime: number
  delayedTime: number | null
  status: 'pending' | 'completed' | 'reverted' | 'refunded'
}) => {
  const t = useTranslations('Time')
  const compactDurationLabels = {
    minute: t('m'),
    second: t('s'),
  }

  const estTime = useMemo(() => {
    return `${formatCompactDuration(
      remainingTime,
      compactDurationLabels
    )} ${t('remaining')}`
  }, [compactDurationLabels, remainingTime, t])

  if (status === 'completed') {
    return <div className="text-sm text-green-400">{t('Complete')}!</div>
  }

  if (status === 'reverted') {
    return (
      <span className="flex items-center space-x-1 text-sm">
        <ExclamationIcon className="w-4 h-4" /> <span>{t('Reverted')}</span>
      </span>
    )
  }

  if (status === 'refunded') {
    return (
      <span className="flex items-center space-x-1 text-sm">
        <ExclamationIcon className="w-4 h-4" /> <span>{t('Refunded')}</span>
      </span>
    )
  }

  if (isDelayed) {
    const absoluteDelayedTime = Math.abs(delayedTime)
    const showDelayedTime = absoluteDelayedTime > 60
    return (
      <div className="text-sm">
        {t('Waiting')}...{' '}
        {showDelayedTime
          ? `(${formatCompactDuration(
              absoluteDelayedTime,
              compactDurationLabels
            )})`
          : null}
      </div>
    )
  }

  return <div className="text-sm">{estTime}</div>
}
