import { Address, readContracts } from '@wagmi/core'
import { SWAP_ABI } from '@/constants/abis/swap'

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
  const data: any = await readContracts({
    contracts: [
      {
        address: poolAddress as Address,
        abi: SWAP_ABI,
        functionName: 'swapStorage',
        chainId,
      },
      {
        address: poolAddress as Address,
        abi: SWAP_ABI,
        functionName: 'getVirtualPrice',
        chainId,
      },
    ],
  })

  const poolFeeData: PoolFeeResponse = data[0]
  const swapFee = poolFeeData?.result[4] ?? 0n

  const virtualPriceResponse: VirtualPriceResponse = data[1]
  const virtualPrice = virtualPriceResponse.result

  return { swapFee, virtualPrice }
}
