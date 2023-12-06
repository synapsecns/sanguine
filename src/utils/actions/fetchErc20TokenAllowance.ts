import { erc20ABI } from 'wagmi'
import { Address } from 'viem'
import { viemPublicClient } from 'index'

export const fetchErc20TokenAllowance = async ({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  chainId,
}: {
  spenderAddress: Address
  tokenAddress: Address
  ownerAddress: Address
  chainId: number
}): Promise<bigint> => {
  try {
    const allowance = await viemPublicClient.readContract({
      address: tokenAddress,
      abi: erc20ABI,
      functionName: 'allowance',
      args: [ownerAddress, spenderAddress],
    })

    return allowance
  } catch (error) {
    console.error('Error in fetchErc20TokenAllowance: ', error)
    return error
  }
}
