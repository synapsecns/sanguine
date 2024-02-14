// Needs to be updated if Wagmi update changes error string
// Currently @wagmi/core@v1.4.12
export const isTransactionReceiptError = (str: string): boolean => {
  return str.includes('Timed out while waiting for transaction')
}
