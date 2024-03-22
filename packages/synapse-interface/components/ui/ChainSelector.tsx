import React, { useState } from 'react'
import { type Chain } from '@/utils/types'
import { SelectSpecificNetworkButton } from '../StateManagedBridge/components/SelectSpecificNetworkButton'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useDispatch } from 'react-redux'
import { SelectorWrapper } from '@/components/ui/SelectorWrapper'
import { ListSectionWrapper } from '@/components/ui/ListSectionWrapper'
import { ChainSelectorTypes } from '@/components/ui/types'

export function ChainSelector({
  dataTestId,
  isOrigin,
  selectedItem,
  label,
  placeholder,
  itemListFunction,
  setFunction,
}: ChainSelectorTypes) {
  const [searchStr, setSearchStr] = useState('')

  const dispatch = useDispatch()

  const handleSetChainId = (chainId) => {
    if (selectedItem?.id !== chainId) {
      const eventTitle = `[Bridge User Action] Sets new fromChainId`
      const eventData = {
        previousFromChainId: selectedItem?.id,
        newFromChainId: chainId,
      }

      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(chainId))
    }
  }

  const itemList = itemListFunction(searchStr)

  return (
    <SelectorWrapper
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      setSearchStr={setSearchStr}
    >
      {Object.entries(itemList).map(([key, value]: [string, Chain[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key}>
            {value.map((chain) => (
              <SelectSpecificNetworkButton
                dataId={dataTestId}
                key={chain.id}
                itemChainId={chain.id}
                isOrigin={isOrigin}
                isCurrentChain={selectedItem?.id === chain.id}
                active={false}
                onClick={() => handleSetChainId(chain.id)}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
  )
}
