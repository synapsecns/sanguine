import { type Address } from 'viem'
import toast from 'react-hot-toast'

import ExplorerToastLink from '@/components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { unstakeLpToken } from '@/actions/unstakeLpToken'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { Token } from '@types'
import { useTranslations } from 'next-intl'
import { TranslatedText } from '@/components/ui/TranslatedText'

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

    const txReceipt = await unstakeLpToken({
      address,
      chainId,
      poolId,
      amount: inputValue,
      lpAddress: miniChefAddress as Address,
    })

    const toastContent = (
      <div>
        <div>
          <TranslatedText key="Pools" text="Withdrawal completed" />:
        </div>
        <ExplorerToastLink
          transactionHash={txReceipt.transactionHash}
          chainId={chainId}
        />
      </div>
    )

    toast.success(toastContent)
    segmentAnalyticsEvent(`[Withdraw Stake] Success`, {
      poolId,
      inputValue,
    })

    return txReceipt
  } catch (err) {
    segmentAnalyticsEvent(`[Withdraw Stake] Error`, {
      poolId,
      inputValue,
      errorCode: err.code,
    })
    txErrorHandler(err)
  }
}
