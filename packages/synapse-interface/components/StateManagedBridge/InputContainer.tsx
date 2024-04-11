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
import { useAccount, useNetwork } from 'wagmi'
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

export const inputRef = React.createRef<HTMLInputElement>()

const getBridgeQuote = async (
  synapseSDK: any,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  toToken: Token,
  amount: string
) => {
  return await synapseSDK.bridgeQuote(
    fromChainId,
    toChainId,
    fromToken.addresses[fromChainId],
    toToken.addresses[toChainId],
    stringToBigInt(amount, fromToken?.decimals[fromChainId])
  )
}

const calculateEstimatedBridgeGasLimit = async (
  synapseSDK: any,
  bridgeQuote: any,
  address: string,
  toAddress: string,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  amount: string
) => {
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

  const publicClient = getPublicClient()

  const gasEstimate = await publicClient.estimateGas({
    value: payload.value,
    to: payload.to,
    account: address as Address,
    data: payload.data,
    chainId: fromChainId,
  })

  return gasEstimate
}

export const InputContainer = () => {
  const dispatch = useAppDispatch()
  const { chain } = useNetwork()
  const { address, isConnected } = useAccount()
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

        setEstimatedGasLimit(gasLimit)

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
    } else if (hasMounted && isConnected && fromChainId === chain.id) {
      return <ConnectedIndicator />
    } else if (hasMounted && isConnected && fromChainId !== chain.id) {
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

const AvailableBalance = ({
  fromChainId,
  fromValue,
  fromToken,
  balance,
  parsedBalance,
  isGasToken = false,
  parsedGasCost,
  onMaxBalance,
  hasMounted,
  isConnected,
  disabled = false,
}: {
  fromChainId: number | null
  fromValue: string
  fromToken: Token | null
  balance?: bigint
  parsedBalance?: string
  isGasToken?: boolean
  parsedGasCost?: string
  onMaxBalance?: () => void
  hasMounted: boolean
  isConnected: boolean
  disabled?: boolean
}) => {
  const parsedBalanceFull = formatBigIntToString(
    balance,
    fromToken?.decimals[fromChainId]
  )

  const isTraceBalance = (): boolean => {
    if (!balance || !parsedBalanceFull) return false
    if (balance && !hasOnlyZeroes(parsedBalanceFull)) return true
    return false
  }

  const isTraceInput = (): boolean => {
    if (!fromValue) return false
    const shortenedFromValue = parseFloat(fromValue).toFixed(4)
    if (Number(shortenedFromValue) === 0 && !hasOnlyZeroes(fromValue)) {
      return true
    } else {
      return false
    }
  }

  const isGasCostCoveredByInput = (): boolean => {
    if (!isGasToken) return true

    if (isGasToken && parsedGasCost && fromValue && parsedBalanceFull) {
      return (
        parseFloat(fromValue) <
        parseFloat(parsedBalanceFull) - parseFloat(parsedGasCost)
      )
    } else {
      return true
    }
  }

  const isGasCostCoveredByBalance = (): boolean => {
    if (!isGasToken) return true

    if (isGasToken && parsedGasCost && parsedBalanceFull) {
      return parseFloat(parsedGasCost) < parseFloat(parsedBalanceFull)
    } else {
      return false
    }
  }

  const showGasReserved = (): boolean => {
    return !hasOnlyZeroes(fromValue) && !isGasCostCoveredByInput()
  }

  const gasReserved = showGasReserved()
    ? isGasCostCoveredByBalance()
      ? parseFloat(fromValue) - parseFloat(parsedGasCost)
      : parseFloat(fromValue)
    : undefined

  let tooltipContent

  if (showGasReserved()) {
    tooltipContent = (
      <div className="space-y-2 whitespace-nowrap">
        <div>You may not have enough to cover gas fees.</div>
        <div>Estimated gas: {parseFloat(parsedGasCost).toFixed(4)}</div>
      </div>
    )
  } else if (!isGasCostCoveredByInput()) {
    tooltipContent = (
      <div className="whitespace-nowrap">
        You may not have enough to cover gas fees.
      </div>
    )
  } else if (!isGasCostCoveredByBalance()) {
    tooltipContent = (
      <div className="whitespace-nowrap">
        Gas fees may exceed your available balance.
      </div>
    )
  }

  const labelClassName = joinClassNames({
    space: 'block',
    textColor: `text-xxs md:text-xs ${
      showGasReserved() ? '!text-yellowText' : ''
    }`,
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  // console.log('showGasReserved:', showGasReserved())
  // console.log('isGasCostCoveredByInput: ', isGasCostCoveredByInput())
  // console.log('isGasCostCoveredByBalance:', isGasCostCoveredByBalance())

  if (showGasReserved()) {
    return (
      <HoverTooltip isActive={true} hoverContent={tooltipContent}>
        <label
          htmlFor="inputRow"
          onClick={onMaxBalance}
          className={labelClassName}
        >
          {isTraceInput() ? '<0.001' : gasReserved.toFixed(4)}
          <span> reserved for gas</span>
        </label>
      </HoverTooltip>
    )
  } else if (hasMounted && isConnected && !disabled) {
    return (
      <HoverTooltip
        isActive={!isGasCostCoveredByBalance() || !isGasCostCoveredByInput()}
        hoverContent={tooltipContent}
      >
        <label
          htmlFor="inputRow"
          onClick={onMaxBalance}
          className={labelClassName}
        >
          {isTraceBalance() ? '<0.001' : parsedBalance ?? '0.0'}
          <span className="text-zinc-500 dark:text-zinc-400"> available</span>
        </label>
      </HoverTooltip>
    )
  } else {
    return null
  }
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
