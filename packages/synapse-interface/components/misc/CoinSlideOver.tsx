import _ from 'lodash'
import { useEffect, useState } from 'react'

import Fuse from 'fuse.js'
import Zero from '@ethersproject/constants'
import { ChevronUpIcon } from '@heroicons/react/outline'
import { useBalance, useAccount } from 'wagmi'

import { useKeyPress } from '@hooks/useKeyPress'
// import { useGenericTokenBalance } from '@hooks/tokens/useTokenBalances'

import TokenMenuItem from '@pages/bridge/TokenMenuItem'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { DrawerButton } from '@components/buttons/DrawerButton'
import { Token } from '@utils/classes/Token'
import { BigNumber } from 'ethers'

export const CoinSlideOver = ({
  chainId,
  tokens,
  selected,
  setDisplayType,
  handleTokenChange,
  address,
  isSwapFrom,
}: {
  address: `0x${string}` | undefined
  chainId: number
  tokens: Token[]
  selected: any
  setDisplayType: (v: string) => void
  handleTokenChange: (token: Token, type: 'from' | 'to') => void
  isSwapFrom: boolean
}) => {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')

  // const [tokenList, setTokenList] = useState<any[]>([])
  // const [tokenBalances, setTokenBalances] = useState<Map<string, BigNumber>>(
  //   new Map([])
  // )
  let tokenBalances: Map<string, BigNumber> = new Map([])
  let tokenList: any[] = []

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

  const { tokenList: rawTokenList, tokenBalances: rawTokenBalances } =
    sortByTokenBalance(tokens, chainId, address)
  tokenList = rawTokenList
  if (searchStr?.length > 0) {
    tokenList = fuse.search(searchStr).map((i) => i.item)
  }
  // useEffect(() => {
  //   let tempTokenList = rawTokenList
  //   let tempTokenBalances: Map<string, BigNumber> = new Map()
  //   console.log('searchStr', searchStr)
  //   if (searchStr?.length > 0) {
  //     tempTokenList = fuse.search(searchStr).map((i) => i.item)
  //   }
  //   setTokenBalances(rawTokenBalances)
  //   setTokenList(tempTokenList)
  // }, [rawTokenList, rawTokenBalances])

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  function onClose() {
    setCurrentIdx(-1)
    setDisplayType('')
  }

  function onMenuItemClick(coin: any) {
    handleTokenChange(coin, isSwapFrom ? 'from' : 'to')
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
            coin={token}
            selected={selected}
            active={idx === currentIdx}
            tokenBalance={tokenBalances.get(token.addresses[chainId])}
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

const sortByTokenBalance = (
  tokens: Token[],
  chainId: number,
  address: any
): any => {
  let tokenBalances = new Map<string, BigNumber>()
  let nonZeroTokens: Token[] = []
  let zeroTokens: Token[] = []
  let i = 0
  // go through all tokens and retrieve token balances
  while (i < tokens.length) {
    const tokenAddr = tokens[i].addresses[chainId as keyof Token['addresses']]
    if (!tokenBalances.get(tokenAddr)) {
      let rawTokenBalance: any
      // Check for native token
      if (tokenAddr === '') {
        const { data } = useBalance({
          address: address,
          chainId: chainId,
        })
        rawTokenBalance = data
      } else if (tokenAddr?.length > 0) {
        const { data } = useBalance({
          address: address,
          token: `0x${tokenAddr.slice(2)}`,
          chainId: chainId,
        })
        rawTokenBalance = data
      }
      // manages two the array of tokens with zero balances and non-zero balances
      if (rawTokenBalance) {
        tokenBalances.set(tokenAddr, rawTokenBalance.value)
        if (rawTokenBalance?.value._hex !== '0x00') {
          zeroTokens.push(tokens[i])
        } else {
          nonZeroTokens.push(tokens[i])
        }
      }
    }
    i++
  }
  let tokenList = zeroTokens.concat(nonZeroTokens)
  // console.log('tokenBalances', tokenBalances)
  return { tokenList, tokenBalances }
}
