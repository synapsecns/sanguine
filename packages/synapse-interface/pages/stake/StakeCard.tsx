import { useState, useEffect } from 'react'
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
import { usePrices } from '@/utils/actions/getPrices'
import { Token } from '@/utils/types'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'

import ButtonLoadingSpinner from '@/components/buttons/ButtonLoadingSpinner'
import InteractiveInputRow from '@/components/InteractiveInputRow'
import LoadingText from '@/components/loading/LoadingText'
import Button from '@/components/ui/tailwind/Button'
import Card from '@/components/ui/tailwind/Card'

import InfoSection from '../pool/PoolInfoSection/InfoSection'
import StakeCardTitle from './StakeCardTitle'
import { formatBigIntToString } from '@/utils/bigint/format'
import { stringToBigInt } from '@/utils/stringToBigNum'

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

  const lpTokenBalance = balance?.data ? BigInt(balance?.data?.value) : 0n

  const prices = usePrices(chainId)
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
          pool.addresses[chainId] as Address,
          address as Address,
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
                {lpTokenBalance === 0n
                  ? '\u2212'
                  : formatBigIntToString(
                      lpTokenBalance,
                      tokenInfo.decimals,
                      4
                    )}{' '}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>Staked</div>
              <div>
                {formatBigIntToString(
                  userStakeData.amount,
                  tokenInfo.decimals,
                  4
                )}{' '}
                <span className="text-[#88818C]">LP</span>
              </div>
            </div>
          </div>
          <div>
            <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
              <div>SYN Earned</div>
              <div className="text-green-400">
                {formatBigIntToString(userStakeData.reward, 18, 8)}{' '}
                <span className="text-[#88818C]">SYN</span>
              </div>
            </div>
          </div>
        </InfoSection>
      </Card>
      {userStakeData.reward === BigInt(0) ? null : (
        <Button
          disabled={userStakeData.reward === 0n}
          className={`
            w-full  my-2 px-4 py-3 tracking-wide
            hover:opacity-80 disabled:opacity-100
            disabled:from-bgLight disabled:to-bgLight
            bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
            ${isPending && 'from-[#622e71] to-[#564071]'}
          `}
          onClick={() =>
            pendingTxWrapFunc(
              claimStake(chainId, address as Address, stakingPoolId)
            )
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
              lpTokenBalance === 0n
                ? 'Insufficient Balance'
                : allowance < deposit.bi
                ? `Approve ${pool?.symbol}`
                : 'Stake'
            }
            buttonWidth="w-full"
            loadingLabel={isPendingApprove ? 'Approving' : 'Staking'}
            isConnected={Boolean(address)}
            balanceStr={formatBigIntToString(
              lpTokenBalance,
              tokenInfo.decimals,
              4
            )}
            onClickBalance={() => {
              setDeposit({
                str: lpTokenBalance.toString(),
                bi: lpTokenBalance,
              })
            }}
            value={formatBigIntToString(deposit.bi, tokenInfo.decimals)}
            placeholder={'0.0'}
            onChange={async (e) => {
              let val = cleanNumberInput(e.target.value)
              const tkAllowance = await getTokenAllowance(
                MINICHEF_ADDRESSES[chainId],
                pool.addresses[chainId] as Address,
                address as Address,
                chainId
              )
              setAllowance(tkAllowance)
              setDeposit({
                str: val,
                bi: stringToBigInt(val, pool.decimals[chainId]),
              })
            }}
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
                        deposit.bi
                      )
                    )
                    if (tx?.status === 1) {
                      setDeposit({ bi: 0n, str: '' })
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
            balanceStr={formatBigIntToString(
              userStakeData.amount,
              tokenInfo.decimals,
              4
            )}
            onClickBalance={() => {
              setWithdraw(
                formatBigIntToString(
                  userStakeData.amount,
                  tokenInfo.decimals,
                  4
                )
              )
            }}
            value={withdraw}
            placeholder={'0.0'}
            onChange={(e) => {
              let val = cleanNumberInput(e.target.value)
              setWithdraw(val)
            }}
            disabled={userStakeData.amount === 0n || withdraw === ''}
            isPending={isPendingUnstake}
            onClickEnter={async () => {
              const tx = await pendingUnstakeTxWrapFunc(
                withdrawStake(
                  address as Address,
                  chainId,
                  stakingPoolId,
                  stringToBigInt(withdraw, 18)
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
