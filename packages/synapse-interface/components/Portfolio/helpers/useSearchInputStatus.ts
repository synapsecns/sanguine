import type { Address } from 'viem'

import { usePortfolioState } from '@/slices/portfolio/hooks'
import { isValidAddress } from '@/utils/isValidAddress'

export const useSearchInputState = () => {
  const { searchInput, searchedBalances } = usePortfolioState()

  const isSearchInputActive = searchInput.length > 0
  const isSearchInputAddress = isValidAddress(searchInput)
  const isMasqueradeActive = Object.keys(searchedBalances).length > 0
  const masqueradeAddress = Object.keys(searchedBalances)[0] as Address

  return {
    isSearchInputActive,
    isSearchInputAddress,
    isMasqueradeActive,
    masqueradeAddress,
  }
}
