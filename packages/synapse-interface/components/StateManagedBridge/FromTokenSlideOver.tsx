import { useEffect, useState } from 'react'
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
import { TokenWithBalanceAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'

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

  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  let tokenList: any[] = []

  const portfolioBalances = usePortfolioBalances()

  tokenList = tokens

  // Hiding this below for now bc its conflicting with tokens w/ balances
  // tokenList = sortTokens(tokenList)

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

  return (
    <div
      data-test-id="token-slide-over"
      className="max-h-full pb-4 -mt-3 overflow-auto scrollbar-hide"
    >
      <div className="absolute z-10 w-full px-2 pt-3 ">
        <div className="flex items-center mb-2 font-medium justfiy-between sm:float-none">
          <SlideSearchBox
            placeholder="Filter by symbol, contract, or name..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
        </div>
      </div>
      <div className="px-2 pt-14 pb-8 bg-[#343036] md:px-2">
        {tokenList.map((token, idx) => {
          const tokenBalanceAndAllowance =
            chainId &&
            portfolioBalances[chainId]?.filter(
              (t: TokenWithBalanceAndAllowances) => t.token === token
            )
          const balance = tokenBalanceAndAllowance
            ? tokenBalanceAndAllowance[0]?.balance
            : 0n

          return (
            <SelectSpecificTokenButton
              isOrigin={true}
              key={idx}
              chainId={chainId}
              token={token}
              selectedToken={selectedToken}
              active={idx === currentIdx}
              tokenBalance={balance}
              onClick={() => {
                const eventTitle = isOrigin
                  ? '[Bridge User Action] Sets new fromToken'
                  : `[Bridge User Action] Sets new toToken`
                const eventData = isOrigin
                  ? {
                      previousFromToken: selectedToken.symbol,
                      newFromToken: token.symbol,
                    }
                  : {
                      previousToToken: selectedToken.symbol,
                      newToToken: token.symbol,
                    }
                segmentAnalyticsEvent(eventTitle, eventData)
                onMenuItemClick(token)
              }}
            />
          )
        })}

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
