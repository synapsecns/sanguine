import { readContract } from '@wagmi/core'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { Token } from '@/utils/types'
import { wagmiConfig } from '@/wagmiConfig'

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

  const minToMint = await readContract(wagmiConfig, {
    chainId: chainId as any,
    address: poolAddress,
    abi,
    functionName: 'calculateTokenAmount',
    args: [Object.values(inputAmounts), true],
  })

  return minToMint
}
