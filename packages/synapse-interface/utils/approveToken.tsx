import { erc20ABI } from 'wagmi'
import { fetchSigner } from '@wagmi/core'
import { Contract } from 'ethers'
import { MaxInt256, AddressZero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { CHAINS_BY_ID } from '@/constants/chains'
import { txErrorHandler } from './txErrorHandler'
import toast from 'react-hot-toast'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { useAnalytics } from '@/contexts/AnalyticsProvider'
import { shortenAddress } from './shortenAddress'

export const approveToken = async (
  address: string,
  chainId: number,
  tokenAddress: string,
  amount?: BigNumber
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  let pendingPopup: any
  let successPopup: any
  const analytics = useAnalytics()

  pendingPopup = toast(`Requesting approval on ${currentChainName}`, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  // TODO store this erc20 and signer retrieval in a state in a parent component
  const signer = await fetchSigner({
    chainId,
  })

  const erc20 = new Contract(tokenAddress, erc20ABI, signer)
  try {
    analytics.track(
      `[Approval] ${shortenAddress(address)} initiates approval`,
      {
        chainId,
        tokenAddress,
        amount,
      }
    )

    const approveTx = await erc20.approve(address, amount ?? MaxInt256)
    await approveTx.wait().then((successTx) => {
      if (successTx) {
        toast.dismiss(pendingPopup)

        analytics.track(
          `[Approval] ${shortenAddress(address)} successfully approves token`,
          {
            chainId,
            tokenAddress,
            amount,
          }
        )

        const successToastContent = (
          <div>
            <div>Successfully approved on {currentChainName}</div>
            <ExplorerToastLink
              transactionHash={approveTx?.hash ?? AddressZero}
              chainId={chainId}
            />
          </div>
        )

        successPopup = toast.success(successToastContent, {
          id: 'approve-success-popup',
          duration: 10000,
        })
      }
    })

    return approveTx
  } catch (error) {
    analytics.track(`[Approval] ${shortenAddress(address)} approval fails`, {
      chainId,
      tokenAddress,
      amount,
      errorCode: error.code,
    })
    toast.dismiss(pendingPopup)
    console.log(`Transaction failed with error: ${error}`)
    txErrorHandler(error)
    throw error;
  }
}
