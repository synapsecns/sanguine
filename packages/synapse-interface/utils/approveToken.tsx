import { erc20ABI } from 'wagmi'
import { fetchSigner } from '@wagmi/core'
import { Contract } from 'ethers'
import { MaxInt256 } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import toast from 'react-hot-toast'
import ExplorerToastLink from '@components/ExplorerToastLink'

export const approveToken = async (
  address: string,
  chainId: number,
  tokenAddress: string,
  amount?: BigNumber
) => {
  // TODO store this erc20 and signer retrieval in a state in a parent component
  const signer = await fetchSigner({
    chainId,
  })

  const erc20 = new Contract(tokenAddress, erc20ABI, signer)
  try {
    const approveTx = await erc20.approve(address, amount ?? MaxInt256)

    await approveTx.wait()
    console.log(`Transaction mined successfully: ${approveTx.hash}`)
    const toastContent = (
      <div>
        <div>Token Approved!</div>
        <ExplorerToastLink {...approveTx} chainId={chainId} />
      </div>
    )

    toast.success(toastContent)
    return approveTx
  } catch (error) {
    console.log(`Transaction failed with error: ${error}`)
  }
}
