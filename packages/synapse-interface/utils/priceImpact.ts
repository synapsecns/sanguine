const powBigInt = (base, exponent) => {
  let result = 1n
  for (let i = 0; i < exponent; i++) {
    result *= base
  }
  return result
}

const BI_1E18 = powBigInt(10n, 18)

export const calculatePriceImpact = (
  tokenInputAmount: bigint, // assumed to be 18d precision
  tokenOutputAmount: bigint,
  virtualPrice: bigint = BI_1E18,
  isWithdraw: boolean = false
) => {
  if (tokenInputAmount <= 0n) {
    return 0n
  }

  return isWithdraw
    ? (tokenOutputAmount * powBigInt(10n, 36)) /
        (tokenInputAmount * virtualPrice) -
        BI_1E18
    : (virtualPrice * tokenOutputAmount) / tokenInputAmount - BI_1E18
}

export const calculatePriceImpactWithdraw = (
  lpTokenInputAmount,
  tokenOutputAmount,
  virtualPrice = powBigInt(10n, 18)
) => {
  const baseSquared = powBigInt(10n, 36)
  if (lpTokenInputAmount > 0n) {
    return (
      (tokenOutputAmount * baseSquared) / (lpTokenInputAmount * virtualPrice) -
      powBigInt(10n, 18)
    )
  } else {
    return 0n
  }
}
