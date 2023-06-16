import { fetchSigner } from '@wagmi/core'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'

export const getCurrentTokenAllowance = async (
  address,
  fromChainId,
  fromToken,
  routerAddress: string
) => {
  const wallet = await fetchSigner({
    chainId: fromChainId,
  })

  const erc20 = new Contract(fromToken.addresses[fromChainId], erc20ABI, wallet)
  const allowance = await erc20.allowance(address, routerAddress)
  return allowance
}
