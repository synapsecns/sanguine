import { Address } from 'wagmi'
import toast from 'react-hot-toast'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import MINI_CHEF_ABI from '@/constants/abis/miniChef.json'

import { BigNumber, Contract } from 'ethers'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { fetchSigner } from '@wagmi/core'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { shortenAddress } from '../shortenAddress'

export const withdrawStake = async (
  address: Address,
  chainId: number,
  poolId: number,
  inputValue: BigNumber
) => {
  const wallet = await fetchSigner({
    chainId,
  })
  const miniChefContract = new Contract(
    MINICHEF_ADDRESSES[chainId],
    MINI_CHEF_ABI,
    wallet
  )
  try {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    segmentAnalyticsEvent(
      `[Withdraw Stake] ${shortenAddress(address)} Attempt`,
      {
        poolId,
        inputValue,
      }
    )

    const withdrawTransactionArgs = [poolId, inputValue, address]

    const stakeTransaction = await miniChefContract.withdraw(
      ...withdrawTransactionArgs
    )

    const tx = await stakeTransaction.wait()

    const toastContent = (
      <div>
        <div>Withdraw completed:</div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    toast.success(toastContent)

    segmentAnalyticsEvent(
      `[Withdraw Stake] ${shortenAddress(address)} Success`,
      {
        poolId,
        inputValue,
      }
    )

    return tx
  } catch (err) {
    segmentAnalyticsEvent(`[Withdraw Stake] ${shortenAddress(address)} Error`, {
      poolId,
      inputValue,
      errorCode: err.code,
    })
    txErrorHandler(err)
  }
}
