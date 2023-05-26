import { useBlockNumber, useAccount, useNetwork, Address, Chain } from 'wagmi'
import toast from 'react-hot-toast'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import MINI_CHEF_ABI from '@/constants/abis/miniChef.json'
import { Contract } from 'ethers'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { fetchSigner } from '@wagmi/core'

export const claimStake = async (
  chainId: number,
  address: string,
  poolId: number
) => {
  let pendingPopup: any
  let successPopup: any

  const wallet = await fetchSigner({
    chainId,
  })
  const miniChefContract = new Contract(
    MINICHEF_ADDRESSES[chainId],
    MINI_CHEF_ABI,
    wallet
  )

  pendingPopup = toast(`Starting your claim...`, {
    id: 'claim-in-progress-popup',
    duration: Infinity,
  })

  try {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    const stakeTransaction = await miniChefContract.harvest(poolId, address)
    const tx = await stakeTransaction.wait()

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

    return tx
  } catch (err) {
    toast.dismiss(pendingPopup)
    txErrorHandler(err)
  }
}
