import { usePortfolioState } from '@/slices/portfolio/hooks'
import { isValidAddress } from '@/utils/isValidAddress'

export const useSearchInputState = () => {
  const { searchInput, searchedBalances } = usePortfolioState()

  const isSearchInputActive = searchInput.length > 0
  const isSearchInputAddress = isValidAddress(searchInput)
  const isMasqueradeActive = Object.keys(searchedBalances).length > 0

  return {
    isSearchInputActive,
    isSearchInputAddress,
    isMasqueradeActive,
  }
}
