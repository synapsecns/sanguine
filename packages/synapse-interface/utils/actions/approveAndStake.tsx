import toast from 'react-hot-toast'
import { erc20ABI } from 'wagmi'
import { Contract, BigNumber } from 'ethers'
import { Address } from 'wagmi'
import { Token } from '../types'
import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import MINI_CHEF_ABI from '@/constants/abis/miniChef.json'
import { approveToken } from '@utils/approveToken'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'
import { fetchSigner } from '@wagmi/core'
export const approve = async (
  pool: Token,
  inputValue: BigNumber,
  chainId: number
) => {
  if (inputValue.isZero()) {
    return
  }

  await approveToken(
    MINICHEF_ADDRESSES[chainId],
    chainId,
    pool.addresses[chainId],
    inputValue
  )
}

export const stake = async (
  address: Address,
  chainId: number,
  poolId: number,
  inputValue: BigNumber
) => {
  const signer = await fetchSigner({
    chainId,
  })
  console.log('signer', address, poolId, inputValue)
  const miniChefContract = new Contract(
    MINICHEF_ADDRESSES[chainId],
    MINI_CHEF_ABI,
    signer
  )
  try {
    if (!address) throw new Error('Wallet must be connected')
    if (!miniChefContract) throw new Error('MMind contract is not loaded')

    const stakeTransaction = await miniChefContract.deposit(
      poolId,
      inputValue,
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
  } catch (err) {
    txErrorHandler(err)
  }
}
