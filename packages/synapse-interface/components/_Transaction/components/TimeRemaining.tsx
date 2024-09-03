import { useMemo } from 'react'
import { useTranslations } from 'next-intl'

import ExclamationIcon from '@components/icons/ExclamationIcon'

export const TimeRemaining = ({
  isDelayed,
  remainingTime,
  delayedTime,
  status,
}: {
  isDelayed: boolean
  remainingTime: number
  delayedTime: number | null
  status: 'pending' | 'completed' | 'reverted'
}) => {
  const t = useTranslations('Time')

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

  if (isDelayed) {
    const delayedTimeInMin = Math.floor(delayedTime / 60)
    const absoluteDelayedTime = Math.abs(delayedTimeInMin)
    const showDelayedTime = delayedTimeInMin < -1
    return (
      <div className="text-sm">
        {t('Waiting')}... {showDelayedTime ? `(${absoluteDelayedTime}m)` : null}
      </div>
    )
  }

  const estTime = useMemo(() => {
    if (remainingTime > 60) {
      return Math.ceil(remainingTime / 60) + 'm remaining'
    } else {
      return remainingTime + 's remaining'
    }
  }, [remainingTime])

  return <div className="text-sm">{estTime}</div>
}
