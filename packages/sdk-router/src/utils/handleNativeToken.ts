import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { PopulatedTransaction } from '@ethersproject/contracts'

export const ETH_NATIVE_TOKEN_ADDRESS =
  '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

export const handleNativeToken = (tokenAddr: string) => {
  return tokenAddr === '' || tokenAddr === AddressZero
    ? ETH_NATIVE_TOKEN_ADDRESS
    : tokenAddr
}

export const isNativeToken = (tokenAddr: string): boolean => {
  return tokenAddr.toLowerCase() === ETH_NATIVE_TOKEN_ADDRESS.toLowerCase()
}

/**
 * Sets the tx.value to the amount if the token is native, otherwise sets it to 0.
 *
 * @param tx - The transaction to adjust.
 * @param tokenAddr - The address of the token to check for being native.
 * @param amountNative - The amount to set if the token is native.
 * @param amountOther - The amount to set if the token is not native (optional, defaults to 0).
 * @returns The adjusted populated transaction.
 */
export const adjustValueIfNative = (
  tx: PopulatedTransaction,
  tokenAddr: string,
  amountNative: BigNumber,
  amountOther: BigNumber = Zero
): PopulatedTransaction => {
  tx.value = isNativeToken(tokenAddr) ? amountNative : amountOther
  return tx
}
