import { createContext, useContext, useEffect, useRef, useState } from 'react'
import { getSynapsePauseData } from '@/components/Maintenance/functions/useSynapsePauseData'
import { useMaintanceState } from '@/slices/maintenance/hooks'
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
