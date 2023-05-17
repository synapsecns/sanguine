import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits, commify } from '@ethersproject/units'
import { POOL_FEE_PRECISION } from '@constants/fees'

export const formatBNToString = (
  bn: BigNumber,
  nativePrecison: number,
  decimalPlaces?: number
) => {
  const fullPrecision = formatUnits(bn, nativePrecison)
  const decimalIdx = fullPrecision.indexOf('.')

  if (decimalPlaces === undefined || decimalIdx === -1) {
    return fullPrecision
  } else {
    // don't include decimal point if places = 0
    const rawNumber = Number(fullPrecision)

    if (rawNumber === 0) {
      return rawNumber.toFixed(1)
    } else {
      return rawNumber.toFixed(decimalPlaces) //.slice(0, num)
    }
  }
}

export const formatBNToPercentString = (
  bn: BigNumber,
  nativePrecison: number,
  decimalPlaces = 2
) => {
  return `${formatBNToString(bn, nativePrecison - 2, decimalPlaces)}%`
}

export const commifyBnToString = (bn: BigNumber, decimals = 2) => {
  return commify(formatBNToString(bn, 18, decimals))
}

export const commifyBnWithDefault = (bn: BigNumber, decimals: number) => {
  return bn ? commifyBnToString(bn, decimals) : '0'
}

export const bnPercentFormat = (bn: BigNumber) => {
  return bn ? formatBNToPercentString(bn, POOL_FEE_PRECISION) : null
}

export function fixNumberToPercentageString(num, numDecimals = 2) {
  return `${num?.toFixed(numDecimals)}%`
}
