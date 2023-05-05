import { useBlockNumber, useAccount, useNetwork, Address, Chain } from 'wagmi'
import toast from 'react-hot-toast'

import { useMiniChefContract } from '../contracts/useMiniChefContract'
import { Contract } from 'ethers'
import ExplorerToastLink from '@/components/ExplorerToastLink'

type ClaimStakeFn = (params: { poolId: number }) => Promise<void>

export const useClaimStake = (): ClaimStakeFn => {
  const [miniChefContract, miniChefAddress]: [Contract, Address] =
    useMiniChefContract()
  const { address } = useAccount()
  const { chain } = useNetwork()

  return async function claimStake({ poolId }: { poolId: number }) {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    const stakeTransaction = await miniChefContract.harvest(poolId, address)

    const tx = await stakeTransaction.wait()

    const toastContent = (
      <div>
        <div>Claim completed:</div>
        <ExplorerToastLink {...tx} chainId={chain.id} />
      </div>
    )

    toast.success(toastContent)

    return Promise.resolve()
  }
}
