import { readContract } from '@wagmi/core'

import { Token } from '@/utils/types'
import { SYNAPSE_ROUTER_ABI } from '@/constants/abis/synapseRouter'

const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'
const BASE_ROUTER_ADDRESS = '0x6C8c6E68604e78B549C96907bfe9EBdaaC04e3B3'

export const swapPoolCalculateAddLiquidity = async ({
  chainId,
  pool,
  inputs,
}: {
  chainId: number
  pool: Token
  inputs: bigint[]
}) => {
  if (chainId === 8453) {
    const amount = await readContract({
      chainId,
      address: BASE_ROUTER_ADDRESS,
      abi: SYNAPSE_ROUTER_ABI,
      functionName: 'calculateAddLiquidity',
      args: [pool.swapAddresses[chainId], inputs],
    })
    return amount
  } else {
    const amount = await readContract({
      chainId,
      address: ROUTER_ADDRESS,
      abi: SYNAPSE_ROUTER_ABI,
      functionName: 'calculateAddLiquidity',
      args: [pool.swapAddresses[chainId], inputs],
    })
    return amount
  }
}
