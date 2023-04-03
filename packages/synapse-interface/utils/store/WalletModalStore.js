import { createContext, useState } from 'react'

import { useLocalStorage } from '@hooks/store/useLocalStorage'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


export function WalletModalStore({ children }) {
  const { account } = useActiveWeb3React()


  const [showWalletModal, setShowWalletModal] = useState()

  return (
    <WalletModalContext.Provider value={[showWalletModal, setShowWalletModal]}>
      {children}
    </WalletModalContext.Provider>
  )
}

export const WalletModalContext = createContext([])



