import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import {
  BridgeState,
  setIsLoading,
  initialState,
  updateDebouncedFromValue,
  updateDebouncedToTokensFromValue,
} from '@/slices/bridge/reducer'

export const useBridgeListener = () => {
  const dispatch = useAppDispatch()
  const { fromValue, debouncedFromValue }: BridgeState = useBridgeState()

  /**
   * Debounce user input to fetch primary bridge quote (in ms)
   * Delay loading animation when user input updates
   */
  useEffect(() => {
    const DEBOUNCE_DELAY = 300
    const ANIMATION_DELAY = 200

    const animationTimer = setTimeout(() => {
      if (debouncedFromValue !== initialState.debouncedFromValue) {
        dispatch(setIsLoading(true))
      }
    }, ANIMATION_DELAY)

    const debounceTimer = setTimeout(() => {
      dispatch(updateDebouncedFromValue(fromValue))
    }, DEBOUNCE_DELAY)

    return () => {
      clearTimeout(debounceTimer)
      clearTimeout(animationTimer)
      dispatch(setIsLoading(false))
    }
  }, [fromValue])

  // Debounce alternative destination token bridge quotes
  useEffect(() => {
    const ALTERNATE_OPTIONS_DEBOUNCE_DELAY = 1000

    const alternativeOptionsDebounceTimer = setTimeout(() => {
      dispatch(updateDebouncedToTokensFromValue(debouncedFromValue))
    }, ALTERNATE_OPTIONS_DEBOUNCE_DELAY)

    return () => {
      clearTimeout(alternativeOptionsDebounceTimer)
    }
  }, [debouncedFromValue])

  return null
}
