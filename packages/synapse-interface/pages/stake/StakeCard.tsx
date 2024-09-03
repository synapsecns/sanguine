import { useState, useEffect } from 'react'
import { type Address } from 'viem'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import {
  fetchAndStoreSingleNetworkPortfolioBalances,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import { usePendingTxWrapper } from '@/utils/hooks/usePendingTxWrapper'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'
import { approve, stake } from '@/utils/actions/approveAndStake'
import { withdrawStake } from '@/utils/actions/withdrawStake'
import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { claimStake } from '@/utils/actions/claimStake'
import { formatBigIntToString } from '@/utils/bigint/format'
import { stringToBigInt } from '@/utils/bigint/format'
import { Token } from '@/utils/types'
import { InteractiveInputRowButton } from '@/components/InteractiveInputRowButton'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'
import InteractiveInputRow from '@/components/InteractiveInputRow'
import Button from '@/components/ui/tailwind/Button'
import TabItem from '@/components/ui/tailwind/TabItem'
import Tabs from '@/components/ui/tailwind/Tabs'
import InfoSectionCard from '../pool/PoolInfoSection/InfoSectionCard'
import StakeCardTitle from './StakeCardTitle'

interface StakeCardProps {
  address: string
  chainId: number
  pool: Token
}

const StakeCard = ({ address, chainId, pool }: StakeCardProps) => {
  const dispatch = useAppDispatch()
  const tokenInfo = getTokenOnChain(chainId, pool)
  const stakingPoolLabel: string = tokenInfo?.poolName
  const stakingPoolTokens: Token[] = tokenInfo?.poolTokens
  const stakingPoolId: number = tokenInfo?.poolId

  const t = useTranslations('Pools')

  const { poolTokenBalances } = usePortfolioState()

  const lpTokenBalance = poolTokenBalances[chainId]?.find(
    (tokenBalance) => tokenBalance.tokenAddress === pool.addresses[chainId]
  )?.balance

  const [deposit, setDeposit] = useState({ str: '', bi: 0n })
  const [withdraw, setWithdraw] = useState<string>('')
  const [showStake, setShowStake] = useState<boolean>(true)
  const [allowance, setAllowance] = useState<bigint>(0n)
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()
  const [isPendingStake, pendingStakeTxWrapFunc] = usePendingTxWrapper()
  const [isPendingUnstake, pendingUnstakeTxWrapFunc] = usePendingTxWrapper()
  const [isPendingApprove, pendingApproveTxWrapFunc] = usePendingTxWrapper()
  const [userStakeData, setUserStakeData] = useState({
    amount: 0n,
    reward: 0n,
  })
  const miniChefAddress = pool.miniChefAddress

  const resetUserStakeData = () => {
    setUserStakeData({
      amount: 0n,
      reward: 0n,
    })
  }

  const getUserStakedBalance = async (
    address: Address,
    stakingPoolId: number,
    pool: Token
  ) => {
    try {
      const data = await getStakedBalance(
        address as Address,
        pool.chainId,
        stakingPoolId,
        pool
      )
      setUserStakeData(data)
    } catch (err) {
      console.error('Error fetching user staked balance:', err)
    }
  }

  const getUserLpTokenAllowance = async (
    address: Address,
    chainId: number,
    pool: Token
  ) => {
    try {
      const tkAllowance = await getTokenAllowance(
        miniChefAddress as Address,
        pool.addresses[chainId] as Address,
        address as Address,
        chainId
      )
      setAllowance(tkAllowance)
    } catch (err) {
      console.error('Error fetching user LP token allowance:', err)
    }
  }

  useEffect(() => {
    if (!address || !chainId || stakingPoolId === null) {
      resetUserStakeData()
      return
    }
    getUserStakedBalance(address as Address, stakingPoolId, pool)
  }, [address, chainId, stakingPoolId])

  useEffect(() => {
    if (!address) {
      resetUserStakeData()
      return
    }
    getUserLpTokenAllowance(address as Address, chainId, pool)
    getUserStakedBalance(address as Address, stakingPoolId, pool)
  }, [lpTokenBalance])

  useEffect(() => {
    if (!address) return
    getUserLpTokenAllowance(address as Address, chainId, pool)
  }, [deposit])

  return (
    <div className="flex-wrap space-y-2">
      <StakeCardTitle
        token={pool}
        poolTokens={stakingPoolTokens}
        poolLabel={stakingPoolLabel}
        lpTokenBalance={lpTokenBalance}
      />
      <InfoSectionCard title="Your balances">
        <div className="flex items-center justify-between my-2">
          <div className="text-[#EEEDEF]">{t('Unstaked')}</div>
          <div className="text-white ">
            {!lpTokenBalance
              ? '\u2212'
              : trimTrailingZeroesAfterDecimal(
                  formatBigIntToString(lpTokenBalance, tokenInfo.decimals, 6)
                )}{' '}
            <span className="text-base text-[#A9A5AD]">
              {pool ? pool.symbol : ''}
            </span>
          </div>
        </div>
        <div className="flex items-center justify-between my-2">
          <div className="text-[#EEEDEF]">{t('Staked')}</div>
          <div className="text-white ">
            {trimTrailingZeroesAfterDecimal(
              formatBigIntToString(userStakeData.amount, tokenInfo.decimals, 6)
            )}{' '}
            <span className="text-base text-[#A9A5AD]">
              {pool ? pool.symbol : ''}
            </span>
          </div>
        </div>
        <div className="flex items-center justify-between my-2">
          <div className="text-[#EEEDEF]">
            {pool?.customRewardToken ?? 'SYN'} {t('Earned')}
          </div>
          <div className="text-white ">
            {!userStakeData.reward
              ? '\u2212'
              : trimTrailingZeroesAfterDecimal(
                  formatBigIntToString(userStakeData.reward, 18, 6)
                )}{' '}
            <span className="text-base text-[#A9A5AD]">
              {pool?.customRewardToken ?? 'SYN'}
            </span>
          </div>
        </div>
        {!userStakeData.reward ? null : (
          <Button
            disabled={!userStakeData.reward}
            className={`
             bg-[#564f58]
              w-full my-2 px-4 py-3 tracking-wide
              rounded-sm
              border border-transparent
              hover:border-[#AC8FFF]
              disabled:opacity-100
              disabled:from-bgLight disabled:to-bgLight
            `}
            onClick={async (e) => {
              const tx = await pendingTxWrapFunc(
                claimStake(chainId, address as Address, stakingPoolId, pool)
              )
              if (tx?.status === 'success') {
                await getUserStakedBalance(
                  address as Address,
                  stakingPoolId,
                  pool
                )
              }
            }}
          >
            {isPending ? (
              <div className="flex items-center justify-center space-x-5">
                <ButtonLoadingDots className="mr-3" />
                <span className="animate-pulse">{t('Claiming')}</span>{' '}
              </div>
            ) : (
              <div className="font-thin">
                {t('Claim')} {pool.customRewardToken ?? 'SYN'}
              </div>
            )}
          </Button>
        )}
      </InfoSectionCard>
      <div className="p-0 rounded-md bg-bgBase">
        <div className="mb-3">
          <Tabs>
            <TabItem
              isActive={showStake}
              onClick={() => {
                setShowStake(true)
              }}
              className="rounded-tl-sm"
            >
              {t('Stake')}
            </TabItem>
            <TabItem
              isActive={!showStake}
              onClick={() => {
                setShowStake(false)
              }}
              className="rounded-tr-sm"
            >
              {t('Unstake')}
            </TabItem>
          </Tabs>
        </div>
        <div className="p-lg">
          {showStake ? (
            <InteractiveInputRow
              title={pool?.symbol}
              isConnected={Boolean(address)}
              balanceStr={
                !lpTokenBalance
                  ? '0.0'
                  : trimTrailingZeroesAfterDecimal(
                      formatBigIntToString(
                        lpTokenBalance,
                        tokenInfo.decimals,
                        18
                      )
                    )
              }
              onClickBalance={() => {
                setDeposit({
                  str: !lpTokenBalance
                    ? '0.0000'
                    : trimTrailingZeroesAfterDecimal(
                        formatBigIntToString(lpTokenBalance, tokenInfo.decimals)
                      ),
                  bi: lpTokenBalance,
                })
              }}
              value={deposit.str}
              placeholder={'0.0000'}
              onChange={async (e) => {
                let val = cleanNumberInput(e.target.value)
                setDeposit({
                  str: val,
                  bi: stringToBigInt(val, pool.decimals[chainId]),
                })
              }}
              disabled={!lpTokenBalance}
              icon={pool?.icon?.src}
            />
          ) : (
            <InteractiveInputRow
              title={pool?.symbol}
              isConnected={Boolean(address)}
              balanceStr={trimTrailingZeroesAfterDecimal(
                formatBigIntToString(
                  userStakeData.amount,
                  tokenInfo.decimals,
                  18
                )
              )}
              onClickBalance={() => {
                setWithdraw(
                  !userStakeData.amount
                    ? '0.00000'
                    : trimTrailingZeroesAfterDecimal(
                        formatBigIntToString(
                          userStakeData.amount,
                          tokenInfo.decimals,
                          18
                        )
                      )
                )
              }}
              value={withdraw}
              placeholder={'0.0000'}
              onChange={(e) => {
                let val = cleanNumberInput(e.target.value)
                setWithdraw(val)
              }}
              disabled={!userStakeData.amount}
              icon={pool?.icon?.src}
            />
          )}
          {showStake ? (
            <InteractiveInputRowButton
              title={pool?.symbol}
              buttonLabel={
                !lpTokenBalance || lpTokenBalance < deposit.bi
                  ? t('Insufficient balance')
                  : allowance < deposit.bi
                  ? `${t('Approve')} ${pool?.symbol}`
                  : t('Stake')
              }
              loadingLabel={isPendingApprove ? t('Approving') : t('Staking')}
              disabled={
                !lpTokenBalance ||
                lpTokenBalance < deposit.bi ||
                deposit.str === ''
              }
              isPending={isPendingStake || isPendingApprove}
              onClickEnter={
                allowance < deposit.bi
                  ? async (e) => {
                      const tx = await pendingApproveTxWrapFunc(
                        approve(pool, deposit.bi, chainId)
                      )
                      if (tx?.status === 'success') {
                        getUserLpTokenAllowance(
                          address as Address,
                          chainId,
                          pool
                        )
                      }
                    }
                  : async (e) => {
                      const tx = await pendingStakeTxWrapFunc(
                        stake(
                          address as Address,
                          chainId,
                          stakingPoolId,
                          pool,
                          deposit.bi
                        )
                      )
                      if (tx?.status === 'success') {
                        setDeposit({ bi: 0n, str: '' })
                        dispatch(
                          fetchAndStoreSingleNetworkPortfolioBalances({
                            address,
                            chainId,
                          })
                        )
                      }
                    }
              }
            />
          ) : (
            <InteractiveInputRowButton
              title={pool?.symbol}
              buttonLabel={
                userStakeData.amount < stringToBigInt(withdraw, 18)
                  ? t('Insufficient balance')
                  : t('Unstake')
              }
              loadingLabel={t('Unstaking')}
              disabled={
                !userStakeData.amount ||
                userStakeData.amount < stringToBigInt(withdraw, 18) ||
                withdraw === ''
              }
              isPending={isPendingUnstake}
              onClickEnter={async () => {
                const tx = await pendingUnstakeTxWrapFunc(
                  withdrawStake(
                    address as Address,
                    chainId,
                    stakingPoolId,
                    pool,
                    stringToBigInt(withdraw, 18)
                  )
                )

                if (tx?.status === 'success') {
                  setWithdraw('')
                  dispatch(
                    fetchAndStoreSingleNetworkPortfolioBalances({
                      address,
                      chainId,
                    })
                  )
                }
              }}
            />
          )}
        </div>
      </div>
    </div>
  )
}

export default StakeCard
