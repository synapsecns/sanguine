import { useAppDispatch } from '@/store/hooks'
import {
  setPausedChainsData,
  setPausedModulesData,
} from '@/slices/maintenance/reducer'
import { fetchJSONData } from './fetchJsonData'

export const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/maintenance/packages/synapse-interface/public/pauses/v1/paused-chains.json'
export const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/maintenance/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

let isFetching = false

export const getSynapsePauseData = () => {
  const dispatch = useAppDispatch()

  const fetchAndStoreData = async () => {
    if (isFetching) {
      return
    }
    try {
      isFetching = true
      const pausedChainsData = await fetchJSONData(PAUSED_CHAINS_URL)
      const pausedModulesData = await fetchJSONData(PAUSED_MODULES_URL)

      dispatch(setPausedChainsData(pausedChainsData))
      dispatch(setPausedModulesData(pausedModulesData))
    } catch (error) {
      console.error('Failed to fetch paused chains/modules: ', error)
    } finally {
      setTimeout(() => {
        isFetching = false
      }, 1000)
    }
  }

  return fetchAndStoreData
}
