import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

// removes common copy/pasted financial characters
export const stringToBigNum = (rawVal, rawDecimals) => {
  let value = rawVal.replace(/[$,]/g, '')
  if (['.', '0.', '', '.0'].includes(value)) {
    return Zero
  }
  if (value.startsWith('.')) {
    value = `0${value}`
  }
  const decimals = BigNumber.from(10).pow(rawDecimals)
  const valueSplit = value.split('.')
  const valueBase = valueSplit[0]
  const valueMantissa = valueSplit?.[1]?.length > 0 ? valueSplit[1] : '0'
  return BigNumber.from(valueBase)
    .mul(decimals)
    .add(
      BigNumber.from(valueMantissa)
        .mul(decimals)
        .div(BigNumber.from(10).pow(valueMantissa.length))
    )
}

export const stringToBigInt = (rawVal, rawDecimals) => {
  let value = rawVal.replace(/[$,]/g, '')
  if (['.', '0.', '', '.0'].includes(value)) {
    return BigInt(0)
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
}
