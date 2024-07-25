import { useEffect } from 'react'

import { useIntervalTimer } from './useIntervalTimer'
import { useSynapsePauseData } from '@/components/Maintenance/hooks/useSynapsePauseData'

export const useMaintenanceListener = () => {
  const interval = useIntervalTimer(60000)
  const fetchMaintenanceData = useSynapsePauseData()

  useEffect(() => {
    fetchMaintenanceData()
  }, [interval])

  return null
}
