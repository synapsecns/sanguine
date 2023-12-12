import { ethers } from 'ethers'
import erc20ABI from '../../constants/abis/erc20.json'

export const fetchErc20TokenAllowance = async ({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  provider,
  signer,
}) => {
  try {
    if (!provider || !signer) {
      console.error('Web3 provider or signer is not available')
      throw new Error('Web3 provider or signer is not available')
    }

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
    console.error('Error in fetchErc20TokenAllowance: ', error)
    return error
  }
}
