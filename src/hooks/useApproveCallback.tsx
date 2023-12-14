import { useMemo, useCallback, useState } from 'react'
import { JsonRpcApiProvider, BrowserProvider } from 'ethers'
import { approveErc20Token } from '@/utils/actions/approveErc20Token'

export enum UseApproveCallbackError {
  REQUIRE_SPENDER_ADDRESS = 'Approve: Missing Spender Address',
  REQUIRE_TOKEN_ADDRESS = 'Approve: Missing Token Address',
  REQUIRE_OWNER_ADDRESS = 'Approve: Missing Owner Address',
  REQUIRE_AMOUNT = 'Approve: Missing valid amount',
  REQUIRE_CHAIN_ID = 'Approve: Missing ChainId',
}

export enum ApproveCallbackState {
  PENDING = 'pending',
  SUCCESS = 'success',
  IDLE = 'idle',
}

export interface UseApproveCallbackProps {
  spenderAddress: string
  tokenAddress: string
  ownerAddress: string
  amount: bigint
  chainId: number
  onSuccess: () => any
  signer: JsonRpcApiProvider | BrowserProvider
}

export function useApproveCallback({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  amount,
  chainId,
  onSuccess,
  signer,
}: UseApproveCallbackProps) {
  const [approveState, setApproveState] = useState<ApproveCallbackState>(
    ApproveCallbackState.IDLE
  )

  const ApproveStateCallback = {
    initiateApproval: async () => {
      setApproveState(ApproveCallbackState.PENDING)
    },
    successApproval: async () => {
      await onSuccess()
      setApproveState(ApproveCallbackState.SUCCESS)
    },
    resetState: () => {
      setApproveState(ApproveCallbackState.IDLE)
    },
  }

  const approveCallback = useCallback(async () => {
    try {
      /**
       * Dependency Guardrail checks
       * Bail on executing Approve Callback if dependencies missing
       * Throw respective error when dependency missing
       */
      if (!tokenAddress) {
        throw new Error(UseApproveCallbackError.REQUIRE_TOKEN_ADDRESS)
      } else if (!spenderAddress) {
        throw new Error(UseApproveCallbackError.REQUIRE_SPENDER_ADDRESS)
      } else if (!ownerAddress) {
        throw new Error(UseApproveCallbackError.REQUIRE_SPENDER_ADDRESS)
      } else if (!amount) {
        throw new Error(UseApproveCallbackError.REQUIRE_SPENDER_ADDRESS)
      } else if (!chainId) {
        throw new Error(UseApproveCallbackError.REQUIRE_SPENDER_ADDRESS)
      }

      ApproveStateCallback.initiateApproval()
      await approveErc20Token({
        spenderAddress,
        tokenAddress,
        amount,
        signer,
      })
      await ApproveStateCallback.successApproval()
    } catch (error) {
      ApproveStateCallback.resetState()
      console.error(error)
    }
  }, [tokenAddress, spenderAddress, amount, chainId])

  return useMemo(() => {
    /**
     * Return error state based on current hook props
     * Callback will throw error when invoked with missing dependencies
     */
    if (!tokenAddress) {
      return {
        state: approveState,
        callback: () => approveCallback(),
        error: UseApproveCallbackError.REQUIRE_TOKEN_ADDRESS,
      }
    } else if (!spenderAddress) {
      return {
        state: approveState,
        callback: () => approveCallback(),
        error: UseApproveCallbackError.REQUIRE_SPENDER_ADDRESS,
      }
    } else if (!ownerAddress) {
      return {
        state: approveState,
        callback: () => approveCallback(),
        error: UseApproveCallbackError.REQUIRE_OWNER_ADDRESS,
      }
    } else if (!amount) {
      return {
        state: approveState,
        callback: () => approveCallback(),
        error: UseApproveCallbackError.REQUIRE_AMOUNT,
      }
    } else if (!chainId) {
      return {
        state: approveState,
        callback: () => approveCallback(),
        error: UseApproveCallbackError.REQUIRE_CHAIN_ID,
      }
    }

    return {
      state: approveState,
      callback: () => approveCallback(),
      error: null,
    }
  }, [
    chainId,
    spenderAddress,
    tokenAddress,
    ownerAddress,
    amount,
    approveCallback,
    approveState,
  ])
}
