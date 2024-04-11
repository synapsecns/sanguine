import { isNull } from 'lodash'
import { useAppSelector, useAppDispatch } from '@/store/hooks'
import { zeroAddress } from 'viem'
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

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useAppDispatch()
  const { chain } = useNetwork()
  const { isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, fromToken, fromValue } = useBridgeState()
  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

  const { gasData } = useAppSelector((state) => state.gasData)
  const { gasPrice, maxFeePerGas } = gasData?.formatted
  const { rawGasCost, parsedGasCost } = calculateGasCost(gasPrice, 500_000)

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

  console.log('rawGasCost:', rawGasCost)
  console.log('parsedGasCost: ', parsedGasCost)
  console.log('parsedGasBalance:', parseFloat(parsedBalance))
  console.log('maxBridgeableGas: ', maxBridgeableGas)

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

  const showMaxButton = (): boolean => {
    if (!hasMounted || !isConnected) return false
    if (isGasToken && isNull(parsedGasCost)) return false
    return true
  }

  const showGasReserved = (): boolean => {
    if (!hasMounted || !isConnected) return false
    if (!parsedGasCost) return false
    if (isGasToken && !isEmpty(fromValue) && !hasOnlyZeroes(fromValue)) {
      return true
    }
  }

  const isGasInputMoreThanBridgeableMax = (): boolean => {
    if (isGasToken && parsedGasCost && fromValue && parsedBalance) {
      return (
        parseFloat(fromValue) >
        parseFloat(parsedBalance) - parseFloat(parsedGasCost)
      )
    } else {
      return false
    }
  }

  const isGasBalanceLessThanCost = (): boolean => {
    if (isGasToken && parsedGasCost && parsedBalance) {
      return parseFloat(parsedGasCost) > parseFloat(parsedBalance)
    } else {
      return false
    }
  }

  console.log('isGasBalanceLessThanCost outside:', isGasBalanceLessThanCost())

  const isTraceBalance = (): boolean => {
    if (!balance || !parsedBalance) return false
    if (balance && hasOnlyZeroes(parsedBalance)) return true
    return false
  }

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
            hasMounted={hasMounted}
            isConnected={isConnected}
            showValue={showValue}
            handleFromValueChange={handleFromValueChange}
            parsedBalance={parsedBalance}
            onMaxBalance={onMaxBalance}
          />
          <AvailableBalance
            fromValue={fromValue}
            balance={balance}
            parsedBalance={parsedBalance}
            isGasToken={isGasToken}
            parsedGasCost={parsedGasCost}
            onMaxBalance={onMaxBalance}
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
  fromValue,
  balance,
  parsedBalance,
  isGasToken = false,
  parsedGasCost,
  onMaxBalance,
  hasMounted,
  isConnected,
  disabled = false,
}: {
  fromValue: string
  balance?: bigint
  parsedBalance?: string
  isGasToken?: boolean
  parsedGasCost?: string
  onMaxBalance?: () => void
  hasMounted: boolean
  isConnected: boolean
  disabled?: boolean
}) => {
  const labelClassName = joinClassNames({
    space: 'block',
    textColor: 'text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  const isTraceBalance = (): boolean => {
    if (!balance || !parsedBalance) return false
    if (balance && hasOnlyZeroes(parsedBalance)) return true
    return false
  }

  const isGasCostCoveredByInput = (): boolean => {
    if (!isGasToken) return true

    if (isGasToken && parsedGasCost && fromValue && parsedBalance) {
      return (
        parseFloat(fromValue) >
        parseFloat(parsedBalance) - parseFloat(parsedGasCost)
      )
    }

    return true
  }

  const isGasBalanceLessThanCost = (): boolean => {
    if (isGasToken && parsedGasCost && parsedBalance) {
      return parseFloat(parsedGasCost) > parseFloat(parsedBalance)
    } else {
      return false
    }
  }

  console.log('isGasBalanceLessThanCost inside:', isGasBalanceLessThanCost())

  if (hasMounted && isConnected && !disabled) {
    return (
      <HoverTooltip
        isActive={isGasBalanceLessThanCost() || !isGasCostCoveredByInput()}
        hoverContent={
          <div className="whitespace-nowrap">
            Gas fees may exceed your available balance
          </div>
        }
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
  }
  return null
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

// {hasMounted &&
//   isConnected &&
//   (isGasToken &&
//   showGasReserved() &&
//   isGasInputMoreThanBridgeableMax() ? (
//     <label
//       htmlFor="inputRow"
//       className={`
//         text-xs text-secondaryTextColor transition-all duration-150 transform-gpu
//         ${
//           (isGasBalanceLessThanCost() ||
//             isGasInputMoreThanBridgeableMax()) &&
//           'text-yellow-500'
//         }
//         `}
//     >
//       <HoverTooltip
//         isActive={
//           isGasBalanceLessThanCost() ||
//           isGasInputMoreThanBridgeableMax()
//         }
//         hoverContent={
//           isGasInputMoreThanBridgeableMax() ? (
//             <div className="whitespace-nowrap">
//               Requested bridge amount may not cover gas fees
//             </div>
//           ) : (
//             <div className="whitespace-nowrap">
//               Gas fees may exceed your available balance
//             </div>
//           )
//         }
//       >
//         {parsedGasCost.toFixed(4)}
//         <span className="text-opacity-50">
//           {' '}
//           estimated gas fee
//         </span>
//       </HoverTooltip>
//     </label>
//   ) : (
//     <label
//       htmlFor="inputRow"
//       onClick={onAvailableBalance}
//       className={`
//         text-xs text-white transition-all duration-150 transform-gpu
//         hover:text-opacity-70 hover:cursor-pointer
//       `}
//     >
//       <HoverTooltip
//         isActive={isGasBalanceLessThanCost()}
//         hoverContent={
//           <div className="whitespace-nowrap">
//             Gas fees may exceed your available balance
//           </div>
//         }
//       >
//         {isTraceBalance()
//           ? '< 0.0001'
//           : trimmedParsedBalance ?? '0.0'}
//         <span className="text-opacity-50 text-secondaryTextColor">
//           {' '}
//           available
//         </span>
//       </HoverTooltip>
//     </label>
//   ))}
