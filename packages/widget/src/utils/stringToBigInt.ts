export const stringToBigInt = (rawVal: string, rawDecimals: number) => {
  if (typeof rawVal !== 'string' && !rawVal) {
    return 0n
  }

  if (typeof rawDecimals !== 'number') {
    return 0n
  }

  try {
    let value = rawVal.replace(/[$,]/g, '')
    if (['.', '0.', '', '.0'].includes(value)) {
      return 0n
    }
    if (value.startsWith('.')) {
      value = `0${value}`
    }

    // Scale the decimal number up by the appropriate number of decimals.
    const scaleFactor = BigInt(10 ** rawDecimals)

    // Split the input into whole and fractional parts.
    const [wholePart, fractionalPart = '0'] = value.split('.')

    // Convert the whole part directly to a BigInt.
    const wholeBigInt = BigInt(wholePart) * scaleFactor

    // For the fractional part, first scale it up to the right size, then trim any excess decimal places.
    const fractionalBigInt = BigInt(
      fractionalPart.padEnd(rawDecimals, '0').slice(0, rawDecimals)
    )

    return wholeBigInt + fractionalBigInt
  } catch (error) {
    console.log(error)
  }
}
