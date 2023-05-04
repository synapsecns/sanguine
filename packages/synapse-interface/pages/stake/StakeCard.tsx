import { useState } from 'react'
import { usePendingTxWrapper } from '@/utils/hooks/usePendingTxWrapper'
import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { Token } from '@/utils/types'
import Card from '@/components/ui/tailwind/Card'
import InfoSection from '../pool/PoolInfoSection/InfoSection'
import StakeCardTitle from './StakeCardTitle'
import { useStakedBalance } from '@/utils/hooks/useStakedBalance'
import { commifyBnToString } from '@/utils/bignumber/format'
interface StakeCardProps {
  chainId: number
  token: Token
}

const StakeCard = ({ chainId, token }: StakeCardProps) => {
  const tokenInfo = getTokenOnChain(chainId, token)
  const stakingPoolLabel: string = tokenInfo?.poolName
  const stakingPoolTokens: Token[] = tokenInfo?.poolTokens
  const stakingPoolId: number = tokenInfo?.poolId
  const { amount, reward, rawAmount } = useStakedBalance({
    poolId: stakingPoolId,
  })

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
      <Card
        title="Your balances"
        className="p-4 rounded-xl bg-bgBase max-h-40"
        titleClassName="text-base font-base text-secondaryTextColor text-opacity-50"
        divider={false}
      >
        <InfoSection showDivider={true} showOutline={false}>
          {' '}
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>Unstaked</div>
              <div>
                {/* {commifyBnToString(lpTokenBalance, tokenInfo, 2)}{' '} */}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>Staked</div>
              <div>
                {/* {commifyBnToString(amount, 2)}{' '} */}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>SYN Earned</div>
              <div className="text-green-400">
                {/* {commifyBnToString(reward, 2)}{' '} */}
                <span className="text-[#88818C]">SYN</span>
              </div>
            </div>
          </div>
        </InfoSection>
      </Card>
    </div>
  )
}

export default StakeCard
