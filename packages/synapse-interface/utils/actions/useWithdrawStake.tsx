import { useBlockNumber, useAccount, useNetwork, Address, Chain } from 'wagmi'
import toast from 'react-hot-toast'

import { useMiniChefContract } from '../contracts/useMiniChefContract'
import { BigNumberish, Contract } from 'ethers'
import ExplorerToastLink from '@/components/ExplorerToastLink'

type WithdrawStakeFn = (params: WithdrawStakeProps) => Promise<void>

interface WithdrawStakeProps {
  poolId: number
  amount: BigNumberish
  account: Address
}

export const useWithdrawStake = (): WithdrawStakeFn => {
  const [miniChefContract, miniChefAddress]: [Contract, Address] =
    useMiniChefContract()
  const { address } = useAccount()
  const { chain } = useNetwork()

  return async function withdrawStake({
    poolId,
    amount,
    account,
  }: WithdrawStakeProps) {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    const stakeTransaction = await miniChefContract.withdraw(
      poolId,
      amount,
      account
    )

    const tx = await stakeTransaction.wait()

    const toastContent = (
      <div>
        <div>Withdraw completed:</div>
        <ExplorerToastLink {...tx} chainId={chain.id} />
      </div>
    )

    toast.success(toastContent)

    return Promise.resolve()
  }
}
