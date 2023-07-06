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

type VirtualPriceResponse = {
  result: bigint
  success: boolean
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
  console.log(`not working on arbitrumethpool`)

  const swapFee = data[0]?.result[4] ?? 0n
  console.log(`swapFee`, swapFee)
  const virtualPrice: VirtualPriceResponse = data[1].result
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

  const virtualPrice: VirtualPriceResponse = data[0]
  return virtualPrice.result
}
