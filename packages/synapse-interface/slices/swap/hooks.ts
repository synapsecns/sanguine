import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useSwapDisplayState = (): RootState['swapDisplay'] => {
  return useAppSelector((state) => state.swapDisplay)
}
export const useSwapState = (): RootState['swap'] => {
  return useAppSelector((state) => state.swap)
}
