import { useState } from 'react'

type usePendingTxWrapperReturnType = [
  isPending: boolean,
  pendingTxWrapFunc: (claimPromise: Promise<any>) => Promise<any>
]

/**
 * Wrapper for telling when a claim tx is pending / goes through
 * @returns isPending=`true` while claiming happens, isPending=`false` otherwise
 */
export function usePendingTxWrapper(): usePendingTxWrapperReturnType {
  const [isPending, setIsPending] = useState<boolean>(false)

  async function pendingTxWrapFunc(claimPromise: Promise<any>): Promise<any> {
    setIsPending(true)
    let tx
    try {
      tx = await claimPromise
    } finally {
      setIsPending(false)
    }
    return tx
  }
  // return { isPending, pendingTxWrapFunc }
  return [isPending, pendingTxWrapFunc]
}
