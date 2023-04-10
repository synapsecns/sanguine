import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

// removes common copy/pasted financial characters
export const stringToBigNum = (rawVal, rawDecimals) => {
  const value = rawVal.replace(/[$,]/g, '')
  if (['.', '0.', ''].includes(value)) {
    return Zero
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
