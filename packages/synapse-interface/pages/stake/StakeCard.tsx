import { useState } from 'react'
import { usePendingTxWrapper } from '@/utils/hooks/usePendingTxWrapper'
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

  const [deposit, setDeposit] = useState('')
  const [withdraw, setWithdraw] = useState('')

  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()
  const [isPendingStake, pendingStakeTxWrapFunc] = usePendingTxWrapper()
  const [isPendingUnstake, pendingUnstakeTxWrapFunc] = usePendingTxWrapper()

  const [showStake, setShowStake] = useState(true)

  return (
    <div className="flex-wrap space-y-2">
      <StakeCardTitle
        token={token}
        poolTokens={stakingPoolTokens}
        poolLabel={stakingPoolLabel}
      />
    </div>
  )
}

export default StakeCard
