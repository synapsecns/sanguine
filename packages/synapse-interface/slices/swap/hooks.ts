import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useSwapState = (): RootState['swap'] => {
  return useAppSelector((state) => state.swap)
}
