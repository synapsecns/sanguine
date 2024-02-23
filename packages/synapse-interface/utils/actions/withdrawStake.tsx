import toast from 'react-hot-toast'

import type { Address } from 'wagmi'
import type { Token } from '@types'

import { txErrorHandler } from '@/utils/txErrorHandler'
import { unstakeLpToken } from '@/actions/unstakeLpToken'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import ExplorerToastLink from '@/components/ExplorerToastLink'

export const withdrawStake = async (
  address: Address,
  chainId: number,
  poolId: number,
  pool: Token,
  inputValue: bigint
) => {
  const miniChefAddress = pool.miniChefAddress
  try {
    if (!address) throw new Error('Wallet must be connected')
    segmentAnalyticsEvent(
      `[Withdraw Stake] Attempt`,
      {
        poolId,
        inputValue,
      },
      true
    )

    const tx = await unstakeLpToken({
      address,
      chainId,
      poolId,
      amount: inputValue,
      lpAddress: miniChefAddress as Address,
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
