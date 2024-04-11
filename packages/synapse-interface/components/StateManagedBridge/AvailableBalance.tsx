import React from 'react'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { joinClassNames } from '@/utils/joinClassNames'
import { HoverTooltip } from './InputContainer'

export const AvailableBalance = ({
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

  const isGasCostCovered = (): boolean => {
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

  if (hasMounted && isConnected && !disabled) {
    return (
      <HoverTooltip
        // isActive={isGasBalanceLessThanCost() || !isGasCostCovered()}
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
