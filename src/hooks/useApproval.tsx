import { useState, useEffect, useMemo, useCallback } from 'react'
import { approveErc20Token } from '@/utils/actions/approveErc20Token'
import { Address } from 'viem'

interface UseApproveProps {
  spenderAddress: Address
  tokenAddress: Address
  ownerAddress: Address
  amount: bigint
  chainId: number
}

export function useApprove({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  amount,
  chainId,
}: UseApproveProps) {
  const approveCallback = useCallback(async () => {
    /**
     * Dependency Guardrail checks
     * Bail on executing Approve Callback if dependencies missing
     */
    if (!tokenAddress) {
      console.error('useApprove missing token address')
      return
    } else if (!spenderAddress) {
      console.error('useApprove missing spender address')
      return
    } else if (!amount) {
      console.error('useApprove missing amount')
      return
    } else if (!chainId) {
      console.error('useApprove missing chainId')
      return
    }

    return approveErc20Token({
      chainId,
      spenderAddress,
      tokenAddress,
      ownerAddress,
      amount,
    })
  }, [tokenAddress, spenderAddress, amount, chainId])

  return approveCallback
}
