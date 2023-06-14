import { fetchBalance, fetchSigner } from '@wagmi/core'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'

export const getTokenAllowance = async (
  routerAddress: string,
  tokenAddress: string,
  address: string,
  chainId: number
) => {
  console.log(`[getTokenAllowance] tokenAddress`, tokenAddress)
  let fetchedBalance
  let allowance
  const wallet = await fetchSigner({
    chainId,
  })

  if (tokenAddress === '0x0000000000000000000000000000000000000000') {
    fetchedBalance = await fetchBalance({
      address: address as `0x${string}`,
      chainId,
    })

    allowance = fetchedBalance.value
  } else {
    const erc20 = new Contract(tokenAddress, erc20ABI, wallet)
    allowance = await erc20.allowance(address, routerAddress)
  }

  console.log(`allowance`, allowance)

  return allowance
}
