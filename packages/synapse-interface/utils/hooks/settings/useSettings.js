import { SettingsContext } from "@store/SettingsStore"
import { useContext } from "react"

import createPersistedState from 'use-persisted-state'

const usePassthroughSettings = createPersistedState('settingsObjStuff')

export function useSettings() {
  // const [settings, setSettings] = useContext(SettingsContext)

  return usePassthroughSettings({
    expertMode: false
  })
}

