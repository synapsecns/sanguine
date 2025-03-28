import { PopulatedTransaction } from 'ethers'

import { PopulatedTx } from '../types'

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
