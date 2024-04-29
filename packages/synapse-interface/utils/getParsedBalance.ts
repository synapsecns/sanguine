import { formatBigIntToString } from './bigint/format'
import { hasOnlyZeroes } from './hasOnlyZeroes'

export const getParsedBalance = (
  balance: bigint,
  decimals: number,
  places?: number
) => {
  const formattedBalance = formatBigIntToString(balance, decimals, places)
  const verySmallBalance = balance > 0n && hasOnlyZeroes(formattedBalance)

  return verySmallBalance ? '< 0.001' : formattedBalance
}
