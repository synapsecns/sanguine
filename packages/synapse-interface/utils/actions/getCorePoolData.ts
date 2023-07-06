import { Address, readContracts } from '@wagmi/core'
import { SWAP_ABI } from '@abis/swap'

const list = {
  0: 'initialA',
  1: 'futureA',
  2: 'initialAtTime',
  3: 'futureAtTime',
  4: 'swapFee',
  5: 'adminFee',
  6: 'lpToken',
}

type PoolFeeResponse = {
  result: [bigint, bigint, bigint, bigint, bigint, bigint, bigint]
  success: boolean
}

type VirtualPriceResponse = {
  result: bigint
  success: boolean
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
