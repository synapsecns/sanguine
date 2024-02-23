import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useDispatch } from 'react-redux'
import type { Token } from '@types'

import {
  setSwapChainId,
  setSwapFromToken,
  setSwapToToken,
} from '@/slices/swap/reducer'
import { SWAP_CHAIN_IDS } from '@/constants/existingSwapRoutes'
import * as BRIDGEABLE from '@/constants/tokens/bridgeable'
import { useSwapState } from '@/slices/swap/hooks'

const useSyncQueryParamsWithSwapState = () => {
  const dispatch = useDispatch()
  const router = useRouter()
  const { swapChainId, swapFromToken, swapToToken } = useSwapState()

  const allowedSwapChainIds = SWAP_CHAIN_IDS

  // Sync URL params to state on load or when URL changes
  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search)
    const swapChainParam = urlParams.get('chainId')
    const swapFromTokenParam = urlParams.get('fromToken')
    const swapToTokenParam = urlParams.get('toToken')

    if (swapChainParam !== null) {
      if (allowedSwapChainIds.includes(Number(swapChainParam))) {
        dispatch(setSwapChainId(Number(swapChainParam)))
      }
    }

    if (swapFromTokenParam !== null) {
      const token = findKeyByRouteSymbol(swapFromTokenParam, BRIDGEABLE)
      dispatch(setSwapFromToken(token))
    }

    if (swapToTokenParam !== null) {
      const token = findKeyByRouteSymbol(swapToTokenParam, BRIDGEABLE)
      dispatch(setSwapToToken(token))
    }
  }, [router.asPath])

  // Sync state to URL params
  useEffect(() => {
    const queryParams = new URLSearchParams()

    if (swapChainId !== null) queryParams.set('chainId', swapChainId.toString())
    if (swapFromToken?.symbol)
      queryParams.set('fromToken', swapFromToken.symbol)
    if (swapToToken?.symbol) queryParams.set('toToken', swapToToken.symbol)

    const newSearch = queryParams.toString()

    const newUrl = {
      pathname: router.pathname,
      search: newSearch,
    }

    router.replace(newUrl, undefined, { shallow: true })
  }, [swapChainId, swapFromToken, swapToToken])

  return { swapChainId, swapFromToken, swapToToken }
}

export default useSyncQueryParamsWithSwapState

const findKeyByRouteSymbol = (symbol, MODULE): Token | null => {
  const searchSymbol = symbol.toLowerCase()
  for (const [key, value] of Object.entries(MODULE)) {
    if ((value as Token).routeSymbol.toLowerCase() === searchSymbol) {
      return MODULE[key]
    }
  }
  return null
}
