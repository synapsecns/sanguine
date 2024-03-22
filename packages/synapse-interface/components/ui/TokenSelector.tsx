import React, { useState } from 'react'
import { type Token } from '@/utils/types'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useDispatch } from 'react-redux'
import SelectSpecificTokenButton from '../StateManagedBridge/components/SelectSpecificTokenButton'

import { TokenSelectorTypes } from '@/components/ui/types'
import { SelectorWrapper } from '@/components/ui/SelectorWrapper'
import { ListSectionWrapper } from '@/components/ui/ListSectionWrapper'

export function TokenSelector({
  dataTestId,
  selectedItem,
  label,
  placeholder,
  itemListFunction,
  setFunction,
  isOrigin,
}: TokenSelectorTypes) {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')

  const dispatch = useDispatch()

  const handleSetFromToken = (token) => {
    if (selectedItem !== token) {
      const eventTitle = `[Bridge User Action] Sets new fromChainId`
      const eventData = {
        previousFromToken: selectedItem?.symbol,
        newFromToken: token?.symbol,
      }
      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(token))
    }
  }

  const itemList = itemListFunction(searchStr) // TODO: Use result instead of variable in context?

  return (
    <SelectorWrapper
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      setSearchStr={setSearchStr}
    >
      {Object.entries(itemList).map(([key, value]: [string, Token[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key}>
            {value.map((token, idx) => (
              <SelectSpecificTokenButton
                isOrigin={isOrigin}
                key={idx}
                token={token}
                selectedToken={selectedItem}
                // active={
                //   idx +
                //     possibleTokens.length +
                //     remainingChainTokens.length ===
                //   currentIdx
                // }
                active={false}
                // showAllChains={true}
                onClick={() => handleSetFromToken(token)}
                alternateBackground={false}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
  )
}
