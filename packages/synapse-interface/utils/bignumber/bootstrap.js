import { BigNumber, BigNumberish } from '@ethersproject/bignumber'

import { Zero } from '@ethersproject/constants'
import { parseUnits } from '@ethersproject/units'

/**
 * @param {BigNumberish} multiplier
 * @param {BigNumberish} divisor
 */
BigNumber.prototype.mulDiv = function (multiplier, divisor) {
  if (BigNumber.from(divisor).gt(0)) {
    return BigNumber.from(this).mul(multiplier).div(divisor)
  } else {
    return Zero
  }
}


/**
 * @param {BigNumberish} decimals
 */
String.prototype.toBigNumber = function (decimals) {
  try {
    return parseUnits(this, decimals)
  } catch (error) {
    console.debug(`Failed to parse input amount: "${this}"`, error)
  }
  return BigNumber.from(0)
}