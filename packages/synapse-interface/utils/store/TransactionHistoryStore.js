import { createContext, useState } from 'react'

import { useLocalStorage } from '@hooks/store/useLocalStorage'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


export function TransactionHistoryStore({ children }) {
  const { account } = useActiveWeb3React()

  const [transactions, setTransactions] = useState([])

  return (
    <TransactionHistoryContext.Provider value={[transactions, setTransactions]}>
      {children}
    </TransactionHistoryContext.Provider>
  )
}

export const TransactionHistoryContext = createContext([])



