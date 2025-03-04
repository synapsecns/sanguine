import toast from 'react-hot-toast'
import {
  type SimulateContractParameters,
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { Address, erc20Abi, ContractFunctionExecutionError, Abi } from 'viem'

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

  let txReceipt

  try {
    txReceipt = await _submitApproval({
      chainId,
      address: tokenAddress,
      abi,
      amount,
      spender,
    })
  } catch (error) {
    // Check if the error was caused by the approve simulation being reverted.
    if (
      error instanceof ContractFunctionExecutionError &&
      error.message.includes('revert')
    ) {
      // TODO: translation
      const msg = 'Resetting allowance to zero first'
      const pendingPopup = toast(msg, {
        id: 'reset-allowance-in-progress-popup',
        duration: Infinity,
      })
      await _submitApproval({
        chainId,
        address: tokenAddress,
        abi,
        spender,
        amount: 0n,
      })
      toast.dismiss(pendingPopup)
      txReceipt = await _submitApproval({
        chainId,
        address: tokenAddress,
        abi,
        amount,
        spender,
      })
    } else {
      // Unrelated error, rethrow.
      throw error
    }
  }

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
