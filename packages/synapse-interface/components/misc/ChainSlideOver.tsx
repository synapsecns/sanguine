import { useCallback, useEffect, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import * as CHAINS from '@constants/chains/master'
import { SelectSpecificNetworkButton } from '@components/buttons/SelectSpecificNetworkButton'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { DrawerButton } from '@components/buttons/DrawerButton'
import { useNetwork } from 'wagmi'
import { DisplayType } from '@/pages/bridge/DisplayType'
import { sortChains } from '@constants/chains'

export const ChainSlideOver = ({
  isOrigin,
  chains,
  chainId,
  onChangeChain,
  setDisplayType,
}: {
  isOrigin: boolean
  chains: string[]
  chainId: number
  onChangeChain: (chainId: number, flip: boolean, type: 'from' | 'to') => void
  setDisplayType: (v: DisplayType) => void
}) => {
  const { chain } = useNetwork()
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const [networks, setNetworks] = useState([])
  const fuse = new Fuse(networks, {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'chainShortName',
      'chainId',
      'nativeCurrency',
    ],
  })
  // let networks: ChainInfo[] = []

  const dataId = isOrigin
    ? 'bridge-origin-chain-list'
    : 'bridge-destination-chain-list'

  useEffect(() => {
    let tempNetworks = []
    Object.values(CHAINS).map((chain) => {
      if (isOrigin || (!isOrigin && chains?.includes(String(chain.id)))) {
        tempNetworks.push(chain)
      }
    })
    tempNetworks = sortChains(tempNetworks)
    if (searchStr?.length > 0) {
      tempNetworks = fuse.search(searchStr).map((i) => i.item)
    }
    setNetworks(tempNetworks)
  }, [chain, searchStr])

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setDisplayType(DisplayType.DEFAULT)
  }, [setDisplayType])

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < networks.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  const enterPressedFunc = () => {
    if (enterPressed && currentIdx > -1) {
      const currentChain = networks[currentIdx]
      onChangeChain(currentChain.chainId, false, isOrigin ? 'from' : 'to')
      onClose()
    }
  }
  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(escFunc, [escPressed])
  useEffect(arrowUpFunc, [arrowUp])
  useEffect(enterPressedFunc, [enterPressed])

  return (
    <div className="max-h-full pb-4 -mt-3 overflow-auto scrollbar-hide rounded-lg">
      <div className="absolute z-10 w-full px-6 pt-3 bg-bgLight rounded-t-md">
        <div className="flex items-center float-right mb-2 font-medium sm:float-none">
          <SlideSearchBox
            placeholder="Search by asset, name, or chainID..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <DrawerButton onClick={onClose} isOrigin={isOrigin} />
        </div>
      </div>
      <div data-test-id={dataId} className="pt-20 pb-8 space-y-4 bg-bgLighter">
        {networks.map(({ id: mapChainId }, idx) => {
          let onClickSpecificNetwork
          if (chainId === mapChainId) {
            onClickSpecificNetwork = () => console.log('INCEPTION') // I think this case is obsolete
          } else {
            onClickSpecificNetwork = () => {
              onChangeChain(mapChainId, false, isOrigin ? 'from' : 'to')
              onClose()
            }
          }
          return (
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={chainId === mapChainId}
              active={idx === currentIdx}
              onClick={onClickSpecificNetwork}
              dataId={dataId}
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
