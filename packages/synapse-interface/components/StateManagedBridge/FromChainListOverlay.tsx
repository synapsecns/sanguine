import _ from 'lodash'
import { useCallback, useEffect, useMemo, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import * as ALL_CHAINS from '@constants/chains/master'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { useDispatch } from 'react-redux'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { resetState, setFromChainId } from '@/slices/bridge/reducer'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { fromChainText } from './helpers/fromChainText'
import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
import { getSymbol } from '@/utils/routeMaker/generateRoutePossibilities'

import * as ALL_TOKENS from '@constants/tokens/master'

export const FromChainListOverlay = () => {
  const { fromChainIds, fromChainId, fromToken, toChainId, toToken } =
    useBridgeState()
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  const dataId = 'bridge-origin-chain-list'

  let possibleChains = _(ALL_CHAINS)
    .pickBy((value) => _.includes(fromChainIds, value.id))
    .values()
    .value()

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
      fromChainIds.map((id) => CHAINS_BY_ID[id])
    )
  )

  const possibleChainsWithSource = possibleChains.map((chain) => ({
    ...chain,
    source: 'possibleChains',
  }))

  const remainingChainsWithSource = remainingChains.map((chain) => ({
    ...chain,
    source: 'remainingChains',
  }))

  const masterList = [...possibleChainsWithSource, ...remainingChainsWithSource]

  const fuseOptions = {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'id',
      'nativeCurrency.symbol',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter((item) => item.source === 'possibleChains')
    remainingChains = results.filter(
      (item) => item.source === 'remainingChains'
    )
  }

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowFromChainListOverlay(false))
  }, [setShowFromChainListOverlay])

  function onMenuItemClick(chainId: number) {
    dispatch(resetState())
    dispatch(setFromChainId(chainId))
  }

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterList.length) {
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
      const chain = masterList[currentIdx]
      dispatch(setFromChainId(chain.id))
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

  const fromChainsText = useMemo(() => {
    return fromChainText({
      fromChainId,
      fromToken,
      toChainId,
      toToken,
    })
  }, [fromChainId, fromToken, toChainId, toToken])

  return (
    <div
      data-test-id="chain-slide-over"
      className="max-h-full pb-4 mt-2 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="flex items-center mb-2 font-medium justfiy-between sm:float-none">
          <SlideSearchBox
            placeholder="Filter by chain name, id, or native currency"
            searchStr={searchStr}
            onSearch={onSearch}
          />
        </div>
      </div>
      <div
        data-test-id={dataId}
        className="px-2 pt-2 pb-8 bg-[#343036] md:px-2"
      >
        <div className="mb-2 text-sm font-normal text-white">
          {fromChainsText}
        </div>
        {possibleChains.map(({ id: mapChainId }, idx) => {
          return (
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={fromChainId === mapChainId}
              active={idx === currentIdx}
              onClick={() => {
                if (fromChainId !== mapChainId) {
                  const eventTitle = `[Bridge User Action] Sets new fromChainId`
                  const eventData = {
                    previousFromChainId: fromChainId,
                    newFromChainId: mapChainId,
                  }

                  segmentAnalyticsEvent(eventTitle, eventData)
                  onMenuItemClick(mapChainId)
                }
                onClose()
              }}
              dataId={dataId}
            />
          )
        })}
        <div className="mt-2 mb-2 text-sm font-normal text-white">
          All other chains
        </div>
        {remainingChains.map(({ id: mapChainId }, idx) => {
          let onClickSpecificNetwork
          if (fromChainId === mapChainId) {
            onClickSpecificNetwork = () => {
              onClose()
            }
          } else {
            onClickSpecificNetwork = () => {
              const eventTitle = `[Bridge User Action] Sets new fromChainId`
              const eventData = {
                previousFromChainId: fromChainId,
                newFromChainId: mapChainId,
              }

              segmentAnalyticsEvent(eventTitle, eventData)
              dispatch(setFromChainId(mapChainId))
              onClose()
            }
          }
          return (
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={fromChainId === mapChainId}
              active={idx === currentIdx}
              onClick={onClickSpecificNetwork}
              dataId={dataId}
            />
          )
        })}
        {searchStr && (
          <div className="px-12 py-4 text-center text-white text-md">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-2 text-white text-opacity-50 align-bottom text-md">
              Want to see a chain supported on Synapse? Let us know!
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
