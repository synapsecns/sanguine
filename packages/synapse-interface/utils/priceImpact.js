import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

/**
 * @param {BigNumber} tokenInputAmount assuing 18d precision
 * @param {BigNumber} tokenOutputAmount assuming 18d precision
 * @param {BigNumber} virtualPrice cause everything is fake anyway
 */
export function calculatePriceImpact(
  tokenInputAmount,
  tokenOutputAmount,
  virtualPrice = BigNumber.from(10).pow(18),
) {
  if (tokenInputAmount.gt(0)) {
    return (
      virtualPrice
        .mul(tokenOutputAmount)
        .div(tokenInputAmount)
        .sub(BigNumber.from(10).pow(18))
    )
  } else {
    return Zero
  }
}

export function calculatePriceImpactWithdraw(
  lpTokenInputAmount,
  tokenOutputAmount,
  virtualPrice = BigNumber.from(10).pow(18),
) {
  const baseSquared = BigNumber.from(10).pow(36)
  if (lpTokenInputAmount.gt(0)) {
    return (
      tokenOutputAmount
        .mul(baseSquared)
        .div(lpTokenInputAmount.mul(virtualPrice))
        .sub(BigNumber.from(10).pow(18))
    )
  } else {
    return Zero
  }
}