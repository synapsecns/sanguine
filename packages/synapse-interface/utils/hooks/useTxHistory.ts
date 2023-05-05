import { useState } from 'react'
import _ from 'lodash'

export const useTxHistory = () => {
  // const { account } = useActiveWeb3React()
  // const [transactionsByAccount, setTransactionsByAccount] = usePersistedTransactionHistory({})

  // const transactions = transactionsByAccount[account] ?? []

  // function setTransactions(txns) {
  //   setTransactionsByAccount({
  //     ...transactionsByAccount,
  //     [account]: txns
  //   })
  // }
  const [transactions, setTransactions] = useState([])

  const addTransaction = ({
    transactionHash,
    hash,
    chainId,
    ...transaction
  }) => {
    const formattedTx = {
      transactionHash: transactionHash ?? hash,
      chainId,
      ...transaction,
    }

    const filteredTransactions = transactions.filter(
      (tx) => tx.transactionHash !== formattedTx.transactionHash
    )
    const arr = [...filteredTransactions, formattedTx]

    setTransactions(arr)
  }

  const updateTransactions = (txns) => {
    setTransactions((oldTransactions) => {
      const newTxnHashes = txns.map((tx) => tx.transactionHash ?? tx.hash)

      const filteredPrevTxns = oldTransactions.filter((tx) => {
        const txOverlap = newTxnHashes.includes(tx.transactionHash ?? tx.hash)
        return !txOverlap
      })

      const txnsToAdd = _.sortBy(
        [
          ...filteredPrevTxns,
          ...txns.map((tx) => {
            return { ...tx, transactionHash: tx.transactionHash ?? tx.hash }
          }),
        ],
        (tx) => -tx.timestamp
      )
      return txnsToAdd
    })
  }

  const clear = () => {
    setTransactions([])
  }

  return {
    transactions,
    setTransactions,
    addTransaction,
    updateTransactions,
    clear,
  }
}
