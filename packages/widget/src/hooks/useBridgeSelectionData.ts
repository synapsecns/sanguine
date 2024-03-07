import { useState, useEffect } from 'react'
import { BridgeSelections } from 'types'

import { CHAINS_ARRAY } from '@/constants/chains'
import { findTokenByRouteSymbol } from '@/utils/findTokenByRouteSymbol'

export const useBridgeSelectionData = (): BridgeSelections => {
  const [originChain, setOriginChain] = useState('')
  const [originToken, setOriginToken] = useState('')
  const [destinationChain, setDestinationChain] = useState('')
  const [destinationToken, setDestinationToken] = useState('')

  useEffect(() => {
    let originChainSelect = null
    let destinationChainSelect = null

    let originTokenSelect = null
    let destinationTokenSelect = null

    const getChainValues = () => {
      if (
        originChainSelect &&
        destinationChainSelect &&
        originTokenSelect &&
        destinationTokenSelect
      ) {
        const originChainName = originChainSelect.textContent?.trim() || ''
        const destinationChainName =
          destinationChainSelect.textContent?.trim() || ''

        const originTokenSymbol = originTokenSelect.textContent?.trim() || ''
        const destinationTokenSymbol =
          destinationTokenSelect.textContent?.trim() || ''

        setOriginChain(originChainName === 'Network' ? null : originChainName)
        setOriginToken(originTokenSymbol === 'Token' ? null : originTokenSymbol)
        setDestinationChain(
          destinationChainName === 'Network' ? null : destinationChainName
        )
        setDestinationToken(
          destinationTokenSymbol === 'Token' ? null : destinationTokenSymbol
        )
      }
    }

    const checkElements = () => {
      originChainSelect = document.getElementById('origin-chain-select')
      destinationChainSelect = document.getElementById(
        'destination-chain-select'
      )
      originTokenSelect = document.getElementById('origin-token-select')
      destinationTokenSelect = document.getElementById(
        'destination-token-select'
      )

      if (
        originChainSelect &&
        destinationChainSelect &&
        originTokenSelect &&
        destinationTokenSelect
      ) {
        // Get initial values when elements are available
        getChainValues()

        const observer = new MutationObserver(getChainValues)
        const config = { childList: true, characterData: true, subtree: true }

        observer.observe(originChainSelect, config)
        observer.observe(destinationChainSelect, config)
        observer.observe(originTokenSelect, config)
        observer.observe(destinationTokenSelect, config)

        return () => {
          observer.disconnect()
        }
      } else {
        // If elements are not available, check again after a short delay
        const timerId = setTimeout(checkElements, 100)
        return () => clearTimeout(timerId)
      }
    }

    checkElements()
  }, [])

  const originChainId = CHAINS_ARRAY.find(
    (chain) => chain.name === originChain
  )?.id
  const destinationChainId = CHAINS_ARRAY.find(
    (chain) => chain.name === destinationChain
  )?.id

  const originTokenAddress =
    findTokenByRouteSymbol(originToken)?.addresses[originChainId] ?? null
  const destinationTokenAddress =
    findTokenByRouteSymbol(destinationToken)?.addresses[destinationChainId] ??
    null

  return {
    originChain: {
      id: originChainId,
      name: originChain,
    },
    destinationChain: {
      id: destinationChainId,
      name: destinationChain,
    },
    originToken: {
      symbol: originToken,
      address: originTokenAddress,
    },
    destinationToken: {
      symbol: destinationToken,
      address: destinationTokenAddress,
    },
  }
}
