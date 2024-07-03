import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useMaintanceState = (): RootState['maintenance'] => {
  return useAppSelector((state) => state.maintenance)
}
