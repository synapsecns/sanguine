import { type Address } from 'viem'
import toast from 'react-hot-toast'

import ExplorerToastLink from '@/components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { harvestLpPool } from '@/actions/harvestLpPool'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { Token } from '@types'
import { TranslatedText } from '@/components/ui/TranslatedText'

export const claimStake = async (
  chainId: number,
  address: Address,
  poolId: number,
  pool: Token
) => {
  let pendingPopup: any
  let successPopup: any
  let miniChefAddress = pool.miniChefAddress

  const msg = (
    <TranslatedText namespace="Pools.Other" id="Starting your claim" />
  )

  pendingPopup = toast(msg, {
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
        <div>
          <TranslatedText namespace="Pools" id="Claim completed" />:
        </div>
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
