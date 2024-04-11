import React from 'react'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { joinClassNames } from '@/utils/joinClassNames'
import { HoverTooltip } from './InputContainer'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'

export const AvailableBalance = ({
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
    if (!isGasToken && !parsedGasCost) return true

    if (isGasToken && parsedGasCost && fromValue && parsedBalanceFull) {
      return (
        parseFloat(fromValue) >
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
      ? parseFloat(parsedGasCost)
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
