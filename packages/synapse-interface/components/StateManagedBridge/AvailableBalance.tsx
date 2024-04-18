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

  const isTraceBalance = (): boolean => {
    return balance && !hasOnlyZeroes(parsedBalanceFull)
  }

  const isTraceGasCost = (): boolean => {
    if (!parsedGasCost) {
      return false
    } else {
      const shortenedGasCost = parseFloat(parsedGasCost).toFixed(4)
      return Number(shortenedGasCost) === 0 && !hasOnlyZeroes(parsedGasCost)
    }
  }

  const isInputGreaterThanBalanceMinusGasFees = (): boolean => {
    if (!isGasToken || !parsedGasCost || !parsedBalanceFull || !fromValue) {
      return true
    } else {
      return (
        parseFloat(fromValue) >
        parseFloat(parsedBalanceFull) - parseFloat(parsedGasCost)
      )
    }
  }

  const isBalanceGreaterThanGasFees = (): boolean => {
    if (!isGasToken || !parsedGasCost || !parsedBalanceFull) {
      return true
    } else {
      return parseFloat(parsedGasCost) < parseFloat(parsedBalanceFull)
    }
  }

  const showGasReserved = (): boolean => {
    return (
      fromValue &&
      !hasOnlyZeroes(fromValue) &&
      parsedGasCost &&
      isGasToken &&
      isInputGreaterThanBalanceMinusGasFees()
    )
  }

  console.log('maxBridgeableBalance:', maxBridgeableBalance)

  const gasReserved = showGasReserved() ? parseFloat(parsedGasCost) : undefined

  console.log('gasReserved:', gasReserved)
  console.log('parsedGasCost:', parsedGasCost)

  let tooltipContent = null

  if (isGasToken && parsedGasCost) {
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
      </div>
    )
  } else if (
    isGasToken &&
    parsedGasCost &&
    isInputGreaterThanBalanceMinusGasFees()
  ) {
    tooltipContent = (
      <div className="whitespace-nowrap">
        You may not have enough to cover gas fees.
      </div>
    )
  } else if (isGasToken && !isBalanceGreaterThanGasFees()) {
    tooltipContent = (
      <div className="whitespace-nowrap">
        Gas fees may exceed your available balance.
      </div>
    )
  }

  const labelClassName = joinClassNames({
    space: 'block',
    textColor: `text-xxs md:text-xs ${
      fromValue &&
      parsedGasCost &&
      balance &&
      !hasOnlyZeroes(fromValue) &&
      isInputGreaterThanBalanceMinusGasFees()
        ? '!text-yellowText'
        : ''
    }`,
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  console.log(
    'isInputGreaterThanBalanceMinusGasFees: ',
    isInputGreaterThanBalanceMinusGasFees()
  )

  if (
    fromValue &&
    !hasOnlyZeroes(fromValue) &&
    showGasReserved() &&
    isInputGreaterThanBalanceMinusGasFees()
  ) {
    return (
      <HoverTooltip isActive={true} hoverContent={tooltipContent}>
        <label
          htmlFor="inputRow"
          onClick={onMaxBalance}
          className={labelClassName}
        >
          <span>Gas est: </span>
          {isTraceGasCost() ? '<0.001' : parsedGasCost}
          <span> {fromToken?.symbol}</span>
        </label>
      </HoverTooltip>
    )
  } else if (hasMounted && isConnected && !disabled) {
    return (
      <HoverTooltip isActive={true} hoverContent={tooltipContent}>
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
