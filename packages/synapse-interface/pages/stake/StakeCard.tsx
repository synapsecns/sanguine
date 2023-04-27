import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { Token } from '@/utils/types'
import StakeCardTitle from './StakeCardTitle'

interface StakeCardProps {
  chainId: number
  token: Token
}

const StakeCard = ({ chainId, token }: StakeCardProps) => {
  const tokenInfo = getTokenOnChain(chainId, token)
  const stakingPoolLabel: string = tokenInfo?.poolName
  const stakingPoolTokens: Token[] = tokenInfo?.poolTokens

  return (
    <div className="flex-wrap space-y-2">
      <StakeCardTitle
        token={token}
        poolTokens={stakingPoolTokens}
        poolLabel={stakingPoolLabel}
      />
      {tokenInfo?.poolName}
    </div>
  )
}

export default StakeCard
