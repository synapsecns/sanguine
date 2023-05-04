import toast from 'react-hot-toast'

import { useMiniChefContract } from '../contracts/useMiniChefContract'
import { Contract } from 'ethers'
import { Address, useAccount } from 'wagmi'

import { checkAndApproveTokenForTrade } from '@utils/checkAndApproveTokenForTrade'

import ExplorerToastLink from '@components/ExplorerToastLink'

export function useApproveAndStake(token) {
  const tokenContract = useTokenContract(token)
  const { address } = useAccount()
  const [miniChefContract, miniChefAddress]: [Contract, Address] =
    useMiniChefContract()

  return async function approveAndStake({ amount, poolId, infiniteApproval }) {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')
    if (!tokenContract) throw new Error('tokenContract not found')

    await checkAndApproveTokenForTrade(
      tokenContract,
      miniChefContract.address,
      address,
      amount,
      infiniteApproval,
      {
        onTransactionStart: () => {},
        onTransactionSuccess: () => {},
        onTransactionError: () => {
          throw new Error('Your transaction could not be completed')
        },
      }
    )

    const stakeTransaction = await miniChefContract.deposit(
      poolId,
      amount,
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
  }
}
