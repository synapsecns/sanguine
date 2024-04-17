import { isNull } from 'lodash'
import { useAppSelector, useAppDispatch } from '@/store/hooks'
import { zeroAddress, Address } from 'viem'
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
import { calculateGasCost } from '../../utils/calculateGasCost'
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
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { AvailableBalance } from './AvailableBalance'
import { wagmiConfig } from '@/wagmiConfig'

export const inputRef = React.createRef<HTMLInputElement>()

const getBridgeQuote = async (
  synapseSDK: any,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  toToken: Token,
  amount: string
) => {
  try {
    return await synapseSDK.bridgeQuote(
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId],
      toToken.addresses[toChainId],
      stringToBigInt(amount, fromToken?.decimals[fromChainId])
    )
  } catch (error) {
    console.error('getBridgeQuote: ', error)
    return null
  }
}

const calculateEstimatedBridgeGasLimit = async (
  synapseSDK: any,
  bridgeQuote: any | null,
  address: string,
  toAddress: string,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  amount: string
) => {
  if (!bridgeQuote) return null

  const data = await synapseSDK.bridge(
    toAddress,
    bridgeQuote.routerAddress,
    fromChainId,
    toChainId,
    fromToken?.addresses[fromChainId as keyof Token['addresses']],
    stringToBigInt(amount, fromToken?.decimals[fromChainId]),
    bridgeQuote.originQuery,
    bridgeQuote.destQuery
  )

  const payload =
    fromToken?.addresses[fromChainId as keyof Token['addresses']] ===
      zeroAddress ||
    fromToken?.addresses[fromChainId as keyof Token['addresses']] === ''
      ? {
          data: data.data,
          to: data.to,
          value: stringToBigInt(amount, fromToken?.decimals[fromChainId]),
        }
      : data

  const publicClient = getPublicClient(wagmiConfig)

  const gasEstimate = await publicClient.estimateGas({
    value: payload.value,
    to: payload.to,
    account: address as Address,
    data: payload.data,
    chainId: fromChainId,
  })

  return gasEstimate
}

const useGasEstimator = () => {
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken, fromValue } =
    useBridgeState()
  const { gasData } = useAppSelector((state) => state.gasData)
  const { gasPrice, maxFeePerGas } = gasData?.formatted

  const [estimatedGasLimit, setEstimatedGasLimit] = useState<bigint>(0n)

  const { rawGasCost, parsedGasCost } = calculateGasCost(
    gasPrice,
    estimatedGasLimit.toString()
  )

  const isGasToken: boolean = fromToken?.addresses[fromChainId] === zeroAddress

  const selectedFromToken: TokenAndBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )

  const { balance, parsedBalance } = selectedFromToken || {}

  const hasRequiredGasEstimateInputs = (): boolean => {
    return Boolean(
      fromChainId &&
        toChainId &&
        fromToken &&
        toToken &&
        isGasToken &&
        parsedBalance
    )
  }

  useEffect(() => {
    if (hasRequiredGasEstimateInputs()) {
      ;(async () => {
        const bridgeQuote = await getBridgeQuote(
          synapseSDK,
          fromChainId,
          toChainId,
          fromToken,
          toToken,
          parsedBalance
        )

        const gasLimit = await calculateEstimatedBridgeGasLimit(
          synapseSDK,
          bridgeQuote,
          address,
          address,
          fromChainId,
          toChainId,
          fromToken,
          parsedBalance
        )

        setEstimatedGasLimit(gasLimit ?? 0n)

        console.log('gasLimit: ', gasLimit)
      })()
    }
  }, [fromChainId, toChainId, isGasToken])

  return { rawGasCost, parsedGasCost }
}

export const InputContainer = () => {
  const dispatch = useDispatch()
  const { address, chain, isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken, fromValue } =
    useBridgeState()
  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

  const [estimatedGasLimit, setEstimatedGasLimit] = useState<bigint>(0n)

  const { gasData } = useAppSelector((state) => state.gasData)
  const { gasPrice, maxFeePerGas } = gasData?.formatted
  const { rawGasCost, parsedGasCost } = calculateGasCost(
    gasPrice,
    estimatedGasLimit.toString()
  )

  const isGasToken: boolean = fromToken?.addresses[fromChainId] === zeroAddress

  const selectedFromToken: TokenAndBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )

  const { balance, parsedBalance } = selectedFromToken || {}

  /** Fetch gasLimit using Wallet's gas balance */

  const { synapseSDK } = useSynapseContext()
  useEffect(() => {
    if (
      fromChainId &&
      toChainId &&
      fromToken &&
      toToken &&
      address &&
      isGasToken &&
      parsedBalance
    ) {
      ;(async () => {
        const bridgeQuote = await getBridgeQuote(
          synapseSDK,
          fromChainId,
          toChainId,
          fromToken,
          toToken,
          parsedBalance
        )

        const gasLimit = await calculateEstimatedBridgeGasLimit(
          synapseSDK,
          bridgeQuote,
          address,
          address,
          fromChainId,
          toChainId,
          fromToken,
          parsedBalance
        )

        setEstimatedGasLimit(gasLimit ?? 0n)

        console.log('gasLimit: ', gasLimit)
      })()
    }
  }, [fromChainId, toChainId, fromToken, toToken, address, isGasToken])

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

  const calculateMaxBridgeableGas = (
    parsedGasBalance: number,
    parsedGasCost: number
  ): number => {
    const maxBridgeable = parsedGasBalance - parsedGasCost
    return maxBridgeable
  }

  const onMaxBalance = () => {
    dispatch(
      updateFromValue(
        formatBigIntToString(balance, fromToken?.decimals[fromChainId])
      )
    )
  }

  const maxBridgeableGas: number | null =
    isGasToken && parsedGasCost
      ? calculateMaxBridgeableGas(
          parseFloat(parsedBalance),
          parseFloat(parsedGasCost)
        )
      : null

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

// TODO: Replace with HoverTooltip in Portfolio once other branch is merged in
export const HoverTooltip = ({ children, hoverContent, isActive }) => {
  const [showTooltip, setShowTooltip] = useState(false)

  const activateTooltip = () => setShowTooltip(true)
  const hideTooltip = () => setShowTooltip(false)

  if (!isActive) {
    return <div>{children}</div>
  } else {
    return (
      <div
        onMouseEnter={activateTooltip}
        onMouseLeave={hideTooltip}
        className="relative w-fit"
      >
        {children}
        <Tooltip isHovered={showTooltip}>{hoverContent}</Tooltip>
      </div>
    )
  }
}

const Tooltip = ({
  isHovered,
  children,
}: {
  isHovered: boolean
  children: React.ReactNode
}) => {
  if (isHovered) {
    return (
      <div
        className={`
           absolute left-1/2 bottom-full translate-x-[-50%]
           z-50 hover-content px-2 py-1 text-white mb-1
           border border-solid border-[#252537]
           bg-[#101018] rounded-md text-left text-sm
         `}
      >
        <data>{children}</data>
      </div>
    )
  }
}
