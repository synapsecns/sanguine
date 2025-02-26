import { type Address } from 'viem'
import toast from 'react-hot-toast'

import { txErrorHandler } from '@utils/txErrorHandler'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Token } from '../types'
import { zeroAddress } from 'viem'
import { approveErc20Token } from '@/actions/approveErc20Token'
import { stakeLpToken } from '@/actions/stakeLpToken'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { TranslatedText } from '@/components/ui/TranslatedText'

export const approve = async (
  pool: Token,
  inputValue: bigint,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  const miniChefAddress = pool.miniChefAddress
  let pendingPopup: any
  let successPopup: any

  if (inputValue === 0n) {
    return
  }

  const msg = (
    <TranslatedText namespace="Pools.Other" id="Requesting approval" />
  )

  pendingPopup = toast(msg, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  try {
    const txReceipt = await approveErc20Token({
      chainId,
      tokenAddress: pool.addresses[chainId] as Address,
      spender: miniChefAddress as Address,
      amount: inputValue,
    })

    if (txReceipt.status === 'success') {
      toast.dismiss(pendingPopup)

      const successToastContent = (
        <div>
          <div>
            <TranslatedText namespace="Pools" id="Successfully approved on" />{' '}
            {currentChainName}
          </div>
          <ExplorerToastLink
            transactionHash={txReceipt?.transactionHash ?? zeroAddress}
            chainId={chainId}
          />
        </div>
      )

      successPopup = toast.success(successToastContent, {
        id: 'approve-success-popup',
        duration: 10000,
      })
    }

    return txReceipt
  } catch (error) {
    toast.dismiss(pendingPopup)
    txErrorHandler(error)
    return error
  }
}

export const stake = async (
  address: Address,
  chainId: number,
  poolId: number,
  pool: Token,
  inputValue: bigint
) => {
  let pendingPopup: any
  let successPopup: any

  const miniChefAddress = pool.miniChefAddress

  const msg = (
    <TranslatedText namespace="Pools.Other" id="Starting your deposit" />
  )

  pendingPopup = toast(msg, {
    id: 'deposit-in-progress-popup',
    duration: Infinity,
  })

  try {
    if (!address) throw new Error('Wallet must be connected')

    segmentAnalyticsEvent(
      `[Stake] Attempt`,
      {
        poolId,
        inputValue,
      },
      true
    )

    const tx = await stakeLpToken({
      address,
      chainId,
      poolId,
      amount: inputValue,
      lpAddress: miniChefAddress as Address,
    })

    toast.dismiss(pendingPopup)

    const successToastContent = (
      <div>
        <div>
          <TranslatedText namespace="Pools" id="Stake completed" />:
        </div>
        <ExplorerToastLink
          transactionHash={tx?.transactionHash}
          chainId={chainId}
        />
      </div>
    )

    successPopup = toast.success(successToastContent, {
      id: 'stake-success-popup',
      duration: 10000,
    })
    segmentAnalyticsEvent(`[Stake] Success`, {
      poolId,
      inputValue,
    })

    return tx
  } catch (err) {
    toast.dismiss(pendingPopup)
    segmentAnalyticsEvent(`[Stake] Error`, {
      poolId,
      inputValue,
      errorCode: err.code,
    })
    txErrorHandler(err)
    return err
  }
}
