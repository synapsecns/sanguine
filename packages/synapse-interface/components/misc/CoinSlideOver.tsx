import _ from 'lodash'
import { useEffect, useRef, useState } from 'react'

import Fuse from 'fuse.js'

import { ChevronUpIcon } from '@heroicons/react/outline'
import { useBalance } from 'wagmi'

import { useKeyPress } from '@hooks/useKeyPress'
// import { useGenericTokenBalance } from '@hooks/tokens/useTokenBalances'

import TokenMenuItem from '@pages/bridge/TokenMenuItem'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { DrawerButton } from '@components/buttons/DrawerButton'
import { Token } from '@utils/classes/Token'

export function CoinSlideOver({
  chainId,
  tokens,
  onChangeSelected,
  selected,
  setDisplayType,
}: {
  chainId: number
  tokens: Token[]
  onChangeSelected: (val: any) => void
  selected: any
  setDisplayType: (val: string) => void
}) {
  const [currentIdx, setCurrentIdx] = useState(-1)

  const [searchStr, setSearchStr] = useState('')

  let balanceSortedTokens = sortByTokenBalance(tokens, chainId)

  const fuse = new Fuse(balanceSortedTokens, {
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

  let resultTokens: any[] = []
  if (searchStr?.length > 0) {
    resultTokens = fuse.search(searchStr).map((i) => i.item)
  } else {
    resultTokens = balanceSortedTokens
  }

  const ref = useRef(null)

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  function onClose() {
    setCurrentIdx(-1)
    setDisplayType('')
  }

  function onMenuItemClick(coin: any) {
    onChangeSelected(coin)
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
    if (arrowDown && nextIdx < resultTokens.length) {
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
      onMenuItemClick(resultTokens[currentIdx])
    }
  }

  useEffect(enterPressedFunc, [enterPressed])

  // useEffect(() => ref?.current?.scrollTo(0, 0), [])
  useEffect(() => window.scrollTo(0, 0), [])

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
          <DrawerButton onClick={onClose} />
        </div>
      </div>
      <div
        ref={ref}
        className={`
          bg-bgLighter
          space-y-4
          pt-20 pb-8 px-2 md:px-6
          rounded-3xl
        `}
      >
        {resultTokens.map((coin, idx) => (
          <TokenMenuItem
            key={idx}
            chainId={chainId}
            coin={coin}
            selected={selected}
            active={idx === currentIdx}
            onClick={() => {
              onMenuItemClick(coin)
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

function sortByTokenBalance(tokens: Token[], chainId: number) {
  let nonZeroTokens = tokens.filter((token) => {
    const tokenAddr = token.addresses[chainId as keyof Token['addresses']]
    const { data: rawTokenBalance } = useBalance({
      address: `0x${tokenAddr.slice(2)}`,
    })
    rawTokenBalance?.value
    return rawTokenBalance?.value._hex !== '0x00'
  })

  let zeroTokens = tokens.filter((token) => {
    const tokenAddr = token.addresses[chainId as keyof Token['addresses']]
    const { data: rawTokenBalance } = useBalance({
      address: `0x${tokenAddr.slice(2)}`,
    })
    rawTokenBalance?.value
    return rawTokenBalance?.value._hex === '0x00'
  })

  return nonZeroTokens.concat(zeroTokens)
}
