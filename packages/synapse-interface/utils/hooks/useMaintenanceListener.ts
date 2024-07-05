import { useEffect } from 'react'

import { useIntervalTimer } from './useIntervalTimer'
import { getSynapsePauseData } from '@/components/Maintenance/functions/getSynapsePauseData'

export const useMaintenanceListener = () => {
  const interval = useIntervalTimer(60000)
  const fetchMaintenanceData = getSynapsePauseData()

  useEffect(() => {
    fetchMaintenanceData()
  }, [interval])

  return null
}
