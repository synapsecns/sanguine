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

  const [currentIdx, setCurrentIdx] = useState(-1)
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

  console.log(
    `flatItemList`,
    flatItemList.map((t, i) => `${i}: ${t.routeSymbol}`)
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
    console.log(
      `flatItemList[currentIdx]`,
      flatItemList[currentIdx]?.routeSymbol
    )
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < flatItemList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    console.log(`in isOrigin: ${isOrigin}, up fn`)
    console.log(
      `flatItemList[currentIdx]`,
      flatItemList[currentIdx]?.routeSymbol
    )
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
      console.log(
        `in enter flatItemList[currentIdx]`,
        flatItemList[currentIdx]?.routeSymbol
      )
      handleSetToken(flatItemList[currentIdx])
    }
    onClose()
  }, [enterPress])

  return (
    <SelectorWrapper
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      setSearchStr={setSearchStr}
      onClose={onClose}
      hover={hover}
      setHover={setHover}
    >
      {Object.entries(itemList).map(
        ([key, value]: [string, Token[]], index) => {
          return value.length ? (
            <ListSectionWrapper sectionKey={key}>
              {value.map((token, tokenIndex) => (
                <SelectSpecificTokenButton
                  isOrigin={isOrigin}
                  key={token.routeSymbol}
                  token={token}
                  selectedToken={selectedItem}
                  active={false}
                  showAllChains={key === 'All other tokens'}
                  onClick={() => handleSetToken(token)}
                  alternateBackground={false}
                  action={action}
                  isCurrentToken={currentIdx === index + tokenIndex}
                />
              ))}
            </ListSectionWrapper>
          ) : null
        }
      )}
    </SelectorWrapper>
  )
}
