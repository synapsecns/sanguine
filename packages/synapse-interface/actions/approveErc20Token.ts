import {
  Address,
  erc20ABI,
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

import { MAX_UINT256 } from '@/constants'
import { USDT } from '@/constants/tokens/bridgeable'
import { USDT_ABI } from '@/constants/abis/usdtAbi'
import { ETH as Ethereum } from '@/constants/chains/master'

export const approveErc20Token = async ({
  chainId,
  amount,
  spender,
  tokenAddress,
}: {
  chainId: number
  amount: bigint
  spender: Address
  tokenAddress: Address
}) => {
  let abi

  if (tokenAddress === USDT.addresses[Ethereum.id]) {
    abi = USDT_ABI
  } else {
    abi = erc20ABI
  }

  const { request } = await prepareWriteContract({
    chainId,
    address: tokenAddress,
    abi,
    functionName: 'approve',
    args: [spender, amount ?? MAX_UINT256],
  })

  const { hash } = await writeContract(request)

  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
