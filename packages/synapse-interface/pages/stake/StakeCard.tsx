import { useState } from 'react'
import { usePendingTxWrapper } from '@/utils/hooks/usePendingTxWrapper'
import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { Token } from '@/utils/types'
import Card from '@/components/ui/tailwind/Card'
import InfoSection from '../pool/PoolInfoSection/InfoSection'
import StakeCardTitle from './StakeCardTitle'
import { useStakedBalance } from '@/utils/hooks/useStakedBalance'
import { useClaimStake } from '@/utils/actions/useClaimStake'
import { useApproveAndStake } from '@/utils/actions/useApproveAndStake'
import { useWithdrawStake } from '@/utils/actions/useWithdrawStake'
import { useAccount, useNetwork } from 'wagmi'
import { commifyBnToString } from '@/utils/bignumber/format'
import { useTokenBalance } from '@/utils/hooks/useTokenBalance'
import Button from '@/components/ui/tailwind/Button'
import ButtonLoadingSpinner from '@/components/buttons/ButtonLoadingSpinner'
import InteractiveInputRow from '@/components/InteractiveInputRow'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { smartParseUnits } from '@/utils/bignumber'
import { formatUnits } from '@ethersproject/units'
import { Zero } from '@ethersproject/constants'

interface StakeCardProps {
  chainId: number
  token: Token
}

const StakeCard = ({ chainId, token }: StakeCardProps) => {
  const tokenInfo = getTokenOnChain(chainId, token)
  const stakingPoolLabel: string = tokenInfo?.poolName
  const stakingPoolTokens: Token[] = tokenInfo?.poolTokens
  const stakingPoolId: number = tokenInfo?.poolId

  const { data } = useTokenBalance(token)
  const lpTokenBalance = data?.value ?? Zero
  const { chain } = useNetwork()
  const { address } = useAccount()
  const { amount, reward } = useStakedBalance({ poolId: stakingPoolId })
  const claimStake = useClaimStake()
  const approveAndStake = useApproveAndStake(token)
  const withdrawStake = useWithdrawStake()

  const [deposit, setDeposit] = useState('')
  const [withdraw, setWithdraw] = useState('')
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()
  const [isPendingStake, pendingStakeTxWrapFunc] = usePendingTxWrapper()
  const [isPendingUnstake, pendingUnstakeTxWrapFunc] = usePendingTxWrapper()
  const [showStake, setShowStake] = useState(true)

  return (
    <div className="flex-wrap space-y-2">
      <StakeCardTitle
        address={address}
        connectedChainId={chain.id}
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
                {commifyBnToString(lpTokenBalance, 2)}{' '}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>Staked</div>
              <div>
                {commifyBnToString(amount, 2)}{' '}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>SYN Earned</div>
              <div className="text-green-400">
                {commifyBnToString(reward, 2)}{' '}
                <span className="text-[#88818C]">SYN</span>
              </div>
            </div>
          </div>
        </InfoSection>
      </Card>
      {reward.eq(0) ? null : (
        <Button
          disabled={reward.eq(0)}
          className={`
          w-full  my-2 px-4 py-3 tracking-wide
          hover:opacity-80 disabled:opacity-100
          disabled:from-bgLight disabled:to-bgLight
          bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
          ${isPending && 'from-[#622e71] to-[#564071]'}
        `}
          onClick={async () => {
            pendingTxWrapFunc(claimStake({ poolId: stakingPoolId }))
          }}
        >
          {isPending ? (
            <>
              <ButtonLoadingSpinner className="mr-2" />
              <span className="animate-pulse">Claiming SYN</span>{' '}
            </>
          ) : (
            <>Claim SYN</>
          )}
        </Button>
      )}
      <Card className="bg-bgBase rounded-xl" divider={false}>
        <div className="flex justify-center space-x-4">
          <Button
            className={`${
              showStake ? 'bg-[#111111]' : 'bg-bgLight hover:bg-opacity-70'
            }  w-full rounded-lg h-[48px] text-white text-xl`}
            onClick={() => setShowStake(true)}
          >
            Stake
          </Button>
          <Button
            className={`${
              !showStake ? 'bg-[#111111]' : 'bg-bgLight hover:bg-opacity-70'
            }  w-full rounded-lg h-[48px] text-white text-xl`}
            onClick={() => setShowStake(false)}
          >
            Unstake
          </Button>
        </div>
        {showStake ? (
          <InteractiveInputRow
            showButton={true}
            title={token.symbol}
            buttonLabel="Stake"
            buttonWidth="w-full"
            loadingLabel="Staking"
            isConnected={Boolean(address)}
            balanceStr={commifyBnToString(lpTokenBalance, 2)}
            onClickBalance={() => {
              setDeposit(formatUnits(lpTokenBalance, 18))
            }}
            value={deposit}
            placeholder={'0.0'}
            onChange={(e) => {
              let val = cleanNumberInput(e.target.value)
              setDeposit(val)
            }}
            disabled={lpTokenBalance.eq(0) || deposit == ''}
            isPending={isPendingStake}
            onClickEnter={async (e) => {
              const tx = await pendingStakeTxWrapFunc(
                approveAndStake({
                  poolId: stakingPoolId,
                  infiniteApproval: true,
                  amount: smartParseUnits(deposit, 18),
                })
              )
              if (tx?.status === 1) {
                setDeposit('')
              }
            }}
            token={token}
            icon={token.icon.src}
          />
        ) : (
          <InteractiveInputRow
            showButton={true}
            title={token.symbol}
            buttonLabel="Unstake"
            buttonWidth="w-full"
            loadingLabel="Unstaking"
            isConnected={Boolean(address)}
            balanceStr={commifyBnToString(amount, 4)}
            onClickBalance={() => {
              setWithdraw(formatUnits(amount, 18))
            }}
            value={withdraw}
            placeholder={'0.0'}
            onChange={(e) => {
              let val = cleanNumberInput(e.target.value)
              setWithdraw(val)
            }}
            disabled={amount.eq(0) || withdraw == ''}
            isPending={isPendingUnstake}
            onClickEnter={async () => {
              const tx = await pendingUnstakeTxWrapFunc(
                withdrawStake({
                  poolId: stakingPoolId,
                  amount: smartParseUnits(withdraw, 18),
                  account: address,
                })
              )
              if (tx?.status === 1) {
                setWithdraw('')
              }
            }}
            token={token}
            icon={token.icon.src}
          />
        )}
      </Card>
    </div>
  )
}

export default StakeCard
