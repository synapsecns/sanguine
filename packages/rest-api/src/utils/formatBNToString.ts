import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'

export const formatBNToString = (
  bn: BigNumber,
  nativePrecision: number,
  decimalPlaces = 18
) => {
  const fullPrecision = formatUnits(bn, nativePrecision)
  const decimalIdx = fullPrecision.indexOf('.')

  if (decimalPlaces === undefined || decimalIdx === -1) {
    return fullPrecision
  } else {
    const rawNumber = Number(fullPrecision)

    if (rawNumber === 0) {
      return rawNumber.toFixed(1)
    }
    return rawNumber.toString()
  }
}
