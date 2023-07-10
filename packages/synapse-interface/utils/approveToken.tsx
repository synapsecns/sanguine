import { Address } from '@wagmi/core'
import { CHAINS_BY_ID } from '@/constants/chains'
import { txErrorHandler } from './txErrorHandler'
import toast from 'react-hot-toast'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { zeroAddress } from 'viem'
import { approveErc20Token } from '@/actions/approveErc20Token'

export const approveToken = async (
  address: string,
  chainId: number,
  tokenAddress: string,
  amount?: bigint
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Requesting approval on ${currentChainName}`, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  try {
    const approveTx = await approveErc20Token({
      spender: address as Address,
      chainId,
      tokenAddress: tokenAddress as Address,
      amount,
    })

    if (approveTx.status === 'success') {
      toast.dismiss(pendingPopup)

      const successToastContent = (
        <div>
          <div>Successfully approved on {currentChainName}</div>
          <ExplorerToastLink
            transactionHash={approveTx?.transactionHash ?? zeroAddress}
            chainId={chainId}
          />
        </div>
      )

      successPopup = toast.success(successToastContent, {
        id: 'approve-success-popup',
        duration: 10000,
      })
    }

    return approveTx.transactionHash
  } catch (error) {
    toast.dismiss(pendingPopup)
    console.log(`Transaction failed with error: ${error}`)
    txErrorHandler(error)
    throw error
  }
}
