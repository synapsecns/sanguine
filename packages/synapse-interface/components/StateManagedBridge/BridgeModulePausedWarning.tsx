import { useMemo } from 'react'
import { useTranslations } from 'next-intl'

import type { BridgeModulePause } from '@/components/Maintenance/Maintenance'
import { WarningMessage } from '@/components/Warning'
import { getPausedBridgeModuleNamesForRoute } from '@/utils/getPausedBridgeModuleNamesForRoute'

export const BridgeModulePausedWarning = ({
  fromChainId,
  toChainId,
  pausedModulesList,
}: {
  fromChainId: number
  toChainId: number
  pausedModulesList: BridgeModulePause[]
}) => {
  const t = useTranslations('Bridge')
  const pausedModuleNames = useMemo(
    () =>
      getPausedBridgeModuleNamesForRoute({
        pausedModules: pausedModulesList,
        fromChainId,
        toChainId,
      }),
    [fromChainId, toChainId, pausedModulesList]
  )

  if (!fromChainId || !toChainId || pausedModuleNames.length === 0) {
    return null
  }

  return (
    <WarningMessage
      twClassName="mb-2 border border-amber-500/30 !bg-amber-50 !text-amber-950 dark:!bg-amber-950/30 dark:!text-amber-100"
      message={t('BridgeModulePauseWarningMessage')}
    />
  )
}
