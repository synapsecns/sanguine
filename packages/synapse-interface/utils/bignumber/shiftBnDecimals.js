import { BigNumber } from '@ethersproject/bignumber'
/**
 * @param {BigNumber} bn
 * @param {number} shiftAmount
 */
export function shiftBnDecimals(bn, shiftAmount) {
  if (shiftAmount < 0) throw new Error('shiftAmount must be positive')
  return bn.mul(BigNumber.from(10).pow(shiftAmount))
}
