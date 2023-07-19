import { useDispatch, useSelector } from 'react-redux'
import { getAccount } from '@wagmi/core'

import { AppDispatch, RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import { fetchAndStorePortfolioBalances } from './reducer'
import { NetworkTokenBalancesAndAllowances } from '@/utils/hooks/usePortfolioBalances'

export const usePortfolioBalances = (): NetworkTokenBalancesAndAllowances => {
  return useAppSelector((state) => state.portfolio.balancesAndAllowances)
}

export const useFetchPortfolioBalances = () => {
  const dispatch: AppDispatch = useDispatch()
  const { address } = getAccount()
  const { balancesAndAllowances, status, error } = useSelector(
    (state: RootState) => state.portfolio
  )

  const fetch = () => {
    if (address) {
      dispatch(fetchAndStorePortfolioBalances(address))
    }
  }

  return { balancesAndAllowances, fetchPortfolioBalances: fetch, status, error }
}
