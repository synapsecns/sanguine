import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { CHAINS_BY_ID } from '@/constants/chains'
import { getTokenFuseOptions } from '@/constants/fuseOptions'

import type { Token } from '@/utils/types'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { sortByPriorityRank } from '@/utils/helpers/sortByPriorityRank'

import { setShowSwapToTokenListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapToToken } from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'

import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'
import { SelectTokenButton } from '@/components/bridgeSwap/SelectTokenButton'

export const SwapToTokenListOverlay = () => {
  const { swapChainId, swapToTokens, swapToToken } = useSwapState()

  const dispatch = useDispatch()

  let possibleTokens = sortByPriorityRank(swapToTokens)

  const { toTokens: allToChainTokens } = getSwapPossibilities({
    fromChainId: swapChainId,
    fromToken: null,
    toChainId: swapChainId,
    toToken: null,
  })

  let remainingChainTokens = swapChainId
    ? sortByPriorityRank(_.difference(allToChainTokens, swapToTokens))
    : []

  const { toTokens: allTokens } = getSwapPossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherToTokens = swapChainId
    ? sortByPriorityRank(_.difference(allTokens, allToChainTokens))
    : sortByPriorityRank(allTokens)

  const possibleTokenswithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))

  const remainingChainTokensWithSource = remainingChainTokens.map((token) => ({
    ...token,
    source: 'remainingChainTokens',
  }))

  const allOtherToTokensWithSource = allOtherToTokens.map((token) => ({
    ...token,
    source: 'allOtherToTokens',
  }))

  const masterList = [
    ...possibleTokenswithSource,
    ...remainingChainTokensWithSource,
    ...allOtherToTokensWithSource,
  ]

  function onCloseOverlay() {
    dispatch(setShowSwapToTokenListOverlay(false))
  }

  const {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose,
  } = useOverlaySearch(masterList.length, onCloseOverlay)


  const fuse = new Fuse(masterList, getTokenFuseOptions(swapChainId))

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingChainTokens = results.filter(
      (item) => item.source === 'remainingChainTokens'
    )
    allOtherToTokens = results.filter(
      (item) => item.source === 'allOtherToTokens'
    )
  }

  const handleSetToToken = (oldToken: Token, newToken: Token) => {
    const eventTitle = `[Swap User Action] Sets new toToken`
    const eventData = {
      previousToToken: oldToken?.symbol,
      newToToken: newToken?.symbol,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setSwapToToken(newToken))
    onClose()
  }

  return (
    <SearchOverlayContent
      overlayRef={overlayRef}
      searchStr={searchStr}
      onSearch={onSearch}
      onClose={onClose}
      type="token"
    >
      {possibleTokens?.length > 0 && (
        <SearchResultsContainer label="Receive…">
          {possibleTokens.map((token, idx) =>
            <SelectTokenButton
              isOrigin={false}
              key={idx}
              token={token}
              chainId={swapChainId}
              selectedToken={swapToToken}
              active={idx === currentIdx}
              showAllChains={false}
              onClick={() => {
                if (token === swapToToken) {
                  onClose()
                } else {
                  handleSetToToken(swapToToken, token)
                }
              }}
            />
          )}
        </SearchResultsContainer>
      )}
      {remainingChainTokens?.length > 0 && (
        <SearchResultsContainer
          label={
            swapChainId
              ? `More on ${CHAINS_BY_ID[swapChainId]?.name}`
              : 'All swapable tokens'
          }
        >
          {remainingChainTokens.map((token, idx) =>
            <SelectTokenButton
              isOrigin={false}
              key={idx}
              token={token}
              chainId={swapChainId}
              selectedToken={swapToToken}
              active={idx + possibleTokens.length === currentIdx}
              showAllChains={false}
              onClick={() => handleSetToToken(swapToToken, token)}
            />
          )}
        </SearchResultsContainer>
      )}
      {allOtherToTokens?.length > 0 && (
        <SearchResultsContainer label="All swapable tokens">
          {allOtherToTokens.map((token, idx) =>
            <SelectTokenButton
              isOrigin={false}
              key={idx}
              token={token}
              chainId={swapChainId}
              selectedToken={swapToToken}
              active={
                idx +
                  possibleTokens.length +
                  remainingChainTokens.length ===
                currentIdx
              }
              showAllChains={true}
              onClick={() => handleSetToToken(swapToToken, token)}
              alternateBackground={true}
            />
          )}
        </SearchResultsContainer>
      )}
    </SearchOverlayContent>
  )
}
