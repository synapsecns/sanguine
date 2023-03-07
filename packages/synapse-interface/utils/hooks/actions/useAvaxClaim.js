import toast from 'react-hot-toast'

import { useMiniChefContract } from '@hooks/contracts/useMiniChefContract'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useBlockNumber } from '@hooks/useBlockNumber'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { useAvaxClaimContract } from '@hooks/contracts/useAvaxClaimContract'


export function useAvaxClaim() {
  const { account, chainId } = useActiveWeb3React()
  const avaxClaimContract = useAvaxClaimContract()
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)

  return async function claimAvaxNusd() {
    if (!account) throw new Error('Wallet must be connected')
    const avaxClaimTransaction = await avaxClaimContract.claim(account)


    const tx = await avaxClaimTransaction.wait()


    toast.success(
      <div>
        <div>
          NUSD Claim Completed
        </div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    setBlockNumber(tx.blockNumber)

    return Promise.resolve()
  }
}
