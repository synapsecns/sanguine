import { Zero } from '@ethersproject/constants'

import { useMemo } from 'react'
import { useSingleContractMultipleMethods } from '@hooks/multicall'
import { useTokenContract } from '@hooks/contracts/useContract'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

/**
 * @param {Token} token
 * @param {string} spending contract addr
 */
export function useTokenAllowance(token, spender) {
  const { chainId, account } = useActiveWeb3React()
  const contract = useTokenContract(token, false)

  const [allowanceResult, totalSupplyResult] = useSingleContractMultipleMethods(
    chainId,
    contract,
    {
      allowance: [account, spender],
      totalSupply: [],
    },
    { resultOnly: true }
  )

  const allowance = allowanceResult?.[0]
  const totalSupply = totalSupplyResult?.[0]

  return useMemo(() => {
    if (token && allowance) {
      // CurrencyAmount.fromRawAmount(token, allowance.toString())
      return { allowance, totalSupply }
    } else {
      return {
        allowance: undefined,
        totalSupply: undefined,
      } // undefined // previously was undefined here
    }
  }, [token, allowance, totalSupply, chainId, account, contract])
}
