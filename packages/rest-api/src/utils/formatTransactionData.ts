import { PopulatedTx } from '@synapsecns/sdk-router'

import { Prettify } from './prettify'

/**
 * Transforms an object with a tx field into one with a callData field.
 * The tx field is removed and replaced with callData containing either the tx value or null
 * depending on the shouldIncludeCallData flag.
 *
 * @param object The source object containing tx field
 * @param shouldIncludeCallData Boolean flag determining if callData should be included
 * @returns Transformed object with tx field removed and callData field added
 */
export const formatTransactionData = <T extends { tx?: PopulatedTx }>(
  object: T,
  shouldIncludeCallData: boolean
): Prettify<Omit<T, 'tx'> & { callData: PopulatedTx | null }> => {
  const { tx, ...rest } = object

  return {
    ...rest,
    callData: shouldIncludeCallData ? tx : null,
  }
}
