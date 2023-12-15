import { RootState } from '@/state/store'
import { useAppSelector } from '@/state/hooks'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}
