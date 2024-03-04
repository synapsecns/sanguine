import { readContract } from '@wagmi/core'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import type { Token } from '@/utils/types'

export const swapPoolCalculateTokenAmount = async ({
  chainId,
  pool,
  inputAmounts,
}: {
  chainId: number
  pool: Token
  inputAmounts: any
}) => {
  const { abi, poolAddress } = getSwapDepositContractFields(pool, chainId)

  const minToMint = await readContract({
    chainId,
    address: poolAddress,
    abi,
    functionName: 'calculateTokenAmount',
    args: [Object.values(inputAmounts), true],
  })

  return minToMint
}
