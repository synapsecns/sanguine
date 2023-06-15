import { readContracts } from '@wagmi/core'
import { Zero } from '@ethersproject/constants'
import SWAP_ABI from '@abis/swap.json'
import { bnPercentFormat } from '@bignumber/format'

export const getPoolFee = async (poolAddress: string, chainId: number) => {
  const data: any = await readContracts({
    contracts: [
      {
        address: `0x${poolAddress.slice(2)}`,
        abi: SWAP_ABI,
        functionName: 'swapStorage',
        chainId,
      },
      {
        address: `0x${poolAddress.slice(2)}`,
        abi: SWAP_ABI,
        functionName: 'getVirtualPrice',
        chainId,
      },
    ],
  })
  const swapFeeRaw = data[0]?.swapFee ?? Zero
  const virtualPrice: any = data[1]
  const swapFee = bnPercentFormat(swapFeeRaw)
  return { swapFee, virtualPrice }
}

export const getVirtualPrice = async (poolAddress: string, chainId: number) => {
  const data: any = await readContracts({
    contracts: [
      {
        address: `0x${poolAddress.slice(2)}`,
        abi: SWAP_ABI,
        functionName: 'getVirtualPrice',
        chainId,
      },
    ],
  })
  const virtualPrice: any = data[0]
  return virtualPrice
}
