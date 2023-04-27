import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { Token } from '@/utils/types'

interface StakeCardProps {
  chainId: number
  token: Token
}

const StakeCard = ({ chainId, token }: StakeCardProps) => {
  const tokenInfo = getTokenOnChain(chainId, token)
  console.log('tokenInfo from swap card: ', tokenInfo)

  return <div className="flex-wrap space-y-2">token</div>
}

export default StakeCard
