import {
  WaitForTransactionReceiptTimeoutError,
  TransactionNotFoundError,
} from 'viem'

export const isTransactionReceiptError = (error: unknown): boolean => {
  if (error instanceof WaitForTransactionReceiptTimeoutError) {
    return true
  }

  if (error instanceof TransactionNotFoundError) {
    return true
  }

  // If the error type check is not possible or not specific enough, use properties or regex
  if (typeof error === 'object' && error !== null) {
    const message = (error as { message?: string }).message

    if (typeof message === 'string') {
      const regex =
        /Timed out while waiting for transaction with hash "0x[0-9a-fA-F]+" | Transaction with hash "0x[0-9a-fA-F]+" could not be found\./
      return regex.test(message)
    }
  }

  return false
}
