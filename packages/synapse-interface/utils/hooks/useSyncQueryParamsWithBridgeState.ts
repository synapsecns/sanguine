import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useDispatch } from 'react-redux'
import { Token } from '@types'
import * as BRIDGEABLE from '@constants/tokens/bridgeable'

import { useBridgeState } from '@/slices/bridge/hooks'
import {
  setFromChainId,
  setFromToken,
  setToChainId,
  setToToken,
} from '@/slices/bridge/reducer'
import { getAllFromChainIds } from '../routeMaker/getFromChainIds'
import { getAllToChainIds } from '../routeMaker/getToChainIds'

const useSyncQueryParamsWithBridgeState = () => {
  const dispatch = useDispatch()
  const router = useRouter()
  const { fromChainId, fromToken, toChainId, toToken } = useBridgeState()

  const allowedFromChainIds = getAllFromChainIds()
  const allowedToChainIds = getAllToChainIds()

  // Sync URL params to state on load or when URL changes
  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search)
    const fromChainParam = urlParams.get('fromChainId')
    const toChainParam = urlParams.get('toChainId')
    const fromTokenParam = urlParams.get('fromToken')
    const toTokenParam = urlParams.get('toToken')

    if (fromChainParam !== null && Number(fromChainParam) !== fromChainId) {
      if (allowedFromChainIds.includes(Number(fromChainParam))) {
        dispatch(setFromChainId(Number(fromChainParam)))
      }
    }

    if (fromTokenParam !== null && fromTokenParam !== fromToken?.symbol) {
      const token = findKeyByRouteSymbol(fromTokenParam, BRIDGEABLE)
      dispatch(setFromToken(token))
    }

    if (toChainParam !== null && Number(toChainParam) !== toChainId) {
      if (allowedToChainIds.includes(Number(toChainParam))) {
        dispatch(setToChainId(Number(toChainParam)))
      }
    }

    if (toTokenParam !== null && toTokenParam !== toToken?.symbol) {
      const token = findKeyByRouteSymbol(toTokenParam, BRIDGEABLE)
      dispatch(setToToken(token))
    }
  }, [router.asPath])

  // Sync state to URL params
  useEffect(() => {
    const queryParams = new URLSearchParams()

    if (fromChainId !== null)
      queryParams.set('fromChainId', fromChainId.toString())
    if (fromToken?.symbol) queryParams.set('fromToken', fromToken.symbol)
    if (toChainId !== null) queryParams.set('toChainId', toChainId.toString())
    if (toToken?.symbol) queryParams.set('toToken', toToken.symbol)

    const newSearch = queryParams.toString()

    const newUrl = {
      pathname: router.pathname,
      search: newSearch,
    }

    router.replace(newUrl, undefined, { shallow: true })
  }, [fromChainId, fromToken, toChainId, toToken])

  return { fromChainId, fromToken, toChainId, toToken }
}

export default useSyncQueryParamsWithBridgeState

const findKeyByRouteSymbol = (symbol, MODULE): Token | null => {
  const searchSymbol = symbol.toLowerCase()
  for (const [key, value] of Object.entries(MODULE)) {
    if ((value as Token).routeSymbol.toLowerCase() === searchSymbol) {
      return MODULE[key]
    }
  }
  return null
}
