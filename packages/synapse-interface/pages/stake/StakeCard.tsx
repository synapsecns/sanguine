import { useState, useEffect } from 'react'
import { formatUnits } from '@ethersproject/units'
import { Zero } from '@ethersproject/constants'
import { Address } from '@wagmi/core'

import { usePendingTxWrapper } from '@/utils/hooks/usePendingTxWrapper'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'
import { approve, stake } from '@/utils/actions/approveAndStake'
import { useTokenBalance } from '@/utils/hooks/useTokenBalance'
import { withdrawStake } from '@/utils/actions/withdrawStake'
import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { commifyBnToString } from '@/utils/bignumber/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { claimStake } from '@/utils/actions/claimStake'
import { usePrices } from '@/utils/actions/getPrices'
import { smartParseUnits } from '@/utils/bignumber'
import { Token } from '@/utils/types'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'

import ButtonLoadingSpinner from '@/components/buttons/ButtonLoadingSpinner'
import InteractiveInputRow from '@/components/InteractiveInputRow'
import LoadingText from '@/components/loading/LoadingText'
import Button from '@/components/ui/tailwind/Button'
import Card from '@/components/ui/tailwind/Card'

import InfoSection from '../pool/PoolInfoSection/InfoSection'
import StakeCardTitle from './StakeCardTitle'

interface StakeCardProps {
  address: string
  chainId: number
  pool: Token
}

const StakeCard = ({ address, chainId, pool }: StakeCardProps) => {
  const tokenInfo = getTokenOnChain(chainId, pool)
  const stakingPoolLabel: string = tokenInfo?.poolName
  const stakingPoolTokens: Token[] = tokenInfo?.poolTokens
  const stakingPoolId: number = tokenInfo?.poolId

  // TODO get rid of this hook
  const balance = useTokenBalance(pool)
  const lpTokenBalance = balance?.data?.value ?? Zero

  const prices = usePrices(chainId)
  const [deposit, setDeposit] = useState({ str: '', bn: Zero })
  const [withdraw, setWithdraw] = useState('')
  const [showStake, setShowStake] = useState(true)
  const [allowance, setAllowance] = useState(Zero)
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()
  const [isPendingStake, pendingStakeTxWrapFunc] = usePendingTxWrapper()
  const [isPendingUnstake, pendingUnstakeTxWrapFunc] = usePendingTxWrapper()
  const [isPendingApprove, pendingApproveTxWrapFunc] = usePendingTxWrapper()
  const [userStakeData, setUserStakeData] = useState({
    amount: Zero,
    reward: Zero,
  })
  const [tx, setTx] = useState(undefined)

  useEffect(() => {
    if (!address || !chainId || stakingPoolId === null) return
    getStakedBalance(address as Address, chainId, stakingPoolId)
      .then((data) => {
        setUserStakeData(data)
      })
      .catch((err) => {
        console.log(err)
      })
  }, [address, chainId, stakingPoolId])

  useEffect(() => {
    if (tx !== undefined) {
      ;(async () => {
        const tkAllowance = await getTokenAllowance(
          MINICHEF_ADDRESSES[chainId],
          pool.addresses[chainId],
          address,
          chainId
        )
        setAllowance(tkAllowance)
      })()
    }
  }, [tx])

  return (
    <div className="flex-wrap space-y-2">
      <StakeCardTitle
        address={address}
        connectedChainId={chainId}
        token={pool}
        poolTokens={stakingPoolTokens}
        poolLabel={stakingPoolLabel}
        prices={prices}
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
                {commifyBnToString(userStakeData.amount, 4)}{' '}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>SYN Earned</div>
              <div className="text-green-400">
                {commifyBnToString(userStakeData.reward, 6)}{' '}
                <span className="text-[#88818C]">SYN</span>
              </div>
            </div>
          </div>
        </InfoSection>
      </Card>
      {userStakeData.reward.eq(0) ? null : (
        <Button
          disabled={userStakeData.reward.eq(0)}
          className={`
          w-full  my-2 px-4 py-3 tracking-wide
          hover:opacity-80 disabled:opacity-100
          disabled:from-bgLight disabled:to-bgLight
          bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
          ${isPending && 'from-[#622e71] to-[#564071]'}
        `}
          onClick={() =>
            pendingTxWrapFunc(claimStake(chainId, address, stakingPoolId))
          }
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
            title={pool?.symbol}
            buttonLabel={
              lpTokenBalance.eq(0)
                ? 'Insufficient Balance'
                : allowance.lt(deposit.bn)
                ? `Approve ${pool?.symbol}`
                : 'Stake'
            }
            buttonWidth="w-full"
            loadingLabel={isPendingApprove ? 'Approving' : 'Staking'}
            isConnected={Boolean(address)}
            balanceStr={commifyBnToString(lpTokenBalance, 2)}
            onClickBalance={() => {
              setDeposit({
                str: formatUnits(lpTokenBalance, 18),
                bn: lpTokenBalance,
              })
            }}
            value={deposit.str}
            placeholder={'0.0'}
            onChange={async (e) => {
              let val = cleanNumberInput(e.target.value)
              const tkAllowance = await getTokenAllowance(
                MINICHEF_ADDRESSES[chainId],
                pool.addresses[chainId],
                address,
                chainId
              )
              setAllowance(tkAllowance)
              setDeposit({
                str: val,
                bn: smartParseUnits(val, pool.decimals[chainId]),
              })
            }}
            disabled={lpTokenBalance.eq(0) || deposit.str == ''}
            isPending={isPendingStake || isPendingApprove}
            onClickEnter={
              allowance.lt(deposit.bn)
                ? async (e) => {
                    const tx = await pendingApproveTxWrapFunc(
                      approve(pool, deposit.bn, chainId)
                    )

                    setTx(tx?.hash)
                  }
                : async (e) => {
                    const tx = await pendingStakeTxWrapFunc(
                      stake(
                        `0x${address.slice(2)}`,
                        chainId,
                        stakingPoolId,
                        deposit.bn
                      )
                    )
                    if (tx?.status === 1) {
                      setDeposit({ bn: Zero, str: '' })
                    }
                  }
            }
            token={pool}
            icon={pool?.icon?.src}
          />
        ) : (
          <InteractiveInputRow
            showButton={true}
            title={pool?.symbol}
            buttonLabel="Unstake"
            buttonWidth="w-full"
            loadingLabel="Unstaking"
            isConnected={Boolean(address)}
            balanceStr={commifyBnToString(userStakeData.amount, 4)}
            onClickBalance={() => {
              setWithdraw(formatUnits(userStakeData.amount, 18))
            }}
            value={withdraw}
            placeholder={'0.0'}
            onChange={(e) => {
              let val = cleanNumberInput(e.target.value)
              setWithdraw(val)
            }}
            disabled={userStakeData.amount.eq(0) || withdraw == ''}
            isPending={isPendingUnstake}
            onClickEnter={async () => {
              const tx = await pendingUnstakeTxWrapFunc(
                withdrawStake(
                  `0x${address.slice(2)}`,
                  chainId,
                  stakingPoolId,
                  smartParseUnits(withdraw, 18)
                )
              )
              if (tx?.status === 1) {
                setWithdraw('')
              }
            }}
            token={pool}
            icon={pool?.icon?.src}
          />
        )}
      </Card>
    </div>
  )
}

export default StakeCard
