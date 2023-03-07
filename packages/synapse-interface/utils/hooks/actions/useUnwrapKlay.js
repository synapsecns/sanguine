import toast from 'react-hot-toast'


import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { DEPRECATED_WKLAY } from "@constants/tokens/basic"
import WETH_ABI from '@abis/weth.json'
import { useContract } from "@hooks/contracts/useContract"
import { useBlockNumber } from '@hooks/useBlockNumber'

export function useUnwrapKlay() {
  const { account, chainId } = useActiveWeb3React()
  const wklayContract = useContract(DEPRECATED_WKLAY.addresses[chainId], WETH_ABI)
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)

  return async function unwrapKlay({ amount }) {
    if (!account) throw new Error('Wallet must be connected')
    let unwrapTransaction
    console.log(wklayContract)
    unwrapTransaction = await wklayContract.withdraw(amount)

    const tx = await unwrapTransaction.wait()


    toast.success(
      <div>
        <div>
          Unwrap WKLAY Completed
        </div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    setBlockNumber(tx.blockNumber)

    return Promise.resolve()
  }
}
