import type { Token } from '@types'
import type { Address } from 'wagmi'
import toast from 'react-hot-toast'

import ExplorerToastLink from '@/components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { harvestLpPool } from '@/actions/harvestLpPool'
import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'


export const claimStake = async (
  chainId: number,
  address: Address,
  poolId: number,
  pool: Token
) => {
  let pendingPopup: any
  let successPopup: any
  let miniChefAddress = pool.miniChefAddress

  pendingPopup = toast(`Starting your claim...`, {
    id: 'claim-in-progress-popup',
    duration: Infinity,
  })

  try {
    if (!address) throw new Error('Wallet must be connected')
    segmentAnalyticsEvent(
      `[Claim Stake] Attempt`,
      {
        poolId,
      },
      true
    )
    const tx = await harvestLpPool({
      address,
      chainId,
      poolId,
      lpAddress: miniChefAddress as Address,
    })

    toast.dismiss(pendingPopup)

    const successToastContent = (
      <div>
        <div>Claim Completed:</div>
        <ExplorerToastLink
          transactionHash={tx?.transactionHash}
          chainId={chainId}
        />
      </div>
    )

    successPopup = toast.success(successToastContent, {
      id: 'claim-success-popup',
      duration: 10000,
    })
    segmentAnalyticsEvent(`[Claim Stake] Success`, {
      poolId,
    })

    return tx
  } catch (err) {
    segmentAnalyticsEvent(`[Claim Stake] Failure`, {
      errorCode: err.code,
    })
    toast.dismiss(pendingPopup)
    txErrorHandler(err)
  }
}
