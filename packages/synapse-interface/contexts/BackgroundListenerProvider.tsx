import { createContext } from 'react'
import { useApplicationListener } from '@/utils/hooks/useApplicationListener'
import { useBridgeListener } from '@/utils/hooks/useBridgeListener'
import { usePortfolioListener } from '@/utils/hooks/usePortfolioListener'
import { useTransactionListener } from '@/utils/hooks/useTransactionListener'
import { use_TransactionsListener } from '@/utils/hooks/use_TransactionsListener'


const BackgroundListenerContext = createContext(null)

export const BackgroundListenerProvider = ({ children }) => {
  useApplicationListener()
  usePortfolioListener()
  useTransactionListener()
  use_TransactionsListener()
  useBridgeListener()

  return (
    <BackgroundListenerContext.Provider value={null}>
      {children}
    </BackgroundListenerContext.Provider>
  )
}
