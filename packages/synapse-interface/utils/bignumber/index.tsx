import { Zero } from '@ethersproject/constants'
import { parseUnits } from '@ethersproject/units'
import { BigNumberish } from '@ethersproject/bignumber'

export function smartParseUnits(
  value: string,
  unitName: number | BigNumberish
) {
  if (value) {
    if (value === '') {
      return Zero
    } else {
      return parseUnits(value.replace(/,/g, ''), unitName)
    }
  } else {
    return Zero
  }
}
