import { Address } from 'wagmi'
import toast from 'react-hot-toast'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { unstakeLpToken } from '@/actions/unstakeLpToken'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

export const withdrawStake = async (
  address: Address,
  chainId: number,
  poolId: number,
  inputValue: bigint
) => {
  try {
    if (!address) throw new Error('Wallet must be connected')
    segmentAnalyticsEvent(`[Withdraw Stake] Attempt`, {
      poolId,
      inputValue,
    })

    const tx = await unstakeLpToken({
      address,
      chainId,
      poolId,
      amount: inputValue,
      lpAddress: MINICHEF_ADDRESSES[chainId],
    })

    const toastContent = (
      <div>
        <div>Withdraw completed:</div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    toast.success(toastContent)
    segmentAnalyticsEvent(`[Withdraw Stake] Success`, {
      poolId,
      inputValue,
    })

    return tx
  } catch (err) {
    segmentAnalyticsEvent(`[Withdraw Stake] Error`, {
      poolId,
      inputValue,
      errorCode: err.code,
    })
    txErrorHandler(err)
  }
}