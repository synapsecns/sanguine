import { useAppDispatch } from '@/store/hooks'
import {
  setPausedChainsData,
  setPausedModulesData,
  setIsFetching,
} from '@/slices/maintenance/reducer'
import { fetchJSONData } from './fetchJsonData'
import { useMaintenanceState } from '@/slices/maintenance/hooks'

export const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/maintenance/packages/synapse-interface/public/pauses/v1/paused-chains.json'
export const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/maintenance/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

export const useSynapsePauseData = () => {
  const dispatch = useAppDispatch()
  const { isFetching } = useMaintenanceState()

  const startFetching = () => dispatch(setIsFetching(true))
  const stopFetching = () => dispatch(setIsFetching(false))

  const fetchAndStoreData = async () => {
    if (isFetching) return
    try {
      startFetching()

      const pausedChainsData = await fetchJSONData(PAUSED_CHAINS_URL)
      const pausedModulesData = await fetchJSONData(PAUSED_MODULES_URL)

      dispatch(setPausedChainsData(pausedChainsData))
      dispatch(setPausedModulesData(pausedModulesData))
    } catch (error) {
      console.error('Failed to fetch paused chains/modules: ', error)
    } finally {
      setTimeout(() => {
        stopFetching()
      }, 1000)
    }
  }

  return fetchAndStoreData
}
