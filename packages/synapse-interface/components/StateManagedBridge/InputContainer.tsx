import { isNull } from 'lodash'
import { useAppDispatch } from '@/store/hooks'
import { zeroAddress, Address, parseGwei } from 'viem'
import {
  initialState,
  updateFromValue,
  setFromChainId,
  setFromToken,
} from '@/slices/bridge/reducer'
import React, { useEffect, useState, useCallback, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
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
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { isEmpty } from 'lodash'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { useFromTokenListArray } from './hooks/useFromTokenListArray'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { joinClassNames } from '@/utils/joinClassNames'
import { Token } from '@/utils/types'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { stringToBigInt } from '@/utils/bigint/format'
import { getPublicClient } from '@wagmi/core'
import { AvailableBalance } from './AvailableBalance'
import { estimateGas } from '@wagmi/core'
import { wagmiConfig } from '@/wagmiConfig'
import { useGasEstimator } from '../../utils/hooks/useGasEstimator'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useDispatch()
  const { address, chain, isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken, fromValue } =
    useBridgeState()
  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

  const { rawGasCost, parsedGasCost, maxBridgeableGas } = useGasEstimator()

  const isGasToken: boolean = fromToken?.addresses[fromChainId] === zeroAddress

  const selectedFromToken: TokenAndBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )

  const { balance, parsedBalance } = selectedFromToken || {}

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
    if (fromToken && fromToken?.decimals[fromChainId]) {
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

  const onMaxBridgeableBalance = useCallback(() => {
    if (maxBridgeableGas) {
      if (maxBridgeableGas < 0) {
        dispatch(
          updateFromValue(
            formatBigIntToString(0n, fromToken?.decimals[fromChainId])
          )
        )
      } else {
        dispatch(updateFromValue(maxBridgeableGas.toString()))
      }
    } else {
      dispatch(
        updateFromValue(
          formatBigIntToString(balance, fromToken?.decimals[fromChainId])
        )
      )
    }
  }, [
    fromChainId,
    fromToken,
    isGasToken,
    parsedGasCost,
    balance,
    parsedBalance,
  ])

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <FromChainSelector />
        {connectedStatus}
      </div>
      <BridgeAmountContainer>
        <FromTokenSelector />
        <div>
          <AmountInput
            inputRef={inputRef}
            showValue={showValue}
            handleFromValueChange={handleFromValueChange}
          />
          <AvailableBalance
            fromChainId={fromChainId}
            fromValue={fromValue}
            fromToken={fromToken}
            balance={balance}
            parsedBalance={parsedBalance}
            maxBridgeableBalance={
              maxBridgeableGas ? maxBridgeableGas.toString() : parsedBalance
            }
            isGasToken={isGasToken}
            parsedGasCost={parsedGasCost}
            onMaxBalance={onMaxBridgeableBalance}
            isConnected={isConnected}
            hasMounted={hasMounted}
          />
        </div>
        {hasMounted && isConnected && (
          <MiniMaxButton
            disabled={!balance || balance === 0n ? true : false}
            onClickBalance={onMaxBridgeableBalance}
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
