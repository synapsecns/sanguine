import _ from 'lodash'

import { useEffect, useMemo, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { sortTokens } from '@constants/tokens'
import { Token } from '@/utils/types'
import { useDispatch } from 'react-redux'
import { setFromToken, setToToken } from '@/slices/bridge/reducer'
import {
  setShowFromTokenSlideOver,
  setShowToTokenSlideOver,
} from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'
import { useAccount } from 'wagmi'
import { CHAINS_BY_ID } from '@/constants/chains'

export const FromTokenSlideOver = ({}: {}) => {
  let setToken
  let setShowSlideOver

  const { fromTokens, fromChainId, fromToken } = useBridgeState()

  const isOrigin = true

  const tokens = fromTokens

  const chainId = fromChainId

  const selectedToken = fromToken

  if (isOrigin) {
    setToken = setFromToken
    setShowSlideOver = setShowFromTokenSlideOver
  } else {
    setToken = setToToken
    setShowSlideOver = setShowToTokenSlideOver
  }

  const [hasMounted, setHasMounted] = useState(false)

  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  let tokenList: any[] = []

  const portfolioBalances = usePortfolioBalances()
  const { isConnected } = useAccount()

  tokenList = tokens

  const fuse = new Fuse(tokenList, {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'symbol',
        weight: 2,
      },
      `addresses.${chainId}`,
      'name',
    ],
  })

  if (searchStr?.length > 0) {
    tokenList = fuse.search(searchStr).map((i) => i.item)
  }

  const tokenListWithBalances = tokenList.filter((t) => {
    const pb = portfolioBalances[fromChainId]
    const token = _(pb)
      .pickBy((value, _key) => value.token === t)
      .value()

    const tokenWithPb = Object.values(token)[0]

    if (tokenWithPb && tokenWithPb.parsedBalance !== '0.0') {
      return true
    } else {
      return false
    }
  })

  const tokenListWithoutBalances = _.difference(
    tokenList,
    tokenListWithBalances
  )

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  function onClose() {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowSlideOver(false))
  }

  function onMenuItemClick(token: Token) {
    dispatch(setToken(token))
    dispatch(setShowSlideOver(false))
    onClose()
  }

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  function arrowDownFunc() {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < tokenList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  useEffect(arrowDownFunc, [arrowDown])

  function arrowUpFunc() {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  useEffect(arrowUpFunc, [arrowUp])

  function enterPressedFunc() {
    if (enterPressed && currentIdx > -1) {
      onMenuItemClick(tokenList[currentIdx])
    }
  }

  useEffect(enterPressedFunc, [enterPressed])

  function onSearch(str: string) {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const fromChainName = useMemo(() => {
    return CHAINS_BY_ID[fromChainId]?.name
  }, [fromChainId])

  return (
    <div
      data-test-id="token-slide-over"
      className="max-h-full pb-4 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="flex items-center mt-2 mb-2 font-medium justfiy-between sm:float-none">
          <SlideSearchBox
            placeholder="Filter by symbol, contract, or name..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
        </div>
      </div>
      {hasMounted && isConnected && fromChainId && (
        <div className="px-2 pt-4 pb-2 text-secondaryTextColor text-sm bg-[#343036]">
          Wallet
        </div>
      )}
      <div className="px-2 pb-2 bg-[#343036] md:px-2">
        {tokenListWithBalances.map((token, idx) => {
          return (
            <SelectSpecificTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              selectedToken={selectedToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = isOrigin
                  ? '[Bridge User Action] Sets new fromToken'
                  : `[Bridge User Action] Sets new toToken`
                const eventData = isOrigin
                  ? {
                      previousFromToken: selectedToken?.symbol,
                      newFromToken: token?.symbol,
                    }
                  : {
                      previousToToken: selectedToken?.symbol,
                      newToToken: token?.symbol,
                    }
                segmentAnalyticsEvent(eventTitle, eventData)
                onMenuItemClick(token)
              }}
            />
          )
        })}
      </div>
      <div className="px-2 pb-2 pt-2 text-secondaryTextColor text-sm bg-[#343036]">
        {fromChainName ? `${fromChainName} tokens` : `All tokens`}
      </div>
      <div className="px-2 pb-8 bg-[#343036] md:px-2 ">
        {tokenListWithoutBalances.map((token, idx) => {
          return (
            <SelectSpecificTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              selectedToken={selectedToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = isOrigin
                  ? '[Bridge User Action] Sets new fromToken'
                  : `[Bridge User Action] Sets new toToken`
                const eventData = isOrigin
                  ? {
                      previousFromToken: selectedToken?.symbol,
                      newFromToken: token?.symbol,
                    }
                  : {
                      previousToToken: selectedToken?.symbol,
                      newToToken: token?.symbol,
                    }
                segmentAnalyticsEvent(eventTitle, eventData)
                onMenuItemClick(token)
              }}
            />
          )
        })}
      </div>
      <div>
        {searchStr && (
          <div className="px-12 py-4 text-center text-white text-md">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-2 text-white text-opacity-50 align-bottom text-md">
              Want to see a token supported on Synapse? Let us know!
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
