import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount, useNetwork } from 'wagmi'

import MiniMaxButton from '../buttons/MiniMaxButton'
import { AmountInput, TokenSelector } from '../ui/BridgeCardComponents'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { ChainSelector } from '../ui/BridgeCardComponents'
import { SwapChainListArray } from './SwapChainListOverlay'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import {
  initialState,
  setSwapChainId,
  setSwapFromToken,
  updateSwapFromValue,
} from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'
import {
  BridgeAmountContainer,
  BridgeSectionContainer,
} from '../ui/BridgeCardComponents'
import { CHAINS_BY_ID } from '@/constants/chains'
import { FromChainListArray } from '../StateManagedBridge/FromChainListOverlay'
import { SwapFromTokenListArray } from './SwapFromTokenListOverlay'

export const SwapInputContainer = () => {
  const inputRef = useRef<HTMLInputElement>(null)
  const { swapChainId, swapFromToken, swapToToken, swapFromValue } =
    useSwapState()
  const [showValue, setShowValue] = useState('')

  const [hasMounted, setHasMounted] = useState(false)

  const { balances } = usePortfolioState()

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const { isConnected } = useAccount()
  const { chain } = useNetwork()

  const dispatch = useDispatch()

  const tokenData = balances[swapChainId]?.find(
    (token) => token.tokenAddress === swapFromToken?.addresses[swapChainId]
  )

  const parsedBalance = tokenData?.parsedBalance

  const balance = tokenData?.balance

  useEffect(() => {
    if (
      swapFromToken &&
      swapFromToken.decimals[swapChainId] &&
      stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]) !== 0n
    ) {
      setShowValue(swapFromValue)
    }

    if (swapFromValue === initialState.swapFromValue) {
      setShowValue(initialState.swapFromValue)
    }
  }, [swapFromValue, swapChainId, swapFromToken])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const swapFromValueString: string = cleanNumberInput(event.target.value)
    try {
      dispatch(updateSwapFromValue(swapFromValueString))
      setShowValue(swapFromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        dispatch(updateSwapFromValue(''))
        setShowValue(inputValue)
      }
    }
  }

  const onMaxBalance = useCallback(() => {
    dispatch(
      updateSwapFromValue(
        formatBigIntToString(balance, swapFromToken.decimals[swapChainId])
      )
    )
  }, [balance, swapChainId, swapFromToken])

  const connectedStatus = useMemo(() => {
    if (hasMounted && isConnected) {
      if (swapChainId === chain.id) {
        return <ConnectedIndicator />
      } else if (swapChainId !== chain.id) {
        return <ConnectToNetworkButton chainId={swapChainId} />
      }
    } else if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    }
  }, [chain, swapChainId, isConnected, hasMounted])

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <SwapChainSelector />
        {connectedStatus}
      </div>
      <BridgeAmountContainer>
        <SwapFromTokenSelector />
        <AmountInput
          inputRef={inputRef}
          hasMounted={hasMounted}
          isConnected={isConnected}
          showValue={showValue}
          handleFromValueChange={handleFromValueChange}
          parsedBalance={parsedBalance}
          onMaxBalance={onMaxBalance}
        />

        {hasMounted && isConnected && (
          <MiniMaxButton
            disabled={!balance || balance === 0n ? true : false}
            onClickBalance={onMaxBalance}
          />
        )}
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const SwapChainSelector = () => (
  <ChainSelector
    dataTestId="swap-origin-chain"
    isOrigin={true}
    selectedItem={CHAINS_BY_ID[useSwapState().swapChainId]}
    label="From"
    itemListFunction={SwapChainListArray}
    setFunction={setSwapChainId}
  />
)

const SwapFromTokenSelector = () => (
  <TokenSelector
    dataTestId="swap-origin-token"
    selectedItem={useSwapState().swapFromToken}
    isOrigin={true}
    placeholder="In"
    itemListFunction={SwapFromTokenListArray}
    setFunction={setSwapFromToken}
  />
)