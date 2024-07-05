import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useMaintenanceState = (): RootState['maintenance'] => {
  return useAppSelector((state) => state.maintenance)
}
