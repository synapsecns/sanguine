import { Zero } from '@ethersproject/constants'
import { BigNumber, PopulatedTransaction } from 'ethers'

import { PopulatedTx } from '../types'
import { isNativeToken } from './addressUtils'

export const stringifyPopulatedTransaction = (
  tx?: PopulatedTransaction
): PopulatedTx | undefined => {
  if (!tx || !tx.to) {
    return undefined
  }
  return {
    to: tx.to,
    data: tx.data ?? '',
    value: tx.value?.toString() ?? '0',
  }
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
