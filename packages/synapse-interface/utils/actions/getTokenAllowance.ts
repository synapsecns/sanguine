import { getBalance } from '@wagmi/core'
import { type Address, zeroAddress } from 'viem'

import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { wagmiConfig } from '@/wagmiConfig'

export const getTokenAllowance = async (
  routerAddress: Address,
  tokenAddress: Address,
  address: Address,
  chainId: number
) => {
  let fetchedBalance
  let allowance

  if (tokenAddress === zeroAddress) {
    fetchedBalance = await getBalance(wagmiConfig, {
      address,
      chainId: chainId as any,
    })

    allowance = fetchedBalance.value
  } else {
    const allowance = await getErc20TokenAllowance({
      address,
      chainId,
      tokenAddress,
      spender: routerAddress,
    })

    return allowance
  }

  return allowance
}
