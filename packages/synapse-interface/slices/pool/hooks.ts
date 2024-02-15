
import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'




export const usePoolDataState = (): RootState['poolData'] => {
  return useAppSelector((state: RootState) => state.poolData)
}

export const usePoolUserDataState = (): RootState['poolUserData'] => {
  return useAppSelector((state: RootState) => state.poolUserData)
}

export const usePoolWithdrawState = (): RootState['poolWithdraw'] => {
  return useAppSelector((state: RootState) => state.poolWithdraw)
}

export const usePoolDepositState = (): RootState['poolDeposit'] => {
  return useAppSelector((state: RootState) => state.poolDeposit)
}