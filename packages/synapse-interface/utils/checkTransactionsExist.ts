export const checkTransactionsExist = (
  transactions: any[] | undefined | null
): boolean => {
  const exists: boolean =
    transactions && Array.isArray(transactions) && transactions.length > 0
  return exists
}
