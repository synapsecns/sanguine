import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const usePriceDataState = (): RootState['priceData'] => {
  return useAppSelector((state) => state.priceData)
}
