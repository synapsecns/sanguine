import { UserRejectedRequestError } from 'viem'

export const isTransactionUserRejectedError = (error: unknown): boolean => {
  if (error instanceof UserRejectedRequestError) {
    return true
  }

  // If the error type check is not possible or not specific enough, use properties or regex
  if (typeof error === 'object' && error !== null) {
    const message = (error as { message?: string }).message

    if (typeof message === 'string') {
      const regex = /User rejected the request/
      return regex.test(message)
    }
  }

  return false
}
