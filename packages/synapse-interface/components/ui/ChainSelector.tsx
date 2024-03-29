import React, { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'

import { type Chain } from '@/utils/types'
import { SelectSpecificNetworkButton } from '@/components/ui/SelectSpecificNetworkButton'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { SelectorWrapper } from '@/components/ui/SelectorWrapper'
import { ListSectionWrapper } from '@/components/ui/ListSectionWrapper'
import { ChainSelectorTypes } from '@/components/ui/types'
import { useKeyPress } from '@/utils/hooks/useKeyPress'

export function ChainSelector({
  dataTestId,
  isOrigin,
  selectedItem,
  label,
  placeholder,
  itemListFunction,
  setFunction,
  action,
}: ChainSelectorTypes) {
  const [searchStr, setSearchStr] = useState('')
  const [hover, setHover] = useState(false)

  const [currentIdx, setCurrentIdx] = useState(-1)
  const dispatch = useDispatch()

  const handleSetChainId = (chainId) => {
    if (selectedItem?.id !== chainId) {
      const eventTitle = `[${action} User Action] Sets new ${
        isOrigin ? 'from' : 'to'
      }ChainId`

      const eventData = isOrigin
        ? {
            previousFromChainId: selectedItem?.id,
            newFromChainId: chainId,
          }
        : {
            previousToChainId: selectedItem?.id,
            newToChainId: chainId,
          }

      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(chainId))
    }
  }

  const itemList = itemListFunction(searchStr)
  const flatItemList = Object.entries(itemList).reduce(
    (acc, [_, value]) => [...acc, ...(value as Chain[])],
    []
  )

  const onClose = () => {
    setSearchStr('')
    setHover(false)
  }

  const arrowUp = useKeyPress('ArrowUp', hover)
  const arrowDown = useKeyPress('ArrowDown', hover)
  const enterPress = useKeyPress('Enter', hover)

  const arrowDownFunc = () => {
    console.log(`in isOrigin: ${isOrigin}, down fn`)
    console.log(`flatItemList[currentIdx]`, flatItemList[currentIdx])
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < flatItemList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    console.log(`in isOrigin: ${isOrigin}, up fn`)
    console.log(`flatItemList[currentIdx]`, flatItemList[currentIdx])
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(arrowUpFunc, [arrowUp])

  useEffect(() => {
    console.log(`currentIdx`, currentIdx)
    console.log(`flatItemList`, flatItemList[currentIdx])
    if (currentIdx >= 0 && flatItemList[currentIdx]) {
      console.log(`flatItemList[currentIdx]`, flatItemList[currentIdx])
      handleSetChainId(flatItemList[currentIdx].id)
    }
    onClose()
  }, [enterPress])

  return (
    <SelectorWrapper
      key={dataTestId}
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      setSearchStr={setSearchStr}
      hover={hover}
      setHover={setHover}
      onClose={onClose}
    >
      {Object.entries(itemList).map(
        ([key, value]: [string, Chain[]], index) => {
          return value.length ? (
            <ListSectionWrapper sectionKey={key}>
              {value.map((chain, chainIndex) => (
                <SelectSpecificNetworkButton
                  dataId={dataTestId}
                  key={chain.id}
                  itemChainId={chain.id}
                  isOrigin={isOrigin}
                  isCurrentChain={currentIdx === index + chainIndex}
                  active={false}
                  onClick={() => handleSetChainId(chain.id)}
                />
              ))}
            </ListSectionWrapper>
          ) : null
        }
      )}
    </SelectorWrapper>
  )
}
