import { erc20ABI } from 'wagmi'
import { fetchSigner } from '@wagmi/core'
import { Contract } from 'ethers'
import { MaxInt256 } from '@ethersproject/constants'

export const approveToken = async (address, chainId, fromTokenAddress) => {
  // TODO store this erc20 and signer retrieval in a state in a parent component
  const wallet = await fetchSigner({
    chainId,
  })

  const erc20 = new Contract(fromTokenAddress, erc20ABI, wallet)
  const approveTx = await erc20.approve(address, MaxInt256)

  try {
    await approveTx.wait()
    console.log(`Transaction mined successfully: ${approveTx.hash}`)
    return approveTx
  } catch (error) {
    console.log(`Transaction failed with error: ${error}`)
  }
}
