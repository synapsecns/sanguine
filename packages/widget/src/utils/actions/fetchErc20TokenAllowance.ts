import { ZeroAddress, ethers } from 'ethers'
import erc20ABI from '../../constants/abis/erc20.json'

export const fetchErc20TokenAllowance = async ({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  provider,
}: {
  spenderAddress: string
  tokenAddress: string
  ownerAddress: string
  provider: any
}) => {
  try {
    if (!spenderAddress) {
      throw new Error('Require Spender Address')
    }
    if (!tokenAddress) {
      throw new Error('Require Token Address')
    }
    if (!ownerAddress) {
      throw new Error('Require Token Owner Address')
    }
    if (!provider) {
      throw new Error('Require Provider')
    }

    if (tokenAddress === ZeroAddress) return

    // Create a new instance of Contract
    const tokenContract = new ethers.Contract(tokenAddress, erc20ABI, provider)

    // Call the allowance function
    const allowance = await tokenContract.allowance(
      ownerAddress,
      spenderAddress
    )

    // Convert the returned value to a bigint for consistency
    return BigInt(allowance.toString())
  } catch (error) {
    console.error('fetchErc20TokenAllowance: ', error)
    return error
  }
}
