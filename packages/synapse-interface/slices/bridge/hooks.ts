import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}

export const useBridgeDisplayState = (): RootState['bridgeDisplay'] => {
  return useAppSelector((state) => state.bridgeDisplay)
}
