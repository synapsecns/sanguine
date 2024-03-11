import React, { createContext } from 'react'

import { useApplicationListener } from '@/utils/hooks/useApplicationListener'
import { useBridgeListener } from '@/utils/hooks/useBridgeListener'
import { usePortfolioListener } from '@/utils/hooks/usePortfolioListener'
import { useRiskEvent } from '@/utils/hooks/useRiskEvent'
import { useTransactionListener } from '@/utils/hooks/useTransactionListener'
import { use_TransactionsListener } from '@/utils/hooks/use_TransactionsListener'
import { useFetchPricesOnInterval } from '@/utils/hooks/useFetchPricesOnInterval'

const BackgroundListenerContext = createContext(null)

export const BackgroundListenerProvider = ({ children }) => {
  useApplicationListener()
  usePortfolioListener()
  useTransactionListener()
  use_TransactionsListener()
  useBridgeListener()
  useRiskEvent()
  useFetchPricesOnInterval()

  return (
    <BackgroundListenerContext.Provider value={null}>
      {children}
    </BackgroundListenerContext.Provider>
  )
}
