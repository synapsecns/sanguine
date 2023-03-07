import { useState } from 'react'

import _ from 'lodash'
import createPersistedState from 'use-persisted-state'


const usePersistedTransactionHistory = createPersistedState("historical_transactions")


export function useTxHistory() {
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

  function addTransaction({ transactionHash, hash, chainId, ...transaction }) {
    let formattedTx
    formattedTx = {
      transactionHash: transactionHash ?? hash,
      chainId,
      ...transaction
    }

    const filteredTransactions = transactions.filter(tx => tx.transactionHash !== formattedTx.transactionHash)
    const arr = [...filteredTransactions, formattedTx]

    setTransactions(arr)

  }

  function updateTransactions(txns) {
    setTransactions((oldTransactions) => {
      const newTxnHashes = txns.map(tx => tx.transactionHash ?? tx.hash)


      const filteredPrevTxns = oldTransactions.filter(tx => {
        const txOverlap = newTxnHashes.includes(tx.transactionHash ?? tx.hash)
        return !txOverlap
      })

      const txnsToAdd = _.sortBy([
        ...filteredPrevTxns,
        ...txns.map(tx => {
          return { ...tx, transactionHash: tx.transactionHash ?? tx.hash }
        })
      ],
        (tx) => -tx.timestamp
      )
      return txnsToAdd
    })
  }

  function clear() {
    setTransactions([])
  }


  return {
    transactions,
    setTransactions,
    addTransaction,
    updateTransactions,
    clear
  }
}
