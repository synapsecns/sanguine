import { createContext } from 'react'
import { getSynapsePauseData } from '@/components/Maintenance/functions/getSynapsePauseData'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

const MaintenanceContext = createContext(null)

export const MaintenanceProvider = ({ children }) => {
  const time = useIntervalTimer(60000)
  const fetchMaintenanceData = getSynapsePauseData()

  fetchMaintenanceData()

  return (
    <MaintenanceContext.Provider value={null}>
      {children}
    </MaintenanceContext.Provider>
  )
}
