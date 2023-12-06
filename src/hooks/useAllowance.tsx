import { useState, useEffect, useMemo, ReactNode } from 'react'
import { Address } from 'viem'
import { fetchErc20TokenAllowance } from '@/utils/actions/fetchErc20TokenAllowance'

export enum UseAllowanceError {
  REQUIRE_SPENDER_ADDRESS = 'Missing Spender Address',
  REQUIRE_TOKEN_ADDRESS = 'Missing Token Address',
  REQUIRE_OWNER_ADDRESS = 'Missing Owner Address',
  REQUIRE_CHAIN_ID = 'Missing ChainId',
}

interface UseAllowanceProps {
  spenderAddress: Address
  tokenAddress: Address
  ownerAddress: Address
  chainId: number
}

export function useAllowance({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  chainId,
}: UseAllowanceProps): {
  allowance: bigint
  checkAllowanceCallback: () => Promise<void>
  error: any
} {
  const [allowance, setAllowance] = useState<bigint>(null)
  const [error, setError] = useState<ReactNode | any>(null)

  const getTokenAllowance: () => Promise<void> = async () => {
    try {
      setError(null)
      console.log('fetching allowance')
      const allowance: bigint = await fetchErc20TokenAllowance({
        spenderAddress: spenderAddress,
        tokenAddress: tokenAddress,
        ownerAddress: ownerAddress,
        chainId: chainId,
      })
      console.log('fetched allowance:', allowance)

      setAllowance(allowance)
    } catch (error) {
      setError(error)
    }
  }

  /** Fetch Token Allowance when props update */
  useEffect(() => {
    if (spenderAddress && tokenAddress && ownerAddress && chainId) {
      getTokenAllowance()
    }
  }, [spenderAddress, tokenAddress, ownerAddress, chainId])

  return useMemo(() => {
    /** Guardrail to check required dependencies to fetch allowance */
    if (!spenderAddress) {
      return {
        allowance: null,
        checkAllowanceCallback: null,
        error: UseAllowanceError.REQUIRE_SPENDER_ADDRESS,
      }
    } else if (!tokenAddress) {
      return {
        allowance: null,
        checkAllowanceCallback: null,
        error: UseAllowanceError.REQUIRE_TOKEN_ADDRESS,
      }
    } else if (!ownerAddress) {
      return {
        allowance: null,
        checkAllowanceCallback: null,
        error: UseAllowanceError.REQUIRE_OWNER_ADDRESS,
      }
    } else if (!chainId) {
      return {
        allowance: null,
        checkAllowanceCallback: null,
        error: UseAllowanceError.REQUIRE_CHAIN_ID,
      }
    }

    return {
      allowance,
      checkAllowanceCallback: () => getTokenAllowance(),
      error,
    }
  }, [spenderAddress, tokenAddress, ownerAddress, chainId])
}
