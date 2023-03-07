import toast from 'react-hot-toast'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useMiniChefContract } from '@hooks/contracts/useMiniChefContract'
import { useTokenContract } from '@hooks/contracts/useContract'

import { checkAndApproveTokenForTrade } from '@utils/checkAndApproveTokenForTrade'

import ExplorerToastLink from '@components/ExplorerToastLink'
import { useBlockNumber } from '@hooks/useBlockNumber'


export function useApproveAndStake(token) {
  const { account, chainId } = useActiveWeb3React()
  const tokenContract = useTokenContract(token)

  const miniChefContract = useMiniChefContract()
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)

  return async function approveAndStake({ amount, poolId, infiniteApproval }) {
    if (!account) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')
    if (!tokenContract) throw new Error("tokenContract not found")

    await checkAndApproveTokenForTrade(
      tokenContract,
      miniChefContract.address,
      account,
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

    const stakeTransaction = await miniChefContract.deposit(poolId, amount, account)

    const tx = await stakeTransaction.wait()

    toast.success(
      <div>
        <div>
          {'Stake completed: '}
        </div>
        <ExplorerToastLink {...tx} />
      </div>
    )

    setBlockNumber(tx.blockNumber)
    return tx
  }
}
