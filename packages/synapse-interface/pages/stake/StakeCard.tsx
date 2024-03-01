import { useState, useEffect, useMemo } from 'react'
import { Address } from '@wagmi/core'

import { usePendingTxWrapper } from '@/utils/hooks/usePendingTxWrapper'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'
import { approve, stake } from '@/utils/actions/approveAndStake'
import { useTokenBalance } from '@/utils/hooks/useTokenBalance'
import { withdrawStake } from '@/utils/actions/withdrawStake'
import { getTokenOnChain } from '@/utils/hooks/useTokenInfo'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { claimStake } from '@/utils/actions/claimStake'
import { Token } from '@/utils/types'

import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'
import InteractiveInputRow from '@/components/InteractiveInputRow'
import Button from '@/components/ui/tailwind/Button'

import StakeCardTitle from './StakeCardTitle'
import { formatBigIntToString } from '@/utils/bigint/format'
import { stringToBigInt } from '@/utils/bigint/format'
import InfoSectionCard from '../pool/PoolInfoSection/InfoSectionCard'
import Tabs from '@/components/ui/tailwind/Tabs'
import TabItem from '@/components/ui/tailwind/TabItem'
import { InteractiveInputRowButton } from '@/components/InteractiveInputRowButton'

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

  const balance = useTokenBalance(pool, pool?.chainId)

  const lpTokenBalance = balance?.data ? BigInt(balance?.data?.value) : 0n

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
  const [tx, setTx] = useState(undefined)
  const miniChefAddress = pool.miniChefAddress

  useEffect(() => {
    if (!address || !chainId || stakingPoolId === null) {
      setUserStakeData({
        amount: 0n,
        reward: 0n,
      })
      return
    }
    getStakedBalance(address as Address, pool.chainId, stakingPoolId, pool)
      .then((data) => {
        setUserStakeData(data)
      })
      .catch((err) => {
        console.log(err)
      })
  }, [address, chainId, stakingPoolId])

  useEffect(() => {
    if (!address) {
      setUserStakeData({
        amount: 0n,
        reward: 0n,
      })
      return
    }
    ;(async () => {
      const tkAllowance = await getTokenAllowance(
        miniChefAddress as Address,
        pool.addresses[chainId] as Address,
        address as Address,
        chainId
      )
      setAllowance(tkAllowance)
      getStakedBalance(address as Address, pool.chainId, stakingPoolId, pool)
        .then((data) => {
          setUserStakeData(data)
        })
        .catch((err) => {
          console.log(err)
        })
    })()
  }, [lpTokenBalance])

  useEffect(() => {
    if (!address) return
    ;(async () => {
      const tkAllowance = await getTokenAllowance(
        miniChefAddress as Address,
        pool.addresses[chainId] as Address,
        address as Address,
        chainId
      )
      setAllowance(tkAllowance)
    })()
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
          <div className="text-[#EEEDEF]">Unstaked</div>
          <div className="text-white ">
            {lpTokenBalance === 0n
              ? '\u2212'
              : formatBigIntToString(
                  lpTokenBalance,
                  tokenInfo.decimals,
                  6
                )}{' '}
            <span className="text-base text-[#A9A5AD]">
              {pool ? pool.symbol : ''}
            </span>
          </div>
        </div>
        <div className="flex items-center justify-between my-2">
          <div className="text-[#EEEDEF]">Staked</div>
          <div className="text-white ">
            {formatBigIntToString(userStakeData.amount, tokenInfo.decimals, 6)}{' '}
            <span className="text-base text-[#A9A5AD]">
              {pool ? pool.symbol : ''}
            </span>
          </div>
        </div>
        <div className="flex items-center justify-between my-2">
          <div className="text-[#EEEDEF]">
            {pool?.customRewardToken ?? 'SYN'} Earned
          </div>
          <div className="text-white ">
            {userStakeData.reward === 0n
              ? '\u2212'
              : formatBigIntToString(userStakeData.reward, 18, 6)}{' '}
            <span className="text-base text-[#A9A5AD]">
              {pool?.customRewardToken ?? 'SYN'}
            </span>
          </div>
        </div>
        {userStakeData.reward === 0n ? null : (
          <Button
            disabled={userStakeData.reward === 0n}
            className={`
             bg-[#564f58]
              w-full my-2 px-4 py-3 tracking-wide
              rounded-sm
              border border-transparent 
              hover:border-[#AC8FFF] 
              disabled:opacity-100
              disabled:from-bgLight disabled:to-bgLight
            `}
            onClick={() =>
              pendingTxWrapFunc(
                claimStake(chainId, address as Address, stakingPoolId, pool)
              )
            }
          >
            {isPending ? (
              <div className="flex items-center justify-center space-x-5">
                <ButtonLoadingDots className="mr-3" />
                <span className="animate-pulse">Claiming</span>{' '}
              </div>
            ) : (
              <div className="font-thin">
                Claim {pool.customRewardToken ?? 'SYN'}
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
              Stake
            </TabItem>
            <TabItem
              isActive={!showStake}
              onClick={() => {
                setShowStake(false)
              }}
              className="rounded-tr-sm"
            >
              Unstake
            </TabItem>
          </Tabs>
        </div>
        <div className="p-lg">
          {showStake ? (
            <InteractiveInputRow
              title={pool?.symbol}
              isConnected={Boolean(address)}
              balanceStr={
                lpTokenBalance === 0n
                  ? '0.0'
                  : formatBigIntToString(lpTokenBalance, tokenInfo.decimals, 18)
              }
              onClickBalance={() => {
                setDeposit({
                  str:
                    lpTokenBalance === 0n
                      ? '0.0000'
                      : formatBigIntToString(
                          lpTokenBalance,
                          tokenInfo.decimals
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
              disabled={lpTokenBalance === 0n}
              icon={pool?.icon?.src}
            />
          ) : (
            <InteractiveInputRow
              title={pool?.symbol}
              isConnected={Boolean(address)}
              balanceStr={formatBigIntToString(
                userStakeData.amount,
                tokenInfo.decimals,
                18
              )}
              onClickBalance={() => {
                setWithdraw(
                  userStakeData.amount === 0n
                    ? '0.00000'
                    : formatBigIntToString(
                        userStakeData.amount,
                        tokenInfo.decimals,
                        18
                      )
                )
              }}
              value={withdraw}
              placeholder={'0.0000'}
              onChange={(e) => {
                let val = cleanNumberInput(e.target.value)
                setWithdraw(val)
              }}
              disabled={userStakeData.amount === 0n}
              icon={pool?.icon?.src}
            />
          )}
          {showStake ? (
            <InteractiveInputRowButton
              title={pool?.symbol}
              buttonLabel={
                lpTokenBalance === 0n
                  ? 'Insufficient Balance'
                  : allowance < deposit.bi
                  ? `Approve ${pool?.symbol}`
                  : 'Stake'
              }
              loadingLabel={isPendingApprove ? 'Approving' : 'Staking'}
              disabled={lpTokenBalance === 0n || deposit.str === ''}
              isPending={isPendingStake || isPendingApprove}
              onClickEnter={
                allowance < deposit.bi
                  ? async (e) => {
                      const tx = await pendingApproveTxWrapFunc(
                        approve(pool, deposit.bi, chainId)
                      )

                      setTx(tx?.transactionHash)
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
                      }
                      setTx(tx?.transactionHash)
                    }
              }
            />
          ) : (
            <InteractiveInputRowButton
              title={pool?.symbol}
              buttonLabel="Unstake"
              loadingLabel="Unstaking"
              disabled={userStakeData.amount === 0n || withdraw === ''}
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
                if (tx?.status === 1) {
                  setWithdraw('')
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
