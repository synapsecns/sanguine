import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

// removes common copy/pasted financial characters
export const stringToBigNum = (rawVal, decimals) => {
  const value = rawVal.replace(/[$,]/g, '')
  if (['.', '0.', ''].includes(value)) {
    return Zero
  }
  let valueSplit = value.split('.')
  let valueBase = valueSplit[0]
  let valueMantissa = valueSplit?.[1]?.length > 0 ? valueSplit[1] : '0'
  return BigNumber.from(valueBase)
    .mul(decimals)
    .add(
      BigNumber.from(valueMantissa)
        .mul(decimals)
        .div(BigNumber.from(10).pow(valueMantissa.length))
    )
}
