import { Address, erc20ABI, waitForTransaction } from '@wagmi/core'
import { TransactionReceipt } from 'viem'
import { MAX_UINT256 } from '@/constants/index'
import { viemPublicClient, viemWalletClient } from 'index'

export const approveErc20Token = async ({
  spenderAddress,
  tokenAddress,
  ownerAddress,
  chainId,
  amount,
}: {
  spenderAddress: Address
  tokenAddress: Address
  ownerAddress: Address
  chainId: number
  amount: bigint
}) => {
  /** @DEV TO-DO: Uncomment once Chains / Tokens constants port over */
  //   let abi
  //   if (tokenAddress === USDT.addresses[Ethereum.id]) {
  //     abi = USDT_ABI
  //   } else {
  //     abi = erc20ABI
  //   }

  const { request } = await viemPublicClient.simulateContract({
    address: tokenAddress,
    account: ownerAddress,
    abi: erc20ABI,
    functionName: 'approve',
    args: [spenderAddress, amount ?? MAX_UINT256],
  })

  const hash = await viemWalletClient.writeContract(request)

  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
