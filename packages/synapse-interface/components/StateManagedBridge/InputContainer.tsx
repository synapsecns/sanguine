import { debounce, isNull } from 'lodash'
import toast from 'react-hot-toast'
import React, { useEffect, useState, useCallback, useMemo } from 'react'
import { useAccount } from 'wagmi'

import { useAppDispatch } from '@/store/hooks'
import { updateDebouncedFromValue } from '@/slices/bridge/reducer'
import { AmountInput } from '@/components/ui/AmountInput'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { AvailableBalance } from './AvailableBalance'
import { useGasEstimator } from '@/utils/hooks/useGasEstimator'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { formatAmount } from '@/utils/formatAmount'
import { useWalletState } from '@/slices/wallet/hooks'
import { FromChainSelector } from '@/components/StateManagedBridge/FromChainSelector'
import { FromTokenSelector } from '@/components/StateManagedBridge/FromTokenSelector'
import { useBridgeSelections } from './hooks/useBridgeSelections'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { useUsdDisplay } from '@hooks/useUsdDisplay'

export const inputRef = React.createRef<HTMLInputElement>()

interface InputContainerProps {
  setIsTyping: React.Dispatch<React.SetStateAction<boolean>>
}

export const InputContainer: React.FC<InputContainerProps> = ({
  setIsTyping,
}) => {
  const dispatch = useAppDispatch()
  const { chain, isConnected } = useAccount()
  const { isWalletPending } = useWalletState()
  const { fromChainId, fromToken, debouncedFromValue } = useBridgeState()
  const [localInputValue, setLocalInputValue] = useState(debouncedFromValue)

  const { hasValidFromSelections, hasValidSelections, onSelectedChain } =
    useBridgeValidations()
  const { fromTokenBalance, fromTokenDecimals } = useBridgeSelections()

  const parsedBalance = getParsedBalance(fromTokenBalance, fromTokenDecimals)
  const formattedBalance = formatAmount(parsedBalance)

  const {
    isLoading,
    isGasToken,
    parsedGasCost,
    maxBridgeableGas,
    hasValidGasEstimateInputs,
    estimateBridgeableBalanceCallback,
  } = useGasEstimator()

  // Fetch token price and calculate USD value
  const usdValue = useUsdDisplay(fromToken, localInputValue, fromChainId)

  const isInputMax =
    maxBridgeableGas?.toString() === debouncedFromValue ||
    parsedBalance === debouncedFromValue

  const debouncedUpdateFromValue = useMemo(
    () =>
      debounce(
        (value: string) => dispatch(updateDebouncedFromValue(value)),
        400
      ),
    [dispatch]
  )

  useEffect(() => {
    return () => {
      debouncedUpdateFromValue.cancel()
    }
  }, [debouncedUpdateFromValue])

  const handleFromValueChange = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      const cleanedValue = cleanNumberInput(event.target.value)
      try {
        setLocalInputValue(cleanedValue)
        debouncedUpdateFromValue(cleanedValue)
      } catch (error) {
        console.log('Invalid value for conversion to BigInteger')
        const inputValue = event.target.value
        const regex = /^[0-9]*[.,]?[0-9]*$/

        if (regex.test(inputValue) || inputValue === '') {
          setLocalInputValue(cleanedValue)
          debouncedUpdateFromValue(cleanedValue)
        }
      }
    },
    [debouncedUpdateFromValue]
  )

  const onMaxBalance = useCallback(async () => {
    if (hasValidGasEstimateInputs()) {
      const bridgeableBalance = await estimateBridgeableBalanceCallback()

      if (isNull(bridgeableBalance)) {
        setLocalInputValue(parsedBalance)
        dispatch(updateDebouncedFromValue(parsedBalance))
      } else if (bridgeableBalance > 0) {
        const bridgeableBalanceString = bridgeableBalance.toString()
        setLocalInputValue(bridgeableBalanceString)
        dispatch(updateDebouncedFromValue(bridgeableBalanceString))
      } else {
        setLocalInputValue('0.0')
        dispatch(updateDebouncedFromValue('0.0'))
        toast.error('Gas fees likely exceeds your balance.', {
          id: 'toast-error-not-enough-gas',
          duration: 10000,
        })
      }
    } else {
      setLocalInputValue(parsedBalance)
      dispatch(updateDebouncedFromValue(parsedBalance))
    }
  }, [
    dispatch,
    fromChainId,
    fromToken,
    parsedBalance,
    hasValidGasEstimateInputs,
    estimateBridgeableBalanceCallback,
  ])

  useEffect(() => {
    setLocalInputValue(debouncedFromValue)
  }, [debouncedFromValue])

  const connectedStatus = useMemo(() => {
    if (!isConnected) {
      return <ConnectWalletButton />
    } else if (isConnected && onSelectedChain) {
      return <ConnectedIndicator />
    } else if (isConnected && !onSelectedChain) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }, [chain, fromChainId, isConnected])

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <FromChainSelector />
        {connectedStatus}
      </div>
      <BridgeAmountContainer>
        <FromTokenSelector />
        <div className="flex flex-col w-full gap-1">
          <AmountInput
            setIsTyping={setIsTyping}
            inputRef={inputRef}
            showValue={localInputValue}
            handleFromValueChange={handleFromValueChange}
            disabled={isWalletPending}
          />
          <div className="flex justify-between items-center">
            <div className="text-xs text-zinc-500 dark:text-zinc-400">
              {usdValue}
            </div>
            <AvailableBalance
              balance={formattedBalance}
              maxBridgeableBalance={maxBridgeableGas}
              gasCost={parsedGasCost}
              isGasToken={isGasToken}
              isGasEstimateLoading={isLoading}
              isDisabled={!isConnected || !hasValidFromSelections}
              onClick={
                !isConnected ||
                !hasValidSelections ||
                isLoading ||
                isInputMax ||
                isWalletPending
                  ? undefined
                  : onMaxBalance
              }
            />
          </div>
        </div>
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}
