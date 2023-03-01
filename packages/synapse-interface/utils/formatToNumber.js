import { formatUnits } from '@ethersproject/units'



/**
 * @param {BigNumber} bn
 */
export function formatToNumber(bn) {
  return Number(formatUnits(bn))
}