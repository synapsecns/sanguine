import { fetchBalance, Address } from '@wagmi/core'
import { zeroAddress } from 'viem'

import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'

export const getTokenAllowance = async (
  routerAddress: Address,
  tokenAddress: Address,
  address: Address,
  chainId: number
) => {
  let fetchedBalance
  let allowance

  if (tokenAddress === zeroAddress) {
    fetchedBalance = await fetchBalance({
      address,
      chainId,
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
