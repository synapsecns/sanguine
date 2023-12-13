import { useState, useEffect, useMemo, ReactNode } from 'react'
import { fetchErc20TokenAllowance } from '@/utils/actions/fetchErc20TokenAllowance'

export enum UseAllowanceError {
  REQUIRE_SPENDER_ADDRESS = 'Allowance: Missing Spender Address',
  REQUIRE_TOKEN_ADDRESS = 'Allowance: Missing Token Address',
  REQUIRE_OWNER_ADDRESS = 'Allowance: Missing Owner Address',
  REQUIRE_CHAIN_ID = 'Allowance: Missing ChainId',
}

interface UseAllowanceProps {
  spenderAddress: string
  tokenAddress: string
  ownerAddress: string
  chainId: number
  signer: any
  provider: any
}

export function useAllowance({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  chainId,
  signer,
  provider,
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
      const allowance: bigint = await fetchErc20TokenAllowance({
        spenderAddress: spenderAddress,
        tokenAddress: tokenAddress,
        ownerAddress: ownerAddress,
        // chainId: chainId,
        signer: signer,
        provider: provider,
      })
      console.log('fetched allowance: ', allowance)
      setAllowance(allowance)
    } catch (error) {
      setError(error)
    }
  }

  /** Fetch Token Allowance when props update */
  useEffect(() => {
    if (spenderAddress && tokenAddress && ownerAddress && chainId) {
      console.log('fetching allowance')
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
  }, [
    spenderAddress,
    tokenAddress,
    ownerAddress,
    chainId,
    getTokenAllowance,
    allowance,
    error,
  ])
}
