import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

/**
 * @param {BigNumber} tokenInputAmount assuing 18d precision
 * @param {BigNumber} tokenOutputAmount assuming 18d precision
 * @param {BigNumber} virtualPrice cause everything is fake anyway
 */

const BN_1E18 = BigNumber.from(10).pow(18)

export const calculatePriceImpact = (
  tokenInputAmount: BigNumber, // assumed to be 18d precision
  tokenOutputAmount: BigNumber,
  virtualPrice = BN_1E18,
  isWithdraw = false
): BigNumber => {
  if (tokenInputAmount.lte(0)) {
    return Zero
  }

  return isWithdraw
    ? tokenOutputAmount
        .mul(BigNumber.from(10).pow(36))
        .div(tokenInputAmount.mul(virtualPrice))
        .sub(BN_1E18)
    : virtualPrice.mul(tokenOutputAmount).div(tokenInputAmount).sub(BN_1E18)
}

export const calculatePriceImpactWithdraw = (
  lpTokenInputAmount,
  tokenOutputAmount,
  virtualPrice = BigNumber.from(10).pow(18)
) => {
  const baseSquared = BigNumber.from(10).pow(36)
  if (lpTokenInputAmount.gt(0)) {
    return tokenOutputAmount
      .mul(baseSquared)
      .div(lpTokenInputAmount.mul(virtualPrice))
      .sub(BigNumber.from(10).pow(18))
  } else {
    return Zero
  }
}
