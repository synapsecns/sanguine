import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  setPausedChainData,
  setPausedModuleData,
} from '@/slices/maintenance/reducer'
import { fetchJSONData } from './fetchJsonData'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const PAUSED_CHAINS_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-chains.json'
export const PAUSED_MODULES_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

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

      dispatch(setPausedChainData(pausedChainsData))
      dispatch(setPausedModuleData(pausedModulesData))
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

  return fetchAndStoreData
}
