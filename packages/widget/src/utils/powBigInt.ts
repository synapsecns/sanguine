export const powBigInt = (base, exponent) => {
  let result = 1n
  for (let i = 0; i < exponent; i++) {
    result *= base
  }
  return result
}
