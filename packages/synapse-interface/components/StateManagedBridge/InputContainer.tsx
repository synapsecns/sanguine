import { debounce, isNull, isNumber } from 'lodash'
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
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { AvailableBalance } from './AvailableBalance'
import { useGasEstimator } from '../../utils/hooks/useGasEstimator'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { MaxButton } from './MaxButton'
import { formatAmount } from '../../utils/formatAmount'
import { useWalletState } from '@/slices/wallet/hooks'
import { FromChainSelector } from '@/components/StateManagedBridge/FromChainSelector'
import { FromTokenSelector } from '@/components/StateManagedBridge/FromTokenSelector'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useAppDispatch()
  const { chain, isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken, debouncedFromValue } =
    useBridgeState()
  const { isWalletPending } = useWalletState()
  const [localInputValue, setLocalInputValue] = useState(debouncedFromValue)

  const { addresses, decimals } = fromToken || {}
  const tokenDecimals = isNumber(decimals) ? decimals : decimals?.[fromChainId]
  const balance: bigint = balances[fromChainId]?.find(
    (token) => token.tokenAddress === addresses?.[fromChainId]
  )?.balance
  const parsedBalance = getParsedBalance(balance, tokenDecimals)
  const formattedBalance = formatAmount(parsedBalance)

  const hasValidFromSelections: boolean = useMemo(() => {
    return Boolean(fromChainId && fromToken)
  }, [fromChainId, fromToken])

  const hasValidInputSelections: boolean = useMemo(() => {
    return Boolean(fromChainId && fromToken && toChainId && toToken)
  }, [fromChainId, toChainId, fromToken, toToken])

  const {
    isLoading,
    isGasToken,
    parsedGasCost,
    maxBridgeableGas,
    hasValidGasEstimateInputs,
    estimateBridgeableBalanceCallback,
  } = useGasEstimator()

  const isInputMax =
    maxBridgeableGas?.toString() === debouncedFromValue ||
    parsedBalance === debouncedFromValue

  const debouncedUpdateFromValue = useMemo(
    () =>
      debounce(
        (value: string) => dispatch(updateDebouncedFromValue(value)),
        300
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
    } else if (isConnected && fromChainId === chain?.id) {
      return <ConnectedIndicator />
    } else if (isConnected && fromChainId !== chain?.id) {
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
        <div className="flex flex-wrap w-full">
          <AmountInput
            inputRef={inputRef}
            showValue={localInputValue}
            handleFromValueChange={handleFromValueChange}
            disabled={isWalletPending}
          />
          <AvailableBalance
            balance={formattedBalance}
            maxBridgeableBalance={maxBridgeableGas}
            gasCost={parsedGasCost}
            isGasToken={isGasToken}
            isGasEstimateLoading={isLoading}
            isDisabled={!isConnected || !hasValidFromSelections}
          />
          <MaxButton
            onClick={onMaxBalance}
            isHidden={
              !isConnected ||
              !hasValidInputSelections ||
              isLoading ||
              isInputMax ||
              isWalletPending
            }
          />
        </div>
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}
