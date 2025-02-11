import { CHAINS_BY_ID } from '@/constants/chains'
import { txErrorHandler } from './txErrorHandler'
import toast from 'react-hot-toast'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { type Address, zeroAddress } from 'viem'
import { approveErc20Token } from '@/actions/approveErc20Token'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { TranslatedText } from '@/components/ui/TranslatedText'

export const approveToken = async (
  address: string,
  chainId: number,
  tokenAddress: string,
  amount?: bigint
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  let pendingPopup: any
  let successPopup: any

  const msg = (
    <TranslatedText namespace="Pools.Other" id="Requesting approval" />
  )

  pendingPopup = toast(msg, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  try {
    segmentAnalyticsEvent(
      `[Approval] initiates approval`,
      {
        chainId,
        tokenAddress,
        amount,
      },
      true
    )
    const approveTx = await approveErc20Token({
      spender: address as Address,
      chainId,
      amount,
      tokenAddress: tokenAddress as Address,
    })

    if (approveTx?.status === 'success') {
      toast.dismiss(pendingPopup)

      segmentAnalyticsEvent(`[Approval] successfully approves token`, {
        chainId,
        tokenAddress,
        amount,
      })
      const successToastContent = (
        <div>
          <div>
            <TranslatedText namespace="Pools" id="Successfully approved on" />{' '}
            {currentChainName}
          </div>
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

    return approveTx?.transactionHash
  } catch (error) {
    segmentAnalyticsEvent(`[Approval] approval fails`, {
      chainId,
      tokenAddress,
      amount,
      errorCode: error.code,
    })
    toast.dismiss(pendingPopup)
    console.log(`Transaction failed with error: ${error}`)
    txErrorHandler(error)
    throw error
  }
}
