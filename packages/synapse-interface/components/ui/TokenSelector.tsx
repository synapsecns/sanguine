import React, { useEffect, useMemo, useState } from 'react'
import { useDispatch } from 'react-redux'

import { type Token } from '@/utils/types'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { SelectSpecificTokenButton } from '@/components/ui/SelectSpecificTokenButton'
import { TokenSelectorTypes } from '@/components/ui/types'
import { SelectorWrapper } from '@/components/ui/SelectorWrapper'
import { ListSectionWrapper } from '@/components/ui/ListSectionWrapper'
import { useKeyPress } from '@/utils/hooks/useKeyPress'

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
  const [hover, setHover] = useState(false)

  const [currentRouteSymbol, setCurrentRouteSymbol] = useState(null)
  const dispatch = useDispatch()

  const handleSetToken = (token) => {
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
  const flatItemList = Object.entries(itemList).reduce(
    (acc, [_, value]) => [...acc, ...(value as Token[])],
    []
  )

  const onClose = () => {
    setSearchStr('')
    setCurrentRouteSymbol(null)
    setHover(false)
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentRouteSymbol(null)
  }

  const arrowUp = useKeyPress('ArrowUp', hover)
  const arrowDown = useKeyPress('ArrowDown', hover)
  const enterPress = useKeyPress('Enter', hover)

  const arrowDownFunc = () => {
    const currentIndex = flatItemList.findIndex(
      (item) => item.routeSymbol === currentRouteSymbol
    )
    const nextIndex = currentIndex + 1
    if (arrowDown && nextIndex < flatItemList.length) {
      setCurrentRouteSymbol(flatItemList[nextIndex].routeSymbol)
    }
  }

  const arrowUpFunc = () => {
    const currentIndex = flatItemList.findIndex(
      (item) => item.routeSymbol === currentRouteSymbol
    )
    const prevIndex = currentIndex - 1
    if (arrowUp && prevIndex >= 0) {
      setCurrentRouteSymbol(flatItemList[prevIndex].routeSymbol)
    }
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(arrowUpFunc, [arrowUp])

  useEffect(() => {
    if (currentRouteSymbol !== null) {
      const token = flatItemList.find(
        (item) => item.routeSymbol === currentRouteSymbol
      )
      handleSetToken(token)
      onClose()
    }
  }, [enterPress])

  return (
    <SelectorWrapper
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      onSearch={onSearch}
      onClose={onClose}
      hover={hover}
      setHover={setHover}
    >
      {Object.entries(itemList).map(([key, value]: [string, Token[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key} key={key}>
            {value.map((token) => (
              <SelectSpecificTokenButton
                isOrigin={isOrigin}
                key={token.routeSymbol}
                token={token}
                showAllChains={key === 'All other tokens'}
                action={action}
                isSelected={selectedItem?.routeSymbol === token.routeSymbol}
                isActive={token.routeSymbol === currentRouteSymbol}
                onClick={() => handleSetToken(token)}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
  )
}
