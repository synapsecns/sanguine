import { createContext } from 'react'

import { useTransactionListener } from '@/hooks/useTransactionListener'

const BackgroundListenerContext = createContext(null)

export const BackgroundListenerProvider = ({ children }) => {
  useTransactionListener()

  return (
    <BackgroundListenerContext.Provider value={null}>
      {children}
    </BackgroundListenerContext.Provider>
  )
}
