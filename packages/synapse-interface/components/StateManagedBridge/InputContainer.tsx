import { isNull, isNumber } from 'lodash'
import toast from 'react-hot-toast'
import React, { useEffect, useState, useCallback, useMemo } from 'react'
import { useAccount } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import {
  initialState,
  updateFromValue,
  setFromChainId,
  setFromToken,
} from '@/slices/bridge/reducer'
import { ChainSelector } from '@/components/ui/ChainSelector'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { AmountInput } from '@/components/ui/AmountInput'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useFromChainListArray } from './hooks/useFromChainListArray'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { useFromTokenListArray } from './hooks/useFromTokenListArray'
import { AvailableBalance } from './AvailableBalance'
import { useGasEstimator } from '../../utils/hooks/useGasEstimator'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { MaxButton } from './MaxButton'
import { formatAmount } from '../../utils/formatAmount'
import { useWalletState } from '@/slices/wallet/hooks'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useAppDispatch()
  const { chain, isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken, fromValue } =
    useBridgeState()
  const { isWalletPending } = useWalletState()
  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

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
    maxBridgeableGas?.toString() === fromValue || parsedBalance === fromValue

  const onMaxBalance = useCallback(async () => {
    if (hasValidGasEstimateInputs()) {
      const bridgeableBalance = await estimateBridgeableBalanceCallback()

      if (isNull(bridgeableBalance)) {
        dispatch(updateFromValue(parsedBalance))
      } else if (bridgeableBalance > 0) {
        dispatch(updateFromValue(bridgeableBalance?.toString()))
      } else {
        dispatch(updateFromValue('0.0'))
        toast.error('Gas fees likely exceeds your balance.', {
          id: 'toast-error-not-enough-gas',
          duration: 10000,
        })
      }
    } else {
      dispatch(updateFromValue(parsedBalance))
    }
  }, [
    fromChainId,
    fromToken,
    parsedBalance,
    hasValidGasEstimateInputs,
    estimateBridgeableBalanceCallback,
  ])

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const connectedStatus = useMemo(() => {
    if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    } else if (hasMounted && isConnected && fromChainId === chain?.id) {
      return <ConnectedIndicator />
    } else if (hasMounted && isConnected && fromChainId !== chain?.id) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }, [chain, fromChainId, isConnected, hasMounted])

  useEffect(() => {
    if (fromToken && tokenDecimals) {
      setShowValue(fromValue)
    }

    if (fromValue === initialState.fromValue) {
      setShowValue(initialState.fromValue)
    }
  }, [fromValue, inputRef, fromChainId, fromToken])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const fromValueString: string = cleanNumberInput(event.target.value)
    try {
      dispatch(updateFromValue(fromValueString))
      setShowValue(fromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        dispatch(updateFromValue(inputValue))
        setShowValue(inputValue)
      }
    }
  }

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
            showValue={showValue}
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
              isInputMax
            }
          />
        </div>
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const FromChainSelector = () => {
  const { fromChainId } = useBridgeState()

  return (
    <ChainSelector
      dataTestId="bridge-origin-chain"
      selectedItem={CHAINS_BY_ID[fromChainId]}
      isOrigin={true}
      label="From"
      itemListFunction={useFromChainListArray}
      setFunction={setFromChainId}
      action="Bridge"
    />
  )
}

const FromTokenSelector = () => {
  const { fromToken } = useBridgeState()

  return (
    <TokenSelector
      dataTestId="bridge-origin-token"
      selectedItem={fromToken}
      isOrigin={true}
      placeholder="Out"
      itemListFunction={useFromTokenListArray}
      setFunction={setFromToken}
      action="Bridge"
    />
  )
}
