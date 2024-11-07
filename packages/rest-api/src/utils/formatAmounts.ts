import { parseUnits } from '@ethersproject/units'
import { BigNumber } from '@ethersproject/bignumber'

export const formatAndValidateAmount = (
  amount: string,
  decimals: number
): BigNumber => {
  // Clean the input first
  const cleanedAmount = amount.replace(/[^0-9.]/g, '')

  // Handle empty or invalid input
  if (!cleanedAmount || cleanedAmount === '.') {
    return BigNumber.from(0)
  }

  // Split into whole and decimal parts
  const [wholePart, fractionalPart = ''] = cleanedAmount.split('.')

  // Truncate decimal places to token's decimals
  const truncatedAmount = `${wholePart}${
    fractionalPart ? '.' + fractionalPart.slice(0, decimals) : ''
  }`

  return parseUnits(truncatedAmount, decimals)
}
