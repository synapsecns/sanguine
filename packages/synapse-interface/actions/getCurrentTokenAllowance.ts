import { getPublicClient } from '@wagmi/core'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'
import { publicClientToProvider } from '@/ethers'

export const getCurrentTokenAllowance = async (
  address,
  fromChainId,
  fromToken,
  routerAddress: string
) => {
  const publicClient = getPublicClient({
    chainId: fromChainId,
  })

  const erc20 = new Contract(
    fromToken.addresses[fromChainId],
    erc20ABI,
    publicClientToProvider(publicClient)
  )
  const allowance = await erc20.allowance(address, routerAddress)
  return allowance
}
