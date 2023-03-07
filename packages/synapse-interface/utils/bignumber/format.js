import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits, commify } from '@ethersproject/units'

import { POOL_FEE_PRECISION } from '@constants/fees'

/**
 * @param {BigNumber} bn
 * @param {number} nativePrecison
 * @param {number} decimalPlaces
 */
export function formatBNToString(bn, nativePrecison, decimalPlaces) {
  const fullPrecision = formatUnits(bn, nativePrecison)
  const decimalIdx = fullPrecision.indexOf('.')

  if (decimalPlaces === undefined || decimalIdx === -1) {

    return fullPrecision
  } else {
    // don't include decimal point if places = 0
    const rawNumber = Number(fullPrecision)

    if (rawNumber == 0) {
      return rawNumber.toFixed(1)
    } else {
      return rawNumber.toFixed(decimalPlaces) //.slice(0, num)
    }

  }
}

/**
 * @param {BigNumber} bn
 * @param {TokenInfo} tokenInfo
 * @param {number} decimalPlaces
 */
export function formatBnMagic(bn, tokenInfo, decimalPlaces) {
  const nativePrecison = tokenInfo.decimals
  return formatBNToString(bn, nativePrecison, decimalPlaces)
}

/**
 * @param {BigNumber} bn
 * @param {number} nativePrecison
 * @param {number} decimalPlaces
 */
export function formatBNToPercentString(bn, nativePrecison, decimalPlaces=2) {
  return `${formatBNToString(bn, nativePrecison - 2, decimalPlaces)}%`
}

/**
 * @param {BigNumber} bn
 * @param {TokenInfo} tokenInfo
 * @param {number} decimalPlaces
 */
export function formatCommifyBn(bn, tokenInfo, decimalPlaces) {
  const nativePrecison = tokenInfo.decimals
  return commify(
    formatBNToString(bn, nativePrecison, decimalPlaces)
  )
}

/**
 * @param {BigNumber} bn
 * @param {number} decimals
 */
export function commifyBnToString(bn, decimals = 2) {
  return (
    commify(formatBNToString(bn, 18, decimals))
  )
}

/**
 * @param {BigNumber?} bn
 * @param {number} decimals
 */
export function commifyBnWithDefault(bn, decimals) {
  return (bn ? commifyBnToString(bn, decimals) : '0')
}

/**
 * @param {BigNumber?} value
 */
export function bnPercentFormat(bn) {
  return (
    bn ? formatBNToPercentString(bn, POOL_FEE_PRECISION) : null
  )
}

