import { erc20ABI } from 'wagmi'
import { Address, readContract } from '@wagmi/core'

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
  const allowance = await readContract({
    chainId,
    address: tokenAddress,
    abi: erc20ABI,
    functionName: 'allowance',
    args: [address, spender],
  })

  return allowance
}
