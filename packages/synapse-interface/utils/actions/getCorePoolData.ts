// @ts-nocheck
import { readContracts } from '@wagmi/core'
import { type Address } from 'viem'

import { SWAP_ABI } from '@/constants/abis/swap'
import { wagmiConfig } from '@/wagmiConfig'

type PoolFeeResponse = {
  result: [
    initialA: bigint,
    futureA: bigint,
    initialAtTime: bigint,
    futureAtTime: bigint,
    swapFee: bigint,
    adminFee: bigint,
    lpToken: Address
  ]
  status: string
}

type VirtualPriceResponse = {
  result: bigint
  status: string
}

export const getCorePoolData = async (poolAddress: string, chainId: number) => {
  const data: any = await readContracts(wagmiConfig, {
    contracts: [
      {
        address: poolAddress as Address,
        abi: SWAP_ABI,
        functionName: 'swapStorage',
        chainId: chainId as any,
      },
      {
        address: poolAddress as Address,
        abi: SWAP_ABI,
        functionName: 'getVirtualPrice',
        chainId: chainId as any,
      },
    ],
  })

  const poolFeeData: PoolFeeResponse = data[0]
  const swapFee = poolFeeData?.result[4] ?? 0n

  const virtualPriceResponse: VirtualPriceResponse = data[1]
  const virtualPrice = virtualPriceResponse.result

  return { swapFee, virtualPrice }
}
