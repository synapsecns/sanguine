import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { formatBigIntToString } from './bigint/format'

export const getParsedBalance = (
  balance: bigint,
  decimals: number,
  places?: number
) => {
  return trimTrailingZeroesAfterDecimal(
    formatBigIntToString(balance, decimals, places)
  )
}
