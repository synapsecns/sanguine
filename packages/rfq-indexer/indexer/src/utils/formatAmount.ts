import { formatUnits } from 'viem'

const ADDRESSES_WITH_18_DECIMALS = [
  '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE', // ETH
  '0x2cFc85d8E48F8EAB294be644d9E25C3030863003', // WLD
  '0x163f8c2467924be0ae7b5347228cabf260318753', // WLD
].map(address => address.toLowerCase())

export function formatAmount(amount: bigint, tokenAddress: string): string {
  const decimals = ADDRESSES_WITH_18_DECIMALS.includes(tokenAddress.toLowerCase()) ? 18 : 6
  return formatUnits(amount, decimals)
}
