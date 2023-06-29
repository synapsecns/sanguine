import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

export const calculateExchangeRate = (
  amountFrom,
  tokenPrecisionFrom,
  amountTo,
  tokenPrecisionTo
) => {
  if (amountFrom > 0) {
    return amountTo
      .mul(BigNumber.from(10).pow(36 - tokenPrecisionTo)) // convert to standard 1e18 precision
      .div(BigNumber.from(amountFrom).mul(BigNumber.from(10).pow(18 - tokenPrecisionFrom)))
  } else {
    return Zero
  }
}
