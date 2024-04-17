import { readContract } from '@wagmi/core'
import { Address, erc20Abi } from 'viem'

import { wagmiConfig } from '@/wagmiConfig'

export const getErc20TokenAllowance = async ({
  address,
  chainId,
  tokenAddress,
  spender,
}: {
  address: Address
  chainId: number
  tokenAddress: Address
  spender: Address
}): Promise<bigint> => {
  try {
    const allowance = await readContract(wagmiConfig, {
      chainId,
      address: tokenAddress,
      abi: erc20Abi,
      functionName: 'allowance',
      args: [address, spender],
    })

    return allowance
  } catch (error) {
    console.log(error)
    return 0n
  }
}
