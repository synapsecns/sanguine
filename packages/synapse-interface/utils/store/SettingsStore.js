import { createContext } from 'react'

import { useLocalStorage } from '@hooks/store/useLocalStorage'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'



const DEFAULT_SETTINGS = {
  expertMode: false
}

export function SettingsStore({ children }) {
  const { account } = useActiveWeb3React()

  const settingsKey = `settings_${account}`

  const [settings, setSettings] = useLocalStorage(settingsKey, {
    expertMode: false
  })

  return (
    <SettingsContext.Provider value={[settings, setSettings]}>
      {children}
    </SettingsContext.Provider>
  )
}

export const SettingsContext = createContext([])



