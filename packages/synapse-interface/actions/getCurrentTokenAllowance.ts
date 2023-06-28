import { getProvider } from '@wagmi/core'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'

export const getCurrentTokenAllowance = async (
  address,
  fromChainId,
  fromToken,
  routerAddress: string
) => {
  const provider = getProvider({
    chainId: fromChainId,
  })

  const erc20 = new Contract(
    fromToken.addresses[fromChainId],
    erc20ABI,
    provider
  )
  const allowance = await erc20.allowance(address, routerAddress)
  return allowance
}
