import { ethers, AbiCoder } from 'ethers'

import { MAX_UINT256 } from '@/constants/index'

export const approveErc20Token = async ({
  spenderAddress,
  tokenAddress,
  amount,
  signer,
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

    const functionSignature = 'approve(address,uint256)'
    const abiCoder = AbiCoder.defaultAbiCoder()
    const data = abiCoder.encode(
      ['address', 'uint256'],
      [spenderAddress, amount ?? MAX_UINT256]
    )

    // The function selector is the first 4 bytes of the hash of the function signature
    const functionSelector = ethers
      .keccak256(ethers.toUtf8Bytes(functionSignature))
      .slice(0, 10)

    // The complete data for the transaction combines the selector and the encoded arguments
    const txData = functionSelector + data.slice(2) // remove 0x from data

    const tx = await signer.sendTransaction({
      to: tokenAddress,
      data: txData,
    })

    // Wait for the transaction to be mined
    const receipt = await tx.wait()

    return receipt
  } catch (error) {
    console.error('approveErc20Token: ', error)
    return error
  }
}
