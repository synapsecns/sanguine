import {
  type SimulateContractParameters,
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { Address, erc20Abi, Abi } from 'viem'

import { MAX_UINT256 } from '@/constants'
import { USDT } from '@/constants/tokens/bridgeable'
import { USDT_ABI } from '@/constants/abis/usdtAbi'
import { ETH as Ethereum } from '@/constants/chains/master'
import { wagmiConfig } from '@/wagmiConfig'

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
    abi = erc20Abi
  }
  amount = amount ?? MAX_UINT256
  const txReceipt = await _submitApproval({
    chainId,
    address: tokenAddress,
    abi,
    amount,
    spender,
  })

  return txReceipt
}

const _submitApproval = async ({
  chainId,
  address,
  abi,
  spender,
  amount,
}: {
  chainId: number
  address: Address
  abi: Abi
  spender: Address
  amount: bigint
}) => {
  const { request } = await simulateContract(wagmiConfig, {
    chainId,
    address,
    abi,
    functionName: 'approve',
    args: [spender, amount],
  } as SimulateContractParameters)

  const hash = await writeContract(wagmiConfig, request)

  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
