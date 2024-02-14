import { WaitForTransactionReceiptTimeoutError } from 'viem'

// Needs to be updated if Wagmi update changes error string
// Currently @wagmi/core@v1.4.12
export const isTransactionReceiptError = (str: string): boolean => {
  return str.includes('Timed out while waiting for transaction')
}

export const _isTransactionReceiptError = (error: unknown): boolean => {
  if (error instanceof WaitForTransactionReceiptTimeoutError) {
    console.log('hit 1')
    return true
  }

  // If the error type check is not possible or not specific enough, use properties or regex
  if (typeof error === 'object' && error !== null) {
    console.log('hit 2')
    const message = (error as { message?: string }).message

    if (typeof message === 'string') {
      console.log('hit 3')
      const regex =
        /Timed out while waiting for transaction with hash "0x[0-9a-fA-F]+"/
      return regex.test(message)
    }
  }

  return false
}
