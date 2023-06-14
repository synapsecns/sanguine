import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

export const calculateExchangeRate = (
  amountFrom,
  tokenPrecisionFrom,
  amountTo,
  tokenPrecisionTo
) => {
  if (amountFrom.gt('0')) {
    const increasedPrecisionAmountTo = amountTo.mul(
      BigNumber.from(10).pow(36 - tokenPrecisionTo)
    )

    const increasedPrecisionAmountFrom = amountFrom.mul(
      BigNumber.from(10).pow(18 - tokenPrecisionFrom)
    )

    const ratio = increasedPrecisionAmountTo.div(increasedPrecisionAmountFrom)
    return ratio
  } else {
    return Zero
  }
}
