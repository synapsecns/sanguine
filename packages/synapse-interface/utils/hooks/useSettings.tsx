import createPersistedState from 'use-persisted-state'

const usePassthroughSettings = createPersistedState('settingsObjStuff')

export const useSettings = () => {
  return usePassthroughSettings({
    expertMode: false,
  })
}
