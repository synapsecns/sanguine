import { Address, readContracts } from '@wagmi/core'
import { Zero } from '@ethersproject/constants'
import { SWAP_ABI } from '@abis/swap'
import { bnPercentFormat } from '@bignumber/format'

const list = {
  0: 'initialA',
  1: 'futureA',
  2: 'initialAtTime',
  3: 'futureAtTime',
  4: 'swapFee',
  5: 'adminFee',
  6: 'lpToken',
}

export const getPoolFee = async (poolAddress: string, chainId: number) => {
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

  console.log(`'think these are now returned as list elements vs named vars?`)
  console.log(`data get poolfee`, data)

  console.log(`swapFee`, data[0].result[4])

  const swapFeeRaw = data[0]?.swapFee ?? Zero
  const virtualPrice: { result: bigint; success: boolean } = data[1]
  const swapFee = bnPercentFormat(swapFeeRaw)
  return { swapFee, virtualPrice }
}

export const getVirtualPrice = async (poolAddress: string, chainId: number) => {
  const data: any = await readContracts({
    contracts: [
      {
        address: poolAddress as Address,
        abi: SWAP_ABI,
        functionName: 'getVirtualPrice',
        chainId,
      },
    ],
  })

  const virtualPrice: any = data[0]
  return virtualPrice
}
