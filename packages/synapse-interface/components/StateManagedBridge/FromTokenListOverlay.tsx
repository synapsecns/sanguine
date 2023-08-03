import _ from 'lodash'

import { useEffect, useMemo, useState } from 'react'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'
import { useAccount } from 'wagmi'

import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { sortTokens } from '@constants/tokens'
import { Token } from '@/utils/types'
import { setFromToken } from '@/slices/bridge/reducer'
import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'
import { CHAINS_BY_ID } from '@/constants/chains'

export const FromTokenListOverlay = () => {
  const [hasMounted, setHasMounted] = useState(false)

  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()

  const { fromTokens, fromChainId, fromToken, toChainId, toToken } =
    useBridgeState()
  const portfolioBalances = usePortfolioBalances()
  const { isConnected } = useAccount()

  let tokenList = sortTokens(fromTokens)

  const fuse = new Fuse(tokenList, {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'symbol',
        weight: 2,
      },
      `addresses.${fromChainId}`,
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

    if (Object.keys(token).length > 0 && tokenWithPb?.balance !== 0n) {
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
    dispatch(setShowFromTokenListOverlay(false))
  }

  function onMenuItemClick(token: Token) {
    dispatch(setFromToken(token))
    dispatch(setShowFromTokenListOverlay(false))
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

  const toChainName = useMemo(() => {
    return CHAINS_BY_ID[toChainId]?.name
  }, [toChainId])

  const tokensWithoutBalancesText = useMemo(() => {
    if (fromChainName && toChainName && toToken) {
      return `Other ${fromChainName} tokens bridgeable to ${toToken.name} on ${toChainName}`
    } else if (!fromChainName && !fromToken && toChainName && !toToken) {
      return `All tokens you can bridge to ${toChainName}`
    } else if (fromChainName && toChainName) {
      return `Other ${fromChainName} tokens bridgeable to ${toChainName}`
    } else if (fromChainName && (!toChainName || !toToken)) {
      return `${fromChainName} tokens`
    } else if (!fromChainName && toChainName && toToken) {
      return `Tokens bridgeable to ${toToken.name} on ${toChainName}`
    } else if (!fromChainName && !toChainName && toToken) {
      return `Tokens bridgeable to ${toToken.name}`
    } else {
      return 'All tokens'
    }
  }, [fromChainId, toChainId, toToken])

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
              selectedToken={fromToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = '[Bridge User Action] Sets new fromToken'
                const eventData = {
                  previousFromToken: fromToken?.symbol,
                  newFromToken: token?.symbol,
                }

                segmentAnalyticsEvent(eventTitle, eventData)
                onMenuItemClick(token)
              }}
            />
          )
        })}
      </div>
      <div className="px-2 pb-2 pt-2 text-secondaryTextColor text-sm bg-[#343036]">
        {tokensWithoutBalancesText}
      </div>
      <div className="px-2 pb-8 bg-[#343036] md:px-2 ">
        {tokenListWithoutBalances.map((token, idx) => {
          return (
            <SelectSpecificTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              selectedToken={fromToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = '[Bridge User Action] Sets new fromToken'
                const eventData = {
                  previousFromToken: fromToken?.symbol,
                  newFromToken: token?.symbol,
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
