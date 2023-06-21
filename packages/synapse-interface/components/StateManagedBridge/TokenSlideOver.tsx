import { useEffect, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import TokenMenuItem from '@pages/bridge/TokenMenuItem'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { DrawerButton } from '@components/buttons/DrawerButton'
import { sortTokens } from '@constants/tokens'

import { Token } from '@/utils/types'
import { useDispatch } from 'react-redux'

export const TokenSlideOver = ({
  isOrigin,
  tokens = [],
  chainId,
  selectedToken,
  setToken,
  setShowSlideOver,
}: {
  isOrigin: boolean
  tokens: any[]
  chainId: number
  selectedToken: Token
  setToken: any
  setShowSlideOver: any
}) => {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  let tokenList: any[] = []

  tokenList = tokens

  tokenList = sortTokens(tokenList)
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
    <div className="max-h-full pb-4 -mt-3 overflow-auto scrollbar-hide rounded-3xl">
      <div className="absolute z-10 w-full px-6 pt-3 bg-bgLight rounded-t-xl">
        <div className="flex items-center float-right mb-2 font-medium sm:float-none">
          <SlideSearchBox
            placeholder="Search by symbol, contract, or name..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <DrawerButton onClick={onClose} isOrigin={isOrigin} />
        </div>
      </div>
      <div
        className={`
          bg-bgLighter
          space-y-4
          pt-20 pb-8 px-2 md:px-6
          rounded-3xl
        `}
      >
        {tokenList.map((token, idx) => (
          <TokenMenuItem
            key={idx}
            chainId={chainId}
            token={token}
            selectedToken={selectedToken}
            active={idx === currentIdx}
            tokenBalance={token.balance}
            onClick={() => {
              onMenuItemClick(token)
            }}
          />
        ))}
        {searchStr && (
          <div className="px-12 py-4 text-xl text-center text-white">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-4 text-lg text-white text-opacity-50 align-bottom text-medium">
              Want to see a token supported on Synapse? Submit a request{' '}
              <span className="text-white text-opacity-70 hover:underline hover:cursor-pointer">
                here
              </span>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
