import { MAX_UINT256 } from '@/constants/index'
import { ethers } from 'ethers'
import erc20ABI from '../../constants/abis/erc20.json'

export const approveErc20Token = async ({
  spenderAddress,
  tokenAddress,
  amount,
  provider,
  signer,
}) => {
  try {
    if (!provider || !signer) {
      console.error('Web3 provider or signer is not available')
      throw new Error('Web3 provider or signer is not available')
    }
    // Create a new instance of Contract with the signer
    const tokenContract = new ethers.Contract(tokenAddress, erc20ABI, signer)

    // If amount is not provided, use maximum uint256
    const approveAmount = amount ?? MAX_UINT256

    // Send approve transaction
    const tx = await tokenContract.approve(spenderAddress, approveAmount)

    // Wait for the transaction to be mined
    const receipt = await tx.wait()

    console.log('receipt:', receipt)

    return receipt
  } catch (error) {
    console.error('Error in approveErc20Token: ', error)
    return error
  }
}
