import { formatUnits } from 'viem'

const ADDRESSES_WITH_18_DECIMALS = [
  '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE', // native chain gas (eg: ETH)
  '0x2cFc85d8E48F8EAB294be644d9E25C3030863003', // WLD
  '0x163f8c2467924be0ae7b5347228cabf260318753', // WLD
  '0xdC6fF44d5d932Cbd77B52E5612Ba0529DC6226F1', // WLD
  '0x2f6f07cdcf3588944bf4c42ac74ff24bf56e7590', // Berachain WETH
  '0x2170Ed0880ac9A755fd29B2688956BD959F933F8', // BSC ETH
].map((address) => address.toLowerCase())

export const formatAmount = (amount: bigint, tokenAddress: string): string => {
  const decimals = ADDRESSES_WITH_18_DECIMALS.includes(
    tokenAddress.toLowerCase()
  )
    ? 18
    : 6
  return formatUnits(amount, decimals)
}
