import toast from 'react-hot-toast'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useMiniChefContract } from '@hooks/contracts/useMiniChefContract'
import { useBlockNumber } from '@hooks/useBlockNumber'
import ExplorerToastLink from '@components/ExplorerToastLink'


export function useWithdrawStake() {
  const { account, chainId } = useActiveWeb3React()
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)
  const miniChefContract = useMiniChefContract()

  return async function withdrawStake({poolId, amount}) {
    if (!account) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MiniChef Contract is not loaded')

    const stakeTransaction = await miniChefContract.withdraw(poolId, amount, account)

    const tx = await stakeTransaction.wait()

    toast.success(
      <div>
        <div>
          {"Withdraw completed: "}
        </div>
        <ExplorerToastLink {...tx} />
      </div>
    )
    setBlockNumber(tx.blockNumber)

    return tx
  }
}
