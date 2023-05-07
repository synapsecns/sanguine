import { fetchSigner } from '@wagmi/core'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'

export const getTokenAllowance = async (
  routerAddress: string,
  tokenAddress: string,
  address: string,
  chainId: number
) => {
  const wallet = await fetchSigner({
    chainId,
  })
  const erc20 = new Contract(tokenAddress, erc20ABI, wallet)
  const allowance = await erc20.allowance(address, routerAddress)
  return allowance
}
