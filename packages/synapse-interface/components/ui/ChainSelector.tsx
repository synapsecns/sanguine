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

  const [currentId, setCurrentId] = useState(null)
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
    setCurrentId(null)
    setHover(false)
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentId(null)
  }

  const arrowUp = useKeyPress('ArrowUp', hover)
  const arrowDown = useKeyPress('ArrowDown', hover)
  const enterPress = useKeyPress('Enter', hover)

  const arrowDownFunc = () => {
    const currentIndex = flatItemList.findIndex((item) => item.id === currentId)
    const nextIndex = currentIndex + 1
    if (arrowDown && nextIndex < flatItemList.length) {
      setCurrentId(flatItemList[nextIndex].id)
    }
  }

  const arrowUpFunc = () => {
    const currentIndex = flatItemList.findIndex((item) => item.id === currentId)
    const prevIndex = currentIndex - 1
    if (arrowUp && prevIndex >= 0) {
      setCurrentId(flatItemList[prevIndex].id)
    }
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(arrowUpFunc, [arrowUp])

  useEffect(() => {
    if (currentId !== null) {
      handleSetChainId(currentId)
      onClose()
    }
  }, [enterPress])

  return (
    <SelectorWrapper
      key={dataTestId}
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      onSearch={onSearch}
      hover={hover}
      setHover={setHover}
      onClose={onClose}
    >
      {Object.entries(itemList).map(([key, value]: [string, Chain[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key} key={key}>
            {value.map((chain) => (
              <SelectSpecificNetworkButton
                dataId={dataTestId}
                key={chain.id}
                itemChainId={chain.id}
                isOrigin={isOrigin}
                isSelected={selectedItem?.id === chain.id}
                isActive={chain.id === currentId}
                onClick={() => handleSetChainId(chain.id)}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
  )
}
