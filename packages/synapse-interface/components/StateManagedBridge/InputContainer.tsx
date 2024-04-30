import _, { isNumber } from 'lodash'
import toast from 'react-hot-toast'
import React, { useEffect, useState, useCallback, useMemo } from 'react'
import { useAccount } from 'wagmi'
import { zeroAddress } from 'viem'
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
import { formatBigIntToString } from '@/utils/bigint/format'
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
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { useFromTokenListArray } from './hooks/useFromTokenListArray'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { AvailableBalance } from './AvailableBalance'
import { useGasEstimator } from '../../utils/hooks/useGasEstimator'
import { getParsedBalance } from '@/utils/getParsedBalance'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useAppDispatch()
  const { chain, isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, fromToken, fromValue } = useBridgeState()
  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

  const { parsedGasCost, maxBridgeableGas, isLoading, gasFeeExceedsBalance } =
    useGasEstimator()

  const { addresses, decimals } = fromToken || {}

  const tokenAddress = addresses?.[fromChainId]
  const tokenDecimals = _.isNumber(decimals)
    ? decimals
    : decimals?.[fromChainId]
  const isGasToken: boolean = tokenAddress === zeroAddress

  const balance: bigint = balances[fromChainId]?.find(
    (token) => token.tokenAddress === addresses?.[fromChainId]
  )?.balance
  const parsedBalance = getParsedBalance(balance, tokenDecimals, 4)

  const maxBalance = formatBigIntToString(balance, tokenDecimals)
  const maxBalanceBridgeable = isNumber(maxBridgeableGas)
    ? maxBridgeableGas?.toString()
    : maxBalance

  const onMaxBalance = useCallback(() => {
    if (gasFeeExceedsBalance) {
      toast.error('Gas fees likely exceeds your balance.', {
        id: 'toast-error-not-enough-gas',
        duration: 10000,
      })
      dispatch(updateFromValue('0.0'))
    } else {
      dispatch(updateFromValue(maxBalanceBridgeable))
    }
  }, [fromChainId, fromToken, maxBalanceBridgeable, gasFeeExceedsBalance])

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
        <div className="mr-auto">
          <AmountInput
            inputRef={inputRef}
            showValue={showValue}
            handleFromValueChange={handleFromValueChange}
          />
          <AvailableBalance
            fromChainId={fromChainId}
            fromToken={fromToken}
            balance={balance}
            parsedBalance={parsedBalance}
            onMaxBalance={onMaxBalance}
            isGasEstimateLoading={isLoading}
            disabled={!isConnected}
          />
        </div>
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
