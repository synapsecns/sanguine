import { Contract, BigNumber } from 'ethers'
import { AddressZero } from '@ethersproject/constants'
import { Address } from 'wagmi'
import { fetchSigner } from '@wagmi/core'
import toast from 'react-hot-toast'

import { txErrorHandler } from '@utils/txErrorHandler'
import { approveToken } from '@utils/approveToken'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { CHAINS_BY_ID } from '@/constants/chains'
import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import MINI_CHEF_ABI from '@/constants/abis/miniChef.json'
import { Token } from '../types'

export const approve = async (
  pool: Token,
  inputValue: BigNumber,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  let pendingPopup: any
  let successPopup: any

  if (inputValue.isZero()) {
    return
  }

  pendingPopup = toast(`Requesting approval on ${currentChainName}`, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  try {
    await approveToken(
      MINICHEF_ADDRESSES[chainId],
      chainId,
      pool.addresses[chainId],
      inputValue
    ).then((successTx) => {
      if (successTx) {
        toast.dismiss(pendingPopup)

        const successToastContent = (
          <div>
            <div>Successfully approved on {currentChainName}</div>
            <ExplorerToastLink
              transactionHash={successTx?.hash ?? AddressZero}
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
  } catch {}
}

export const stake = async (
  address: Address,
  chainId: number,
  poolId: number,
  inputValue: BigNumber
) => {
  const signer = await fetchSigner({
    chainId,
  })
  const miniChefContract = new Contract(
    MINICHEF_ADDRESSES[chainId],
    MINI_CHEF_ABI,
    signer
  )
  try {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    const stakeTransaction = await miniChefContract.deposit(
      poolId,
      inputValue,
      address
    )

    const tx = await stakeTransaction.wait()

    toast.success(
      <div>
        <div>{'Stake completed: '}</div>
        <ExplorerToastLink {...tx} />
      </div>
    )

    return tx
  } catch (err) {
    txErrorHandler(err)
  }
}
