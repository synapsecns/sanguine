import { PAUSED_CHAINS_URL, PAUSED_MODULES_URL } from '@/constants/index'
import { fetchJSONData } from '@/utils/fetchJsonData'

enum SessionStorageKey {
  CHAIN_PAUSE = 'synapse-paused-chains',
  MODULE_PAUSE = 'synapse-paused-modules',
}

let isFetching = false

export const getSynapsePauseData = () => {
  const fetchAndStoreData = async () => {
    if (isFetching) {
      return
    }
    try {
      isFetching = true
      const chainsData = await fetchJSONData(PAUSED_CHAINS_URL)
      const modulesData = await fetchJSONData(PAUSED_MODULES_URL)

      sessionStorage.setItem(
        SessionStorageKey.CHAIN_PAUSE,
        JSON.stringify(chainsData)
      )
      sessionStorage.setItem(
        SessionStorageKey.MODULE_PAUSE,
        JSON.stringify(modulesData)
      )
    } catch (error) {
      console.error(
        '[Synapse Widget] Failed to fetch paused chains/modules: ',
        error
      )
    } finally {
      setTimeout(() => {
        isFetching = false
      }, 1000)
    }
  }

  const readData = () => {
    if (sessionStorage) {
      const pausedChainsData = JSON.parse(
        sessionStorage.getItem(SessionStorageKey.CHAIN_PAUSE)
      )
      const pausedModulesData = JSON.parse(
        sessionStorage.getItem(SessionStorageKey.MODULE_PAUSE)
      )
      return { pausedChainsData, pausedModulesData }
    }
    return { pausedChainsData: null, pausedModulesData: null }
  }

  fetchAndStoreData()

  return readData()
}
