import { useAppDispatch } from '@/store/hooks'
import {
  setPausedChainsData,
  setPausedModulesData,
} from '@/slices/maintenance/reducer'
import { fetchJsonData } from './fetchJsonData'
import pausedChains from '@/public/pauses/v1/paused-chains.json'
import pausedBridgeModules from '@/public/pauses/v1/paused-bridge-modules.json'

export const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-chains.json'
export const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

export const PAUSED_CHAINS_CDN_URL =
  'https://cdn.jsdelivr.net/gh/synapsecns/sanguine@master/packages/synapse-interface/public/pauses/v1/paused-chains.json'
export const PAUSED_MODULES_CDN_URL =
  'https://cdn.jsdelivr.net/gh/synapsecns/sanguine@master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

let isFetching = false

export const getSynapsePauseData = () => {
  const dispatch = useAppDispatch()

  const fetchAndStoreData = async () => {
    if (isFetching) return
    try {
      isFetching = true

      const pausedChainsData = await fetchJsonData(PAUSED_CHAINS_URL)
      const pausedModulesData = await fetchJsonData(PAUSED_MODULES_URL)

      dispatch(setPausedChainsData(pausedChainsData))
      dispatch(setPausedModulesData(pausedModulesData))
    } catch (error) {
      console.error(
        'Using local source, failed to fetch paused chains/modules: ',
        error
      )

      /** Read local source if fetch fails as backup */
      dispatch(setPausedChainsData(pausedChains))
      dispatch(setPausedModulesData(pausedBridgeModules))
    } finally {
      setTimeout(() => {
        isFetching = false
      }, 1000)
    }
  }

  return fetchAndStoreData
}
