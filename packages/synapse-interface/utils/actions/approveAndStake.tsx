import { Address } from 'wagmi'
import toast from 'react-hot-toast'

import { txErrorHandler } from '@utils/txErrorHandler'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { CHAINS_BY_ID } from '@/constants/chains'
import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import { Token } from '../types'
import { TransactionReceipt, zeroAddress } from 'viem'
import { approveErc20Token } from '@/actions/approveErc20Token'
import { stakeLpToken } from '@/actions/stakeLpToken'

export const approve = async (
  pool: Token,
  inputValue: bigint,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  let pendingPopup: any
  let successPopup: any

  if (inputValue === 0n) {
    return
  }

  pendingPopup = toast(`Requesting approval on ${currentChainName}`, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  try {
    const txReceipt: TransactionReceipt = await approveErc20Token({
      chainId,
      tokenAddress: pool.addresses[chainId] as Address,
      spender: MINICHEF_ADDRESSES[chainId],
      amount: inputValue,
    })

    if (txReceipt.status === 'success') {
      toast.dismiss(pendingPopup)

      const successToastContent = (
        <div>
          <div>Successfully approved on {currentChainName}</div>
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
  inputValue: bigint
) => {
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Starting your deposit...`, {
    id: 'deposit-in-progress-popup',
    duration: Infinity,
  })

  try {
    if (!address) throw new Error('Wallet must be connected')

    const tx: TransactionReceipt = await stakeLpToken({
      address,
      chainId,
      poolId,
      amount: inputValue,
      lpAddress: MINICHEF_ADDRESSES[chainId],
    })

    toast.dismiss(pendingPopup)

    const successToastContent = (
      <div>
        <div>Stake Completed:</div>
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

    return tx
  } catch (err) {
    toast.dismiss(pendingPopup)
    txErrorHandler(err)
    return err
  }
}
