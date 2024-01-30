import { MAX_UINT256 } from '@/constants/index'
import { JsonRpcApiProvider, BrowserProvider, ethers } from 'ethers'
import erc20ABI from '../../constants/abis/erc20.json'

export const approveErc20Token = async ({
  spenderAddress,
  tokenAddress,
  amount,
  signer,
}: {
  spenderAddress: string
  tokenAddress: string
  amount: bigint
  signer: JsonRpcApiProvider | BrowserProvider
}) => {
  try {
    if (!spenderAddress) {
      throw new Error('Require Spender Address')
    }
    if (!tokenAddress) {
      throw new Error('Require Token Address')
    }
    if (!signer) {
      throw new Error('Require Signer')
    }
    // Create a new instance of Contract with the signer
    const tokenContract = new ethers.Contract(tokenAddress, erc20ABI, signer)

    // If amount is not provided, use maximum uint256
    const approveAmount = amount ?? MAX_UINT256

    // Send approve transaction
    const tx = await tokenContract.approve(spenderAddress, approveAmount)

    // Wait for the transaction to be mined
    const receipt = await tx.wait()

    return receipt
  } catch (error) {
    console.error('approveErc20Token: ', error)
    return error
  }
}
