import toast from 'react-hot-toast'

import { useMiniChefContract } from '@hooks/contracts/useMiniChefContract'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useBlockNumber } from '@hooks/useBlockNumber'
import ExplorerToastLink from '@components/ExplorerToastLink'


export function useClaimStake() {
  const { account, chainId } = useActiveWeb3React()
  const miniChefContract = useMiniChefContract()
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)
  return async function claimStake({poolId}) {
    if (!account) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    const stakeTransaction = await miniChefContract.harvest(poolId, account)

    const tx = await stakeTransaction.wait()


    toast.success(
      <div>
        <div>
          Claim completed:
        </div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    setBlockNumber(tx.blockNumber)

    return Promise.resolve()
  }
}
