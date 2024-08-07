import { useState, useCallback } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  setPausedChainsData,
  setPausedModulesData,
} from '@/slices/maintenance/reducer'
import { fetchJsonData } from '../functions/fetchJsonData'
import pausedChains from '@/public/pauses/v1/paused-chains.json'
import pausedBridgeModules from '@/public/pauses/v1/paused-bridge-modules.json'

const PAUSED_CHAINS_URL =
  'https://bigboydiamonds.github.io/sanguine/packages/synapse-interface/public/pauses/v1/paused-chains.json'
const PAUSED_MODULES_URL =
  'https://synapsecns.github.io/sanguine/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json'

export const useSynapsePauseData = () => {
  const dispatch = useAppDispatch()
  const [isFetching, setIsFetching] = useState(false)

  const fetchAndStoreData = useCallback(async () => {
    if (isFetching) return
    setIsFetching(true)

    try {
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
        setIsFetching(false)
      }, 1000)
    }
  }, [dispatch, isFetching])

  return fetchAndStoreData
}
