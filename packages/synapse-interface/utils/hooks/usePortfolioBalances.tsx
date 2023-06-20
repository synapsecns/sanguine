import { useAccount } from 'wagmi'
import { Address, multicall } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'

//move to constants file later
const MULTICALL3_ADDRESS: Address = '0xcA11bde05977b3631167028862bE2a173976CA11'

export const usePortfolioBalances = () => {}

const useTokenBalances = () => {}

const useTokenApprovals = () => {}
