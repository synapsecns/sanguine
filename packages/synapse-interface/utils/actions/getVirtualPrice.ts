import { readContracts } from '@wagmi/core'
import SWAP_ABI from '@abis/swap.json'

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
