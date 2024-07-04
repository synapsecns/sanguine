import { createContext } from 'react'
import { getSynapsePauseData } from '@/components/Maintenance/functions/getSynapsePauseData'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

const MaintenanceContext = createContext(null)

export const MaintenanceProvider = ({ children }) => {
  const fetchMaintenanceData = getSynapsePauseData()
  fetchMaintenanceData()
  useIntervalTimer(60000)

  return (
    <MaintenanceContext.Provider value={null}>
      {children}
    </MaintenanceContext.Provider>
  )
}
