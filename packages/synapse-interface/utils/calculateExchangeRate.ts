import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

export const calculateExchangeRate = (
  amountFrom: bigint,
  tokenPrecisionFrom: number,
  amountTo: bigint,
  tokenPrecisionTo: number
) => {
  if (amountFrom > 0n) {
    return (amountTo
      * (10n ** BigInt(36 - tokenPrecisionTo))) // convert to standard 1e18 precision
      / (amountFrom * (10n ** BigInt(18 - tokenPrecisionFrom)));
  } else {
    return 0n;
  }
}
