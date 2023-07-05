const powBigInt = (base, exponent) => {
  let result = 1n
  for (let i = 0; i < exponent; i++) {
    result *= base
  }
  return result
}

export const calculateExchangeRate = (
  amountFrom: bigint,
  tokenPrecisionFrom: number,
  amountTo: bigint,
  tokenPrecisionTo: number
) => {
  if (amountFrom > 0n) {
    return (
      (amountTo * powBigInt(10n, BigInt(36 - tokenPrecisionTo))) / // convert to standard 1e18 precision
      (amountFrom * powBigInt(10n, BigInt(18 - tokenPrecisionFrom)))
    )
  } else {
    return 0n
  }
}
