import React, { useState } from 'react'
import { useDispatch } from 'react-redux'

import { type Token } from '@/utils/types'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { SelectSpecificTokenButton } from '@/components/ui/SelectSpecificTokenButton'
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
  action,
}: TokenSelectorTypes) {
  const [searchStr, setSearchStr] = useState('')

  const dispatch = useDispatch()

  const handleSetFromToken = (token) => {
    if (selectedItem !== token) {
      const eventTitle = `[${action} User Action] Sets new ${
        isOrigin ? 'from' : 'to'
      }Token`

      const eventData = isOrigin
        ? {
            previousFromToken: selectedItem?.symbol,
            newFromToken: token?.symbol,
          }
        : {
            previousToToken: selectedItem?.symbol,
            newToToken: token?.symbol,
          }
      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(token))
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
      {Object.entries(itemList).map(([key, value]: [string, Token[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key}>
            {value.map((token, idx) => (
              <SelectSpecificTokenButton
                isOrigin={isOrigin}
                key={idx}
                token={token}
                selectedToken={selectedItem}
                active={false}
                showAllChains={key === 'All other tokens'}
                onClick={() => handleSetFromToken(token)}
                alternateBackground={false}
                action={action}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
  )
}
