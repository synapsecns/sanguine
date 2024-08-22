import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useWalletState = (): RootState['wallet'] => {
  return useAppSelector((state) => state.wallet)
}
