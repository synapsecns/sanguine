import { useState } from 'react'

/**
 * Wrapper for telling when a claim tx is pending / goes through
 * @returns isPending=`true` while claiming happens, isPending=`false` otherwise
 */
export function usePendingTxWrapper() {
  const [isPending, setIsPending] = useState(false)

  async function pendingTxWrapFunc(claimPromise) {
    setIsPending(true)
    let tx
    try {
      tx = await claimPromise
    } finally {
      setIsPending(false)
    }
    return tx
  }
  return [isPending, pendingTxWrapFunc]
}