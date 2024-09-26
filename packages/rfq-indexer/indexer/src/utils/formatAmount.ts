import { formatUnits } from 'viem'

const ETH_ADDRESS = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

export function formatAmount(amount: bigint, tokenAddress: string): string {
  const decimals = tokenAddress.toLowerCase() === ETH_ADDRESS.toLowerCase() ? 18 : 6
  return formatUnits(amount, decimals)
}