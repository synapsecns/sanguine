import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { ChainSelector } from '@/components/ui/ChainSelector'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import {
  initialState,
  setSwapChainId,
  setSwapFromToken,
  updateSwapFromValue,
} from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useSwapChainListArray } from '@/components/StateManagedSwap//hooks/useSwapChainListArray'
import { useSwapFromTokenListArray } from '@/components/StateManagedSwap/hooks/useSwapFromTokenListOverlay'
import { AmountInput } from '@/components/ui/AmountInput'
import { joinClassNames } from '@/utils/joinClassNames'
import { MaxButton } from '../StateManagedBridge/MaxButton'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { formatAmount } from '@/utils/formatAmount'
import { getParsedBalance } from '@/utils/getParsedBalance'

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
  const { chain } = useAccount()

  const dispatch = useDispatch()

  const tokenData = balances[swapChainId]?.find(
    (token) => token.tokenAddress === swapFromToken?.addresses[swapChainId]
  )

  const balance = tokenData?.balance
  const decimals = tokenData?.token?.decimals[swapChainId]
  const parsedBalance = getParsedBalance(balance, decimals)
  const formattedBalance = formatAmount(parsedBalance)

  const isInputMax = parsedBalance === swapFromValue

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
        trimTrailingZeroesAfterDecimal(
          formatBigIntToString(balance, swapFromToken.decimals[swapChainId])
        )
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

  const labelClassName = joinClassNames({
    space: 'block',
    textColor: 'text-xxs md:text-xs',
    cursor: 'cursor-default',
  })

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <SwapChainSelector />
        {connectedStatus}
      </div>
      <BridgeAmountContainer>
        <SwapFromTokenSelector />
        <div className="flex flex-col">
          <AmountInput
            inputRef={inputRef}
            showValue={showValue}
            handleFromValueChange={handleFromValueChange}
          />
          <div className="flex">
            {hasMounted && isConnected && (
              <label htmlFor="inputRow" className={labelClassName}>
                <span className="text-zinc-500 dark:text-zinc-400">
                  Available:{' '}
                </span>
                {formattedBalance ?? '0.0'}
              </label>
            )}
            <MaxButton
              onClick={onMaxBalance}
              isHidden={!isConnected || !balance || isInputMax}
            />
          </div>
        </div>
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const SwapChainSelector = () => {
  const { swapChainId } = useSwapState()

  return (
    <ChainSelector
      dataTestId="swap-origin-chain"
      isOrigin={true}
      selectedItem={CHAINS_BY_ID[swapChainId]}
      label="From"
      itemListFunction={useSwapChainListArray}
      setFunction={setSwapChainId}
      action="Swap"
    />
  )
}

const SwapFromTokenSelector = () => {
  const { swapFromToken } = useSwapState()

  return (
    <TokenSelector
      dataTestId="swap-origin-token"
      selectedItem={swapFromToken}
      isOrigin={true}
      placeholder="In"
      itemListFunction={useSwapFromTokenListArray}
      setFunction={setSwapFromToken}
      action="Swap"
    />
  )
}
