import { Zero } from '@ethersproject/constants'
import { parseUnits } from '@ethersproject/units'




/**
 * @param {string} value
 * @param {number|BigNumberish} unitName
 */
export function smartParseUnits(value, unitName) {
  if (value) {
    if (value === '') {
      return Zero
    } else {
      return parseUnits(value.replace(/,/g, ''), unitName)
    }
  } else {
    return Zero
  }
}

