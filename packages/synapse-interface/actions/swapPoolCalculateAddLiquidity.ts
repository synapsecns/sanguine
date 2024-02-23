import { readContract } from '@wagmi/core'

import type { Token } from '@/utils/types'
import { SYNAPSE_ROUTER_ABI } from '@/constants/abis/synapseRouter'

const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'

export const swapPoolCalculateAddLiquidity = async ({
  chainId,
  pool,
  inputs,
}: {
  chainId: number
  pool: Token
  inputs: bigint[]
}) => {
  const amount = await readContract({
    chainId,
    address: ROUTER_ADDRESS,
    abi: SYNAPSE_ROUTER_ABI,
    functionName: 'calculateAddLiquidity',
    args: [pool.swapAddresses[chainId], inputs],
  })
  return amount
}
