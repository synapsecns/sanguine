import toast from 'react-hot-toast'

import { useMiniChefContract } from '@hooks/contracts/useMiniChefContract'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useBlockNumber } from '@hooks/useBlockNumber'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { useAirdropContracts } from '@hooks/contracts/useAirdropContracts'


export function useClaimAirdrop() {
  const { account, chainId } = useActiveWeb3React()
  const {nrvAirdropContract, synAirdropContract} = useAirdropContracts()
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)
  return async function claimAirdrop({ type }) {
    if (!account) throw new Error('Wallet must be connected')
    let airdropTransaction
    if (type == "SYN") {
      airdropTransaction = await synAirdropContract.claim(account)
    } else {
      airdropTransaction = await nrvAirdropContract.claim(account)
    }

    const tx = await airdropTransaction.wait()


    toast.success(
      <div>
        <div>
          {type} Claim Completed
        </div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    setBlockNumber(tx.blockNumber)

    return Promise.resolve()
  }
}
