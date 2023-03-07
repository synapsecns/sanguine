import _ from 'lodash'

import toast from 'react-hot-toast'

import { txErrorHandler } from '@utils/txErrorHandler'


import { ChainId } from '@constants/networks'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useBlockNumber } from '@hooks/useBlockNumber'
import { useTokenContract } from '@hooks/contracts/useContract'

import ExplorerToastLink from '@components/ExplorerToastLink'



export function useRevokeToken(token) {
  const tokenContract = useTokenContract(token)

  const { account, chainId } = useActiveWeb3React()

  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)



  return async function revokeToken({ addrToRevoke }) {
    try {
      if (!account) throw new Error('Wallet must be connected')
      if (!tokenContract) throw new Error('Swap contract is not loaded')

      // For each token being deposited, check the allowance and approve it if necessary
      if (tokenContract == null) return


      const revokeTransaction = await tokenContract.approve(addrToRevoke, 0)

      toast(`Revoking Approval`)

      const tx = await revokeTransaction.wait()
      if (tx?.status === 1) {

        toast.success(
          <div>
            <div className="w-full">
              {`Successful `}
            </div>
            <ExplorerToastLink {...tx} chainId={chainId} />
          </div>
        )
      }
      setBlockNumber(tx.blockNumber)
      return tx
    } catch (err) {
      txErrorHandler(err)
    }
  }
}
