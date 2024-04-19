import React from 'react'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { joinClassNames } from '@/utils/joinClassNames'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { HoverTooltip } from '../HoverTooltip'

export const AvailableBalance = ({
  fromChainId,
  fromValue,
  fromToken,
  balance,
  parsedBalance,
  maxBridgeableBalance,
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
  maxBridgeableBalance?: string
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

  const isValidUserInput = () => {
    return Boolean(fromValue && !hasOnlyZeroes(fromValue))
  }

  const isTraceBalance = () => {
    return Boolean(
      balance &&
        hasOnlyZeroes(parsedBalance) &&
        !hasOnlyZeroes(parsedBalanceFull)
    )
  }

  const isTraceGasCost = () => {
    return Boolean(
      parsedGasCost &&
        hasOnlyZeroes(parseFloat(parsedGasCost).toFixed(4)) &&
        !hasOnlyZeroes(parsedGasCost)
    )
  }

  const hasUserInput = isValidUserInput()
  const hasBalance = Boolean(balance || isTraceBalance())
  const hasGasCost = Boolean(parsedGasCost || isTraceGasCost())

  const isUserInputGreaterThanBalanceMinusGasFees = (): boolean => {
    if (isGasToken && hasUserInput && hasBalance && hasGasCost) {
      return (
        parseFloat(fromValue) >
        parseFloat(parsedBalanceFull) - parseFloat(parsedGasCost)
      )
    } else {
      return false
    }
  }

  const isUserBalanceGreaterThanGasFees = (): boolean => {
    if (isGasToken && hasGasCost && hasBalance) {
      return parseFloat(parsedGasCost) < parseFloat(parsedBalanceFull)
    } else {
      return true
    }
  }

  const hasNotEnoughGas =
    isUserInputGreaterThanBalanceMinusGasFees() ||
    !isUserBalanceGreaterThanGasFees()

  const showGasReserved =
    isGasToken && hasGasCost && hasUserInput && hasBalance && hasNotEnoughGas

  let tooltipContent = null

  if (isGasToken && hasGasCost) {
    tooltipContent = (
      <div className="flex flex-col space-y-2 whitespace-nowrap">
        <span>
          Available balance: {parsedBalanceFull} {fromToken?.symbol}
        </span>
        <span>
          Estimated bridgeable balance:{' '}
          {Number(maxBridgeableBalance) > 0 ? maxBridgeableBalance : '0.0'}{' '}
          {fromToken?.symbol}
        </span>
        <span>
          Estimated gas fee: {parsedGasCost} {fromToken?.symbol}
        </span>
        {hasNotEnoughGas && (
          <span>You may not have enough to cover gas fees.</span>
        )}
      </div>
    )
  }

  const labelClassName = joinClassNames({
    space: 'block',
    textColor: `text-xxs md:text-xs ${
      showGasReserved ? '!text-yellowText' : ''
    }`,
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  if (showGasReserved) {
    return (
      <HoverTooltip isActive={true} hoverContent={tooltipContent}>
        <label
          onClick={onMaxBalance}
          className={labelClassName}
          htmlFor="inputRow"
        >
          <span>Gas est: </span>
          {isTraceGasCost() ? '<0.001' : Number(parsedGasCost).toFixed(4)}
          <span> {fromToken?.symbol}</span>
        </label>
      </HoverTooltip>
    )
  } else if (hasMounted && isConnected && !disabled) {
    return (
      <HoverTooltip isActive={true} hoverContent={tooltipContent}>
        <label
          onClick={onMaxBalance}
          className={labelClassName}
          htmlFor="inputRow"
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
