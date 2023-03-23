import _ from 'lodash'
import { useEffect, useRef, useState, useCallback } from 'react'

import Fuse from 'fuse.js'

import { useKeyPress } from '@hooks/useKeyPress'

import {
  CHAIN_ID_DISPLAY_ORDER,
  CHAIN_INFO_MAP,
  ChainInfoMap,
  ChainInfo,
} from '@constants/networks'

import { SelectSpecificNetworkButton } from '@components/buttons/SelectSpecificNetworkButton'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { DrawerButton } from '@components/buttons/DrawerButton'
import { useNetwork } from 'wagmi'

export function NetworkSlideOver({
  selectedChainId,
  isSwapFrom,
  onChangeChain,
  setDisplayType,
  possibleChains,
}: {
  selectedChainId: number
  isSwapFrom: boolean
  onChangeChain: (v: number) => void
  setDisplayType: (v: string) => void
  possibleChains: string[] | undefined
}) {
  const { chain } = useNetwork()
  const [currentIdx, setCurrentIdx] = useState(-1)

  const [searchStr, setSearchStr] = useState('')
  const [networks, setNetworks] = useState<ChainInfo[]>([])
  const fuse = new Fuse(networks, {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'chainName',
        weight: 2,
      },
      'chainShortName',
      'chainId',
      'nativeCurrency',
    ],
  })
  // let networks: ChainInfo[] = []

  useEffect(() => {
    let tempNetworks: ChainInfo[] = []
    CHAIN_ID_DISPLAY_ORDER.map((cid) => {
      console.log('CIIIID', possibleChains, cid)
      if (
        isSwapFrom ||
        (!isSwapFrom &&
          possibleChains?.includes(String(cid)) &&
          cid !== Number(chain?.id))
      ) {
        tempNetworks.push(CHAIN_INFO_MAP[cid])
      }
    })
    console.log('tempNetworks', tempNetworks, isSwapFrom, Number(chain?.id))
    if (searchStr?.length > 0) {
      tempNetworks = fuse.search(searchStr).map((i) => i.item)
    }
    setNetworks(tempNetworks)
  }, [chain, searchStr])

  const ref = useRef(null)

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setDisplayType('')
  }, [setDisplayType])

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  function arrowDownFunc() {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < networks.length) {
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
      let currentChain = networks[currentIdx]
      onChangeChain(currentChain.chainId)
      onClose()
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
            placeholder="Search by asset, name, or chainID..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <DrawerButton onClick={onClose} />
        </div>
      </div>
      <div
        ref={ref}
        className="px-3 pt-20 pb-8 space-y-4 bg-bgLighter md:px-6 rounded-xl"
      >
        {networks.map((chainData, idx) => {
          const itemChainId = chainData.chainId
          let chaindata = itemChainId
          const isCurrentChain = selectedChainId === itemChainId

          let onClickSpecificNetwork
          if (isCurrentChain) {
            onClickSpecificNetwork = () => console.log('INCEPTION')
          } else {
            onClickSpecificNetwork = () => {
              onChangeChain(chainData.chainId)
              onClose()
            }
          }
          return (
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={itemChainId}
              isCurrentChain={isCurrentChain}
              active={idx === currentIdx}
              onClick={onClickSpecificNetwork}
            />
          )
        })}
        {searchStr && (
          <div className="px-12 py-4 text-xl text-center text-white">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-4 text-lg text-white text-opacity-50 align-bottom text-medium">
              Want to see a chain supported on Synapse? Submit a request{' '}
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
